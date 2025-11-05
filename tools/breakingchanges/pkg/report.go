package breakingchanges

import (
	"fmt"
	"sort"
	"strings"
)

// FindingWithProvider wraps a Finding with provider context
type FindingWithProvider struct {
	Finding  Finding
	Provider string
	Version  string
}

// GroupedFinding represents a finding that may affect multiple providers
type GroupedFinding struct {
	Finding   Finding
	Providers map[string]string // provider -> version
}

// GenerateReport creates a markdown report from findings
func GenerateReport(findings []Finding, prNumber, commitSHA string) string {
	if len(findings) == 0 {
		return generateNoFindingsReport(prNumber, commitSHA)
	}

	var sb strings.Builder

	// Header
	sb.WriteString("## 🤖 Breaking Change Analysis\n\n")
	sb.WriteString(fmt.Sprintf("**PR:** #%s | **Commit:** `%s`\n\n", prNumber, commitSHA[:8]))

	// Summary
	critical := countBySeverity(findings, "critical")
	high := countBySeverity(findings, "high")
	medium := countBySeverity(findings, "medium")
	low := countBySeverity(findings, "low")

	sb.WriteString("### 📊 Summary\n\n")
	sb.WriteString(fmt.Sprintf("Found **%d** potential breaking change(s):\n\n", len(findings)))

	if critical > 0 {
		sb.WriteString(fmt.Sprintf("- 🔴 **%d Critical** - Action required before upgrade\n", critical))
	}
	if high > 0 {
		sb.WriteString(fmt.Sprintf("- 🟠 **%d High** - Review strongly recommended\n", high))
	}
	if medium > 0 {
		sb.WriteString(fmt.Sprintf("- 🟡 **%d Medium** - Review recommended\n", medium))
	}
	if low > 0 {
		sb.WriteString(fmt.Sprintf("- 🟢 **%d Low** - Informational\n", low))
	}

	sb.WriteString("\n---\n\n")

	// Group findings by severity
	severityOrder := []string{"critical", "high", "medium", "low"}
	for _, severity := range severityOrder {
		findingsForSeverity := filterBySeverity(findings, severity)
		if len(findingsForSeverity) == 0 {
			continue
		}

		sb.WriteString(fmt.Sprintf("### %s %s Priority\n\n", getSeverityEmoji(severity), strings.Title(severity)))

		for i, finding := range findingsForSeverity {
			sb.WriteString(formatFinding(finding, i+1))
			sb.WriteString("\n")
		}
	}

	// Footer
	sb.WriteString("\n---\n\n")
	sb.WriteString("### ℹ️ What to do next\n\n")
	sb.WriteString("1. **Review** each breaking change above\n")
	sb.WriteString("2. **Verify** if the breaking change applies to your environment\n")
	sb.WriteString("3. **Test** critical and high severity changes in a dev environment before rolling out\n")
	sb.WriteString("4. **Document** confirmed breaking changes in the release README and migration guide\n")
	sb.WriteString("5. **Notify** customers about breaking changes in release announcements\n\n")

	sb.WriteString("> ⚠️ **Note:** This analysis is automated and may include false positives or miss breaking changes. ")
	sb.WriteString("Always review the full changelogs and test in a non-production environment.\n\n")

	sb.WriteString("<details>\n<summary>ℹ️ About this analysis</summary>\n\n")
	sb.WriteString("This automated breaking change detector uses AI to analyze:\n")
	sb.WriteString("- Kubernetes and Flatcar changelogs (especially Urgent Upgrade Notes)\n")
	sb.WriteString("- Component version changes and GitHub diffs\n")
	sb.WriteString("- Configuration changes in values.yaml and manifests\n")
	sb.WriteString("- RBAC, CRD, and API changes\n\n")
	sb.WriteString("The AI looks for breaking changes such as:\n")
	sb.WriteString("- **Removed or renamed APIs/CRDs** (will fail immediately)\n")
	sb.WriteString("- **Required configuration fields** (pods won't start)\n")
	sb.WriteString("- **RBAC permission changes** (authorization failures)\n")
	sb.WriteString("- **Security context changes** (pod security violations)\n")
	sb.WriteString("- **Behavior changes** (unexpected failures or degraded performance)\n")
	sb.WriteString("- **Default value changes** (resource constraints, timeouts, etc.)\n\n")
	sb.WriteString("Severity indicates likelihood of customer impact:\n")
	sb.WriteString("- 🔴 **Critical**: WILL break customer workloads without action\n")
	sb.WriteString("- 🟠 **High**: LIKELY to break depending on customer setup\n")
	sb.WriteString("- 🟡 **Medium**: MIGHT break in specific scenarios\n")
	sb.WriteString("- 🟢 **Low**: Low risk, but worth customer awareness\n\n")
	sb.WriteString("</details>\n")

	return sb.String()
}

