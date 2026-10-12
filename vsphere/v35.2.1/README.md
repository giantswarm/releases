# :zap: Giant Swarm Release v35.2.1 for vSphere :zap:

## Changes compared to v35.2.0

### Components

- cluster-vsphere from v9.4.1 to v9.3.2
- cluster from v8.4.2 to v8.3.2

### cluster [v8.4.2...v8.3.2](https://github.com/giantswarm/cluster/compare/v8.4.2...v8.3.2)

#### Changed

- Stop deleting and recreating `/etc/ssl/certs` on nodes for fixing SELinux labeling, so any custom certificates placed there directly are preserved.

