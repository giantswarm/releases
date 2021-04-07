# :zap: Giant Swarm Release v14.1.0 for AWS :zap:

This release provides security and bug fixes for various components.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

## Change details


### cluster-operator [3.6.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.6.0)

#### Fixed
- Fix cluster status computation to correctly display rollbacks, version changes and multiple updates.
#### Added
- Add unit tests for cluster status computation
- Check existence of chart tarball for `release` CR `apps` in catalog.
- Add vertical pod autoscaler support.
- Add `appversionlabel` resource to update version labels for optional app CRs.



### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

#### Added
- Include `apiVersion`, `restrictions.compatibleProviders` in appcatalogentry CRs.
#### Changed
- Limit the number of AppCatalogEntry per app.
- Delete legacy finalizers on app CRs.
- Reconciling appCatalog CRs only if pod is unique.
#### Fixed
- Updating status as cordoned if app CR has cordoned annotation.


### kubernetes [1.19.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.9)

### API Change
- Kubernetes is now built using go1.15.8 ([#99093](https://github.com/kubernetes/kubernetes/pull/99093), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
### Feature
- Add a new flag to set priority for the kubelet on Windows nodes so that workloads cannot overwhelm the node there by disrupting kubelet process. ([#96157](https://github.com/kubernetes/kubernetes/pull/96157), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node and Windows]
#### Failing Test
- Fix handing special characters in the volume path on Windows ([#99137](https://github.com/kubernetes/kubernetes/pull/99137), [@yujuhong](https://github.com/yujuhong)) [SIG Storage]
- Resolves an issue running Ingress conformance tests on clusters which use finalizers on Ingress objects to manage releasing load balancer resources ([#96742](https://github.com/kubernetes/kubernetes/pull/96742), [@spencerhance](https://github.com/spencerhance)) [SIG Network and Testing]
#### Bug or Regression
- Count pod overhead against an entity's ResourceQuota ([#99600](https://github.com/kubernetes/kubernetes/pull/99600), [@gjkim42](https://github.com/gjkim42)) [SIG API Machinery and Node]
- EndpointSlice controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100114](https://github.com/kubernetes/kubernetes/pull/100114), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSliceMirroring controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100144](https://github.com/kubernetes/kubernetes/pull/100144), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Fixed bug that caused cAdvisor to incorrectly detect single-socket multi-NUMA topology. ([#99209](https://github.com/kubernetes/kubernetes/pull/99209), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixes kubelet to retrieve number of sockets from cAdvisor MachineInfo, instead of assuming it to be equal to number of NUMA nodes. ([#99771](https://github.com/kubernetes/kubernetes/pull/99771), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#98140](https://github.com/kubernetes/kubernetes/pull/98140), [@CKchen0726](https://github.com/CKchen0726)) [SIG Apps and Node]
- Kubelet now cleans up orphaned volume directories automatically ([#95301](https://github.com/kubernetes/kubernetes/pull/95301), [@lorenz](https://github.com/lorenz)) [SIG Node and Storage]
- Resolves spurious `Failed to list *v1.Secret` or `Failed to list *v1.ConfigMap` messages in kubelet logs. ([#99538](https://github.com/kubernetes/kubernetes/pull/99538), [@liggitt](https://github.com/liggitt)) [SIG Auth and Node]
- Using NUMA nodes instead of sockets for CPU manager hints. ([#99276](https://github.com/kubernetes/kubernetes/pull/99276), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- We will no longer automatically delete all data when a failure is detected during creation of the volume data file on a CSI volume. Now we will only remove the data file and volume path. ([#96021](https://github.com/kubernetes/kubernetes/pull/96021), [@huffmanca](https://github.com/huffmanca)) [SIG Storage]
- Aggregate errors when putting vmss ([#98350](https://github.com/kubernetes/kubernetes/pull/98350), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Avoid marking node as Ready until node has synced with API servers at least once ([#97996](https://github.com/kubernetes/kubernetes/pull/97996), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Cleanup subnet in frontend IP configs to prevent huge subnet request bodies in some scenarios. ([#98288](https://github.com/kubernetes/kubernetes/pull/98288), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix CSI-migrated inline EBS volumes failing to mount if their volumeID is prefixed by aws:// ([#96821](https://github.com/kubernetes/kubernetes/pull/96821), [@wongma7](https://github.com/wongma7)) [SIG Storage]
- Fix azure file migration issue ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx)) [SIG Auth, Cloud Provider and Storage]
- Fix the description of command line flags that can override --config ([#98873](https://github.com/kubernetes/kubernetes/pull/98873), [@changshuchao](https://github.com/changshuchao)) [SIG Scheduling]
- Fix to recover CSI volumes from certain dangling attachments ([#96617](https://github.com/kubernetes/kubernetes/pull/96617), [@yuga711](https://github.com/yuga711)) [SIG Apps and Storage]
- Fixed a bug that the kubelet cannot start on BtrfS. ([#98015](https://github.com/kubernetes/kubernetes/pull/98015), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Fixed a bug where aggregator_unavailable_apiservice metrics were reported for deleted apiservices. ([#96421](https://github.com/kubernetes/kubernetes/pull/96421), [@dgrisonnet](https://github.com/dgrisonnet)) [SIG API Machinery and Instrumentation]
- Fixed provisioning of Cinder volumes migrated to CSI when StorageClass with AllowedTopologies was used. ([#98311](https://github.com/kubernetes/kubernetes/pull/98311), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixes a panic in the disruption budget controller for PDB objects with invalid selectors ([#98776](https://github.com/kubernetes/kubernetes/pull/98776), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Apps]
- Kubeadm: get k8s CI version markers from k8s infra bucket ([#98836](https://github.com/kubernetes/kubernetes/pull/98836), [@hasheddan](https://github.com/hasheddan)) [SIG Cluster Lifecycle and Release]
- Kubelet should ignore cgroup driver check on Windows node. ([#98385](https://github.com/kubernetes/kubernetes/pull/98385), [@pacoxu](https://github.com/pacoxu)) [SIG Node]
- Performance regression [#97685](https://github.com/kubernetes/kubernetes/pull/97685) has been fixed ([#98432](https://github.com/kubernetes/kubernetes/pull/98432), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Static pods will be deleted gracefully. ([#98103](https://github.com/kubernetes/kubernetes/pull/98103), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Truncates a message if it hits the NoteLengthLimit when the scheduler records an event for the pod that indicates the pod has failed to schedule. ([#98715](https://github.com/kubernetes/kubernetes/pull/98715), [@carlory](https://github.com/carlory)) [SIG Scheduling]
- Warning about using a deprecated volume plugin is logged only once. ([#96751](https://github.com/kubernetes/kubernetes/pull/96751), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- AcceleratorStats will be available in the Summary API of kubelet when cri_stats_provider is used. ([#97017](https://github.com/kubernetes/kubernetes/pull/97017), [@ruiwen-zhao](https://github.com/ruiwen-zhao)) [SIG Node]
- Exposes and sets a default timeout for the SubjectAccessReview client for DelegatingAuthorizationOptions ([#95910](https://github.com/kubernetes/kubernetes/pull/95910), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Fix the panic when kubelet registers if a node object already exists with no Status.Capacity or Status.Allocatable ([#96297](https://github.com/kubernetes/kubernetes/pull/96297), [@SataQiu](https://github.com/SataQiu)) [SIG Node]
- Fixed FibreChannel volume plugin corrupting filesystems on detach of multipath volumes. ([#97013](https://github.com/kubernetes/kubernetes/pull/97013), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixed a bug in kubelet that will saturate CPU utilization after containerd got restarted. ([#97176](https://github.com/kubernetes/kubernetes/pull/97176), [@hanlins](https://github.com/hanlins)) [SIG Node]
- Remove ready file and its directory (which is created during volume SetUp) during emptyDir volume TearDown. ([#95770](https://github.com/kubernetes/kubernetes/pull/95770), [@jingxu97](https://github.com/jingxu97)) [SIG Storage]
- Volumebinding: report UnschedulableAndUnresolvable status instead of an error when PVC not found ([#96850](https://github.com/kubernetes/kubernetes/pull/96850), [@cofyc](https://github.com/cofyc)) [SIG Scheduling and Storage]
- Bump node-problem-detector version to v0.8.5 to fix OOM detection in with Linux kernels 5.1+ ([#96716](https://github.com/kubernetes/kubernetes/pull/96716), [@tosi3k](https://github.com/tosi3k)) [SIG Cloud Provider, Scalability and Testing]
- Exposes and sets a default timeout for the SubjectAccessReview client for DelegatingAuthorizationOptions ([#95910](https://github.com/kubernetes/kubernetes/pull/95910), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Fix a bug that DefaultPreemption plugin is disabled when using (legacy) scheduler policy. ([#96472](https://github.com/kubernetes/kubernetes/pull/96472), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Testing]
- Fix bug in JSON path parser where an error occurs when a range is empty ([#95933](https://github.com/kubernetes/kubernetes/pull/95933), [@brianpursley](https://github.com/brianpursley)) [SIG API Machinery]
- Fix memory leak in kube-apiserver when underlying time goes forth and back. ([#96266](https://github.com/kubernetes/kubernetes/pull/96266), [@chenyw1990](https://github.com/chenyw1990)) [SIG API Machinery]
- Fix pull image error from multiple ACRs using azure managed identity ([#96355](https://github.com/kubernetes/kubernetes/pull/96355), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: resize Azure disk issue when it's in attached state ([#96705](https://github.com/kubernetes/kubernetes/pull/96705), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a bug that prevents kubectl to validate CRDs with schema using x-kubernetes-preserve-unknown-fields on object fields.
  Fix kubectl SchemaError on CRDs with schema using x-kubernetes-preserve-unknown-fields on array types. ([#96562](https://github.com/kubernetes/kubernetes/pull/96562), [@gautierdelorme](https://github.com/gautierdelorme)) [SIG API Machinery and Testing]
- Fixes an issue with the max-in-flight API server filter where the filter could take a long time to process an incoming request if it had been a long time since the last request. ([#96282](https://github.com/kubernetes/kubernetes/pull/96282), [@staebler](https://github.com/staebler)) [SIG API Machinery]
- HTTP/2 connection health check is enabled by default in all Kubernetes clients. The feature should work out-of-the-box. If needed, users can tune the feature via the HTTP2_READ_IDLE_TIMEOUT_SECONDS and HTTP2_PING_TIMEOUT_SECONDS environment variables. The feature is disabled if HTTP2_READ_IDLE_TIMEOUT_SECONDS is set to 0. ([#96778](https://github.com/kubernetes/kubernetes/pull/96778), [@caesarxuchao](https://github.com/caesarxuchao)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Node]
- Kubeadm: Fixes a kubeadm upgrade bug that could cause a custom CoreDNS configuration to be replaced with the default. ([#97016](https://github.com/kubernetes/kubernetes/pull/97016), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cluster Lifecycle]
- Kubeadm: fix coredns migration should be triggered when there are newdefault configs during kubeadm upgrade ([#96970](https://github.com/kubernetes/kubernetes/pull/96970), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Metric names for CSI and flexvolume drivers will include the driver name as well as the CSI plugin name. ([#96477](https://github.com/kubernetes/kubernetes/pull/96477), [@mattcary](https://github.com/mattcary)) [SIG Instrumentation and Storage]
- New Azure instance types do now have correct max data disk count information. ([#94340](https://github.com/kubernetes/kubernetes/pull/94340), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Cloud Provider and Storage]
- Resolves a regression in 1.19+ with workloads targeting deprecated beta os/arch labels getting stuck in NodeAffinity status on node startup. ([#96810](https://github.com/kubernetes/kubernetes/pull/96810), [@liggitt](https://github.com/liggitt)) [SIG Node]
- Volume binding: report UnschedulableAndUnresolvable status instead of an error when bound PVs not found ([#96291](https://github.com/kubernetes/kubernetes/pull/96291), [@cofyc](https://github.com/cofyc)) [SIG Apps, Scheduling and Storage]

### Other (Cleanup or Flake)
- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Resolves flakes in the Ingress conformance tests due to conflicts with controllers updating the Ingress object ([#98430](https://github.com/kubernetes/kubernetes/pull/98430), [@liggitt](https://github.com/liggitt)) [SIG Network and Testing]
- Fix Azure file share not deleted issue when the namespace is deleted ([#97417](https://github.com/kubernetes/kubernetes/pull/97417), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix counting error in service/nodeport/loadbalancer quota check ([#97828](https://github.com/kubernetes/kubernetes/pull/97828), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery and Network]
- Fix missing cadvisor machine metrics. ([#97006](https://github.com/kubernetes/kubernetes/pull/97006), [@lingsamuel](https://github.com/lingsamuel)) [SIG Node]
- Fix: azure file latency issue for metadata-heavy workloads ([#97082](https://github.com/kubernetes/kubernetes/pull/97082), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed bug in CPUManager with race on container map access ([#97427](https://github.com/kubernetes/kubernetes/pull/97427), [@klueska](https://github.com/klueska)) [SIG Node]
- GCE Internal LoadBalancer sync loop will now release the ILB IP address upon sync failure. An error in ILB forwarding rule creation will no longer leak IP addresses. ([#97740](https://github.com/kubernetes/kubernetes/pull/97740), [@prameshj](https://github.com/prameshj)) [SIG Cloud Provider and Network]
- Kubeadm: avoid detection of the container runtime for commands that do not need it ([#97848](https://github.com/kubernetes/kubernetes/pull/97848), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Client-go header logging (at verbosity levels >= 9) now masks `Authorization` header contents ([#95316](https://github.com/kubernetes/kubernetes/pull/95316), [@sfowl](https://github.com/sfowl)) [SIG API Machinery]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.37.0 → v0.37.5](https://github.com/google/cadvisor/compare/v0.37.0...v0.37.5)
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.9 → v0.0.15
- golang.org/x/net: ab34263 → 69a7880
- golang.org/x/sys: ed371f2 → 5cba982
#### Removed
_Nothing has changed._



### cert-manager [2.4.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.4.3)

#### Changed
- Set docker.io as the default registry
- Made CRD install Job backoffLimit configurable (and increased the default value). ([#129](https://github.com/giantswarm/cert-manager-app/pull/129))
### Added
- Enabled configuration of certificate Secret deletion when the parent Certificate is deleted. ([#127](https://github.com/giantswarm/cert-manager-app/pull/127))



### external-dns [2.2.2](https://github.com/giantswarm/external-dns-app/releases/tag/v2.2.2)

#### Changed
- Set docker.io as the default registry
### Fixed
- Adds additional options required for vmware installations. ([#74](https://github.com/giantswarm/external-dns-app/pull/74))
### Added
- Add crd source if the provider is vmware. ([#72](https://github.com/giantswarm/external-dns-app/pull/72))



### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry
### Added
- Add exceptions in NetworkPolicies to allow DNS to work correctly through port 53.



### chart-operator [2.13.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.13.0)

### Changed
- `giantswarm-critical` PriorityClass only managed when E2E.
- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.
- Deploy `giantswarm-critical` PriorityClass when it's not found.
### Added
- Updating namespace metadata using namespaceConfig in `Chart` CRs.
- Pause Chart CR reconciliation when it has chart-operator.giantswarm.io/paused=true annotation.
- Use diff key when logging differences between the current and desired release.
- Add support for skip CRD flag when installing Helm releases.
- Added last reconciled timestamp as metrics.
### Fixed
- Stop updating Helm release if it has failed the previous 5 attempts.
- Only create VPA if autoscaling API group is present.



### kiam [1.7.1](https://github.com/giantswarm/kiam-app/releases/tag/v1.7.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.2.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.2.2)

#### Changed
- Set docker.io as the default registry



### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed
- Set docker.io as the default registry
- Update `coredns` to upstream version [1.8.0](https://coredns.io/2020/10/22/coredns-1.8.0-release/).
- Update `coredns` to upstream version [1.7.1](https://coredns.io/2020/09/21/coredns-1.7.1-release/) (including changes introduced in version [1.7.0](https://coredns.io/2020/06/15/coredns-1.7.0-release/)).
- Update `coredns` to upstream version [1.6.9](https://coredns.io/2020/03/24/coredns-1.6.9-release/).
### Added
- Added monitoring annotations and common labels.


### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### cluster-autoscaler [1.19.2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.2)

Not found


### net-exporter [1.9.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.3)

#### Changed
- Set docker.io as the default registry
- Update kubectl image to v1.18.8.
