# Retagging Cluster Provider Charts

This document describes the automated workflow that creates `release-<provider>` charts from `cluster-<provider>` charts with Release versions.

## Problem Statement

Flux cannot easily modify a specific field in a ConfigMap (`global.release.version`). The previous flow required:
1. User updates `global.release.version` in ConfigMap
2. Mutating webhook updates App's `.spec.version`
3. Chart does runtime lookup of Release CR

This workflow enables the App CR's `.spec.version` field to be the **actual** version being deployed, enabling standard Helm/Flux upgrade patterns.

## How It Works

When a Release (e.g., `v34.0.0`) references `cluster-aws:7.2.5`, this workflow:

1. **Pulls** the original chart `cluster-aws:7.2.5` from `gsoci.azurecr.io/charts/giantswarm`
2. **Renames** the chart from `cluster-aws` to `release-aws`
3. **Updates** `Chart.yaml` version to `34.0.0`
4. **Injects** `global.release.version: "34.0.0"` into `values.yaml`
5. **Packages** with App Build Suite (adds metadata and validation)
6. **Pushes** `release-aws:34.0.0` to both:
   - OCI registry: `gsoci.azurecr.io/charts/giantswarm`
   - GitHub Pages catalog: `cluster-catalog` (required while App/Chart CRs are still in use)

### Why Rename to release-<provider>?

1. **No version collisions**: Original chart versions (7.2.5) and Release versions (34.0.0) coexist without conflicts
2. **OCI registry ready**: When migrating to pure OCI (no catalogs), different chart names prevent conflicts
3. **Webhook compatibility**: The mutating webhook only triggers for charts starting with `cluster-`, so `release-*` charts bypass mutation

### Why Inject global.release.version?

The `cluster` subchart (bundled inside provider charts) uses `.Chart.Version` to look up the Release CR. However, in Helm subcharts, `.Chart.Version` refers to the subchart's version (`5.1.2`), not the parent chart's version (`34.0.0`).

The subchart's lookup logic in `_helpers.tpl` already checks `$.Values.global.release.version` first, then falls back to `.Chart.Version`. By injecting this value, the subchart correctly resolves the Release CR (e.g., `aws-34.0.0`).

## Triggers

### Automatic (on merge to main/master)

When `release.yaml` files are changed and merged, the workflow automatically creates release charts.

**Filters:**
- Excludes archived releases (paths containing `/archived/`)
- Only processes releases with `state: active` (skips deprecated)

### Manual (via CircleCI)

**Via UI:**
1. Go to [CircleCI project](https://app.circleci.com/pipelines/github/giantswarm/releases)
2. Click "Trigger Pipeline"
3. Add parameters:
   - `provider`: `capa`, `azure`, `vsphere`, `cloud-director`, or `eks`
   - `release_version`: e.g., `v34.0.0`

**Via API:**
```bash
curl -X POST https://circleci.com/api/v2/project/gh/giantswarm/releases/pipeline \
  -H "Circle-Token: $CIRCLE_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"parameters": {"provider": "capa", "release_version": "v34.0.0"}}'
```

## Provider to Chart Mapping

| Provider Directory | Source Chart | Target Chart |
|-------------------|--------------|--------------|
| `capa` | `cluster-aws` | `release-aws` |
| `azure` | `cluster-azure` | `release-azure` |
| `vsphere` | `cluster-vsphere` | `release-vsphere` |
| `cloud-director` | `cluster-cloud-director` | `release-cloud-director` |
| `eks` | `cluster-eks` | `release-eks` |

## Related

- [Issue: Enable Automated Cluster Upgrades](https://github.com/giantswarm/giantswarm/issues/33311)
- [Roadmap: Flux Migration](https://github.com/giantswarm/roadmap/issues/4184)
