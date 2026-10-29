package main

import (
	"testing"
	"time"
)

func testDefs() suiteDefinitions {
	return suiteDefinitions{
		ExpectedSuites: map[string][]string{
			"capa":           {"standard", "upgrade", "private"},
			"azure":          {"standard"},
			"cloud-director": {"standard"},
			"aks":            {"standard", "byo-cni"},
		},
		CheckNameFormats: map[string]string{
			"standard": "{capi} Standard Suite",
			"upgrade":  "{capi} Upgrade Suite",
			"private":  "{capi} Private Suite",
			"byo-cni":  "{capi} BYO CNI Suite",
		},
		SuitePathProviders: map[string]string{"capa": "capa", "azure": "capz", "cloud-director": "capvcd", "aks": "aks"},
		RequiredSince:      "2026-08-04",
		SuiteRequiredSince: map[string]string{"capa/private": "2026-09-15"},
	}
}

// allRequired marks every wanted suite as required.
func allRequired(wanted map[string]string) map[string]bool {
	required := map[string]bool{}
	for key := range wanted {
		required[key] = true
	}
	return required
}

func TestExpectedChecks(t *testing.T) {
	checks := testDefs().expectedChecks([]string{"azure", "capa"})
	want := map[string]string{
		"capa/standard":  "Release Tests / CAPA Standard Suite",
		"capa/upgrade":   "Release Tests / CAPA Upgrade Suite",
		"capa/private":   "Release Tests / CAPA Private Suite",
		"azure/standard": "Release Tests / CAPZ Standard Suite",
	}
	if len(checks) != len(want) {
		t.Fatalf("got %d checks, want %d: %v", len(checks), len(want), checks)
	}
	for k, v := range want {
		if checks[k] != v {
			t.Errorf("%s = %q, want %q", k, checks[k], v)
		}
	}
	aks := testDefs().expectedChecks([]string{"aks"})
	if aks["aks/byo-cni"] != "Release Tests / AKS BYO CNI Suite" {
		t.Errorf("aks/byo-cni = %q, want %q", aks["aks/byo-cni"], "Release Tests / AKS BYO CNI Suite")
	}
}

func TestProviderDir(t *testing.T) {
	defs := testDefs()
	cases := map[string]string{"azure": "azure", "capz": "azure", "CAPVCD": "cloud-director", "aws": "capa", "capa": "capa", "nope": ""}
	for in, want := range cases {
		if got := defs.providerDir(in); got != want {
			t.Errorf("providerDir(%s) = %q, want %q", in, got, want)
		}
	}
}

func comment(user, body, at string) issueComment {
	c := issueComment{Body: body, CreatedAt: ts(at), AuthorAssociation: "MEMBER"}
	c.User.Login = user
	return c
}

