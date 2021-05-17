# :zap: Tenant Cluster Release v13.0.0 for KVM :zap:

This release upgrades Kubernetes to 1.18 and Calico to 3.15. It also includes other minor component updates summarized below.

## Change details


### kubernetes [1.18.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.12)

Kubernetes 1.18 includes a large number of features, deprecations, and bug fixes. We have attempted to summarize the changes relevant to Giant Swarm customers in the following section. The full list of changes can be viewed [here](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md).

Please discuss with your Giant Swarm solution engineer if you are unsure about a cluster's Kubernetes 1.18 upgrade readiness.

#### kube-apiserver
- The following features are unconditionally enabled and the corresponding `--feature-gates` flags have been removed: `PodPriority`, `TaintNodesByCondition`, `ResourceQuotaScopeSelectors` and `ScheduleDaemonSetPods`.
- the following deprecated APIs can no longer be served:
  - All resources under `apps/v1beta1` and `apps/v1beta2` - use `apps/v1` instead
  - `daemonsets`, `deployments`, `replicasets` resources under `extensions/v1beta1` - use `apps/v1` instead
  - `networkpolicies` resources under `extensions/v1beta1` - use `networking.k8s.io/v1` instead
  - `podsecuritypolicies` resources under `extensions/v1beta1` - use `policy/v1beta1` instead

#### kube-scheduler
- The `scheduling_duration_seconds` summary metric is deprecated.
- The `scheduling_algorithm_predicate_evaluation_seconds` and
  `scheduling_algorithm_priority_evaluation_seconds` metrics are deprecated, replaced by `framework_extension_point_duration_seconds[extension_point="Filter"]` and `framework_extension_point_duration_seconds[extension_point="Score"]`.
- `AlwaysCheckAllPredicates` is deprecated in scheduler Policy API.

#### kubelet
- `--enable-cadvisor-json-endpoints` is now disabled by default. If you need access to the cAdvisor v1 Json API please enable it explicitly in the kubelet command line. Please note that this flag was deprecated in 1.15 and will be removed in 1.19.
- The StreamingProxyRedirects feature and `--redirect-container-streaming` flag are deprecated, and will be removed in a future release. The default behavior (proxy streaming requests through the kubelet) will be the only supported option. If you are setting `--redirect-container-streaming=true`, then you must migrate off this configuration. The flag will no longer be able to be enabled starting in v1.20. If you are not setting the flag, no action is necessary.
- resource metrics endpoint `/metrics/resource/v1alpha1` as well as all metrics under this endpoint have been deprecated. Please convert to the following metrics emitted by endpoint `/metrics/resource`:
  - scrape_error --> scrape_error
  - node_cpu_usage_seconds_total --> node_cpu_usage_seconds
  - node_memory_working_set_bytes --> node_memory_working_set_bytes
  - container_cpu_usage_seconds_total --> container_cpu_usage_seconds
  - container_memory_working_set_bytes --> container_memory_working_set_bytes
  - scrape_error --> scrape_error
- In a future release, kubelet will no longer create the CSI NodePublishVolume target directory, in accordance with the CSI specification. CSI drivers may need to be updated accordingly to properly create and process the target path.

### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)

#### Changes
 - Add FelixConfiguration parameters to explicitly allow encapsulated packets from workloads.
 - Respect explicit configuration for drop rules for encapsulated packets originating from workloads.

#### Bug fixes
- Added monitor-addresses option to calico-node to continually monitor IP addresses.
- Fix issue with service IP advertisement breaking host service connectivity.
- Felix FV tests now run with Go’s race detector enabled and a couple of low-impact data races have been fixed.
- Fix config inheritance so that the BPF kernel version check takes precedence over environment variables.
- In BPF mode, fix spurious “Failed to run bpftool” logs.
- Fixed capitalization of WireGuard interfaceIPv4Address (was interfaceIpv4Address).
- Fix race condition during block affinity deletion.

#### Other changes
- Handle panics in the CNI plugin more gracefully cni-plugin.
- Remove unnecessary packages from docker image to address CVEs.
- By default, exclude cni.* from node IP auto detection.
- Added conditional check for FELIX_HEALTHHOST env variable node.
- The Typha port is now included in the failsafe port lists by default.
- Felix can now run in active/passive modes.
- For NetworkPolicy and GlobalNetworkPolicy, the use of floating point values for the spec.Order field is now deprecated, and will be removed entirely in a future release. Please update your policies to use integer values for ordering.
- Update included CustomResourceDefinitions to use the apiextensions/v1 API group and version, and include schemas for basic validation.
- Improve scaling characteristics when using host-local IPAM - perform fewer List API calls.
- Network policy now has the global() namespace selector which selects host endpoints or global network sets.
- Program blackhole routes for full rejectcidrs to avoid route loops.
- install-cni.sh now also fails if calico -v doesn’t work after copying the calico binary cni-plugin.
- Upstream CNI plugins updated to v0.8.6.



