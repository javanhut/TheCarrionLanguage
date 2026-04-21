// Package update implements the `carrion update` subcommand. It fetches the
// latest release (or latest main-branch commit with --experimental) from
// GitHub, compares against the running binary's baked-in version metadata,
// prompts the user, and swaps the on-disk binary in place.
package update

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/javanhut/TheCarrionLanguage/src/version"
)

// Run parses update-subcommand args and executes the update flow. It returns
// an error (suitable for printing) on any failure; the caller handles exit.
// args is everything after the "update" keyword, e.g. ["--experimental"].
func Run(args []string) error {
	fs := flag.NewFlagSet("update", flag.ContinueOnError)
	fs.SetOutput(os.Stderr)
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: carrion update [flags]")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Updates the carrion binary in place.")
		fmt.Fprintln(os.Stderr, "Default channel is stable releases (patch/minor/major).")
		fmt.Fprintln(os.Stderr, "")
		fmt.Fprintln(os.Stderr, "Flags:")
		fs.PrintDefaults()
	}
	experimental := fs.Bool("experimental", false, "Track latest commit on main (builds from source, requires go)")
	check := fs.Bool("check", false, "Check for updates and print the result; do not install")
	yes := fs.Bool("yes", false, "Skip confirmation prompt")
	yesShort := fs.Bool("y", false, "Skip confirmation prompt (short form)")
	verbose := fs.Bool("verbose", false, "Print extra diagnostic information")
	if err := fs.Parse(args); err != nil {
		return err
	}

	opts := options{
		experimental: *experimental,
		check:        *check,
		yes:          *yes || *yesShort,
		verbose:      *verbose,
	}

	if opts.experimental {
		return runExperimental(opts)
	}
	return runStable(opts)
}

type options struct {
	experimental bool
	check        bool
	yes          bool
	verbose      bool
}

// checkAvailable tracks whether `--check` mode found an available update. The
// caller exposes this via CheckExitCode() so `carrion update --check` can
// distinguish "up to date" (exit 0) from "update available" (exit 1) in CI
// scripts.
var checkAvailable bool

func setCheckAvailable() { checkAvailable = true }

// CheckExitCode returns the exit code the main binary should use after Run()
// returns nil. Non-zero only when --check was used and an update was found.
func CheckExitCode() int {
	if checkAvailable {
		return 1
	}
	return 0
}

