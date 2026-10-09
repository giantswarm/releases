package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"
)

// Time to all green tests.
//
// For every release PR, the moment is determined at which every expected E2E
// suite of the PR's new releases had passed for the first time, in any run on
// any commit of the PR. Suite results are the "Release Tests / <suite>" check
// runs that the check-run-results-to-pr Tekton task attaches to the tested
// commit. Test runs are triggered with "/run releases-test-suites" comments.
//
// A "neutral" conclusion means the suite was skipped entirely, e.g. the plain
// upgrade suite of a major release, which has no previous minor to upgrade
// from. That is not test friction, so it counts as green here. The merge gate
// in .github/scripts/e2e-coverage.js is stricter and wants a /waive-suite.
//
// See https://github.com/giantswarm/roadmap/issues/4385

const (
	suiteCheckPrefix = "Release Tests / "
	runCommand       = "/run releases-test-suites"

	// maxCommitsScanned caps the check run lookups per PR.
	maxCommitsScanned = 200
)

var (
	waiverRe = regexp.MustCompile(`(?im)^\s*/waive-suite\s+(?:\./providers/)?([a-z-]+)/([a-z0-9-]+)\s+\S`)
)

// suiteDefinitions is the content of .github/scripts/e2e-suites.json, shared
// with the E2E coverage merge gate.
type suiteDefinitions struct {
	ExpectedSuites     map[string][]string `json:"expectedSuites"`
	CheckNameFormats   map[string]string   `json:"checkNameFormats"`
	SuitePathProviders map[string]string   `json:"suitePathProviders"`
}

func loadSuiteDefinitions(path string) (suiteDefinitions, error) {
	var defs suiteDefinitions
	data, err := os.ReadFile(path)
	if err != nil {
		return defs, err
	}
	if err := json.Unmarshal(data, &defs); err != nil {
		return defs, fmt.Errorf("parsing %s: %w", path, err)
	}
	return defs, nil
}

// expectedChecks returns the check run names of every suite expected for the
// given release directories, keyed by "<dir>/<suite>".
func (d suiteDefinitions) expectedChecks(dirs []string) map[string]string {
	checks := map[string]string{}
	for _, dir := range dirs {
		capi, ok := providers[dir]
		if !ok {
			capi = strings.ToUpper(dir)
		}
		for _, suite := range d.ExpectedSuites[dir] {
			format, ok := d.CheckNameFormats[suite]
			if !ok {
				continue
			}
			checks[dir+"/"+suite] = suiteCheckPrefix + strings.ReplaceAll(format, "{capi}", capi)
		}
	}
	return checks
}

// providerDir resolves a provider as written in a /waive-suite command
// ("azure", "capz", "aws", ...) to the release directory.
func (d suiteDefinitions) providerDir(name string) string {
	name = strings.ToLower(name)
	if name == "aws" {
		return "capa"
	}
	if _, ok := d.ExpectedSuites[name]; ok {
		return name
	}
	for dir, p := range d.SuitePathProviders {
		if p == name {
			return dir
		}
	}
	return ""
}

// issueComment is the subset of the GitHub issue comments API response we need.
type issueComment struct {
	Body              string    `json:"body"`
	CreatedAt         time.Time `json:"created_at"`
	AuthorAssociation string    `json:"author_association"`
	User              struct {
		Login string `json:"login"`
	} `json:"user"`
}

// checkRun is the subset of the GitHub check runs API response we need.
type checkRun struct {
	Name        string     `json:"name"`
	Conclusion  string     `json:"conclusion"`
	CompletedAt *time.Time `json:"completed_at"`
}

// testResults is what the test-related fields of a Release are computed from.
type testResults struct {
	runs        []issueComment       // /run releases-test-suites comments, oldest first
	waivers     map[string]time.Time // "<dir>/<suite>" -> time of the waiver
	firstGreen  map[string]time.Time // check run name -> first successful completion
	checkedShas int
}

// collectTestResults gathers the test runs, waivers and first successful suite
// runs of a PR.
func (c *githubClient) collectTestResults(number int, wanted map[string]string, defs suiteDefinitions) (testResults, error) {
	res := testResults{waivers: map[string]time.Time{}, firstGreen: map[string]time.Time{}}

	comments, err := c.issueComments(number)
	if err != nil {
		return res, err
	}
	for _, cm := range comments {
		body := strings.TrimSpace(cm.Body)
		if strings.HasPrefix(body, runCommand) {
			res.runs = append(res.runs, cm)
		}
		if cm.AuthorAssociation != "OWNER" && cm.AuthorAssociation != "MEMBER" && cm.AuthorAssociation != "COLLABORATOR" {
			continue
		}
		for _, m := range waiverRe.FindAllStringSubmatch(cm.Body, -1) {
			dir := defs.providerDir(m[1])
			if dir == "" {
				continue
			}
			key := dir + "/" + strings.ToLower(m[2])
			if t, ok := res.waivers[key]; !ok || cm.CreatedAt.Before(t) {
				res.waivers[key] = cm.CreatedAt
			}
		}
	}
	sort.Slice(res.runs, func(i, j int) bool { return res.runs[i].CreatedAt.Before(res.runs[j].CreatedAt) })

	wantedNames := map[string]bool{}
	for _, name := range wanted {
		wantedNames[name] = true
	}

	shas, err := c.pullRequestCommits(number)
	if err != nil {
		return res, err
	}
	for i, sha := range shas {
		if i >= maxCommitsScanned {
			log.Printf("Warning: #%d has more than %d commits, older results are not counted", number, maxCommitsScanned)
			break
		}
		runs, err := c.checkRuns(sha)
		if err != nil {
			return res, err
		}
		res.checkedShas++
		for _, run := range runs {
			if !isGreen(run.Conclusion) || run.CompletedAt == nil || !wantedNames[run.Name] {
				continue
			}
			if t, ok := res.firstGreen[run.Name]; !ok || run.CompletedAt.Before(t) {
				res.firstGreen[run.Name] = *run.CompletedAt
			}
		}
	}
	return res, nil
}

