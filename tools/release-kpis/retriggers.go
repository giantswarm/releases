package main

import "time"

// Retrigger rate.
//
// A retrigger is a "/run releases-test-suites" comment after the first one on
// a release PR. Each retrigger is classified by what happened since the
// previous run: a suite failed, new commits were pushed, or neither, in which
// case the tests were rerun while green or still running. Retriggers posted
// by the automation (first run on the active stage, runs after automatic
// updates) are kept apart since nobody decided to rerun.
//
// On the suite level, a rerun is every completed run of a suite after its
// first one on the PR. It is forced when the previous run failed or was
// cancelled, caused by a new commit when the previous run tested another
// commit, and redundant when the suite had already passed on the same commit.
// Redundant reruns are the cost of rerunning the whole matrix instead of the
// suites that failed.
//
// See https://github.com/giantswarm/roadmap/issues/4406

// retriggerCounts are the retriggers of one release PR by reason.
type retriggerCounts struct {
	total, automated, afterFailure, afterCommit, other int
}

// classifyRetriggers classifies every run comment after the first. results
// are the completed suite runs of the PR and commits the times its commits
// were pushed.
func classifyRetriggers(runs []issueComment, results map[string][]checkRun, commits []time.Time) retriggerCounts {
	var c retriggerCounts
	for i := 1; i < len(runs); i++ {
		prev, cur := runs[i-1].CreatedAt, runs[i].CreatedAt
		c.total++
		switch {
		case isAutomation(runs[i].User.Login):
			c.automated++
		case failedBetween(results, prev, cur):
			c.afterFailure++
		case anyBetween(commits, prev, cur):
			c.afterCommit++
		default:
			c.other++
		}
	}
	return c
}

// failedBetween reports whether a suite run failed after from and up to to.
// Cancelled runs do not count: they are the result of a retrigger, not a
// reason for one.
func failedBetween(results map[string][]checkRun, from, to time.Time) bool {
	for _, runs := range results {
		for _, run := range runs {
			if outcome(run.Conclusion) == outcomeFailure && between(*run.CompletedAt, from, to) {
				return true
			}
		}
	}
	return false
}

func anyBetween(times []time.Time, from, to time.Time) bool {
	for _, t := range times {
		if between(t, from, to) {
			return true
		}
	}
	return false
}

// between reports whether t is after from and not after to.
func between(t, from, to time.Time) bool {
	return t.After(from) && !t.After(to)
}

// suiteReruns classifies the reruns of one suite. runs are its completed runs
// on the PR, oldest first.
func suiteReruns(runs []checkRun) (reruns, afterFailure, afterCommit, redundant int) {
	for i := 1; i < len(runs); i++ {
		prev := runs[i-1]
		reruns++
		switch {
		case outcome(prev.Conclusion) == outcomeFailure || outcome(prev.Conclusion) == outcomeCancelled:
			afterFailure++
		case prev.HeadSHA != runs[i].HeadSHA:
			afterCommit++
		default:
			redundant++
		}
	}
	return reruns, afterFailure, afterCommit, redundant
}

// applySuiteTotals sums the suite rows of a release into its cost fields.
func applySuiteTotals(r *Release, rows []SuiteResult) {
	for _, row := range rows {
		r.SuiteRuns += row.Runs
		r.SuiteReruns += row.Reruns
		r.SuiteRerunsRedundant += row.RerunsRedundant
	}
}
