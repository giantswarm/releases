# :zap: Giant Swarm Release v36.0.0 for EKS :zap:

## Changes compared to v35.0.0

### Components

- cluster-eks from v3.0.1-dev.migrate-app--elmreleases.2026-08-19.15-46-04.h1888387 to v3.0.0
- Kubernetes from v1.35.7 to [v1.36.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.3)

### Apps

- Added priority-classes v0.3.1

### priority-classes [v0.3.1](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.1)

#### Fixed

- Sanitize `Chart.Version` used in labels. This is needed because flux apapends the digest to the version using the `+` character which is not allowed in labels.
