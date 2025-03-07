# :zap: Giant Swarm Release v30.1.0 for CAPA :zap:

## Changes compared to v30.0.0

### Components

- cluster-aws from 3.0.0 to [3.1.1](https://github.com/giantswarm/cluster-aws/compare/v3.0.0...v3.1.1)

### cluster-aws [v3.0.0...v3.1.1](https://github.com/giantswarm/cluster-aws/compare/v3.0.0...v3.1.1)

#### Added

- Add ingress rule in nodes Security Group to allow access to the Cilium Relay when using ENI mode.

#### Changed

- Chart: Update `cluster` to v2.1.1.
- Add option `global.providerSpecific.nodeTerminationHandlerEnabled` to disable the AWS Node Termination Handler (NTH).
