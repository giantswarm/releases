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

func TestParseCodeownersTeam(t *testing.T) {
	cases := []struct {
		name    string
		content string
		want    string
	}{
		{
			name:    "single team",
			content: "* @giantswarm/team-phoenix",
			want:    "team-phoenix",
		},
		{
			name: "skips comments and blank lines",
			content: `# Owners
# of this repo

* @giantswarm/team-rocket
`,
			want: "team-rocket",
		},
		{
			name:    "scoped path before team",
			content: "/helm/ @giantswarm/team-cabbage",
			want:    "team-cabbage",
		},
		{
			name:    "first non-comment line wins",
			content: "*.go @giantswarm/team-phoenix\n*.md @giantswarm/team-rocket",
			want:    "team-phoenix",
		},
		{
			name:    "no giantswarm team handle",
			content: "* @some-user",
			want:    "",
		},
		{
			name:    "empty",
			content: "",
			want:    "",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := parseCodeownersTeam(tc.content); got != tc.want {
				t.Errorf("parseCodeownersTeam = %q, want %q", got, tc.want)
			}
		})
	}
}

func TestBuildTeamIssuesPayload(t *testing.T) {
	findings := []FindingWithProvider{
		{
			Finding: Finding{
				Severity: "high",
				Title:    "`cilium` v1.3.4 does not declare support for Kubernetes v1.35.0",
				Action:   "Bump cilium",
				Team:     "team-cabbage",
				AppName:  "cilium",
			},
			Provider: "aws",
			Version:  "v35.0.0",
		},
		{
			// Same app+team as above from a different provider — should dedupe.
			Finding: Finding{
				Severity: "high",
				Title:    "`cilium` v1.3.4 does not declare support for Kubernetes v1.35.0",
				Action:   "Bump cilium",
				Team:     "team-cabbage",
				AppName:  "cilium",
			},
			Provider: "azure",
			Version:  "v35.0.0",
		},
		{
			Finding: Finding{
				Severity: "high",
				Title:    "`cloud-provider-aws` v2.0.0 needs a new release for Kubernetes v1.35.0",
				Action:   "Bump cloud-provider-aws",
				Team:     "team-phoenix",
				AppName:  "cloud-provider-aws",
			},
			Provider: "aws",
			Version:  "v35.0.0",
		},
		{
			// Finding without team — should be excluded.
			Finding: Finding{
				Severity:  "high",
				Title:     "Kubernetes 1.35 deprecates v1beta1 FlowSchema",
				Component: "Kubernetes",
			},
			Provider: "aws",
			Version:  "v35.0.0",
		},
	}

	payload := BuildTeamIssuesPayload(findings, "1.35.0")

	if payload.KubernetesVersion != "1.35.0" {
		t.Errorf("KubernetesVersion = %q, want %q", payload.KubernetesVersion, "1.35.0")
	}
	if payload.Version != "v35.0.0" {
		t.Errorf("Version = %q, want %q", payload.Version, "v35.0.0")
	}
	if len(payload.Teams) != 2 {
		t.Fatalf("Teams length = %d, want 2 (cabbage + phoenix)", len(payload.Teams))
	}

	// Sorted alphabetically: cabbage, phoenix
	if payload.Teams[0].Team != "team-cabbage" {
		t.Errorf("Teams[0] = %q, want team-cabbage", payload.Teams[0].Team)
	}
	if len(payload.Teams[0].Apps) != 1 {
		t.Errorf("team-cabbage apps = %d, want 1 (deduped)", len(payload.Teams[0].Apps))
	}
	if payload.Teams[0].Apps[0].Version != "1.3.4" {
		t.Errorf("cilium version = %q, want 1.3.4", payload.Teams[0].Apps[0].Version)
	}
	if payload.Teams[1].Team != "team-phoenix" {
		t.Errorf("Teams[1] = %q, want team-phoenix", payload.Teams[1].Team)
	}
}

func TestImplicitlyK8sCoupledAppsContainsCloudControllers(t *testing.T) {
	required := []string{
		"cloud-provider-aws",
		"azure-cloud-controller-manager",
		"azure-cloud-node-manager",
		"cloud-provider-vsphere",
		"cloud-provider-cloud-director",
		"cloud-provider-proxmox",
	}
	for _, app := range required {
		if !implicitlyK8sCoupledApps[app] {
			t.Errorf("expected %q in implicitlyK8sCoupledApps fallback", app)
		}
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
