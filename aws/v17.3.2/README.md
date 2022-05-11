# :zap: Giant Swarm Release v17.3.2 for AWS :zap:

This release reintroduces tagging private subnets on node pools to enable autodiscovery for internal ELBs by setting the annotation `alpha.aws.giantswarm.io/internal-elb: ""` on AWSMachineDeployment CR's.

## Change details


### aws-operator [11.10.0](https://github.com/giantswarm/aws-operator/releases/tag/v11.10.0)

#### Added
 - Set optionally the `kubernetes.io/role/internal-elb` tag to machine deployment subnets.

