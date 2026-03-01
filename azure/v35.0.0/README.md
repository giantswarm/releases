# :zap: Giant Swarm Release v35.0.0 for Azure :zap:

<< Add description here >>

## Changes compared to v34.1.0

### Components

- cluster-azure from v5.3.0 to v6.2.0
- cluster from v5.1.2 to v6.2.0
- Flatcar from v4459.2.3 to [v4459.2.4](https://www.flatcar.org/releases/#release-4459.2.4)
- Kubernetes from v1.34.5 to [v1.35.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.2)

### cluster-azure [v5.3.0...v6.2.0](https://github.com/giantswarm/cluster-azure/compare/v5.3.0...v6.2.0)

#### Added

- Node pools: Add `minSize` and `maxSize` fields to enable per-node-pool autoscaling configuration.
- Autoscaling: Deploy cluster-autoscaler to the management cluster org-namespace using in-cluster credentials to scale MachineDeployments.
- Add support for `network.giantswarm.io/wildcard-cname-target` annotation on `AzureCluster` via `global.connectivity.dns.wildcardCnameTarget`.

#### Changed

- Add proxy values to azuredisk-csi-driver and azurefile-csi-driver HelmRelease default values.
- Add `_clusterautoscaler_app_config.yaml` to provide Azure `subscriptionID` to cluster-autoscaler.
- Migrate `external-dns-private` from App CR to HelmRelease, removing cluster-values ConfigMap reference.
- Add proxy env vars and ciliumNetworkPolicy to external-dns-private default values.
- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

### cluster [v5.1.2...v6.2.0](https://github.com/giantswarm/cluster/compare/v5.1.2...v6.2.0)

#### Added

- Apps: Deploy `cluster-autoscaler` inCluster in Azure.
- MachineDeployment: Add CAPI autoscaler annotations (`cluster-api-autoscaler-node-group-min-size`/`max-size`) when `minSize`/`maxSize` are set on a node pool (only in Azure).
- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))
- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.

#### Changed

- Apps: Skip `kyverno-crds` dependency for `cluster-autoscaler` when deployed inCluster.
- Apps: Add cluster-probes HelmRelease to deploy ServiceMonitors for probing workload cluster API server endpoint from the management cluster. Configurable via `global.apps.clusterProbes` with default module `http_2xx_insecure` for self-signed certificates.
- Helpers: Use `.Chart.AppVersion` in `app.kubernetes.io/version` label.
- Cluster API: Migrate to API `v1beta2`.
- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.

#### Fixed

- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.

### Apps

- azure-cloud-controller-manager from v2.0.0 to v2.1.0
- azure-cloud-node-manager from v2.0.0 to v2.1.0
- cert-exporter from v2.9.16 to v2.10.0
- Added cluster-autoscaler v1.34.3-2
- observability-bundle from v2.6.0 to v2.8.0
- prometheus-blackbox-exporter from v0.5.1 to v0.7.0

### azure-cloud-controller-manager [v2.0.0...v2.1.0](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v2.0.0...v2.1.0)

#### Changed

- Migrate to App Build Suite (ABS).
- Bump to upstream image v1.35.1

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### azure-cloud-node-manager [v2.0.0...v2.1.0](https://github.com/giantswarm/azure-cloud-node-manager-app/compare/v2.0.0...v2.1.0)

#### Changed

- Migrate to App Build Suite (ABS).
- Bump to upstream image v1.35.1

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### cert-exporter [v2.9.16...v2.10.0](https://github.com/giantswarm/cert-exporter/compare/v2.9.16...v2.10.0)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Values: Tune resources.

### cluster-autoscaler [v1.34.3-2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.34.3-2)

#### Added

- Validate that `managementCluster` (when `isManagementCluster=true`) or `clusterID` (otherwise) are set, failing early with a clear error message.
- Add support for CAPI mode (`kubeconfig-incluster`): run cluster-autoscaler on the management cluster using a pre-existing kubeconfig to connect to the workload cluster.
- Add `clusterAPI` values section for configuring CAPI mode (autodiscovery, kubeconfig secret, configmaps namespace).
- Add `rbac.clusterScoped` toggle to support namespace-scoped RBAC (no ClusterRole/ClusterRoleBinding) for CAPI deployments.

#### Changed

- Migrate test infrastructure from pipenv to uv.
- Deploy the Kyverno policy exception in the `policy-exceptions` Namespace.
- Deploy the Kyverno PolicyException as a Helm `pre-install,pre-upgrade` hook so it takes effect before chart resources are created.

### observability-bundle [v2.6.0...v2.8.0](https://github.com/giantswarm/observability-bundle/compare/v2.6.0...v2.8.0)

#### Added

- Add KSM metrics for Envoy Gateway resources.
- Add `application.giantswarm.io/team` annotation from HelmReleases as label to KSM emitted metrics.

#### Changed

- Update kube-prometheus-stack to 20.1.0
- Change team annotation in `Chart.yaml` to OpenContainers format (`io.giantswarm.application.team`).
- Update alloy-app to 0.17.1
- Update kube-prometheus-stack to 20.0.0
- Update prometheus-operator-crd to 20.0.0

### prometheus-blackbox-exporter [v0.5.1...v0.7.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.5.1...v0.7.0)

#### Added

- Add `http_2xx_insecure` module with `insecure_skip_verify: true` to support probing workload cluster API servers from the management cluster. The MC's service account CA (`http_2xx_k8sca`) only covers the MC itself; workload clusters have their own CA which is not available to the blackbox exporter, making TLS verification impossible without this module.

#### Changed

- Set `priorityClassName` to `system-node-critical` to ensure DaemonSet pods are scheduled even on full nodes.
