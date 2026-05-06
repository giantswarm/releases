package breakingchanges

import (
	"testing"

	"github.com/Masterminds/semver/v3"
)

func TestEvaluateKubeVersionConstraint(t *testing.T) {
	target := semver.MustParse("1.35.0")

	cases := []struct {
		name        string
		kubeVersion string
		wantFinding bool
		wantSev     string
	}{
		{
			name:        "open-ended supports target",
			kubeVersion: ">= 1.21.0-0",
			wantFinding: false,
		},
		{
			name:        "upper bound excludes target",
			kubeVersion: ">= 1.27.0-0 < 1.34.0-0",
			wantFinding: true,
			wantSev:     "high",
		},
		{
			name:        "explicit minor range includes target",
			kubeVersion: ">= 1.34.0-0 < 1.36.0-0",
			wantFinding: false,
		},
		{
			name:        "unparseable surfaces as low-confidence info",
			kubeVersion: "this is not a constraint",
			wantFinding: true,
			wantSev:     "low",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			f := evaluateKubeVersionConstraint(tc.kubeVersion, target, "test-app", "1.0.0", "test")
			if tc.wantFinding && f == nil {
				t.Fatalf("expected a finding, got nil")
			}
			if !tc.wantFinding && f != nil {
				t.Fatalf("expected no finding, got %+v", f)
			}
			if tc.wantFinding && f.Severity != tc.wantSev {
				t.Errorf("severity = %q, want %q", f.Severity, tc.wantSev)
			}
		})
	}
}

func TestIsExternalHelmRepo(t *testing.T) {
	cases := map[string]bool{
		"https://helm.cilium.io": true,
		"http://charts.example":  true,
		"oci://ghcr.io/x/y":      false,
		"":                       false,
		"file://./local-chart":   false,
	}
	for in, want := range cases {
		if got := isExternalHelmRepo(in); got != want {
			t.Errorf("isExternalHelmRepo(%q) = %v, want %v", in, got, want)
		}
	}
}