func generateNoFindingsReport(prNumber, commitSHA string) string {
	var sb strings.Builder

	sb.WriteString("## Breaking Change Analysis\n\n")
	sb.WriteString(fmt.Sprintf("**PR:** #%s | **Commit:** `%s`\n\n", prNumber, commitSHA[:8]))
	sb.WriteString("### No Breaking Changes Detected\n\n")
	sb.WriteString("The automated analysis did not identify any obvious breaking changes in this release.\n\n")
	sb.WriteString("> **Note:** This does not guarantee there are no breaking changes. Please still:\n")
	sb.WriteString("> - Review component changelogs for behavior changes\n")
	sb.WriteString("> - Check Kubernetes and Flatcar upgrade notes\n")
	sb.WriteString("> - Test in a non-production environment, especially for major/minor version bumps\n")

	return sb.String()
}

func formatFinding(f Finding, index int) string {
	var sb strings.Builder

	// Finding header
	sb.WriteString(fmt.Sprintf("<details>\n<summary><strong>%d. %s</strong> ", index, f.Title))
	sb.WriteString(fmt.Sprintf("<code>%s</code> ", f.Component))
	sb.WriteString(fmt.Sprintf("<em>(%s confidence)</em>", f.Confidence))
	sb.WriteString("</summary>\n\n")

	// Description
	sb.WriteString("#### 📝 Description\n\n")
	sb.WriteString(f.Description)
	sb.WriteString("\n\n")

	// Impact
	if f.Impact != "" {
		sb.WriteString("#### 👥 Impact\n\n")
		sb.WriteString(f.Impact)
		sb.WriteString("\n\n")
	}

	// Required Actions
	if f.Action != "" {
		sb.WriteString("#### ✅ Required Actions\n\n")
		sb.WriteString(f.Action)
		sb.WriteString("\n\n")
	}

	// Raw text (in code block for context)
	if f.RawText != "" {
		sb.WriteString("#### 📄 Original Text\n\n")
		sb.WriteString("```\n")
		// Truncate if too long
		if len(f.RawText) > 500 {
			sb.WriteString(f.RawText[:500])
			sb.WriteString("\n[...truncated...]\n")
		} else {
			sb.WriteString(f.RawText)
			sb.WriteString("\n")
		}
		sb.WriteString("```\n\n")
	}

	sb.WriteString("</details>\n")

	return sb.String()
}

func getSeverityEmoji(severity string) string {
	switch severity {
	case "critical":
		return "🔴"
	case "high":
		return "🟠"
	case "medium":
		return "🟡"
	case "low":
		return "🟢"
	default:
		return "⚪"
	}
}

func countBySeverity(findings []Finding, severity string) int {
	count := 0
	for _, f := range findings {
		if strings.EqualFold(f.Severity, severity) {
			count++
		}
	}
	return count
}

func filterBySeverity(findings []Finding, severity string) []Finding {
	var result []Finding
	for _, f := range findings {
		if strings.EqualFold(f.Severity, severity) {
			result = append(result, f)
		}
	}
	return result
}

