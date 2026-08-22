# Verifying CAPI Release apps against provider cluster charts

This document describes how to manually (or programmatically) verify that a
provider's Release CR (`release.yaml` in this repository) lists exactly the
apps — with exactly the right `dependsOn` — that the provider's cluster
chart would actually render, given its current default configuration.

It is meant to be reproducible by a human or by an agent with no other
context than this file and read access to the repositories below.

## Repositories involved

This process reads from three kinds of repositories, expected to be checked
out locally as siblings (e.g. all directly under the same parent directory
as this `releases` repository):

- **`releases`** (this repo) — contains the Release CRs we are verifying.
- **`cluster`** — the shared Helm chart included by every provider cluster
  chart. Defines the default set of apps common to all providers.
- **`cluster-<provider>`** (e.g. `cluster-aws`, `cluster-azure`,
  `cluster-eks`, `cluster-cloud-director`, `cluster-vsphere`, ...) — one repo
  per supported provider ("provider cluster chart"), defining
  provider-specific defaults and provider-specific apps.

All of these repos should be checked out on the same feature/release branch
when comparing against an *in-progress* (not yet released) version — e.g.
at the time of writing, all provider cluster charts and the shared `cluster`
chart were on a branch called `migrate-apps-helmreleases`.

### Provider name → directory mapping

The directory name inside `releases/` matches the provider identifier used
throughout the tooling, with one exception:

| Provider identifier | `releases/` directory | Provider chart repo |
|---|---|---|
| `aws`  | `capa`           | `cluster-aws` |
| `eks`  | `eks`            | `cluster-eks` |
| `azure`| `azure`          | `cluster-azure` |
| `cloud-director` | `cloud-director` | `cluster-cloud-director` |
| `vsphere` | `vsphere`     | `cluster-vsphere` |
| `aks`  | `aks`            | *(not verified in this pass — assumed analogous)* |
| `proxmox` | `proxmox`     | *(not verified in this pass — assumed analogous)* |

Only AWS breaks the `cluster-<provider>` / `releases/<provider>` naming
convention (`capa` vs. `aws`). Everything else follows
`cluster-<provider>` ↔ `releases/<provider>`.

## Terminology

- **Provider cluster chart**: a Helm chart named `cluster-<provider>`,
  living in the repo of the same name, under `helm/cluster-<provider>/`.
  Renders Cluster API resources, Flux `HelmRelease` CRs, etc. for one
  provider, using the shared `cluster` chart as a dependency.
- **`cluster` chart**: the shared chart (repo `cluster`, chart at
  `helm/cluster/`) included by every provider cluster chart. Renders the
  default set of apps common to (most) providers via Go-templated
  pseudo-YAML files.
- **`configKey`**: the key used in a provider chart's `values.yaml` under
  `cluster.providerIntegration.apps.<configKey>` to enable/disable and
  configure one app rendered by the shared `cluster` chart.
- **`appName`**: the canonical app name as it appears in a HelmRelease
  definition and, ultimately, in a Release CR's `spec.apps[].name`.
- **HelmRelease CR**: a Flux `HelmRelease` resource (`kind: HelmRelease`),
  either defined as a Go-templated pseudo-YAML file in the shared `cluster`
  chart's `files/helmreleases/`, or as an actual Helm template under a
  provider chart's `templates/`.
- **Release CR**: the `kind: Release` resource in this repo's
  `release.yaml` files, listing the resolved apps for one provider release.

## Step-by-step process

### 1. Read the provider chart's `values.yaml`

File: `cluster-<provider>/helm/cluster-<provider>/values.yaml`

