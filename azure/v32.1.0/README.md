# :zap: Giant Swarm Release v32.1.0 for Azure :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- cluster-azure from v3.0.0 to v4.0.0
- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### cluster-azure [v3.0.0...v4.0.0](https://github.com/giantswarm/cluster-azure/compare/v3.0.0...v4.0.0)

#### Changed

- Chart: Update `cluster` to v4.0.1.

### Apps

- azure-cloud-controller-manager from v1.32.7-1 to v1.32.7
- cilium from v1.3.0 to v1.3.1

### azure-cloud-controller-manager [v1.32.7-1...v1.32.7](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v1.32.7-1...v1.32.7)

#### Changed

- Switch to semver-compatible release name
- Chart: Update to upstream v1.32.7. ([#114](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/114))

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).
