package main

import (
	"encoding/json"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

const oursReleases = `{
  "releases": [
    {"version": "33.1.4", "isDeprecated": false, "releaseTimestamp": "2026-01-22T20:26:21Z", "changelogUrl": "x/v33.1.4", "isStable": true},
    {"version": "33.2.0", "isDeprecated": false, "releaseTimestamp": "2026-04-21T13:45:36Z", "changelogUrl": "x/v33.2.0", "isStable": true}
  ]
}`

const theirsReleases = `{
  "releases": [
    {"version": "33.1.4", "isDeprecated": false, "releaseTimestamp": "2026-01-22T20:26:21Z", "changelogUrl": "x/v33.1.4", "isStable": true},
    {"version": "34.0.0", "isDeprecated": false, "releaseTimestamp": "2026-04-01T00:00:00Z", "changelogUrl": "x/v34.0.0", "isStable": true}
  ]
}`

func TestMergeReleasesJSON_UnionAndSort(t *testing.T) {
	merged, err := mergeReleasesJSON([]byte(oursReleases), []byte(theirsReleases))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got releasesFile
	if err := json.Unmarshal(merged, &got); err != nil {
		t.Fatalf("merged output is not valid JSON: %v", err)
	}

	versions := make([]string, len(got.Releases))
	for i, r := range got.Releases {
		versions[i] = r.Version
	}
	want := []string{"33.1.4", "33.2.0", "34.0.0"}
	if !equalSlices(versions, want) {
		t.Errorf("versions = %v, want %v", versions, want)
	}
}

func TestMergeReleasesJSON_PrefersTheirsOnVersionCollision(t *testing.T) {
	ours := `{"releases": [{"version": "1.0.0", "isDeprecated": false, "releaseTimestamp": "2020", "changelogUrl": "old", "isStable": false}]}`
	theirs := `{"releases": [{"version": "1.0.0", "isDeprecated": true,  "releaseTimestamp": "2021", "changelogUrl": "new", "isStable": true}]}`

	merged, err := mergeReleasesJSON([]byte(ours), []byte(theirs))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got releasesFile
	if err := json.Unmarshal(merged, &got); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}
	if len(got.Releases) != 1 {
		t.Fatalf("expected 1 entry after dedup, got %d", len(got.Releases))
	}
	r := got.Releases[0]
	if r.ChangelogUrl != "new" || !r.IsDeprecated || !r.IsStable {
		t.Errorf("expected theirs to win, got %+v", r)
	}
}

func TestMergeReleasesJSON_SortsBySemverNotLex(t *testing.T) {
	ours := `{"releases": [{"version": "9.0.0", "isDeprecated": false, "releaseTimestamp": "", "changelogUrl": "", "isStable": true}]}`
	theirs := `{"releases": [{"version": "10.0.0", "isDeprecated": false, "releaseTimestamp": "", "changelogUrl": "", "isStable": true}]}`

	merged, err := mergeReleasesJSON([]byte(ours), []byte(theirs))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var got releasesFile
	_ = json.Unmarshal(merged, &got)
	if got.Releases[0].Version != "9.0.0" || got.Releases[1].Version != "10.0.0" {
		t.Errorf("semver sort wrong, got %s then %s", got.Releases[0].Version, got.Releases[1].Version)
	}
}

const oursKustomization = `commonAnnotations:
  giantswarm.io/docs: https://docs.example
resources:
- v33.1.4
- v33.2.0
transformers:
- |
  apiVersion: builtin
  kind: PrefixSuffixTransformer
`

const theirsKustomization = `commonAnnotations:
  giantswarm.io/docs: https://docs.example
resources:
- v33.1.4
- v34.0.0
transformers:
- |
  apiVersion: builtin
  kind: PrefixSuffixTransformer
`

func TestMergeKustomization_UnionAndSort(t *testing.T) {
	merged, err := mergeKustomization([]byte(oursKustomization), []byte(theirsKustomization))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var got struct {
		CommonAnnotations map[string]string `yaml:"commonAnnotations"`
		Resources         []string          `yaml:"resources"`
		Transformers     []string          `yaml:"transformers"`
	}
	if err := yaml.Unmarshal(merged, &got); err != nil {
		t.Fatalf("merged output is not valid YAML: %v\n%s", err, merged)
	}

	want := []string{"v33.1.4", "v33.2.0", "v34.0.0"}
	if !equalSlices(got.Resources, want) {
		t.Errorf("resources = %v, want %v", got.Resources, want)
	}
	if got.CommonAnnotations["giantswarm.io/docs"] != "https://docs.example" {
		t.Errorf("commonAnnotations not preserved: %+v", got.CommonAnnotations)
	}
	if len(got.Transformers) != 1 || !strings.Contains(got.Transformers[0], "PrefixSuffixTransformer") {
		t.Errorf("transformers not preserved: %+v", got.Transformers)
	}
}

func TestMergeKustomization_PreservesColumnZeroListStyle(t *testing.T) {
	// Matches the style used across giantswarm/releases kustomization files.
	ours := "commonAnnotations:\n  giantswarm.io/docs: x\nresources:\n- v33.1.4\n- v33.2.0\ntransformers:\n- |\n  apiVersion: builtin\n"
	theirs := "commonAnnotations:\n  giantswarm.io/docs: x\nresources:\n- v33.1.4\n- v34.0.0\ntransformers:\n- |\n  apiVersion: builtin\n"

	merged, err := mergeKustomization([]byte(ours), []byte(theirs))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got := string(merged)
	// The merged resources must be at column zero, not indented.
	want := "resources:\n- v33.1.4\n- v33.2.0\n- v34.0.0\ntransformers:"
	if !strings.Contains(got, want) {
		t.Errorf("expected column-zero list style preserved.\n--- want contains ---\n%s\n--- got ---\n%s", want, got)
	}
}

func TestMergeKustomization_DedupSameVersion(t *testing.T) {
	ours := "resources:\n- v1.0.0\n- v2.0.0\n"
	theirs := "resources:\n- v1.0.0\n- v2.0.0\n- v3.0.0\n"

	merged, err := mergeKustomization([]byte(ours), []byte(theirs))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	var got struct {
		Resources []string `yaml:"resources"`
	}
	_ = yaml.Unmarshal(merged, &got)
	want := []string{"v1.0.0", "v2.0.0", "v3.0.0"}
	if !equalSlices(got.Resources, want) {
		t.Errorf("resources = %v, want %v", got.Resources, want)
	}
}

func TestSemverLess(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"v9.0.0", "v10.0.0", true},
		{"v10.0.0", "v9.0.0", false},
		{"33.1.4", "33.2.0", true},
		{"v33.2.0", "v34.0.0", true},
		{"v34.0.0", "v34.0.0", false},
	}
	for _, c := range cases {
		if got := semverLess(c.a, c.b); got != c.want {
			t.Errorf("semverLess(%q, %q) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}

func equalSlices(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
