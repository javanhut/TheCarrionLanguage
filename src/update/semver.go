package update

import (
	"fmt"
	"strconv"
	"strings"
)

// semver is a minimal major.minor.patch representation. Carrion releases don't
// use build metadata or pre-release identifiers beyond an optional commit
// suffix (which we strip before comparison), so a richer parser isn't needed.
type semver struct {
	Major, Minor, Patch int
}

// parseSemver accepts inputs like "v0.1.9", "0.1.9", "v0.1.9-a1b2c3d", and
// returns the semver with any pre-release suffix discarded.
func parseSemver(s string) (semver, error) {
	s = strings.TrimSpace(s)
	s = strings.TrimPrefix(s, "v")
	if i := strings.IndexAny(s, "-+"); i >= 0 {
		s = s[:i]
	}
	parts := strings.Split(s, ".")
	if len(parts) != 3 {
		return semver{}, fmt.Errorf("not a major.minor.patch version: %q", s)
	}
	out := semver{}
	vals := []*int{&out.Major, &out.Minor, &out.Patch}
	for i, p := range parts {
		n, err := strconv.Atoi(p)
		if err != nil {
			return semver{}, fmt.Errorf("invalid number in version %q: %s", s, p)
		}
		*vals[i] = n
	}
	return out, nil
}

// compare returns -1, 0, 1 like strings.Compare.
func (a semver) compare(b semver) int {
	if a.Major != b.Major {
		if a.Major < b.Major {
			return -1
		}
		return 1
	}
	if a.Minor != b.Minor {
		if a.Minor < b.Minor {
			return -1
		}
		return 1
	}
	if a.Patch != b.Patch {
		if a.Patch < b.Patch {
			return -1
		}
		return 1
	}
	return 0
}

// bumpKind classifies the difference between two versions so the updater can
// tell the user "this is a major update" vs "this is a patch".
type bumpKind int

const (
	bumpNone bumpKind = iota
	bumpPatch
	bumpMinor
	bumpMajor
)

func (b bumpKind) String() string {
	switch b {
	case bumpPatch:
		return "patch"
	case bumpMinor:
		return "minor"
	case bumpMajor:
		return "major"
	default:
		return "none"
	}
}

func classifyBump(from, to semver) bumpKind {
	switch {
	case to.Major > from.Major:
		return bumpMajor
	case to.Minor > from.Minor:
		return bumpMinor
	case to.Patch > from.Patch:
		return bumpPatch
	default:
		return bumpNone
	}
}
