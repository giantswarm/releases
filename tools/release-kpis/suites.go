package main

import (
	"sort"
	"time"
)

// Failure rate by suite.
//
// Every completed "Release Tests / <suite>" check run of a release PR is
// classified by its conclusion and counted per suite, so that the dashboard
// can show which suites fail most often and which ones are the last to go
// green. The rows are per release and suite; the aggregation happens in
// Grafana so that the time range and provider filters apply.
//
// See https://github.com/giantswarm/roadmap/issues/4398

// SuiteResult is the outcome of one E2E suite on one release PR.
type SuiteResult struct {
	Number        int       `json:"number"`
	Name          string    `json:"name"` // "CAPI v36.1.0", the release
	Provider      string    `json:"provider"`
	ReleaseType   string    `json:"release_type"`
	MergedAt      time.Time `json:"merged_at"`
	MergedQuarter string    `json:"merged_quarter"`
	Suite         string    `json:"suite"` // "capa/china"
	Check         string    `json:"check"` // the check run name

	// Completed check runs of the suite on the PR, by conclusion. Runs is
	// their sum. Failed is 1 when the suite failed at least once, so that its
	// mean is the share of releases on which the suite failed.
	Runs      int `json:"runs"`
	Successes int `json:"successes"`
	Failures  int `json:"failures"`
	Cancelled int `json:"cancelled"`
	Skipped   int `json:"skipped"`
	Failed    int `json:"failed"`
	Waived    int `json:"waived"`

	// FirstGreenAt is the suite's first pass, or its waiver. RunsToGreen is
	// the number of runs up to and including that pass. LastToGreen is 1 when
	// this suite's first green set the PR's all-green time, that is, it was the
	// suite the release waited for.
	FirstGreenAt *time.Time `json:"first_green_at"`
	RunsToGreen  *int       `json:"runs_to_green"`
	LastToGreen  int        `json:"last_to_green"`
}

// Outcome classes of a check run conclusion.
const (
	outcomeSuccess   = "success"
	outcomeFailure   = "failure"
	outcomeCancelled = "cancelled"
	outcomeSkipped   = "skipped"
)

// outcome classifies a check run conclusion. In-progress runs and unknown
// conclusions return "" and are not counted.
func outcome(conclusion string) string {
	switch conclusion {
	case "success":
		return outcomeSuccess
	case "neutral", "skipped":
		return outcomeSkipped
	case "cancelled":
		return outcomeCancelled
	case "failure", "timed_out", "action_required", "stale":
		return outcomeFailure
	}
	return ""
}

// suiteResults returns one row per suite that counts for the release, with
// the same rule as applyTestResults: required suites, plus suites that ran
// or were waived anyway. It must be called after applyTestResults so that
// r.AllGreenAt is set.
func suiteResults(r Release, wanted map[string]string, required map[string]bool, res testResults) []SuiteResult {
	var rows []SuiteResult
	for key, name := range wanted {
		waivedAt, waived := res.waivers[key]
		if !required[key] && !res.ran[name] && !waived {
			continue
		}
		row := SuiteResult{
			Number:        r.Number,
			Name:          r.Name,
			Provider:      r.Provider,
			ReleaseType:   r.ReleaseType,
			MergedAt:      r.MergedAt,
			MergedQuarter: r.MergedQuarter,
			Suite:         key,
			Check:         name,
		}

		runs := append([]checkRun(nil), res.results[name]...)
		sort.Slice(runs, func(i, j int) bool { return runs[i].CompletedAt.Before(*runs[j].CompletedAt) })
		for _, run := range runs {
			switch outcome(run.Conclusion) {
			case outcomeSuccess:
				row.Successes++
			case outcomeFailure:
				row.Failures++
			case outcomeCancelled:
				row.Cancelled++
			case outcomeSkipped:
				row.Skipped++
			default:
				continue
			}
			row.Runs++
			if row.FirstGreenAt == nil && isGreen(run.Conclusion) {
				t := run.CompletedAt.UTC()
				row.FirstGreenAt = &t
				n := row.Runs
				row.RunsToGreen = &n
			}
		}
		if row.Failures > 0 {
			row.Failed = 1
		}
		if waived {
			row.Waived = 1
			if row.FirstGreenAt == nil {
				t := waivedAt.UTC()
				row.FirstGreenAt = &t
			}
		}
		if r.AllGreenAt != nil && row.FirstGreenAt != nil && row.FirstGreenAt.Equal(*r.AllGreenAt) {
			row.LastToGreen = 1
		}
		rows = append(rows, row)
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Suite < rows[j].Suite })
	return rows
}
