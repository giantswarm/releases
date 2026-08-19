# :zap: Giant Swarm Release v36.0.0 for  :zap:

## Changes compared to v35.1.0

### Components

- Kubernetes from v1.35.7 to [v1.36.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.3)

### Apps

- node-exporter from v1.20.11 to v1.20.13
- observability-bundle from v3.3.0 to v3.3.1

### node-exporter [v1.20.11...v1.20.13](https://github.com/giantswarm/node-exporter-app/compare/v1.20.11...v1.20.13)

#### Changed

- CircleCI: Do not override app version.
- Chart: Move PolicyException to `kube-system` namespace.

### observability-bundle [v3.3.0...v3.3.1](https://github.com/giantswarm/observability-bundle/compare/v3.3.0...v3.3.1)

#### Changed

- Values: Generate schema for Alloy PodLogs CRDs.
- Values: Add Cilium as dependency for Alloy apps & Kube Prometheus Stack.
- Values: Update Alloy apps to v0.21.2.

#### Removed

- Values: Remove unused catalog.
