# :zap: Tenant Cluster Release v11.5.2 for AWS :zap:

This release provides a new aws-operator which is fixing an issue with NetworkPolicies and custom Pod CIDRs.

## Change details

### aws-operator [8.7.6](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.6)

- Add release version tag for ec2 instances.
- Update Cloudformation Stack when components version differ.
- Fix IAM policy on Tenant Clusters to manages IAM Role tags.
- Fixed passing custom Pod CIDR to k8scloudconfig.
