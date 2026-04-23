package breakingchanges

import (
	"strings"
	"testing"
)

func TestParseSupportedEtcdVersion(t *testing.T) {
	constantsGo := `
	// SupportedEtcdVersion lists officially supported etcd versions.
	SupportedEtcdVersion = map[uint8]string{
		31: "3.5.21-0",
		32: "3.5.21-0",
		33: "3.5.21-0",
		34: "3.6.4-0",
	}
`
	cases := []struct {
		minor int
		want  string
	}{
		{33, "3.5.21"},
		{34, "3.6.4"},
	}
	for _, tc := range cases {
		got, err := parseSupportedEtcdVersion(constantsGo, tc.minor)
		if err != nil {
			t.Fatalf("minor %d: unexpected error: %v", tc.minor, err)
		}
		if got != tc.want {
			t.Errorf("minor %d: got %q, want %q", tc.minor, got, tc.want)
		}
	}

	if _, err := parseSupportedEtcdVersion(constantsGo, 99); err == nil {
		t.Error("expected error for missing minor 99")
	}
}

func TestStripEtcdImageSuffix(t *testing.T) {
	cases := map[string]string{
		"3.6.4-0": "3.6.4",
		"3.5.21":  "3.5.21",
	}
	for in, want := range cases {
		if got := stripEtcdImageSuffix(in); got != want {
			t.Errorf("stripEtcdImageSuffix(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestCompareSemver(t *testing.T) {
	cases := []struct {
		a, b string
		sign int // -1, 0, +1
	}{
		{"3.6.4", "3.6.4", 0},
		{"3.6.4", "3.6.5", -1},
		{"3.6.10", "3.6.9", 1},
		{"3.5.21", "3.6.0", -1},
	}
	for _, tc := range cases {
		got := compareSemver(tc.a, tc.b)
		if (got == 0 && tc.sign != 0) || (got < 0 && tc.sign >= 0) || (got > 0 && tc.sign <= 0) {
			t.Errorf("compareSemver(%q, %q) = %d, want sign %d", tc.a, tc.b, got, tc.sign)
		}
	}
}

func TestExtractEtcdBetween(t *testing.T) {
	changelog := `
## v3.6.11 (TBC)

- should be excluded (greater than toVer)

---

## v3.6.10 (2026-04-01)

- should be included

---

## v3.6.9 (2026-03-20)

- should be included

---

## v3.6.8 (2026-02-13)

- should be excluded (equal to fromVer)
`
	got := extractEtcdBetween(changelog, "3.6.8", "3.6.10")
	if !strings.Contains(got, "v3.6.10") || !strings.Contains(got, "v3.6.9") {
		t.Errorf("expected 3.6.9 and 3.6.10 in slice, got:\n%s", got)
	}
	if strings.Contains(got, "v3.6.11") {
		t.Errorf("expected 3.6.11 (TBC) to be excluded, got:\n%s", got)
	}
	if strings.Contains(got, "v3.6.8") {
		t.Errorf("expected fromVer 3.6.8 to be excluded, got:\n%s", got)
	}
}

func TestExtractEtcdUpTo(t *testing.T) {
	changelog := `
## v3.6.5 (2025-07-31)

- included

---

## v3.6.4 (2025-06-15)

- included

---

## v3.6.3 (2025-05-01)

- included
`
	got := extractEtcdUpTo(changelog, "3.6.4")
	if !strings.Contains(got, "v3.6.4") || !strings.Contains(got, "v3.6.3") {
		t.Errorf("expected 3.6.3 and 3.6.4 in slice, got:\n%s", got)
	}
	if strings.Contains(got, "v3.6.5") {
		t.Errorf("expected 3.6.5 (greater than toVer) to be excluded, got:\n%s", got)
	}
}
