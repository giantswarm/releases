# :zap: Giant Swarm Release v28.5.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v28.4.0

### Components



* cluster-aws from 1.3.5 to [1.3.6](https://github.com/giantswarm/cluster-aws/compare/v1.3.5...v1.3.6)






### cluster-aws [1.3.6](https://github.com/giantswarm/cluster-aws/compare/v1.3.5...v1.3.6)

#### Changed

- Chart: Reduce default etcd volume size to 50 GB.
- Explicitly set Ignition user data storage type to S3 bucket objects for machine pools
- Use reduced IAM permissions on worker nodes instance profile. This can be toggled back with `global.providerSpecific.reducedInstanceProfileIamPermissionsForWorkers`.

#### Fixed

- Explicitly set aws-node-termination-handler queue region so crash-loops are avoided, allowing faster startup





### Apps



* aws-nth-bundle from 1.2.0 to [1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)





### aws-nth-bundle [1.2.1](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.0...v1.2.1)

#### Added

- Forward proxy settings to `aws-node-termination-handler-app` as environment variables




