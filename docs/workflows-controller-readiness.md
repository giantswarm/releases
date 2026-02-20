# Controller Readiness Tracking

This document describes the automated controller readiness tracking system for major releases. It covers the composite action that creates tracking issues, the daily readiness check workflow, and how major release PRs are held in development until controllers are updated.

## Overview

When a major release is created (e.g., v35.0.0 for a new Kubernetes version), cloud controller apps need new versions compatible with that Kubernetes version. This system automates the notification, tracking, and follow-up:

1. **Issue creation** -- tracking issues are created in `giantswarm/roadmap` for the responsible teams
2. **PR tracking** -- a structured comment on the release PR links to the issues
3. **Bump protection** -- the weekly bump job skips major releases in development
4. **Daily readiness check** -- once all issues are closed, the release PR automatically receives a component bump

## Provider-to-Team Mapping

| Provider | Team | Controller Apps |
|----------|------|-----------------|
| aws | team-phoenix | `cloud-provider-aws` |
| azure | team-phoenix | `azure-cloud-controller-manager`, `azure-cloud-node-manager` |
| vsphere | team-rocket | `cloud-provider-vsphere` |
| cloud-director | team-rocket | `cloud-provider-cloud-director` |
| eks | _(excluded)_ | No cloud controller manager needed |

One issue is created per team (not per provider). If a release includes both aws and azure, team-phoenix gets a single issue listing both.

## Components

### 1. Composite Action (`create-controller-readiness-issues`)

**Path**: `.github/actions/create-controller-readiness-issues/action.yml`

Called automatically after a major release PR is created (both consolidated and individual paths).

**Inputs**:

| Input | Required | Description |
|-------|----------|-------------|
| `providers` | yes | Space-separated provider list |
| `version` | yes | Release version (e.g. `v35.0.0`) |
| `pr_numbers` | yes | Comma-separated PR numbers |
| `github_token` | yes | GitHub token for API access |

**What it does**:

1. Maps providers to teams (see table above)
2. For each team, checks for an existing open issue (deduplication)
3. Creates a new issue in `giantswarm/roadmap` with:
   - Labels: `controller-readiness` + team label (e.g. `team/phoenix`)
   - Checklist of controller apps to update
   - CC to the team
4. Adds the issue to GitHub Project #273 (Roadmap board) with Team and Status fields
5. Posts a tracking comment on the release PR:

```markdown
<!-- CONTROLLER_READINESS_START -->
## Controller readiness tracker

| Team | Providers | Issue | Status |
|------|-----------|-------|--------|
| team-phoenix | aws, azure | giantswarm/roadmap#123 | open |
| team-rocket | vsphere | giantswarm/roadmap#124 | open |
<!-- CONTROLLER_READINESS_END -->
```

If the action runs again (e.g. workflow re-run), it finds the existing issues and updates the tracking comment instead of creating duplicates.

### 2. Bump Protection (`bump-open-releases.yaml`)

The weekly bump job (`bump-open-releases.yaml`) excludes PRs that have **both** `release/major` and `stage/development` labels. This prevents the weekly job from bumping components before controllers are ready.

- Minor and patch releases in development are still bumped normally
- Manual runs with a specific PR number bypass this filter entirely
- Once the PR moves out of development (e.g. to active), it will be included in weekly bumps again

### 3. Daily Readiness Check (`check-controller-readiness.yaml`)

**Path**: `.github/workflows/check-controller-readiness.yaml`

**Schedule**: Daily at 08:00 UTC + manual dispatch

**What it does**:

1. Finds all open PRs with labels `release/major` + `stage/development`
2. For each PR, looks for the `<!-- CONTROLLER_READINESS_START -->` tracking comment
3. Extracts issue references (`giantswarm/roadmap#NNN`) from the comment
4. Checks the state of each issue via `gh issue view`
5. If **all** issues are closed:
   - Posts `/update-release` on the PR (triggers a full component bump with new controller versions)
   - Updates the tracking comment statuses from "open" to "closed"

## Lifecycle

```
Major release PR created
    │
    ├─ Controller readiness issues created in giantswarm/roadmap
    ├─ Tracking comment posted on PR
    ├─ Weekly bump skipped (major + development)
    │
    │  ... teams update controller apps ...
    │
    ├─ Teams close their readiness issues
    │
    ├─ Daily check detects all issues closed
    ├─ /update-release posted → component bump picks up new controllers
    │
    └─ PR ready for stage progression (Development → Active → Freeze → Merge)
```

## Troubleshooting

### Issues not created

- Check that providers other than `eks` are included in the release
- Check the workflow logs for the "Create controller readiness issues" step
- Verify the `controller-readiness` label exists in `giantswarm/roadmap`

### Readiness check not triggering bump

- Ensure the PR has **both** `release/major` and `stage/development` labels
- Ensure a tracking comment with `CONTROLLER_READINESS_START` exists on the PR
- Ensure all referenced issues are actually closed (not just checked off)
- Run the workflow manually via workflow dispatch to test

### Weekly bump still running on major development PR

- Verify the PR has both `release/major` and `stage/development` labels
- The filter only applies to scheduled runs; manual runs with a specific PR number always process

## Related Files

- `.github/actions/create-controller-readiness-issues/action.yml`: Composite action
- `.github/workflows/check-controller-readiness.yaml`: Daily readiness check
- `.github/workflows/create-release.yaml`: Calls the composite action after major PR creation
- `.github/workflows/bump-open-releases.yaml`: Excludes major development PRs