// runStable resolves the latest GitHub release, compares semver against the
// current binary, and (if newer) downloads the release asset or falls back to
// a source build at the release tag.
func runStable(opts options) error {
	fmt.Fprintf(os.Stderr, "Current version: %s\n", version.Full())

	rel, err := fetchLatestRelease()
	if err != nil {
		return fmt.Errorf("fetch latest release: %w", err)
	}
	if rel.Draft {
		return fmt.Errorf("latest release is a draft; nothing to update to")
	}

	latest, err := parseSemver(rel.TagName)
	if err != nil {
		return fmt.Errorf("parse release tag %q: %w", rel.TagName, err)
	}
	current, err := parseSemver(version.Short())
	if err != nil {
		return fmt.Errorf("parse current version %q: %w", version.Short(), err)
	}

	cmp := current.compare(latest)
	var prompt string
	switch {
	case cmp == 0 && !version.IsExperimental():
		fmt.Fprintln(os.Stderr, "You are already on the latest stable release.")
		return nil
	case cmp == 0 && version.IsExperimental():
		fmt.Fprintf(os.Stderr,
			"You are on %s, an experimental build at the same semver as stable %s.\n",
			version.Full(), rel.TagName)
		if opts.check {
			setCheckAvailable()
			return nil
		}
		prompt = fmt.Sprintf("Pin back to the clean stable release %s?", rel.TagName)
	case cmp > 0:
		fmt.Fprintf(os.Stderr,
			"Your version (%s) is newer than the latest stable release (%s).\n"+
				"Nothing to update.\n", version.Full(), rel.TagName)
		return nil
	default:
		bump := classifyBump(current, latest)
		fmt.Fprintf(os.Stderr, "A new %s release is available: %s (published %s)\n",
			bump, rel.TagName, rel.PublishedAt.Format("2006-01-02"))
		if rel.HTMLURL != "" {
			fmt.Fprintf(os.Stderr, "  %s\n", rel.HTMLURL)
		}
		if opts.check {
			setCheckAvailable()
			return nil
		}
		prompt = fmt.Sprintf("Update to %s now?", rel.TagName)
	}

	if !opts.yes {
		ok, err := confirm(prompt)
		if err != nil {
			return err
		}
		if !ok {
			fmt.Fprintln(os.Stderr, "Aborted.")
			return nil
		}
	}

	binPath, err := currentBinaryPath()
	if err != nil {
		return err
	}
	if !canWrite(binPath) {
		return fmt.Errorf("cannot write to %s — re-run with sudo or as the install owner", binPath)
	}

	tmp, err := os.MkdirTemp("", "carrion-update-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	bins, err := acquireReleaseBundle(rel, tmp, latest, opts)
	if err != nil {
		return err
	}

	installed, err := installBundle(bins, filepath.Dir(binPath))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Updated to %s (%s) at %s\n", rel.TagName, strings.Join(installed, ", "), filepath.Dir(binPath))
	return nil
}

// runExperimental tracks main-branch HEAD: fetch the latest commit SHA, short-
// circuit if already on that commit, otherwise clone the source, build with
// version/commit baked in, and install.
func runExperimental(opts options) error {
	fmt.Fprintf(os.Stderr, "Current version: %s\n", version.Full())

	commit, err := fetchLatestCommit("main")
	if err != nil {
		return fmt.Errorf("fetch main HEAD: %w", err)
	}
	latestShort := commit.SHA
	if len(latestShort) > 7 {
		latestShort = latestShort[:7]
	}

	if version.CommitShort() == latestShort {
		fmt.Fprintf(os.Stderr, "You are already on main@%s.\n", latestShort)
		return nil
	}

	fmt.Fprintf(os.Stderr, "New commit on main: %s\n", latestShort)
	if subj := strings.SplitN(commit.Commit.Message, "\n", 2)[0]; subj != "" {
		fmt.Fprintf(os.Stderr, "  %s\n", subj)
	}
	if commit.HTMLURL != "" {
		fmt.Fprintf(os.Stderr, "  %s\n", commit.HTMLURL)
	}

	if opts.check {
		setCheckAvailable()
		return nil
	}

	if !opts.yes {
		ok, err := confirm(fmt.Sprintf("Update to v%s-%s now?", version.Short(), latestShort))
		if err != nil {
			return err
		}
		if !ok {
			fmt.Fprintln(os.Stderr, "Aborted.")
			return nil
		}
	}

	binPath, err := currentBinaryPath()
	if err != nil {
		return err
	}
	if !canWrite(binPath) {
		return fmt.Errorf("cannot write to %s — re-run with sudo or as the install owner", binPath)
	}

	tmp, err := os.MkdirTemp("", "carrion-update-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmp)

	srcDir := filepath.Join(tmp, "src")
	outDir := filepath.Join(tmp, "built")
	fmt.Fprintln(os.Stderr, "Cloning source…")
	if err := cloneAt(commit.SHA, srcDir); err != nil {
		return err
	}
	semver := readSourceVersion(srcDir)
	if semver == "" {
		semver = version.Short()
	}
	fmt.Fprintln(os.Stderr, "Building…")
	bins, err := buildFromSource(srcDir, outDir, semver, commit.SHA, "experimental")
	if err != nil {
		return err
	}
	installed, err := installBundle(bins, filepath.Dir(binPath))
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "Updated to v%s-%s (%s) at %s\n", semver, latestShort, strings.Join(installed, ", "), filepath.Dir(binPath))
	return nil
}

// acquireReleaseBundle picks the right strategy for obtaining the stable-
// release binary set (carrion + sindri + mimir): prefer a prebuilt asset
// matching the host OS/arch, fall back to a source build at the release tag.
func acquireReleaseBundle(rel *releaseInfo, tmp string, latest semver, opts options) (map[string]string, error) {
	assetName, supported := assetNameForHost()
	if supported {
		for _, a := range rel.Assets {
			if a.Name == assetName {
				archivePath := filepath.Join(tmp, a.Name)
				fmt.Fprintf(os.Stderr, "Downloading %s…\n", a.Name)
				if err := downloadTo(a.BrowserDownloadURL, archivePath); err != nil {
					return nil, err
				}
				extractDir := filepath.Join(tmp, "extracted")
				bins, err := extractAsset(archivePath, extractDir)
				if err != nil {
					return nil, err
				}
				if bins["carrion"] == "" {
					return nil, fmt.Errorf("release asset %s did not contain a carrion binary", a.Name)
				}
				return bins, nil
			}
		}
		if opts.verbose {
			fmt.Fprintf(os.Stderr, "No asset named %s in release %s — falling back to source build.\n",
				assetName, rel.TagName)
		}
	} else if opts.verbose {
		fmt.Fprintf(os.Stderr, "Host OS %q has no prebuilt asset convention — falling back to source build.\n",
			runtime.GOOS)
	}

	srcDir := filepath.Join(tmp, "src")
	outDir := filepath.Join(tmp, "built")
	fmt.Fprintln(os.Stderr, "Cloning source…")
	if err := cloneAt(rel.TagName, srcDir); err != nil {
		return nil, err
	}
	fmt.Fprintln(os.Stderr, "Building…")
	return buildFromSource(srcDir, outDir, fmt.Sprintf("%d.%d.%d", latest.Major, latest.Minor, latest.Patch), "", "release")
}

func extractAsset(archivePath, destDir string) (map[string]string, error) {
	if strings.HasSuffix(archivePath, ".zip") {
		return extractZip(archivePath, destDir)
	}
	return extractTarGz(archivePath, destDir)
}

func confirm(prompt string) (bool, error) {
	fmt.Fprintf(os.Stderr, "%s [y/N] ", prompt)
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil && err != io.EOF {
		return false, err
	}
	line = strings.ToLower(strings.TrimSpace(line))
	return line == "y" || line == "yes", nil
}

