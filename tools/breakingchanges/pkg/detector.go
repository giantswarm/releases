package breakingchanges

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
)

// Detector orchestrates the breaking change analysis
type Detector struct {
	anthropicAPIKey      string
	githubToken          string
	httpClient           *http.Client
	debugContextFile     string
	componentMappings    map[string]ComponentMapping
	componentMappingDone bool

	// Caches to avoid redundant fetches in consolidated PRs
	flatcarCache    map[string]string // "fromVer->toVer" -> changelog
	kubernetesCache map[string]string // "fromVer->toVer" -> changelog
	componentCache  map[string]string // "component@fromVer->toVer" -> diff
}

// NewDetector creates a new Detector instance
func NewDetector(anthropicAPIKey, githubToken string) (*Detector, error) {
	if anthropicAPIKey == "" {
		return nil, fmt.Errorf("anthropic API key is required")
	}

	return &Detector{
		anthropicAPIKey:      anthropicAPIKey,
		githubToken:          githubToken,
		httpClient:           &http.Client{},
		debugContextFile:     "",
		componentMappings:    make(map[string]ComponentMapping),
		componentMappingDone: false,
		flatcarCache:         make(map[string]string),
		kubernetesCache:      make(map[string]string),
		componentCache:       make(map[string]string),
	}, nil
}

// SetDebugContextFile sets the path for saving debug context
func (d *Detector) SetDebugContextFile(path string) {
	d.debugContextFile = path
}

// AnalyzeMultipleReleases performs consolidated analysis of multiple releases in a single LLM call
func (d *Detector) AnalyzeMultipleReleases(ctx context.Context, releases []Release) ([]FindingWithProvider, error) {
	fmt.Println("\n=== Analyzing Multiple Releases (Consolidated) ===")

	// Fetch component mappings from devctl once
	if err := d.fetchComponentMappings(ctx); err != nil {
		fmt.Printf("Warning: Could not fetch component mappings: %v\n", err)
	}

	// Collect all version changes from all releases
	allVersionChanges := make(map[string][]VersionChange) // provider -> version changes
	allExternalChanges := make(map[string]string)         // Shared external changelogs (Flatcar, K8s)
	allComponentDiffs := make(map[string]string)          // Deduplicated component diffs

	for _, release := range releases {
		fmt.Printf("\nProcessing %s/%s...\n", release.Provider, release.Version)

		// Extract version changes
		versionChanges := extractVersionChanges(release.README)
		allVersionChanges[release.Provider] = versionChanges
		fmt.Printf("  Found %d version changes\n", len(versionChanges))

		// Collect external changelogs (deduplicated)
		externalChanges, _ := d.fetchExternalChangelogs(ctx, versionChanges)
		for key, content := range externalChanges {
			if _, exists := allExternalChanges[key]; !exists {
				allExternalChanges[key] = content
			}
		}

		// Collect component diffs (deduplicated by component name)
		componentDiffs, _ := d.fetchComponentDiffs(ctx, versionChanges)
		for component, diff := range componentDiffs {
			if _, exists := allComponentDiffs[component]; !exists {
				allComponentDiffs[component] = diff
			}
		}
	}

	fmt.Printf("\n=== Consolidated Summary ===\n")
	fmt.Printf("External changelogs: %d\n", len(allExternalChanges))
	fmt.Printf("Component diffs: %d (deduplicated)\n", len(allComponentDiffs))

	// Build consolidated analysis context
	analysisContext := d.buildConsolidatedContext(releases, allVersionChanges, allExternalChanges, allComponentDiffs)
	fmt.Printf("Consolidated context size: ~%d characters\n", len(analysisContext))

	// Make ONE LLM call
	fmt.Println("\nSending consolidated analysis to AI (this may take 20-40 seconds)...")
	findings, err := d.analyzeWithLLM(ctx, analysisContext)
	if err != nil {
		return nil, fmt.Errorf("LLM analysis failed: %w", err)
	}

	// Attribute findings to providers
	// For consolidated releases, findings apply to all providers unless specifically mentioned
	var findingsWithProviders []FindingWithProvider
	for _, finding := range findings {
		// Check if finding mentions specific provider
		mentionedProvider := d.extractProviderFromFinding(finding)

		if mentionedProvider != "" {
			// Finding is specific to one provider
			for _, release := range releases {
				if release.Provider == mentionedProvider {
					findingsWithProviders = append(findingsWithProviders, FindingWithProvider{
						Finding:  finding,
						Provider: release.Provider,
						Version:  release.Version,
					})
					break
				}
			}
		} else {
			// Finding applies to all providers (shared component like Kubernetes/Flatcar)
			for _, release := range releases {
				findingsWithProviders = append(findingsWithProviders, FindingWithProvider{
					Finding:  finding,
					Provider: release.Provider,
					Version:  release.Version,
				})
			}
		}
	}

	fmt.Printf("\n✓ Consolidated analysis complete. Found %d breaking change(s) affecting %d provider(s)\n",
		len(findings), len(releases))

	return findingsWithProviders, nil
}