// isGreen reports whether a suite check run conclusion counts as passed. See
// the package comment for why "neutral" does.
func isGreen(conclusion string) bool {
	return conclusion == "success" || conclusion == "neutral"
}

// applyTestResults fills the test-related fields of a release.
func applyTestResults(r *Release, wanted map[string]string, res testResults) {
	r.SuitesExpected = len(wanted)
	r.TestRuns = len(res.runs)
	for _, run := range res.runs {
		if run.User.Login == automationUser {
			r.TestRunsAutomated++
		}
	}
	if len(res.runs) > 0 {
		t := res.runs[0].CreatedAt.UTC()
		r.TestsFirstRunAt = &t
	}

	var allGreen time.Time
	complete := true
	for key, name := range wanted {
		t, ok := res.firstGreen[name]
		if ok {
			r.SuitesGreen++
		} else if w, waived := res.waivers[key]; waived {
			t, ok = w, true
			r.SuitesWaived++
		}
		if !ok {
			complete = false
			continue
		}
		if t.After(allGreen) {
			allGreen = t
		}
	}
	if !complete || len(wanted) == 0 {
		return
	}

	allGreen = allGreen.UTC()
	r.AllGreenAt = &allGreen
	d := days(allGreen.Sub(r.CreatedAt))
	r.TimeToGreenDays = &d
	if r.TestsFirstRunAt != nil {
		f := days(allGreen.Sub(*r.TestsFirstRunAt))
		r.TestFrictionDays = &f
	}
}

// issueComments lists the comments of a pull request, oldest first.
func (c *githubClient) issueComments(number int) ([]issueComment, error) {
	var result []issueComment
	for page := 1; ; page++ {
		var comments []issueComment
		u := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/comments?per_page=100&page=%d", c.repo, number, page)
		if err := c.get(u, &comments); err != nil {
			return nil, err
		}
		result = append(result, comments...)
		if len(comments) < 100 {
			return result, nil
		}
	}
}

// pullRequestCommits returns the commits that were ever the head of a pull
// request: its current commits plus the commits its branch was force-pushed away
// from, which the timeline records. Check runs stay attached to those commits.
func (c *githubClient) pullRequestCommits(number int) ([]string, error) {
	seen := map[string]bool{}
	var shas []string
	add := func(sha string) {
		if sha != "" && !seen[sha] {
			seen[sha] = true
			shas = append(shas, sha)
		}
	}

	for page := 1; ; page++ {
		var commits []struct {
			SHA string `json:"sha"`
		}
		u := fmt.Sprintf("https://api.github.com/repos/%s/pulls/%d/commits?per_page=100&page=%d", c.repo, number, page)
		if err := c.get(u, &commits); err != nil {
			return nil, err
		}
		for _, cm := range commits {
			add(cm.SHA)
		}
		if len(commits) < 100 {
			break
		}
	}

	for page := 1; ; page++ {
		var events []struct {
			Event    string `json:"event"`
			CommitID string `json:"commit_id"`
		}
		u := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d/timeline?per_page=100&page=%d", c.repo, number, page)
		if err := c.get(u, &events); err != nil {
			return nil, err
		}
		for _, e := range events {
			if e.Event == "head_ref_force_pushed" {
				add(e.CommitID)
			}
		}
		if len(events) < 100 {
			break
		}
	}
	return shas, nil
}

// checkRuns lists the check runs of a commit.
func (c *githubClient) checkRuns(sha string) ([]checkRun, error) {
	var result []checkRun
	for page := 1; ; page++ {
		var res struct {
			CheckRuns []checkRun `json:"check_runs"`
		}
		u := fmt.Sprintf("https://api.github.com/repos/%s/commits/%s/check-runs?per_page=100&page=%d", c.repo, sha, page)
		if err := c.get(u, &res); err != nil {
			return nil, err
		}
		result = append(result, res.CheckRuns...)
		if len(res.CheckRuns) < 100 {
			return result, nil
		}
	}
}
