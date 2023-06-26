# :zap: Giant Swarm Release v18.4.1 for AWS :zap:

This release contains small improvements. It disables ETCD compaction request from apiserver which is included in etcd by default and upgrades `observability-bundle`.

***Important for IRSA***
When upgrading to AWS Release `v18.4.1` you can additionally set a annotation on AWSCluster CR `alpha.aws.giantswarm.io/enable-cloudfront-alias: ""` to enable the usage of the Cloudfront alternate domain name before v19 where it will be a default. This is useful if you want to take immeditate actions replacing `Kiam`.

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.2.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#320---2023-04-27) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Change details


### aws-operator [14.13.2](https://github.com/giantswarm/aws-operator/releases/tag/v14.13.2)

#### Changed
- Disable ETCD compaction request from apiserver.



### observability-bundle [0.4.3](https://github.com/giantswarm/observability-bundle/releases/tag/v0.4.3)

- Upgrade `prometheus-operator-app` to 4.2.3.



