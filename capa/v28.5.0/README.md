# :zap: Giant Swarm Release v28.5.0 for CAPA :zap:

Most notable change in this release is the reduction of IAM permissions on the worker nodes instance profile, aiming at improving the general security of the clusters. Additional changes include reducing the size of the ETCD volume to 50GB targetting costs saving initiatives, as well as improvements for the `node-termination-handler` application for smoother upgrades and operations.

## Changes compared to v28.4.0

### Components

- cluster-aws from v1.3.5 to v1.3.6

### cluster-aws [v1.3.5...v1.3.6](https://github.com/giantswarm/cluster-aws/compare/v1.3.5...v1.3.6)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools
- Use reduced IAM permissions on worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.

#### Fixed

- Explicitly set aws-node-termination-handler queue region so crash-loops are avoided, allowing faster startup

### Apps

- aws-nth-bundle from v1.2.0 to v1.2.1

### aws-nth-bundle [v1.2.0...v1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)

#### Added

- Forward proxy settings to `aws-node-termination-handler-app` as environment variables
