# :zap: Giant Swarm Release v30.1.2 for vSphere :zap:

This release reconfigures the Cilium HelmRelease to avoid disruptions during cluster upgrades.

## Changes compared to v30.1.1

### Components

- cluster-vsphere from v1.1.0 to v1.1.1

### cluster-vsphere [v1.1.0...v1.1.1](https://github.com/giantswarm/cluster-vsphere/compare/v1.1.0...v1.1.1)

#### Changed

- Chart: Update `cluster` to v2.2.1.

### Apps

- cilium from v0.31.1 to v0.31.4

### cilium [v0.31.1...v0.31.4](https://github.com/giantswarm/cilium-app/compare/v0.31.1...v0.31.4)

#### Changed

- Upgrade Cilium to [v1.16.10](https://github.com/cilium/cilium/releases/tag/v1.16.10).