// AnalyzeRelease performs the complete analysis of a release
func (d *Detector) AnalyzeRelease(ctx context.Context, release Release) ([]Finding, error) {
	fmt.Println("\n=== Analyzing Release ===")

	// Fetch component mappings from devctl
	if err := d.fetchComponentMappings(ctx); err != nil {
		fmt.Printf("Warning: Could not fetch component mappings: %v\n", err)
	}

	// Extract version changes from README
	versionChanges := d.extractAndLogVersionChanges(release.README)

	// Fetch external changelogs
	externalChanges := d.fetchAndLogExternalChangelogs(ctx, versionChanges)

	// Fetch component diffs from GitHub
	componentDiffs := d.fetchAndLogComponentDiffs(ctx, versionChanges)

	// Build comprehensive context for LLM
	analysisContext := d.buildAndSaveAnalysisContext(release, versionChanges, externalChanges, componentDiffs)

	// Log content summary
	d.logContentSummary(externalChanges, componentDiffs)

	// Send to LLM for analysis
	findings := d.analyzeAndAnnotateFindings(ctx, analysisContext, release)

	fmt.Printf("\n✓ Analysis complete. Found %d breaking change(s)\n", len(findings))
	return findings, nil
}

// extractAndLogVersionChanges extracts version changes and logs them
func (d *Detector) extractAndLogVersionChanges(readme string) []VersionChange {
	fmt.Println("Extracting version changes...")
	versionChanges := extractVersionChanges(readme)
	fmt.Printf("Found %d version changes\n", len(versionChanges))
	for _, vc := range versionChanges {
		fmt.Printf("  - %s: v%s -> v%s\n", vc.Component, vc.FromVersion, vc.ToVersion)
	}
	return versionChanges
}

// fetchAndLogExternalChangelogs fetches external changelogs with error handling
func (d *Detector) fetchAndLogExternalChangelogs(ctx context.Context, versionChanges []VersionChange) map[string]string {
	fmt.Println("\nFetching external changelogs (Flatcar, Kubernetes)...")
	externalChanges, err := d.fetchExternalChangelogs(ctx, versionChanges)
	if err != nil {
		fmt.Printf("Warning: Failed to fetch some external changelogs: %v\n", err)
	}
	fmt.Printf("Fetched %d external changelogs\n", len(externalChanges))
	return externalChanges
}

// fetchAndLogComponentDiffs fetches component diffs with error handling
func (d *Detector) fetchAndLogComponentDiffs(ctx context.Context, versionChanges []VersionChange) map[string]string {
	fmt.Println("\nFetching component diffs...")
	componentDiffs, err := d.fetchComponentDiffs(ctx, versionChanges)
	if err != nil {
		fmt.Printf("Warning: Failed to fetch some component diffs: %v\n", err)
	}
	fmt.Printf("Fetched %d component diffs\n", len(componentDiffs))
	return componentDiffs
}

