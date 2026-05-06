package breakingchanges

import (
	"context"
	"fmt"
	"strings"

	"github.com/Masterminds/semver/v3"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/releases/sdk/api/v1alpha1"
)

// App Kubernetes-compatibility check.
//
// For every app in release.yaml we read the Helm chart's `kubeVersion` field
// (a Helm primitive) and evaluate it against the release's target Kubernetes
// version. We also walk one level of `dependencies:` so that giantswarm app
// charts which wrap an upstream chart (cilium-app -> upstream cilium) are
// covered via the upstream's own declaration.
//
// Apps whose chart does not declare a kubeVersion are silently skipped —
// without an authoritative declaration there is no trustworthy signal to
// surface. This avoids needing a hand-maintained whitelist of "K8s-coupled"
// apps: charts that care declare it, charts that don't, won't.

type chartYAML struct {
	Name         string            `json:"name"`
	Version      string            `json:"version"`
	KubeVersion  string            `json:"kubeVersion,omitempty"`
	Dependencies []chartDependency `json:"dependencies,omitempty"`
}

type chartDependency struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	Repository string `json:"repository"`
}

type helmIndexEntry struct {
	Name        string `json:"name"`
	Version     string `json:"version"`
	KubeVersion string `json:"kubeVersion,omitempty"`
}

type helmIndex struct {
	Entries map[string][]helmIndexEntry `json:"entries"`
}

// checkAppKubernetesCompatibility evaluates every app in the release against
// the target Kubernetes version using declared Helm kubeVersion constraints.
func (d *Detector) checkAppKubernetesCompatibility(ctx context.Context, release Release) []Finding {
	if strings.TrimSpace(release.YAML) == "" {
		return nil
	}

	var rel v1alpha1.Release
	if err := yaml.Unmarshal([]byte(release.YAML), &rel); err != nil {
		fmt.Printf("Warning: parse release.yaml for %s/%s: %v\n", release.Provider, release.Version, err)
		return nil
	}

	targetStr, err := rel.GetKubernetesVersion()
	if err != nil {
		fmt.Printf("Warning: get kubernetes version: %v\n", err)
		return nil
	}
	target, err := semver.NewVersion(strings.TrimPrefix(targetStr, "v"))
	if err != nil {
		fmt.Printf("Warning: parse target kubernetes version %q: %v\n", targetStr, err)
		return nil
	}

	fmt.Printf("\nChecking app kubeVersion compatibility against Kubernetes v%s...\n", target.String())

	var findings []Finding
	for _, app := range rel.Spec.Apps {
		findings = append(findings, d.checkAppCharts(ctx, app, target)...)
	}
	fmt.Printf("✓ App compatibility check complete: %d finding(s)\n", len(findings))
	return findings
}

func (d *Detector) checkAppCharts(ctx context.Context, app v1alpha1.ReleaseSpecApp, target *semver.Version) []Finding {
	chart := d.fetchGSAppChart(ctx, app.Name, app.Version)
	if chart == nil {
		return nil
	}

	var findings []Finding
	if chart.KubeVersion != "" {
		if f := evaluateKubeVersionConstraint(chart.KubeVersion, target, app.Name, app.Version, "Helm chart kubeVersion (giantswarm app)"); f != nil {
			findings = append(findings, *f)
		}
	}
	for _, dep := range chart.Dependencies {
		if !isExternalHelmRepo(dep.Repository) {
			continue
		}
		depKubeVer := d.fetchUpstreamKubeVersion(ctx, dep.Repository, dep.Name, dep.Version)
		if depKubeVer == "" {
			continue
		}
		component := fmt.Sprintf("%s (upstream %s)", app.Name, dep.Name)
		source := fmt.Sprintf("Upstream chart kubeVersion (%s %s from %s)", dep.Name, dep.Version, dep.Repository)
		if f := evaluateKubeVersionConstraint(depKubeVer, target, component, dep.Version, source); f != nil {
			findings = append(findings, *f)
		}
	}
	return findings
}

