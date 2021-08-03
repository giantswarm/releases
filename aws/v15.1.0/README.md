# :zap: Giant Swarm Release v15.1.0 for AWS :zap:

This release provides stability and bug fixes for various components.

**Highlights**
- Kubernetes 1.20.9;
- Kiam 4.1

## Change details


### kubernetes [1.20.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.9)

#### Feature
- Kubernetes 1.20.x is now built using Go 1.15.14 ([#103677](https://github.com/kubernetes/kubernetes/pull/103677), [@puerco](https://github.com/puerco)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
#### Bug or Regression
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102999](https://github.com/kubernetes/kubernetes/pull/102999), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/structured-merge-diff/v4: v4.0.3 → v4.1.2
#### Removed
_Nothing has changed._



### etcd [3.14.16](https://github.com/etcd-io/etcd/releases/tag/v3.14.16)

#### etcd server
- Log [successful etcd server-side health check in debug level](https://github.com/etcd-io/etcd/pull/12677).
- Fix [64 KB websocket notification message limit](https://github.com/etcd-io/etcd/pull/12402).
- Add [`--experimental-warning-apply-duration`](https://github.com/etcd-io/etcd/pull/12448) flag which allows apply duration threshold to be configurable.
- Fix [`--unsafe-no-fsync`](https://github.com/etcd-io/etcd/pull/12751) to still write-out data avoiding corruption (most of the time).
- Reduce [around 30% memory allocation by logging range response size without marshal](https://github.com/etcd-io/etcd/pull/12871).
- Add [exclude alarms from health check conditionally](https://github.com/etcd-io/etcd/pull/12880).

#### Package `fileutil`
- Fix [`F_OFD_` constants](https://github.com/etcd-io/etcd/pull/12444).

#### Metrics
- Fix [incorrect metrics generated when clients cancel watches](https://github.com/etcd-io/etcd/pull/12803) back-ported from (https://github.com/etcd-io/etcd/pull/12196).

#### Dependency
- Bump up [`gorilla/websocket` to v1.4.2](https://github.com/etcd-io/etcd/pull/12645).

#### Go
- Compile with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).



### calico [3.15.5](https://github.com/projectcalico/calico/releases/tag/v3.15.5)

#### Bug fixes
 - Fix that calico/node would fail to set NetworkUnavailable to false for etcd clusters with mismatched nodenames [node #949](https://github.com/projectcalico/node/pull/949) (@caseydavenport)
 - Fixes a bug where IPv6 networks were not handled properly by the failsafe rules [felix #2748](https://github.com/projectcalico/felix/pull/2748) (@mgleung)
 - Fix that, after a netlink read failure, Felix would tight loop reading from a closed channel.  Restart the event poll in that case. [felix #2713](https://github.com/projectcalico/felix/pull/2713) (@fasaxc)
#### Other changes
 - FailsafeInboundHostPorts & FailsafeOutboundHostPorts now support restricting to specific cidrs. New format <protocol>:<net>:<port> [felix #2721](https://github.com/projectcalico/felix/pull/2721) (@mgleung)



### app-operator [5.1.0](https://github.com/giantswarm/app-operator/releases/tag/v5.1.0)

#### Changed
- Create `AppCatalogEntry` CRs into the same namespace of Catalog CR.
- Include `chart.keywords`, `chart.description` and `chart.upstreamChartVersion` in `AppCatalogEntry` CRs.
- Create `AppCatalog` CRs from `Catalog` CRs for compatibility with existing app-operator releases.
- Prepare helm values to configuration management.
- Use `Catalog` CRs in `App` controller.
- Reconcile to `Catalog` CRs instead of `AppCatalog`.
- Get `Chart` CRD from the GitHub resources.
- Get metadata constants from k8smetadata library not apiextensions.

#### Fixed
- For the chart CR watcher get the kubeconfig secret from the chart-operator app CR to avoid hardcoding it.
- Quote namespace in helm templates to handle numeric workload cluster IDs.



### aws-ebs-csi-driver [2.2.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.2.0)

#### Added
- CRD for snapshot-controller.

### Changed
- Update aws-ebs-csi-driver to v1.1.1.
- Reduce default log level to 2.
- Default volume resizing.



### kiam [2.0.0](https://github.com/giantswarm/kiam-app/releases/tag/v2.0.0)

#### Changed
- Upgrade `kiam` version to 4.1.
- Update RBAC API version from `v1beta1` to `v1`.
- Add `kind: Issuer` and `group: cert-manager.io` to `Certificate` templates.



### cert-manager [2.8.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.8.0)

#### Changed
- Label deployments with `giantswarm.io/monitoring_basic_sli: "true"`. ([#171](https://github.com/giantswarm/cert-manager-app/pull/171))
- Migrate values file structure to match `config` repo. ([#172](https://github.com/giantswarm/cert-manager-app/pull/172))



### coredns [1.6.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.0)

#### Changed
- Make `targetCPUUtilizationPercentage` in HPA configurable.
- Update `coredns` to upstream version [1.8.3](https://coredns.io/2021/02/24/coredns-1.8.3-release/).
- Increase maximum replica count to 50 when using horizontal pod autoscaling.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

#### Added
- Added new configuration value `extraArgs`.



### net-exporter [1.10.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.2)

#### Changed
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.



### external-dns [2.4.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.4.0)

#### Changed
- Upgrade upstream external-dns from v0.7.6 to [v0.8.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.8.0).
- Allow to configure the minimum interval between two consecutive synchronizations triggered from kubernetes events through `externalDNS.minEventSyncInterval`.



### cert-exporter [1.7.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.7.1)

#### Fixed
- Fix configuration version in `Chart.yaml`.



### chart-operator [2.18.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.18.0)

#### Added
- Add releasemaxhistory resource which ensures we retry at a reduced rate when there are repeated failed upgrades.
- Proxy support in helm template.

#### Changed
- Upgrade Helm release when failed even if version or values have not changed to handle situations like failed webhooks where we should retry.
- Prepare helm values to configuration management.
- Update architect-orb to v3.0.0.
- [CAPI] Add tolerations to start on `NotReady` nodes for installing CNI.
- [CAPI] Create `giantswarm-critical` priority class.
- [CAPI] Use host network to allow installing CNI packaged as an app.

#### Fixed
- Improve status message when helm release has failed max number of attempts.
