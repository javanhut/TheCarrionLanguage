package update

import "testing"

func TestParseSemver(t *testing.T) {
	cases := []struct {
		in      string
		want    semver
		wantErr bool
	}{
		{"0.1.9", semver{0, 1, 9}, false},
		{"v0.1.9", semver{0, 1, 9}, false},
		{" v1.2.3 ", semver{1, 2, 3}, false},
		{"v0.1.9-a1b2c3d", semver{0, 1, 9}, false},
		{"v1.0.0+build.42", semver{1, 0, 0}, false},
		{"v10.20.30", semver{10, 20, 30}, false},
		{"", semver{}, true},
		{"v1.2", semver{}, true},
		{"v1.2.x", semver{}, true},
		{"not-a-version", semver{}, true},
	}
	for _, tc := range cases {
		got, err := parseSemver(tc.in)
		if tc.wantErr {
			if err == nil {
				t.Errorf("parseSemver(%q): want error, got %+v", tc.in, got)
			}
			continue
		}
		if err != nil {
			t.Errorf("parseSemver(%q): unexpected error %v", tc.in, err)
			continue
		}
		if got != tc.want {
			t.Errorf("parseSemver(%q) = %+v, want %+v", tc.in, got, tc.want)
		}
	}
}

func TestSemverCompare(t *testing.T) {
	cases := []struct {
		a, b semver
		want int
	}{
		{semver{0, 1, 9}, semver{0, 1, 9}, 0},
		{semver{0, 1, 9}, semver{0, 1, 10}, -1},
		{semver{0, 1, 10}, semver{0, 1, 9}, 1},
		{semver{0, 2, 0}, semver{0, 1, 99}, 1},
		{semver{1, 0, 0}, semver{0, 99, 99}, 1},
		{semver{0, 1, 9}, semver{0, 2, 0}, -1},
	}
	for _, tc := range cases {
		if got := tc.a.compare(tc.b); got != tc.want {
			t.Errorf("%+v.compare(%+v) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

func TestClassifyBump(t *testing.T) {
	cases := []struct {
		from, to semver
		want     bumpKind
	}{
		{semver{0, 1, 9}, semver{0, 1, 9}, bumpNone},
		{semver{0, 1, 9}, semver{0, 1, 10}, bumpPatch},
		{semver{0, 1, 9}, semver{0, 2, 0}, bumpMinor},
		{semver{0, 1, 9}, semver{1, 0, 0}, bumpMajor},
		{semver{0, 1, 9}, semver{2, 0, 0}, bumpMajor},
	}
	for _, tc := range cases {
		if got := classifyBump(tc.from, tc.to); got != tc.want {
			t.Errorf("classifyBump(%+v, %+v) = %v, want %v", tc.from, tc.to, got, tc.want)
		}
	}
}
