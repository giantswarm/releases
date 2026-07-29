package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

const lowestActiveReleaseScript = "tools/lowest-active-release.sh"

// writeFixtureRelease creates a minimal release directory for the given version and state.
func writeFixtureRelease(t *testing.T, providerDir string, version string, state string, archived bool) {
	t.Helper()

	releaseDir := providerDir
	if archived {
		releaseDir = filepath.Join(providerDir, "archived")
	}
	releaseDir = filepath.Join(releaseDir, version)

	if err := os.MkdirAll(releaseDir, 0o755); err != nil {
		t.Fatal(err)
	}

	content := "apiVersion: release.giantswarm.io/v1alpha1\nkind: Release\nspec:\n  state: " + state + "\n"
	if err := os.WriteFile(filepath.Join(releaseDir, releaseFilename), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}

// Test_ArchiveFloor covers the archive floor used to decide which deprecated releases
// may be archived. Anything above the lowest active release has to stay available,
// because clusters on that lower release may still need it as an intermediate upgrade
// step. See https://github.com/giantswarm/roadmap/issues/4325
func Test_ArchiveFloor(t *testing.T) {
	scriptPath, err := filepath.Abs(lowestActiveReleaseScript)
	if err != nil {
		t.Fatal(err)
	}

	testCases := []struct {
		name string
		// releases maps a version to its state, "archived" marks an already archived release.
		releases map[string]string
		expected string
	}{
		{
			name:     "case 0: no releases at all yields no floor",
			expected: "",
		},
		{
			name: "case 1: only archived releases yields no floor",
			releases: map[string]string{
				"v32.1.0": "archived",
				"v32.2.0": "archived",
			},
			expected: "",
		},
		{
			name: "case 2: single active release is the floor",
			releases: map[string]string{
				"v33.4.0": "active",
			},
			expected: "33.4.0",
		},
		{
			name: "case 3: lowest active release wins over higher ones",
			releases: map[string]string{
				"v33.1.1": "active",
				"v33.4.0": "active",
				"v34.5.0": "active",
			},
			expected: "33.1.1",
		},
		{
			name: "case 4: deprecated releases do not lower the floor",
			releases: map[string]string{
				"v33.1.1": "deprecated",
				"v33.4.0": "active",
			},
			expected: "33.4.0",
		},
		{
			name: "case 5: archived releases do not lower the floor",
			releases: map[string]string{
				"v32.1.0": "archived",
				"v33.4.0": "active",
			},
			expected: "33.4.0",
		},
		{
			name: "case 6: floor is compared by version, not lexically",
			releases: map[string]string{
				"v33.9.0":  "active",
				"v33.10.0": "active",
			},
			expected: "33.9.0",
		},
		{
			name: "case 7: patch versions are ordered correctly",
			releases: map[string]string{
				"v33.1.2":  "active",
				"v33.1.10": "active",
			},
			expected: "33.1.2",
		},
		{
			// The situation that caused https://github.com/giantswarm/roadmap/issues/4325:
			// clusters lagging on v33.1.1 while v33.2.0 was deprecated as unused.
			name: "case 8: deprecated intermediate release stays above the floor",
			releases: map[string]string{
				"v33.1.1": "active",
				"v33.2.0": "deprecated",
				"v33.3.0": "deprecated",
				"v33.4.0": "active",
			},
			expected: "33.1.1",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// The script resolves providers relative to the working directory, so build
			// the fixture provider inside a temporary directory and run from there.
			workDir := t.TempDir()
			providerDir := filepath.Join(workDir, "capa")
			if err := os.MkdirAll(providerDir, 0o755); err != nil {
				t.Fatal(err)
			}

			for version, state := range tc.releases {
				writeFixtureRelease(t, providerDir, version, state, state == "archived")
			}

			cmd := exec.Command(scriptPath, "capa")
			cmd.Dir = workDir
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Fatalf("running %s failed: %v\noutput: %s", lowestActiveReleaseScript, err, output)
			}

			actual := strings.TrimSpace(string(output))
			if actual != tc.expected {
				t.Errorf("expected archive floor %q, got %q", tc.expected, actual)
			}
		})
	}
}

// Test_ArchiveFloorHoldsForProviders checks that the archive floor can be determined
// for every provider that has releases, so the archive workflow never falls back to
// archiving without a floor by accident.
func Test_ArchiveFloorHoldsForProviders(t *testing.T) {
	providers := []string{"azure", "capa", "vsphere", "eks", "proxmox", "aks"}

	for _, provider := range providers {
		t.Run(provider, func(t *testing.T) {
			releases, err := findReleases(provider, false)
			if err != nil {
				t.Fatal(err)
			}

			hasActive := false
			for _, release := range releases {
				if release.Spec.State == "active" {
					hasActive = true
					break
				}
			}

			output, err := exec.Command(lowestActiveReleaseScript, provider).CombinedOutput()
			if err != nil {
				t.Fatalf("running %s failed: %v\noutput: %s", lowestActiveReleaseScript, err, output)
			}
			floor := strings.TrimSpace(string(output))

			if hasActive && floor == "" {
				t.Errorf("provider %s has active releases but no archive floor was determined", provider)
			}
			if !hasActive && floor != "" {
				t.Errorf("provider %s has no active releases but archive floor %q was determined", provider, floor)
			}
		})
	}
}
