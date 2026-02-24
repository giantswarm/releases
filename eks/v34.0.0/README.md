# :zap: Giant Swarm Release v34.0.0 for EKS :zap:

<< Add description here >>

## Changes compared to v33.0.0

### Components

- Kubernetes from v1.33.8 to [v1.34.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.4)

### Apps

- Added priority-classes v0.3.0
- teleport-kube-agent from v0.10.7 to v0.10.8

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).

### teleport-kube-agent [v0.10.7...v0.10.8](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.7...v0.10.8)

#### Added

- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.
