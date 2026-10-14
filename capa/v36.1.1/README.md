# :zap: Giant Swarm Release v36.1.1 for CAPA :zap:

## Changes compared to v36.1.0

### Components

- cluster-aws from v11.0.1 to v11.0.2

### cluster-aws [v11.0.1...v11.0.2](https://github.com/giantswarm/cluster-aws/compare/v11.0.1...v11.0.2)

#### Changed

- Bump `OCIRepository` resources from the `source.toolkit.fluxcd.io/v1beta2` API to `source.toolkit.fluxcd.io/v1`.

### Apps

- aws-pod-identity-webhook from v2.3.0 to v2.3.1
- cluster-autoscaler from v2.1.0 to v2.1.1

### aws-pod-identity-webhook [v2.3.0...v2.3.1](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v2.3.0...v2.3.1)

#### Added

- Build and publish `linux/arm64` images of `aws-pod-identity-restarter`.

#### Fixed

- Bump `golang.org/x/net` to v0.61.0.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### cluster-autoscaler [v2.1.0...v2.1.1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v2.1.0...v2.1.1)

#### Fixed

- RBAC: Allow reading `infrastructure.cluster.x-k8s.io` resources when `clusterAPI.enabled` is set.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
