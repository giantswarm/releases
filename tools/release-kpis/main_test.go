package main

import (
	"strings"
	"testing"
	"time"
)

func ts(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func mergedPR(number int, title, body, author, created, merged string) pullRequest {
	m := ts(merged)
	pr := pullRequest{Number: number, Title: title, Body: body, CreatedAt: ts(created), MergedAt: &m}
	pr.User.Login = author
	return pr
}

func v(provider string, major, minor, patch int) releaseVersion {
	return releaseVersion{provider: provider, major: major, minor: minor, patch: patch}
}

const gitLog = `commit	aaaa	CAPA: Release v29.0.0. (#1318)
A	capa/v29.0.0/release.yaml
A	capa/v29.0.0/README.md

commit	bbbb	CAPI: Release v35.0.1. (#2421)
D	azure/v35.0.0/release.yaml
A	azure/archived/v35.0.0/release.yaml
A	azure/v35.0.1/release.yaml
D	capa/v35.0.0/release.yaml
A	capa/archived/v35.0.0/release.yaml
A	capa/v35.0.1/release.yaml

commit	cccc	CAPA: Bring back v33.2.0 (#2371)
D	capa/archived/v33.2.0/release.yaml
A	capa/v33.2.0/release.yaml

commit	dddd	Merge pull request #1234 from giantswarm/release-v30.0.0
A	capa/v30.0.0/release.yaml

commit	eeee	Azure release 20.0.0 (#999)
A	azure/v20.0.0/release.yaml
`

func TestParseGitLog(t *testing.T) {
	commits, err := parseGitLog(strings.NewReader(gitLog))
	if err != nil {
		t.Fatal(err)
	}
	if len(commits) != 5 {
		t.Fatalf("got %d commits, want 5", len(commits))
	}

	if got := newReleases(commits[0], 25); len(got) != 1 || got[0] != v("capa", 29, 0, 0) {
		t.Errorf("commit aaaa: new releases = %v", got)
	}

	// Archiving the previous release must not count as a new or deleted release.
	if got := newReleases(commits[1], 25); len(got) != 2 || got[0] != v("azure", 35, 0, 1) || got[1] != v("capa", 35, 0, 1) {
		t.Errorf("commit bbbb: new releases = %v", got)
	}

	// Moving a release back from the archive is not a new release.
	if got := newReleases(commits[2], 25); len(got) != 0 {
		t.Errorf("commit cccc: new releases = %v, want none", got)
	}

	// Vintage releases below the minimum major version are ignored.
	if got := newReleases(commits[4], 25); len(got) != 0 {
		t.Errorf("commit eeee: new releases = %v, want none", got)
	}
	if got := newReleases(commits[4], 0); len(got) != 1 {
		t.Errorf("commit eeee with min-major 0: new releases = %v, want one", got)
	}

	if n, ok := prNumber(commits[0].subject); !ok || n != 1318 {
		t.Errorf("prNumber(aaaa) = %d, %v", n, ok)
	}
	if _, ok := prNumber(commits[3].subject); ok {
		t.Error("prNumber(dddd) should not match a merge commit subject")
	}
}

func TestProviderAndVersion(t *testing.T) {
	cases := []struct {
		versions []releaseVersion
		provider string
		version  string
	}{
		{[]releaseVersion{v("capa", 29, 0, 0)}, "CAPA", "v29.0.0"},
		{[]releaseVersion{v("azure", 29, 1, 0)}, "CAPZ", "v29.1.0"},
		{[]releaseVersion{v("cloud-director", 31, 0, 0)}, "CAPVCD", "v31.0.0"},
		{[]releaseVersion{v("azure", 34, 1, 1), v("capa", 34, 1, 1), v("cloud-director", 34, 1, 2), v("vsphere", 34, 1, 1)}, "CAPI", "v34.1.1"},
		{[]releaseVersion{v("capa", 26, 4, 3), v("capa", 26, 4, 4)}, "CAPA", "v26.4.3"},
	}
	for _, c := range cases {
		if got := provider(c.versions); got != c.provider {
			t.Errorf("provider(%v) = %s, want %s", c.versions, got, c.provider)
		}
		if got := mainVersion(c.versions).String(); got != c.version {
			t.Errorf("mainVersion(%v) = %s, want %s", c.versions, got, c.version)
		}
	}
}

func TestToReleaseScheduledWithMarker(t *testing.T) {
	pr := mergedPR(2426, "CAPI: Release v36.0.0.",
		"Some body\n<!-- PLANNED_MERGE_DATE: 2026-10-01 -->\n", "taylorbot",
		"2026-09-01T06:35:58Z", "2026-09-25T04:31:25Z")
	pr.Labels = []struct{ Name string }{{"aws"}, {"release/major"}}

	r, ok := toRelease(pr, []releaseVersion{v("azure", 36, 0, 0), v("capa", 36, 0, 0)})
	if !ok {
		t.Fatal("expected a release")
	}
	if r.Provider != "CAPI" || r.Version != "v36.0.0" || r.ReleaseType != "major" || r.Author != "taylorbot" {
		t.Errorf("unexpected row: %+v", r)
	}
	if r.LeadTimeDays != 23.9 {
		t.Errorf("lead time = %v, want 23.9", r.LeadTimeDays)
	}
	if !r.Automated || r.PlannedMergeDate == nil || *r.PlannedMergeDate != "2026-10-01" {
		t.Fatalf("planned merge date = %v, want 2026-10-01", r.PlannedMergeDate)
	}
	if r.DelayDays == nil || *r.DelayDays != -5.8 {
		t.Errorf("delay = %v, want -5.8", r.DelayDays)
	}
	if r.OnTime == nil || *r.OnTime != 1 {
		t.Errorf("on time = %v, want 1", r.OnTime)
	}
	if r.MergedMonth != "2026-09" || r.MergedQuarter != "2026-Q3" {
		t.Errorf("merged month/quarter = %s/%s, want 2026-09/2026-Q3", r.MergedMonth, r.MergedQuarter)
	}
}

func TestToReleaseScheduledWithoutMarker(t *testing.T) {
	// PR created by the automation on the 1st, before the marker existed.
	pr := mergedPR(2215, "CAPI: Release v35.0.0.", "", "taylorbot",
		"2026-03-01T06:24:59Z", "2026-08-22T21:55:48Z")

	r, ok := toRelease(pr, []releaseVersion{v("capa", 35, 0, 0)})
	if !ok {
		t.Fatal("expected a release")
	}
	if !r.Automated || r.PlannedMergeDate == nil || *r.PlannedMergeDate != "2026-04-01" {
		t.Fatalf("planned merge date = %v, want 2026-04-01", r.PlannedMergeDate)
	}
	if r.DelayDays == nil || *r.DelayDays != 143.9 {
		t.Errorf("delay = %v, want 143.9", r.DelayDays)
	}
	if r.OnTime == nil || *r.OnTime != 0 {
		t.Errorf("on time = %v, want 0", r.OnTime)
	}
}

func TestToReleaseManual(t *testing.T) {
	cases := []struct {
		name    string
		author  string
		created string
	}{
		{"human author", "njuettner", "2026-08-01T06:30:00Z"},
		{"not the first of the month", "taylorbot", "2026-09-23T06:30:00Z"},
		{"outside the scheduled window", "taylorbot", "2026-09-01T13:42:46Z"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			pr := mergedPR(1, "CAPA: Release v35.1.0.", "", c.author, c.created, "2026-09-25T00:00:00Z")
			r, ok := toRelease(pr, []releaseVersion{v("capa", 35, 1, 0)})
			if !ok {
				t.Fatal("expected a release")
			}
			if r.Automated || r.PlannedMergeDate != nil || r.DelayDays != nil || r.OnTime != nil {
				t.Errorf("manual release must not have a planned merge date, got %+v", r)
			}
		})
	}
}

