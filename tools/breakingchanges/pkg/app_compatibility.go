package breakingchanges

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/Masterminds/semver/v3"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/releases/sdk/api/v1alpha1"
)

// implicitlyK8sCoupledApps lists apps known to need a new release for every
// Kubernetes minor bump but which don't (yet) declare kubeVersion in their
// chart. Each entry should be temporary: the long-term fix is to add a
// kubeVersion declaration to the corresponding chart and remove the entry.
var implicitlyK8sCoupledApps = map[string]bool{
	"cloud-provider-aws":             true,
	"azure-cloud-controller-manager": true,
	"azure-cloud-node-manager":       true,
	"cloud-provider-vsphere":         true,
	"cloud-provider-cloud-director":  true,
	"cloud-provider-proxmox":         true,
}

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

	// Detect whether this release crosses a Kubernetes minor boundary.
	// The fallback list (implicitlyK8sCoupledApps) only fires on minor bumps
	// since cloud-controllers don't need a new release for a patch upgrade.
	isMinorBump := false
	for _, vc := range extractKubernetesVersion(release.README) {
		fromMinor := parseK8sMinorVersion(vc.FromVersion)
		toMinor := parseK8sMinorVersion(vc.ToVersion)
		if fromMinor != 0 && toMinor != 0 && fromMinor != toMinor {
			isMinorBump = true
		}
		break
	}

	fmt.Printf("\nChecking app kubeVersion compatibility against Kubernetes v%s (minor bump: %v)...\n", target.String(), isMinorBump)

	var findings []Finding
	for _, app := range rel.Spec.Apps {
		findings = append(findings, d.checkAppCharts(ctx, app, target, isMinorBump)...)
	}
	fmt.Printf("✓ App compatibility check complete: %d finding(s)\n", len(findings))
	return findings
}

func (d *Detector) checkAppCharts(ctx context.Context, app v1alpha1.ReleaseSpecApp, target *semver.Version, isK8sMinorBump bool) []Finding {
	chart := d.fetchGSAppChart(ctx, app.Name, app.Version)

	var findings []Finding
	hasDeclaredConstraint := false

	if chart != nil && chart.KubeVersion != "" {
		hasDeclaredConstraint = true
		if f := evaluateKubeVersionConstraint(chart.KubeVersion, target, app.Name, app.Version, "Helm chart kubeVersion (giantswarm app)"); f != nil {
			findings = append(findings, *f)
		}
	}
	if chart != nil {
		for _, dep := range chart.Dependencies {
			if !isExternalHelmRepo(dep.Repository) {
				continue
			}
			depKubeVer := d.fetchUpstreamKubeVersion(ctx, dep.Repository, dep.Name, dep.Version)
			if depKubeVer == "" {
				continue
			}
			hasDeclaredConstraint = true
			component := fmt.Sprintf("%s (upstream %s)", app.Name, dep.Name)
			source := fmt.Sprintf("Upstream chart kubeVersion (%s %s from %s)", dep.Name, dep.Version, dep.Repository)
			if f := evaluateKubeVersionConstraint(depKubeVer, target, component, dep.Version, source); f != nil {
				findings = append(findings, *f)
			}
		}
	}

	// Fallback for apps known to be implicitly coupled to Kubernetes that
	// don't (yet) declare kubeVersion. Only fires on minor bumps.
	if !hasDeclaredConstraint && isK8sMinorBump && implicitlyK8sCoupledApps[app.Name] {
		findings = append(findings, makeImplicitlyCoupledFinding(app, target))
	}

	// Resolve owning team for any findings produced for this app.
	if len(findings) > 0 {
		team := d.fetchAppTeam(ctx, app.Name)
		for i := range findings {
			findings[i].Team = team
			findings[i].AppName = app.Name
		}
	}

	return findings
}

// makeImplicitlyCoupledFinding produces a finding for an app on the implicit
// fallback list when no declared kubeVersion was found and the release crosses
// a Kubernetes minor.
func makeImplicitlyCoupledFinding(app v1alpha1.ReleaseSpecApp, target *semver.Version) Finding {
	return Finding{
		Severity:    "high",
		Component:   app.Name,
		Title:       fmt.Sprintf("`%s` v%s needs a new release for Kubernetes v%s", app.Name, app.Version, target.String()),
		Description: fmt.Sprintf("`%s` is implicitly coupled to Kubernetes (no declared `kubeVersion`) and a new release is required for every Kubernetes minor bump.", app.Name),
		Impact:      fmt.Sprintf("Cluster components managed by `%s` may not function correctly on Kubernetes v%s.", app.Name, target.String()),
		Action:      fmt.Sprintf("Bump `%s` to a release built against Kubernetes v%s; long-term, add a `kubeVersion` to the chart so this fallback can be removed.", app.Name, target.String()),
		Confidence:  "high",
		Source:      "Implicit Kubernetes coupling (fallback list)",
		RawText:     fmt.Sprintf("app %s has no declared kubeVersion", app.Name),
	}
}

// fetchAppTeam looks up the owning team for an app by reading CODEOWNERS in
// the giantswarm app repo. The first @giantswarm/team-* handle wins (the GS
// convention is one team per repo). Returns "" when no team can be resolved.
func (d *Detector) fetchAppTeam(ctx context.Context, appName string) string {
	if cached, ok := d.teamCache[appName]; ok {
		return cached
	}

	for _, repo := range d.candidateRepositories(appName) {
		for _, branch := range []string{"main", "master"} {
			for _, path := range []string{"CODEOWNERS", ".github/CODEOWNERS", "docs/CODEOWNERS"} {
				url := fmt.Sprintf("https://raw.githubusercontent.com/giantswarm/%s/%s/%s", repo, branch, path)
				body, err := d.fetchURL(ctx, url)
				if err != nil {
					continue
				}
				if team := parseCodeownersTeam(body); team != "" {
					d.teamCache[appName] = team
					return team
				}
			}
		}
	}

	d.teamCache[appName] = ""
	return ""
}

// codeownersTeamRe matches the first @giantswarm/team-* handle on a line.
var codeownersTeamRe = regexp.MustCompile(`@giantswarm/(team-[A-Za-z0-9_-]+)`)

// parseCodeownersTeam extracts the first @giantswarm/team-* handle from a
// CODEOWNERS file body. Comment lines and blank lines are skipped.
func parseCodeownersTeam(content string) string {
	for _, line := range strings.Split(content, "\n") {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}
		if match := codeownersTeamRe.FindStringSubmatch(trimmed); len(match) >= 2 {
			return match[1]
		}
	}
	return ""
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
