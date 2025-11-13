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
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Apps</summary>\n\n")
		sb.WriteString("| App | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----|----------------------------------------|----------------------------------------|--------|\n")
		for _, w := range appWarnings {
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n</details>\n\n")
	}

	if len(componentWarnings) > 0 {
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Components</summary>\n\n")
		sb.WriteString("| Component | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----------|----------------------------------------|----------------------------------------|--------|\n")
		for _, w := range componentWarnings {
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n</details>\n\n")
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

	providerNames := map[string]string{
		"aws":            "AWS",
		"azure":          "Azure",
		"cloud-director": "Cloud Director",
		"eks":            "EKS",
		"vsphere":        "vSphere",
	}

	var sb strings.Builder

	sb.WriteString("## ⚠️ Upgrade Path Warning\n\n")
	sb.WriteString(fmt.Sprintf("The following components/apps in **%s** have versions that are **newer** than those in **%s** (the next major release on the upgrade path).\n\n", currentVersion, nextMajorVersion))
	sb.WriteString("This means customers upgrading from **" + currentVersion + "** to **" + nextMajorVersion + "** would experience component/app **downgrades**, which may cause issues.\n\n")

	// Collect all items and count occurrences across providers
	appCounts := make(map[string]int)
	componentCounts := make(map[string]int)
	appVersions := make(map[string]VersionWarning)
	componentVersions := make(map[string]VersionWarning)

	totalProviders := len(allWarnings)

	for _, pw := range allWarnings {
		for _, w := range pw.Warnings {
			if w.ItemType == "App" {
				appCounts[w.Name]++
				if _, exists := appVersions[w.Name]; !exists {
					appVersions[w.Name] = w
				}
			} else {
				componentCounts[w.Name]++
				if _, exists := componentVersions[w.Name]; !exists {
					componentVersions[w.Name] = w
				}
			}
		}
	}

	// Separate shared vs provider-specific
	sharedApps := []string{}
	sharedComponents := []string{}
	providerSpecificApps := make(map[string][]VersionWarning)
	providerSpecificComponents := make(map[string][]VersionWarning)

	for name, count := range appCounts {
		if count == totalProviders {
			sharedApps = append(sharedApps, name)
		} else {
			// Provider-specific - collect which providers have it
			for _, pw := range allWarnings {
				for _, w := range pw.Warnings {
					if w.Name == name && w.ItemType == "App" {
						providerSpecificApps[pw.Provider] = append(providerSpecificApps[pw.Provider], w)
					}
				}
			}
		}
	}

	for name, count := range componentCounts {
		if count == totalProviders {
			sharedComponents = append(sharedComponents, name)
		} else {
			// Provider-specific - collect which providers have it
			for _, pw := range allWarnings {
				for _, w := range pw.Warnings {
					if w.Name == name && w.ItemType == "Component" {
						providerSpecificComponents[pw.Provider] = append(providerSpecificComponents[pw.Provider], w)
					}
				}
			}
		}
	}

	sort.Strings(sharedApps)
	sort.Strings(sharedComponents)

	// Generate report sections
	if len(sharedApps) > 0 {
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Apps (All Providers)</summary>\n\n")
		sb.WriteString("| App | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----|----------------------------------------|----------------------------------------|--------|\n")
		for _, name := range sharedApps {
			w := appVersions[name]
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n</details>\n\n")
	}

	if len(sharedComponents) > 0 {
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Components (All Providers)</summary>\n\n")
		sb.WriteString("| Component | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
		sb.WriteString("|-----------|----------------------------------------|----------------------------------------|--------|\n")
		for _, name := range sharedComponents {
			w := componentVersions[name]
			sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
				w.Name, w.CurrentVersion, w.NextMajorVersion))
		}
		sb.WriteString("\n</details>\n\n")
	}

	if len(providerSpecificApps) > 0 {
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Provider-Specific Apps</summary>\n\n")

		// Sort providers for consistent output
		providers := make([]string, 0, len(providerSpecificApps))
		for p := range providerSpecificApps {
			providers = append(providers, p)
		}
		sort.Strings(providers)

		for _, provider := range providers {
			warnings := providerSpecificApps[provider]
			if len(warnings) == 0 {
				continue
			}

			// Sort warnings by name
			sort.Slice(warnings, func(i, j int) bool {
				return warnings[i].Name < warnings[j].Name
			})

			providerLabel := providerNames[provider]
			if providerLabel == "" {
				providerLabel = strings.ToUpper(provider)
			}

			sb.WriteString(fmt.Sprintf("**%s:**\n\n", providerLabel))
			sb.WriteString("| App | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
			sb.WriteString("|-----|----------------------------------------|----------------------------------------|--------|\n")

			for _, w := range warnings {
				sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
					w.Name, w.CurrentVersion, w.NextMajorVersion))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("</details>\n\n")
	}

	if len(providerSpecificComponents) > 0 {
		sb.WriteString("<details>\n")
		sb.WriteString("<summary>Provider-Specific Components</summary>\n\n")

		// Sort providers for consistent output
		providers := make([]string, 0, len(providerSpecificComponents))
		for p := range providerSpecificComponents {
			providers = append(providers, p)
		}
		sort.Strings(providers)

		for _, provider := range providers {
			warnings := providerSpecificComponents[provider]
			if len(warnings) == 0 {
				continue
			}

			// Sort warnings by name
			sort.Slice(warnings, func(i, j int) bool {
				return warnings[i].Name < warnings[j].Name
			})

			providerLabel := providerNames[provider]
			if providerLabel == "" {
				providerLabel = strings.ToUpper(provider)
			}

			sb.WriteString(fmt.Sprintf("**%s:**\n\n", providerLabel))
			sb.WriteString("| Component | This Release (" + currentVersion + ") | Next Major (" + nextMajorVersion + ") | Impact |\n")
			sb.WriteString("|-----------|----------------------------------------|----------------------------------------|--------|\n")

			for _, w := range warnings {
				sb.WriteString(fmt.Sprintf("| `%s` | **%s** | %s | ⬇️ Downgrade on upgrade |\n",
					w.Name, w.CurrentVersion, w.NextMajorVersion))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("</details>\n\n")
	}

	sb.WriteString("> **Note:** This check was triggered because components/apps were automatically bumped to their latest versions.\n")
	sb.WriteString("> If this was intentional, you can ignore this warning and proceed with appropriate documentation.\n")

	return sb.String()
}
