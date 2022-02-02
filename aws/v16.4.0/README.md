# :zap: Giant Swarm Release v16.4.0 for AWS :zap:

This release includes a new feature that performs automated rotation of the Kubernetes secret encryption key in etcd.
It also includes latest version of Calico.

## Change details


### aws-operator [10.14.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.14.0)

#### Changed
- Changes to EncryptionConfig in order to fully work with `encryption=provider-operator`.


### cluster-operator [3.13.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.13.0)

#### Changed
- Removed encryption key creation. Encryption keys will be managed by `encryption-provider-operator`.


### calico [3.21.3](https://github.com/projectcalico/calico/releases/tag/v3.21.3)


### external-dns [2.9.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.9.0)

This release contains some changes to mitigate rate limiting on AWS clusters. Please take note of the defaults
for values `aws.batchChangeInterval`, `aws.zonesCacheDuration`, `externalDNS.interval`
and `externalDNS.minEventSyncInterval`.
If you already specify `--aws-batch-change-interval` or `--aws-zones-cache-duration`, please migrate to the new values `aws.batchChangeInterval` and `aws.zonesCacheDuration`.
#### Added
- Allow to set `--aws-batch-change-interval` through `aws.batchChangeInterval` value. Default `10s`.
- Allow to set `--aws-zones-cache-duration` through `aws.zonesCacheDuration` value. Default `3h`.
#### Changed
- Set default `externalDNS.interval` to `5m`.
- Set default `externalDNS.minEventSyncInterval` to `30s`.
- Allow setting Route53 credentials (`externalDNS.aws_access_key_id` and `externalDNS.aws_secret_access_key`) indepentent from `aws.access` value.
- Allow setting the AWS default region (`aws.region`) indepentent from `aws.access` value.
- Allow to omit the `--domain-filter` flag completely by setting `externalDNS.domainFilterList` to `null`.
