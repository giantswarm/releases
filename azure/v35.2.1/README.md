# :zap: Giant Swarm Release v35.2.1 for Azure :zap:

## Changes compared to v35.2.0


### Apps

- cluster-autoscaler from v2.0.4 to v2.1.1

### cluster-autoscaler [v2.0.4...v2.1.1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v2.0.4...v2.1.1)

#### Changed

- Chart: Update to upstream v1.36.1.

#### Fixed

- RBAC: Allow reading `infrastructure.cluster.x-k8s.io` resources when `clusterAPI.enabled` is set.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
