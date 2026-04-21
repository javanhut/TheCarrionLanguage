// versionsync keeps version references in documentation aligned with the
// single source of truth in src/version/version.go. Run without flags to
// rewrite stale references. Run with --check to exit non-zero if anything is
// out of sync (intended for CI and pre-commit hooks).
//
// Invocation:
//
//	go run ./cmd/versionsync          # apply updates in place
//	go run ./cmd/versionsync --check  # report without modifying
//
// Add a new file that should track the Carrion version by appending an entry
// to the `rules` slice below — there's no YAML/config file intentionally, so
// the set of synced files stays grep-able from Go source.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
)

const versionFile = "src/version/version.go"

// rule describes one file and the set of regex→format replacements the tool
// should apply whenever the version bumps. Exactly one verb (%s) is required
// per replacement so we can substitute the parsed version string.
type rule struct {
	path     string
	replace  []replacement
}

type replacement struct {
	// re matches the substring to replace. Use capturing groups sparingly;
	// the replacement format string receives only the semver, not captures.
	re     *regexp.Regexp
	format string // must contain exactly one %s for the version
}

func mustRe(pattern string) *regexp.Regexp {
	return regexp.MustCompile(pattern)
}

// rules is the full set of doc locations kept in sync with the Version const.
// Add entries here rather than sprinkling version mentions ad-hoc throughout
// docs — if it isn't in this list, it won't stay current.
var rules = []rule{
	{
		path: "docs/README.md",
		replace: []replacement{
			{mustRe(`Latest Version: \d+\.\d+\.\d+`), "Latest Version: %s"},
			{mustRe(`img\.shields\.io/badge/version-\d+\.\d+\.\d+-blue`), "img.shields.io/badge/version-%s-blue"},
			{mustRe(`badge/version-\d+\.\d+\.\d+\.svg`), "badge/version-%s.svg"},
			{mustRe(`- Current Version: \d+\.\d+\.\d+`), "- Current Version: %s"},
		},
	},
	{
		path: "docs/DOCUMENTATION.md",
		replace: []replacement{
			{mustRe("Base semver \\(major\\.minor\\.patch\\), e\\.g\\. `\\d+\\.\\d+\\.\\d+`"), "Base semver (major.minor.patch), e.g. `%s`"},
		},
	},
}

// readVersion parses `var Version = "x.y.z"` from src/version/version.go.
// Returns a descriptive error if the pattern is missing so a future rename of
// the field fails loudly rather than silently writing blanks into the docs.
func readVersion(root string) (string, error) {
	path := filepath.Join(root, versionFile)
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("read %s: %w", path, err)
	}
	re := regexp.MustCompile(`var\s+Version\s*=\s*"([^"]+)"`)
	m := re.FindSubmatch(data)
	if len(m) < 2 {
		return "", fmt.Errorf("could not locate `var Version = \"...\"` in %s", path)
	}
	return string(m[1]), nil
}

// sync applies all rules and returns the list of files it changed. In check
// mode it returns the list of files that *would* change without writing.
func sync(root, version string, check bool) ([]string, error) {
	var changed []string
	for _, r := range rules {
		full := filepath.Join(root, r.path)
		data, err := os.ReadFile(full)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", full, err)
		}
		orig := data
		for _, rep := range r.replace {
			repl := []byte(fmt.Sprintf(rep.format, version))
			data = rep.re.ReplaceAll(data, repl)
		}
		if !bytes.Equal(orig, data) {
			changed = append(changed, r.path)
			if !check {
				if err := os.WriteFile(full, data, 0o644); err != nil {
					return nil, fmt.Errorf("write %s: %w", full, err)
				}
			}
		}
	}
	return changed, nil
}

func main() {
	check := flag.Bool("check", false, "Report files that would change; exit 1 if any. Do not write.")
	rootFlag := flag.String("root", ".", "Repository root (defaults to current directory)")
	flag.Parse()

	root, err := filepath.Abs(*rootFlag)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	v, err := readVersion(root)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	changed, err := sync(root, v, *check)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(2)
	}

	if *check {
		if len(changed) > 0 {
			fmt.Fprintln(os.Stderr, "Out of sync with version", v+":")
			for _, f := range changed {
				fmt.Fprintln(os.Stderr, "  -", f)
			}
			fmt.Fprintln(os.Stderr, "Run `make sync-version` (or `go run ./cmd/versionsync`) to fix.")
			os.Exit(1)
		}
		fmt.Fprintln(os.Stderr, "Documentation is in sync with version", v+".")
		return
	}

	if len(changed) == 0 {
		fmt.Fprintln(os.Stderr, "No changes — docs already match version", v+".")
		return
	}
	fmt.Fprintln(os.Stderr, "Synced to version", v+":")
	for _, f := range changed {
		fmt.Fprintln(os.Stderr, "  -", f)
	}
}
