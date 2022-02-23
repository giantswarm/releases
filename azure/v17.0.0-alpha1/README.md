# :zap: Giant Swarm Release v17.0.0-alpha1 for Azure :zap:

<< Add description here >>

## Change details


### kubernetes [1.22.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.6)

#### Feature
- Kube-apiserver: when merging lists, Server Side Apply now prefers the order of the submitted request instead of the existing persisted object ([#107568](https://github.com/kubernetes/kubernetes/pull/107568), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]
#### Bug or Regression
- An inefficient lock in EndpointSlice controller metrics cache has been reworked. Network programming latency may be significantly reduced in certain scenarios, especially in clusters with a large number of Services. ([#107168](https://github.com/kubernetes/kubernetes/pull/107168), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Client-go: fix that paged list calls with ResourceVersionMatch set would fail once paging kicked in. ([#107335](https://github.com/kubernetes/kubernetes/pull/107335), [@fasaxc](https://github.com/fasaxc)) [SIG API Machinery]
- Fix a panic when using invalid output format in kubectl create secret command ([#107346](https://github.com/kubernetes/kubernetes/pull/107346), [@rikatz](https://github.com/rikatz)) [SIG CLI]
- Fix: azuredisk parameter lowercase translation issue ([#107429](https://github.com/kubernetes/kubernetes/pull/107429), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixes a rare race condition handling requests that timeout ([#107459](https://github.com/kubernetes/kubernetes/pull/107459), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Mount-utils: Detect potential stale file handle ([#107039](https://github.com/kubernetes/kubernetes/pull/107039), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage]
#### Other (Cleanup or Flake)
- Updates konnectivity-network-proxy to v0.0.27. This includes a memory leak fix for the network proxy ([#107187](https://github.com/kubernetes/kubernetes/pull/107187), [@rata](https://github.com/rata)) [SIG API Machinery, Auth and Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.39.2 → v0.39.3](https://github.com/google/cadvisor/compare/v0.39.2...v0.39.3)
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.22 → v0.0.27
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.1
#### Removed
_Nothing has changed._



### cert-operator [1.3.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.3.0)

#### Changed
- Use `RenewSelf` instead of `LookupSelf` to prevent expiration of `Vault token`.



### cluster-operator [3.13.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.13.0)

#### Changed
- Removed encryption key creation. Encryption keys will be managed by `encryption-provider-operator`.



### app-operator [5.6.0](https://github.com/giantswarm/app-operator/releases/tag/v5.6.0)

#### Changed
- Get tarball URL for chart CRs from index.yaml for better community app catalog support.
#### Fixed
- Fix error handling in chart CR watcher when chart CRD not installed.



### azure-operator [5.16.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.16.0)

#### Added
- Add support for feature that enables forcing cgroups v1 for Flatcar version `3033.2.0` and above.
#### Changed
- Upgraded to giantswarm/exporterkit v1.0.0
- Upgraded to giantswarm/microendpoint v1.0.0
- Upgraded to giantswarm/microkit v1.0.0
- Upgraded to giantswarm/micrologger v0.6.0
- Upgraded to giantswarm/versionbundle v1.0.0
- Upgraded to spf13/viper v1.10.0
- Make nodepool nodes roll in case the user switches between cgroups v1 and v2
- Drop dependency on giantswarm/apiextensions/v2
- Bump k8scloudconfig to disable rpc-statd





### calico [3.21.3](https://github.com/projectcalico/calico/releases/tag/v3.21.3)

Not found


### chart-operator [2.20.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.20.0)

#### Changed
- Update Helm to v3.6.3.
- Use controller-runtime client to remove CAPI dependency.
#### Removed
- Remove unused helm 2 release collector.



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



### cluster-autoscaler [1.22.2-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs2)

Not found


### azure-scheduled-events [0.6.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.6.0)

#### Added
- Add `priorityClassName: "system-node-critical"` to Daemonset to give higher priority during scheduling.



