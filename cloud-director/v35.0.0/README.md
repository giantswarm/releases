# :zap: Giant Swarm Release v35.0.0 for VMware Cloud Director :zap:

<< Add description here >>

## Changes compared to v34.1.0

### Components

- cluster-cloud-director from v3.1.3 to v4.0.0
- cluster from v5.1.2 to v6.0.0
- Flatcar from v4459.2.3 to [v4459.2.4](https://www.flatcar.org/releases/#release-4459.2.4)
- Kubernetes from v1.34.5 to [v1.35.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.2)

### cluster-cloud-director [v3.1.3...v4.0.0](https://github.com/giantswarm/cluster-cloud-director/compare/v3.1.3...v4.0.0)

#### Changed

- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

### cluster [v5.1.2...v6.0.0](https://github.com/giantswarm/cluster/compare/v5.1.2...v6.0.0)

#### Added

- Control Plane: Add Kamaji control plane support with `KamajiControlPlane` resource, Kamaji etcd HelmRelease, automation RBAC, and cleanup jobs. ([#740](https://github.com/giantswarm/cluster/pull/740))
- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.

#### Changed

- Cluster API: Migrate to API `v1beta2`.
- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.

#### Fixed

- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.

#### Removed

- Cluster API: Remove `strategy.rollingUpdate.deletePolicy` from node pools.

### Apps

- cert-exporter from v2.9.16 to v2.10.0
- observability-bundle from v2.6.0 to v2.8.0

### cert-exporter [v2.9.16...v2.10.0](https://github.com/giantswarm/cert-exporter/compare/v2.9.16...v2.10.0)

#### Added

- DaemonSet: Add VPA.

#### Changed

- Values: Tune resources.

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
