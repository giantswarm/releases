# :zap: Giant Swarm Release v18.3.1 for AWS :zap:

This release contains a bugfix for Giantswarm `aws-operator`. It fixes an issue where the `aws-operator` would not be able to create any new cluster because of the recent changes for S3 buckets. [Official documentation available here.](https://aws.amazon.com/blogs/aws/heads-up-amazon-s3-security-changes-are-coming-in-april-of-2023/)

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.1.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#310---2023-04-27) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Change details


### aws-operator [14.12.2](https://github.com/giantswarm/aws-operator/releases/tag/v14.12.2)

#### Fixed
- Allow to enable ACLs for a S3 buckets.