func TestToReleaseMarkerWinsOverHeuristic(t *testing.T) {
	pr := mergedPR(1, "CAPI: Release v37.0.0.", "<!-- PLANNED_MERGE_DATE: 2026-12-01 -->", "someone",
		"2026-10-15T10:00:00Z", "2026-11-20T00:00:00Z")
	r, _ := toRelease(pr, []releaseVersion{v("capa", 37, 0, 0)})
	if r.PlannedMergeDate == nil || *r.PlannedMergeDate != "2026-12-01" {
		t.Fatalf("planned merge date = %v, want 2026-12-01", r.PlannedMergeDate)
	}
}

func TestToReleaseSkipsUnmerged(t *testing.T) {
	pr := mergedPR(2, "CAPI: Release v36.0.0.", "", "x", "2026-09-01T06:30:00Z", "2026-09-02T00:00:00Z")
	pr.MergedAt = nil
	if _, ok := toRelease(pr, []releaseVersion{v("capa", 36, 0, 0)}); ok {
		t.Error("expected unmerged pull request to be skipped")
	}
}

func TestReleaseType(t *testing.T) {
	cases := map[releaseVersion]string{
		v("capa", 29, 0, 0): "major",
		v("capa", 29, 1, 0): "minor",
		v("capa", 28, 0, 1): "patch",
	}
	for ver, want := range cases {
		if got := releaseType(pullRequest{}, ver); got != want {
			t.Errorf("releaseType(%s) = %s, want %s", ver, got, want)
		}
	}

	pr := pullRequest{Labels: []struct{ Name string }{{"release/patch"}}}
	if got := releaseType(pr, v("capa", 29, 1, 0)); got != "patch" {
		t.Errorf("release type = %s, want patch (from label)", got)
	}
}

