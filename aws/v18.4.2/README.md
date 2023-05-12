# :zap: Giant Swarm Release v18.4.2 for AWS :zap:

This is a patch release that fixes a problem causing CNI downtime when upgrading from v18 to v19.

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.3.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#330---2023-05-11) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Change details


### aws-operator [14.13.3](https://github.com/giantswarm/aws-operator/releases/tag/v14.13.3)

#### Fixed
- Ensure `net.ipv4.conf.eth0.rp_filter` is set to `2` if aws-CNI is used.



