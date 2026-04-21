package update

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// readSourceVersion extracts the default `var Version = "x.y.z"` value from
// the cloned source tree. Used so an experimental build reports the repo's
// current semver, not the running binary's semver (main may be ahead).
// Returns "" if the file is missing or unparseable — caller must supply a
// fallback.
func readSourceVersion(sourceDir string) string {
	path := filepath.Join(sourceDir, "src", "version", "version.go")
	data, err := os.ReadFile(path)
	if err != nil {
		return ""
	}
	re := regexp.MustCompile(`var\s+Version\s*=\s*"([^"]+)"`)
	m := re.FindSubmatch(data)
	if len(m) < 2 {
		return ""
	}
	return string(m[1])
}

// assetNameForHost returns the expected release asset filename for the host
// OS/arch, matching the naming used by the Makefile targets.
// Windows builds are shipped as .zip; Linux/macOS as .tar.gz.
func assetNameForHost() (string, bool) {
	switch runtime.GOOS {
	case "linux":
		return fmt.Sprintf("carrion_linux_%s.tar.gz", runtime.GOARCH), true
	case "darwin":
		return fmt.Sprintf("carrion_darwin_%s.tar.gz", runtime.GOARCH), true
	case "windows":
		return fmt.Sprintf("carrion_windows_%s.zip", runtime.GOARCH), true
	default:
		return "", false
	}
}

func currentBinaryPath() (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", fmt.Errorf("locate current binary: %w", err)
	}
	resolved, err := filepath.EvalSymlinks(exe)
	if err != nil {
		return exe, nil
	}
	return resolved, nil
}

// canWrite reports whether the calling user can replace the file at path.
// We probe with O_WRONLY rather than checking mode bits because effective
// permissions depend on ownership, ACLs, and mount options we can't read.
func canWrite(path string) bool {
	f, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return false
	}
	_ = f.Close()
	return true
}

// replaceBinary atomically swaps targetPath with the file at newPath,
// preserving the target's mode bits so the replacement stays executable.
// On Linux/macOS, renaming over a running binary is safe; on Windows, we move
// the current binary aside and then rename the new one in.
func replaceBinary(newPath, targetPath string) error {
	info, err := os.Stat(targetPath)
	if err != nil {
		return fmt.Errorf("stat %s: %w", targetPath, err)
	}
	mode := info.Mode().Perm()
	if err := os.Chmod(newPath, mode); err != nil {
		return fmt.Errorf("chmod new binary: %w", err)
	}

	if runtime.GOOS == "windows" {
		backup := targetPath + ".old"
		_ = os.Remove(backup)
		if err := os.Rename(targetPath, backup); err != nil {
			return fmt.Errorf("move old binary aside: %w", err)
		}
		if err := os.Rename(newPath, targetPath); err != nil {
			_ = os.Rename(backup, targetPath)
			return fmt.Errorf("install new binary: %w", err)
		}
		return nil
	}

	if err := os.Rename(newPath, targetPath); err != nil {
		return fmt.Errorf("install new binary: %w", err)
	}
	return nil
}

// extractTarGz extracts a gzipped tar archive into destDir and returns the
// paths of the Carrion-related binaries it wrote (carrion, sindri, mimir).
func extractTarGz(archivePath, destDir string) (map[string]string, error) {
	f, err := os.Open(archivePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	gz, err := gzip.NewReader(f)
	if err != nil {
		return nil, fmt.Errorf("gzip open: %w", err)
	}
	defer gz.Close()
	tr := tar.NewReader(gz)
	return extractStream(tr, destDir)
}

// extractZip handles the Windows release format.
func extractZip(archivePath, destDir string) (map[string]string, error) {
	r, err := zip.OpenReader(archivePath)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	binaries := map[string]string{}
	for _, f := range r.File {
		if f.FileInfo().IsDir() {
			continue
		}
		outPath, name := placeInto(destDir, f.Name)
		if outPath == "" {
			continue
		}
		src, err := f.Open()
		if err != nil {
			return nil, err
		}
		dst, err := createFile(outPath)
		if err != nil {
			src.Close()
			return nil, err
		}
		if _, err := io.Copy(dst, src); err != nil {
			src.Close()
			dst.Close()
			return nil, err
		}
		src.Close()
		dst.Close()
		if name != "" {
			binaries[name] = outPath
		}
	}
	return binaries, nil
}

type tarReader interface {
	Next() (*tar.Header, error)
	io.Reader
}

func extractStream(tr tarReader, destDir string) (map[string]string, error) {
	binaries := map[string]string{}
	for {
		h, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if h.FileInfo().IsDir() {
			continue
		}
		outPath, name := placeInto(destDir, h.Name)
		if outPath == "" {
			continue
		}
		dst, err := createFile(outPath)
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(dst, tr); err != nil {
			dst.Close()
			return nil, err
		}
		dst.Close()
		if name != "" {
			binaries[name] = outPath
		}
	}
	return binaries, nil
}

// placeInto maps an archive entry path to a destination path and reports the
// logical binary name (carrion, sindri, mimir) if it's one we care about.
// Returns ("", "") for entries to ignore.
func placeInto(destDir, entryPath string) (string, string) {
	base := filepath.Base(entryPath)
	switch strings.TrimSuffix(base, ".exe") {
	case "carrion":
		return filepath.Join(destDir, base), "carrion"
	case "sindri":
		return filepath.Join(destDir, base), "sindri"
	case "mimir":
		return filepath.Join(destDir, base), "mimir"
	}
	return "", ""
}

func createFile(path string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, err
	}
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o755)
}