func TestNextFirstOfMonth(t *testing.T) {
	cases := map[string]string{
		"2026-09-01T06:35:58Z": "2026-10-01",
		"2026-12-01T06:35:58Z": "2027-01-01",
		"2026-01-31T23:59:59Z": "2026-02-01",
	}
	for in, want := range cases {
		if got := nextFirstOfMonth(ts(in)).Format("2006-01-02"); got != want {
			t.Errorf("nextFirstOfMonth(%s) = %s, want %s", in, got, want)
		}
	}
}

func TestDays(t *testing.T) {
	if got := days(36 * time.Hour); got != 1.5 {
		t.Errorf("days(36h) = %v, want 1.5", got)
	}
	if got := days(-36 * time.Hour); got != -1.5 {
		t.Errorf("days(-36h) = %v, want -1.5", got)
	}
}

func TestQuarter(t *testing.T) {
	cases := map[string]string{
		"2026-01-01T00:00:00Z": "2026-Q1",
		"2026-03-31T23:59:59Z": "2026-Q1",
		"2026-04-01T00:00:00Z": "2026-Q2",
		"2026-09-25T04:31:25Z": "2026-Q3",
		"2026-12-31T23:59:59Z": "2026-Q4",
	}
	for in, want := range cases {
		if got := quarter(ts(in)); got != want {
			t.Errorf("quarter(%s) = %s, want %s", in, got, want)
		}
	}
}

func labelEvent(event, label, at string) issueEvent {
	e := issueEvent{Event: event, CreatedAt: ts(at)}
	e.Label.Name = label
	return e
}

func TestStageDurations(t *testing.T) {
	// Taken from CAPI v36.0.0 (#2426).
	pr := mergedPR(2426, "CAPI: Release v36.0.0.", "", "taylorbot", "2026-09-01T06:35:58Z", "2026-09-25T04:31:25Z")
	events := []issueEvent{
		labelEvent("labeled", "release/major", "2026-09-01T06:36:00Z"),
		labelEvent("labeled", "stage/development", "2026-09-01T06:36:00Z"),
		labelEvent("unlabeled", "stage/development", "2026-09-10T10:08:10Z"),
		labelEvent("labeled", "stage/active", "2026-09-10T10:08:11Z"),
		labelEvent("unlabeled", "stage/active", "2026-09-24T15:36:08Z"),
		labelEvent("labeled", "stage/freeze", "2026-09-24T15:36:08Z"),
	}
	d, a, f := stageDurations(pr, events)
	if d == nil || *d != 9.1 {
		t.Errorf("development = %v, want 9.1", d)
	}
	if a == nil || *a != 14.2 {
		t.Errorf("active = %v, want 14.2", a)
	}
	if f == nil || *f != 0.5 {
		t.Errorf("freeze = %v, want 0.5", f)
	}
}

func TestStageDurationsRepeatedStage(t *testing.T) {
	pr := mergedPR(1, "CAPI: Release v37.0.0.", "", "x", "2026-10-01T00:00:00Z", "2026-10-11T00:00:00Z")
	events := []issueEvent{
		labelEvent("labeled", "stage/development", "2026-10-01T00:00:00Z"),
		labelEvent("labeled", "stage/active", "2026-10-03T00:00:00Z"),
		labelEvent("labeled", "stage/development", "2026-10-04T00:00:00Z"), // back to development
		labelEvent("labeled", "stage/freeze", "2026-10-09T00:00:00Z"),
	}
	d, a, f := stageDurations(pr, events)
	if *d != 7 || *a != 1 || *f != 2 {
		t.Errorf("development/active/freeze = %v/%v/%v, want 7/1/2", *d, *a, *f)
	}
}

func TestStageDurationsWithoutStages(t *testing.T) {
	pr := mergedPR(1, "CAPA: Release v29.0.0.", "", "x", "2024-07-01T00:00:00Z", "2024-07-02T00:00:00Z")
	if d, a, f := stageDurations(pr, []issueEvent{labelEvent("labeled", "aws", "2024-07-01T00:00:00Z")}); d != nil || a != nil || f != nil {
		t.Errorf("expected no stage durations, got %v/%v/%v", d, a, f)
	}
	if hasStageLabel(pr) {
		t.Error("PR without stage label reported as having one")
	}
	pr.Labels = []struct{ Name string }{{"stage/freeze"}}
	if !hasStageLabel(pr) {
		t.Error("PR with stage label not detected")
	}
}
