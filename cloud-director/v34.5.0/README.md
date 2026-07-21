# :zap: Giant Swarm Release v34.5.0 for VMware Cloud Director :zap:

## Changes compared to v34.4.0

### Components

- Flatcar from v4593.2.2 to [v4593.2.4](https://www.flatcar.org/releases/#release-4593.2.4)
- Kubernetes from v1.34.7 to [v1.34.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.9)
- os-tooling from v1.31.0 to v1.33.1

### Apps

- cert-exporter from v2.10.1 to v2.11.1
- cert-manager from v3.13.0 to v4.1.0
- cilium from v1.4.3 to v1.4.5
- cilium-servicemonitors from v0.1.4 to v0.2.0
- coredns from v1.30.0 to v1.31.0
- etcd-defrag from v1.2.6 to v1.2.9
- net-exporter from v1.23.1 to v1.24.0
- network-policies from v0.1.3 to v0.2.0
- observability-bundle from v2.8.0 to v3.3.0
- teleport-kube-agent from v0.10.8 to v0.11.1

### cert-exporter [v2.10.1...v2.11.1](https://github.com/giantswarm/cert-exporter/compare/v2.10.1...v2.11.1)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.

#### Fixed

- Add a `serialnumber` label to the `cert_exporter_not_after` and `cert_exporter_secret_not_after` metrics so concatenated certificates no longer collide into identical series. The collision made the registry fail `Gather()`, which blanked out the entire `/metrics` endpoint (regression from v2.10.1).
- Serve `/metrics` with `ContinueOnError` so a single problematic metric can no longer fail the whole scrape.

### cert-manager [v3.13.0...v4.1.0](https://github.com/giantswarm/cert-manager-app/compare/v3.13.0...v4.1.0)

#### Changed

- Updated `cert-manager` to upstream version `v1.20.3`.
- **Notes:** `cert-manager-edit` ClusterRole no longer grants `create` on `challenges.acme.cert-manager.io`, nor `create`/`patch`/`update` on `orders.acme.cert-manager.io`.
- Improved proxy settings by adding a proxy ConfigMap and setting upstream `envFrom` values for `controller`, `webhook` and `cainjector`.
- **Breaking:** Helm values to be passed to the upstream `cert-manager` chart will now need to use the `cert-manager` path instead of root. For example, the value `crds.enabled: true` must now be set with `cert-manager.crds.enabled: true`.
- Moved vendored chart to `helm/cert-manager/charts/` and adapted sync scripts to follow new structure.

### cilium [v1.4.3...v1.4.5](https://github.com/giantswarm/cilium-app/compare/v1.4.3...v1.4.5)

#### Changed

- Upgrade Cilium to [v1.19.5](https://github.com/cilium/cilium/releases/tag/v1.19.5).
- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).

### cilium-servicemonitors [v0.1.4...v0.2.0](https://github.com/giantswarm/cilium-servicemonitors-app/compare/v0.1.4...v0.2.0)

#### Added

- Add per-monitor `enabled` flag for the agent, hubble and operator.

#### Changed

- Switch all monitors from ServiceMonitor to PodMonitor.

### coredns [v1.30.0...v1.31.0](https://github.com/giantswarm/coredns-app/compare/v1.30.0...v1.31.0)

#### Added

- Wire up the full set of CoreDNS `forward`, `cache`, and `kubernetes` block parameters in the structured zone config:
  - `forward`: `maxIdleConns`, `maxConnectAttempts`, `dohMethod`, `tls`, `tlsServername`, `next`, `nextOnNodata`, `failfastAllUnhealthyUpstreams`, `failover`, `resolver`.
  - `cache`: `zones`, `serveStale.verifyTimeout`, `disable.successZones`, `disable.denialZones`.
  - `kubernetes`: `endpoint`, `tls`, `kubeconfig`, `apiserverQPS`, `apiserverBurst`, `apiserverMaxInflight`, `namespaceLabels`, `fallthroughZones`, `multicluster`, `startupTimeout`.

#### Changed

- Standardize `values.yaml` comments to the `# @schema` / `# --` (helm-docs) convention and remove section-header dividers, so `values.schema.json` and the chart `README.md` are generated from the values file. Rendered manifests are unchanged.
- Update `coredns` image to [1.14.4](https://github.com/coredns/coredns/releases/tag/v1.14.4).
- update CoreDNS icon to light version
- Update `coredns` image to [1.14.3](https://github.com/coredns/coredns/releases/tag/v1.14.3).

#### Fixed

- Render the `health` directive in only the `.` server block. The health plugin is process-wide and can be enabled in just one Server Block, so emitting it in every zone block was invalid. `ready` is kept in every block (its readiness is aggregated across blocks).
- Correct the `coredns.*.cache.serveStale.refreshMode` schema enum to `immediate`/`verify` (was `immediate`/`background`), matching the CoreDNS cache plugin.

### etcd-defrag [v1.2.6...v1.2.9](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.6...v1.2.9)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.42.0. ([#120](https://github.com/giantswarm/etcd-defrag-app/pull/120))
- Chart: Update dependency ahrtr/etcd-defrag to v0.41.0. ([#108](https://github.com/giantswarm/etcd-defrag-app/pull/108))
- Chart: Update dependency ahrtr/etcd-defrag to v0.40.0. ([#94](https://github.com/giantswarm/etcd-defrag-app/pull/94))

### net-exporter [v1.23.1...v1.24.0](https://github.com/giantswarm/net-exporter/compare/v1.23.1...v1.24.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the net-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Bump `docker-kubectl` init container from `1.25.4` to `1.36.0`.

### network-policies [v0.1.3...v0.2.0](https://github.com/giantswarm/network-policies-app/compare/v0.1.3...v0.2.0)

#### Added

- Add support for AKS selector labels.

#### Changed

- Deprecated the .Values.kamaji in favour of the more generic .Values.konnectivityAgent to control the behaviour for the `konnectivity-agent`.

### observability-bundle [v2.8.0...v3.3.0](https://github.com/giantswarm/observability-bundle/compare/v2.8.0...v3.3.0)

#### Added

- Add KSM metrics for Gateway API `ListenerSet` and `ReferenceGrant` resources.
- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle

#### Changed

- Update Gateway API KSM configs to `v1` for `Gateway`, `GatewayClass`, `HTTPRoute`, `GRPCRoute`, `TLSRoute` and `BackendTLSPolicy`.
- Update `kube-prometheus-stack` and `prometheus-operator-crd` to 22.0.0
- Update `alloy-app` to 0.21.0
- HelmReleases: honor the App platform `priority` field (1-150, default 25) on `extraConfigs` entries. `spec.valuesFrom` now reproduces the App platform merge order — all configMaps before all secrets (a secret always overrides a configMap), each kind ordered by priority around the user-config layer — preserving the App CR merge semantics after the migration. ([giantswarm#36096](https://github.com/giantswarm/giantswarm/issues/36096))
- Migrate sub-apps from App CRs to Flux HelmRelease CRs.
- Remove 'cluster-values' ConfigMap reference from HelmReleases.
- Add new `alloy-podlogs-crds` chart.
- Update alloy-app to 0.20.0
- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0

### teleport-kube-agent [v0.10.8...v0.11.1](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.1)

#### Changed

- Values: Tolerate `node.cloudprovider.kubernetes.io/uninitialized`.
- Values: Ignore taints regardless of value.
- Values: Pass HTTP proxy settings to sub-chart.
- Updated `teleport-kube-agent` to upstream version `v18.7.6`.