// cloneAt fetches the repo into destDir and checks out the given ref.
// Uses --depth=1 for speed; falls back to a full fetch if the ref isn't in the
// shallow clone (rare — HEAD of main always is).
func cloneAt(ref, destDir string) error {
	if _, err := exec.LookPath("git"); err != nil {
		return fmt.Errorf("git not found in PATH — experimental updates require git")
	}
	url := fmt.Sprintf("https://github.com/%s/%s.git", repoOwner, repoName)
	cmd := exec.Command("git", "clone", "--depth", "50", url, destDir)
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git clone: %w", err)
	}
	if ref == "" {
		return nil
	}
	checkout := exec.Command("git", "-C", destDir, "checkout", ref)
	checkout.Stdout = os.Stderr
	checkout.Stderr = os.Stderr
	if err := checkout.Run(); err != nil {
		unshallow := exec.Command("git", "-C", destDir, "fetch", "--unshallow")
		unshallow.Stdout = os.Stderr
		unshallow.Stderr = os.Stderr
		if uerr := unshallow.Run(); uerr != nil {
			return fmt.Errorf("checkout %s: %w (unshallow also failed: %v)", ref, err, uerr)
		}
		retry := exec.Command("git", "-C", destDir, "checkout", ref)
		retry.Stdout = os.Stderr
		retry.Stderr = os.Stderr
		if err := retry.Run(); err != nil {
			return fmt.Errorf("checkout %s after unshallow: %w", ref, err)
		}
	}
	return nil
}

// buildFromSource compiles carrion plus the sindri and mimir companion
// binaries from sourceDir into outDir, baking the given version/commit/channel
// into carrion via ldflags. Returns a map of binary-name → absolute path for
// whatever built successfully. sindri and mimir are best-effort: if their
// packages are missing from the tree (older commits), they're skipped with a
// warning rather than failing the whole update.
func buildFromSource(sourceDir, outDir, semver, commit, channel string) (map[string]string, error) {
	if _, err := exec.LookPath("go"); err != nil {
		return nil, fmt.Errorf("go toolchain not found in PATH — experimental updates require go 1.24+")
	}
	ldflags := fmt.Sprintf(
		"-X github.com/javanhut/TheCarrionLanguage/src/version.Version=%s "+
			"-X github.com/javanhut/TheCarrionLanguage/src/version.Commit=%s "+
			"-X github.com/javanhut/TheCarrionLanguage/src/version.Channel=%s "+
			"-X github.com/javanhut/TheCarrionLanguage/src/version.BuildDate=%s",
		semver, commit, channel, time.Now().UTC().Format(time.RFC3339),
	)

	if err := os.MkdirAll(outDir, 0o755); err != nil {
		return nil, err
	}

	bins := map[string]string{}
	// carrion is the primary binary — failure here aborts the update.
	carrionOut := filepath.Join(outDir, binName("carrion"))
	cmd := exec.Command("go", "build", "-ldflags", ldflags, "-o", carrionOut, "./src")
	cmd.Dir = sourceDir
	cmd.Stdout = os.Stderr
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("go build carrion: %w", err)
	}
	bins["carrion"] = carrionOut

	// sindri and mimir are companions — skip with a warning if not present.
	for _, companion := range []struct{ name, pkg string }{
		{"sindri", "./cmd/sindri"},
		{"mimir", "./cmd/mimir"},
	} {
		if _, err := os.Stat(filepath.Join(sourceDir, companion.pkg)); err != nil {
			fmt.Fprintf(os.Stderr, "  (skipping %s — %s not in source tree)\n", companion.name, companion.pkg)
			continue
		}
		out := filepath.Join(outDir, binName(companion.name))
		c := exec.Command("go", "build", "-o", out, companion.pkg)
		c.Dir = sourceDir
		c.Stdout = os.Stderr
		c.Stderr = os.Stderr
		if err := c.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "  (warning: %s build failed: %v)\n", companion.name, err)
			continue
		}
		bins[companion.name] = out
	}
	return bins, nil
}

// binName returns the platform-appropriate executable filename.
func binName(name string) string {
	if runtime.GOOS == "windows" {
		return name + ".exe"
	}
	return name
}

// installBundle swaps each staged binary into destDir, replacing any existing
// carrion/sindri/mimir at that path. Binaries whose counterparts don't exist
// in destDir are skipped — the user may have only installed carrion, and we
// don't want to silently create new binaries they didn't ask for.
// Returns the list of binary names actually installed.
func installBundle(bins map[string]string, destDir string) ([]string, error) {
	// Install carrion last: if a companion fails, we haven't yet replaced the
	// main binary, so the user's install is still coherent.
	order := []string{"sindri", "mimir", "carrion"}
	var installed []string
	for _, name := range order {
		src, ok := bins[name]
		if !ok {
			continue
		}
		dest := filepath.Join(destDir, binName(name))
		if _, err := os.Stat(dest); err != nil {
			if os.IsNotExist(err) {
				if name == "carrion" {
					// carrion MUST exist — we found it via os.Executable().
					return installed, fmt.Errorf("carrion binary missing from %s (unexpected)", destDir)
				}
				fmt.Fprintf(os.Stderr, "  (skipping %s — not previously installed at %s)\n", name, destDir)
				continue
			}
			return installed, fmt.Errorf("stat %s: %w", dest, err)
		}
		if err := replaceBinary(src, dest); err != nil {
			return installed, fmt.Errorf("install %s: %w", name, err)
		}
		installed = append(installed, name)
	}
	return installed, nil
}