// buildAndSaveAnalysisContext builds the analysis context and optionally saves it
func (d *Detector) buildAndSaveAnalysisContext(release Release, versionChanges []VersionChange, externalChanges, componentDiffs map[string]string) string {
	fmt.Println("\nBuilding analysis context...")
	analysisContext := buildAnalysisContext(release, versionChanges, externalChanges, componentDiffs)
	fmt.Printf("Context size: ~%d characters\n", len(analysisContext))

	// Save context to file if debug mode enabled
	if d.debugContextFile != "" {
		if err := os.WriteFile(d.debugContextFile, []byte(analysisContext), 0644); err != nil {
			fmt.Printf("Warning: Failed to save debug context: %v\n", err)
		}
	}

	return analysisContext
}

// logContentSummary logs what content was fetched
func (d *Detector) logContentSummary(externalChanges, componentDiffs map[string]string) {
	fmt.Println("\n=== Content Summary ===")

	if flatcarContent, ok := externalChanges["Flatcar"]; ok {
		fmt.Printf("Flatcar changelog: %d chars (preview: %s...)\n", len(flatcarContent), truncate(flatcarContent, 100))
	} else {
		fmt.Println("⚠️  Flatcar changelog: NOT FETCHED")
	}

	if k8sContent, ok := externalChanges["Kubernetes"]; ok {
		fmt.Printf("Kubernetes changelog: %d chars (preview: %s...)\n", len(k8sContent), truncate(k8sContent, 100))
	} else {
		fmt.Println("⚠️  Kubernetes changelog: NOT FETCHED")
	}

	fmt.Printf("Component diffs: %d components\n", len(componentDiffs))
	for name := range componentDiffs {
		fmt.Printf("  - %s\n", name)
	}
}

// analyzeAndAnnotateFindings sends to LLM and adds source info
func (d *Detector) analyzeAndAnnotateFindings(ctx context.Context, analysisContext string, release Release) []Finding {
	fmt.Println("\nSending to AI for analysis (this may take 10-30 seconds)...")
	findings, err := d.analyzeWithLLM(ctx, analysisContext)
	if err != nil {
		fmt.Printf("Error: LLM analysis failed: %v\n", err)
		return nil
	}

	return findings
}

// buildConsolidatedContext builds analysis context for multiple releases
func (d *Detector) buildConsolidatedContext(releases []Release, allVersionChanges map[string][]VersionChange, externalChanges, componentDiffs map[string]string) string {
	ctx := "# Consolidated Release Analysis\n\n"
	ctx += "## Environment\n"
	ctx += "- **Platform**: Linux-only (Flatcar Container Linux)\n"
	ctx += "- **Windows Support**: NO - Do NOT report Windows-related changes\n"
	ctx += "- **Experimental Features**: Generally NOT enabled - Be cautious about alpha/beta API warnings\n\n"

	ctx += fmt.Sprintf("## Analyzing %d provider releases:\n", len(releases))
	for _, release := range releases {
		ctx += fmt.Sprintf("- %s %s\n", release.Provider, release.Version)
	}
	ctx += "\n"

	// Add version changes per provider
	ctx += "## Version Changes by Provider\n\n"
	for provider, versionChanges := range allVersionChanges {
		if len(versionChanges) > 0 {
			ctx += fmt.Sprintf("### %s\n\n", provider)
			for _, vc := range versionChanges {
				ctx += fmt.Sprintf("- %s: %s → %s\n", vc.Component, vc.FromVersion, vc.ToVersion)
			}
			ctx += "\n"
		}
	}

	// Add external changelogs (shared across providers)
	ctx += buildExternalChangelogsSection(externalChanges)

	// Add component diffs (deduplicated)
	ctx += buildComponentDiffsSection(componentDiffs)

	return ctx
}

