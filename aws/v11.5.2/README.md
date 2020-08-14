# :zap: Giant Swarm Release v11.5.2 for AWS :zap:

Giant Swarm release v11.5.2 is fixing an issue with NetworkPolicies and custom pod cidrs.

## Change details

### aws-operator [8.7.6](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.6)
- Add release version tag for ec2 instances.
- Update Cloudformation Stack when components version differ.
- Fix IAM policy on Tenant Clusters to manages IAM Role tags.
- Fixed passing custom pod CIDR to k8scloudconfig.
