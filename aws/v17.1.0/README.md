# :zap: Giant Swarm Release v17.1.0 for AWS :zap:

This release introduces IAM roles for service accounts (IRSA) as an alternative to Kiam. More details are available in the [documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/).

> **_Warning:_** IAM roles for service accounts requires the following additional permissions to be granted:
- `iam:CreateOpenIDConnectProvider`
- `iam:DeleteOpenIDConnectProvider`
- `iam:TagOpenIDConnectProvider`
- `iam:UntagOpenIDConnectProvider`
- `s3:PutObjectAcl`

All the AWS prerequisites are available in the [giantswarm-aws-account-prerequisites repository](https://github.com/giantswarm/giantswarm-aws-account-prerequisites).

## Change details


### cluster-operator [3.14.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.14.1)

#### Added
- Add IAM Roles for Service Accounts feature support for AWS.

### Changed
- Update `aws-pod-identity-webhook` app version.



### aws-operator [10.18.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.18.0)

#### Added
- Add support for IAM Roles for Service Accounts feature.



### net-exporter [1.11.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.11.0)

#### Added
- Add networkpolicy to allow egress towards `k8s-dns-node-cache-app` endpoints.



### kiam [2.2.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.2.0)

#### Changed
- Updated `whiteListRouteRegexp` to default to `/latest/meta-data/placement/availability-zone`

#### Fixed
- Merged two release workflows into one to handle both tags

#### Added
- Build script to generate an IRSA compatible version of each release
