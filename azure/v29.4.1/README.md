# :zap: Giant Swarm Release v29.4.1 for Azure :zap:

## Changes compared to v29.4.0

### Components

- cluster-azure from v1.5.0 to v1.5.0-82c9e6dded539b2f311c75dd49718cbd5be8c201

### cluster-aws [v1.5.0-v1.5.0-82c9e6dded539b2f311c75dd49718cbd5be8c201](https://github.com/giantswarm/cluster-azure/compare/v1.5.0...82c9e6dded539b2f311c75dd49718cbd5be8c201)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.

### Apps

- etcd-defrag v0.1.1

### etcd-defrag [v0.1.1](https://github.com/giantswarm/etcd-defrag-app/releases/tag/v0.1.1)

#### Changed

- Values: Rename `cluster` into `useClusterEndpoints`. ([#8](https://github.com/giantswarm/etcd-defrag-app/pull/8))
