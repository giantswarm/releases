package pkg

import (
	"fmt"
	"sort"
	"strings"
)

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