// extractProviderFromFinding checks if a finding mentions a specific provider
func (d *Detector) extractProviderFromFinding(finding Finding) string {
	// Check if component name contains provider-specific identifier
	component := strings.ToLower(finding.Component)

	providerIdentifiers := map[string]string{
		"aws":                    "aws",
		"capa":                   "aws",
		"cluster-aws":            "aws",
		"azure":                  "azure",
		"capz":                   "azure",
		"cluster-azure":          "azure",
		"vsphere":                "vsphere",
		"capv":                   "vsphere",
		"cluster-vsphere":        "vsphere",
		"cloud-director":         "cloud-director",
		"capvcd":                 "cloud-director",
		"cluster-cloud-director": "cloud-director",
	}

	for identifier, provider := range providerIdentifiers {
		if strings.Contains(component, identifier) {
			return provider
		}
	}

	// Check title and description for provider mentions
	text := strings.ToLower(finding.Title + " " + finding.Description)
	for identifier, provider := range providerIdentifiers {
		if strings.Contains(text, identifier) {
			return provider
		}
	}

	// No specific provider mentioned - applies to all
	return ""
}

// buildAnalysisContext constructs the full context string for LLM analysis
func buildAnalysisContext(release Release, versionChanges []VersionChange, externalChanges, componentDiffs map[string]string) string {
	ctx := fmt.Sprintf("# Release Analysis: %s %s\n\n", release.Provider, release.Version)

	// Add environment context
	ctx += "## Environment\n"
	ctx += "- **Platform**: Linux-only (Flatcar Container Linux)\n"
	ctx += "- **Windows Support**: NO - Do NOT report Windows-related changes\n"
	ctx += "- **Experimental Features**: Generally NOT enabled - Be cautious about alpha/beta API warnings\n\n"

	ctx += buildReadmeSection(release.README)
	ctx += buildVersionChangesSection(versionChanges)
	ctx += buildExternalChangelogsSection(externalChanges)
	ctx += buildComponentDiffsSection(componentDiffs)
	return ctx
}

// buildReadmeSection builds the README section
func buildReadmeSection(readme string) string {
	return fmt.Sprintf("## Release README\n\n%s\n\n", readme)
}

// buildVersionChangesSection builds the version changes section
func buildVersionChangesSection(versionChanges []VersionChange) string {
	if len(versionChanges) == 0 {
		return ""
	}

	result := "## Component Version Changes\n\n"
	for _, vc := range versionChanges {
		result += fmt.Sprintf("- %s: %s → %s\n", vc.Component, vc.FromVersion, vc.ToVersion)
	}
	return result + "\n"
}

// buildExternalChangelogsSection builds the external changelogs section
func buildExternalChangelogsSection(externalChanges map[string]string) string {
	if len(externalChanges) == 0 {
		return ""
	}

	result := "## External Changelogs\n\n"
	result += "**IMPORTANT**: Extract SPECIFIC breaking changes from these changelogs, don't just say 'review the changelog'.\n\n"

	for component, changelog := range externalChanges {
		result += fmt.Sprintf("### %s Changelog\n\n", component)

		// Different limits for different components
		limit := 10000
		if component == "Kubernetes" {
			limit = 50000 // K8s changelogs need more space for Urgent Upgrade Notes and detailed API changes
		}

		if len(changelog) > limit {
			result += changelog[:limit] + "\n\n[...truncated...]\n\n"
		} else {
			result += changelog + "\n\n"
		}
	}

	return result
}

// buildComponentDiffsSection builds the component diffs section
func buildComponentDiffsSection(componentDiffs map[string]string) string {
	if len(componentDiffs) == 0 {
		return ""
	}

	result := "## Component Diffs\n\n"
	for component, diff := range componentDiffs {
		result += fmt.Sprintf("### %s Diff\n\n", component)

		// Limit diff size to avoid token limits
		if len(diff) > 5000 {
			result += diff[:5000] + "\n\n[...truncated...]\n\n"
		} else {
			result += diff + "\n\n"
		}
	}

	return result
}

// truncate truncates a string to maxLen characters
func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen]
}
