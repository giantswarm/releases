# :zap: Giant Swarm Release v35.0.0 for Azure :zap:

## Changes compared to v34.4.0

### Components

- cluster-azure from v5.4.2 to v6.7.0
- cluster from v5.3.2 to v6.7.0
- Kubernetes from v1.34.7 to [v1.35.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.6)
- os-tooling from v1.31.0 to v1.32.1

### cluster-azure [v5.4.2...v6.7.0](https://github.com/giantswarm/cluster-azure/compare/v5.4.2...v6.7.0)

#### Added

- Node pools: Add `minSize` and `maxSize` fields to enable per-node-pool autoscaling configuration.
- Autoscaling: Deploy cluster-autoscaler to the management cluster org-namespace using in-cluster credentials to scale MachineDeployments.
- Add support for `network.giantswarm.io/wildcard-cname-target` annotation on `AzureCluster` via `global.connectivity.dns.wildcardCnameTarget`.

#### Changed

- Enable cert-manager DNS challenges by default.
- Make: Fix application variable.
- Move `network.giantswarm.io/wildcard-cname-target` annotation from `AzureCluster` CR to `Cluster` CR.
- Add proxy values to azuredisk-csi-driver and azurefile-csi-driver HelmRelease default values.
- Add `_clusterautoscaler_app_config.yaml` to provide Azure `subscriptionID` to cluster-autoscaler.
- Migrate `external-dns-private` from App CR to HelmRelease, removing cluster-values ConfigMap reference.
- Add proxy env vars and ciliumNetworkPolicy to external-dns-private default values.

#### Removed

- Templates: Remove dead container registry cache schema.
- Chart: Remove unused `cluster-shared` library chart dependency.

### cluster [v5.3.2...v6.7.0](https://github.com/giantswarm/cluster/compare/v5.3.2...v6.7.0)

#### Added

- Feature Gates: Add support for defining maximum Kubernetes version.
- Apps: Add External DNS Crossplane Resources.
- Apps: Deploy `cluster-autoscaler` inCluster in Azure.
- MachineDeployment: Add CAPI autoscaler annotations (`cluster-api-autoscaler-node-group-min-size`/`max-size`) when `minSize`/`maxSize` are set on a node pool (only in Azure).
- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))

#### Changed

