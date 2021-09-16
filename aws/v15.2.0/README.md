# :zap: Giant Swarm Release v15.2.0 for AWS :zap:

This release adds support to `aws-operator` to comply with the following additional AWS S3 policies:
* `s3-bucket-public-read-prohibited`
* `s3-bucket-ssl-requests-only`
* `s3-bucket-public-write-prohibited`
* `s3-bucket-server-side-encryption-enabled`
* `s3-bucket-logging-enabled`.

## Change details


### aws-operator [10.7.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.7.1)

### Added

- Add security settings to S3 bucket to comply with aws policies `s3-bucket-public-read-prohibited,s3-bucket-ssl-requests-only,s3-bucket-public-write-prohibited,s3-bucket-server-side-encryption-enabled,s3-bucket-logging-enabled`, `aws-operator` will need additonal permissions `s3:PutBucketPublicAccessBlock` and `s3:PutBucketPolicy`.