func TestApplyTestResults(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"capa"})
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z")}
	res := testResults{
		runs: []issueComment{
			comment("Gacko", "/run releases-test-suites TARGET_SUITES=./providers/capa/standard", "2026-09-09T08:00:00Z"),
			comment("taylorbot", "/run releases-test-suites", "2026-09-10T10:00:00Z"),
			comment("giantswarm-release[bot]", "/run releases-test-suites", "2026-09-10T11:00:00Z"),
			comment("Gacko", "/run releases-test-suites", "2026-09-11T10:00:00Z"),
		},
		waivers: map[string]time.Time{"capa/private": ts("2026-09-12T12:00:00Z")},
		firstGreen: map[string]time.Time{
			"Release Tests / CAPA Standard Suite": ts("2026-09-09T12:00:00Z"),
			"Release Tests / CAPA Upgrade Suite":  ts("2026-09-11T18:00:00Z"),
		},
	}
	applyTestResults(&r, wanted, allRequired(wanted), res)

	if r.SuitesExpected != 3 || r.SuitesGreen != 2 || r.SuitesWaived != 1 {
		t.Errorf("suites expected/green/waived = %d/%d/%d, want 3/2/1", r.SuitesExpected, r.SuitesGreen, r.SuitesWaived)
	}
	if r.TestRuns != 4 || r.TestRunsAutomated != 2 {
		t.Errorf("test runs = %d (%d automated), want 4 (2)", r.TestRuns, r.TestRunsAutomated)
	}
	if r.TestsFirstRunAt == nil || !r.TestsFirstRunAt.Equal(ts("2026-09-09T08:00:00Z")) {
		t.Errorf("first run = %v", r.TestsFirstRunAt)
	}
	// All green once the waiver, the last of the three, is in place.
	if r.AllGreenAt == nil || !r.AllGreenAt.Equal(ts("2026-09-12T12:00:00Z")) {
		t.Fatalf("all green = %v, want 2026-09-12T12:00:00Z", r.AllGreenAt)
	}
	if *r.TimeToGreenDays != 11.3 || *r.TestFrictionDays != 3.2 {
		t.Errorf("time to green/friction = %v/%v, want 11.3/3.2", *r.TimeToGreenDays, *r.TestFrictionDays)
	}
	// No stage labels: friction starts at the first manual run, which here is
	// also the first run.
	if r.FrictionStartAt == nil || !r.FrictionStartAt.Equal(ts("2026-09-09T08:00:00Z")) {
		t.Errorf("friction start = %v, want 2026-09-09T08:00:00Z", r.FrictionStartAt)
	}
	// Four runs before all green: the two manual ones and both automated ones
	// (taylorbot and giantswarm-release[bot]).
	if r.MergedAllGreen == nil || *r.MergedAllGreen != 1 || r.RunsToGreen == nil || *r.RunsToGreen != 4 {
		t.Errorf("merged all green/runs to green = %v/%v, want 1/4", r.MergedAllGreen, r.RunsToGreen)
	}
}

func TestApplyTestResultsFrictionFromActiveStage(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"azure"})
	active := ts("2026-09-10T09:00:00Z")
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z"), ActiveAt: &active}
	res := testResults{
		runs: []issueComment{
			comment("taylorbot", "/run releases-test-suites", "2026-09-02T10:00:00Z"), // weekly bump during development
			comment("Gacko", "/run releases-test-suites", "2026-09-05T10:00:00Z"),     // manual, still in development
			comment("taylorbot", "/run releases-test-suites", "2026-09-10T09:05:00Z"), // stage change to active
			comment("Gacko", "/run releases-test-suites", "2026-09-14T10:00:00Z"),
		},
		waivers:    map[string]time.Time{},
		firstGreen: map[string]time.Time{"Release Tests / CAPZ Standard Suite": ts("2026-09-12T12:00:00Z")},
	}
	applyTestResults(&r, wanted, allRequired(wanted), res)
	if r.FrictionStartAt == nil || !r.FrictionStartAt.Equal(ts("2026-09-10T09:05:00Z")) {
		t.Fatalf("friction start = %v, want the first run after the active stage", r.FrictionStartAt)
	}
	if r.TestFrictionDays == nil || *r.TestFrictionDays != 2.1 {
		t.Errorf("friction = %v, want 2.1", r.TestFrictionDays)
	}
	if r.RunsToGreen == nil || *r.RunsToGreen != 3 {
		t.Errorf("runs to green = %v, want 3", r.RunsToGreen)
	}
}

func TestApplyTestResultsGreenBeforeActiveStage(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"azure"})
	active := ts("2026-09-20T09:00:00Z")
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z"), ActiveAt: &active}
	res := testResults{
		runs:       []issueComment{comment("taylorbot", "/run releases-test-suites", "2026-09-21T09:00:00Z")},
		waivers:    map[string]time.Time{},
		firstGreen: map[string]time.Time{"Release Tests / CAPZ Standard Suite": ts("2026-09-05T12:00:00Z")},
	}
	applyTestResults(&r, wanted, allRequired(wanted), res)
	if r.TestFrictionDays == nil || *r.TestFrictionDays != 0 {
		t.Errorf("friction = %v, want 0 when green before the friction start", r.TestFrictionDays)
	}
}

