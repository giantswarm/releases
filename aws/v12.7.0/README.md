# :zap: Giant Swarm Release v12.7.0 for AWS :zap:

This release offers the possibility to configure the subnet size of Network Pools, the size and wait time of batches during tenant cluster upgrades. More details about the upgrade improvments can be found in our [Fine-tuning upgrade disruption on AWS](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/) guide.

## Change details

### aws-cni [1.7.5](https://github.com/aws/amazon-vpc-cni-k8s/blob/master/CHANGELOG.md#v175)

Bug - Match primary ENI IP correctly ([#1247](https://github.com/aws/amazon-vpc-cni-k8s/pull/1247) , @mogren)

### aws-operator [9.3.1](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.1)

#### Changed
- Update dependencies to next major versions.

#### Fixed
- During a deletion of a cluster, ignore volumes that are mounted to an instance in a different cluster.

#### Added
- Annotation `alpha.aws.giantswarm.io/metadata-v2` to enable AWS Metadata API v2
- Annotation `alpha.aws.giantswarm.io/aws-subnet-size` to customize subnet size of Control Plane and Node Pools
- Annotation `alpha.aws.giantswarm.io/update-max-batch-size` to configure max batch size in ASG update policy on cluster or machine deployment CR.
- Annotation `alpha.aws.giantswarm.io/update-pause-time` to configure pause between batches in ASG update on cluster or machine deployment CR.

### cert-manager [2.3.2](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.3.2)

#### Added

- Added values.schema.json for validation of default values. ([#90](https://github.com/giantswarm/cert-manager-app/pull/90))
- Made cert-manager version configurable. ([#91](https://github.com/giantswarm/cert-manager-app/pull/91))

#### Changed

- Updated `cert-manager` to v1.0.4. ([#95](https://github.com/giantswarm/cert-manager-app/pull/95))
- Update RBAC API versions. ([#84](https://github.com/giantswarm/cert-manager-app/pull/84))

#### Fixed

- Updated app version in Chart.yaml metadata to `v1.0.3`. ([#91](https://github.com/giantswarm/cert-manager-app/pull/91))
