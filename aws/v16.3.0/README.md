# :zap: Giant Swarm Release v16.3.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [5.4.0](https://github.com/giantswarm/app-operator/releases/tag/v5.4.0)

#### Changed
- Update Helm to `v3.6.3`.
- Use controller-runtime client to remove CAPI dependency.
- Use `apptestctl` to install CRDs in integration tests to avoid hitting GitHub rate limits.
#### Removed
- Remove `releasemigration` resource now migration to Helm 3 is complete.



### aws-cni [1.10.1](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.10.1)

Not found


### cert-operator [1.3.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.3.0)

#### Changed
- Use `RenewSelf` instead of `LookupSelf` to prevent expiration of `Vault token`.



### aws-operator [10.10.2-dev](https://github.com/giantswarm/aws-operator/releases/tag/v10.10.2-dev)

Not found


### chart-operator [2.20.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.20.0)

#### Changed
- Update Helm to v3.6.3.
- Use controller-runtime client to remove CAPI dependency.
#### Removed
- Remove unused helm 2 release collector.



### coredns [1.7.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.7.0)

#### Changed
- Update `coredns` to upstream version [1.8.6](https://coredns.io/2021/10/07/coredns-1.8.6-release/).



### kiam [2.1.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.1.0)

#### Changed
- Upgrade `kiam` version to 4.2.
- Increase AWS session duration to `60m`.
- Increase AWS session refresh to `10m`.



### kube-state-metrics [1.6.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.6.0)

#### Changed
- Bumped to upstream version v2.3.0.



### metrics-server [1.6.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.6.0)

#### Changed
- Updated metrics-server version to 0.5.2.



### node-exporter [1.9.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.9.0)

#### Updated
- Updated node-exporter version to 1.3.1.



