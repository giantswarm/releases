# :zap: Giant Swarm Release v35.1.0 for  :zap:

## Changes compared to v35.0.0

### Components

- cluster-aks from v0.3.0 to v0.5.0
- Kubernetes from v1.35.5 to [v1.35.7](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.7)

### cluster-aks [v0.3.0...v0.5.0](https://github.com/giantswarm/cluster-aks/compare/v0.3.0...v0.5.0)

#### Added

- Add default tag `giantswarm-cluster` to all resources.
- Allow adding custom tags to resources using `providerSpecific.additionalResourceTags` value.

### Apps

- cert-exporter from v2.11.0 to v2.12.0
- cert-manager from v3.13.0 to v4.1.1
- Added cert-manager-crossplane-resources v0.2.0
- cilium-servicemonitors from v0.1.4 to v0.2.0
- external-dns-crossplane-resources from v0.3.0 to v0.4.0
- observability-bundle from v2.9.1 to v3.3.0
- prometheus-blackbox-exporter from v0.8.0 to v0.9.0
- security-bundle from v1.17.1 to v2.2.0

### cert-exporter [v2.11.0...v2.12.0](https://github.com/giantswarm/cert-exporter/compare/v2.11.0...v2.12.0)

#### Added

- Regression tests covering the v2.11.1 metrics endpoint fix: concatenated and repeated certificates in secrets and files, and an endpoint level test asserting `/metrics` keeps returning 200 and serving the remaining metrics when a duplicate series is emitted.
- ATS: End to end test for a TLS secret whose `tls.crt` holds a concatenated certificate chain.

#### Changed

- ATS: `cert_gen` now sets a certificate serial number, which defaulted to `0` for every generated certificate.
- Go: Update dependencies.
- Templates: Move PolicyException to `kube-system` namespace.

#### Fixed

- Add a `serialnumber` label to the `cert_exporter_not_after` and `cert_exporter_secret_not_after` metrics so concatenated certificates no longer collide into identical series. The collision made the registry fail `Gather()`, which blanked out the entire `/metrics` endpoint (regression from v2.10.1).
- Serve `/metrics` with `ContinueOnError` so a single problematic metric can no longer fail the whole scrape.

### cert-manager [v3.13.0...v4.1.1](https://github.com/giantswarm/cert-manager-app/compare/v3.13.0...v4.1.1)

#### Changed

- Upgrade `docker-kubectl` image to support arm64 architecture
- Updated `cert-manager` to upstream version `v1.20.3`.
- **Notes:** `cert-manager-edit` ClusterRole no longer grants `create` on `challenges.acme.cert-manager.io`, nor `create`/`patch`/`update` on `orders.acme.cert-manager.io`.
- Improved proxy settings by adding a proxy ConfigMap and setting upstream `envFrom` values for `controller`, `webhook` and `cainjector`.
- **Breaking:** Helm values to be passed to the upstream `cert-manager` chart will now need to use the `cert-manager` path instead of root. For example, the value `crds.enabled: true` must now be set with `cert-manager.crds.enabled: true`.
- Moved vendored chart to `helm/cert-manager/charts/` and adapted sync scripts to follow new structure.

### cert-manager-crossplane-resources [v0.2.0](https://github.com/giantswarm/cert-manager-crossplane-resources/releases/tag/v0.2.0)

#### Added

- Add `WorkloadIdentity` resource for AKS provider.

### cilium-servicemonitors [v0.1.4...v0.2.0](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.4...v0.2.0)

#### Added

- Add per-monitor `enabled` flag for the agent, hubble and operator.

#### Changed

- Switch all monitors from ServiceMonitor to PodMonitor.

### external-dns-crossplane-resources [v0.3.0...v0.4.0](https://github.com/giantswarm/external-dns-crossplane-resources/compare/v0.3.0...v0.4.0)

#### Added

- Add configurable `serviceAccount.name`/`serviceAccount.namespace` and `azure.mountPath` values.

#### Changed

- Scope AWS values under `aws.*`. Root-level values remain supported as a fallback for backward compatibility.
- Templatize hardcoded resource, service account and tag values.
- Scope the IAM trust policy to the exact `serviceAccount` subject instead of a wildcard match.
- Enable Azure Workload Identity for the `azure` provider in addition to `aks`.
- Update chart metadata to reflect Azure support.

#### Removed

- Remove unused `baseDomain` value.

### observability-bundle [v2.9.1...v3.3.0](https://github.com/giantswarm/observability-bundle/compare/v2.9.1...v3.3.0)

#### Added

- Add KSM metrics for Gateway API `ListenerSet` and `ReferenceGrant` resources.

#### Changed

- Update Gateway API KSM configs to `v1` for `Gateway`, `GatewayClass`, `HTTPRoute`, `GRPCRoute`, `TLSRoute` and `BackendTLSPolicy`.
- Update `kube-prometheus-stack` and `prometheus-operator-crd` to 22.0.0
- Update `alloy-app` to 0.21.0
- HelmReleases: honor the App platform `priority` field (1-150, default 25) on `extraConfigs` entries. `spec.valuesFrom` now reproduces the App platform merge order — all configMaps before all secrets (a secret always overrides a configMap), each kind ordered by priority around the user-config layer — preserving the App CR merge semantics after the migration. ([giantswarm#36096](https://github.com/giantswarm/giantswarm/issues/36096))
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- Remove 'cluster-values' ConfigMap reference from HelmReleases.
- Add new `alloy-podlogs-crds` chart.
- Update alloy-app to 0.20.0

### prometheus-blackbox-exporter [v0.8.0...v0.9.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.8.0...v0.9.0)

#### Added

- Add VPA for `blackbox-exporter`. Uses `updateMode: Initial` for DaemonSet and `updateMode: Auto` for Deployment.

#### Fixed

- Add `probe_target` label to ensure unique synthetic metrics

### security-bundle [v1.17.1...v2.2.0](https://github.com/giantswarm/security-bundle/compare/v1.17.1...v2.2.0)

#### Changed

- Update `kyverno-policies` (app) to v0.26.1.
- Update `starboard-exporter` (app) to v1.1.4.
- Update `trivy` (app) to v0.16.0.
- Update `trivy-operator` (app) to v0.13.2.
- Update `cloudnative-pg` (app) to v0.1.0.
- Update `trivy` (app) to v0.15.0.
- Update `falco` (app) to v0.12.0.
- HelmReleases: honor the App platform `priority` field (1-150, default 25) on `extraConfigs` entries. `spec.valuesFrom` now reproduces the App platform merge order — all configMaps before all secrets (a secret always overrides a configMap), each kind ordered by priority around the user-config layer — preserving the App CR merge semantics after the migration. ([giantswarm#36096](https://github.com/giantswarm/giantswarm/issues/36096))
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- No longer pass the 'cluster-values' ConfigMap to the applications inside the bundle.
- Update `kyverno` (app) to v0.24.2.
  - This release includes a new Kyverno minor version. Please refer to the upstream release notes for the latest changes:
  - https://github.com/kyverno/kyverno/releases/tag/v1.17.0
- Update `kyverno-crds` (app) to v1.17.0.
- Update `kyverno-policies` (app) to v0.25.0.
- Update `kyverno-policy-operator` (app) to v0.2.2.
- Update `kubescape` (app) to v0.1.0.