// GenerateReportWithProviders creates a report that groups findings by provider and deduplicates shared findings
func GenerateReportWithProviders(findingsWithProviders []FindingWithProvider, prNumber, commitSHA string) string {
	if len(findingsWithProviders) == 0 {
		return generateNoFindingsReport(prNumber, commitSHA)
	}

	// Group findings by their unique key (title + component)
	findingGroups := make(map[string]*GroupedFinding)
	for _, fwp := range findingsWithProviders {
		// Create a key based on title and component to identify duplicates
		key := fmt.Sprintf("%s|%s", fwp.Finding.Title, fwp.Finding.Component)

		if existing, ok := findingGroups[key]; ok {
			// Add this provider to existing finding
			existing.Providers[fwp.Provider] = fwp.Version
		} else {
			// New finding
			findingGroups[key] = &GroupedFinding{
				Finding: fwp.Finding,
				Providers: map[string]string{
					fwp.Provider: fwp.Version,
				},
			}
		}
	}

	// Convert back to list
	var groupedFindings []GroupedFinding
	for _, gf := range findingGroups {
		groupedFindings = append(groupedFindings, *gf)
	}

	// Sort by severity
	sort.Slice(groupedFindings, func(i, j int) bool {
		severityOrder := map[string]int{"critical": 0, "high": 1, "medium": 2, "low": 3}
		return severityOrder[strings.ToLower(groupedFindings[i].Finding.Severity)] <
			severityOrder[strings.ToLower(groupedFindings[j].Finding.Severity)]
	})

	var sb strings.Builder

	// Header
	sb.WriteString("## Breaking Change Analysis\n\n")
	sb.WriteString(fmt.Sprintf("**PR:** #%s | **Commit:** `%s`\n\n", prNumber, commitSHA[:min(8, len(commitSHA))]))

	// Summary
	critical := 0
	high := 0
	medium := 0
	low := 0
	for _, gf := range groupedFindings {
		switch strings.ToLower(gf.Finding.Severity) {
		case "critical":
			critical++
		case "high":
			high++
		case "medium":
			medium++
		case "low":
			low++
		}
	}

	sb.WriteString("### Summary\n\n")
	sb.WriteString(fmt.Sprintf("Found **%d** potential breaking change(s):\n\n", len(groupedFindings)))

	if critical > 0 {
		sb.WriteString(fmt.Sprintf("- 🔴 **%d Critical** - Action required before upgrade\n", critical))
	}
	if high > 0 {
		sb.WriteString(fmt.Sprintf("- 🟠 **%d High** - Review strongly recommended\n", high))
	}
	if medium > 0 {
		sb.WriteString(fmt.Sprintf("- 🟡 **%d Medium** - Review recommended\n", medium))
	}
	if low > 0 {
		sb.WriteString(fmt.Sprintf("- 🟢 **%d Low** - Informational\n", low))
	}

	sb.WriteString("\n---\n\n")

	// Group by severity
	severityOrder := []string{"critical", "high", "medium", "low"}
	for _, severity := range severityOrder {
		findings := filterGroupedBySeverity(groupedFindings, severity)
		if len(findings) == 0 {
			continue
		}

		// Severity header
		switch severity {
		case "critical":
			sb.WriteString("### 🔴 Critical\n\n")
		case "high":
			sb.WriteString("### 🟠 High Priority\n\n")
		case "medium":
			sb.WriteString("### 🟡 Medium Priority\n\n")
		case "low":
			sb.WriteString("### 🟢 Low Priority\n\n")
		}

		// List findings
		for i, gf := range findings {
			sb.WriteString(formatGroupedFinding(i+1, gf))

			// Add separator between findings for better readability
			if i < len(findings)-1 {
				sb.WriteString("---\n\n")
			}
		}

		sb.WriteString("\n")
	}

	// Footer
	sb.WriteString("\n\n---\n\n")
	sb.WriteString("> **Note:** This analysis is automated and may include false positives. Please verify each finding before taking action.\n\n")

	// About section
	sb.WriteString("<details>\n")
	sb.WriteString("<summary>About this analysis</summary>\n\n")
	sb.WriteString("This automated analysis uses AI to scan:\n")
	sb.WriteString("- Release changelogs and READMEs\n")
	sb.WriteString("- Component version changes\n")
	sb.WriteString("- External dependency changelogs (Flatcar, Kubernetes)\n")
	sb.WriteString("- GitHub diffs for component changes\n\n")
	sb.WriteString("The AI looks for patterns indicating breaking changes such as:\n")
	sb.WriteString("- Configuration flag removals\n")
	sb.WriteString("- API deprecations and removals\n")
	sb.WriteString("- Default behavior changes\n")
	sb.WriteString("- RBAC and permission changes\n")
	sb.WriteString("- OS-level changes (cgroups, kernel, etc.)\n\n")
	sb.WriteString("</details>\n")

	return sb.String()
}

func formatGroupedFinding(index int, gf GroupedFinding) string {
	var sb strings.Builder
	f := gf.Finding

	// Title with component badge
	confidence := ""
	if f.Confidence != "" {
		confidence = fmt.Sprintf(" · %s confidence", f.Confidence)
	}

	sb.WriteString(fmt.Sprintf("#### ⚠️ %d. %s\n\n", index, f.Title))
	sb.WriteString(fmt.Sprintf("`%s`%s\n\n", f.Component, confidence))

	// Description
	if f.Description != "" {
		sb.WriteString(f.Description)
		sb.WriteString("\n\n")
	}

	// Impact and Action in blockquotes for visual emphasis
	if f.Impact != "" {
		sb.WriteString("> **Impact:** ")
		sb.WriteString(f.Impact)
		sb.WriteString("\n\n")
	}

	if f.Action != "" {
		sb.WriteString("> **Action:** ")
		sb.WriteString(f.Action)
		sb.WriteString("\n\n")
	}

	return sb.String()
}

func filterGroupedBySeverity(findings []GroupedFinding, severity string) []GroupedFinding {
	var result []GroupedFinding
	for _, gf := range findings {
		if strings.EqualFold(gf.Finding.Severity, severity) {
			result = append(result, gf)
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
