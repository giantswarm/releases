# :zap: Giant Swarm Release v30.1.2 for CAPA :zap:

This release reconfigures the Cilium HelmRelease to avoid disruptions during cluster upgrades.

## Changes compared to v30.1.0

### Apps

- cluster-aws from v3.2.1 to v3.2.2
- cilium from v0.31.1 to v0.31.3

### cluster-aws [v3.2.1...v3.2.2](https://github.com/giantswarm/cluster-aws/compare/v3.2.1...v3.2.2)

#### Changed

- Chart: Update cluster to v2.2.1.

### cluster [v2.2.0...v2.2.1](https://github.com/giantswarm/cluster/compare/v2.2.0...v2.2.1)

#### Changed

- Make HelmRelease options configurable per app.
- Set Cilium HelmRelease timeout to 1hs and disable remediation.

### cilium [v0.31.1...v0.31.3](https://github.com/giantswarm/cilium-app/compare/v0.31.1...v0.31.3)

#### Changed

- Upgrade Cilium to v1.16.9.

