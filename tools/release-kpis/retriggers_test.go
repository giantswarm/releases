package main

import (
	"testing"
	"time"
)

func at(s string) time.Time {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		panic(err)
	}
	return t
}

func runOn(sha, conclusion, completed string) checkRun {
	r := run("Release Tests / CAPA Standard Suite", conclusion, completed)
	r.HeadSHA = sha
	return r
}

func TestClassifyRetriggers(t *testing.T) {
	runs := []issueComment{
		comment("taylorbot", runCommand, "2026-09-01T10:00:00Z"),               // first run
		comment("alice", runCommand, "2026-09-01T12:00:00Z"),                   // after a failure
		comment("giantswarm-release[bot]", runCommand, "2026-09-02T09:00:00Z"), // automation, active stage
		comment("alice", runCommand, "2026-09-03T09:00:00Z"),                   // after a commit
		comment("alice", runCommand, "2026-09-03T09:05:00Z"),                   // nothing happened: other
		comment("alice", runCommand, "2026-09-04T09:00:00Z"),                   // only a cancellation since: other
	}
	results := map[string][]checkRun{
		"Release Tests / CAPA Standard Suite": {
			run("Release Tests / CAPA Standard Suite", "failure", "2026-09-01T11:00:00Z"),
			run("Release Tests / CAPA Standard Suite", "success", "2026-09-01T13:00:00Z"),
			run("Release Tests / CAPA Standard Suite", "cancelled", "2026-09-03T10:00:00Z"),
		},
	}
	commits := []time.Time{at("2026-08-31T08:00:00Z"), at("2026-09-02T15:00:00Z")}

	c := classifyRetriggers(runs, results, commits)
	want := retriggerCounts{total: 5, automated: 1, afterFailure: 1, afterCommit: 1, other: 2}
	if c != want {
		t.Errorf("got %+v, want %+v", c, want)
	}

	if c := classifyRetriggers(runs[:1], results, commits); c != (retriggerCounts{}) {
		t.Errorf("a single run has no retriggers, got %+v", c)
	}
}

func TestSuiteReruns(t *testing.T) {
	runs := []checkRun{
		runOn("aaa", "failure", "2026-09-01T11:00:00Z"),
		runOn("aaa", "success", "2026-09-01T13:00:00Z"),   // after a failure
		runOn("aaa", "success", "2026-09-01T15:00:00Z"),   // redundant: passed on the same commit
		runOn("bbb", "cancelled", "2026-09-02T10:00:00Z"), // after a commit
		runOn("bbb", "neutral", "2026-09-02T12:00:00Z"),   // after a cancellation counts as after a failure
		runOn("bbb", "success", "2026-09-02T14:00:00Z"),   // redundant: skipped counts as passed
	}
	reruns, afterFailure, afterCommit, redundant := suiteReruns(runs)
	if reruns != 5 || afterFailure != 2 || afterCommit != 1 || redundant != 2 {
		t.Errorf("got reruns %d, after failure %d, after commit %d, redundant %d; want 5, 2, 1, 2", reruns, afterFailure, afterCommit, redundant)
	}
	if reruns, _, _, _ := suiteReruns(runs[:1]); reruns != 0 {
		t.Errorf("a single run is no rerun, got %d", reruns)
	}
}

func TestApplyTestResultsRetriggers(t *testing.T) {
	wanted := map[string]string{"capa/standard": "Release Tests / CAPA Standard Suite"}
	res := testResults{
		runs: []issueComment{
			comment("taylorbot", runCommand, "2026-09-01T10:00:00Z"),
			comment("alice", runCommand, "2026-09-01T12:00:00Z"),
		},
		ran: map[string]bool{"Release Tests / CAPA Standard Suite": true},
		results: map[string][]checkRun{
			"Release Tests / CAPA Standard Suite": {
				runOn("aaa", "failure", "2026-09-01T11:00:00Z"),
				runOn("aaa", "success", "2026-09-01T13:00:00Z"),
			},
		},
		firstGreen: map[string]time.Time{"Release Tests / CAPA Standard Suite": at("2026-09-01T13:00:00Z")},
		waivers:    map[string]time.Time{},
	}
	r := Release{CreatedAt: at("2026-09-01T09:00:00Z")}
	applyTestResults(&r, wanted, allRequired(wanted), res)
	if r.Retriggers != 1 || r.RetriggersAfterFailure != 1 || r.Retriggered == nil || *r.Retriggered != 1 {
		t.Errorf("got retriggers %d, after failure %d, retriggered %v", r.Retriggers, r.RetriggersAfterFailure, r.Retriggered)
	}
	rows := suiteResults(r, wanted, allRequired(wanted), res)
	applySuiteTotals(&r, rows)
	if r.SuiteRuns != 2 || r.SuiteReruns != 1 || r.SuiteRerunsRedundant != 0 {
		t.Errorf("got suite runs %d, reruns %d, redundant %d; want 2, 1, 0", r.SuiteRuns, r.SuiteReruns, r.SuiteRerunsRedundant)
	}
	if rows[0].Reruns != 1 || rows[0].RerunsAfterFailure != 1 {
		t.Errorf("suite row: got reruns %d, after failure %d", rows[0].Reruns, rows[0].RerunsAfterFailure)
	}

	// Without per-suite results the reasons cannot be told.
	old := Release{CreatedAt: at("2026-09-01T09:00:00Z")}
	applyTestResults(&old, map[string]string{}, map[string]bool{}, testResults{runs: res.runs, waivers: map[string]time.Time{}})
	if old.Retriggers != 1 || old.RetriggersAfterFailure != 0 || old.RetriggersOther != 0 || old.Retriggered != nil {
		t.Errorf("got retriggers %d, after failure %d, other %d, retriggered %v; want 1, 0, 0, nil", old.Retriggers, old.RetriggersAfterFailure, old.RetriggersOther, old.Retriggered)
	}
}
