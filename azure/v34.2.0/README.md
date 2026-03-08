# :zap: Giant Swarm Release v34.2.0 for Azure :zap:

<< Add description here >>

## Changes compared to v34.1.0

### Components

- cluster-azure from v5.3.0 to v5.4.1
- cluster from v5.1.2 to v5.3.1

### cluster-azure [v5.3.0...v5.4.1](https://github.com/giantswarm/cluster-azure/compare/v5.3.0...v5.4.1)

#### Changed

- Apps: Enable `rbac-bootstrap` as a default HelmRelease app.

### cluster [v5.1.2...v5.3.1](https://github.com/giantswarm/cluster/compare/v5.1.2...v5.3.1)

#### Added

- Apps: Add `rbac-bootstrap` as a default HelmRelease app with a default ClusterRoleBinding for `giantswarm:giantswarm-admins`.

#### Changed

- Apps: Use OCIRepository source for `rbac-bootstrap` HelmRelease.

#### Fixed

- Apps: Change `rbac-bootstrap` default role from `read-all` to `view` and add additional groups for token forwarded cases.

### Apps

- observability-bundle from v2.6.0 to v2.8.0

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
