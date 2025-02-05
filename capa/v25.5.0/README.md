# :zap: Giant Swarm Release v25.5.0 for CAPA :zap:

Most notable change in this release is the reduction of IAM permissions on the worker nodes instance profile, aiming at improving the general security of the clusters. Additional changes include reducing the size of the ETCD volume to 50GB targetting costs saving initiatives, improvements for the `node-termination-handler` application for smoother upgrades and operations as well as fixes for ENI mode targetting the CAPA migration.

## Changes compared to v25.4.0

### Components

- cluster-aws from v1.3.5 to v1.3.7

### cluster-aws [v1.3.5...v1.3.7](https://github.com/giantswarm/cluster-aws/compare/v1.3.5...v1.3.7)

#### Added

- Add ingress rule in nodes Security Group to allow access to the Kubelet API when using ENI mode. This is needed by the metrics server to gather metrics from the Kubelet

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools
- Use reduced IAM permissions on worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.
- Explicitly set aws-node-termination-handler queue region so crash-loops are avoided, allowing faster startup

### Apps

- aws-nth-bundle from v1.2.0 to v1.2.1
- cilium-crossplane-resources from v0.1.0 to v0.2.0

### aws-nth-bundle [v1.2.0...v1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)

#### Added

- Forward proxy settings to `aws-node-termination-handler-app` as environment variables

### cilium-crossplane-resources [v0.1.0...v0.2.0](https://github.com/giantswarm/cilium-crossplane-resources/compare/v0.1.0...v0.2.0)

#### Added

- Add a Security Group rule for node to pod communication
