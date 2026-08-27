# :zap: Giant Swarm Release v34.5.1 for Azure :zap:

## Changes compared to v34.5.0

### Components

- cluster-azure from v5.4.2 to v5.4.4
- Added containerd [v2.3.2](https://github.com/containerd/containerd/releases/tag/v2.3.2)

### cluster-azure [v5.4.2...v5.4.4](https://github.com/giantswarm/cluster-azure/compare/v5.4.2...v5.4.4)

#### Added

- Add option to be able to disable the Private Link on private clusters (`.global.connectivity.network.enablePrivateLinkWithPrivateMode`). Useful for BYON scenarios.
- Add initial Bring-Your-Own-Network support.
- Add `global.providerSpecific.failureDomains` to restrict the availability zones of the region the control plane nodes are allowed to run in.

