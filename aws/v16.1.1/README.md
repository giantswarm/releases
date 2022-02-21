# :zap: Giant Swarm Release v16.1.1 for AWS :zap:

This release provides a fix for `aws-ebs-csi-driver` to ensure all taints for `ebs-node` are tolerated as well as selecting the right node selector for worker nodes.

## Change details


### aws-ebs-csi-driver [2.8.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.8.1)

#### Fixed
- Use node selector according to control-plane and nodepool labels.



