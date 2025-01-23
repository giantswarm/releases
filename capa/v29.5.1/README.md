# :zap: Giant Swarm Release v29.5.1 for CAPA :zap:

## Changes compared to v29.5.0

### Components

- cluster-aws from v2.5.0 to v2.5.0-c415effd69a6e2294f5f5ad4f2677deb48887255

### cluster-aws [v2.5.0-v2.5.0-c415effd69a6e2294f5f5ad4f2677deb48887255](https://github.com/giantswarm/cluster-aws/compare/v2.5.0...c415effd69a6e2294f5f5ad4f2677deb48887255)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools

### Apps

- etcd-defrag v0.1.1

### etcd-defrag [v0.1.1](https://github.com/giantswarm/etcd-defrag-app/releases/tag/v0.1.1)

#### Changed

- Values: Rename `cluster` into `useClusterEndpoints`. ([#8](https://github.com/giantswarm/etcd-defrag-app/pull/8))
