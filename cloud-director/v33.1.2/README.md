# :zap: Giant Swarm Release v33.1.2 for VMware Cloud Director :zap:

## Changes compared to v33.1.1

### Components

- cluster-cloud-director from v2.4.0 to v2.4.1

### cluster-cloud-director [v2.4.0...v2.4.1](https://github.com/giantswarm/cluster-cloud-director/compare/v2.4.0...v2.4.1)

#### Added

- Added `fix-dns-nic-allocation.sh` Ignition script to attach DNS servers to correct network interfaces.

#### Fixed

- Fix a race condition when populating `/run/metadata/coreos`.
- Fix race condition in `ntpd` unit.