// evaluateKubeVersionConstraint returns a Finding when the chart's kubeVersion
// constraint does not include the target Kubernetes version. A nil return
// means the chart explicitly supports the target.
func evaluateKubeVersionConstraint(kubeVersion string, target *semver.Version, component, version, source string) *Finding {
	constraint, err := semver.NewConstraint(kubeVersion)
	if err != nil {
		return &Finding{
			Severity:    "low",
			Component:   component,
			Title:       fmt.Sprintf("Could not parse Helm kubeVersion constraint for %s", component),
			Description: fmt.Sprintf("kubeVersion %q failed to parse as a semver constraint.", kubeVersion),
			Impact:      "Compatibility with the target Kubernetes version could not be verified automatically.",
			Action:      fmt.Sprintf("Verify manually that `%s` v%s supports Kubernetes v%s.", component, version, target.String()),
			Confidence:  "low",
			Source:      source,
			RawText:     fmt.Sprintf("kubeVersion: %q", kubeVersion),
		}
	}

	if ok, _ := constraint.Validate(target); ok {
		return nil
	}

	return &Finding{
		Severity:    "high",
		Component:   component,
		Title:       fmt.Sprintf("`%s` v%s does not declare support for Kubernetes v%s", component, version, target.String()),
		Description: fmt.Sprintf("Helm chart `kubeVersion` `%q` does not include target Kubernetes v%s.", kubeVersion, target.String()),
		Impact:      fmt.Sprintf("Helm may refuse to deploy `%s` on Kubernetes v%s, or the chart may rely on APIs unavailable on this version.", component, target.String()),
		Action:      fmt.Sprintf("Bump `%s` to a version whose chart supports Kubernetes v%s, or wait for a compatible upstream release.", component, target.String()),
		Confidence:  "high",
		Source:      source,
		RawText:     fmt.Sprintf("kubeVersion: %q", kubeVersion),
	}
}

// isExternalHelmRepo returns true for HTTP(S) Helm repositories. Same-cluster
// (catalog: cluster) and oci:// references are skipped: cluster-catalog charts
// are covered by GS component diffs and OCI registries don't expose index.yaml.
func isExternalHelmRepo(repo string) bool {
	return strings.HasPrefix(repo, "http://") || strings.HasPrefix(repo, "https://")
}

// fetchGSAppChart resolves the giantswarm app chart for a given app name and
// version. The repo name is taken from devctl mappings when available and
// falls back to the conventional `<name>-app` and `<name>` patterns.
func (d *Detector) fetchGSAppChart(ctx context.Context, appName, appVersion string) *chartYAML {
	cacheKey := fmt.Sprintf("%s@%s", appName, appVersion)
	if cached, ok := d.chartCache[cacheKey]; ok {
		return cached
	}

	repoCandidates := d.candidateRepositories(appName)
	chartDirCandidates := []string{appName, appName + "-app"}

	for _, repo := range repoCandidates {
		for _, chartDir := range chartDirCandidates {
			url := fmt.Sprintf("https://raw.githubusercontent.com/giantswarm/%s/v%s/helm/%s/Chart.yaml",
				repo, appVersion, chartDir)
			body, err := d.fetchURL(ctx, url)
			if err != nil {
				continue
			}
			var c chartYAML
			if err := yaml.Unmarshal([]byte(body), &c); err != nil {
				continue
			}
			d.chartCache[cacheKey] = &c
			return &c
		}
	}

	fmt.Printf("  → Could not locate Chart.yaml for app %s v%s — skipping\n", appName, appVersion)
	d.chartCache[cacheKey] = nil
	return nil
}

// candidateRepositories returns the GitHub repo names to try for a given app.
// devctl mappings (already loaded by the detector) are authoritative; the
// `<name>-app` and `<name>` patterns cover everything else.
func (d *Detector) candidateRepositories(appName string) []string {
	var candidates []string
	if mapping, ok := d.componentMappings[appName]; ok && mapping.RepositoryName != "" {
		candidates = append(candidates, mapping.RepositoryName)
	}
	candidates = appendUniqueRepo(candidates, appName+"-app")
	candidates = appendUniqueRepo(candidates, appName)
	return candidates
}

func appendUniqueRepo(slice []string, val string) []string {
	for _, s := range slice {
		if s == val {
			return slice
		}
	}
	return append(slice, val)
}

// fetchUpstreamKubeVersion fetches index.yaml from a Helm repository and
// returns the kubeVersion declared by the entry matching depName + depVersion.
// Returns "" when no kubeVersion is declared or the index can't be parsed.
func (d *Detector) fetchUpstreamKubeVersion(ctx context.Context, repoURL, depName, depVersion string) string {
	indexURL := strings.TrimRight(repoURL, "/") + "/index.yaml"

	index, ok := d.helmIndexCache[indexURL]
	if !ok {
		body, err := d.fetchURL(ctx, indexURL)
		if err != nil {
			fmt.Printf("  → Could not fetch helm index %s: %v\n", indexURL, err)
			d.helmIndexCache[indexURL] = nil
			return ""
		}
		var parsed helmIndex
		if err := yaml.Unmarshal([]byte(body), &parsed); err != nil {
			fmt.Printf("  → Could not parse helm index %s: %v\n", indexURL, err)
			d.helmIndexCache[indexURL] = nil
			return ""
		}
		index = &parsed
		d.helmIndexCache[indexURL] = index
	}
	if index == nil {
		return ""
	}

	targetVer := strings.TrimPrefix(depVersion, "v")
	for _, e := range index.Entries[depName] {
		if strings.TrimPrefix(e.Version, "v") == targetVer {
			return e.KubeVersion
		}
	}
	return ""
}
