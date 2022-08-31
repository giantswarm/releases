# :zap: Giant Swarm Release v18.0.0 for Azure :zap:

This is the first Azure release featuring Kubernetes 1.23.
Furthermore, this release is the first to use out-of-tree controller manager and CSI providers.

## Change details

### azure-operator [6.0.0](https://github.com/giantswarm/azure-operator/releases/tag/v6.0.0)

#### Changed
- Support for k8s 1.23.
- Config changes needed to run out-of-tree controller manager and CSI providers. 

### kubernetes [1.23.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.9)

#### Bug or Regression
- Fix a bug that caused the wrong result length when using --chunk-size and --selector together ([#110757](https://github.com/kubernetes/kubernetes/pull/110757), [@Abirdcfly](https://github.com/Abirdcfly)) [SIG API Machinery and Testing]
- Fix bug that prevented the job controller from enforcing activeDeadlineSeconds when set ([#110545](https://github.com/kubernetes/kubernetes/pull/110545), [@harshanarayana](https://github.com/harshanarayana)) [SIG Apps]
- Fix image pulling failure when IMDS is unavailable in kubelet startup ([#110523](https://github.com/kubernetes/kubernetes/pull/110523), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix printing resources with int64 fields ([#110602](https://github.com/kubernetes/kubernetes/pull/110602), [@sanchezl](https://github.com/sanchezl)) [SIG API Machinery]
- Fixed a regression introduced in 1.23.0 where Azure load balancers were not kept up to date with the state of cluster nodes. In particular, nodes that are not in the ready state and are not newly created (i.e. not having the `node.cloudprovider.kubernetes.io/uninitialized` taint) now get removed from Azure load balancers. ([#109932](https://github.com/kubernetes/kubernetes/pull/109932), [@ricky-rav](https://github.com/ricky-rav)) [SIG Cloud Provider]
- Fixed potential scheduler crash when scheduling with unsatisfied nodes in PodTopologySpread. ([#110853](https://github.com/kubernetes/kubernetes/pull/110853), [@kerthcet](https://github.com/kerthcet)) [SIG Scheduling]
- Kubeadm: fix the bug that configurable KubernetesVersion not respected during kubeadm join ([#111022](https://github.com/kubernetes/kubernetes/pull/111022), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Reduced time taken to sync proxy rules on Windows kube-proxy with kernelspace mode ([#110702](https://github.com/kubernetes/kubernetes/pull/110702), [@daschott](https://github.com/daschott)) [SIG Network and Windows]
- Updated cAdvisor to v0.43.1 to pick up a kubelet fix where network metrics can be missing in some cases when used with containerd ([#111013](https://github.com/kubernetes/kubernetes/pull/111013), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.43.0 → v0.43.1](https://github.com/google/cadvisor/compare/v0.43.0...v0.43.1)
#### Removed
_Nothing has changed._



### chart-operator [2.29.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.29.0)

Upgraded from version 2.24.1.

#### Changed

- Added support for k8s 1.24.
- Allow running on clusters where critical workloads are not running yet.
 
#### Fixed

- Better handling of priority class creation.


### coredns [1.11.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.11.0)

#### Changed

- Update coredns to upstream version 1.9.3.
 

### external-dns [2.15.2](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.2)

#### Changed

- Bumped base image to address security issues.


### kube-state-metrics [1.12.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.12.1)

#### Changed

- Bump to upstream version 2.5.0.


### metrics-server [1.8.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.8.0)

#### Changed

- Updated metrics-server version to 0.6.1.


### cluster-autoscaler [1.23.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.23.1)

#### Changed

- Update cluster-autoscaler to version 1.23.1.


### vertical-pod-autoscaler [2.5.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.5.0)

### Changed

- Upgrade vertical-pod-autoscaler to [0.11.0](https://github.com/kubernetes/autoscaler/releases/tag/vertical-pod-autoscaler-0.11.0)

  Potentially breaking change:
  - Added validation - CPU values will be accepted only with resolution of 1 mCPU, memory with resolution of 1 b

  Other changes:
  - Switch to go 1.16
  - Admission controller now logs when it fails to start
  - Increase resolution of admission_latency_seconds metric
  - Reduce verbosity of some logs


### app-operator [6.3.0](https://github.com/giantswarm/app-operator/releases/tag/v6.3.0)

#### Added

- Added support for k8s 1.24.
- Watch config maps and secrets listed in the extraConfigs section of App CR for multi layer configs, see: https://github.com/giantswarm/rfc/tree/main/multi-layer-app-config#enhancing-app-cr
- If no userconfig configmap or secret reference is specified but one is found following the default naming convention (*-user-values / *-user-secrets) then the App resource is updated to reference the found configmap/secret.


### cluster-operator [4.6.1](https://github.com/giantswarm/cluster-operator/releases/tag/v4.6.1)

#### Changed

- Enable cni.install mode for Chart Operator.

### Fixed

- Fix handling of managed apps renamed to renove the `-app` suffix.
 

### calico [3.21.6](https://github.com/projectcalico/calico/releases/tag/v3.21.6)

Please refer to the [upstream changelog](https://projectcalico.docs.tigera.io/archive/v3.21/release-notes/#v3216).


### containerlinux [3227.2.1](https://www.flatcar-linux.org/releases/#release-3227.2.1)

New Stable Release 3227.2.1

Changes since Stable 3227.2.0

#### Security fixes:

- Linux ([CVE-2022-23816](https://nvd.nist.gov/vuln/detail/CVE-2022-23816), [CVE-2022-23825](https://nvd.nist.gov/vuln/detail/CVE-2022-23825), [CVE-2022-29900](https://nvd.nist.gov/vuln/detail/CVE-2022-29900), [CVE-2022-29901](https://nvd.nist.gov/vuln/detail/CVE-2022-29901))

#### Bug fixes:

- Added support for Openstack for cloud-init activation ([flatcar-linux/init#76](https://github.com/flatcar-linux/init/pull/76))
- Excluded Wireguard interface from `systemd-networkd` default management ([Flatcar#808](https://github.com/flatcar-linux/Flatcar/issues/808))
- Fixed `/etc/resolv.conf` symlink by pointing it at `resolv.conf` instead of `stub-resolv.conf`. This bug was present since the update to systemd v250 ([coreos-overlay#2057](https://github.com/flatcar-linux/coreos-overlay/pull/2057))
- Fixed excluded interface type from default systemd-networkd configuration ([flatcar-linux/init#78](https://github.com/flatcar-linux/init/pull/78))
- Fixed space escaping in the `networkd` Ignition translation ([Flatcar#812](https://github.com/flatcar-linux/Flatcar/issues/812))

#### Updates:

- Linux ([5.15.58](https://lwn.net/Articles/902917) (includes [5.15.57](https://lwn.net/Articles/902317), [5.15.56](https://lwn.net/Articles/902101)))
- ca-certificates ([3.81](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_81.html))

### kubernetes [1.23.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.9)

Upgraded from version 1.22. Please refer to the [official changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.23.md) for all details.

#### What's New (Major Themes)

##### Deprecation of FlexVolume

FlexVolume is deprecated. Out-of-tree CSI driver is the recommended way to write volume drivers in Kubernetes.
See [this doc](https://github.com/kubernetes/community/blob/master/sig-storage/volume-plugin-faq.md#kubernetes-volume-plugin-faq-for-storage-vendors) for more information.
Maintainers of FlexVolume drivers should implement a CSI driver and move users of FlexVolume to CSI.
Users of FlexVolume should move their workloads to CSI driver.

##### Deprecation of klog specific flags

To simplify the code base, several [logging flags got marked as deprecated](https://kubernetes.io/docs/concepts/cluster-administration/system-logs/#klog) in Kubernetes 1.23.
The code which implements them will be removed in a future release, so users of those need to start replacing the deprecated flags with some alternative solutions.

##### Software Supply Chain SLSA Level 1 Compliance in the Kubernetes Release Process

Kubernetes releases are now generating provenance attestation files describing the staging and release phases of the release process and artifacts are verified as they are handed over from one phase to the next.
This final piece completes the work needed to comply with Level 1 of the [SLSA security framework](https://slsa.dev/) (Supply-chain Levels for Software Artifacts).

##### IPv4/IPv6 Dual-stack Networking graduates to GA

[IPv4/IPv6 dual-stack networking](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/563-dual-stack) graduates to GA.
Since 1.21, Kubernetes clusters are enabled to support dual-stack networking by default.
In 1.23, the `IPv6DualStack` feature gate is removed.
The use of dual-stack networking is not mandatory.
Although clusters are enabled to support dual-stack networking, Pods and Services continue to default to single-stack.
To use dual-stack networking: Kubernetes nodes have routable IPv4/IPv6 network interfaces, a dual-stack capable CNI network plugin is used, Pods are configured to be dual-stack and Services have their `.spec.ipFamilyPolicy` field set to either `PreferDualStack` or `RequireDualStack`.

##### HorizontalPodAutoscaler v2 graduates to GA

Version 2 of the HorizontalPodAutoscaler API graduates to stable in the 1.23 release. The HorizontalPodAutoscaler `autoscaling/v2beta2` API is deprecated in favor of the new `autoscaling/v2` API, which the Kubernetes project recommends for all use cases.

This release does *not* deprecate the v1 HorizontalPodAutoscaler API.

##### Generic Ephemeral Volume feature graduates to GA

The generic ephemeral volume feature moved to GA in 1.23.
This feature allows any existing storage driver that supports dynamic provisioning to be used as an ephemeral volume with the volume’s lifecycle bound to the Pod.
All StorageClass parameters for volume provisioning and all features supported with PersistentVolumeClaims are supported.

##### Skip Volume Ownership change graduates to GA

The feature to configure volume permission and ownership change policy for Pods moved to GA in 1.23.
This allows users to skip recursive permission changes on mount and speeds up the pod start up time.

##### Allow CSI drivers to opt-in to volume ownership and permission change graduates to GA

The feature to allow CSI Drivers to declare support for fsGroup based permissions graduates to GA in 1.23.

##### PodSecurity graduates to Beta

[PodSecurity](https://kubernetes.io/docs/concepts/security/pod-security-admission/) moves to Beta.
`PodSecurity` replaces the deprecated `PodSecurityPolicy` admission controller.
`PodSecurity` is an admission controller that enforces Pod Security Standards on Pods in a Namespace based on specific namespace labels that set the enforcement level.
In 1.23, the `PodSecurity` feature gate is enabled by default.

##### Container Runtime Interface (CRI) v1 is default

The Kubelet now supports the CRI `v1` API, which is now the project-wide default.
If a container runtime does not support the `v1` API, Kubernetes will fall back to the `v1alpha2` implementation.
There is no intermediate action required by end-users, because `v1` and `v1alpha2` do not differ in their implementation.
It is likely that `v1alpha2` will be removed in one of the future Kubernetes releases to be able to develop `v1`.

##### Structured logging graduate to Beta

Structured logging reached its Beta milestone. Most log messages from kubelet and kube-scheduler have been converted. Users are encouraged to try out JSON output or parsing of the structured text format and provide feedback on possible solutions for the open issues, such as handling of multi-line strings in log values.

##### Simplified Multi-point plugin configuration for scheduler

The kube-scheduler is adding a new, simplified config field for Plugins to allow multiple extension points to be enabled in one spot.
The new `multiPoint` plugin field is intended to simplify most scheduler setups for administrators.
Plugins that are enabled via `multiPoint` will automatically be registered for each individual extension point that they implement.
For example, a plugin that implements Score and Filter extensions can be simultaneously enabled for both.
This means entire plugins can be enabled and disabled without having to manually edit individual extension point settings.
These extension points can now be abstracted away due to their irrelevance for most users.

##### CSI Migration updates

CSI Migration enables the replacement of existing in-tree storage plugins such as `kubernetes.io/gce-pd` or `kubernetes.io/aws-ebs` with a corresponding CSI driver.
If CSI Migration is working properly, Kubernetes end users shouldn’t notice a difference.
After migration, Kubernetes users may continue to rely on all the functionality of in-tree storage plugins using the existing interface.
- CSI Migration feature is turned on by default but stays in Beta for GCE PD, AWS EBS, and Azure Disk in 1.23.
- CSI Migration is introduced as an Alpha feature for Ceph RBD and Portworx in 1.23.

#### Urgent Upgrade Notes

##### (No, really, you MUST read this before you upgrade)

- Kubeadm: remove the deprecated flag `--experimental-patches` for the `init|join|upgrade` commands. The flag `--patches` is no longer allowed in a mixture with the flag `--config`. Please use the kubeadm configuration for setting patches for a node using `{Init|Join}Configuration.patches`. ([#104065](https://github.com/kubernetes/kubernetes/pull/104065), [@pacoxu](https://github.com/pacoxu))
- Log messages in JSON format are written to stderr by default now (same as text format) instead of stdout. Users who expected JSON output on stdout must now capture stderr instead or in addition to stdout. ([#106146](https://github.com/kubernetes/kubernetes/pull/106146), [@pohly](https://github.com/pohly)) [SIG API Machinery, Architecture, Cluster Lifecycle and Instrumentation]
- Support for the seccomp annotations `seccomp.security.alpha.kubernetes.io/pod` and `container.seccomp.security.alpha.kubernetes.io/[name]` has been deprecated since 1.19, will be dropped in 1.25. Transition to using the `seccompProfile` API field. ([#104389](https://github.com/kubernetes/kubernetes/pull/104389), [@saschagrunert](https://github.com/saschagrunert))
- [kube-log-runner](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/component-base/logs/kube-log-runner) is included in release tar balls. It can be used to replace the deprecated `--log-file` parameter. ([#106123](https://github.com/kubernetes/kubernetes/pull/106123), [@pohly](https://github.com/pohly)) [SIG API Machinery, Architecture, Cloud Provider, Cluster Lifecycle and Instrumentation]
- Kubernetes is built using golang 1.17. This version of go removes the ability to use a `GODEBUG=x509ignoreCN=0` environment setting to re-enable deprecated legacy behavior of treating the CommonName of X.509 serving certificates as a host name. This behavior has been disabled by default since Kubernetes 1.19 / go 1.15. Serving certificates used by admission webhooks, custom resource conversion webhooks, and aggregated API servers must now include valid Subject Alternative Names. If you are running Kubernetes 1.22 with `GODEBUG=x509ignoreCN=0` set, check the `apiserver_kube_aggregator_x509_missing_san_total` and `apiserver_webhooks_x509_missing_san_total` metrics for non-zero values to see if the API server is connecting to webhooks or aggregated API servers using certificates that will be considered invalid in Kubernetes 1.23+.

#### Known Issues

##### Etcd v3.5.[0-2] data corruption

Data corruption issue was found in etcd v3.5.0 release that was shipped with 1.22 Kubernetes release. Please read up-to-date [production recommendations for etcd](https://github.com/etcd-io/etcd/tree/main/CHANGELOG).

##### Deprecation

- A deprecation notice has been added when using the kube-proxy userspace proxier, which will be removed in v1.25. (#103860) ([#104631](https://github.com/kubernetes/kubernetes/pull/104631), [@perithompson](https://github.com/perithompson))
- Added `apiserver_longrunning_requests` metric to replace the soon to be deprecated `apiserver_longrunning_gauge` metric. ([#103799](https://github.com/kubernetes/kubernetes/pull/103799), [@jyz0309](https://github.com/jyz0309))
- Controller-manager: the following flags have no effect and would be removed in v1.24:
  - `--port`
  - `--address`
    The insecure port flags `--port` may only be set to 0 now.
- Kube-scheduler: the `--port` and `--address` flags have no effect and would be removed in v1.24.
  The insecure port flags `--port` may only be set to 0 now.
  Also `metricsBindAddress` and `healthzBindAddress` fields from `kubescheduler.config.k8s.io/v1beta1` are no-op and expected to be empty. Removed in `kubescheduler.config.k8s.io/v1beta2` completely. ([#96345](https://github.com/kubernetes/kubernetes/pull/96345), [@ingvagabund](https://github.com/ingvagabund))
  In addition, please be careful that:
  - kube-scheduler MUST start with `--authorization-kubeconfig` and `--authentication-kubeconfig` correctly set to get authentication/authorization working.
  - liveness/readiness probes to kube-scheduler MUST use HTTPS now, and the default port has been changed to 10259.
  - Applications that fetch metrics from kube-scheduler should use a dedicated service account which is allowed to access nonResourceURLs `/metrics`. ([#96345](https://github.com/kubernetes/kubernetes/pull/96345), [@ingvagabund](https://github.com/ingvagabund)) [SIG Cloud Provider, Scheduling and Testing]
- Feature-gate `VolumeSubpath` has been deprecated and cannot be disabled. It will be completely removed in 1.25 ([#105474](https://github.com/kubernetes/kubernetes/pull/105474), [@mauriciopoppe](https://github.com/mauriciopoppe))
- Kubeadm: add a new `output/v1alpha2` API that is identical to the `output/v1alpha1`, but attempts to resolve some internal dependencies with the `kubeadm/v1beta2` API. The `output/v1alpha1` API is now deprecated and will be removed in a future release. ([#105295](https://github.com/kubernetes/kubernetes/pull/105295), [@neolit123](https://github.com/neolit123))
- Kubeadm: add the kubeadm specific, Alpha (disabled by default) feature gate UnversionedKubeletConfigMap. When this feature is enabled kubeadm will start using a new naming format for the ConfigMap where it stores the KubeletConfiguration structure. The old format included the Kubernetes version - "kube-system/kubelet-config-1.22", while the new format does not - "kube-system/kubelet-config". A similar formatting change is done for the related RBAC rules. The old format is now DEPRECATED and will be removed after the feature graduates to GA. When writing the ConfigMap kubeadm (init, upgrade apply) will respect the value of UnversionedKubeletConfigMap, while when reading it (join, reset, upgrade), it would attempt to use new format first and fallback to the legacy format if needed. ([#105741](https://github.com/kubernetes/kubernetes/pull/105741), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle and Testing]
- Kubeadm: remove the deprecated / NO-OP phase `update-cluster-status` in `kubeadm reset` ([#105888](https://github.com/kubernetes/kubernetes/pull/105888), [@neolit123](https://github.com/neolit123))
- Remove 'master' as a valid EgressSelection type in the EgressSelectorConfiguration API. ([#102242](https://github.com/kubernetes/kubernetes/pull/102242), [@pacoxu](https://github.com/pacoxu))
- Removed `kubectl --dry-run` empty default value and boolean values. `kubectl --dry-run` usage must be specified with `--dry-run=(server|client|none)`. ([#105327](https://github.com/kubernetes/kubernetes/pull/105327), [@julianvmodesto](https://github.com/julianvmodesto))
- Removed deprecated metric `scheduler_volume_scheduling_duration_seconds`. ([#104518](https://github.com/kubernetes/kubernetes/pull/104518), [@dntosas](https://github.com/dntosas))
- The deprecated `--experimental-bootstrap-kubeconfig` flag has been removed.
  This can be set via `--bootstrap-kubeconfig`. ([#103172](https://github.com/kubernetes/kubernetes/pull/103172), [@niulechuan](https://github.com/niulechuan))

#### API Change

- A new field `omitManagedFields` has been added to both `audit.Policy` and `audit.PolicyRule`
  so cluster operators can opt in to omit managed fields of the request and response bodies from
  being written to the API audit log. ([#94986](https://github.com/kubernetes/kubernetes/pull/94986), [@tkashem](https://github.com/tkashem)) [SIG API Machinery, Auth, Cloud Provider and Testing]
- A small regression in Service updates was fixed. The circumstances are so unlikely that probably nobody would ever hit it. ([#104601](https://github.com/kubernetes/kubernetes/pull/104601), [@thockin](https://github.com/thockin))
- Added a feature gate `StatefulSetAutoDeletePVC`, which allows PVCs automatically created for StatefulSet pods to be automatically deleted. ([#99728](https://github.com/kubernetes/kubernetes/pull/99728), [@mattcary](https://github.com/mattcary))
- Client-go impersonation config can specify a UID to pass impersonated uid information through in requests. ([#104483](https://github.com/kubernetes/kubernetes/pull/104483), [@margocrawf](https://github.com/margocrawf))
- Create HPA v2 from v2beta2 with some fields changed. ([#102534](https://github.com/kubernetes/kubernetes/pull/102534), [@wangyysde](https://github.com/wangyysde)) [SIG API Machinery, Apps, Auth, Autoscaling and Testing]
- Ephemeral containers graduated to beta and are now available by default. ([#105405](https://github.com/kubernetes/kubernetes/pull/105405), [@verb](https://github.com/verb))
- Fix kube-proxy regression on UDP services because the logic to detect stale connections was not considering if the endpoint was ready. ([#106163](https://github.com/kubernetes/kubernetes/pull/106163), [@aojea](https://github.com/aojea)) [SIG API Machinery, Apps, Architecture, Auth, Autoscaling, CLI, Cloud Provider, Contributor Experience, Instrumentation, Network, Node, Release, Scalability, Scheduling, Storage, Testing and Windows]
- If a conflict occurs when creating an object with `generateName`, the server now returns an "AlreadyExists" error with a retry option. ([#104699](https://github.com/kubernetes/kubernetes/pull/104699), [@vincepri](https://github.com/vincepri))
- Implement support for recovering from volume expansion failures ([#106154](https://github.com/kubernetes/kubernetes/pull/106154), [@gnufied](https://github.com/gnufied)) [SIG API Machinery, Apps and Storage]
- In kubelet, log verbosity and flush frequency can also be configured via the configuration file and not just via command line flags. In other commands (kube-apiserver, kube-controller-manager), the flags are listed in the "Logs flags" group and not under "Global" or "Misc". The type for `-vmodule` was made a bit more descriptive (`pattern=N,...` instead of `moduleSpec`). ([#106090](https://github.com/kubernetes/kubernetes/pull/106090), [@pohly](https://github.com/pohly)) [SIG API Machinery, Architecture, CLI, Cluster Lifecycle, Instrumentation, Node and Scheduling]
- Introduce `OS` field in the PodSpec ([#104693](https://github.com/kubernetes/kubernetes/pull/104693), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Introduce `v1beta3` API for scheduler. This version
  - increases the weight of user specifiable priorities.
    The weights of following priority plugins are increased
    - `TaintTolerations` to 3 - as leveraging node tainting to group nodes in the cluster is becoming a widely-adopted practice
    - `NodeAffinity` to 2
    - `InterPodAffinity` to 2

  - Won't have `HealthzBindAddress`, `MetricsBindAddress` fields ([#104251](https://github.com/kubernetes/kubernetes/pull/104251), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Introduce v1beta2 for Priority and Fairness with no changes in API spec. ([#104399](https://github.com/kubernetes/kubernetes/pull/104399), [@tkashem](https://github.com/tkashem))
- JSON log output is configurable and now supports writing info messages to stdout and error messages to stderr. Info messages can be buffered in memory. The default is to write both to stdout without buffering, as before. ([#104873](https://github.com/kubernetes/kubernetes/pull/104873), [@pohly](https://github.com/pohly))
- JobTrackingWithFinalizers graduates to beta. Feature is enabled by default. ([#105687](https://github.com/kubernetes/kubernetes/pull/105687), [@alculquicondor](https://github.com/alculquicondor))
- Kube-apiserver: Fixes handling of CRD schemas containing literal null values in enums. ([#104969](https://github.com/kubernetes/kubernetes/pull/104969), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: The `rbac.authorization.k8s.io/v1alpha1` API version is removed; use the `rbac.authorization.k8s.io/v1` API, available since v1.8. The `scheduling.k8s.io/v1alpha1` API version is removed; use the `scheduling.k8s.io/v1` API, available since v1.14. ([#104248](https://github.com/kubernetes/kubernetes/pull/104248), [@liggitt](https://github.com/liggitt))
- Kube-scheduler: support for configuration file version `v1beta1` is removed. Update configuration files to v1beta2(xref: https://github.com/kubernetes/enhancements/issues/2901) or v1beta3 before upgrading to 1.23. ([#104782](https://github.com/kubernetes/kubernetes/pull/104782), [@kerthcet](https://github.com/kerthcet))
- KubeSchedulerConfiguration provides a new field `MultiPoint` which will register a plugin for all valid extension points ([#105611](https://github.com/kubernetes/kubernetes/pull/105611), [@damemi](https://github.com/damemi)) [SIG Scheduling and Testing]
- Kubelet should reject pods whose OS doesn't match the node's OS label. ([#105292](https://github.com/kubernetes/kubernetes/pull/105292), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Apps and Node]
- Kubelet: turn the KubeletConfiguration v1beta1 `ResolverConfig` field from a `string` to `*string`. ([#104624](https://github.com/kubernetes/kubernetes/pull/104624), [@Haleygo](https://github.com/Haleygo))
- Kubernetes is now built using go 1.17. ([#103692](https://github.com/kubernetes/kubernetes/pull/103692), [@justaugustus](https://github.com/justaugustus))
- Performs strict server side schema validation requests via the `fieldValidation=[Strict,Warn,Ignore]`. ([#105916](https://github.com/kubernetes/kubernetes/pull/105916), [@kevindelgado](https://github.com/kevindelgado))
- Promote `IPv6DualStack` feature to stable.
  Controller Manager flags for the node IPAM controller have slightly changed:
  1. When configuring a dual-stack cluster, the user must specify both `--node-cidr-mask-size-ipv4` and `--node-cidr-mask-size-ipv6` to set the per-node IP mask sizes, instead of the previous `--node-cidr-mask-size` flag.
  2. The `--node-cidr-mask-size` flag is mutually exclusive with `--node-cidr-mask-size-ipv4` and `--node-cidr-mask-size-ipv6`.
  3. Single-stack clusters do not need to change, but may choose to use the more specific flags.  Users can use either the older `--node-cidr-mask-size` flag or one of the newer `--node-cidr-mask-size-ipv4` or `--node-cidr-mask-size-ipv6` flags to configure the per-node IP mask size, provided that the flag's IP family matches the cluster's IP family (--cluster-cidr). ([#104691](https://github.com/kubernetes/kubernetes/pull/104691), [@khenidak](https://github.com/khenidak))
- Remove `NodeLease` feature gate that was graduated and locked to stable in 1.17 release. ([#105222](https://github.com/kubernetes/kubernetes/pull/105222), [@cyclinder](https://github.com/cyclinder))
- Removed deprecated `--seccomp-profile-root`/`seccompProfileRoot` config. ([#103941](https://github.com/kubernetes/kubernetes/pull/103941), [@saschagrunert](https://github.com/saschagrunert))
- Since golang 1.17 both net.ParseIP and net.ParseCIDR rejects leading zeros in the dot-decimal notation of IPv4 addresses,
  Kubernetes will keep allowing leading zeros on IPv4 address to not break the compatibility.
  IMPORTANT: Kubernetes interprets leading zeros on IPv4 addresses as decimal, users must not rely on parser alignment to not being impacted by the associated security advisory:
  CVE-2021-29923 golang standard library "net" - Improper Input Validation of octal literals in golang 1.16.2 and below standard library "net" results in indeterminate SSRF & RFI vulnerabilities.
  Reference: https://nvd.nist.gov/vuln/detail/CVE-2021-29923 ([#104368](https://github.com/kubernetes/kubernetes/pull/104368), [@aojea](https://github.com/aojea))
- StatefulSet `minReadySeconds` is promoted to beta. ([#104045](https://github.com/kubernetes/kubernetes/pull/104045), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- Support pod priority based node graceful shutdown. ([#102915](https://github.com/kubernetes/kubernetes/pull/102915), [@wzshiming](https://github.com/wzshiming))
- The "Generic Ephemeral Volume" feature graduates to GA. It is now enabled unconditionally. ([#105609](https://github.com/kubernetes/kubernetes/pull/105609), [@pohly](https://github.com/pohly))
- The Kubelet's `--register-with-taints` option is now available via the Kubelet config file field registerWithTaints ([#105437](https://github.com/kubernetes/kubernetes/pull/105437), [@cmssczy](https://github.com/cmssczy)) [SIG Node and Scalability]
- The `CSIDriver.Spec.StorageCapacity` can now be modified. ([#101789](https://github.com/kubernetes/kubernetes/pull/101789), [@pohly](https://github.com/pohly))
- The `CSIVolumeFSGroupPolicy` feature has moved from beta to GA. ([#105940](https://github.com/kubernetes/kubernetes/pull/105940), [@dobsonj](https://github.com/dobsonj))
- The `IngressClass.Spec.Parameters.Namespace` field is now GA. ([#104636](https://github.com/kubernetes/kubernetes/pull/104636), [@hbagdi](https://github.com/hbagdi))
- The `Service.spec.ipFamilyPolicy` field is now *required* in order to create or update a Service as dual-stack.  This is a breaking change from the beta behavior.  Previously the server would try to infer the value of that field from either `ipFamilies` or `clusterIPs`, but that caused ambiguity on updates.  Users who want a dual-stack Service MUST specify `ipFamilyPolicy` as either "PreferDualStack" or "RequireDualStack". ([#96684](https://github.com/kubernetes/kubernetes/pull/96684), [@thockin](https://github.com/thockin))
- The `TTLAfterFinished` feature gate is now GA and enabled by default. ([#105219](https://github.com/kubernetes/kubernetes/pull/105219), [@sahilvv](https://github.com/sahilvv))
- The `kube-controller-manager` supports `--concurrent-ephemeralvolume-syncs` flag to set the number of ephemeral volume controller workers. ([#102981](https://github.com/kubernetes/kubernetes/pull/102981), [@SataQiu](https://github.com/SataQiu))
- The legacy scheduler policy config is removed in v1.23, the associated flags `policy-config-file`, `policy-configmap`, `policy-configmap-namespace` and `use-legacy-policy-config` are also removed. Migrate to Component Config instead, see https://kubernetes.io/docs/reference/scheduling/config/ for details. ([#105424](https://github.com/kubernetes/kubernetes/pull/105424), [@kerthcet](https://github.com/kerthcet))
- Track the number of Pods with a Ready condition in Job status. The feature is alpha and needs the feature gate JobReadyPods to be enabled. ([#104915](https://github.com/kubernetes/kubernetes/pull/104915), [@alculquicondor](https://github.com/alculquicondor))
- Users of `LogFormatRegistry` in component-base must update their code to use the logr v1.0.0 API. The JSON log output now uses the format from go-logr/zapr (no `v` field for error messages, additional information for invalid calls) and has some fixes (correct source code location for warnings about invalid log calls). ([#104103](https://github.com/kubernetes/kubernetes/pull/104103), [@pohly](https://github.com/pohly))
- Validation rules for Custom Resource Definitions can be written in the [CEL expression language](https://github.com/google/cel-spec) using the `x-kubernetes-validations` extension in OpenAPIv3 schemas (alpha). This is gated by the alpha "CustomResourceValidationExpressions" feature gate. ([#106051](https://github.com/kubernetes/kubernetes/pull/106051), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Fix OpenAPI serialization of the x-kubernetes-validations field ([#108030](https://github.com/kubernetes/kubernetes/pull/108030), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Fixes a regression in v1beta1 PodDisruptionBudget handling of "strategic merge patch"-type API requests for the `selector` field. Prior to 1.21, these requests would merge `matchLabels` content and replace `matchExpressions` content. In 1.21, patch requests touching the `selector` field started replacing the entire selector. This is consistent with server-side apply and the v1 PodDisruptionBudget behavior, but should not have been changed for v1beta1. ([#108139](https://github.com/kubernetes/kubernetes/pull/108139), [@liggitt](https://github.com/liggitt)) [SIG Auth and Testing]
- Omits alpha-level enums from the static openapi file captured in api/openapi-spec ([#109179](https://github.com/kubernetes/kubernetes/pull/109179), [@liggitt](https://github.com/liggitt)) [SIG Apps and Auth]
- Sets JobTrackingWithFinalizers, beta feature, as disabled by default, due to unresolved bug https://github.com/kubernetes/kubernetes/issues/109485 ([#109491](https://github.com/kubernetes/kubernetes/pull/109491), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps, Auth, CLI, Network, Node, Scheduling, Storage and Testing]

#### Feature

- (beta feature) If the CSI driver supports the NodeServiceCapability `VOLUME_MOUNT_GROUP` and the `DelegateFSGroupToCSIDriver` feature gate is enabled, kubelet will delegate applying FSGroup to the driver by passing it to NodeStageVolume and NodePublishVolume, regardless of what other FSGroup policies are set. ([#106330](https://github.com/kubernetes/kubernetes/pull/106330), [@verult](https://github.com/verult)) [SIG Storage]
- Add a new `distribute-cpus-across-numa` option to the static `CPUManager` policy. When enabled, this will trigger the `CPUManager` to evenly distribute CPUs across NUMA nodes in cases where more than one NUMA node is required to satisfy the allocation. ([#105631](https://github.com/kubernetes/kubernetes/pull/105631), [@klueska](https://github.com/klueska))
- Add fish shell completion to kubectl. ([#92989](https://github.com/kubernetes/kubernetes/pull/92989), [@WLun001](https://github.com/WLun001))
- Add mechanism to load simple sniffer class into fluentd-elasticsearch image ([#92853](https://github.com/kubernetes/kubernetes/pull/92853), [@cosmo0920](https://github.com/cosmo0920))
- Add support for Portworx plugin to csi-translation-lib. Alpha release

  Portworx CSI driver is required to enable migration.
  This PR adds support of the `CSIMigrationPortworx` feature gate, which can be enabled by:

  1. Adding the feature flag to the kube-controller-manager `--feature-gates=CSIMigrationPortworx=true`
  2. Adding the feature flag to the kubelet config:

  featureGates:
  CSIMigrationPortworx: true ([#103447](https://github.com/kubernetes/kubernetes/pull/103447), [@trierra](https://github.com/trierra)) [SIG API Machinery, Apps, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Release, Scalability, Scheduling, Storage, Testing and Windows]
- Add support to generate client-side binaries for windows/arm64 platform ([#104894](https://github.com/kubernetes/kubernetes/pull/104894), [@pacoxu](https://github.com/pacoxu))
- Added PowerShell completion generation by running `kubectl completion powershell`. ([#103758](https://github.com/kubernetes/kubernetes/pull/103758), [@zikhan](https://github.com/zikhan))
- Added a `Processing` condition for the `workqueue` API.
  Changed `Shutdown` for the `workqueue` API to wait until the work queue finishes processing all in-flight items. ([#101928](https://github.com/kubernetes/kubernetes/pull/101928), [@alexanderConstantinescu](https://github.com/alexanderConstantinescu))
- Added a new feature gate `CustomResourceValidationExpressions` to enable expression validation for Custom Resource. ([#105107](https://github.com/kubernetes/kubernetes/pull/105107), [@cici37](https://github.com/cici37))
- Added a new flag `--append-server-path` to `kubectl proxy` that will automatically append the kube context server path to each request. ([#97350](https://github.com/kubernetes/kubernetes/pull/97350), [@FabianKramm](https://github.com/FabianKramm))
- Added ability for `kubectl wait` to wait on arbitary JSON path ([#105776](https://github.com/kubernetes/kubernetes/pull/105776), [@lauchokyip](https://github.com/lauchokyip))
- Added support for `PodAndContainerStatsFromCRI` feature gate, which allows a user to specify their pod stats must also come from the CRI, not `cAdvisor`. ([#103095](https://github.com/kubernetes/kubernetes/pull/103095), [@haircommander](https://github.com/haircommander))
- Added support for setting controller-manager log level online. ([#104571](https://github.com/kubernetes/kubernetes/pull/104571), [@h4ghhh](https://github.com/h4ghhh))
- Added the ability to specify whether to use an RFC7396 JSON Merge Patch, an RFC6902 JSON Patch, or a Strategic Merge Patch to perform an override of the resources created by `kubectl run` and `kubectl expose`. ([#105140](https://github.com/kubernetes/kubernetes/pull/105140), [@brianpursley](https://github.com/brianpursley))
- Adding option for `kubectl cp` to resume on network errors until completion, requires tar in addition to tail inside the container image ([#104792](https://github.com/kubernetes/kubernetes/pull/104792), [@matthyx](https://github.com/matthyx))
- Adding support for multiple `--from-env-file flags`. ([#104232](https://github.com/kubernetes/kubernetes/pull/104232), [@lauchokyip](https://github.com/lauchokyip))
- Adding support for multiple `--from-env-file` flags. ([#101646](https://github.com/kubernetes/kubernetes/pull/101646), [@lauchokyip](https://github.com/lauchokyip))
- Adds `--as-uid` flag to `kubectl` to allow uid impersonation in the same way as user and group impersonation. ([#105794](https://github.com/kubernetes/kubernetes/pull/105794), [@margocrawf](https://github.com/margocrawf))
- Adds new [alpha] command 'kubectl events' ([#99557](https://github.com/kubernetes/kubernetes/pull/99557), [@bboreham](https://github.com/bboreham))
- Allow node expansion of local volumes. ([#102886](https://github.com/kubernetes/kubernetes/pull/102886), [@gnufied](https://github.com/gnufied))
- Allow to build kubernetes with a custom `kube-cross` image. ([#104185](https://github.com/kubernetes/kubernetes/pull/104185), [@dims](https://github.com/dims))
- Allows users to prevent garbage collection on pinned images ([#103299](https://github.com/kubernetes/kubernetes/pull/103299), [@wgahnagl](https://github.com/wgahnagl)) [SIG Node]
- CRI v1 is now the project default. If a container runtime does not support the v1 API, Kubernetes will fall back to the v1alpha2 implementation. ([#106501](https://github.com/kubernetes/kubernetes/pull/106501), [@ehashman](https://github.com/ehashman))
- Changed feature `CSIMigrationAWS` to on by default. This feature requires the AWS EBS CSI driver to be installed. ([#106098](https://github.com/kubernetes/kubernetes/pull/106098), [@wongma7](https://github.com/wongma7))
- Client-go: pass `DeleteOptions` down to the fake client `Reactor` ([#102945](https://github.com/kubernetes/kubernetes/pull/102945), [@chenchun](https://github.com/chenchun))
- Cloud providers can set service account names for cloud controllers. ([#103178](https://github.com/kubernetes/kubernetes/pull/103178), [@nckturner](https://github.com/nckturner))
- Display Labels when kubectl describe ingress. ([#103894](https://github.com/kubernetes/kubernetes/pull/103894), [@kabab](https://github.com/kabab))
- Enhance scheduler `VolumeBinding` plugin to handle Lost PVC as `UnschedulableAndUnresolvable` ([#105245](https://github.com/kubernetes/kubernetes/pull/105245), [@yibozhuang](https://github.com/yibozhuang))
- Ensures that volume is deleted from the storage backend when the user tries to delete the PV object manually and the PV `ReclaimPolicy` is set to `Delete`. ([#105773](https://github.com/kubernetes/kubernetes/pull/105773), [@deepakkinni](https://github.com/deepakkinni))
- Expose a `NewUnstructuredExtractor` from apply configurations `meta/v1` package that enables extracting objects into unstructured apply configurations. ([#103564](https://github.com/kubernetes/kubernetes/pull/103564), [@kevindelgado](https://github.com/kevindelgado))
- Feature gate `StorageObjectInUseProtection` has been deprecated and cannot be disabled. It will be completely removed in 1.25 ([#105495](https://github.com/kubernetes/kubernetes/pull/105495), [@ikeeip](https://github.com/ikeeip))
- Graduating `controller_admission_duration_seconds`, `step_admission_duration_seconds`, `webhook_admission_duration_seconds`, `apiserver_current_inflight_requests` and `apiserver_response_sizes` metrics to stable. ([#106122](https://github.com/kubernetes/kubernetes/pull/106122), [@rezakrimi](https://github.com/rezakrimi)) [SIG API Machinery, Instrumentation and Testing]
- Graduating `pending_pods`, `preemption_attempts_total`, `preemption_victims` and `schedule_attempts_total` metrics to stable. Also `e2e_scheduling_duration_seconds` is renamed to `scheduling_attempt_duration_seconds` and the latter is graduated to stable. ([#105941](https://github.com/kubernetes/kubernetes/pull/105941), [@rezakrimi](https://github.com/rezakrimi)) [SIG Instrumentation, Scheduling and Testing]
- Health check of kube-controller-manager now includes each controller. ([#104667](https://github.com/kubernetes/kubernetes/pull/104667), [@jiahuif](https://github.com/jiahuif))
- Integration testing now takes periodic Prometheus scrapes from the etcd server.
  There is a new script ,`hack/run-prometheus-on-etcd-scrapes.sh`, that runs a containerized Prometheus server against an archive of such scrapes. ([#106190](https://github.com/kubernetes/kubernetes/pull/106190), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery and Testing]
- Introduce a feature gate `DisableKubeletCloudCredentialProviders` which allows disabling the in-tree kubelet credential providers.

  The feature gate `DisableKubeletCloudCredentialProviders` is currently in Alpha, which means is currently disabled by default. Once this feature gate moves to beta, in-tree credential providers will be disabled by default, and users will need to migrate to use external credential providers. ([#102507](https://github.com/kubernetes/kubernetes/pull/102507), [@ostrain](https://github.com/ostrain))
- Introduces a new metric: `admission_webhook_request_total` with the following labels: name (string) - the webhook name, type (string) - the admission type, operation (string) - the requested verb, code (int) - the HTTP status code, rejected (bool) - whether the request was rejected, namespace (string) - the namespace of the requested resource. ([#103162](https://github.com/kubernetes/kubernetes/pull/103162), [@rmoriar1](https://github.com/rmoriar1))
- Kubeadm: add support for dry running `kubeadm join`. The new flag `kubeadm join --dry-run` is similar to the existing flag for `kubeadm init/upgrade` and allows you to see what changes would be applied. ([#103027](https://github.com/kubernetes/kubernetes/pull/103027), [@Haleygo](https://github.com/Haleygo))
- Kubeadm: do not check if the `/etc/kubernetes/manifests` folder is empty on joining worker nodes during preflight ([#104942](https://github.com/kubernetes/kubernetes/pull/104942), [@SataQiu](https://github.com/SataQiu))
- Kubectl will now provide shell completion choices for the `--output/-o` flag ([#105851](https://github.com/kubernetes/kubernetes/pull/105851), [@marckhouzam](https://github.com/marckhouzam))
- Kubelet should reconcile `kubernetes.io/os` and `kubernetes.io/arch` labels on the node object. The side-effect of this is kubelet would deny admission to pod which has nodeSelector with label `kubernetes.io/os` or `kubernetes.io/arch` which doesn't match the underlying OS or arch on the host OS.
  - The label reconciliation happens as part of periodic status update which can be configured via flag `--node-status-update-frequency` ([#104613](https://github.com/kubernetes/kubernetes/pull/104613), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node, Testing and Windows]
- Kubernetes is now built with Golang 1.16.7. ([#104199](https://github.com/kubernetes/kubernetes/pull/104199), [@cpanato](https://github.com/cpanato))
- Kubernetes is now built with Golang 1.17.1. ([#104904](https://github.com/kubernetes/kubernetes/pull/104904), [@cpanato](https://github.com/cpanato))
- Kubernetes is now built with Golang 1.17.2 ([#105563](https://github.com/kubernetes/kubernetes/pull/105563), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Kubernetes is now built with Golang 1.17.3 ([#106209](https://github.com/kubernetes/kubernetes/pull/106209), [@cpanato](https://github.com/cpanato)) [SIG API Machinery, Cloud Provider, Instrumentation, Release and Testing]
- Move `ConfigurableFSGroupPolicy` to GA and rename metric `volume_fsgroup_recursive_apply` to `volume_apply_access_control` ([#105885](https://github.com/kubernetes/kubernetes/pull/105885), [@gnufied](https://github.com/gnufied))
- Move the `GetAllocatableResources` Endpoint in PodResource API to the beta that will make it enabled by default. ([#105003](https://github.com/kubernetes/kubernetes/pull/105003), [@swatisehgal](https://github.com/swatisehgal))
- Moving `WindowsHostProcessContainers` feature to beta ([#106058](https://github.com/kubernetes/kubernetes/pull/106058), [@marosset](https://github.com/marosset))
- Node affinity, Node selectors, and tolerations are now mutable for Jobs that are suspended and have never been started ([#105479](https://github.com/kubernetes/kubernetes/pull/105479), [@ahg-g](https://github.com/ahg-g))
- Pod template annotations and labels are now mutable for Jobs that are suspended and have never been started ([#105980](https://github.com/kubernetes/kubernetes/pull/105980), [@ahg-g](https://github.com/ahg-g))
- PodSecurity: in 1.23+ restricted policy levels, Pods and containers which set `runAsUser=0` are forbidden at admission-time; previously, they would be rejected at runtime ([#105857](https://github.com/kubernetes/kubernetes/pull/105857), [@liggitt](https://github.com/liggitt))
- Shell completion now knows to continue suggesting resource names when the command supports it.  For example `kubectl get pod pod1 <TAB>` will suggest more Pod names. ([#105711](https://github.com/kubernetes/kubernetes/pull/105711), [@marckhouzam](https://github.com/marckhouzam))
- Support to enable Hyper-V in GCE Windows Nodes created with `kube-up` ([#105999](https://github.com/kubernetes/kubernetes/pull/105999), [@mauriciopoppe](https://github.com/mauriciopoppe))
- The CPUManager policy options are now enabled, and we introduce a graduation path for the new CPU Manager policy options. ([#105012](https://github.com/kubernetes/kubernetes/pull/105012), [@fromanirh](https://github.com/fromanirh))
- The Pods and Pod controllers that are exempted from the PodSecurity admission process are now marked with the `pod-security.kubernetes.io/exempt: user/namespace/runtimeClass` annotation, based on what caused the exemption.

  The enforcement level that allowed or denied a Pod during PodSecurity admission is now marked by the `pod-security.kubernetes.io/enforce-policy` annotation.

  The annotation that informs about audit policy violations changed from `pod-security.kubernetes.io/audit` to `pod-security.kubernetes.io/audit-violation`. ([#105908](https://github.com/kubernetes/kubernetes/pull/105908), [@stlaz](https://github.com/stlaz))
- The `/openapi/v3` endpoint will be populated with OpenAPI v3 if the feature flag is enabled ([#105945](https://github.com/kubernetes/kubernetes/pull/105945), [@Jefftree](https://github.com/Jefftree))
- The `CSIMigrationGCE` feature flag is turned `ON` by default ([#104722](https://github.com/kubernetes/kubernetes/pull/104722), [@leiyiz](https://github.com/leiyiz))
- The `DownwardAPIHugePages` feature is now enabled by default. ([#106271](https://github.com/kubernetes/kubernetes/pull/106271), [@mysunshine92](https://github.com/mysunshine92))
- The `PodSecurity `admission plugin has graduated to `beta` and is enabled by default. The admission configuration version has been promoted to `pod-security.admission.config.k8s.io/v1beta1`. See https://kubernetes.io/docs/concepts/security/pod-security-admission/ for usage guidelines. ([#106089](https://github.com/kubernetes/kubernetes/pull/106089), [@liggitt](https://github.com/liggitt))
- The `ServiceAccountIssuerDiscovery` feature gate is removed. It reached GA in Kubernetes 1.21. ([#103685](https://github.com/kubernetes/kubernetes/pull/103685), [@mengjiao-liu](https://github.com/mengjiao-liu))
- The `constants/variables` from k8s.io for STABLE metrics is now supported. ([#103654](https://github.com/kubernetes/kubernetes/pull/103654), [@coffeepac](https://github.com/coffeepac))
- The `kubectl describe namespace` now shows Conditions ([#106219](https://github.com/kubernetes/kubernetes/pull/106219), [@dlipovetsky](https://github.com/dlipovetsky))
- The etcd container image now supports Windows. ([#92433](https://github.com/kubernetes/kubernetes/pull/92433), [@claudiubelu](https://github.com/claudiubelu))
- The kube-apiserver's Prometheus metrics have been extended with some that describe the costs of handling LIST requests.  They are as follows.
  - *apiserver_cache_list_total*: Counter of LIST requests served from watch cache, broken down by resource_prefix and index_name
  - *apiserver_cache_list_fetched_objects_total*: Counter of objects read from watch cache in the course of serving a LIST request, broken down by resource_prefix and index_name
  - *apiserver_cache_list_evaluated_objects_total*: Counter of objects tested in the course of serving a LIST request from watch cache, broken down by resource_prefix
  - *apiserver_cache_list_returned_objects_total*: Counter of objects returned for a LIST request from watch cache, broken down by resource_prefix
  - *apiserver_storage_list_total*: Counter of LIST requests served from etcd, broken down by resource
  - *apiserver_storage_list_fetched_objects_total*: Counter of objects read from etcd in the course of serving a LIST request, broken down by resource
  - *apiserver_storage_list_evaluated_objects_total*: Counter of objects tested in the course of serving a LIST request from etcd, broken down by resource
  - *apiserver_storage_list_returned_objects_total*: Counter of objects returned for a LIST request from etcd, broken down by resource ([#104983](https://github.com/kubernetes/kubernetes/pull/104983), [@MikeSpreitzer](https://github.com/MikeSpreitzer))
- The pause image list now contains Windows Server 2022. ([#104438](https://github.com/kubernetes/kubernetes/pull/104438), [@nick5616](https://github.com/nick5616))
- The script `kube-up.sh` installs `csi-proxy v1.0.1-gke.0`. ([#104426](https://github.com/kubernetes/kubernetes/pull/104426), [@mauriciopoppe](https://github.com/mauriciopoppe))
- This PR adds the following metrics for API Priority and Fairness.
  - **apiserver_flowcontrol_priority_level_seat_count_samples**: histograms of seats occupied by executing requests (both regular and final-delay phases included), broken down by priority_level; the observations are taken once per millisecond.
  - **apiserver_flowcontrol_priority_level_seat_count_watermarks**: histograms of high and low watermarks of number of seats occupied by executing requests (both regular and final-delay phases included), broken down by priority_level.
  - **apiserver_flowcontrol_watch_count_samples**: histograms of number of watches relevant to a given mutating request, broken down by that request's priority_level and flow_schema. ([#105873](https://github.com/kubernetes/kubernetes/pull/105873), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery, Instrumentation and Testing]
- Topology Aware Hints have graduated to beta. ([#106433](https://github.com/kubernetes/kubernetes/pull/106433), [@robscott](https://github.com/robscott)) [SIG Network]
- Turn on CSIMigrationAzureDisk by default on 1.23 ([#104670](https://github.com/kubernetes/kubernetes/pull/104670), [@andyzhangx](https://github.com/andyzhangx))
- Update the system-validators library to v1.6.0 ([#106323](https://github.com/kubernetes/kubernetes/pull/106323), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle and Node]
- Updated Cluster Autosaler to version `1.22.0`. Release notes: https://github.com/kubernetes/autoscaler/releases/tag/cluster-autoscaler-1.22.0. ([#104293](https://github.com/kubernetes/kubernetes/pull/104293), [@x13n](https://github.com/x13n))
- Updates `debian-iptables` to v1.6.7 to pick up CVE fixes. ([#104970](https://github.com/kubernetes/kubernetes/pull/104970), [@PushkarJ](https://github.com/PushkarJ))
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.9.0
  - `debian-iptables` to v1.6.6
  - `setcap` to v2.0.4 ([#104142](https://github.com/kubernetes/kubernetes/pull/104142), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Upgrade etcd to 3.5.1 ([#105706](https://github.com/kubernetes/kubernetes/pull/105706), [@uthark](https://github.com/uthark)) [SIG Cloud Provider, Cluster Lifecycle and Testing]
- When feature gate `JobTrackingWithFinalizers` is enabled:
  - Limit the number of Pods tracked in a single Job sync to avoid starvation of small Jobs.
  - The metric `job_pod_finished_total` counts the number of finished Pods tracked by the Job controller. ([#105197](https://github.com/kubernetes/kubernetes/pull/105197), [@alculquicondor](https://github.com/alculquicondor))
- When using `RequestedToCapacityRatio` ScoringStrategy, empty shape will cause error. ([#106169](https://github.com/kubernetes/kubernetes/pull/106169), [@kerthcet](https://github.com/kerthcet)) [SIG Scheduling]
- `client-go` event library allows customizing spam filtering function.
  It is now possible to override `SpamKeyFunc`, which is used by event filtering to detect spam in the events. ([#103918](https://github.com/kubernetes/kubernetes/pull/103918), [@olagacek](https://github.com/olagacek))
- `client-go`, using log level 9, traces the following events of a HTTP request:
  - DNS lookup
  - TCP dialing
  - TLS handshake
  - Time to get a connection from the pool
  - Time to process a request ([#105156](https://github.com/kubernetes/kubernetes/pull/105156), [@aojea](https://github.com/aojea))
- Allow KUBE_TEST_REPO_LIST to be a remote url ([#109512](https://github.com/kubernetes/kubernetes/pull/109512), [@eddiezane](https://github.com/eddiezane)) [SIG Cloud Provider and Testing]
- Kubernetes is now built with Golang 1.17.11 ([#110423](https://github.com/kubernetes/kubernetes/pull/110423), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kube-apiserver: when merging lists, Server Side Apply now prefers the order of the submitted request instead of the existing persisted object ([#107567](https://github.com/kubernetes/kubernetes/pull/107567), [@jiahuif](https://github.com/jiahuif)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Storage and Testing]

#### Documentation

- Graduating `pod_scheduling_duration_seconds`, `pod_scheduling_attempts`, `framework_extension_point_duration_seconds`, `plugin_execution_duration_seconds` and `queue_incoming_pods_total` metrics to stable. ([#106266](https://github.com/kubernetes/kubernetes/pull/106266), [@ahg-g](https://github.com/ahg-g)) [SIG Instrumentation, Scheduling and Testing]
- The test "[sig-network] EndpointSlice should have Endpoints and EndpointSlices pointing to API Server [Conformance]" only requires that there is an EndpointSlice that references the "kubernetes.default" service, it no longer requires that its named "kubernetes". ([#104664](https://github.com/kubernetes/kubernetes/pull/104664), [@aojea](https://github.com/aojea))
- Update description of `--audit-log-maxbackup` to describe behavior when `value = 0`. ([#103843](https://github.com/kubernetes/kubernetes/pull/103843), [@Arkessler](https://github.com/Arkessler))
- Users should not rely on unsupported CRON_TZ variable when specifying schedule, both the API server and cronjob controller will emit warnings pointing to https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/ containing explanation ([#106455](https://github.com/kubernetes/kubernetes/pull/106455), [@soltysh](https://github.com/soltysh)) [SIG Apps]

#### Failing Test

- Fixes hostPath storage E2E tests within SELinux enabled env ([#104551](https://github.com/kubernetes/kubernetes/pull/104551), [@Elbehery](https://github.com/Elbehery))
- Allow KUBE_TEST_REPO_LIST to be a remote url ([#109512](https://github.com/kubernetes/kubernetes/pull/109512), [@eddiezane](https://github.com/eddiezane)) [SIG Cloud Provider and Testing]

#### Bug or Regression

- (PodSecurity admission) errors validating workload resources (deployment, replicaset, etc.) no longer block admission. ([#106017](https://github.com/kubernetes/kubernetes/pull/106017), [@tallclair](https://github.com/tallclair)) [SIG Auth]
- A pod that the Kubelet rejects was still considered as being accepted for a brief period of time after rejection, which might cause some pods to be rejected briefly that could fit on the node.  A pod that is still terminating (but has status indicating it has failed) may also still be consuming resources and so should also be considered. ([#104817](https://github.com/kubernetes/kubernetes/pull/104817), [@smarterclayton](https://github.com/smarterclayton))
- Add Kubernetes Events to the `Kubelet Graceful Shutdown` feature. ([#101081](https://github.com/kubernetes/kubernetes/pull/101081), [@rphillips](https://github.com/rphillips))
- Add Pod Security admission metrics: `pod_security_evaluations_total`, `pod_security_exemptions_total`, `pod_security_errors_total` ([#105898](https://github.com/kubernetes/kubernetes/pull/105898), [@tallclair](https://github.com/tallclair))
- Add support for Windows Network stats in Containerd ([#105744](https://github.com/kubernetes/kubernetes/pull/105744), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node, Testing and Windows]
- Added show-capacity option to `kubectl top node` to show `Capacity` resource usage ([#102917](https://github.com/kubernetes/kubernetes/pull/102917), [@bysnupy](https://github.com/bysnupy)) [SIG CLI]
- Apimachinery: Pretty printed JSON and YAML output is now indented consistently. ([#105466](https://github.com/kubernetes/kubernetes/pull/105466), [@liggitt](https://github.com/liggitt))
- Be able to create a Pod with Generic Ephemeral Volumes as raw block devices. ([#105682](https://github.com/kubernetes/kubernetes/pull/105682), [@pohly](https://github.com/pohly))
- CA, certificate and key bundles for the `generic-apiserver` based servers will be reloaded immediately after the files are changed. ([#104102](https://github.com/kubernetes/kubernetes/pull/104102), [@tnqn](https://github.com/tnqn))
- Change `kubectl diff --invalid-arg` status code from 1 to 2 to match docs ([#105445](https://github.com/kubernetes/kubernetes/pull/105445), [@ardaguclu](https://github.com/ardaguclu))
- Changed kubectl describe to compute age of an event using the `EventSeries.count` and `EventSeries.lastObservedTime`. ([#104482](https://github.com/kubernetes/kubernetes/pull/104482), [@harjas27](https://github.com/harjas27))
- Changes behaviour of kube-proxy start; does not attempt to set specific `sysctl` values (which does not work in recent Kernel versions anymore in non-init namespaces), when the current sysctl values are already set higher. ([#103174](https://github.com/kubernetes/kubernetes/pull/103174), [@Napsty](https://github.com/Napsty))
- Client-go uses the same HTTP client for all the generated groups and versions, allowing to share customized transports for multiple groups versions. ([#105490](https://github.com/kubernetes/kubernetes/pull/105490), [@aojea](https://github.com/aojea))
- Disable aufs module for gce clusters. ([#103831](https://github.com/kubernetes/kubernetes/pull/103831), [@lizhuqi](https://github.com/lizhuqi))
- Do not unmount and mount subpath bind mounts during container creation unless bind mount changes ([#105512](https://github.com/kubernetes/kubernetes/pull/105512), [@gnufied](https://github.com/gnufied)) [SIG Storage]
- Don't prematurely close reflectors in case of slow initialization in watch based manager to fix issues with inability to properly mount secrets/configmaps. ([#104604](https://github.com/kubernetes/kubernetes/pull/104604), [@wojtek-t](https://github.com/wojtek-t))
- Don't use a custom dialer for the kubelet if is not rotating certificates, so we can reuse TCP connections and have only one between the apiserver and the kubelet. If users experiment problems with stale connections  using HTTP1.1, they can force the previous behavior of the kubelet by setting the environment variable DISABLE_HTTP2. ([#104844](https://github.com/kubernetes/kubernetes/pull/104844), [@aojea](https://github.com/aojea)) [SIG API Machinery, Auth and Node]
- EndpointSlice Mirroring controller now cleans up managed EndpointSlices when a Service selector is added ([#105997](https://github.com/kubernetes/kubernetes/pull/105997), [@robscott](https://github.com/robscott)) [SIG Apps, Network and Testing]
- Enhanced event messages for pod failed for exec probe timeout ([#106201](https://github.com/kubernetes/kubernetes/pull/106201), [@yxxhero](https://github.com/yxxhero)) [SIG Node]
- Ensure Pods are removed from the scheduler cache when the scheduler misses deletion events due to transient errors ([#106102](https://github.com/kubernetes/kubernetes/pull/106102), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Ensure `InstanceShutdownByProviderID` return false for creating Azure VMs. ([#104382](https://github.com/kubernetes/kubernetes/pull/104382), [@feiskyer](https://github.com/feiskyer))
- Evicted and other terminated Pods will no longer revert to the `Running` phase. ([#105462](https://github.com/kubernetes/kubernetes/pull/105462), [@ehashman](https://github.com/ehashman))
- Fix `kube-apiserver` metric reporting for the deprecated watch path of `/api/<version>/watch/...`. ([#104161](https://github.com/kubernetes/kubernetes/pull/104161), [@wojtek-t](https://github.com/wojtek-t))
- Fix a regression where the Kubelet failed to exclude already completed pods from calculations about how many resources it was currently using when deciding whether to allow more pods. ([#104577](https://github.com/kubernetes/kubernetes/pull/104577), [@smarterclayton](https://github.com/smarterclayton))
- Fix detach disk issue on deleting vmss node. ([#104572](https://github.com/kubernetes/kubernetes/pull/104572), [@andyzhangx](https://github.com/andyzhangx))
- Fix job controller syncs: In case of conflicts, ensure that the sync happens with the most up to date information. Improves reliability of JobTrackingWithFinalizers. ([#105214](https://github.com/kubernetes/kubernetes/pull/105214), [@alculquicondor](https://github.com/alculquicondor))
- Fix job tracking with finalizers for more than 500 pods, ensuring all finalizers are removed before counting the Pod. ([#104666](https://github.com/kubernetes/kubernetes/pull/104666), [@alculquicondor](https://github.com/alculquicondor))
- Fix pod name of NonIndexed Jobs to not include rogue `-1` substring ([#105676](https://github.com/kubernetes/kubernetes/pull/105676), [@alculquicondor](https://github.com/alculquicondor))
- Fix scoring for `NodeResourcesBalancedAllocation` plugins when nodes have containers with no requests. ([#105845](https://github.com/kubernetes/kubernetes/pull/105845), [@ahmad-diaa](https://github.com/ahmad-diaa))
- Fix system default topology spreading when nodes don't have zone labels. Pods correctly spread by default now. ([#105046](https://github.com/kubernetes/kubernetes/pull/105046), [@alculquicondor](https://github.com/alculquicondor))
- Fix: do not try to delete a LoadBalancer that does not exist ([#105777](https://github.com/kubernetes/kubernetes/pull/105777), [@nilo19](https://github.com/nilo19))
- Fix: ignore non-VMSS error for VMAS nodes in `reconcileBackendPools`. ([#103997](https://github.com/kubernetes/kubernetes/pull/103997), [@nilo19](https://github.com/nilo19))
- Fix: leave the probe path empty for TCP probes ([#105253](https://github.com/kubernetes/kubernetes/pull/105253), [@nilo19](https://github.com/nilo19))
- Fix: remove VMSS and VMSS instances from SLB backend pool only when necessary ([#105839](https://github.com/kubernetes/kubernetes/pull/105839), [@nilo19](https://github.com/nilo19))
- Fix: skip `instance not found` when decoupling VMSSs from LB ([#105666](https://github.com/kubernetes/kubernetes/pull/105666), [@nilo19](https://github.com/nilo19))
- Fix: skip case sensitivity when checking Azure NSG rules. ([#104384](https://github.com/kubernetes/kubernetes/pull/104384), [@feiskyer](https://github.com/feiskyer))
- Fixed a bug that prevents a PersistentVolume that has a PersistentVolumeClaim UID which doesn't exist in local cache but exists in etcd from being updated to the Released phase. ([#105211](https://github.com/kubernetes/kubernetes/pull/105211), [@xiaopingrubyist](https://github.com/xiaopingrubyist))
- Fixed a bug where using `kubectl patch` with `$deleteFromPrimitiveList` on a nonexistent or empty list would add the item to the list ([#105421](https://github.com/kubernetes/kubernetes/pull/105421), [@brianpursley](https://github.com/brianpursley))
- Fixed a bug which could cause webhooks to have an incorrect copy of the old object after an Apply or Update ([#106195](https://github.com/kubernetes/kubernetes/pull/106195), [@alexzielenski](https://github.com/alexzielenski)) [SIG API Machinery]
- Fixed a bug which kubectl would emit duplicate warning messages for flag names that contain an underscore and recommend using a nonexistent flag in some cases. ([#103852](https://github.com/kubernetes/kubernetes/pull/103852), [@brianpursley](https://github.com/brianpursley))
- Fixed a panic in `kubectl` when creating secrets with an improper output type ([#106317](https://github.com/kubernetes/kubernetes/pull/106317), [@lauchokyip](https://github.com/lauchokyip))
- Fixed a regression setting `--audit-log-path=-` to log to stdout in 1.22 pre-release. ([#103875](https://github.com/kubernetes/kubernetes/pull/103875), [@andrewrynhard](https://github.com/andrewrynhard))
- Fixed an issue which didn't append OS's environment variables with the one provided in Credential Provider Config file, which may fail execution of external credential provider binary. See https://github.com/kubernetes/kubernetes/issues/102750. ([#103231](https://github.com/kubernetes/kubernetes/pull/103231), [@n4j](https://github.com/n4j))
- Fixed applying of SELinux labels to CSI volumes on very busy systems (with "error checking for SELinux support: could not get consistent content of /proc/self/mountinfo after 3 attempts") ([#105934](https://github.com/kubernetes/kubernetes/pull/105934), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixed architecture within manifest for non `amd64` etcd images. ([#104116](https://github.com/kubernetes/kubernetes/pull/104116), [@saschagrunert](https://github.com/saschagrunert))
- Fixed architecture within manifest for non `amd64` etcd images. ([#105484](https://github.com/kubernetes/kubernetes/pull/105484), [@saschagrunert](https://github.com/saschagrunert))
- Fixed azure disk translation issue due to lower case `managed` kind. ([#103439](https://github.com/kubernetes/kubernetes/pull/103439), [@andyzhangx](https://github.com/andyzhangx))
- Fixed client IP preservation for NodePort service with protocol SCTP in ipvs. ([#104756](https://github.com/kubernetes/kubernetes/pull/104756), [@tnqn](https://github.com/tnqn))
- Fixed concurrent map access causing panics when logging timed-out API calls. ([#105734](https://github.com/kubernetes/kubernetes/pull/105734), [@marseel](https://github.com/marseel))
- Fixed consolidate logs for `instance not found` error
- Fixed skip `not found` nodes when reconciling LB backend address pools ([#105188](https://github.com/kubernetes/kubernetes/pull/105188), [@nilo19](https://github.com/nilo19))
- Fixed occasional pod cgroup freeze when using cgroup v1 and systemd driver. ([#104528](https://github.com/kubernetes/kubernetes/pull/104528), [@kolyshkin](https://github.com/kolyshkin))
- Fixed the issue where logging output of kube-scheduler configuration files included line breaks and escape characters. The output also attempted to output the configuration file in one section without showing the user a more readable format. ([#106228](https://github.com/kubernetes/kubernetes/pull/106228), [@sanchayanghosh](https://github.com/sanchayanghosh)) [SIG Scheduling]
- Fixes a bug that could result in the EndpointSlice controller unnecessarily updating EndpointSlices associated with a Service that had Topology Aware Hints enabled. ([#105267](https://github.com/kubernetes/kubernetes/pull/105267), [@llhuii](https://github.com/llhuii))
- Fixes a regression that could cause panics in LRU caches in controller-manager, kubelet, kube-apiserver, or client-go. ([#104466](https://github.com/kubernetes/kubernetes/pull/104466), [@stbenjam](https://github.com/stbenjam))
- Fixes an issue where an admission webhook can observe a v1 Pod object that does not have the `defaultMode` field set in the injected service account token volume in kube-api-server. ([#104523](https://github.com/kubernetes/kubernetes/pull/104523), [@liggitt](https://github.com/liggitt))
- Fixes the `should support building a client with a CSR` E2E test to work with clusters configured with short certificate lifetimes ([#105396](https://github.com/kubernetes/kubernetes/pull/105396), [@liggitt](https://github.com/liggitt))
- Graceful node shutdown, allow the actual inhibit delay to be greater than the expected inhibit delay. ([#103137](https://github.com/kubernetes/kubernetes/pull/103137), [@wzshiming](https://github.com/wzshiming))
- Handle Generic Ephemeral Volumes properly in the node limits scheduler filter and the kubelet `hostPath` check. ([#100482](https://github.com/kubernetes/kubernetes/pull/100482), [@pohly](https://github.com/pohly))
- Headless Services with no selector which were created without dual-stack enabled will be defaulted to RequireDualStack instead of PreferDualStack.  This is consistent with such Services which are created with dual-stack enabled. ([#104986](https://github.com/kubernetes/kubernetes/pull/104986), [@thockin](https://github.com/thockin))
- Ignore `not a vmss instance` error for VMAS nodes in `EnsureBackendPoolDeleted`. ([#105185](https://github.com/kubernetes/kubernetes/pull/105185), [@ialidzhikov](https://github.com/ialidzhikov))
- Ignore the case when comparing azure tags in service annotation. ([#104705](https://github.com/kubernetes/kubernetes/pull/104705), [@nilo19](https://github.com/nilo19))
- Ignore the case when updating Azure tags. ([#104593](https://github.com/kubernetes/kubernetes/pull/104593), [@nilo19](https://github.com/nilo19))
- Introduce a new server run option 'shutdown-send-retry-after'. If true the HTTP Server will continue listening until all non longrunning request(s) in flight have been drained, during this window all incoming requests will be rejected with a status code `429` and a 'Retry-After' response header. ([#101257](https://github.com/kubernetes/kubernetes/pull/101257), [@tkashem](https://github.com/tkashem))
- Kube-apiserver: Avoid unnecessary repeated calls to `admission webhooks` that reject an update or delete request. ([#104182](https://github.com/kubernetes/kubernetes/pull/104182), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: Server Side Apply merge order is reverted to match v1.22 behavior until `http://issue.k8s.io/104641` is resolved. ([#106661](https://github.com/kubernetes/kubernetes/pull/106661), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: events created via the `events.k8s.io` API group for cluster-scoped objects are now permitted in the default namespace as well for compatibility with events clients and the `v1` API ([#100125](https://github.com/kubernetes/kubernetes/pull/100125), [@h4ghhh](https://github.com/h4ghhh))
- Kube-apiserver: fix a memory leak when deleting multiple objects with a `deletecollection`. ([#105606](https://github.com/kubernetes/kubernetes/pull/105606), [@sxllwx](https://github.com/sxllwx))
- Kube-proxy health check ports used to listen to `:<port>` for each of the services. This is not needed and opens ports in addresses the cluster user may not have intended. The PR limits listening to all node address which are controlled by `--nodeport-addresses` flag. if no addresses are provided then we default to existing behavior by listening to `:<port>` for each service ([#104742](https://github.com/kubernetes/kubernetes/pull/104742), [@khenidak](https://github.com/khenidak))
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104009](https://github.com/kubernetes/kubernetes/pull/104009), [@aojea](https://github.com/aojea))
- Kube-scheduler now doesn't print any usage message when unknown flag is specified. ([#104503](https://github.com/kubernetes/kubernetes/pull/104503), [@sanposhiho](https://github.com/sanposhiho))
- Kube-up now includes CoreDNS version v1.8.6 ([#106091](https://github.com/kubernetes/kubernetes/pull/106091), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cloud Provider]
- Kubeadm: When adding an etcd peer to an existing cluster, if an error is returned indicating the peer has already been added, this is accepted and a `ListMembers` call is used instead to return the existing cluster. This helps to diminish the exponential backoff when the first AddMember call times out, while still retaining a similar performance when the peer has already been added from a previous call. ([#104134](https://github.com/kubernetes/kubernetes/pull/104134), [@ihgann](https://github.com/ihgann))
- Kubeadm: do not allow empty `--config` paths to be passed to `kubeadm kubeconfig user` ([#105649](https://github.com/kubernetes/kubernetes/pull/105649), [@navist2020](https://github.com/navist2020))
- Kubeadm: fix a bug on Windows worker nodes, where the downloaded KubeletConfiguration from the cluster can contain Linux paths that do not work on Windows and can trip the kubelet binary. ([#105992](https://github.com/kubernetes/kubernetes/pull/105992), [@hwdef](https://github.com/hwdef)) [SIG Cluster Lifecycle and Windows]
- Kubeadm: switch the preflight check (called 'Swap') that verifies if swap is enabled on Linux hosts to report a warning instead of an error. This is related to the graduation of the NodeSwap feature gate in the kubelet to Beta and being enabled by default in 1.23 - allows swap support on Linux hosts. In the next release of kubeadm (1.24) the preflight check will be removed, thus we recommend that you stop using it - e.g. via `--ignore-preflight-errors` or the kubeadm config. ([#104854](https://github.com/kubernetes/kubernetes/pull/104854), [@pacoxu](https://github.com/pacoxu))
- Kubelet did not report `kubelet_volume_stats_*` metrics for Generic Ephemeral Volumes. ([#105569](https://github.com/kubernetes/kubernetes/pull/105569), [@pohly](https://github.com/pohly))
- Kubelet's Node Grace Shutdown will terminate probes when shutting down ([#105215](https://github.com/kubernetes/kubernetes/pull/105215), [@rphillips](https://github.com/rphillips))
- Kubelet: fixes a file descriptor leak in log rotation ([#106382](https://github.com/kubernetes/kubernetes/pull/106382), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Kubelet: the printing of flags at the start of kubelet now uses the final logging configuration. ([#106520](https://github.com/kubernetes/kubernetes/pull/106520), [@pohly](https://github.com/pohly))
- Make the etcd client (used by the API server) retry certain types of errors. The full list of retriable (codes.Unavailable) errors can be found at https://github.com/etcd-io/etcd/blob/main/api/v3rpc/rpctypes/error.go#L72 ([#105069](https://github.com/kubernetes/kubernetes/pull/105069), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric. ([#100720](https://github.com/kubernetes/kubernetes/pull/100720), [@dntosas](https://github.com/dntosas))
- Migrated kubernetes object references (= name + namespace) to structured logging when using JSON as log output format ([#104877](https://github.com/kubernetes/kubernetes/pull/104877), [@pohly](https://github.com/pohly))
- Pass additional flags to subpath mount to avoid flakes in certain conditions. ([#104253](https://github.com/kubernetes/kubernetes/pull/104253), [@mauriciopoppe](https://github.com/mauriciopoppe))
- Pod SecurityContext sysctls name parameter for update requests where the existing object's sysctl contains slashes and kubelet sysctl whitelist support contains slashes. ([#102393](https://github.com/kubernetes/kubernetes/pull/102393), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG Apps, Auth, Node, Storage and Testing]
- Pod will not start when Init container was OOM killed. ([#104650](https://github.com/kubernetes/kubernetes/pull/104650), [@yxxhero](https://github.com/yxxhero)) [SIG Node]
- PodResources interface was changed, now it returns only isolated CPUs ([#97415](https://github.com/kubernetes/kubernetes/pull/97415), [@AlexeyPerevalov](https://github.com/AlexeyPerevalov))
- Provide IPv6 support for internal load balancer. ([#103794](https://github.com/kubernetes/kubernetes/pull/103794), [@nilo19](https://github.com/nilo19))
- Reduce the number of calls to docker for stats via dockershim. For Windows this reduces the latency when calling docker, for Linux this saves cpu cycles. ([#104287](https://github.com/kubernetes/kubernetes/pull/104287), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node and Windows]
- Removed the error message label from the `kubelet_started_pods_errors_total` metric ([#105213](https://github.com/kubernetes/kubernetes/pull/105213), [@yxxhero](https://github.com/yxxhero))
- Resolves a potential issue with GC and NS controllers which may delete objects after getting a 404 response from the server during its startup. This PR ensures that requests to aggregated APIs will get 503, not 404 while the APIServiceRegistrationController hasn't finished its job. ([#104748](https://github.com/kubernetes/kubernetes/pull/104748), [@p0lyn0mial](https://github.com/p0lyn0mial))
- Respect grace period when updating static pods. ([#104743](https://github.com/kubernetes/kubernetes/pull/104743), [@gjkim42](https://github.com/gjkim42)) [SIG Node and Testing]
- Revert building binaries with PIE mode. ([#105352](https://github.com/kubernetes/kubernetes/pull/105352), [@ehashman](https://github.com/ehashman))
- Reverts adding namespace label to admission metrics (and histogram exansion) due to cardinality issues. ([#104033](https://github.com/kubernetes/kubernetes/pull/104033), [@s-urbaniak](https://github.com/s-urbaniak))
- Reverts the CRI API version surfaced by dockershim to `v1alpha2`. ([#106808](https://github.com/kubernetes/kubernetes/pull/106808), [@saschagrunert](https://github.com/saschagrunert))
- Scheduler resource metrics over fractional binary quantities (2.5Gi, 1.1Ki) were incorrectly reported as very small values. ([#103751](https://github.com/kubernetes/kubernetes/pull/103751), [@y-tag](https://github.com/y-tag))
- Support more than 100 disk mounts on Windows ([#105673](https://github.com/kubernetes/kubernetes/pull/105673), [@andyzhangx](https://github.com/andyzhangx))
- Support using negative array index in JSON patch replace operations. ([#105896](https://github.com/kubernetes/kubernetes/pull/105896), [@zqzten](https://github.com/zqzten))
- The `--leader-elect*` CLI args are now honored in scheduler. ([#105915](https://github.com/kubernetes/kubernetes/pull/105915), [@Huang-Wei](https://github.com/Huang-Wei))
- The `--leader-elect*` CLI args are now honored in the scheduler. ([#105712](https://github.com/kubernetes/kubernetes/pull/105712), [@Huang-Wei](https://github.com/Huang-Wei))
- The `client-go` dynamic client sets the header `Content-Type: application/json` by default ([#104327](https://github.com/kubernetes/kubernetes/pull/104327), [@sxllwx](https://github.com/sxllwx))
- The `kube-Proxy` now correctly filters out unready endpoints for Services with Topology. ([#106507](https://github.com/kubernetes/kubernetes/pull/106507), [@robscott](https://github.com/robscott))
- The `pods/binding` subresource now honors `metadata.uid` and `metadata.resourceVersion` ([#105913](https://github.com/kubernetes/kubernetes/pull/105913), [@aholic](https://github.com/aholic))
- The kube-proxy sync_proxy_rules_iptables_total metric now gives
  the correct number of rules, rather than being off by one.
  Fixed multiple iptables proxy regressions introduced in 1.22:
  - When using Services with SessionAffinity, client affinity for an
    endpoint now gets broken when that endpoint becomes non-ready
    (rather than continuing until the endpoint is fully deleted).
  - Traffic to a service IP now starts getting rejected (as opposed to
    merely dropped) as soon as there are no longer any *usable*
    endpoints, rather than waiting until all of the terminating
    endpoints have terminated even when those terminating endpoints
    were not being used.
  - Chains for endpoints that won't be used are no longer output to
    iptables, saving a bit of memory/time/cpu. ([#106030](https://github.com/kubernetes/kubernetes/pull/106030), [@danwinship](https://github.com/danwinship)) [SIG Network]
- Topology Aware Hints now ignores unready endpoints when assigning hints. ([#106510](https://github.com/kubernetes/kubernetes/pull/106510), [@robscott](https://github.com/robscott))
- Topology Hints now excludes control plane notes from capacity calculations. ([#104744](https://github.com/kubernetes/kubernetes/pull/104744), [@robscott](https://github.com/robscott))
- Update Go used to build migrate script in etcd image to v1.16.7. ([#104301](https://github.com/kubernetes/kubernetes/pull/104301), [@serathius](https://github.com/serathius))
- Updated json representation for a conflicted taint to `Key=Effect` when a conflicted taint occurs in kubectl taint. ([#104011](https://github.com/kubernetes/kubernetes/pull/104011), [@manugupt1](https://github.com/manugupt1))
- Upgrades functionality of `kubectl kustomize` as described at
  https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.4.1 ([#106389](https://github.com/kubernetes/kubernetes/pull/106389), [@natasha41575](https://github.com/natasha41575)) [SIG CLI]
- Watch requests that are delegated to aggregated API servers no longer reserve concurrency units (seats) in the API Priority and Fairness dispatcher for their entire duration. ([#105511](https://github.com/kubernetes/kubernetes/pull/105511), [@benluddy](https://github.com/benluddy))
- When a static pod file is deleted and recreated while using a fixed UID, the pod was not properly restarted. ([#104847](https://github.com/kubernetes/kubernetes/pull/104847), [@smarterclayton](https://github.com/smarterclayton))
- XFS-filesystems are now force-formatted (option `-f`) in order to avoid problems being formatted due to detection of magic super-blocks. This aligns with the behaviour of formatting of ext3/4 filesystems. ([#104923](https://github.com/kubernetes/kubernetes/pull/104923), [@davidkarlsen](https://github.com/davidkarlsen))
- `--log-flush-frequency` had no effect in several commands or was missing. Help and warning texts were not always using the right format for a command (`add_dir_header` instead of `add-dir-header`). Fixing this included cleaning up flag handling in component-base/logs: that package no longer adds flags to the global flag sets. Commands which want the klog and `--log-flush-frequency` flags must explicitly call `logs.AddFlags`; the new `cli.Run` does that for commands. That helper function also covers flag normalization and printing of usage and errors in a consistent way (print usage text first if parsing failed, then the error). ([#105076](https://github.com/kubernetes/kubernetes/pull/105076), [@pohly](https://github.com/pohly))
- Kubeadm: allow the "certs check-expiration" command to not require the existence of the cluster CA key (ca.key file) when checking the expiration of managed certificates in kubeconfig files. ([#106931](https://github.com/kubernetes/kubernetes/pull/106931), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: during execution of the "check expiration" command, treat the etcd CA as external if there is a missing etcd CA key file (etcd/ca.key) and perform the proper validation on certificates signed by the etcd CA. Additionally, make sure that the CA for all entries in the output table is included - for both certificates on disk and in kubeconfig files. ([#106926](https://github.com/kubernetes/kubernetes/pull/106926), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubectl: restores `--dry-run`, `--dry-run=true`, and `--dry-run=false` for compatibility with pre-1.23 invocations. ([#107021](https://github.com/kubernetes/kubernetes/pull/107021), [@liggitt](https://github.com/liggitt)) [SIG CLI and Testing]
- Reverts graceful node shutdown to match 1.21 behavior of setting pods that have not yet successfully completed to "Failed" phase if the GracefulNodeShutdown feature is enabled in kubelet. The GracefulNodeShutdown feature is beta and must be explicitly configured via kubelet config to be enabled in 1.21+. This changes 1.22 and 1.23 behavior on node shutdown to match 1.21. If you do not want pods to be marked terminated on node shutdown in 1.22 and 1.23, disable the GracefulNodeShutdown feature. ([#106900](https://github.com/kubernetes/kubernetes/pull/106900), [@bobbypage](https://github.com/bobbypage)) [SIG Node and Testing]
- An inefficient lock in EndpointSlice controller metrics cache has been reworked. Network programming latency may be significantly reduced in certain scenarios, especially in clusters with a large number of Services. ([#107167](https://github.com/kubernetes/kubernetes/pull/107167), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Client-go: fix that paged list calls with ResourceVersionMatch set would fail once paging kicked in. ([#107334](https://github.com/kubernetes/kubernetes/pull/107334), [@fasaxc](https://github.com/fasaxc)) [SIG API Machinery]
- Fix a panic when using invalid output format in kubectl create secret command ([#107347](https://github.com/kubernetes/kubernetes/pull/107347), [@rikatz](https://github.com/rikatz)) [SIG CLI]
- Fix: azuredisk parameter lowercase translation issue ([#107429](https://github.com/kubernetes/kubernetes/pull/107429), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug that a pod's .status.nominatedNodeName is not cleared properly, and thus over-occupied system resources. ([#107109](https://github.com/kubernetes/kubernetes/pull/107109), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Testing]
- Fixes a rare race condition handling requests that timeout ([#107458](https://github.com/kubernetes/kubernetes/pull/107458), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Mount-utils: Detect potential stale file handle ([#106988](https://github.com/kubernetes/kubernetes/pull/106988), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage]
- The feature gate was mentioned as `csiMigrationRBD` where it should have been `CSIMigrationRBD` to be in parity with other migration plugins. This release correct the same and keep it as `CSIMigrationRBD`. Users who have configured this feature gate as `csiMigrationRBD` has to reconfigure the same to `CSIMigrationRBD` from this release. ([#107554](https://github.com/kubernetes/kubernetes/pull/107554), [@humblec](https://github.com/humblec)) [SIG Storage]
- Fix: delete non existing Azure disk issue ([#107406](https://github.com/kubernetes/kubernetes/pull/107406), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixes a regression in 1.23 that incorrectly pruned data from array items of a custom resource that set `x-kubernetes-preserve-unknown-fields: true` ([#107689](https://github.com/kubernetes/kubernetes/pull/107689), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Fix Azurefile volumeid collision issue in csi migration ([#107575](https://github.com/kubernetes/kubernetes/pull/107575), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix e2e test "Services should respect internalTrafficPolicy=Local Pod and Node, to Pod (hostNetwork: true)" ([#107902](https://github.com/kubernetes/kubernetes/pull/107902), [@xueqzhan](https://github.com/xueqzhan)) [SIG Network and Testing]
- Fixes a regression in 1.23 where update requests to previously persisted `Service` objects that have not been modified since 1.19 can be rejected with an incorrect `spec.clusterIPs: Required value` error ([#107875](https://github.com/kubernetes/kubernetes/pull/107875), [@liggitt](https://github.com/liggitt)) [SIG Network and Testing]
- Fixes static pod add and removes restarts in certain cases. ([#107761](https://github.com/kubernetes/kubernetes/pull/107761), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Bump sigs.k8s.io/apiserver-network-proxy/konnectivity-client to v0.0.30, fixing goroutine leaks in kube-apiserver. ([#108438](https://github.com/kubernetes/kubernetes/pull/108438), [@andrewsykim](https://github.com/andrewsykim)) [SIG API Machinery, Auth and Cloud Provider]
- Fix kubectl config flags incorrectly setting burst and discovery limits ([#108401](https://github.com/kubernetes/kubernetes/pull/108401), [@ulucinar](https://github.com/ulucinar)) [SIG CLI]
- Fix static pod restarts in cases where the container is not present. ([#108164](https://github.com/kubernetes/kubernetes/pull/108164), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Fixes a bug where a partial EndpointSlice update could cause node name information to be dropped from endpoints that were not updated. ([#108201](https://github.com/kubernetes/kubernetes/pull/108201), [@robscott](https://github.com/robscott)) [SIG Network]
- Fixes a regression in the kubelet restarting static pods. ([#107931](https://github.com/kubernetes/kubernetes/pull/107931), [@rphillips](https://github.com/rphillips)) [SIG Node and Testing]
- Fixes error handling in a kubectl method used in downstream packages. ([#107938](https://github.com/kubernetes/kubernetes/pull/107938), [@heybronson](https://github.com/heybronson)) [SIG CLI]
- Increase Azure ACR credential provider timeout ([#108209](https://github.com/kubernetes/kubernetes/pull/108209), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Kube-apiserver: removed apf_fd from server logs (added in 1.23.0) which could contain data identifying the requesting user ([#108634](https://github.com/kubernetes/kubernetes/pull/108634), [@jupblb](https://github.com/jupblb)) [SIG API Machinery and Scalability]
- Bug: client-go clientset was not defaulting the user agent, using the default golang agent for all the requests. ([#108791](https://github.com/kubernetes/kubernetes/pull/108791), [@aojea](https://github.com/aojea)) [SIG API Machinery and Instrumentation]
- E2e tests wait for kube-root-ca.crt to be populated in namespaces for use with projected service account tokens, reducing delays starting those test pods and errors in the logs. ([#108860](https://github.com/kubernetes/kubernetes/pull/108860), [@eddiezane](https://github.com/eddiezane)) [SIG Testing]
- Failure to start a container cannot accidentally result in the pod being considered "Succeeded" in the presence of deletion. ([#108882](https://github.com/kubernetes/kubernetes/pull/108882), [@rphillips](https://github.com/rphillips)) [SIG Node]
- Fix indexer bug that resulted in incorrect index updates if number of index values for a given object was changing during update ([#109137](https://github.com/kubernetes/kubernetes/pull/109137), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- Fix the overestimated cost of delegated API requests in kube-apiserver API priority&fairness ([#109216](https://github.com/kubernetes/kubernetes/pull/109216), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- Fixed a regression that could incorrectly reject pods with OutOfCpu errors if they were rapidly scheduled after other pods were reported as complete in the API. The Kubelet now waits to report the phase of a pod as terminal in the API until all running containers are guaranteed to have stopped and no new containers can be started.  Short-lived pods may take slightly longer (~1s) to report Succeeded or Failed after this change. ([#108723](https://github.com/kubernetes/kubernetes/pull/108723), [@bobbypage](https://github.com/bobbypage)) [SIG Apps, Node and Testing]
- Correct event registration for multiple scheduler plugins; this fixes a potential significant delay in re-queueing unschedulable pods. ([#109446](https://github.com/kubernetes/kubernetes/pull/109446), [@ahg-g](https://github.com/ahg-g)) [SIG Scheduling and Testing]
- Existing InTree AzureFile PVs which don't have a secret namespace defined will now work properly after enabling CSI migration - the namespace will be obtained from ClaimRef. ([#108000](https://github.com/kubernetes/kubernetes/pull/108000), [@RomanBednar](https://github.com/RomanBednar)) [SIG Cloud Provider and Storage]
- Fix JobTrackingWithFinalizers that:
  - was declaring a job finished before counting all the created pods in the status
  - was leaving pods with finalizers, blocking pod and job deletions
  - JobTrackingWithFinalizers is still disabled by default. ([#109486](https://github.com/kubernetes/kubernetes/pull/109486), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and Testing]
- Fix a bug that out-of-tree plugin is misplaced when using scheduler v1beta3 config ([#108890](https://github.com/kubernetes/kubernetes/pull/108890), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix kubectl completion zsh to use any command name rather than hardcoded kubectl ([#109235](https://github.com/kubernetes/kubernetes/pull/109235), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Kubeadm: add the flag "--experimental-initial-corrupt-check" to etcd static Pod manifests to ensure etcd member data consistency ([#109075](https://github.com/kubernetes/kubernetes/pull/109075), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- EndpointSlices marked for deletion are now ignored during reconciliation. ([#110483](https://github.com/kubernetes/kubernetes/pull/110483), [@aryan9600](https://github.com/aryan9600)) [SIG Apps and Network]
- Fixed a kubelet issue that could result in invalid pod status updates to be sent to the api-server where pods would be reported in a terminal phase but also report a ready condition of true in some cases. ([#110480](https://github.com/kubernetes/kubernetes/pull/110480), [@bobbypage](https://github.com/bobbypage)) [SIG Node and Testing]
- Pods will now post their readiness during termination. ([#110417](https://github.com/kubernetes/kubernetes/pull/110417), [@aojea](https://github.com/aojea)) [SIG Network, Node and Testing]
- Fix a bug that caused the wrong result length when using --chunk-size and --selector together ([#110757](https://github.com/kubernetes/kubernetes/pull/110757), [@Abirdcfly](https://github.com/Abirdcfly)) [SIG API Machinery and Testing]
- Fix bug that prevented the job controller from enforcing activeDeadlineSeconds when set ([#110545](https://github.com/kubernetes/kubernetes/pull/110545), [@harshanarayana](https://github.com/harshanarayana)) [SIG Apps]
- Fix image pulling failure when IMDS is unavailable in kubelet startup ([#110523](https://github.com/kubernetes/kubernetes/pull/110523), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix printing resources with int64 fields ([#110602](https://github.com/kubernetes/kubernetes/pull/110602), [@sanchezl](https://github.com/sanchezl)) [SIG API Machinery]
- Fixed a regression introduced in 1.23.0 where Azure load balancers were not kept up to date with the state of cluster nodes. In particular, nodes that are not in the ready state and are not newly created (i.e. not having the `node.cloudprovider.kubernetes.io/uninitialized` taint) now get removed from Azure load balancers. ([#109932](https://github.com/kubernetes/kubernetes/pull/109932), [@ricky-rav](https://github.com/ricky-rav)) [SIG Cloud Provider]
- Fixed potential scheduler crash when scheduling with unsatisfied nodes in PodTopologySpread. ([#110853](https://github.com/kubernetes/kubernetes/pull/110853), [@kerthcet](https://github.com/kerthcet)) [SIG Scheduling]
- Kubeadm: fix the bug that configurable KubernetesVersion not respected during kubeadm join ([#111022](https://github.com/kubernetes/kubernetes/pull/111022), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Reduced time taken to sync proxy rules on Windows kube-proxy with kernelspace mode ([#110702](https://github.com/kubernetes/kubernetes/pull/110702), [@daschott](https://github.com/daschott)) [SIG Network and Windows]
- Updated cAdvisor to v0.43.1 to pick up a kubelet fix where network metrics can be missing in some cases when used with containerd ([#111013](https://github.com/kubernetes/kubernetes/pull/111013), [@bobbypage](https://github.com/bobbypage)) [SIG Node]

#### Other (Cleanup or Flake)

- All `klog` flags except for `-v` and `-vmodule` are deprecated. Support for `-vmodule` is only guaranteed for the text log format. ([#105042](https://github.com/kubernetes/kubernetes/pull/105042), [@pohly](https://github.com/pohly))
- Better pod events ("waiting for ephemeral volume controller to create the persistentvolumeclaim"" instead of "persistentvolumeclaim not found") when using generic ephemeral volumes. ([#104605](https://github.com/kubernetes/kubernetes/pull/104605), [@pohly](https://github.com/pohly))
- Changed buckets in apiserver_request_duration_seconds metric from [0.05, 0.1, 0.15, 0.2, 0.25, 0.3, 0.35, 0.4, 0.45, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0,1.25, 1.5, 1.75, 2.0, 2.5, 3.0, 3.5, 4.0, 4.5, 5, 6, 7, 8, 9, 10, 15, 20, 25, 30, 40, 50, 60] to [0.05, 0.1, 0.2, 0.4, 0.6, 0.8, 1.0, 1.25, 1.5, 2, 3, 4, 5, 6, 8, 10, 15, 20, 30, 45, 60] ([#106306](https://github.com/kubernetes/kubernetes/pull/106306), [@pawbana](https://github.com/pawbana)) [SIG API Machinery, Instrumentation and Testing]
- Deprecate `apiserver_longrunning_gauge` and `apiserver_register_watchers` in 1.23.0. ([#103793](https://github.com/kubernetes/kubernetes/pull/103793), [@yan-lgtm](https://github.com/yan-lgtm))
- Enhanced error message for nodes not selected by scheduler due to pod's PersistentVolumeClaim(s) bound to PersistentVolume(s) that do not exist. ([#105196](https://github.com/kubernetes/kubernetes/pull/105196), [@yibozhuang](https://github.com/yibozhuang))
- Fix an issue in cleaning up `CertificateSigningRequest` objects with an unparseable `status.certificate` field. ([#103823](https://github.com/kubernetes/kubernetes/pull/103823), [@liggitt](https://github.com/liggitt))
- Kube-apiserver: requests to node, Service, and Pod `/proxy` subresources with no additional URL path now only automatically redirect GET and HEAD requests. ([#95128](https://github.com/kubernetes/kubernetes/pull/95128), [@Riaankl](https://github.com/Riaankl))
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and the time to read the headers of incoming requests. ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt))
- Kubeadm: external etcd endpoints passed in the `ClusterConfiguration` that have Unicode characters are no longer IDNA encoded (converted to Punycode). They are now just URL encoded as per Go's implementation of RFC-3986, have duplicate "/" removed from the URL paths, and passed like that directly to the `kube-apiserver` `--etcd-servers` flag. If you have etcd endpoints that have Unicode characters, it is advisable to encode them in advance with tooling that is fully IDNA compliant. If you don't do that, the Go standard library (used in k8s and etcd) would do it for you when making requests to the endpoints. ([#103801](https://github.com/kubernetes/kubernetes/pull/103801), [@gkarthiks](https://github.com/gkarthiks))
- Kubeadm: remove the `--port` flag from the manifest for the `kube-controller-manager` since the flag has been a NO-OP since 1.22 and insecure serving was removed for the component. ([#104157](https://github.com/kubernetes/kubernetes/pull/104157), [@knight42](https://github.com/knight42))
- Kubeadm: remove the `--port` flag from the manifest for the kube-scheduler since the flag has been a NO-OP since 1.23 and insecure serving was removed for the component. ([#105034](https://github.com/kubernetes/kubernetes/pull/105034), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: update references to legacy artifacts locations, the `ci-cross` prefix has been removed from the version match as it does not exist in the new `gs://k8s-release-dev` bucket. ([#103813](https://github.com/kubernetes/kubernetes/pull/103813), [@SataQiu](https://github.com/SataQiu))
- Kubectl: deprecated command line flags (like several of the klog flags) now have a `DEPRECATED: <explanation>` comment. ([#106172](https://github.com/kubernetes/kubernetes/pull/106172), [@pohly](https://github.com/pohly)) [SIG CLI]
- Kubemark is now built as a portable, static binary. ([#106150](https://github.com/kubernetes/kubernetes/pull/106150), [@pohly](https://github.com/pohly)) [SIG Scalability and Testing]
- Migrate `cmd/proxy/{config, healthcheck, winkernel}` to structured logging ([#104944](https://github.com/kubernetes/kubernetes/pull/104944), [@jyz0309](https://github.com/jyz0309))
- Migrate `pkg/proxy` to structured logging ([#104908](https://github.com/kubernetes/kubernetes/pull/104908), [@CIPHERTron](https://github.com/CIPHERTron))
- Migrate `pkg/scheduler/framework/plugins/interpodaffinity/filtering.go`,`pkg/scheduler/framework/plugins/podtopologyspread/filtering.go`, `pkg/scheduler/framework/plugins/volumezone/volume_zone.go` to structured logging ([#105931](https://github.com/kubernetes/kubernetes/pull/105931), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Migrate `pkg/scheduler` to structured logging. ([#99273](https://github.com/kubernetes/kubernetes/pull/99273), [@yangjunmyfm192085](https://github.com/yangjunmyfm192085))
- Migrate cmd/proxy/app and pkg/proxy/meta_proxier to structured logging ([#104928](https://github.com/kubernetes/kubernetes/pull/104928), [@jyz0309](https://github.com/jyz0309))
- Migrated `cmd/kube-scheduler/app/server.go`, `pkg/scheduler/framework/plugins/nodelabel/node_label.go`, `pkg/scheduler/framework/plugins/nodevolumelimits/csi.go`, `pkg/scheduler/framework/plugins/nodevolumelimits/non_csi.go` to structured logging ([#105855](https://github.com/kubernetes/kubernetes/pull/105855), [@shivanshu1333](https://github.com/shivanshu1333))
- Migrated `pkg/proxy/ipvs` to structured logging ([#104932](https://github.com/kubernetes/kubernetes/pull/104932), [@shivanshu1333](https://github.com/shivanshu1333))
- Migrated `pkg/proxy/userspace` to structured logging. ([#104931](https://github.com/kubernetes/kubernetes/pull/104931), [@shivanshu1333](https://github.com/shivanshu1333))
- Migrated `pkg/proxy` to structured logging ([#104891](https://github.com/kubernetes/kubernetes/pull/104891), [@shivanshu1333](https://github.com/shivanshu1333))
- Migrated `pkg/scheduler/framework/plugins/volumebinding/assume_cache.go` to structured logging. ([#105904](https://github.com/kubernetes/kubernetes/pull/105904), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG Instrumentation, Scheduling and Storage]
- Migrated `pkg/scheduler/framework/preemption/preemption.go`, `pkg/scheduler/framework/plugins/examples/stateful/stateful.go`, and `pkg/scheduler/framework/plugins/noderesources/resource_allocation.go` to structured logging ([#105967](https://github.com/kubernetes/kubernetes/pull/105967), [@shivanshu1333](https://github.com/shivanshu1333)) [SIG Instrumentation, Node and Scheduling]
- Migrated pkg/proxy/winuserspace to structured logging ([#105035](https://github.com/kubernetes/kubernetes/pull/105035), [@shivanshu1333](https://github.com/shivanshu1333))
- Migrated scheduler file `cache.go` to structured logging ([#105969](https://github.com/kubernetes/kubernetes/pull/105969), [@shivanshu1333](https://github.com/shivanshu1333)) [SIG Instrumentation and Scheduling]
- Migrated scheduler files `comparer.go`, `dumper.go`, `node_tree.go` to structured logging ([#105968](https://github.com/kubernetes/kubernetes/pull/105968), [@shivanshu1333](https://github.com/shivanshu1333)) [SIG Instrumentation and Scheduling]
- More detailed logging has been added to the EndpointSlice controller for Topology. ([#104741](https://github.com/kubernetes/kubernetes/pull/104741), [@robscott](https://github.com/robscott))
- Remove deprecated and not supported old cronjob controller. ([#106126](https://github.com/kubernetes/kubernetes/pull/106126), [@soltysh](https://github.com/soltysh)) [SIG Apps]
- Remove ignore error flag for drain, and set this feature as default ([#105571](https://github.com/kubernetes/kubernetes/pull/105571), [@yuzhiquan](https://github.com/yuzhiquan)) [SIG CLI]
- Remove the deprecated flags `--csr-only` and `--csr-dir` from `kubeadm certs renew`. Please use `kubeadm certs generate-csr` instead. ([#104796](https://github.com/kubernetes/kubernetes/pull/104796), [@RA489](https://github.com/RA489))
- Support allocating whole NUMA nodes in the CPUManager when there is not a 1:1 mapping between socket and NUMA node ([#102015](https://github.com/kubernetes/kubernetes/pull/102015), [@klueska](https://github.com/klueska))
- Support for Windows Server 2022 was added to the `k8s.gcr.io/pause:3.6` image. ([#104711](https://github.com/kubernetes/kubernetes/pull/104711), [@claudiubelu](https://github.com/claudiubelu))
- Surface warning when users don't set `propagationPolicy` for jobs while deleting. ([#104080](https://github.com/kubernetes/kubernetes/pull/104080), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla))
- The `AllowInsecureBackendProxy` feature gate is removed. It reached GA in Kubernetes 1.21. ([#103796](https://github.com/kubernetes/kubernetes/pull/103796), [@mengjiao-liu](https://github.com/mengjiao-liu))
- The `BoundServiceAccountTokenVolume` feature gate that is GA since v1.22 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#104167](https://github.com/kubernetes/kubernetes/pull/104167), [@ialidzhikov](https://github.com/ialidzhikov))
- The `StartupProbe` feature gate that is GA since v1.20 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#104168](https://github.com/kubernetes/kubernetes/pull/104168), [@ialidzhikov](https://github.com/ialidzhikov))
- The `SupportPodPidsLimit` and  `SupportNodePidsLimit` feature gates that are GA since v1.20 are unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#104163](https://github.com/kubernetes/kubernetes/pull/104163), [@ialidzhikov](https://github.com/ialidzhikov))
- The `apiserver` exposes 4 new metrics that allow to track the status of the Service CIDRs allocations:
  - current number of available IPs per Service CIDR
  - current number of used IPs per Service CIDR
  - total number of allocation per Service CIDR
  - total number of allocation errors per ServiceCIDR ([#104119](https://github.com/kubernetes/kubernetes/pull/104119), [@aojea](https://github.com/aojea))
- The flag `--deployment-controller-sync-period` has been deprecated and will be removed in v1.24. ([#103538](https://github.com/kubernetes/kubernetes/pull/103538), [@Pingan2017](https://github.com/Pingan2017))
- The image `gcr.io/kubernetes-e2e-test-images` will no longer be used in E2E / CI testing, `k8s.gcr.io/e2e-test-images` will be used instead. ([#103724](https://github.com/kubernetes/kubernetes/pull/103724), [@claudiubelu](https://github.com/claudiubelu))
- The kube-proxy image contains `/go-runner` as a replacement for deprecated klog flags. ([#106301](https://github.com/kubernetes/kubernetes/pull/106301), [@pohly](https://github.com/pohly))
- The maximum length of the `CSINode` id field has increased to 256 bytes to match the CSI spec. ([#104160](https://github.com/kubernetes/kubernetes/pull/104160), [@pacoxu](https://github.com/pacoxu))
- Troubleshooting: informers log handlers that take more than 100 milliseconds to process an object if the `DeltaFIFO` queue starts to grow beyond 10 elements. ([#103917](https://github.com/kubernetes/kubernetes/pull/103917), [@aojea](https://github.com/aojea))
- Update `cri-tools` dependency to v1.22.0. ([#104430](https://github.com/kubernetes/kubernetes/pull/104430), [@saschagrunert](https://github.com/saschagrunert))
- Update `migratecmd/kube-proxy/app` logs to structured logging. ([#98913](https://github.com/kubernetes/kubernetes/pull/98913), [@yxxhero](https://github.com/yxxhero))
- Update build images to Debian 11 (Bullseye)
  - debian-base:bullseye-v1.0.0
  - debian-iptables:bullseye-v1.0.0
  - go-runner:v2.3.1-go1.17.1-bullseye.0
  - kube-cross:v1.23.0-go1.17.1-bullseye.1
  - setcap:bullseye-v1.0.0
  - cluster/images/etcd: Build 3.5.0-2 image
  - test/conformance/image: Update runner image to base-debian11 ([#105158](https://github.com/kubernetes/kubernetes/pull/105158), [@justaugustus](https://github.com/justaugustus))
- Update conformance image to use `debian-base:buster-v1.9.0`. ([#104696](https://github.com/kubernetes/kubernetes/pull/104696), [@PushkarJ](https://github.com/PushkarJ))
- `volume.kubernetes.io/storage-provisioner` annotation will be added to dynamic provisioning required PVC. `volume.beta.kubernetes.io/storage-provisioner` annotation is deprecated. ([#104590](https://github.com/kubernetes/kubernetes/pull/104590), [@Jiawei0227](https://github.com/Jiawei0227))
- Updates konnectivity-network-proxy to v0.0.27. This includes a memory leak fix for the network proxy ([#107037](https://github.com/kubernetes/kubernetes/pull/107037), [@jdnurme](https://github.com/jdnurme)) [SIG API Machinery, Auth and Cloud Provider]

### Dependencies

#### Added
- bazil.org/fuse: 371fbbd
- github.com/OneOfOne/xxhash: [v1.2.2](https://github.com/OneOfOne/xxhash/tree/v1.2.2)
- github.com/antlr/antlr4/runtime/Go/antlr: [b48c857](https://github.com/antlr/antlr4/runtime/Go/antlr/tree/b48c857)
- github.com/cespare/xxhash: [v1.1.0](https://github.com/cespare/xxhash/tree/v1.1.0)
- github.com/cncf/xds/go: [fbca930](https://github.com/cncf/xds/go/tree/fbca930)
- github.com/getkin/kin-openapi: [v0.76.0](https://github.com/getkin/kin-openapi/tree/v0.76.0)
- github.com/go-logr/zapr: [v1.2.0](https://github.com/go-logr/zapr/tree/v1.2.0)
- github.com/google/cel-go: [v0.9.0](https://github.com/google/cel-go/tree/v0.9.0)
- github.com/google/cel-spec: [v0.6.0](https://github.com/google/cel-spec/tree/v0.6.0)
- github.com/google/martian/v3: [v3.1.0](https://github.com/google/martian/v3/tree/v3.1.0)
- github.com/kr/fs: [v0.1.0](https://github.com/kr/fs/tree/v0.1.0)
- github.com/pkg/sftp: [v1.10.1](https://github.com/pkg/sftp/tree/v1.10.1)
- github.com/spaolacci/murmur3: [f09979e](https://github.com/spaolacci/murmur3/tree/f09979e)
- sigs.k8s.io/json: c049b76

#### Changed
- cloud.google.com/go/bigquery: v1.4.0 → v1.8.0
- cloud.google.com/go/storage: v1.6.0 → v1.10.0
- cloud.google.com/go: v0.54.0 → v0.81.0
- github.com/GoogleCloudPlatform/k8s-cloud-provider: [7901bc8 → ea6160c](https://github.com/GoogleCloudPlatform/k8s-cloud-provider/compare/7901bc8...ea6160c)
- github.com/Microsoft/go-winio: [v0.4.15 → v0.4.17](https://github.com/Microsoft/go-winio/compare/v0.4.15...v0.4.17)
- github.com/Microsoft/hcsshim: [5eafd15 → v0.8.22](https://github.com/Microsoft/hcsshim/compare/5eafd15...v0.8.22)
- github.com/benbjohnson/clock: [v1.0.3 → v1.1.0](https://github.com/benbjohnson/clock/compare/v1.0.3...v1.1.0)
- github.com/bketelsen/crypt: [5cbc8cc → v0.0.4](https://github.com/bketelsen/crypt/compare/5cbc8cc...v0.0.4)
- github.com/containerd/cgroups: [0dbf7f0 → v1.0.1](https://github.com/containerd/cgroups/compare/0dbf7f0...v1.0.1)
- github.com/containerd/containerd: [v1.4.4 → v1.4.11](https://github.com/containerd/containerd/compare/v1.4.4...v1.4.11)
- github.com/containerd/continuity: [aaeac12 → v0.1.0](https://github.com/containerd/continuity/compare/aaeac12...v0.1.0)
- github.com/containerd/fifo: [a9fb20d → v1.0.0](https://github.com/containerd/fifo/compare/a9fb20d...v1.0.0)
- github.com/containerd/go-runc: [5a6d9f3 → v1.0.0](https://github.com/containerd/go-runc/compare/5a6d9f3...v1.0.0)
- github.com/containerd/typeurl: [v1.0.1 → v1.0.2](https://github.com/containerd/typeurl/compare/v1.0.1...v1.0.2)
- github.com/coredns/corefile-migration: [v1.0.12 → v1.0.14](https://github.com/coredns/corefile-migration/compare/v1.0.12...v1.0.14)
- github.com/docker/docker: [v20.10.2+incompatible → v20.10.7+incompatible](https://github.com/docker/docker/compare/v20.10.2...v20.10.7)
- github.com/envoyproxy/go-control-plane: [668b12f → 63b5d3c](https://github.com/envoyproxy/go-control-plane/compare/668b12f...63b5d3c)
- github.com/evanphx/json-patch: [v4.11.0+incompatible → v4.12.0+incompatible](https://github.com/evanphx/json-patch/compare/v4.11.0...v4.12.0)
- github.com/go-logr/logr: [v0.4.0 → v1.2.0](https://github.com/go-logr/logr/compare/v0.4.0...v1.2.0)
- github.com/golang/glog: [23def4e → v1.0.0](https://github.com/golang/glog/compare/23def4e...v1.0.0)
- github.com/golang/mock: [v1.4.4 → v1.5.0](https://github.com/golang/mock/compare/v1.4.4...v1.5.0)
- github.com/google/cadvisor: [v0.39.2 → v0.43.0](https://github.com/google/cadvisor/compare/v0.39.2...v0.43.0)
- github.com/google/pprof: [1ebb73c → cbba55b](https://github.com/google/pprof/compare/1ebb73c...cbba55b)
- github.com/hashicorp/golang-lru: [v0.5.1 → v0.5.0](https://github.com/hashicorp/golang-lru/compare/v0.5.1...v0.5.0)
- github.com/ianlancetaylor/demangle: [5e5cf60 → 28f6c0f](https://github.com/ianlancetaylor/demangle/compare/5e5cf60...28f6c0f)
- github.com/json-iterator/go: [v1.1.11 → v1.1.12](https://github.com/json-iterator/go/compare/v1.1.11...v1.1.12)
- github.com/magiconair/properties: [v1.8.1 → v1.8.5](https://github.com/magiconair/properties/compare/v1.8.1...v1.8.5)
- github.com/mitchellh/go-homedir: [v1.1.0 → v1.0.0](https://github.com/mitchellh/go-homedir/compare/v1.1.0...v1.0.0)
- github.com/mitchellh/mapstructure: [v1.1.2 → v1.4.1](https://github.com/mitchellh/mapstructure/compare/v1.1.2...v1.4.1)
- github.com/modern-go/reflect2: [v1.0.1 → v1.0.2](https://github.com/modern-go/reflect2/compare/v1.0.1...v1.0.2)
- github.com/opencontainers/runc: [v1.0.1 → v1.0.2](https://github.com/opencontainers/runc/compare/v1.0.1...v1.0.2)
- github.com/pelletier/go-toml: [v1.2.0 → v1.9.3](https://github.com/pelletier/go-toml/compare/v1.2.0...v1.9.3)
- github.com/prometheus/common: [v0.26.0 → v0.28.0](https://github.com/prometheus/common/compare/v0.26.0...v0.28.0)
- github.com/spf13/afero: [v1.2.2 → v1.6.0](https://github.com/spf13/afero/compare/v1.2.2...v1.6.0)
- github.com/spf13/cast: [v1.3.0 → v1.3.1](https://github.com/spf13/cast/compare/v1.3.0...v1.3.1)
- github.com/spf13/cobra: [v1.1.3 → v1.2.1](https://github.com/spf13/cobra/compare/v1.1.3...v1.2.1)
- github.com/spf13/jwalterweatherman: [v1.0.0 → v1.1.0](https://github.com/spf13/jwalterweatherman/compare/v1.0.0...v1.1.0)
- github.com/spf13/viper: [v1.7.0 → v1.8.1](https://github.com/spf13/viper/compare/v1.7.0...v1.8.1)
- github.com/yuin/goldmark: [v1.3.5 → v1.4.0](https://github.com/yuin/goldmark/compare/v1.3.5...v1.4.0)
- go.opencensus.io: v0.22.3 → v0.23.0
- go.uber.org/zap: v1.17.0 → v1.19.0
- golang.org/x/crypto: 5ea612d → 32db794
- golang.org/x/net: 37e1c6a → e898025
- golang.org/x/oauth2: bf48bf1 → 2bc19b1
- golang.org/x/sys: 59db8d7 → f4d4317
- golang.org/x/term: 6a3ed07 → 6886f2d
- golang.org/x/text: v0.3.6 → v0.3.7
- golang.org/x/tools: v0.1.2 → d4cc65f
- google.golang.org/api: v0.20.0 → v0.46.0
- google.golang.org/appengine: v1.6.5 → v1.6.7
- google.golang.org/genproto: f16073e → fe13028
- google.golang.org/grpc: v1.38.0 → v1.40.0
- google.golang.org/protobuf: v1.26.0 → v1.27.1
- gopkg.in/ini.v1: v1.51.0 → v1.62.0
- honnef.co/go/tools: v0.0.1-2020.1.3 → v0.0.1-2020.1.4
- k8s.io/gengo: b6c5ce2 → 485abfe
- k8s.io/klog/v2: v2.9.0 → v2.30.0
- k8s.io/kube-openapi: 9528897 → e816edb
- k8s.io/system-validators: v1.5.0 → v1.6.0
- sigs.k8s.io/kustomize/api: v0.8.11 → v0.10.1
- sigs.k8s.io/kustomize/cmd/config: v0.9.13 → v0.10.2
- sigs.k8s.io/kustomize/kustomize/v4: v4.2.0 → v4.4.1
- sigs.k8s.io/kustomize/kyaml: v0.11.0 → v0.13.0
- golang.org/x/net: e898025 → 491a49a
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.25 → v0.0.27
- sigs.k8s.io/structured-merge-diff/v4: v4.1.2 → v4.2.1
- k8s.io/utils: cb0fa31 → 6203023
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.27 → v0.0.30
- github.com/google/cadvisor: [v0.43.0 → v0.43.1](https://github.com/google/cadvisor/compare/v0.43.0...v0.43.1)

#### Removed
- cloud.google.com/go/datastore: v1.1.0
- cloud.google.com/go/pubsub: v1.2.0
- github.com/alecthomas/units: [f65c72e](https://github.com/alecthomas/units/tree/f65c72e)
- github.com/coreos/bbolt: [v1.3.2](https://github.com/coreos/bbolt/tree/v1.3.2)
- github.com/coreos/etcd: [v3.3.13+incompatible](https://github.com/coreos/etcd/tree/v3.3.13)
- github.com/coreos/go-systemd: [95778df](https://github.com/coreos/go-systemd/tree/95778df)
- github.com/coreos/pkg: [399ea9e](https://github.com/coreos/pkg/tree/399ea9e)
- github.com/dgrijalva/jwt-go: [v3.2.0+incompatible](https://github.com/dgrijalva/jwt-go/tree/v3.2.0)
- github.com/google/martian: [v2.1.0+incompatible](https://github.com/google/martian/tree/v2.1.0)
- github.com/jpillora/backoff: [v1.0.0](https://github.com/jpillora/backoff/tree/v1.0.0)
- gotest.tools: v2.2.0+incompatible
