# E2E Test Coverage

This document describes the `E2E Coverage` merge gate (`e2e-coverage.yaml`) and the `/waive-suite` slash command.

## Overview

`releases-test-suites` picks which test suites to run based on the release stage, so a PR in `stage/development` only runs `standard`, `upgrade` and `upgrade-major`. The variant suites (`private`, `china`, `cilium-eni-mode`, `on-capa`, `on-capz`) only start automatically once the PR reaches `stage/freeze`.

Releases were being merged before reaching Freeze, which meant those variant suites never ran at all — v34.5.0 was merged from `stage/development` having tested only 12 of the 18 suites expected for its providers.

The `E2E Coverage` check closes that gap. It requires the **full** suite set for every newly added release in the PR, whatever stage the PR is currently in, and only reports success once all of them have passed.

See [giantswarm/roadmap#4334](https://github.com/giantswarm/roadmap/issues/4334).

## What is required

Expected suites per provider directory:

| Provider directory | Expected suites |
|--------------------|-----------------|
| `capa` | `standard`, `upgrade`, `upgrade-major`, `cilium-eni-mode`, `china`, `private` |
| `azure` | `standard`, `upgrade`, `upgrade-major`, `private` |
| `vsphere` | `standard`, `upgrade`, `upgrade-major`, `on-capa`, `on-capz` |
| `cloud-director` | `standard`, `upgrade`, `upgrade-major` |
| `eks` | `standard`, `upgrade`, `upgrade-major` |
| `aks`, `proxmox` | _none_ — no `releases-test-suites` mapping exists yet |

This mirrors what `releases-test-suites` expands to during Freeze. A suite counts as covered when its per-suite check run (`Release Tests / <suite description>`, published by the `check-run-results-to-pr` Tekton task) has concluded successfully.

Only **newly added** releases are considered. A PR that deprecates (modifies) or archives (renames) an existing `release.yaml` requires no coverage and reports `No new releases added`.

## How coverage is evaluated

Coverage is tied to the release **content**, not to the head commit:

- A suite that passed on an earlier commit of the PR still counts, as long as no newly added `release.yaml` has changed since.
- Editing a README, announcement or `release.diff` does **not** invalidate results.
- Changing a component or app version — via `/update-release`, `/pin-version`, or the weekly `bump-open-releases.yaml` job — **does** invalidate them, because the suites then tested different content. They have to run again.

Reaching Freeze early therefore protects coverage: `stage/freeze` restricts `/update-release` to Team Tenet and excludes the PR from the weekly bump, so results stop being invalidated underneath you.

Within one release content, a successful run is never overridden by a later failing run of the same suite — a flaky re-run cannot pull coverage back down. A genuinely failing run is still visible through the `Release Test Suites` check, which the PR gatekeeper requires separately.

## Stage interaction

The gate does not require the PR to pass through any particular stage. It reports the current stage and which suites a plain `/run releases-test-suites` covers there, because that is the usual reason suites are missing:

- While most of what the current stage runs is still outstanding, the check suggests a plain `/run releases-test-suites`, and separately lists the suites that stage will never run.
- Once the stage's suites are mostly green, it lists only the outstanding suites as an explicit `TARGET_SUITES` command, so already-passed suites are not re-run.
- It suggests the next stage command (`/stage active` from Development, `/stage freeze` from Active). Stages advance one step at a time, so `/stage freeze` directly from Development is rejected.

Switching stage does not affect existing results: labels do not change the release content.

## Commands

### `/waive-suite <provider>/<suite> <reason>`

Marks a single suite as covered when it genuinely cannot pass — for example because its test environment is down.

```
/waive-suite capa/china Beijing environment down, see giantswarm/roadmap#1234
```

**Rules:**

- A reason is **required**. `/waive-suite capa/china` with no reason is ignored.
- The comment author must be an organisation member, owner or collaborator.
- The provider can be written either as the release directory or as the cluster-test-suites name (`azure/private` and `capz/private` are equivalent), with or without a `./providers/` prefix.
- The waiver applies to that one suite only; everything else stays enforced.
- The waiver is recorded in the check output (`⚠️ waived by @user`) and becomes **void** as soon as a release in the PR changes, exactly like a test result.

## Check states

| State | Meaning |
|-------|---------|
| ✅ success | Every expected suite passed or is waived |
| 🚧 in progress | Suites are still missing or running — blocks merge without marking the PR as broken |
| ❌ failure | An expected suite failed |

Per-suite rows are reported as `✅ passed`, `⚠️ waived`, `❌ <conclusion>`, `🚧 in progress` or `⬜ not run`. Suites the current stage does not run are marked `runs automatically at stage/freeze`, so a missing result is not mistaken for a problem.

## Enforcement

The check run is required by [pr-gatekeeper](https://github.com/giantswarm/pr-gatekeeper) (Heimdall) for this repo, and Heimdall is a required status check on `master`.

Note that a `/skip-ci <reason>` comment bypasses **all** Heimdall requirements, including this one. When that happens Heimdall records who skipped it and why, and the `E2E Coverage` output still shows which suites never ran.

To make coverage impossible to skip, add the `E2E Coverage` **commit status** to the required status checks on `master`. Branch protection evaluates it directly, so `/skip-ci` — which only affects Heimdall's own verdict — cannot bypass it.

## Workflow details

### Triggers

```yaml
on:
  pull_request:
    types: [opened, synchronize, reopened, ready_for_review, labeled, unlabeled]
  check_run:
    types: [completed]
  issue_comment:
    types: [created]
  workflow_dispatch:
```

The `check_run` trigger re-evaluates coverage as each suite finishes, filtered to `Release Tests / *` so the workflow does not react to its own check run. The `issue_comment` trigger only runs for comments containing `/waive-suite`.

`check_run` and `issue_comment` always use the workflow file from the default branch, so changes only take effect once merged. Use `workflow_dispatch` with a PR number to evaluate a PR on demand.

### What you see in the Checks tab

The result is published twice, so the workflow contributes three rows:

| Row | Meaning |
|-----|---------|
| `E2E Coverage` (commit status) | The gate, as a top-level row next to the `ci/circleci: ...` rows. Shows the tally, e.g. `12/21 expected suites covered · stage/development`, and links to the full report. |
| `gitleaks / E2E Coverage` (check run) | The same result with the full per-provider report. This is the check `pr-gatekeeper` requires. |
| `E2E Coverage / Publish coverage report` | The job that ran the script. Green just means the evaluation completed. |

Why both: a check run can hold a markdown report but cannot control where it appears. GitHub adopts API-created check runs into an existing check suite for the same app and commit, and names the group after the first workflow that ran in it — so the check ends up filed under an unrelated workflow such as `gitleaks`, which is easy to miss. A commit status is always its own row and can be required in branch protection directly, but carries only a 140-character description, so it links to the check run for the detail.

Status states map to the check states above: `pending` while suites are missing, `success` when complete, `failure` when an expected suite failed.

Unrelated to this workflow, `gitleaks` also appears twice on every PR in this repo, because `zz_generated.gitleaks.yaml` is generated with both `push` and `pull_request` triggers.

### Implementation

The logic lives in `.github/scripts/e2e-coverage.js` and talks to the REST API directly rather than through `actions/github-script`, whose pinned SHA trips the repo's gitleaks rule (it matches any long alphanumeric run next to the word "github").

The check run is always updated in place rather than recreated, because `pr-gatekeeper` refuses to evaluate a check name that matches several runs on the same commit.

## Related Files

- `.github/workflows/e2e-coverage.yaml`: Workflow definition
- `.github/scripts/e2e-coverage.js`: Coverage evaluation and check run publishing
- `.github/actions/get-providers/action.yml`: Provider directory → CAPI name mapping
- `docs/workflows-release-stages.md`: Stage lifecycle and what each stage runs
- [`tekton-resources/tekton-pipelines/pipelines/releases-test-suites.yaml`](https://github.com/giantswarm/tekton-resources/blob/main/tekton-resources/tekton-pipelines/pipelines/releases-test-suites.yaml): Stage-based suite selection
- [`tekton-resources/tekton-pipelines/tasks/check-run-results-to-pr.yaml`](https://github.com/giantswarm/tekton-resources/blob/main/tekton-resources/tekton-pipelines/tasks/check-run-results-to-pr.yaml): Publishes the per-suite check runs
- [`pr-gatekeeper/config.yaml`](https://github.com/giantswarm/pr-gatekeeper/blob/main/config.yaml): Required checks per repo
