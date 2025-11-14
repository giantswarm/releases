package pkg

import (
	"fmt"
	"sort"
	"strings"
)

// ProviderWarnings holds warnings for a specific provider
type ProviderWarnings struct {
	Provider string
	Warnings []VersionWarning
}

func GenerateReport(warnings []VersionWarning, currentVersion, nextMajorVersion string) string {
	var sb strings.Builder

	sb.WriteString("## ⚠️ Upgrade Path Warning\n\n")
	sb.WriteString(fmt.Sprintf("The following components/apps in **%s** have versions that are **newer** than those in **%s** (the next major release on the upgrade path).\n\n", currentVersion, nextMajorVersion))
	sb.WriteString("This means customers upgrading from **" + currentVersion + "** to **" + nextMajorVersion + "** would experience component/app **downgrades**, which may cause issues.\n\n")

	// Sort warnings by type and name
	sort.Slice(warnings, func(i, j int) bool {
		if warnings[i].ItemType != warnings[j].ItemType {
			return warnings[i].ItemType < warnings[j].ItemType
		}
		return warnings[i].Name < warnings[j].Name
	})

	// Group by type
	appWarnings := []VersionWarning{}
	componentWarnings := []VersionWarning{}

	for _, w := range warnings {
		if w.ItemType == "App" {
			appWarnings = append(appWarnings, w)
		} else {
			componentWarnings = append(componentWarnings, w)
		}
	}

	// Show apps first, then components
	if len(appWarnings) > 0 {
		sb.WriteString("### Apps\n\n")
		sb.WriteString("| App | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----|----------------------------------------|----------------------------------------|--------|\n")
		for _, w := range appWarnings {
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n")
	}

	if len(componentWarnings) > 0 {
		sb.WriteString("### Components\n\n")
		sb.WriteString("| Component | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----------|----------------------------------------|----------------------------------------|--------|\n")
		for _, w := range componentWarnings {
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("> **Note:** This check was triggered because components/apps were automatically bumped to their latest versions.\n")
	sb.WriteString("> If this was intentional, you can ignore this warning and proceed with appropriate documentation.\n")

	return sb.String()
}

// GenerateConsolidatedReport generates a consolidated report across multiple providers
func GenerateConsolidatedReport(allWarnings []ProviderWarnings, currentVersion, nextMajorVersion string) string {
	if len(allWarnings) == 0 {
		return ""
	}

	var sb strings.Builder

	sb.WriteString("## ⚠️ Upgrade Path Warning\n\n")
	sb.WriteString(fmt.Sprintf("The following components/apps in **%s** have versions that are **newer** than those in **%s** (the next major release on the upgrade path).\n\n", currentVersion, nextMajorVersion))
	sb.WriteString("This means customers upgrading from **" + currentVersion + "** to **" + nextMajorVersion + "** would experience component/app **downgrades**, which may cause issues.\n\n")

	// Collect all apps and components (no distinction between shared/provider-specific)
	allApps := make(map[string]VersionWarning)
	allComponents := make(map[string]VersionWarning)

	for _, pw := range allWarnings {
		for _, w := range pw.Warnings {
			if w.ItemType == "App" {
				// Only add if not already present (de-duplicate across providers)
				if _, exists := allApps[w.Name]; !exists {
					allApps[w.Name] = w
				}
			} else {
				if _, exists := allComponents[w.Name]; !exists {
					allComponents[w.Name] = w
				}
			}
		}
	}

	// Sort apps and components by name
	appNames := make([]string, 0, len(allApps))
	for name := range allApps {
		appNames = append(appNames, name)
	}
	sort.Strings(appNames)

	componentNames := make([]string, 0, len(allComponents))
	for name := range allComponents {
		componentNames = append(componentNames, name)
	}
	sort.Strings(componentNames)

	// Generate Apps section
	if len(allApps) > 0 {
		sb.WriteString("### Apps\n\n")
		sb.WriteString("| App | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----|----------------------------------------|----------------------------------------|--------|\n")
		for _, name := range appNames {
			w := allApps[name]
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n")
	}

	// Generate Components section
	if len(allComponents) > 0 {
		sb.WriteString("### Components\n\n")
		sb.WriteString("| Component | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----------|----------------------------------------|----------------------------------------|--------|\n")
		for _, name := range componentNames {
			w := allComponents[name]
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("> **Note:** This check was triggered because components/apps were automatically bumped to their latest versions.\n")
	sb.WriteString("> If this was intentional, you can ignore this warning and proceed with appropriate documentation.\n")

	return sb.String()
}
