# :zap: Giant Swarm Release v35.0.0 for VMware Cloud Director :zap:

<< Add description here >>

## Changes compared to v34.1.1

### Components

- cluster-cloud-director from v3.1.4 to v4.4.0
- cluster from v5.1.2 to v6.4.0
- Flatcar from v4459.2.3 to [v4593.2.0](https://www.flatcar.org/releases/#release-4593.2.0)
- Kubernetes from v1.34.5 to [v1.35.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.4)
- os-tooling from v1.26.4 to v1.28.0

### cluster-cloud-director [v3.1.4...v4.4.0](https://github.com/giantswarm/cluster-cloud-director/compare/v3.1.4...v4.4.0)

#### Added

- Add support for `network.giantswarm.io/wildcard-cname-target` annotation on the `Cluster` CR via `global.connectivity.dns.wildcardCnameTarget`.

#### Changed

- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

#### Fixed

- Fix ntpd failing permanently on boot due to systemd rate limiting (**node rolling**).

#### Removed

- Chart: Remove unused `cluster-shared` library chart dependency.

### cluster [v5.1.2...v6.4.0](https://github.com/giantswarm/cluster/compare/v5.1.2...v6.4.0)

#### Added

- Apps: Add External DNS Crossplane Resources.
- Apps: Deploy `cluster-autoscaler` inCluster in Azure.
- MachineDeployment: Add CAPI autoscaler annotations (`cluster-api-autoscaler-node-group-min-size`/`max-size`) when `minSize`/`maxSize` are set on a node pool (only in Azure).
- Apps: Add Cluster Autoscaler Crossplane Resources.
- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))
- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.

#### Changed

- Configure `observability-bundle` with the management cluster name.
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

- cert-exporter from v2.9.16 to v2.10.1
- cert-manager from v3.11.0 to v3.13.0
- coredns from v1.29.1 to v1.30.0
- etcd-defrag from v1.2.4 to v1.2.6
- k8s-dns-node-cache from v2.9.2 to v2.11.0
- observability-bundle from v2.6.0 to v2.8.0
- security-bundle from v1.17.0 to v1.17.1

### cert-exporter [v2.9.16...v2.10.1](https://github.com/giantswarm/cert-exporter/compare/v2.9.16...v2.10.1)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Values: Tune resources.

#### Fixed

- Parse all PEM blocks in secrets and certificate files, not just the first one. This fixes false alerts when multiple certificates are concatenated (e.g. Kyverno webhook cert rotation).

### cert-manager [v3.11.0...v3.13.0](https://github.com/giantswarm/cert-manager-app/compare/v3.11.0...v3.13.0)

#### Added

- Add control plane node toleration to CA injector deployment.

#### Changed

- Upgrade cert-manager to v1.19.4.

#### Removed

- Remove PodSecurityPolicy (PSP) and related resources.
- Remove Giant Swarm PSP to PSS migration logic.

### coredns [v1.29.1...v1.30.0](https://github.com/giantswarm/coredns-app/compare/v1.29.1...v1.30.0)

#### Added

- Add `coredns-adopter` job to adopt default CoreDNS resources on EKS clusters (disabled by default).

#### Changed

- Update `coredns` image to [1.14.2](https://github.com/coredns/coredns/releases/tag/v1.14.2).

### etcd-defrag [v1.2.4...v1.2.6](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.4...v1.2.6)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.39.0. ([#86](https://github.com/giantswarm/etcd-defrag-app/pull/86))
- Chart: Update dependency ahrtr/etcd-defrag to v0.38.0. ([#84](https://github.com/giantswarm/etcd-defrag-app/pull/84))

### k8s-dns-node-cache [v2.9.2...v2.11.0](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.2...v2.11.0)

#### Added

- Add `configmap.log.enabled` helm value to toggle CoreDNS query logging (default: `false`).
- Make `AAAA NOERROR` configurable for IPv6.

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

### security-bundle [v1.17.0...v1.17.1](https://github.com/giantswarm/security-bundle/compare/v1.17.0...v1.17.1)

#### Added

- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.

#### Changed

- Update `falco` (app) to v0.11.2.
- Update `gel` (app) to v1.0.2.
- Update `kubescape` (app) to v0.0.6.
- Update `reports-server` (app) to v0.1.3.
- Update `starboard-exporter` (app) to v1.0.3.
- Update `trivy` (app) to v0.14.2.
- Update `trivy-operator` (app) to v0.12.2.
- Migrate chart annotations to OCI-compatible format.
