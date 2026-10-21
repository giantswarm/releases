package main

import (
	"testing"
	"time"
)

func run(name, conclusion, at string) checkRun {
	t := ts(at)
	return checkRun{Name: name, Conclusion: conclusion, CompletedAt: &t}
}

func TestOutcome(t *testing.T) {
	cases := map[string]string{
		"success": outcomeSuccess, "neutral": outcomeSkipped, "skipped": outcomeSkipped,
		"cancelled": outcomeCancelled, "failure": outcomeFailure, "timed_out": outcomeFailure,
		"action_required": outcomeFailure, "stale": outcomeFailure, "": "", "startup_failure": "",
	}
	for in, want := range cases {
		if got := outcome(in); got != want {
			t.Errorf("outcome(%q) = %q, want %q", in, got, want)
		}
	}
}

func TestSuiteResults(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"capa"})
	standard, upgrade := "Release Tests / CAPA Standard Suite", "Release Tests / CAPA Upgrade Suite"
	res := testResults{
		runs:    []issueComment{comment("Gacko", "/run releases-test-suites", "2026-09-09T08:00:00Z")},
		waivers: map[string]time.Time{"capa/private": ts("2026-09-12T12:00:00Z")},
		ran:     map[string]bool{standard: true, upgrade: true},
		results: map[string][]checkRun{
			// Out of order on purpose: a rerun on the same commit, then the
			// first failure.
			standard: {run(standard, "success", "2026-09-09T12:00:00Z"), run(standard, "failure", "2026-09-09T10:00:00Z"), run(standard, "cancelled", "2026-09-09T09:00:00Z")},
			upgrade:  {run(upgrade, "timed_out", "2026-09-10T18:00:00Z"), run(upgrade, "success", "2026-09-11T18:00:00Z"), run(upgrade, "failure", "2026-09-12T18:00:00Z")},
		},
		firstGreen: map[string]time.Time{standard: ts("2026-09-09T12:00:00Z"), upgrade: ts("2026-09-11T18:00:00Z")},
	}
	r := Release{Number: 1, Name: "CAPA v35.1.0", Provider: "CAPA", ReleaseType: "minor", CreatedAt: ts("2026-09-01T06:00:00Z"), MergedAt: ts("2026-09-15T00:00:00Z"), MergedQuarter: "2026-Q3"}
	applyTestResults(&r, wanted, allRequired(wanted), res)
	rows := suiteResults(r, wanted, allRequired(wanted), res)

	if len(rows) != 3 {
		t.Fatalf("got %d rows, want 3: %+v", len(rows), rows)
	}
	byKey := map[string]SuiteResult{}
	for _, row := range rows {
		byKey[row.Suite] = row
	}

	s := byKey["capa/standard"]
	if s.Runs != 3 || s.Successes != 1 || s.Failures != 1 || s.Cancelled != 1 || s.Failed != 1 {
		t.Errorf("standard counts: %+v", s)
	}
	if s.RunsToGreen == nil || *s.RunsToGreen != 3 || s.FirstGreenAt == nil || !s.FirstGreenAt.Equal(ts("2026-09-09T12:00:00Z")) {
		t.Errorf("standard first green: %+v", s)
	}
	if s.LastToGreen != 0 || s.Name != "CAPA v35.1.0" || s.Check != standard || s.MergedQuarter != "2026-Q3" {
		t.Errorf("standard row: %+v", s)
	}

	u := byKey["capa/upgrade"]
	// A failure after the first pass still counts as a failure, but does not
	// move the first green.
	if u.Runs != 3 || u.Failures != 2 || u.Successes != 1 || *u.RunsToGreen != 2 {
		t.Errorf("upgrade counts: %+v", u)
	}

	p := byKey["capa/private"]
	if p.Runs != 0 || p.Waived != 1 || p.Failed != 0 || p.RunsToGreen != nil || p.FirstGreenAt == nil || !p.FirstGreenAt.Equal(ts("2026-09-12T12:00:00Z")) {
		t.Errorf("private (waived) row: %+v", p)
	}
	// The waiver was the last thing the release waited for.
	if p.LastToGreen != 1 || u.LastToGreen != 0 {
		t.Errorf("last to green: private %d, upgrade %d, want 1, 0", p.LastToGreen, u.LastToGreen)
	}
}

func TestSuiteResultsSkipsSuitesThatDoNotCount(t *testing.T) {
	defs := testDefs()
	wanted := defs.expectedChecks([]string{"capa"})
	required := defs.requiredSuites(wanted, ts("2026-07-01T06:00:00Z")) // before the merge gate
	standard := "Release Tests / CAPA Standard Suite"
	res := testResults{
		waivers: map[string]time.Time{},
		ran:     map[string]bool{standard: true},
		results: map[string][]checkRun{standard: {run(standard, "neutral", "2026-07-03T12:00:00Z")}},
	}
	rows := suiteResults(Release{Number: 2}, wanted, required, res)
	if len(rows) != 1 || rows[0].Suite != "capa/standard" || rows[0].Skipped != 1 || rows[0].Runs != 1 {
		t.Errorf("rows = %+v, want only the standard suite with one skipped run", rows)
	}
}
