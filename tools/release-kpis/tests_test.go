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
			"aks":            {},
		},
		CheckNameFormats: map[string]string{
			"standard": "{capi} Standard Suite",
			"upgrade":  "{capi} Upgrade Suite",
			"private":  "{capi} Private Suite",
		},
		SuitePathProviders: map[string]string{"capa": "capa", "azure": "capz", "cloud-director": "capvcd"},
	}
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
	if n := len(testDefs().expectedChecks([]string{"aks"})); n != 0 {
		t.Errorf("aks expects %d checks, want 0", n)
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
			comment("Gacko", "/run releases-test-suites", "2026-09-11T10:00:00Z"),
		},
		waivers: map[string]time.Time{"capa/private": ts("2026-09-12T12:00:00Z")},
		firstGreen: map[string]time.Time{
			"Release Tests / CAPA Standard Suite": ts("2026-09-09T12:00:00Z"),
			"Release Tests / CAPA Upgrade Suite":  ts("2026-09-11T18:00:00Z"),
		},
	}
	applyTestResults(&r, wanted, res)

	if r.SuitesExpected != 3 || r.SuitesGreen != 2 || r.SuitesWaived != 1 {
		t.Errorf("suites expected/green/waived = %d/%d/%d, want 3/2/1", r.SuitesExpected, r.SuitesGreen, r.SuitesWaived)
	}
	if r.TestRuns != 3 || r.TestRunsAutomated != 1 {
		t.Errorf("test runs = %d (%d automated), want 3 (1)", r.TestRuns, r.TestRunsAutomated)
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
}

func TestApplyTestResultsIncomplete(t *testing.T) {
	wanted := testDefs().expectedChecks([]string{"capa"})
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z")}
	res := testResults{
		runs:       []issueComment{comment("Gacko", "/run releases-test-suites", "2026-09-09T08:00:00Z")},
		waivers:    map[string]time.Time{},
		firstGreen: map[string]time.Time{"Release Tests / CAPA Standard Suite": ts("2026-09-09T12:00:00Z")},
	}
	applyTestResults(&r, wanted, res)
	if r.AllGreenAt != nil || r.TimeToGreenDays != nil || r.TestFrictionDays != nil {
		t.Errorf("incomplete suites must not produce an all-green time: %+v", r)
	}
	if r.SuitesGreen != 1 || r.TestRuns != 1 || r.TestsFirstRunAt == nil {
		t.Errorf("partial counts missing: %+v", r)
	}
}

func TestApplyTestResultsNoSuitesExpected(t *testing.T) {
	r := Release{CreatedAt: ts("2026-09-01T06:00:00Z")}
	applyTestResults(&r, map[string]string{}, testResults{waivers: map[string]time.Time{}, firstGreen: map[string]time.Time{}})
	if r.AllGreenAt != nil || r.SuitesExpected != 0 {
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