- Add support for AKS clusters
- Add support for disabling the external autoscaler annotation (`cluster.x-k8s.io/replicas-managed-by: "external-autoscaler"`) on `MachinePools`.
- Refactor `providerIntegration.resourcesApi.machinePoolResourcesEnabled` into an object `providerIntegration.resourcesApi.machinePoolResources` and move `externalAutoscaler` under it (now `providerIntegration.resourcesApi.machinePoolResources.externalAutoscaler`).
- Bump Flux OCIRepository version to v1.
- Control Plane: Remove handling of clusterRole resources for kamaji Datastore CRs and create kamaji-etcd polex here.
- Control Plane: Make etcd image tag configurable. ([#841](https://github.com/giantswarm/cluster/pull/841))
- Chart: Require `global.release.version` if using Releases to give a better rendering error message.
- Chart: Fix validation errors.
- Configure `observability-bundle` with the management cluster name.
- Apps: Skip `kyverno-crds` dependency for `cluster-autoscaler` when deployed inCluster.
- Apps: Add cluster-probes HelmRelease to deploy ServiceMonitors for probing workload cluster API server endpoint from the management cluster. Configurable via `global.apps.clusterProbes` with default module `http_2xx_insecure` for self-signed certificates.
- Helpers: Use `.Chart.AppVersion` in `app.kubernetes.io/version` label.
- Cluster API: Migrate to API `v1beta2`.

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.

#### Fixed

- Control Plane: Ensure components start correctly when SELinux is set to enforcing.

### Apps

- cert-exporter from v2.10.1 to v2.11.1
- cilium from v1.4.3 to v1.4.5
- Added cluster-autoscaler v2.0.2
- coredns from v1.30.0 to v1.31.0
- etcd-defrag from v1.2.6 to v1.2.8
- external-dns from v3.4.0 to v3.5.0
- net-exporter from v1.23.1 to v1.24.0
- network-policies from v0.1.3 to v0.2.0
- observability-bundle from v2.8.0 to v2.9.1
- prometheus-blackbox-exporter from v0.7.0 to v0.8.0
- Added rbac-bootstrap v0.3.0
- teleport-kube-agent from v0.10.8 to v0.11.1

### cert-exporter [v2.10.1...v2.11.1](https://github.com/giantswarm/cert-exporter/compare/v2.10.1...v2.11.1)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the cert-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.

#### Fixed

- Add a `serialnumber` label to the `cert_exporter_not_after` and `cert_exporter_secret_not_after` metrics so concatenated certificates no longer collide into identical series. The collision made the registry fail `Gather()`, which blanked out the entire `/metrics` endpoint (regression from v2.10.1).
- Serve `/metrics` with `ContinueOnError` so a single problematic metric can no longer fail the whole scrape.

### cilium [v1.4.3...v1.4.5](https://github.com/giantswarm/cilium-app/compare/v1.4.3...v1.4.5)

#### Changed

- Upgrade Cilium to [v1.19.5](https://github.com/cilium/cilium/releases/tag/v1.19.5).
- Upgrade Cilium to [v1.19.4](https://github.com/cilium/cilium/releases/tag/v1.19.4).

### cluster-autoscaler [v2.0.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v2.0.2)

#### Changed

- Templates: Move PolicyException to `giantswarm` namespace.

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

### etcd-defrag [v1.2.6...v1.2.8](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.6...v1.2.8)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.41.0. ([#108](https://github.com/giantswarm/etcd-defrag-app/pull/108))
- Chart: Update dependency ahrtr/etcd-defrag to v0.40.0. ([#94](https://github.com/giantswarm/etcd-defrag-app/pull/94))

### external-dns [v3.4.0...v3.5.0](https://github.com/giantswarm/external-dns-app/compare/v3.4.0...v3.5.0)

#### Changed

- Update VPA `updatePolicy.updateMode` from deprecated `Auto` to `Recreate`.
- Upgrade external-dns to [v0.21.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.21.0).
- Sync to upstream helm chart [1.21.1](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.21.1).
  - Add `namespaceOverride` value to support deploying external-dns into a different namespace than the Helm release (useful for subchart usage).
  - Add `enableGatewayListenerSets` value to opt into Gateway API ListenerSet resource support.
  - Add `sourceNamespace` value (used with `namespaced=true`) to watch resources in a namespace different from the deployment namespace.
  - Avoid creating cluster-scoped namespace RBAC when `gatewayNamespace` is set, reducing required permissions.
  - Fix `extraArgs` map handling: boolean values now render as `--flag` / `--no-flag` and string values are properly quoted.
- Use external-dns.namespace in VPA and NetworkPolicy resources.

### net-exporter [v1.23.1...v1.24.0](https://github.com/giantswarm/net-exporter/compare/v1.23.1...v1.24.0)

#### Changed

- Build and publish a multi-arch (linux/amd64 + linux/arm64) container image. Required so the net-exporter daemonset can run on Graviton/arm64 nodepools without `exec format error`.
- Bump `docker-kubectl` init container from `1.25.4` to `1.36.0`.

### network-policies [v0.1.3...v0.2.0](https://github.com/giantswarm/network-policies-app/compare/v0.1.3...v0.2.0)

#### Added

- Add support for AKS selector labels.

#### Changed

- Deprecated the .Values.kamaji in favour of the more generic .Values.konnectivityAgent to control the behaviour for the `konnectivity-agent`.

### observability-bundle [v2.8.0...v2.9.1](https://github.com/giantswarm/observability-bundle/compare/v2.8.0...v2.9.1)

#### Added

- Add Backstage audience annotations.
- Add managementCluster: "" as a top-level value (populated from the cluster chart via defaultValues)
- Moves full KSM metricRelabelings ownership from kube-prometheus-stack-app into observability-bundle

#### Changed

- Update `alloy-app` to 0.20.1
- Update dependency kube-prometheus-stack-app and prometheus-operator-crd to v21.0.0
- Update alloy-app to 0.19.0

### prometheus-blackbox-exporter [v0.7.0...v0.8.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.7.0...v0.8.0)

#### Added

- Add toleration for `kubernetes.io/arch=arm64:NoSchedule` so the DaemonSet schedules on ARM worker nodes.

### rbac-bootstrap [v0.3.0](https://github.com/giantswarm/rbac-bootstrap-app/releases/tag/v0.3.0)

#### Added

- Add `io.giantswarm.application.managed` chart annotation for Backstage visibility.
- Add optional `cluster-reader` ClusterRole (off by default, enabled via `clusterReader.enabled: true`) that aggregates into the built-in `view` ClusterRole and grants read access (`get`/`list`/`watch`) on cluster-scoped resources.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.

### teleport-kube-agent [v0.10.8...v0.11.1](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.8...v0.11.1)

#### Changed

- Values: Tolerate `node.cloudprovider.kubernetes.io/uninitialized`.
- Values: Ignore taints regardless of value.
- Values: Pass HTTP proxy settings to sub-chart.
- Updated `teleport-kube-agent` to upstream version `v18.7.6`.
