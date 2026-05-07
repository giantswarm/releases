package breakingchanges

import "sort"

// TeamIssuesPayload is the JSON shape consumed by the issue-creation workflow
// step. It is emitted by the breakingchanges tool so the action only has to
// render markdown — the detection logic stays in Go.
type TeamIssuesPayload struct {
	Version           string       `json:"version"`            // Release version (any of the analyzed releases — used for label-based dedupe)
	KubernetesVersion string       `json:"kubernetes_version"` // Target Kubernetes version
	Teams             []TeamIssues `json:"teams"`              // One entry per team with at least one finding
}

// TeamIssues is the per-team grouping in the payload.
type TeamIssues struct {
	Team string             `json:"team"` // CODEOWNERS-resolved team handle (e.g. "team-phoenix")
	Apps []TeamIssuesAppRef `json:"apps"` // Apps owned by this team that need attention
}

// TeamIssuesAppRef is a single app on a team's checklist.
type TeamIssuesAppRef struct {
	Name     string `json:"name"`     // App name (release.yaml app name)
	Version  string `json:"version"`  // Currently pinned version
	Severity string `json:"severity"` // critical/high/medium/low
	Reason   string `json:"reason"`   // Short explanation (taken from the Finding title)
	Action   string `json:"action"`   // Recommended action
}

// BuildTeamIssuesPayload turns the raw FindingWithProvider list into a
// team-grouped, app-deduplicated payload suitable for issue creation. Any
// finding without a Team field is excluded — those don't have an owner to
// route to. Apps are deduplicated per team (multiple providers sharing the
// same app produce one checklist item per team).
func BuildTeamIssuesPayload(findings []FindingWithProvider, kubernetesVersion string) TeamIssuesPayload {
	type appKey struct{ team, app string }
	seen := make(map[appKey]bool)
	teamApps := make(map[string][]TeamIssuesAppRef)

	var version string
	for _, fwp := range findings {
		f := fwp.Finding
		if f.Team == "" || f.AppName == "" {
			continue
		}
		key := appKey{team: f.Team, app: f.AppName}
		if seen[key] {
			continue
		}
		seen[key] = true

		if version == "" {
			version = fwp.Version
		}

		teamApps[f.Team] = append(teamApps[f.Team], TeamIssuesAppRef{
			Name:     f.AppName,
			Version:  versionFromFinding(f),
			Severity: f.Severity,
			Reason:   f.Title,
			Action:   f.Action,
		})
	}

	teams := make([]TeamIssues, 0, len(teamApps))
	for team, apps := range teamApps {
		sort.Slice(apps, func(i, j int) bool { return apps[i].Name < apps[j].Name })
		teams = append(teams, TeamIssues{Team: team, Apps: apps})
	}
	sort.Slice(teams, func(i, j int) bool { return teams[i].Team < teams[j].Team })

	return TeamIssuesPayload{
		Version:           version,
		KubernetesVersion: kubernetesVersion,
		Teams:             teams,
	}
}

// versionFromFinding tries to extract the app version from the finding's
// raw text or component. The kubeVersion findings contain the version in the
// title (e.g. "`cilium` v1.3.4 does not declare ..."). Fall back to "" so the
// caller renders something sensible.
func versionFromFinding(f Finding) string {
	// The Title carries "`<app>` v<version>" — extract the v<version> token.
	// Keep it simple; if extraction fails, return empty and let the consumer
	// omit the version from its output.
	parts := splitOnFirst(f.Title, " v")
	if len(parts) < 2 {
		return ""
	}
	rest := parts[1]
	// Stop at first space or backtick to isolate the version token.
	for i, r := range rest {
		if r == ' ' || r == '`' {
			return rest[:i]
		}
	}
	return rest
}

func splitOnFirst(s, sep string) []string {
	idx := indexOf(s, sep)
	if idx < 0 {
		return []string{s}
	}
	return []string{s[:idx], s[idx+len(sep):]}
}

func indexOf(s, sub string) int {
	if len(sub) == 0 || len(sub) > len(s) {
		return -1
	}
	for i := 0; i+len(sub) <= len(s); i++ {
		if s[i:i+len(sub)] == sub {
			return i
		}
	}
	return -1
}
