# :zap: Giant Swarm Release v12.7.0 for AWS :zap:

This minor release adds an alpha feature that allows to configure values for AWS CF ASG UpdatePolicy values. Check guide [Fine-tuning upgrade disruption on AWS](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/) for more information on this feature.

## Change details

### aws-cni [1.7.5](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/CHANGELOG.md#v175)

Bug - Match primary ENI IP correctly ([#1247](https://github.com/aws/amazon-vpc-cni-k8s/pull/1247) , @mogren)

### aws-operator [9.3.1](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.0)

#### Changed
- Update dependencies to next major versions.

#### Fixed
- During a deletion of a cluster, ignore volumes that are mounted to an instance in a different cluster.

#### Added
- Annotation `alpha.aws.giantswarm.io/metadata-v2` to enable AWS Metadata API v2
- Annotation `alpha.aws.giantswarm.io/aws-subnet-size` to customize subnet size of Control Plane and Node Pools
- Annotation `alpha.aws.giantswarm.io/update-max-batch-size` to configure max batch size in ASG update policy on cluster or machine deployment CR.
- Annotation `alpha.aws.giantswarm.io/update-pause-time` to configure pause between batches in ASG update on cluster or machine deployment CR.
