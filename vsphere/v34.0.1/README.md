# :zap: Giant Swarm Release v34.0.1 for vSphere :zap:

<< Add description here >>

## Changes compared to v34.0.0

### Components

- cluster-vsphere from v4.1.2 to v4.1.5

### cluster-vsphere [v4.1.2...v4.1.5](https://github.com/giantswarm/cluster-vsphere/compare/v4.1.2...v4.1.5)

#### Changed

- HelmReleases: Reduce hard-coded default interval from 10m to 5m.
- Values: Enable management cluster registry cache for `gsoci.azurecr.io`.

#### Fixed

- Allow adding additional properties into global.metadata.

