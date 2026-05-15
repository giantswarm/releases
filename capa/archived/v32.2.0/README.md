# :zap: Giant Swarm Release v32.2.0 for CAPA :zap:

This release updates the `cluster-aws` chart to address an issue with Karpenter nodes not working properly with ingress load balancers.

## Changes compared to v32.1.0

### Components

- cluster-aws from v5.3.0 to v5.4.0

### cluster-aws [v5.3.0...v5.4.0](https://github.com/giantswarm/cluster-aws/compare/v5.3.0...v5.4.0)

#### Added

- *This change will roll the nodes on Karpenter node pools* Attach the `lb` Security Group to Karpenter nodes.
- *This change will roll the nodes on Karpenter node pools* Name instance on AWS after the nodepool name.
