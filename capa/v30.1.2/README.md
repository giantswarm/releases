# :zap: Giant Swarm Release v30.1.2 for CAPA :zap:

This release reconfigures the Cilium HelmRelease to avoid disruptions during cluster upgrades.

## Changes compared to v30.1.0

### Apps

- cluster-aws from v3.2.1 to v3.2.2

### cluster-aws [v3.2.1...v3.2.2](https://github.com/giantswarm/cluster-aws/compare/v3.2.1...v3.2.2)

#### Changed

- Chart: Update cluster to v2.2.1.

### cluster [v2.2.0...v2.2.1](https://github.com/giantswarm/cluster/compare/v2.2.0...v2.2.1)

#### Changed

- Make HelmRelease options configurable per app.
- Set Cilium HelmRelease timeout to 1hs and disable remediation.

