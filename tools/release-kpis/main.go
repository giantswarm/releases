// release-kpis collects release KPIs from merged release pull requests.
//
// A release PR is a merged pull request that added a new
// <provider>/vX.Y.Z/release.yaml file on the default branch. Release PRs are
// found in the git history, since PR titles and labels have not been
// consistent over time. For every release PR the lead time (open -> merge) in
// days is recorded. Releases created by the scheduled automation on the 1st of
// the month additionally get a planned merge date (the next 1st of the month)
// and the delay against it.
//
// The result is written as JSON. The Release KPIs workflow embeds it into the
// Grafana dashboard through the Infinity datasource.
//
// See https://github.com/giantswarm/roadmap/issues/4382
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

const (
	// plannedMergeDateMarker is written into the PR body by create-release.yaml
	// for scheduled releases, e.g. "<!-- PLANNED_MERGE_DATE: 2026-10-01 -->".
	plannedMergeDateMarker = "PLANNED_MERGE_DATE"

	// Scheduled release PRs are created by the create-release workflow, which
	// runs at 06:00 UTC on the 1st of every month. PRs created before this
	// marker existed are detected through this window instead.
	scheduledRunDay       = 1
	scheduledRunStartHour = 6
	scheduledRunEndHour   = 9
	automationUser        = "taylorbot"

	// consolidatedProvider is used for PRs that release several providers at once.
	consolidatedProvider = "CAPI"

	// defaultMinMajor is the first major version of the CAPI release process.
	// The same directories also hold vintage releases (e.g. azure/v20.0.0),
	// which are not part of this KPI.
	defaultMinMajor = 25
)

var (
	// providers maps release directories to the provider acronym used in PR titles.
	providers = map[string]string{
		"aks":            "AKS",
		"azure":          "CAPZ",
		"capa":           "CAPA",
		"cloud-director": "CAPVCD",
		"eks":            "EKS",
		"proxmox":        "CAPMOX",
		"vsphere":        "CAPV",
	}

	// Matches "capa/v29.0.0/release.yaml" and "capa/archived/v29.0.0/release.yaml".
	releaseFileRe = regexp.MustCompile(`^([a-z-]+)/(archived/)?v(\d+)\.(\d+)\.(\d+)/release\.yaml$`)

	// Matches the PR number in a squash or merge commit subject, e.g. "CAPI: Release v36.0.0. (#2426)".
	prNumberRe = regexp.MustCompile(`\(#(\d+)\)\s*$`)

	plannedMergeDateRe = regexp.MustCompile(`<!--\s*` + plannedMergeDateMarker + `:\s*(\d{4}-\d{2}-\d{2})\s*-->`)

	// Release PRs carry one of these labels since late 2025. Older release PRs
	// get their type derived from the version instead.
	releaseTypeLabels = map[string]string{
		"release/major": "major",
		"release/minor": "minor",
		"release/patch": "patch",
	}
)

