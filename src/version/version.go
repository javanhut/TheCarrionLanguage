// Package version exposes the build-time version metadata for the Carrion
// language binary. All fields are intended to be overridden at build time via
// `go build -ldflags "-X github.com/javanhut/TheCarrionLanguage/src/version.Version=..."`
// so the running binary can report exactly what source it was built from.
package version

// Version is the base Carrion semver: major.minor.patch (no leading "v", no
// pre-release suffix). The release workflow rewrites this value on tag push
// so bumping is automatic — do not edit by hand unless you're doing something
// unusual. Local `go build` without ldflags picks up this default.
var Version = "0.1.9"

// Commit is the short git SHA the binary was built from. Empty string means
// "tagged release" — omit the suffix when formatting the version. Any non-empty
// value means the binary was built from an untagged commit (experimental).
var Commit = ""

// Channel identifies the release stream:
//   - "release"      — built from a tagged release
//   - "experimental" — built from a main-branch commit via `carrion update --experimental`
//   - "dev"          — local developer build (default)
var Channel = "dev"

// BuildDate is the UTC build timestamp in RFC3339 format. Optional; empty when
// not injected.
var BuildDate = ""

// Full returns the user-facing version string.
//
//	v0.1.9          — tagged release (Commit empty)
//	v0.1.9-a1b2c3d  — experimental build (Commit set; trimmed to 7 chars)
func Full() string {
	if Commit == "" {
		return "v" + Version
	}
	return "v" + Version + "-" + shortCommit()
}

// Short returns the raw semver without the leading "v" and without the commit
// suffix. Useful for semver comparisons.
func Short() string {
	return Version
}

// CommitShort returns the 7-character abbreviated commit hash, or "" if Commit
// was not injected.
func CommitShort() string {
	return shortCommit()
}

// IsExperimental reports whether this binary was built from a non-tagged commit.
func IsExperimental() bool {
	return Commit != ""
}

func shortCommit() string {
	if len(Commit) <= 7 {
		return Commit
	}
	return Commit[:7]
}
