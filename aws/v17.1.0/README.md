# :zap: Giant Swarm Release v17.1.0 for AWS :zap:

 A release that is introducing a new alpha feature IAM roles for service accounts (IRSA). Documentation - [IAM roles for service accounts](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)

## Change details


### cluster-operator [3.14.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.14.0)

#### Added
- Add IAM Roles for Service Accounts feature support for AWS.



### aws-operator [10.18.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.18.0)

#### Added
- Add support for IAM Roles for Service Accounts feature.



### net-exporter [1.11.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.11.0)

#### Added
- Add networkpolicy to allow egress towards `k8s-dns-node-cache-app` endpoints.



### kiam [2.2.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.2.0)

#### Changed
- Updated `whiteListRouteRegexp` to default to `/latest/meta-data/placement/availability-zone`



