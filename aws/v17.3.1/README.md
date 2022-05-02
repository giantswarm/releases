# :zap: Giant Swarm Release v17.3.1 for AWS :zap:

This release fixes a bug when external SNAT is enabled. 

## Change details


### aws-operator [11.9.3](https://github.com/giantswarm/aws-operator/releases/tag/v11.9.3)

#### Fixed
- Set `AWS_VPC_K8S_CNI_RANDOMIZESNAT` to `prng` when SNAT is enabled.