// Release is one row of the KPI data set.
type Release struct {
	Number      int    `json:"number"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Provider    string `json:"provider"`
	Version     string `json:"version"`
	ReleaseType string `json:"release_type"`
	Author      string `json:"author"`

	CreatedAt    time.Time `json:"created_at"`
	MergedAt     time.Time `json:"merged_at"`
	LeadTimeDays float64   `json:"lead_time_days"`

	// Automated is true when the PR was created by the scheduled automation on
	// the 1st of the month. Only those releases have a planned merge date.
	Automated bool `json:"automated"`

	// PlannedMergeDate and DelayDays are only set for automated releases.
	PlannedMergeDate *string  `json:"planned_merge_date"`
	DelayDays        *float64 `json:"delay_days"`
}

// Output is the document written to the output file.
type Output struct {
	GeneratedAt time.Time `json:"generated_at"`
	Repository  string    `json:"repository"`
	Releases    []Release `json:"releases"`
}

// pullRequest is the subset of the GitHub pulls API response we need.
type pullRequest struct {
	Number    int                     `json:"number"`
	Title     string                  `json:"title"`
	Body      string                  `json:"body"`
	HTMLURL   string                  `json:"html_url"`
	CreatedAt time.Time               `json:"created_at"`
	MergedAt  *time.Time              `json:"merged_at"`
	User      struct{ Login string }  `json:"user"`
	Labels    []struct{ Name string } `json:"labels"`
}

// releaseVersion is one release added by a PR.
type releaseVersion struct {
	provider            string // release directory, e.g. "capa"
	major, minor, patch int
}

func (v releaseVersion) String() string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}

func (v releaseVersion) less(o releaseVersion) bool {
	if v.major != o.major {
		return v.major < o.major
	}
	if v.minor != o.minor {
		return v.minor < o.minor
	}
	return v.patch < o.patch
}

// commit is one commit on the default branch that touched release files.
type commit struct {
	sha     string
	subject string
	added   []releaseVersion // new releases
	deleted []releaseVersion // deleted archived releases, i.e. reactivations
}

func main() {
	var (
		repo     string
		repoDir  string
		ref      string
		minMajor int
		output   string
		token    string
		verbose  bool
	)

	flag.StringVar(&repo, "repo", "giantswarm/releases", "GitHub repository in owner/name form")
	flag.StringVar(&repoDir, "repo-dir", ".", "Local clone of the repository to read the git history from")
	flag.StringVar(&ref, "ref", "HEAD", "Git ref of the default branch to read the history from")
	flag.IntVar(&minMajor, "min-major", defaultMinMajor, "Ignore releases below this major version")
	flag.StringVar(&output, "output", "release-lead-time.json", "Path of the JSON file to write")
	flag.StringVar(&token, "token", os.Getenv("GITHUB_TOKEN"), "GitHub token (defaults to GITHUB_TOKEN)")
	flag.BoolVar(&verbose, "verbose", false, "Print the collected releases")
	flag.Parse()

	commits, err := releaseCommits(repoDir, ref)
	if err != nil {
		log.Fatalf("Error reading git history: %v", err)
	}

	client := &githubClient{repo: repo, token: token, http: &http.Client{Timeout: 30 * time.Second}}

	// Group the added releases by pull request.
	added := map[int][]releaseVersion{}
	for _, c := range commits {
		versions := newReleases(c, minMajor)
		if len(versions) == 0 {
			if verbose && len(c.added) > 0 {
				log.Printf("Skipping %s %q: only reactivates archived releases", c.sha[:8], c.subject)
			}
			continue
		}

		number, ok := prNumber(c.subject)
		if !ok {
			number, ok, err = client.pullRequestForCommit(c.sha)
			if err != nil {
				log.Fatalf("Error looking up pull request for commit %s: %v", c.sha, err)
			}
			if !ok {
				log.Printf("Warning: no pull request found for commit %s %q, skipping", c.sha[:8], c.subject)
				continue
			}
		}
		added[number] = append(added[number], versions...)
	}

	var releases []Release
	for number, versions := range added {
		pr, err := client.pullRequest(number)
		if err != nil {
			log.Fatalf("Error fetching pull request #%d: %v", number, err)
		}
		r, ok := toRelease(pr, versions)
		if !ok {
			log.Printf("Warning: pull request #%d is not merged, skipping", number)
			continue
		}
		releases = append(releases, r)
	}

	sort.Slice(releases, func(i, j int) bool {
		return releases[i].MergedAt.Before(releases[j].MergedAt)
	})

	if verbose {
		for _, r := range releases {
			planned := "-"
			if r.PlannedMergeDate != nil {
				planned = fmt.Sprintf("%s (%+.1f days)", *r.PlannedMergeDate, *r.DelayDays)
			}
			log.Printf("#%d %-6s %-8s %-8s lead time %6.1f days, planned %s", r.Number, r.Provider, r.Version, r.ReleaseType, r.LeadTimeDays, planned)
		}
	}

	out := Output{GeneratedAt: time.Now().UTC().Truncate(time.Second), Repository: repo, Releases: releases}
	data, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		log.Fatalf("Error encoding output: %v", err)
	}
	if err := os.WriteFile(output, append(data, '\n'), 0o644); err != nil {
		log.Fatalf("Error writing %s: %v", output, err)
	}
	log.Printf("Wrote %d releases to %s", len(releases), output)
}

// releaseCommits lists the commits on ref that added or deleted release files,
// oldest first.
func releaseCommits(repoDir, ref string) ([]commit, error) {
	args := []string{"log", ref, "--reverse", "--no-renames", "--diff-filter=AD", "--name-status",
		"--format=commit%x09%H%x09%s", "--"}
	for dir := range providers {
		args = append(args, dir+"/v*/release.yaml", dir+"/archived/v*/release.yaml")
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = repoDir
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("git %s: %w: %s", strings.Join(args[:2], " "), err, strings.TrimSpace(stderr.String()))
	}
	return parseGitLog(bytes.NewReader(out))
}

// parseGitLog parses the output of releaseCommits' git log command.
func parseGitLog(r io.Reader) ([]commit, error) {
	var commits []commit
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		fields := strings.SplitN(line, "\t", 3)
		switch {
		case fields[0] == "commit" && len(fields) == 3:
			commits = append(commits, commit{sha: fields[1], subject: fields[2]})
		case (fields[0] == "A" || fields[0] == "D") && len(fields) == 2 && len(commits) > 0:
			m := releaseFileRe.FindStringSubmatch(fields[1])
			if m == nil {
				continue
			}
			v := releaseVersion{provider: m[1]}
			v.major, _ = strconv.Atoi(m[3])
			v.minor, _ = strconv.Atoi(m[4])
			v.patch, _ = strconv.Atoi(m[5])
			archived := m[2] != ""
			c := &commits[len(commits)-1]
			if fields[0] == "A" && !archived {
				c.added = append(c.added, v)
			} else if fields[0] == "D" && archived {
				c.deleted = append(c.deleted, v)
			}
		}
	}
	return commits, scanner.Err()
}

// newReleases returns the releases a commit added, minus releases it moved
// back from the archive and releases below minMajor.
func newReleases(c commit, minMajor int) []releaseVersion {
	var result []releaseVersion
	for _, a := range c.added {
		if a.major < minMajor {
			continue
		}
		reactivated := false
		for _, d := range c.deleted {
			if a == d {
				reactivated = true
				break
			}
		}
		if !reactivated {
			result = append(result, a)
		}
	}
	return result
}

// prNumber extracts the pull request number from a merge commit subject.
func prNumber(subject string) (int, bool) {
	m := prNumberRe.FindStringSubmatch(subject)
	if m == nil {
		return 0, false
	}
	n, err := strconv.Atoi(m[1])
	return n, err == nil
}

// toRelease converts a merged release PR into a KPI row. ok is false when the
// PR is not merged.
func toRelease(pr pullRequest, versions []releaseVersion) (Release, bool) {
	if pr.MergedAt == nil || len(versions) == 0 {
		return Release{}, false
	}

	created := pr.CreatedAt.UTC()
	merged := pr.MergedAt.UTC()
	version := mainVersion(versions)

	r := Release{
		Number:       pr.Number,
		URL:          pr.HTMLURL,
		Title:        pr.Title,
		Provider:     provider(versions),
		Version:      version.String(),
		ReleaseType:  releaseType(pr, version),
		Author:       pr.User.Login,
		CreatedAt:    created,
		MergedAt:     merged,
		LeadTimeDays: days(merged.Sub(created)),
	}

	if planned, ok := plannedMergeDate(pr); ok {
		r.Automated = true
		p := planned.Format("2006-01-02")
		r.PlannedMergeDate = &p
		d := days(merged.Sub(planned))
		r.DelayDays = &d
	}

	return r, true
}

// provider returns the provider acronym for the released versions, or the
// consolidated name when several providers were released at once.
func provider(versions []releaseVersion) string {
	dirs := map[string]bool{}
	for _, v := range versions {
		dirs[v.provider] = true
	}
	if len(dirs) != 1 {
		return consolidatedProvider
	}
	if p, ok := providers[versions[0].provider]; ok {
		return p
	}
	return strings.ToUpper(versions[0].provider)
}

// mainVersion returns the version released by most providers in a PR, the
// lowest one on a tie. Consolidated PRs occasionally release a different patch
// version for one provider.
func mainVersion(versions []releaseVersion) releaseVersion {
	count := map[releaseVersion]int{}
	for _, v := range versions {
		v.provider = ""
		count[v]++
	}
	var best releaseVersion
	bestCount := 0
	for v, n := range count {
		if n > bestCount || (n == bestCount && v.less(best)) {
			best, bestCount = v, n
		}
	}
	return best
}

// releaseType returns major, minor or patch from the release/* label of the
// PR, or from the version when the PR has no such label.
func releaseType(pr pullRequest, v releaseVersion) string {
	for _, l := range pr.Labels {
		if t, ok := releaseTypeLabels[l.Name]; ok {
			return t
		}
	}
	switch {
	case v.patch != 0:
		return "patch"
	case v.minor != 0:
		return "minor"
	default:
		return "major"
	}
}

// plannedMergeDate returns the planned merge date of a scheduled release PR.
//
// The date is taken from the PLANNED_MERGE_DATE marker that create-release.yaml
// writes into the body of scheduled release PRs. PRs created before the marker
// existed are recognised by their creation time instead: the automation runs on
// the 1st of the month at 06:00 UTC as the automation user, and the planned
// merge date is the next 1st of the month.
func plannedMergeDate(pr pullRequest) (time.Time, bool) {
	if m := plannedMergeDateRe.FindStringSubmatch(pr.Body); m != nil {
		t, err := time.Parse("2006-01-02", m[1])
		if err == nil {
			return t, true
		}
	}

	created := pr.CreatedAt.UTC()
	if pr.User.Login != automationUser || created.Day() != scheduledRunDay ||
		created.Hour() < scheduledRunStartHour || created.Hour() >= scheduledRunEndHour {
		return time.Time{}, false
	}
	return nextFirstOfMonth(created), true
}

// nextFirstOfMonth returns midnight UTC of the 1st of the month after t.
func nextFirstOfMonth(t time.Time) time.Time {
	t = t.UTC()
	return time.Date(t.Year(), t.Month()+1, 1, 0, 0, 0, 0, time.UTC)
}

// days converts a duration to days with one decimal.
func days(d time.Duration) float64 {
	return math.Round(d.Hours()/24*10) / 10
}

type githubClient struct {
	repo  string
	token string
	http  *http.Client
}

// pullRequest fetches one pull request.
func (c *githubClient) pullRequest(number int) (pullRequest, error) {
	var pr pullRequest
	err := c.get(fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d", c.repo, number), &pr)
	return pr, err
}

// pullRequestForCommit returns the merged pull request that contains a commit.
// This is only needed for merge commits whose subject has no PR number.
func (c *githubClient) pullRequestForCommit(sha string) (int, bool, error) {
	var prs []pullRequest
	if err := c.get(fmt.Sprintf("https://api.github.com/repos/%s/commits/%s/pulls", c.repo, sha), &prs); err != nil {
		return 0, false, err
	}
	for _, pr := range prs {
		if pr.MergedAt != nil {
			return pr.Number, true, nil
		}
	}
	return 0, false, nil
}

func (c *githubClient) get(u string, v interface{}) error {
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	if c.token != "" {
		req.Header.Set("Authorization", "Bearer "+c.token)
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %s: %s: %s", u, resp.Status, strings.TrimSpace(string(body)))
	}
	return json.Unmarshal(body, v)
}