func TestApplyTestResultsNotRequiredSuites(t *testing.T) {
	defs := testDefs()
	wanted := defs.expectedChecks([]string{"capa"})
	// Created before the merge gate: only suites that ran count.
	required := defs.requiredSuites(wanted, ts("2026-07-01T06:00:00Z"))
	for key, v := range required {
		if v {
			t.Errorf("%s must not be required before the merge gate", key)
		}
	}
	r := Release{CreatedAt: ts("2026-07-01T06:00:00Z")}
	res := testResults{
		runs:       []issueComment{comment("Gacko", "/run releases-test-suites", "2026-07-02T08:00:00Z")},
		waivers:    map[string]time.Time{},
		ran:        map[string]bool{"Release Tests / CAPA Standard Suite": true, "Release Tests / CAPA Upgrade Suite": true},
		firstGreen: map[string]time.Time{"Release Tests / CAPA Standard Suite": ts("2026-07-03T12:00:00Z")},
	}
	applyTestResults(&r, wanted, required, res)
	if r.SuitesExpected != 2 || r.SuitesGreen != 1 {
		t.Errorf("suites expected/green = %d/%d, want 2/1 (private never ran)", r.SuitesExpected, r.SuitesGreen)
	}
	if r.MergedAllGreen == nil || *r.MergedAllGreen != 0 || r.AllGreenAt != nil {
		t.Errorf("upgrade suite ran but never passed, must not be all green: %+v", r)
	}

	// Created after the merge gate but before the private suite was required.
	required = defs.requiredSuites(wanted, ts("2026-09-01T06:00:00Z"))
	if !required["capa/standard"] || !required["capa/upgrade"] || required["capa/private"] {
		t.Errorf("unexpected required suites: %v", required)
	}
	required = defs.requiredSuites(wanted, ts("2026-09-15T00:00:00Z"))
	if !required["capa/private"] {
		t.Errorf("private suite must be required from its required-since date")
	}
}

func TestApplyTestResultsIncomplete(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"capa"})
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z")}
	res := testResults{
		runs:       []issueComment{comment("Gacko", "/run releases-test-suites", "2026-09-09T08:00:00Z")},
		waivers:    map[string]time.Time{},
		firstGreen: map[string]time.Time{"Release Tests / CAPA Standard Suite": ts("2026-09-09T12:00:00Z")},
	}
	applyTestResults(&r, wanted, allRequired(wanted), res)
	if r.AllGreenAt != nil || r.TimeToGreenDays != nil || r.TestFrictionDays != nil || r.RunsToGreen != nil {
		t.Errorf("incomplete suites must not produce an all-green time: %+v", r)
	}
	if r.MergedAllGreen == nil || *r.MergedAllGreen != 0 {
		t.Errorf("merged all green = %v, want 0", r.MergedAllGreen)
	}
	if r.SuitesGreen != 1 || r.TestRuns != 1 || r.TestsFirstRunAt == nil {
		t.Errorf("partial counts missing: %+v", r)
	}
}

func TestApplyTestResultsNoSuitesExpected(t *testing.T) {
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z")}
	applyTestResults(&r, map[string]string{}, map[string]bool{}, testResults{waivers: map[string]time.Time{}, firstGreen: map[string]time.Time{}})
	if r.AllGreenAt != nil || r.SuitesExpected != 0 || r.MergedAllGreen != nil {
		t.Errorf("no expected suites must not produce an all-green time: %+v", r)
	}
}

func TestWaiverPattern(t *testing.T) {
	body := "some text\n/waive-suite ./providers/capz/private no private env\n/waive-suite capa/china   blocked\n"
	m := waiverRe.FindAllStringSubmatch(body, -1)
	if len(m) != 2 || m[0][1] != "capz" || m[0][2] != "private" || m[1][1] != "capa" || m[1][2] != "china" {
		t.Errorf("unexpected waiver matches: %v", m)
	}
}

func TestIsGreen(t *testing.T) {
	for conclusion, want := range map[string]bool{"success": true, "neutral": true, "failure": false, "cancelled": false, "timed_out": false, "": false} {
		if got := isGreen(conclusion); got != want {
			t.Errorf("isGreen(%q) = %v, want %v", conclusion, got, want)
		}
	}
}