### kvm-operator [3.14.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.14.0)

### Added

- Roll nodes when versions of `calico`, `containerlinux`, `etcd`, `kubernetes` change in release and `kvm-operator` version is unchanged.

### Changed

- Update Kubernetes libraries to 1.18 along with all other client-go-dependent libraries.
- Use InternalIP from TC node's status instead of label for dead endpoints detection.
- Shorten `calico-node` wait timeout in `k8s-addons` and add retry for faster cluster initialization.
- Remove unused Kubernetes scheduler configuration fields preventing strict YAML unmarshalling.



### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)

#### Security

- A [log warning](https://github.com/etcd-io/etcd/pull/12242) is added when etcd use any existing directory that has a permission different than 700 on Linux and 777 on Windows.

#### Package `clientv3`

- Remove [excessive watch cancel logging messages](https://github.com/etcd-io/etcd/pull/12187).
  - See [kubernetes/kubernetes#93450](https://github.com/kubernetes/kubernetes/issues/93450).

#### Package `runtime`

- Optimize [`runtime.FDUsage` by removing unnecessary sorting](https://github.com/etcd-io/etcd/pull/12214).

#### Metrics, Monitoring

- Add [`os_fd_used` and `os_fd_limit` to monitor current OS file descriptors](https://github.com/etcd-io/etcd/pull/12214).

#### Package `etcd server`

- Fix [server panic in slow writes warnings](https://github.com/etcd-io/etcd/issues/12197).
  - Fixed via [PR#12238](https://github.com/etcd-io/etcd/pull/12238).
- Improve [`runtime.FDUsage` call pattern to reduce objects malloc of Memory Usage and CPU Usage](https://github.com/etcd-io/etcd/pull/11986).
- Add [`etcd --experimental-watch-progress-notify-interval`](https://github.com/etcd-io/etcd/pull/12216) flag to make watch progress notify interval configurable.
- Add [`--unsafe-no-fsync`](https://github.com/etcd-io/etcd/pull/11946) flag.
  - Setting the flag disables all uses of fsync, which is unsafe and will cause data loss. This flag makes it possible to run an etcd node for testing and development without placing lots of load on the file system.
- Add [etcd --auth-token-ttl](https://github.com/etcd-io/etcd/pull/11980) flag to customize `simpleTokenTTL` settings.
- Improve [runtime.FDUsage objects malloc of Memory Usage and CPU Usage](https://github.com/etcd-io/etcd/pull/11986).
- Improve [mvcc.watchResponse channel Memory Usage](https://github.com/etcd-io/etcd/pull/11987).
- Fix [`int64` convert panic in raft logger](https://github.com/etcd-io/etcd/pull/12106).
  - Fix [kubernetes/kubernetes#91937](https://github.com/kubernetes/kubernetes/issues/91937).

#### Breaking Changes

- Changed behavior on [existing dir permission](https://github.com/etcd-io/etcd/pull/11798). Previously, the permission was not checked on existing data directory and the directory used for automatically generating self-signed certificates for TLS connections with clients. Now a check is added to make sure those directories, if already exist, has a desired permission of 700 on Linux and 777 on Windows.



### kube-state-metrics [1.3.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.0)

#### Added
- Added monitoring annotations and common labels.

#### Changed
- Deploy `kube-state-metrics-app` on installations as part of app collection.



### app-operator [2.7.0](https://github.com/giantswarm/app-operator/releases/tag/v2.7.0)

### Changed

- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.

### Fixed

- Use resourceVersion of configmap for comparison instead of listing option.

### Added

- Secure the webhook with token value from control plane catalog.
- Adding webhook URL as annotation into chart CRs.
- Added Status update endpoint.
- Watch secrets referenced in app CRs to reduce latency when applying config
changes.
- Create appcatalogentry CRs for public app catalogs.
- Watch configmaps referenced in app CRs to reduce latency when applying config
changes.



### chart-operator [2.5.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.5.0)

### Added

- Validate the cache in helmclient to avoid state requests when pulling tarballs.
- Call status webhook with token values.

### Fixed

- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.



### cert-exporter [1.3.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.3.0)

#### Added
- Add Network Policy.
#### Changed
- Remove `hostNetwork` and `hostPID` capabilities.



### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.2)

#### Changed
- Updated backward incompatible Kubernetes dependencies to v1.18.5.
#### Fixed
- Fixed indentation problem with the daemonset template.



### node-exporter [1.7.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.0)

#### Changed
- Disable `btrfs`,`softnet`,`rapl` and `thermal_zone` to reduce memory usage.
- Increase memory limit to `75Mi`.
