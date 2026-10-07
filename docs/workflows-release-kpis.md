# Release KPIs

This document describes how release process KPIs are collected and presented.
The first KPI is **lead time for release**: how long a release PR stays open before it is merged.

Towards https://github.com/giantswarm/roadmap/issues/4382

## Overview

Lead time is computed from the pull requests that create releases, because they have a clear
start (PR opened) and end (PR merged). A release PR is a merged PR that added a new
`<provider>/vX.Y.Z/release.yaml` file on `master`. Release PRs are found in the git history rather
than by title or label, because neither has been consistent over time (e.g. `CAPA: Release 30.1.3`,
`release 25.5.4`, `Releases for vsphere`). PRs that only move a release back from `archived/` are
not counted. Releases below v25, the first CAPI major version, are ignored because the same
directories also hold vintage releases.

For every release PR the data set records:

| Field | Description |
|-------|-------------|
| `number`, `url`, `title`, `author` | The pull request |
| `provider` | From the release directory: `CAPA`, `CAPZ`, `CAPV`, `CAPVCD`, `EKS`, `AKS`, `CAPMOX`, or `CAPI` when the PR released several providers at once |
| `version`, `release_type` | Release version and `major` / `minor` / `patch`, taken from the `release/*` label or, for older PRs without a label, derived from the version. A PR releasing several versions reports the one shared by most providers |
| `created_at`, `merged_at` | When the PR was opened and merged |
| `lead_time_days` | `merged_at - created_at` in days |
| `automated` | `true` when the PR was created by the scheduled automation on the 1st of the month |
| `planned_merge_date` | Only for automated releases: the next 1st of the month after the PR was created |
| `delay_days` | Only for automated releases: `merged_at - planned_merge_date` in days, negative when merged early |

Manually created releases are merged as soon as possible and have no planned date, so the planned
merge date and the delay stay empty for them.

## Workflow: Release KPIs (`release-kpis.yaml`)

**Triggers:**
- Daily at 7:00 AM UTC
- Manual trigger via `workflow_dispatch`
- When a release PR is merged
- On pushes to `master` that change the workflow or `tools/release-kpis/`

**Steps:**
1. Builds `tools/release-kpis` and runs it: release PRs come from the git history, their dates from the GitHub API
2. Writes a summary table of the last 10 releases to the job summary
3. Embeds the data into the dashboard template and uploads the dashboard to Grafana Cloud

The data is stored inside the dashboard itself, as inline data of the Infinity datasource query of
the table panel. The other panels read that panel's result through Grafana's built-in Dashboard
datasource, so the data is embedded only once. This keeps generated data out of the repository and
needs nothing besides the Grafana API key.
The dashboard is overwritten on every run, so changes to the panels have to be made in
[`tools/release-kpis/dashboard.json`](../tools/release-kpis/dashboard.json), not in the Grafana UI.

**Required Secrets:**
- `GITHUB_TOKEN`: Reading pull requests
- `GRAFANA_API_KEY`: Creating and updating the dashboard, needs the Editor role

## Planned merge date

Scheduled releases are created by `create-release.yaml` on the 1st of every month and are planned
to be merged on the next 1st. For those PRs the workflow now adds a line to the PR body:

```
📅 **Planned merge date:** 2026-11-01 (scheduled release)
<!-- PLANNED_MERGE_DATE: 2026-11-01 -->
```

The HTML comment is what the KPI tool reads. PRs created before this marker existed are
recognised by their creation time instead: created by `taylorbot` on the 1st of the month between
06:00 and 09:00 UTC, which is the window of the scheduled run. Manually triggered runs by the same
bot fall outside that window.

## Grafana dashboard

The dashboard template lives in [`tools/release-kpis/dashboard.json`](../tools/release-kpis/dashboard.json)
and is published as [CAPI Release KPIs](https://giantswarm.grafana.net/d/capi-release-kpis/capi-release-kpis).
It uses the Infinity datasource in Grafana Cloud, the same datasource the
[CAPI Releases Dashboard](https://giantswarm.grafana.net/d/be9a0bh8mbwn4e/capi-releases) uses to
read `<provider>/releases.json`, so no new datasource is needed.

Panels:
- Median lead time per release type (major, minor, patch)
- Median delay from the planned merge date (scheduled releases only)
- Lead time per release over time, split by release type
- Delay from the planned merge date over time (scheduled releases only)
- Table of all merged release PRs with links

To change the dashboard, edit the template and merge it. The workflow publishes it on the next
push to `master`. To try a change before merging, run the workflow steps by hand:

```bash
cd tools && go build -o release-kpis ./release-kpis && cd ..
GITHUB_TOKEN=$(gh auth token) tools/release-kpis -repo-dir . -ref origin/master -output /tmp/release-lead-time.json
jq --slurpfile data /tmp/release-lead-time.json \
  '(.panels[].targets[] | select(.datasource.type == "yesoreyeram-infinity-datasource"))
     |= (.source = "inline" | .data = ($data[0] | tojson))
   | {dashboard: ., folderUid: "", overwrite: true}' tools/release-kpis/dashboard.json \
  | curl --silent --fail-with-body -X POST https://giantswarm.grafana.net/api/dashboards/db \
      -H "Authorization: Bearer $GRAFANA_API_KEY" -H "Content-Type: application/json" --data @-
```

Further KPIs from the investigation in [roadmap#4336](https://github.com/giantswarm/roadmap/issues/4336)
are meant to land on the same dashboard.

## Running the tool locally

```bash
cd tools
go build -o release-kpis ./release-kpis
GITHUB_TOKEN=$(gh auth token) ./release-kpis -repo-dir .. -ref origin/master -verbose -output release-lead-time.json
```

The tool fetches every release PR from the GitHub API, so a token is needed to stay within the rate limit.

## Troubleshooting

**No data or an outdated dashboard in Grafana:**
- Check the latest run of the Release KPIs workflow
- Verify `GRAFANA_API_KEY` is valid and has the Editor role, uploading needs dashboard write access

**A release is missing:**
- The PR must be merged and must have added a new `<provider>/vX.Y.Z/release.yaml` file
- The merge commit subject must end with the PR number, e.g. `(#2426)`, as squash merges do; otherwise the PR is looked up through the commit

**A scheduled release has no planned merge date:**
- The PR body must contain the `PLANNED_MERGE_DATE` marker, or
- for older PRs, it must have been created by `taylorbot` on the 1st between 06:00 and 09:00 UTC
