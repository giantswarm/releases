# :zap: Giant Swarm Release v29.0.1 for vSphere :zap:

## Changes compared to v29.0.0

### cluster-vsphere [v0.65.2...v0.66.0](https://github.com/giantswarm/cluster-vsphere/compare/v0.65.2...v0.66.0)

### Components

- cluster-vsphere from v0.65.2 to v0.66.0

### cluster-vsphere [v0.65.1...v0.65.2](https://github.com/giantswarm/cluster-vsphere/compare/v0.65.1...v0.65.2)

#### Changed

- Use Renovate to update `kube-vip` static pod manifest.
- Updated `giantswarm/cluster` to `v1.6.0`.
- Update `kubectl` image used by IPAM job to `1.29.9`.
- Use init-container to prepare `/etc/hosts` file for `kube-vip`.