This is plain YAML (no Go templating). Extract, under the
**`cluster.providerIntegration`** key (not `global.providerIntegration` —
the shared `cluster` chart is included as a Helm subchart, so its own
template values, like `$.Values.providerIntegration.provider` seen inside
the shared chart's HelmRelease files, are the parent chart's `cluster:` key
remapped to the subchart's root):

- **`provider`** (string, e.g. `aws`, `azure`, `eks`, `vsphere`,
  `cloud-director`, `aks`) — used later to resolve provider-guarded
  `dependsOn` blocks.
- **`apps`**: an object keyed by `configKey`. For every `configKey` whose
  value has `enable: true`, that app is expected to be rendered.

Result of this step: `provider` string + set of enabled `configKey`s.

### 2. Resolve each enabled `configKey` via the shared `cluster` chart

Directory: `cluster/helm/cluster/files/helmreleases/*.yaml`

Each file is a **Go-templated pseudo-YAML** HelmRelease definition (not
valid YAML on its own — do not try to fully parse it as YAML). For each
file, extract:

- **`appName`** (top-level key, plain string).
- **`configKey`** (top-level key, optional). If absent, derive it from
  `appName` by converting kebab-case to lower camelCase (e.g.
  `cert-exporter` → `certExporter`). Note this derivation is **not always
  correct** for names with no dashes — e.g. `appName: coredns` would
  naively camel-case to `coredns`, but the actual configKey is `coreDns`;
  that's exactly why some files carry an explicit `configKey:` override.
  Always prefer the explicit field when present.
- **`dependsOn`** (top-level key, optional): a plain YAML list of app
  names (strings, not objects).

  Take care: a `dependsOn:` block may be wrapped in a Go template `if`
  condition, e.g.:

  ```
  {{- if ne $.Values.providerIntegration.provider "aks" }}
  dependsOn:
  - cilium
  {{- end }}
  ```

  We only care about `if` conditions that guard the **entire `dependsOn`
  block** based on `$.Values.providerIntegration.provider` (via `eq` or
  `ne`). Evaluate that condition using the `provider` string from step 1 to
  decide whether the block applies at all. **Any other `if` condition found
  elsewhere in the file** (around `defaultValues`, proxy settings, version
  checks, etc.) is irrelevant here and must be ignored.

For every enabled `configKey` from step 1, look up its matching file here
(by explicit or derived `configKey`) and record `(appName, dependsOn set)`.

### 3. Find provider-specific HelmRelease CRs

Directory: `cluster-<provider>/helm/cluster-<provider>/templates/`
(recursively — the exact subdirectory layout varies per provider chart:
some nest HelmReleases under `templates/helmreleases/`, others place them
directly in `templates/`; this doesn't matter since we scan recursively).

**Do not identify HelmRelease CRs by filename.** Filenames are inconsistent
(sometimes suffixed `-helmrelease.yaml`, sometimes not) and a single file
can define **more than one** HelmRelease CR (e.g. `cluster-aws`'s
`karpenter.yaml` defines both `karpenter` and `karpenter-taint-remover`).
Also, some files that *do* have "helmrelease" in their filename (e.g.
`_cilium_helmrelease_config.yaml`) are not HelmRelease CRs at all — they're
Go template partials (`{{- define "..." }}`) supplying values referenced via
a `configTemplateName` in `values.yaml`, and must be skipped.

Correct approach: read every `.yaml` file under `templates/` and split it
into `---`-separated YAML documents. For each document that literally
contains `kind: HelmRelease`, it's a real HelmRelease CR. To get its
`appName`, scan backwards (or track state while reading top-to-bottom
through the whole file) for the most recent line matching:

```
{{- $_ := set $ "appName" "<app-name>" }}
```

(this is typically set in the preceding `OCIRepository` document within the
same file — the Go template's root context `$` is shared across all
documents rendered from one file, so the last value set before the
`HelmRelease` document is the correct one).

`dependsOn` here has a **different shape** than in step 2: it's a list of
objects with a `name` field, e.g.:

```yaml
dependsOn:
- name: {{ include "resource.default.name" . }}-cilium
  namespace: {{ .Release.Namespace }}
- name: {{ include "resource.default.name" . }}-vertical-pod-autoscaler-crd
  namespace: {{ .Release.Namespace }}
```

Each `name` is the rendered cluster-name prefix (the templated
`{{ include "resource.default.name" . }}-` part) followed by a hard-coded
suffix — only that suffix is the dependency's `appName` (e.g. `cilium`,
`vertical-pod-autoscaler-crd`). Watch out for multi-line list items (a
`namespace:` line following each `- name:` line, at greater indentation) —
a naive line-by-line match on consecutive `- name:` lines only will
under-count entries after the first.

### 4. Build one "expected apps" list

Combine the results of steps 2 and 3 into a single list of entries, each
`(source, appName, dependsOn set)` — do **not** collapse them into a map
keyed by `appName` yet. The same `appName` can legitimately appear twice
here: once from a shared-chart file (step 2) and once from a
provider-template HelmRelease (step 3). This was observed for
`cluster-azure`: the shared chart's `external-dns.yaml` and the provider
template's `external-dns-private-helmrelease.yaml` both define
`appName: external-dns`.

When that happens, both entries are kept and checked — neither one is
authoritative over the other. Instead:

- The two entries' `dependsOn` sets are expected to be **identical**. They
  describe the same app; the provider-specific HelmRelease is expected to
  declare the same dependencies as the shared-chart default, just rendered
  through a different template. A difference between them is itself a real
  inconsistency worth flagging — between the shared cluster chart and the
  provider cluster chart, independent of anything in `release.yaml`.
- The `release.yaml` Release CR, on the other hand, should list that
  `appName` exactly **once** (Release CRs don't have a "source" concept —
  there's only one app either way), and its `dependsOn` should match what
  both of the above entries agree on.

### 5. Read the Release CR

File: `releases/<provider-dir>/<latest-version>/release.yaml`

"Latest version" = the highest-numbered `vX.Y.Z` directory under
`releases/<provider-dir>/` (this is the release currently in progress).

Read `spec.apps` (a plain YAML list — **ignore `spec.components`**, that's
out of scope here). Each entry has:

- `name` — this is the `appName`.
- `dependsOn` (optional) — a plain list of app name strings (unlike the
  provider-template shape in step 3, these are already plain strings, not
  objects with a `name` field).

Build a map `appName → dependsOn set` from this.

### 6. Compare

Compare the "expected apps" list (step 4) against the Release CR's apps map
(step 5):

- First, for every `appName` that appears more than once in the expected
  apps list (i.e. defined by both the shared chart and a provider
  template), check that all of its entries agree on the same `dependsOn`
  set. Flag any disagreement here as an inconsistency between the shared
  cluster chart and the provider cluster chart — independent of
  `release.yaml`.
- Reduce the expected apps list to a map of `appName → dependsOn set` (safe
  to do once the check above passed, since all entries for a given
  `appName` agree).
- Every expected `appName` should be present in the Release CR (flag any
  that are missing).
- The Release CR should not contain `appName`s that aren't expected (flag
  any extras), and should not list the same `appName` more than once.
- For apps present in both, their `dependsOn` sets should match exactly
  (flag any mismatch, in either direction).

A clean result is an exact match: same set of app names, same `dependsOn`
per app on both sides (shared chart vs. provider template, and expected vs.
`release.yaml`), nothing more and nothing less.

## Known accepted exceptions

- **`vsphere` / `kamaji-etcd`**: `cluster-vsphere`'s `values.yaml` has
  `kamajiEtcd.enable: true`, which resolves (via the shared chart's
  `kamaji-etcd.yaml`) to an expected app named `kamaji-etcd`. As of this
  writing it is **not yet** present in `releases/vsphere`'s latest
  `release.yaml`. This is a known, accepted gap — not a bug.

Any other missing app, extra app, or `dependsOn` mismatch found by this
process should be treated as a real discrepancy between the provider
chart's current state and its Release CR, worth investigating before the
release is cut — unless it's already known to be an intentional
work-in-progress state (e.g. a release prepared ahead of a chart migration
that was later postponed, as happened with `eks` at one point).
