# :zap: Giant Swarm Release v36.1.1 for Azure :zap:

## Changes compared to v36.1.0

### Components

- cluster-azure from v10.1.0 to v10.1.1

### cluster-azure [v10.1.0...v10.1.1](https://github.com/giantswarm/cluster-azure/compare/v10.1.0...v10.1.1)

#### Changed

- Bump `OCIRepository` resources from the `source.toolkit.fluxcd.io/v1beta2` API to `source.toolkit.fluxcd.io/v1`.

### Apps

- cluster-autoscaler from v2.1.0 to v2.1.1

### cluster-autoscaler [v2.1.0...v2.1.1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v2.1.0...v2.1.1)

#### Fixed

- RBAC: Allow reading `infrastructure.cluster.x-k8s.io` resources when `clusterAPI.enabled` is set.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
