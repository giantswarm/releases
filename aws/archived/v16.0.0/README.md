# :zap: Giant Swarm Release v16.0.0 for AWS :zap:

This release provides support for Kubernetes 1.21.

**Highlights**
- Kubernetes 1.21 support;
- New clusters must be created in the organization's namespaces. Clusters which don't respect this rule will be rejected;
- Tags on custom resources are propagated to S3 buckets and AWS CNI ENIs;
- `aws-operator` and `cluster-operator` have support for Cluster API `v1alpha3` resources;
- Security fixes:
    * 7 Linux CVEs;
    * 2 Systemd CVEs;
    * 2 openssl CVEs;
    * 1 Kubernetes CVE;
    * 1 Go CVE.

> **_Warning:_** Starting with this release new clusters have to be created in the organization's namespace. Any new cluster defined in a different namespace is going to be rejected. Automation that creates Cluster CRs directly needs to be updated to define these clusters in the organization's namespace. Existing clusters are not impacted by this change and will be migrated later.

## Change details


### aws-operator [10.9.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.9.1)

#### Added
- Add cloud tags propagation to S3 buckets.
- Add provider tags to the AWS CNI ENIs.
- Add configuration for `systemd-networkd` to ignore network interfaces used for AWS CNI.
- Add changes to run properly on Flatcar 2905 and newer.

#### Changed
-  Update `aws-attach-etcd-dep` image version to `0.2.0` to include bugfixes.
- Upgrade `k8scloudconfig` which is required for k8s 1.21.
- Introducing `v1alpha3` CR's.
- Update Flatcar AMI's to the latest stable releases.



### containerlinux [2905.2.3](https://www.flatcar-linux.org/releases/#release-2905.2.3)

**Security fixes**

* Linux
    - [CVE-2021-3653](https://nvd.nist.gov/vuln/detail/CVE-2021-3653)
    - [CVE-2021-3656](https://nvd.nist.gov/vuln/detail/CVE-2021-3656)
    - [CVE-2021-38166](https://nvd.nist.gov/vuln/detail/CVE-2021-38166)
    - [CVE-2021-34556](https://nvd.nist.gov/vuln/detail/CVE-2021-34556)
    - [CVE-2021-35477](https://nvd.nist.gov/vuln/detail/CVE-2021-35477)
    - [CVE-2021-38205](https://nvd.nist.gov/vuln/detail/CVE-2021-38205)
    - [CVE-2021-37576](https://nvd.nist.gov/vuln/detail/CVE-2021-37576)
* openssl
    - [CVE-2021-3711](https://nvd.nist.gov/vuln/detail/CVE-2021-3711)
    - [CVE-2021-3712](https://nvd.nist.gov/vuln/detail/CVE-2021-3712)
* Go
    - [CVE-2021-36221](https://nvd.nist.gov/vuln/detail/CVE-2021-36221)
* Systemd
    - [CVE-2020-13529](https://nvd.nist.gov/vuln/detail/CVE-2020-13529)
    - [CVE-2021-33910](https://nvd.nist.gov/vuln/detail/CVE-2021-33910)


**Bug Fixes**

* Re-enabled kernel config FS_ENCRYPTION ([coreos-overlay#1212](https://github.com/kinvolk/coreos-overlay/pull/1212/))
* Fixed Perl in dev-container ([coreos-overlay#1238](https://github.com/kinvolk/coreos-overlay/pull/1238))
* Fixed pam.d sssd LDAP auth with sudo ([coreos-overlay#1170](https://github.com/kinvolk/coreos-overlay/pull/1170))
* Let network-cleanup.service finish before entering rootfs ([coreos-overlay#1182](https://github.com/kinvolk/coreos-overlay/pull/1182))
* Set the cilium_vxlan interface to be not managed by networkd’s default setup with DHCP as it’s managed by Cilium. ([init#43](https://github.com/kinvolk/init/pull/43))
* Disabled SELinux by default on dockerd wrapper script ([coreos-overlay#1149](https://github.com/kinvolk/coreos-overlay/pull/1149))
* GCE: Granted CAP_NET_ADMIN to set routes for the TCP LB when starting oem-gce.service ([coreos-overlay#1146](https://github.com/kinvolk/coreos-overlay/pull/1146))



**Changes**

* Switched to zstd for the initramfs ([coreos-overlay#1136](https://github.com/kinvolk/coreos-overlay/pull/1136))
* Embedded new subkey in flatcar-install ([coreos-overlay#1180](https://github.com/kinvolk/coreos-overlay/pull/1180))



**Updates**

* Linux ([5.10.61](https://lwn.net/Articles/867497/))
* openssl ([1.1.1l](https://mta.openssl.org/pipermail/openssl-announce/2021-August/000206.html))
* Systemd ([247.9](https://github.com/systemd/systemd-stable/releases/tag/v247.9))
* Go ([1.16.7](https://golang.org/doc/devel/release#go1.16.minor))
* portage-utils ([0.90](https://github.com/gentoo/portage-utils/releases/tag/v0.90))



### kubernetes [1.21.5](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.5)

#### What's New (Major Themes)

##### Deprecation of PodSecurityPolicy

PSP as an admission controller resource is being deprecated. Deployed PodSecurityPolicy's will keep working until version 1.25, their target removal from the codebase. A new feature, with a working title of "PSP replacement policy", is being developed in [KEP-2579](https://features.k8s.io/2579). To learn more, read [PodSecurityPolicy Deprecation: Past, Present, and Future](https://blog.k8s.io/2021/04/06/podsecuritypolicy-deprecation-past-present-and-future/).

##### Kubernetes API Reference Documentation

The API reference is now generated with [`gen-resourcesdocs`](https://github.com/kubernetes-sigs/reference-docs/tree/c96658d89fb21037b7d00d27e6dbbe6b32375837/gen-resourcesdocs) and it is moving to [Kubernetes API](https://docs.k8s.io/reference/kubernetes-api/)

##### Kustomize Updates in Kubectl

[Kustomize](https://github.com/kubernetes-sigs/kustomize) version in kubectl had a jump from v2.0.3 to [v4.0.5](https://github.com/kubernetes/kubernetes/pull/98946). Kustomize is now treated as a library and future updates will be less sporadic.

##### Default Container Annotation

Pod with multiple containers can use `kubectl.kubernetes.io/default-container` annotation to have a container preselected for kubectl commands. More can be read in [KEP-2227](https://github.com/kubernetes/enhancements/blob/master/keps/sig-cli/2227-kubectl-default-container/README.md).

##### Immutable Secrets and ConfigMaps

Immutable Secrets and ConfigMaps graduates to GA. This feature allows users to specify that the contents of a particular Secret or ConfigMap is immutable for its object lifetime. For such instances, Kubelet will not watch/poll for changes and therefore reducing apiserver load.

##### Structured Logging in Kubelet

Kubelet has adopted structured logging, thanks to community effort in accomplishing this within the release timeline. Structured logging in the project remains an ongoing effort -- for folks interested in participating, [keep an eye / chime in to the mailing list discussion](https://groups.google.com/g/kubernetes-dev/c/y4WIw-ntUR8).

##### Storage Capacity Tracking

Traditionally, the Kubernetes scheduler was based on the assumptions that additional persistent storage is available everywhere in the cluster and has infinite capacity. Topology constraints addressed the first point, but up to now pod scheduling was still done without considering that the remaining storage capacity may not be enough to start a new pod. [Storage capacity tracking](https://docs.k8s.io/concepts/storage/storage-capacity/) addresses that by adding an API for a CSI driver to report storage capacity and uses that information in the Kubernetes scheduler when choosing a node for a pod. This feature serves as a stepping stone for supporting dynamic provisioning for local volumes and other volume types that are more capacity constrained.

##### Generic Ephemeral Volumes

[Generic ephermeral volumes](https://docs.k8s.io/concepts/storage/ephemeral-volumes/#generic-ephemeral-volumes) feature allows any existing storage driver that supports dynamic provisioning to be used as an ephemeral volume with the volume’s lifecycle bound to the Pod. It can be used to provide scratch storage that is different from the root disk, for example persistent memory, or a separate local disk on that node. All StorageClass parameters for volume provisioning are supported. All features supported with PersistentVolumeClaims are supported, such as storage capacity tracking, snapshots and restore, and volume resizing.

##### CSI Service Account Token

CSI Service Account Token feature moves to Beta in 1.21. This feature improves the security posture and allows CSI drivers to receive pods' [bound service account tokens](https://github.com/kubernetes/enhancements/blob/master/keps/sig-auth/1205-bound-service-account-tokens/README.md). This feature also provides a knob to re-publish volumes so that short-lived volumes can be refreshed.

##### CSI Health Monitoring

The CSI health monitoring feature is being released as a second Alpha in Kubernetes 1.21. This feature enables CSI Drivers to share abnormal volume conditions from the underlying storage systems with Kubernetes so that they can be reported as events on PVCs or Pods. This feature serves as a stepping stone towards programmatic detection and resolution of individual volume health issues by Kubernetes.

#### Important Security Information

This release contains changes that address the following vulnerabilities:

##### CVE-2021-25741: Symlink Exchange Can Allow Host Filesystem Access

A security issue was discovered in Kubernetes where a user may be able to
create a container with subpath volume mounts to access files &
directories outside of the volume, including on the host filesystem.

#### API Change

- "Auto" is now a valid value for the `service.kubernetes.io/topology-aware-hints` annotation. ([#100728](https://github.com/kubernetes/kubernetes/pull/100728), [@robscott](https://github.com/robscott)) [SIG Apps, Instrumentation and Network]
- We have added a new Priority & Fairness rule that exempts all probes (/readyz, /healthz, /livez) to prevent 
  restarting of "healthy" kube-apiserver instance(s) by kubelet. ([#101111](https://github.com/kubernetes/kubernetes/pull/101111), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]

#### Known Issues

##### `TopologyAwareHints` feature falls back to default behavior

The feature gate currently falls back to the default behavior in most cases. Enabling the feature gate will add hints to `EndpointSlices`, but functional differences are only observed in non-dual stack kube-proxy implementation. This is fixed by [#101054](https://github.com/kubernetes/kubernetes/pull/101054).

#### Urgent Upgrade Notes 

##### (No, really, you MUST read this before you upgrade)

- Kube-proxy's IPVS proxy mode no longer sets the net.ipv4.conf.all.route_localnet sysctl parameter. Nodes upgrading will have net.ipv4.conf.all.route_localnet set to 1 but new nodes will inherit the system default (usually 0). If you relied on any behavior requiring net.ipv4.conf.all.route_localnet, you must set ensure it is enabled as kube-proxy will no longer set it automatically. This change helps to further mitigate CVE-2020-8558. ([#92938](https://github.com/kubernetes/kubernetes/pull/92938), [@lbernail](https://github.com/lbernail)) [SIG Network and Release]
 - Kubeadm: during "init" an empty cgroupDriver value in the KubeletConfiguration is now always set to "systemd" unless the user is explicit about it. This requires existing machine setups to configure the container runtime to use the "systemd" driver. Documentation on this topic can be found here: https://kubernetes.io/docs/setup/production-environment/container-runtimes/. When upgrading existing clusters / nodes using "kubeadm upgrade" the old cgroupDriver value is preserved, but in 1.22 this change will also apply to "upgrade". For more information on migrating to the "systemd" driver or remaining on the "cgroupfs" driver see: https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/configure-cgroup-driver/. ([#99471](https://github.com/kubernetes/kubernetes/pull/99471), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
 - Newly provisioned PVs by EBS plugin will no longer use the deprecated "failure-domain.beta.kubernetes.io/zone" and "failure-domain.beta.kubernetes.io/region" labels. It will use "topology.kubernetes.io/zone" and "topology.kubernetes.io/region" labels instead. ([#99130](https://github.com/kubernetes/kubernetes/pull/99130), [@ayberk](https://github.com/ayberk)) [SIG Cloud Provider, Storage and Testing]
 - Newly provisioned PVs by OpenStack Cinder plugin will no longer use the deprecated "failure-domain.beta.kubernetes.io/zone" and "failure-domain.beta.kubernetes.io/region" labels. It will use "topology.kubernetes.io/zone" and "topology.kubernetes.io/region" labels instead. ([#99719](https://github.com/kubernetes/kubernetes/pull/99719), [@jsafrane](https://github.com/jsafrane)) [SIG Cloud Provider and Storage]
 - Newly provisioned PVs by gce-pd will no longer have the beta FailureDomain label. gce-pd volume plugin will start to have GA topology label instead. ([#98700](https://github.com/kubernetes/kubernetes/pull/98700), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Cloud Provider, Storage and Testing]
 - OpenStack Cinder CSI migration is on by default, Clinder CSI driver must be installed on clusters on OpenStack for Cinder volumes to work. ([#98538](https://github.com/kubernetes/kubernetes/pull/98538), [@dims](https://github.com/dims)) [SIG Storage]
 - Remove alpha `CSIMigrationXXComplete` flag and add alpha `InTreePluginXXUnregister` flag. Deprecate `CSIMigrationvSphereComplete` flag and it will be removed in v1.22. ([#98243](https://github.com/kubernetes/kubernetes/pull/98243), [@Jiawei0227](https://github.com/Jiawei0227))
 - Remove storage metrics `storage_operation_errors_total`, since we already have `storage_operation_status_count`.And add new field `status` for `storage_operation_duration_seconds`, so that we can know about all status storage operation latency. ([#98332](https://github.com/kubernetes/kubernetes/pull/98332), [@JornShen](https://github.com/JornShen)) [SIG Instrumentation and Storage]
 - The metric `storage_operation_errors_total` is not removed, but is marked deprecated, and the metric `storage_operation_status_count` is marked deprecated. In both cases the `storage_operation_duration_seconds` metric can be used to recover equivalent counts (using `status=fail-unknown` in the case of `storage_operations_errors_total`). ([#99045](https://github.com/kubernetes/kubernetes/pull/99045), [@mattcary](https://github.com/mattcary))
 - `ServiceNodeExclusion`, `NodeDisruptionExclusion` and `LegacyNodeRoleBehavior` features have been promoted to GA. `ServiceNodeExclusion` and `NodeDisruptionExclusion` are now unconditionally enabled, while `LegacyNodeRoleBehavior` is unconditionally disabled. To prevent control plane nodes from being added to load balancers automatically, upgrade users need to add "node.kubernetes.io/exclude-from-external-load-balancers" label to control plane nodes. ([#97543](https://github.com/kubernetes/kubernetes/pull/97543), [@pacoxu](https://github.com/pacoxu))

#### Deprecation

- Aborting the drain command in a list of nodes will be deprecated. The new behavior will make the drain command go through all nodes even if one or more nodes failed during the drain. For now, users can try such experience by enabling --ignore-errors flag. ([#98203](https://github.com/kubernetes/kubernetes/pull/98203), [@yuzhiquan](https://github.com/yuzhiquan))
- Delete deprecated `service.beta.kubernetes.io/azure-load-balancer-mixed-protocols` mixed procotol annotation in favor of the MixedProtocolLBService feature ([#97096](https://github.com/kubernetes/kubernetes/pull/97096), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Deprecate the `topologyKeys` field in Service. This capability will be replaced with upcoming work around Topology Aware Subsetting and Service Internal Traffic Policy. ([#96736](https://github.com/kubernetes/kubernetes/pull/96736), [@andrewsykim](https://github.com/andrewsykim)) [SIG Apps]
- Kube-proxy: remove deprecated --cleanup-ipvs flag of kube-proxy, and make --cleanup flag always to flush IPVS ([#97336](https://github.com/kubernetes/kubernetes/pull/97336), [@maaoBit](https://github.com/maaoBit)) [SIG Network]
- Kubeadm: deprecated command "alpha selfhosting pivot" is now removed. ([#97627](https://github.com/kubernetes/kubernetes/pull/97627), [@knight42](https://github.com/knight42))
- Kubeadm: graduate the command `kubeadm alpha kubeconfig user` to `kubeadm kubeconfig user`. The `kubeadm alpha kubeconfig user` command is deprecated now. ([#97583](https://github.com/kubernetes/kubernetes/pull/97583), [@knight42](https://github.com/knight42)) [SIG Cluster Lifecycle]
- Kubeadm: the "kubeadm alpha certs" command is removed now, please use "kubeadm certs" instead. ([#97706](https://github.com/kubernetes/kubernetes/pull/97706), [@knight42](https://github.com/knight42)) [SIG Cluster Lifecycle]
- Kubeadm: the deprecated kube-dns is no longer supported as an option. If "ClusterConfiguration.dns.type" is set to "kube-dns" kubeadm will now throw an error. ([#99646](https://github.com/kubernetes/kubernetes/pull/99646), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cluster Lifecycle]
- Kubectl: The deprecated `kubectl alpha debug` command is removed. Use `kubectl debug` instead. ([#98111](https://github.com/kubernetes/kubernetes/pull/98111), [@pandaamanda](https://github.com/pandaamanda)) [SIG CLI]
- Official support to build kubernetes with docker-machine / remote docker is removed. This change does not affect building kubernetes with docker locally. ([#97935](https://github.com/kubernetes/kubernetes/pull/97935), [@adeniyistephen](https://github.com/adeniyistephen)) [SIG Release and Testing]
- Remove deprecated `--generator, --replicas, --service-generator, --service-overrides, --schedule` from `kubectl run`
  Deprecate `--serviceaccount, --hostport, --requests, --limits` in `kubectl run` ([#99732](https://github.com/kubernetes/kubernetes/pull/99732), [@soltysh](https://github.com/soltysh))
- Remove the deprecated metrics "scheduling_algorithm_preemption_evaluation_seconds" and "binding_duration_seconds", suggest to use "scheduler_framework_extension_point_duration_seconds" instead. ([#96447](https://github.com/kubernetes/kubernetes/pull/96447), [@chendave](https://github.com/chendave)) [SIG Cluster Lifecycle, Instrumentation, Scheduling and Testing]
- Removing experimental windows container hyper-v support with Docker ([#97141](https://github.com/kubernetes/kubernetes/pull/97141), [@wawa0210](https://github.com/wawa0210)) [SIG Node and Windows]
- Rename metrics `etcd_object_counts` to `apiserver_storage_object_counts` and mark it as stable. The original `etcd_object_counts` metrics name is marked as "Deprecated" and will be removed in the future. ([#99785](https://github.com/kubernetes/kubernetes/pull/99785), [@erain](https://github.com/erain)) [SIG API Machinery, Instrumentation and Testing]
- The GA TokenRequest and TokenRequestProjection feature gates have been removed and are unconditionally enabled. Remove explicit use of those feature gates in CLI invocations. ([#97148](https://github.com/kubernetes/kubernetes/pull/97148), [@wawa0210](https://github.com/wawa0210)) [SIG Node]
- The PodSecurityPolicy API is deprecated in 1.21, and will no longer be served starting in 1.25. ([#97171](https://github.com/kubernetes/kubernetes/pull/97171), [@deads2k](https://github.com/deads2k)) [SIG Auth and CLI]
- The `batch/v2alpha1` CronJob type definitions and clients are deprecated and removed. ([#96987](https://github.com/kubernetes/kubernetes/pull/96987), [@soltysh](https://github.com/soltysh)) [SIG API Machinery, Apps, CLI and Testing]
- The `export` query parameter (inconsistently supported by API resources and deprecated in v1.14) is fully removed.  Requests setting this query parameter will now receive a 400 status response. ([#98312](https://github.com/kubernetes/kubernetes/pull/98312), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Auth and Testing]
- `audit.k8s.io/v1beta1` and `audit.k8s.io/v1alpha1` audit policy configuration and audit events are deprecated in favor of `audit.k8s.io/v1`, available since v1.13. kube-apiserver invocations that specify alpha or beta policy configurations with `--audit-policy-file`, or explicitly request alpha or beta audit events with `--audit-log-version` / `--audit-webhook-version` must update to use `audit.k8s.io/v1` and accept `audit.k8s.io/v1` events prior to v1.24. ([#98858](https://github.com/kubernetes/kubernetes/pull/98858), [@carlory](https://github.com/carlory)) [SIG Auth]
- `discovery.k8s.io/v1beta1` EndpointSlices are deprecated in favor of `discovery.k8s.io/v1`, and will no longer be served in Kubernetes v1.25. ([#100472](https://github.com/kubernetes/kubernetes/pull/100472), [@liggitt](https://github.com/liggitt))
- `diskformat` storage class parameter for in-tree vSphere volume plugin is deprecated as of v1.21 release. Please consider updating storageclass and remove `diskformat` parameter. vSphere CSI Driver does not support diskformat storageclass parameter.
  
  vSphere releases less than 67u3 are deprecated as of v1.21. Please consider upgrading vSphere to 67u3 or above. vSphere CSI Driver requires minimum vSphere 67u3.
  
  VM Hardware version less than 15 is deprecated as of v1.21. Please consider upgrading the Node VM Hardware version to 15 or above. vSphere CSI Driver recommends Node VM's Hardware version set to at least vmx-15.
  
  Multi vCenter support is deprecated as of v1.21. If you have a Kubernetes cluster spanning across multiple vCenter servers, please consider moving all k8s nodes to a single vCenter Server. vSphere CSI Driver does not support Kubernetes deployment spanning across multiple vCenter servers.
  
  Support for these deprecations will be available till Kubernetes v1.24. ([#98546](https://github.com/kubernetes/kubernetes/pull/98546), [@divyenpatel](https://github.com/divyenpatel))

#### API Change

- 1. PodAffinityTerm includes a namespaceSelector field to allow selecting eligible namespaces based on their labels. 
  2. A new CrossNamespacePodAffinity quota scope API that allows restricting which namespaces allowed to use PodAffinityTerm with corss-namespace reference via namespaceSelector or namespaces fields. ([#98582](https://github.com/kubernetes/kubernetes/pull/98582), [@ahg-g](https://github.com/ahg-g)) [SIG API Machinery, Apps, Auth and Testing]
- Add Probe-level terminationGracePeriodSeconds field ([#99375](https://github.com/kubernetes/kubernetes/pull/99375), [@ehashman](https://github.com/ehashman)) [SIG API Machinery, Apps, Node and Testing]
- Added `.spec.completionMode` field to Job, with accepted values `NonIndexed` (default) and `Indexed`. This is an alpha field and is only honored by servers with the `IndexedJob` feature gate enabled. ([#98441](https://github.com/kubernetes/kubernetes/pull/98441), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and CLI]
- Adds support for endPort field in NetworkPolicy ([#97058](https://github.com/kubernetes/kubernetes/pull/97058), [@rikatz](https://github.com/rikatz)) [SIG Apps and Network]
- CSIServiceAccountToken graduates to Beta and enabled by default. ([#99298](https://github.com/kubernetes/kubernetes/pull/99298), [@zshihang](https://github.com/zshihang))
- Cluster admins can now turn off `/debug/pprof` and `/debug/flags/v` endpoint in kubelet by setting `enableProfilingHandler` and `enableDebugFlagsHandler` to `false` in the Kubelet configuration file. Options `enableProfilingHandler` and `enableDebugFlagsHandler` can be set to `true` only when `enableDebuggingHandlers` is also set to `true`. ([#98458](https://github.com/kubernetes/kubernetes/pull/98458), [@SaranBalaji90](https://github.com/SaranBalaji90))
- DaemonSets accept a MaxSurge integer or percent on their rolling update strategy that will launch the updated pod on nodes and wait for those pods to go ready before marking the old out-of-date pods as deleted. This allows workloads to avoid downtime during upgrades when deployed using DaemonSets. This feature is alpha and is behind the DaemonSetUpdateSurge feature gate. ([#96441](https://github.com/kubernetes/kubernetes/pull/96441), [@smarterclayton](https://github.com/smarterclayton)) [SIG Apps and Testing]
- Enable SPDY pings to keep connections alive, so that `kubectl exec` and `kubectl portforward` won't be interrupted. ([#97083](https://github.com/kubernetes/kubernetes/pull/97083), [@knight42](https://github.com/knight42)) [SIG API Machinery and CLI]
- FieldManager no longer owns fields that get reset before the object is persisted (e.g. "status wiping"). ([#99661](https://github.com/kubernetes/kubernetes/pull/99661), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Auth and Testing]
- Fixes server-side apply for APIService resources. ([#98576](https://github.com/kubernetes/kubernetes/pull/98576), [@kevindelgado](https://github.com/kevindelgado))
- Generic ephemeral volumes are beta. ([#99643](https://github.com/kubernetes/kubernetes/pull/99643), [@pohly](https://github.com/pohly)) [SIG API Machinery, Apps, Auth, CLI, Node, Storage and Testing]
- Hugepages request values are limited to integer multiples of the page size. ([#98515](https://github.com/kubernetes/kubernetes/pull/98515), [@lala123912](https://github.com/lala123912)) [SIG Apps]
- Implement the GetAvailableResources in the podresources API. ([#95734](https://github.com/kubernetes/kubernetes/pull/95734), [@fromanirh](https://github.com/fromanirh)) [SIG Instrumentation, Node and Testing]
- IngressClass resource can now reference a resource in a specific namespace 
  for implementation-specific configuration (previously only Cluster-level resources were allowed). 
  This feature can be enabled using the IngressClassNamespacedParams feature gate. ([#99275](https://github.com/kubernetes/kubernetes/pull/99275), [@hbagdi](https://github.com/hbagdi))
- Jobs API has a new `.spec.suspend` field that can be used to suspend and resume Jobs. This is an alpha field which is only honored by servers with the `SuspendJob` feature gate enabled. ([#98727](https://github.com/kubernetes/kubernetes/pull/98727), [@adtac](https://github.com/adtac))
- Kubelet Graceful Node Shutdown feature graduates to Beta and enabled by default. ([#99735](https://github.com/kubernetes/kubernetes/pull/99735), [@bobbypage](https://github.com/bobbypage))
- Kubernetes is now built using go1.15.7 ([#98363](https://github.com/kubernetes/kubernetes/pull/98363), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Node, Release and Testing]
- Namespace API objects now have a `kubernetes.io/metadata.name` label matching their metadata.name field to allow selecting any namespace by its name using a label selector. ([#96968](https://github.com/kubernetes/kubernetes/pull/96968), [@jayunit100](https://github.com/jayunit100)) [SIG API Machinery, Apps, Cloud Provider, Storage and Testing]
- One new field "InternalTrafficPolicy" in Service is added.
  It specifies if the cluster internal traffic should be routed to all endpoints or node-local endpoints only.
  "Cluster" routes internal traffic to a Service to all endpoints.
  "Local" routes traffic to node-local endpoints only, and traffic is dropped if no node-local endpoints are ready.
  The default value is "Cluster". ([#96600](https://github.com/kubernetes/kubernetes/pull/96600), [@maplain](https://github.com/maplain)) [SIG API Machinery, Apps and Network]
- PodDisruptionBudget API objects can now contain conditions in status. ([#98127](https://github.com/kubernetes/kubernetes/pull/98127), [@mortent](https://github.com/mortent)) [SIG API Machinery, Apps, Auth, CLI, Cloud Provider, Cluster Lifecycle and Instrumentation]
- PodSecurityPolicy only stores "generic" as allowed volume type if the GenericEphemeralVolume feature gate is enabled ([#98918](https://github.com/kubernetes/kubernetes/pull/98918), [@pohly](https://github.com/pohly)) [SIG Auth and Security]
- Promote CronJobs to batch/v1 ([#99423](https://github.com/kubernetes/kubernetes/pull/99423), [@soltysh](https://github.com/soltysh)) [SIG API Machinery, Apps, CLI and Testing]
- Promote Immutable Secrets/ConfigMaps feature to Stable. This allows to set `immutable` field in Secret or ConfigMap object to mark their contents as immutable. ([#97615](https://github.com/kubernetes/kubernetes/pull/97615), [@wojtek-t](https://github.com/wojtek-t)) [SIG Apps, Architecture, Node and Testing]
- Remove support for building Kubernetes with bazel. ([#99561](https://github.com/kubernetes/kubernetes/pull/99561), [@BenTheElder](https://github.com/BenTheElder)) [SIG API Machinery, Apps, Architecture, Auth, Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Release, Scalability, Scheduling, Storage, Testing and Windows]
- Scheduler extender filter interface now can report unresolvable failed nodes in the new field `FailedAndUnresolvableNodes` of  `ExtenderFilterResult` struct. Nodes in this map will be skipped in the preemption phase. ([#92866](https://github.com/kubernetes/kubernetes/pull/92866), [@cofyc](https://github.com/cofyc)) [SIG Scheduling]
- Services can specify loadBalancerClass to use a custom load balancer ([#98277](https://github.com/kubernetes/kubernetes/pull/98277), [@XudongLiuHarold](https://github.com/XudongLiuHarold))
- Storage capacity tracking (= the CSIStorageCapacity feature) graduates to Beta and enabled by default, storage.k8s.io/v1alpha1/VolumeAttachment and storage.k8s.io/v1alpha1/CSIStorageCapacity objects are deprecated ([#99641](https://github.com/kubernetes/kubernetes/pull/99641), [@pohly](https://github.com/pohly))
- Support for Indexed Job: a Job that is considered completed when Pods associated to indexes from 0 to (.spec.completions-1) have succeeded. ([#98812](https://github.com/kubernetes/kubernetes/pull/98812), [@alculquicondor](https://github.com/alculquicondor)) [SIG Apps and CLI]
- The BoundServiceAccountTokenVolume feature has been promoted to beta, and enabled by default.
  - This changes the tokens provided to containers at `/var/run/secrets/kubernetes.io/serviceaccount/token` to be time-limited, auto-refreshed, and invalidated when the containing pod is deleted.
  - Clients should reload the token from disk periodically (once per minute is recommended) to ensure they continue to use a valid token. `k8s.io/client-go` version v11.0.0+ and v0.15.0+ reload tokens automatically.
  - By default, injected tokens are given an extended lifetime so they remain valid even after a new refreshed token is provided. The metric `serviceaccount_stale_tokens_total` can be used to monitor for workloads that are depending on the extended lifetime and are continuing to use tokens even after a refreshed token is provided to the container. If that metric indicates no existing workloads are depending on extended lifetimes, injected token lifetime can be shortened to 1 hour by starting `kube-apiserver` with `--service-account-extend-token-expiration=false`. ([#95667](https://github.com/kubernetes/kubernetes/pull/95667), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Auth, Cluster Lifecycle and Testing]
- The EndpointSlice Controllers are now GA. The `EndpointSliceController` will not populate the `deprecatedTopology` field and will only provide topology information through the `zone` and `nodeName` fields. ([#99870](https://github.com/kubernetes/kubernetes/pull/99870), [@swetharepakula](https://github.com/swetharepakula))
- The Endpoints controller will now set the `endpoints.kubernetes.io/over-capacity` annotation to "warning" when an Endpoints resource contains more than 1000 addresses. In a future release, the controller will truncate Endpoints that exceed this limit. The EndpointSlice API can be used to support significantly larger number of addresses. ([#99975](https://github.com/kubernetes/kubernetes/pull/99975), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- The PodDisruptionBudget API has been promoted to policy/v1 with no schema changes. The only functional change is that an empty selector (`{}`) written to a policy/v1 PodDisruptionBudget now selects all pods in the namespace. The behavior of the policy/v1beta1 API remains unchanged. The policy/v1beta1 PodDisruptionBudget API is deprecated and will no longer be served in 1.25+. ([#99290](https://github.com/kubernetes/kubernetes/pull/99290), [@mortent](https://github.com/mortent)) [SIG API Machinery, Apps, Auth, Autoscaling, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Scheduling and Testing]
- The `EndpointSlice` API is now GA. The `EndpointSlice` topology field has been removed from the GA API and will be replaced by a new per Endpoint Zone field. If the topology field was previously used, it will be converted into an annotation in the v1 Resource. The `discovery.k8s.io/v1alpha1` API is removed. ([#99662](https://github.com/kubernetes/kubernetes/pull/99662), [@swetharepakula](https://github.com/swetharepakula))
- The `controller.kubernetes.io/pod-deletion-cost` annotation can be set to offer a hint on the cost of deleting a `Pod` compared to other pods belonging to the same ReplicaSet. Pods with lower deletion cost are deleted first. This is an alpha feature. ([#99163](https://github.com/kubernetes/kubernetes/pull/99163), [@ahg-g](https://github.com/ahg-g))
- The kube-apiserver now resets `managedFields` that got corrupted by a mutating admission controller. ([#98074](https://github.com/kubernetes/kubernetes/pull/98074), [@kwiesmueller](https://github.com/kwiesmueller))
- Topology Aware Hints are now available in alpha and can be enabled with the `TopologyAwareHints` feature gate. ([#99522](https://github.com/kubernetes/kubernetes/pull/99522), [@robscott](https://github.com/robscott)) [SIG API Machinery, Apps, Auth, Instrumentation, Network and Testing]
- Users might specify the `kubectl.kubernetes.io/default-exec-container` annotation in a Pod to preselect container for kubectl commands. ([#97099](https://github.com/kubernetes/kubernetes/pull/97099), [@pacoxu](https://github.com/pacoxu)) [SIG CLI]

#### Feature
- Kubernetes is now built with Golang 1.16.8 ([#104906](https://github.com/kubernetes/kubernetes/pull/104906), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update the setcap image to buster-v2.0.1 ([#102377](https://github.com/kubernetes/kubernetes/pull/102377), [@xmudrii](https://github.com/xmudrii)) [SIG Release]
- Add NeedResize function to kubernetes/mount-utils, user can call this function to determine if fs need to be resized ([#101253](https://github.com/kubernetes/kubernetes/pull/101253), [@AndyXiangLi](https://github.com/AndyXiangLi)) [SIG Storage]
- Base image updates to mitigate kube-proxy and etcd container image CVEs
  - debian-base to buster-v1.6.0
  - debian-iptables to buster-v1.6.0 ([#100976](https://github.com/kubernetes/kubernetes/pull/100976), [@jindijamie](https://github.com/jindijamie)) [SIG Release and Testing]
- A client-go metric, rest_client_exec_plugin_call_total, has been added to track total calls to client-go credential plugins. ([#98892](https://github.com/kubernetes/kubernetes/pull/98892), [@ankeesler](https://github.com/ankeesler)) [SIG API Machinery, Auth, Cluster Lifecycle and Instrumentation]
- A new histogram metric to track the time it took to delete a job by the `TTLAfterFinished` controller ([#98676](https://github.com/kubernetes/kubernetes/pull/98676), [@ahg-g](https://github.com/ahg-g))
- AWS cloud provider supports auto-discovering subnets without any `kubernetes.io/cluster/<clusterName>` tags. It also supports additional service annotation `service.beta.kubernetes.io/aws-load-balancer-subnets` to manually configure the subnets. ([#97431](https://github.com/kubernetes/kubernetes/pull/97431), [@kishorj](https://github.com/kishorj))
- Aborting the drain command in a list of nodes will be deprecated. The new behavior will make the drain command go through all nodes even if one or more nodes failed during the drain. For now, users can try such experience by enabling --ignore-errors flag. ([#98203](https://github.com/kubernetes/kubernetes/pull/98203), [@yuzhiquan](https://github.com/yuzhiquan))
- Add --permit-address-sharing flag to `kube-apiserver` to listen with `SO_REUSEADDR`. While allowing to listen on wildcard IPs like 0.0.0.0 and specific IPs in parallel, it avoids waiting for the kernel to release socket in `TIME_WAIT` state, and hence, considerably reducing `kube-apiserver` restart times under certain conditions. ([#93861](https://github.com/kubernetes/kubernetes/pull/93861), [@sttts](https://github.com/sttts))
- Add `csi_operations_seconds` metric on kubelet that exposes CSI operations duration and status for node CSI operations. ([#98979](https://github.com/kubernetes/kubernetes/pull/98979), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Instrumentation and Storage]
- Add `migrated` field into `storage_operation_duration_seconds` metric ([#99050](https://github.com/kubernetes/kubernetes/pull/99050), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps, Instrumentation and Storage]
- Add flag --lease-reuse-duration-seconds for kube-apiserver to config etcd lease reuse duration. ([#97009](https://github.com/kubernetes/kubernetes/pull/97009), [@lingsamuel](https://github.com/lingsamuel)) [SIG API Machinery and Scalability]
- Add metric etcd_lease_object_counts for kube-apiserver to observe max objects attached to a single etcd lease. ([#97480](https://github.com/kubernetes/kubernetes/pull/97480), [@lingsamuel](https://github.com/lingsamuel)) [SIG API Machinery, Instrumentation and Scalability]
- Add support to generate client-side binaries for new darwin/arm64 platform ([#97743](https://github.com/kubernetes/kubernetes/pull/97743), [@dims](https://github.com/dims)) [SIG Release and Testing]
- Added `ephemeral_volume_controller_create[_failures]_total` counters to kube-controller-manager metrics ([#99115](https://github.com/kubernetes/kubernetes/pull/99115), [@pohly](https://github.com/pohly)) [SIG API Machinery, Apps, Cluster Lifecycle, Instrumentation and Storage]
- Added support for installing `arm64` node artifacts. ([#99242](https://github.com/kubernetes/kubernetes/pull/99242), [@liu-cong](https://github.com/liu-cong))
- Adds alpha feature `VolumeCapacityPriority` which makes the scheduler prioritize nodes based on the best matching size of statically provisioned PVs across multiple topologies. ([#96347](https://github.com/kubernetes/kubernetes/pull/96347), [@cofyc](https://github.com/cofyc)) [SIG Apps, Network, Scheduling, Storage and Testing]
- Adds the ability to pass --strict-transport-security-directives to the kube-apiserver to set the HSTS header appropriately.  Be sure you understand the consequences to browsers before setting this field. ([#96502](https://github.com/kubernetes/kubernetes/pull/96502), [@249043822](https://github.com/249043822)) [SIG Auth]
- Adds two new metrics to cronjobs, a histogram to track the time difference when a job is created and the expected time when it should be created, as well as a gauge for the missed schedules of a cronjob ([#99341](https://github.com/kubernetes/kubernetes/pull/99341), [@alaypatel07](https://github.com/alaypatel07))
- Alpha implementation of Kubectl Command Headers: SIG CLI KEP 859 enabled when KUBECTL_COMMAND_HEADERS environment variable set on the client command line. ([#98952](https://github.com/kubernetes/kubernetes/pull/98952), [@seans3](https://github.com/seans3))
- Base-images: Update to debian-iptables:buster-v1.4.0
  - Uses iptables 1.8.5
  - base-images: Update to debian-base:buster-v1.3.0
  - cluster/images/etcd: Build etcd:3.4.13-2 image
    - Uses debian-base:buster-v1.3.0 ([#98401](https://github.com/kubernetes/kubernetes/pull/98401), [@pacoxu](https://github.com/pacoxu)) [SIG Testing]
- CRIContainerLogRotation graduates to GA and unconditionally enabled. ([#99651](https://github.com/kubernetes/kubernetes/pull/99651), [@umohnani8](https://github.com/umohnani8))
- Component owner can configure the allowlist of metric label with flag '--allow-metric-labels'. ([#99385](https://github.com/kubernetes/kubernetes/pull/99385), [@YoyinZyc](https://github.com/YoyinZyc)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Release]
- Component owner can configure the allowlist of metric label with flag '--allow-metric-labels'. ([#99738](https://github.com/kubernetes/kubernetes/pull/99738), [@YoyinZyc](https://github.com/YoyinZyc)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- EmptyDir memory backed volumes are sized as the the minimum of pod allocatable memory on a host and an optional explicit user provided value. ([#100319](https://github.com/kubernetes/kubernetes/pull/100319), [@derekwaynecarr](https://github.com/derekwaynecarr)) [SIG Node]
- Enables Kubelet to check volume condition and log events to corresponding pods. ([#99284](https://github.com/kubernetes/kubernetes/pull/99284), [@fengzixu](https://github.com/fengzixu)) [SIG Apps, Instrumentation, Node and Storage]
- EndpointSliceNodeName graduates to GA and thus will be unconditionally enabled -- NodeName will always be available in the v1beta1 API. ([#99746](https://github.com/kubernetes/kubernetes/pull/99746), [@swetharepakula](https://github.com/swetharepakula))
- Export `NewDebuggingRoundTripper` function and `DebugLevel` options in the k8s.io/client-go/transport package. ([#98324](https://github.com/kubernetes/kubernetes/pull/98324), [@atosatto](https://github.com/atosatto))
- Kube-proxy iptables: new metric sync_proxy_rules_iptables_total that exposes the number of rules programmed per table in each iteration ([#99653](https://github.com/kubernetes/kubernetes/pull/99653), [@aojea](https://github.com/aojea)) [SIG Instrumentation and Network]
- Kube-scheduler now logs plugin scoring summaries at --v=4 ([#99411](https://github.com/kubernetes/kubernetes/pull/99411), [@damemi](https://github.com/damemi)) [SIG Scheduling]
- Kubeadm now includes CoreDNS v1.8.0. ([#96429](https://github.com/kubernetes/kubernetes/pull/96429), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cluster Lifecycle]
- Kubeadm: IPv6DualStack feature gate graduates to Beta and enabled by default ([#99294](https://github.com/kubernetes/kubernetes/pull/99294), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: a warning to user as ipv6 site-local is deprecated ([#99574](https://github.com/kubernetes/kubernetes/pull/99574), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle and Network]
- Kubeadm: add support for certificate chain validation. When using kubeadm in external CA mode, this allows an intermediate CA to be used to sign the certificates. The intermediate CA certificate must be appended to each signed certificate for this to work correctly. ([#97266](https://github.com/kubernetes/kubernetes/pull/97266), [@robbiemcmichael](https://github.com/robbiemcmichael)) [SIG Cluster Lifecycle]
- Kubeadm: amend the node kernel validation to treat CGROUP_PIDS, FAIR_GROUP_SCHED as required and CFS_BANDWIDTH, CGROUP_HUGETLB as optional ([#96378](https://github.com/kubernetes/kubernetes/pull/96378), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle and Node]
- Kubeadm: apply the "node.kubernetes.io/exclude-from-external-load-balancers" label on control plane nodes during "init", "join" and "upgrade" to preserve backwards compatibility with the lagacy LB mode where nodes labeled as "master" where excluded. To opt-out you can remove the label from a node. See #97543 and the linked KEP for more details. ([#98269](https://github.com/kubernetes/kubernetes/pull/98269), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: if the user has customized their image repository via the kubeadm configuration, pass the custom pause image repository and tag to the kubelet via --pod-infra-container-image not only for Docker but for all container runtimes. This flag tells the kubelet that it should not garbage collect the image. ([#99476](https://github.com/kubernetes/kubernetes/pull/99476), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: perform pre-flight validation on host/node name upon `kubeadm init` and `kubeadm join`, showing warnings on non-compliant names ([#99194](https://github.com/kubernetes/kubernetes/pull/99194), [@pacoxu](https://github.com/pacoxu))
- Kubectl version changed to write a warning message to stderr if the client and server version difference exceeds the supported version skew of +/-1 minor version. ([#98250](https://github.com/kubernetes/kubernetes/pull/98250), [@brianpursley](https://github.com/brianpursley)) [SIG CLI]
- Kubectl: Add `--use-protocol-buffers` flag to kubectl top pods and nodes. ([#96655](https://github.com/kubernetes/kubernetes/pull/96655), [@serathius](https://github.com/serathius))
- Kubectl: `kubectl get` will omit managed fields by default now. Users could set `--show-managed-fields` to true to show managedFields when the output format is either `json` or `yaml`. ([#96878](https://github.com/kubernetes/kubernetes/pull/96878), [@knight42](https://github.com/knight42)) [SIG CLI and Testing]
- Kubectl: a Pod can be preselected as default container using `kubectl.kubernetes.io/default-container` annotation ([#99833](https://github.com/kubernetes/kubernetes/pull/99833), [@mengjiao-liu](https://github.com/mengjiao-liu))
- Kubectl: add bash-completion for comma separated list on `kubectl get` ([#98301](https://github.com/kubernetes/kubernetes/pull/98301), [@phil9909](https://github.com/phil9909))
- Kubernetes is now built using go1.15.8 ([#98834](https://github.com/kubernetes/kubernetes/pull/98834), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built with Golang 1.16 ([#98572](https://github.com/kubernetes/kubernetes/pull/98572), [@justaugustus](https://github.com/justaugustus)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Release and Testing]
- Kubernetes is now built with Golang 1.16.1 ([#100106](https://github.com/kubernetes/kubernetes/pull/100106), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Metrics can now be disabled explicitly via a command line flag (i.e. '--disabled-metrics=metric1,metric2') ([#99217](https://github.com/kubernetes/kubernetes/pull/99217), [@logicalhan](https://github.com/logicalhan))
- New admission controller `DenyServiceExternalIPs` is available.  Clusters which do not *need* the Service `externalIPs` feature should enable this controller and be more secure. ([#97395](https://github.com/kubernetes/kubernetes/pull/97395), [@thockin](https://github.com/thockin))
- Overall, enable the feature of `PreferNominatedNode` will  improve the performance of scheduling where preemption might frequently happen, but in theory, enable the feature of `PreferNominatedNode`, the pod might not be scheduled to the best candidate node in the cluster. ([#93179](https://github.com/kubernetes/kubernetes/pull/93179), [@chendave](https://github.com/chendave)) [SIG Scheduling and Testing]
- Persistent Volumes formatted with the btrfs filesystem will now automatically resize when expanded. ([#99361](https://github.com/kubernetes/kubernetes/pull/99361), [@Novex](https://github.com/Novex)) [SIG Storage]
- Port the devicemanager to Windows node to allow device plugins like directx ([#93285](https://github.com/kubernetes/kubernetes/pull/93285), [@aarnaud](https://github.com/aarnaud)) [SIG Node, Testing and Windows]
- Removes cAdvisor JSON metrics (/stats/container, /stats/<podname>/<containername>, /stats/<namespace>/<podname>/<poduid>/<containername>) from the kubelet. ([#99236](https://github.com/kubernetes/kubernetes/pull/99236), [@pacoxu](https://github.com/pacoxu))
- Rename metrics `etcd_object_counts` to `apiserver_storage_object_counts` and mark it as stable. The original `etcd_object_counts` metrics name is marked as "Deprecated" and will be removed in the future. ([#99785](https://github.com/kubernetes/kubernetes/pull/99785), [@erain](https://github.com/erain)) [SIG API Machinery, Instrumentation and Testing]
- Sysctls graduates to General Availability and thus unconditionally enabled. ([#99158](https://github.com/kubernetes/kubernetes/pull/99158), [@wgahnagl](https://github.com/wgahnagl))
- The Kubernetes pause image manifest list now contains an image for Windows Server 20H2. ([#97322](https://github.com/kubernetes/kubernetes/pull/97322), [@claudiubelu](https://github.com/claudiubelu)) [SIG Windows]
- The NodeAffinity plugin implements the PreFilter extension, offering enhanced performance for Filter. ([#99213](https://github.com/kubernetes/kubernetes/pull/99213), [@AliceZhang2016](https://github.com/AliceZhang2016)) [SIG Scheduling]
- The `CronJobControllerV2` feature flag graduates to Beta and set to be enabled by default. ([#98878](https://github.com/kubernetes/kubernetes/pull/98878), [@soltysh](https://github.com/soltysh))
- The `EndpointSlice` mirroring controller mirrors endpoints annotations and labels to the generated endpoint slices, it also ensures that updates on any of these fields are mirrored. 
  The well-known annotation `endpoints.kubernetes.io/last-change-trigger-time` is skipped and not mirrored. ([#98116](https://github.com/kubernetes/kubernetes/pull/98116), [@aojea](https://github.com/aojea))
- The `RunAsGroup` feature has been promoted to GA in this release. ([#94641](https://github.com/kubernetes/kubernetes/pull/94641), [@krmayankk](https://github.com/krmayankk)) [SIG Auth and Node]
- The `ServiceAccountIssuerDiscovery` feature has graduated to GA, and is unconditionally enabled. The `ServiceAccountIssuerDiscovery` feature-gate will be removed in 1.22. ([#98553](https://github.com/kubernetes/kubernetes/pull/98553), [@mtaufen](https://github.com/mtaufen)) [SIG API Machinery, Auth and Testing]
- The `TTLAfterFinished` feature flag is now beta and enabled by default ([#98678](https://github.com/kubernetes/kubernetes/pull/98678), [@ahg-g](https://github.com/ahg-g))
- The apimachinery util/net function used to detect the bind address `ResolveBindAddress()` takes into consideration global IP addresses on loopback interfaces when 1) the host has default routes, or 2) there are no global IPs on those interfaces in order to support more complex network scenarios like BGP Unnumbered RFC 5549 ([#95790](https://github.com/kubernetes/kubernetes/pull/95790), [@aojea](https://github.com/aojea)) [SIG Network]
- The feature gate `RootCAConfigMap` graduated to GA in v1.21 and therefore will be unconditionally enabled. This flag will be removed in v1.22 release. ([#98033](https://github.com/kubernetes/kubernetes/pull/98033), [@zshihang](https://github.com/zshihang))
- The pause image upgraded to `v3.4.1` in kubelet and kubeadm for both Linux and Windows. ([#98205](https://github.com/kubernetes/kubernetes/pull/98205), [@pacoxu](https://github.com/pacoxu))
- Update pause container to run as pseudo user and group `65535:65535`. This implies the release of version 3.5 of the container images. ([#97963](https://github.com/kubernetes/kubernetes/pull/97963), [@saschagrunert](https://github.com/saschagrunert)) [SIG CLI, Cloud Provider, Cluster Lifecycle, Node, Release, Security and Testing]
- Update the latest validated version of Docker to 20.10 ([#98977](https://github.com/kubernetes/kubernetes/pull/98977), [@neolit123](https://github.com/neolit123)) [SIG CLI, Cluster Lifecycle and Node]
- Upgrade node local dns to 1.17.0 for better IPv6 support ([#99749](https://github.com/kubernetes/kubernetes/pull/99749), [@pacoxu](https://github.com/pacoxu)) [SIG Cloud Provider and Network]
- Upgrades `IPv6Dualstack` to `Beta` and turns it on by default. New clusters or existing clusters are not be affected until an actor starts adding secondary Pods and service CIDRS CLI flags as described here: [IPv4/IPv6 Dual-stack](https://github.com/kubernetes/enhancements/tree/master/keps/sig-network/563-dual-stack) ([#98969](https://github.com/kubernetes/kubernetes/pull/98969), [@khenidak](https://github.com/khenidak))
- Users might specify the `kubectl.kubernetes.io/default-container` annotation in a Pod to preselect container for kubectl commands. ([#99581](https://github.com/kubernetes/kubernetes/pull/99581), [@mengjiao-liu](https://github.com/mengjiao-liu)) [SIG CLI]
- When downscaling ReplicaSets, ready and creation timestamps are compared in a logarithmic scale. ([#99212](https://github.com/kubernetes/kubernetes/pull/99212), [@damemi](https://github.com/damemi)) [SIG Apps and Testing]
- When the kubelet is watching a ConfigMap or Secret purely in the context of setting environment variables
  for containers, only hold that watch for a defined duration before cancelling it. This change reduces the CPU
  and memory usage of the kube-apiserver in large clusters. ([#99393](https://github.com/kubernetes/kubernetes/pull/99393), [@chenyw1990](https://github.com/chenyw1990)) [SIG API Machinery, Node and Testing]
- WindowsEndpointSliceProxying feature gate has graduated to beta and is enabled by default. This means kube-proxy will  read from EndpointSlices instead of Endpoints on Windows by default. ([#99794](https://github.com/kubernetes/kubernetes/pull/99794), [@robscott](https://github.com/robscott)) [SIG Network]
- `kubectl wait` ensures that observedGeneration >= generation to prevent stale state reporting. An example scenario can be found on CRD updates. ([#97408](https://github.com/kubernetes/kubernetes/pull/97408), [@KnicKnic](https://github.com/KnicKnic))

#### Failing Test
- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Fixed generic ephemeal volumes with OwnerReferencesPermissionEnforcement admission plugin enabled. ([#101186](https://github.com/kubernetes/kubernetes/pull/101186), [@jsafrane](https://github.com/jsafrane)) [SIG Auth and Storage]
- Resolves an issue with the "ServiceAccountIssuerDiscovery should support OIDC discovery" conformance test failing on clusters which are configured with issuers outside the cluster ([#101725](https://github.com/kubernetes/kubernetes/pull/101725), [@mtaufen](https://github.com/mtaufen)) [SIG Auth and Testing]
- Escape the special characters like `[`, `]` and ` ` that exist in vsphere windows path ([#98830](https://github.com/kubernetes/kubernetes/pull/98830), [@liyanhui1228](https://github.com/liyanhui1228)) [SIG Storage and Windows]
- Kube-proxy: fix a bug on UDP `NodePort` Services where stale connection tracking entries may blackhole the traffic directed to the `NodePort` ([#98305](https://github.com/kubernetes/kubernetes/pull/98305), [@aojea](https://github.com/aojea))
- Kubelet: fixes a bug in the HostPort dockershim implementation that caused the conformance test "HostPort validates that there is no conflict between pods with same hostPort but different hostIP and protocol" to fail. ([#98755](https://github.com/kubernetes/kubernetes/pull/98755), [@aojea](https://github.com/aojea)) [SIG Cloud Provider, Network and Node]


#### Bug or Regression
- Fix NodeAuthenticator tests in dualstack ([#104840](https://github.com/kubernetes/kubernetes/pull/104840), [@ardaguclu](https://github.com/ardaguclu)) [SIG Auth and Testing]
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104447](https://github.com/kubernetes/kubernetes/pull/104447), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fixed occasional pod cgroup freeze when using cgroup v1 and systemd driver.
  Fixed "failed to create container ... unit already exists" when using cgroup v1 and systemd driver. ([#104530](https://github.com/kubernetes/kubernetes/pull/104530), [@kolyshkin](https://github.com/kolyshkin)) [SIG CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104151](https://github.com/kubernetes/kubernetes/pull/104151), [@aojea](https://github.com/aojea)) [SIG Network]
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric ([#100720](https://github.com/kubernetes/kubernetes/pull/100720), [@dntosas](https://github.com/dntosas)) [SIG Apps, Instrumentation, Scheduling and Storage]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104347](https://github.com/kubernetes/kubernetes/pull/104347), [@mauriciopoppe](https://github.com/mauriciopoppe)) [SIG Storage]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104673](https://github.com/kubernetes/kubernetes/pull/104673), [@thockin](https://github.com/thockin)) [SIG Network]
- Disable aufs module for gce clusters ([#103831](https://github.com/kubernetes/kubernetes/pull/103831), [@lizhuqi](https://github.com/lizhuqi)) [SIG Cloud Provider]
- Fix kube-apiserver metric reporting for the deprecated watch path of /api/<version>/watch/... ([#104190](https://github.com/kubernetes/kubernetes/pull/104190), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- Fix the code is leaking the defaulting between unrelated pod instances. ([#103284](https://github.com/kubernetes/kubernetes/pull/103284), [@kebe7jun](https://github.com/kebe7jun)) [SIG CLI]
- Fix: Provide IPv6 support for internal load balancer ([#103794](https://github.com/kubernetes/kubernetes/pull/103794), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: cleanup outdated routes ([#102935](https://github.com/kubernetes/kubernetes/pull/102935), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: ignore not a VMSS error for VMAS nodes in reconcileBackendPools ([#103997](https://github.com/kubernetes/kubernetes/pull/103997), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: return empty VMAS name if using standalone VM ([#103470](https://github.com/kubernetes/kubernetes/pull/103470), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed a bug that scheduler extenders are not called on preemptions ([#103019](https://github.com/kubernetes/kubernetes/pull/103019), [@ordovicia](https://github.com/ordovicia)) [SIG Scheduling]
- Fixes an issue cleaning up CertificateSigningRequest objects with an unparseable `status.certificate` field ([#103948](https://github.com/kubernetes/kubernetes/pull/103948), [@liggitt](https://github.com/liggitt)) [SIG Apps and Auth]
- Fixes issue with websocket-based watches of Service objects not closing correctly on timeout ([#102541](https://github.com/kubernetes/kubernetes/pull/102541), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102969](https://github.com/kubernetes/kubernetes/pull/102969), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider]
- Added jitter factor to lease controller that better smears load on kube-apiserver over time. ([#101652](https://github.com/kubernetes/kubernetes/pull/101652), [@marseel](https://github.com/marseel)) [SIG API Machinery and Scalability]
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: avoid setting cached Sku when updating VMSS and VMSS instances ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix a bug on the endpoint slices mirroring controller where endpoint NotReadyAddresses were mirrored as Ready to the corresponding EndpointSlice ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix removing pods from podTopologyHints mapping ([#101892](https://github.com/kubernetes/kubernetes/pull/101892), [@aheng-ch](https://github.com/aheng-ch)) [SIG Node]
- Fix resource enforcement when using systemd cgroup driver ([#102147](https://github.com/kubernetes/kubernetes/pull/102147), [@kolyshkin](https://github.com/kolyshkin)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Fix: avoid nil-pointer panic when checking the frontend IP configuration ([#101739](https://github.com/kubernetes/kubernetes/pull/101739), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: not tagging static public IP ([#101752](https://github.com/kubernetes/kubernetes/pull/101752), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed garbage collection of dangling VolumeAttachments for PersistentVolumes migrated to CSI on startup of kube-controller-manager. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt)) [SIG Apps and Storage]
- Kube-proxy log now shows the "Skipping topology aware endpoint filtering since no hints were provided for zone" warning under the right conditions ([#101857](https://github.com/kubernetes/kubernetes/pull/101857), [@dervoeti](https://github.com/dervoeti)) [SIG Network]
- Kubeadm: remove the "ephemeral_storage" request from the etcd static pod that kubeadm deploys on stacked etcd control plane nodes. This request has caused sporadic failures on some setups due to a problem in the kubelet with cadvisor and the LocalStorageCapacityIsolation feature gate. See this issue for more details: https://github.com/kubernetes/kubernetes/issues/99305 ([#102673](https://github.com/kubernetes/kubernetes/pull/102673), [@jackfrancis](https://github.com/jackfrancis)) [SIG Cluster Lifecycle]
- Kubeadm: when using a custom image repository for CoreDNS kubeadm now will append the "coredns"  image name instead of  "coredns/coredns", thus restoring the behaviour existing before the v1.21 release. Users who rely on nested folder for the coredns image should set the "clusterConfiguration.dns.imageRepository" value including the nested path name (e.g using "registry.company.xyz/coredns" will force kubeadm to use "registry.company.xyz/coredns/coredns" image). No action is needed if using the default registry (k8s.gcr.io). ([#102502](https://github.com/kubernetes/kubernetes/pull/102502), [@ykakarap](https://github.com/ykakarap)) [SIG Cluster Lifecycle]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- Reverted the previous fix for portforward cleanup because it introduced a kubelet regression which can lead into segmentation faults. ([#102587](https://github.com/kubernetes/kubernetes/pull/102587), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Adds node event handlers to dual stack kube-proxy implementation to fix Topology Aware Hints. ([#101054](https://github.com/kubernetes/kubernetes/pull/101054), [@robscott](https://github.com/robscott)) [SIG Network]
- Azurefile: Normalize share name to not include capital letters ([#100731](https://github.com/kubernetes/kubernetes/pull/100731), [@kassarl](https://github.com/kassarl)) [SIG Cloud Provider and Storage]
- EndpointSlice IP validation now matches Endpoints IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Ensure service deleted when the Azure resource group has been deleted ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix EndpointSlice describe panic when an Endpoint doesn't have zone ([#101025](https://github.com/kubernetes/kubernetes/pull/101025), [@tnqn](https://github.com/tnqn)) [SIG CLI]
- Fix display of Job completion mode on kubectl describe ([#101198](https://github.com/kubernetes/kubernetes/pull/101198), [@alculquicondor](https://github.com/alculquicondor)) [SIG CLI]
- Fix: azure file inline volume namespace issue in csi migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx)) [SIG Apps, Cloud Provider, Node and Storage]
- Fix: set "host is down" as corrupted mount ([#101398](https://github.com/kubernetes/kubernetes/pull/101398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug where startupProbe stopped working after a container's first restart ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed port-forward memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Kubectl create service now respects namespace flag ([#101005](https://github.com/kubernetes/kubernetes/pull/101005), [@zxh326](https://github.com/zxh326)) [SIG CLI]
- Kubelet: improve the performance when waiting for a synchronization of the node list with the kube-apiserver ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123)) [SIG Node]
- No support endpointslice in linux userpace mode ([#101504](https://github.com/kubernetes/kubernetes/pull/101504), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Renames the timeout field for the DelegatingAuthenticationOptions to TokenRequestTimeout and set the timeout only for the token review client. Previously the timeout was also applied to watches making them reconnecting every 10 seconds. ([#101102](https://github.com/kubernetes/kubernetes/pull/101102), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth and Cloud Provider]
- Respect ExecProbeTimeout=false for dockershim ([#101127](https://github.com/kubernetes/kubernetes/pull/101127), [@jackfrancis](https://github.com/jackfrancis)) [SIG Node and Testing]
- Upgrades functionality of `kubectl kustomize` as described at https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.1.0 ([#101177](https://github.com/kubernetes/kubernetes/pull/101177), [@KnVerey](https://github.com/KnVerey)) [SIG CLI]
- AcceleratorStats will be available in the Summary API of kubelet when cri_stats_provider is used. ([#96873](https://github.com/kubernetes/kubernetes/pull/96873), [@ruiwen-zhao](https://github.com/ruiwen-zhao)) [SIG Node]
- All data is no longer automatically deleted when a failure is detected during creation of the volume data file on a CSI volume. Now only the data file and volume path is removed. ([#96021](https://github.com/kubernetes/kubernetes/pull/96021), [@huffmanca](https://github.com/huffmanca))
- Clean ReplicaSet by revision instead of creation timestamp in deployment controller ([#97407](https://github.com/kubernetes/kubernetes/pull/97407), [@waynepeking348](https://github.com/waynepeking348)) [SIG Apps]
- Cleanup subnet in frontend IP configs to prevent huge subnet request bodies in some scenarios. ([#98133](https://github.com/kubernetes/kubernetes/pull/98133), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Client-go exec credential plugins will pass stdin only when interactive terminal is detected on stdin. This fixes a bug where previously it was checking if **stdout** is an interactive terminal. ([#99654](https://github.com/kubernetes/kubernetes/pull/99654), [@ankeesler](https://github.com/ankeesler))
- Cloud-controller-manager: routes controller should not depend on --allocate-node-cidrs ([#97029](https://github.com/kubernetes/kubernetes/pull/97029), [@andrewsykim](https://github.com/andrewsykim)) [SIG Cloud Provider and Testing]
- Cluster Autoscaler version bump to v1.20.0 ([#97011](https://github.com/kubernetes/kubernetes/pull/97011), [@towca](https://github.com/towca))
- Creating a PVC with DataSource should fail for non-CSI plugins. ([#97086](https://github.com/kubernetes/kubernetes/pull/97086), [@xing-yang](https://github.com/xing-yang)) [SIG Apps and Storage]
- EndpointSlice controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#99345](https://github.com/kubernetes/kubernetes/pull/99345), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSlice controllers are less likely to create duplicate EndpointSlices. ([#100103](https://github.com/kubernetes/kubernetes/pull/100103), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSliceMirroring controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#99756](https://github.com/kubernetes/kubernetes/pull/99756), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Ensure all vSphere nodes are are tracked by volume attach-detach controller ([#96689](https://github.com/kubernetes/kubernetes/pull/96689), [@gnufied](https://github.com/gnufied))
- Ensure empty string annotations are copied over in rollbacks. ([#94858](https://github.com/kubernetes/kubernetes/pull/94858), [@waynepeking348](https://github.com/waynepeking348))
- Ensure only one LoadBalancer rule is created when HA mode is enabled ([#99825](https://github.com/kubernetes/kubernetes/pull/99825), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Ensure that client-go's EventBroadcaster is safe (non-racy) during shutdown. ([#95664](https://github.com/kubernetes/kubernetes/pull/95664), [@DirectXMan12](https://github.com/DirectXMan12)) [SIG API Machinery]
- Explicitly pass `KUBE_BUILD_CONFORMANCE=y` in `package-tarballs` to reenable building the conformance tarballs. ([#100571](https://github.com/kubernetes/kubernetes/pull/100571), [@puerco](https://github.com/puerco))
- Fix Azure file migration e2e test failure when CSIMigration is turned on. ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx))
- Fix CSI-migrated inline EBS volumes failing to mount if their volumeID is prefixed by aws:// ([#96821](https://github.com/kubernetes/kubernetes/pull/96821), [@wongma7](https://github.com/wongma7)) [SIG Storage]
- Fix CVE-2020-8555 for Gluster client connections. ([#97922](https://github.com/kubernetes/kubernetes/pull/97922), [@liggitt](https://github.com/liggitt)) [SIG Storage]
- Fix NPE in ephemeral storage eviction ([#98261](https://github.com/kubernetes/kubernetes/pull/98261), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fix PermissionDenied issue on SMB mount for Windows ([#99550](https://github.com/kubernetes/kubernetes/pull/99550), [@andyzhangx](https://github.com/andyzhangx))
- Fix bug that would let the Horizontal Pod Autoscaler scale down despite at least one metric being unavailable/invalid ([#99514](https://github.com/kubernetes/kubernetes/pull/99514), [@mikkeloscar](https://github.com/mikkeloscar)) [SIG Apps and Autoscaling]
- Fix cgroup handling for systemd with cgroup v2 ([#98365](https://github.com/kubernetes/kubernetes/pull/98365), [@odinuge](https://github.com/odinuge)) [SIG Node]
- Fix counting error in service/nodeport/loadbalancer quota check ([#97451](https://github.com/kubernetes/kubernetes/pull/97451), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery, Network and Testing]
- Fix errors when accessing Windows container stats for Dockershim ([#98510](https://github.com/kubernetes/kubernetes/pull/98510), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node and Windows]
- Fix kube-proxy container image architecture for non amd64 images. ([#98526](https://github.com/kubernetes/kubernetes/pull/98526), [@saschagrunert](https://github.com/saschagrunert))
- Fix missing cadvisor machine metrics. ([#97006](https://github.com/kubernetes/kubernetes/pull/97006), [@lingsamuel](https://github.com/lingsamuel)) [SIG Node]
- Fix nil VMSS name when setting service to auto mode ([#97366](https://github.com/kubernetes/kubernetes/pull/97366), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix privileged config of Pod Sandbox which was previously ignored. ([#96877](https://github.com/kubernetes/kubernetes/pull/96877), [@xeniumlee](https://github.com/xeniumlee))
- Fix the panic when kubelet registers if a node object already exists with no Status.Capacity or Status.Allocatable ([#95269](https://github.com/kubernetes/kubernetes/pull/95269), [@SataQiu](https://github.com/SataQiu)) [SIG Node]
- Fix the regression with the slow pods termination. Before this fix pods may take an additional time to terminate - up to one minute. Reversing the change that ensured that CNI resources cleaned up when the pod is removed on API server. ([#97980](https://github.com/kubernetes/kubernetes/pull/97980), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Node]
- Fix to recover CSI volumes from certain dangling attachments ([#96617](https://github.com/kubernetes/kubernetes/pull/96617), [@yuga711](https://github.com/yuga711)) [SIG Apps and Storage]
- Fix: azure file latency issue for metadata-heavy workloads ([#97082](https://github.com/kubernetes/kubernetes/pull/97082), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed Cinder volume IDs on OpenStack Train ([#96673](https://github.com/kubernetes/kubernetes/pull/96673), [@jsafrane](https://github.com/jsafrane)) [SIG Cloud Provider]
- Fixed FibreChannel volume plugin corrupting filesystems on detach of multipath volumes. ([#97013](https://github.com/kubernetes/kubernetes/pull/97013), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixed a bug in kubelet that will saturate CPU utilization after containerd got restarted. ([#97174](https://github.com/kubernetes/kubernetes/pull/97174), [@hanlins](https://github.com/hanlins)) [SIG Node]
- Fixed a bug that causes smaller number of conntrack-max being used under CPU static policy. (#99225, @xh4n3) ([#99613](https://github.com/kubernetes/kubernetes/pull/99613), [@xh4n3](https://github.com/xh4n3)) [SIG Network]
- Fixed a bug that on k8s nodes, when the policy of INPUT chain in filter table is not ACCEPT, healthcheck nodeport would not work.
  Added iptables rules to allow healthcheck nodeport traffic. ([#97824](https://github.com/kubernetes/kubernetes/pull/97824), [@hanlins](https://github.com/hanlins)) [SIG Network]
- Fixed a bug that the kubelet cannot start on BtrfS. ([#98042](https://github.com/kubernetes/kubernetes/pull/98042), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Fixed a race condition on API server startup ensuring previously created webhook configurations are effective before the first write request is admitted. ([#95783](https://github.com/kubernetes/kubernetes/pull/95783), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixed an issue with garbage collection failing to clean up namespaced children of an object also referenced incorrectly by cluster-scoped children ([#98068](https://github.com/kubernetes/kubernetes/pull/98068), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Fixed authentication_duration_seconds metric scope. Previously, it included whole apiserver request duration which yields inaccurate results. ([#99944](https://github.com/kubernetes/kubernetes/pull/99944), [@marseel](https://github.com/marseel))
- Fixed bug in CPUManager with race on container map access ([#97427](https://github.com/kubernetes/kubernetes/pull/97427), [@klueska](https://github.com/klueska)) [SIG Node]
- Fixed bug that caused cAdvisor to incorrectly detect single-socket multi-NUMA topology. ([#99315](https://github.com/kubernetes/kubernetes/pull/99315), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixed cleanup of block devices when /var/lib/kubelet is a symlink. ([#96889](https://github.com/kubernetes/kubernetes/pull/96889), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixed no effect namespace when exposing deployment with --dry-run=client. ([#97492](https://github.com/kubernetes/kubernetes/pull/97492), [@masap](https://github.com/masap)) [SIG CLI]
- Fixed provisioning of Cinder volumes migrated to CSI when StorageClass with AllowedTopologies was used. ([#98311](https://github.com/kubernetes/kubernetes/pull/98311), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixes a bug of identifying the correct containerd process. ([#97888](https://github.com/kubernetes/kubernetes/pull/97888), [@pacoxu](https://github.com/pacoxu))
- Fixes add-on manager leader election to use leases instead of endpoints, similar to what kube-controller-manager does in 1.20 ([#98968](https://github.com/kubernetes/kubernetes/pull/98968), [@liggitt](https://github.com/liggitt))
- Fixes connection errors when using `--volume-host-cidr-denylist` or `--volume-host-allow-local-loopback` ([#98436](https://github.com/kubernetes/kubernetes/pull/98436), [@liggitt](https://github.com/liggitt)) [SIG Network and Storage]
- Fixes problem where invalid selector on `PodDisruptionBudget` leads to a nil pointer dereference that causes the Controller manager to crash loop. ([#98750](https://github.com/kubernetes/kubernetes/pull/98750), [@mortent](https://github.com/mortent))
- Fixes spurious errors about IPv6 in `kube-proxy` logs on nodes with IPv6 disabled. ([#99127](https://github.com/kubernetes/kubernetes/pull/99127), [@danwinship](https://github.com/danwinship))
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#96876](https://github.com/kubernetes/kubernetes/pull/96876), [@howieyuen](https://github.com/howieyuen)) [SIG Apps and Node]
- GCE Internal LoadBalancer sync loop will now release the ILB IP address upon sync failure. An error in ILB forwarding rule creation will no longer leak IP addresses. ([#97740](https://github.com/kubernetes/kubernetes/pull/97740), [@prameshj](https://github.com/prameshj)) [SIG Cloud Provider and Network]
- Ignore update pod with no new images in alwaysPullImages admission controller ([#96668](https://github.com/kubernetes/kubernetes/pull/96668), [@pacoxu](https://github.com/pacoxu)) [SIG Apps, Auth and Node]
- Improve speed of vSphere PV provisioning and reduce number of API calls ([#100054](https://github.com/kubernetes/kubernetes/pull/100054), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- KUBECTL_EXTERNAL_DIFF now accepts equal sign for additional parameters. ([#98158](https://github.com/kubernetes/kubernetes/pull/98158), [@dougsland](https://github.com/dougsland)) [SIG CLI]
- Kube-apiserver: an update of a pod with a generic ephemeral volume dropped that volume if the feature had been disabled since creating the pod with such a volume ([#99446](https://github.com/kubernetes/kubernetes/pull/99446), [@pohly](https://github.com/pohly)) [SIG Apps, Node and Storage]
- Kube-proxy: remove deprecated --cleanup-ipvs flag of kube-proxy, and make --cleanup flag always to flush IPVS ([#97336](https://github.com/kubernetes/kubernetes/pull/97336), [@maaoBit](https://github.com/maaoBit)) [SIG Network]
- Kubeadm installs etcd v3.4.13 when creating cluster v1.19 ([#97244](https://github.com/kubernetes/kubernetes/pull/97244), [@pacoxu](https://github.com/pacoxu))
- Kubeadm: Fixes a kubeadm upgrade bug that could cause a custom CoreDNS configuration to be replaced with the default. ([#97016](https://github.com/kubernetes/kubernetes/pull/97016), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cluster Lifecycle]
- Kubeadm: Some text in the `kubeadm upgrade plan` output has changed. If you have scripts or other automation that parses this output, please review these changes and update your scripts to account for the new output. ([#98728](https://github.com/kubernetes/kubernetes/pull/98728), [@stmcginnis](https://github.com/stmcginnis)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug in the host memory detection code on 32bit Linux platforms ([#97403](https://github.com/kubernetes/kubernetes/pull/97403), [@abelbarrera15](https://github.com/abelbarrera15)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug where "kubeadm join" would not properly handle missing names for existing etcd members. ([#97372](https://github.com/kubernetes/kubernetes/pull/97372), [@ihgann](https://github.com/ihgann)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug where "kubeadm upgrade" commands can fail if CoreDNS v1.8.0 is installed. ([#97919](https://github.com/kubernetes/kubernetes/pull/97919), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug where external credentials in an existing admin.conf prevented the CA certificate to be written in the cluster-info ConfigMap. ([#98882](https://github.com/kubernetes/kubernetes/pull/98882), [@kvaps](https://github.com/kvaps)) [SIG Cluster Lifecycle]
- Kubeadm: get k8s CI version markers from k8s infra bucket ([#98836](https://github.com/kubernetes/kubernetes/pull/98836), [@hasheddan](https://github.com/hasheddan)) [SIG Cluster Lifecycle and Release]
- Kubeadm: skip validating pod subnet against node-cidr-mask when allocate-node-cidrs is set to be false ([#98984](https://github.com/kubernetes/kubernetes/pull/98984), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubectl logs: `--ignore-errors` is now honored by all containers, maintaining consistency with parallelConsumeRequest behavior. ([#97686](https://github.com/kubernetes/kubernetes/pull/97686), [@wzshiming](https://github.com/wzshiming))
- Kubectl-convert: Fix `no kind "Ingress" is registered for version` error ([#97754](https://github.com/kubernetes/kubernetes/pull/97754), [@wzshiming](https://github.com/wzshiming))
- Kubectl: Fixed panic when describing an ingress backend without an API Group ([#100505](https://github.com/kubernetes/kubernetes/pull/100505), [@lauchokyip](https://github.com/lauchokyip)) [SIG CLI]
- Kubelet now cleans up orphaned volume directories automatically ([#95301](https://github.com/kubernetes/kubernetes/pull/95301), [@lorenz](https://github.com/lorenz)) [SIG Node and Storage]
- Kubelet.exe on Windows now checks that the process running as administrator and the executing user account is listed in the built-in administrators group.  This is the equivalent to checking the process is running as uid 0. ([#96616](https://github.com/kubernetes/kubernetes/pull/96616), [@perithompson](https://github.com/perithompson)) [SIG Node and Windows]
- Kubelet: Fix kubelet from panic after getting the wrong signal ([#98200](https://github.com/kubernetes/kubernetes/pull/98200), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Kubelet: Fix repeatedly acquiring the inhibit lock ([#98088](https://github.com/kubernetes/kubernetes/pull/98088), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Kubelet: Fixed the bug of getting the number of cpu when the number of cpu logical processors is more than 64 in windows ([#97378](https://github.com/kubernetes/kubernetes/pull/97378), [@hwdef](https://github.com/hwdef)) [SIG Node and Windows]
- Limits lease to have 1000 maximum attached objects. ([#98257](https://github.com/kubernetes/kubernetes/pull/98257), [@lingsamuel](https://github.com/lingsamuel))
- Mitigate CVE-2020-8555 for kube-up using GCE by preventing local loopback folume hosts. ([#97934](https://github.com/kubernetes/kubernetes/pull/97934), [@mattcary](https://github.com/mattcary)) [SIG Cloud Provider and Storage]
- On single-stack configured (IPv4 or IPv6, but not both) clusters, Services which are both headless (no clusterIP) and selectorless (empty or undefined selector) will report `ipFamilyPolicy RequireDualStack` and will have entries in `ipFamilies[]` for both IPv4 and IPv6.  This is a change from alpha, but does not have any impact on the manually-specified Endpoints and EndpointSlices for the Service. ([#99555](https://github.com/kubernetes/kubernetes/pull/99555), [@thockin](https://github.com/thockin)) [SIG Apps and Network]
- Performance regression #97685 has been fixed. ([#97860](https://github.com/kubernetes/kubernetes/pull/97860), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery]
- Pod Log stats for windows now reports metrics ([#99221](https://github.com/kubernetes/kubernetes/pull/99221), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node, Storage, Testing and Windows]
- Pod status updates faster when reacting on probe results. The first readiness probe will be called faster when startup probes succeeded, which will make Pod status as ready faster. ([#98376](https://github.com/kubernetes/kubernetes/pull/98376), [@matthyx](https://github.com/matthyx))
- Readjust `kubelet_containers_per_pod_count` buckets to only show metrics greater than 1. ([#98169](https://github.com/kubernetes/kubernetes/pull/98169), [@wawa0210](https://github.com/wawa0210))
- Remove CSI topology from migrated in-tree gcepd volume. ([#97823](https://github.com/kubernetes/kubernetes/pull/97823), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Cloud Provider and Storage]
- Requests with invalid timeout parameters in the request URL now appear in the audit log correctly. ([#96901](https://github.com/kubernetes/kubernetes/pull/96901), [@tkashem](https://github.com/tkashem)) [SIG API Machinery and Testing]
- Resolve a "concurrent map read and map write" crashing error in the kubelet ([#95111](https://github.com/kubernetes/kubernetes/pull/95111), [@choury](https://github.com/choury)) [SIG Node]
- Resolves spurious `Failed to list *v1.Secret` or `Failed to list *v1.ConfigMap` messages in kubelet logs. ([#99538](https://github.com/kubernetes/kubernetes/pull/99538), [@liggitt](https://github.com/liggitt)) [SIG Auth and Node]
- ResourceQuota of an entity now inclusively calculate Pod overhead ([#99600](https://github.com/kubernetes/kubernetes/pull/99600), [@gjkim42](https://github.com/gjkim42))
- Return zero time (midnight on Jan. 1, 1970) instead of negative number when reporting startedAt and finishedAt of the not started or a running Pod when using `dockershim` as a runtime. ([#99585](https://github.com/kubernetes/kubernetes/pull/99585), [@Iceber](https://github.com/Iceber))
- Reverts breaking change to inline AzureFile volumes; referenced secrets are now searched for in the same namespace as the pod as in previous releases. ([#100563](https://github.com/kubernetes/kubernetes/pull/100563), [@msau42](https://github.com/msau42))
- Scores from InterPodAffinity have stronger differentiation. ([#98096](https://github.com/kubernetes/kubernetes/pull/98096), [@leileiwan](https://github.com/leileiwan)) [SIG Scheduling]
- Specifying the KUBE_TEST_REPO environment variable when e2e tests are executed will instruct the test infrastructure to load that image from a location within the specified repo, using a predefined pattern. ([#93510](https://github.com/kubernetes/kubernetes/pull/93510), [@smarterclayton](https://github.com/smarterclayton)) [SIG Testing]
- Static pods will be deleted gracefully. ([#98103](https://github.com/kubernetes/kubernetes/pull/98103), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Sync node status during kubelet node shutdown.
  Adds an pod admission handler that rejects new pods when the node is in progress of shutting down. ([#98005](https://github.com/kubernetes/kubernetes/pull/98005), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- The calculation of pod UIDs for static pods has changed to ensure each static pod gets a unique value - this will cause all static pod containers to be recreated/restarted if an in-place kubelet upgrade from 1.20 to 1.21 is performed. Note that draining pods before upgrading the kubelet across minor versions is the supported upgrade path. ([#87461](https://github.com/kubernetes/kubernetes/pull/87461), [@bboreham](https://github.com/bboreham)) [SIG Node]
- The maximum number of ports allowed in EndpointSlices has been increased from 100 to 20,000 ([#99795](https://github.com/kubernetes/kubernetes/pull/99795), [@robscott](https://github.com/robscott)) [SIG Network]
- Truncates a message if it hits the `NoteLengthLimit` when the scheduler records an event for the pod that indicates the pod has failed to schedule. ([#98715](https://github.com/kubernetes/kubernetes/pull/98715), [@carlory](https://github.com/carlory))
- Updated k8s.gcr.io/ingress-gce-404-server-with-metrics-amd64 to a version that serves /metrics endpoint on a non-default port. ([#97621](https://github.com/kubernetes/kubernetes/pull/97621), [@vbannai](https://github.com/vbannai)) [SIG Cloud Provider]
- Updates the commands `
   - kubectl kustomize {arg}
   - kubectl apply -k {arg}
  `to use same code as kustomize CLI [v4.0.5](https://github.com/kubernetes-sigs/kustomize/releases/tag/kustomize%2Fv4.0.5) ([#98946](https://github.com/kubernetes/kubernetes/pull/98946), [@monopole](https://github.com/monopole))
- Use force unmount for NFS volumes if regular mount fails after 1 minute timeout ([#96844](https://github.com/kubernetes/kubernetes/pull/96844), [@gnufied](https://github.com/gnufied)) [SIG Storage]
- Use network.Interface.VirtualMachine.ID to get the binded VM
  Skip standalone VM when reconciling LoadBalancer ([#97635](https://github.com/kubernetes/kubernetes/pull/97635), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Using exec auth plugins with kubectl no longer results in warnings about constructing many client instances from the same exec auth config. ([#97857](https://github.com/kubernetes/kubernetes/pull/97857), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Auth]
- When a CNI plugin returns dual-stack pod IPs, kubelet will now try to respect the
  "primary IP family" of the cluster by picking a primary pod IP of the same family
  as the (primary) node IP, rather than assuming that the CNI plugin returned the IPs
  in the order the administrator wanted (since some CNI plugins don't allow
  configuring this). ([#97979](https://github.com/kubernetes/kubernetes/pull/97979), [@danwinship](https://github.com/danwinship)) [SIG Network and Node]
- When dynamically provisioning Azure File volumes for a premium account, the requested size will be set to 100GB if the request is initially lower than this value to accommodate Azure File requirements. ([#99122](https://github.com/kubernetes/kubernetes/pull/99122), [@huffmanca](https://github.com/huffmanca)) [SIG Cloud Provider and Storage]
- When using `Containerd` on Windows, the `C:\Windows\System32\drivers\etc\hosts` file will now be managed by kubelet. ([#83730](https://github.com/kubernetes/kubernetes/pull/83730), [@claudiubelu](https://github.com/claudiubelu))
- `VolumeBindingArgs` now allow `BindTimeoutSeconds` to be set as zero, while the value zero indicates no waiting for the checking of volume binding operation. ([#99835](https://github.com/kubernetes/kubernetes/pull/99835), [@chendave](https://github.com/chendave)) [SIG Scheduling and Storage]
- `kubectl exec` and `kubectl attach` now honor the `--quiet` flag which suppresses output from the local binary that could be confused by a script with the remote command output (all non-failure output is hidden). In addition, print inline with exec and attach the list of alternate containers when we default to the first spec.container. ([#99004](https://github.com/kubernetes/kubernetes/pull/99004), [@smarterclayton](https://github.com/smarterclayton)) [SIG CLI]




#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]
- Client-go: reduce verbosity of "Starting/Stopping reflector" messages to 3 again ([#102788](https://github.com/kubernetes/kubernetes/pull/102788), [@pohly](https://github.com/pohly)) [SIG API Machinery]
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102340](https://github.com/kubernetes/kubernetes/pull/102340), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]
- APIs for kubelet annotations and labels from `k8s.io/kubernetes/pkg/kubelet/apis` are now moved under `k8s.io/kubelet/pkg/apis/` ([#98931](https://github.com/kubernetes/kubernetes/pull/98931), [@michaelbeaumont](https://github.com/michaelbeaumont))
- Apiserver_request_duration_seconds is promoted to stable status. ([#99925](https://github.com/kubernetes/kubernetes/pull/99925), [@logicalhan](https://github.com/logicalhan)) [SIG API Machinery, Instrumentation and Testing]
- Bump github.com/Azure/go-autorest/autorest to v0.11.12 ([#97033](https://github.com/kubernetes/kubernetes/pull/97033), [@patrickshan](https://github.com/patrickshan)) [SIG API Machinery, CLI, Cloud Provider and Cluster Lifecycle]
- Clients required to use go1.15.8+ or go1.16+ if kube-apiserver has the goaway feature enabled to avoid unexpected data race condition. ([#98809](https://github.com/kubernetes/kubernetes/pull/98809), [@answer1991](https://github.com/answer1991))
- Delete deprecated `service.beta.kubernetes.io/azure-load-balancer-mixed-protocols` mixed procotol annotation in favor of the MixedProtocolLBService feature ([#97096](https://github.com/kubernetes/kubernetes/pull/97096), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- EndpointSlice generation is now incremented when labels change. ([#99750](https://github.com/kubernetes/kubernetes/pull/99750), [@robscott](https://github.com/robscott)) [SIG Network]
- Featuregate AllowInsecureBackendProxy graduates to GA and unconditionally enabled. ([#99658](https://github.com/kubernetes/kubernetes/pull/99658), [@deads2k](https://github.com/deads2k))
- Increase timeout for pod lifecycle test to reach pod status=ready ([#96691](https://github.com/kubernetes/kubernetes/pull/96691), [@hh](https://github.com/hh))
- Increased `CSINodeIDMaxLength` from 128 bytes to 192 bytes. ([#98753](https://github.com/kubernetes/kubernetes/pull/98753), [@Jiawei0227](https://github.com/Jiawei0227))
- Kube-apiserver: The OIDC authenticator no longer waits 10 seconds before attempting to fetch the metadata required to verify tokens. ([#97693](https://github.com/kubernetes/kubernetes/pull/97693), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]
- Kube-proxy: Traffic from the cluster directed to ExternalIPs is always sent directly to the Service. ([#96296](https://github.com/kubernetes/kubernetes/pull/96296), [@aojea](https://github.com/aojea)) [SIG Network and Testing]
- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubectl: The deprecated `kubectl alpha debug` command is removed. Use `kubectl debug` instead. ([#98111](https://github.com/kubernetes/kubernetes/pull/98111), [@pandaamanda](https://github.com/pandaamanda)) [SIG CLI]
- Kubelet command line flags related to dockershim are now showing deprecation message as they will be removed along with dockershim in future release. ([#98730](https://github.com/kubernetes/kubernetes/pull/98730), [@dims](https://github.com/dims))
- Official support to build kubernetes with docker-machine / remote docker is removed. This change does not affect building kubernetes with docker locally. ([#97618](https://github.com/kubernetes/kubernetes/pull/97618), [@jherrera123](https://github.com/jherrera123)) [SIG Release and Testing]
- Process start time on Windows now uses current process information ([#97491](https://github.com/kubernetes/kubernetes/pull/97491), [@jsturtevant](https://github.com/jsturtevant)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Windows]
- Resolves flakes in the Ingress conformance tests due to conflicts with controllers updating the Ingress object ([#98430](https://github.com/kubernetes/kubernetes/pull/98430), [@liggitt](https://github.com/liggitt)) [SIG Network and Testing]
- The `AttachVolumeLimit` feature gate (GA since v1.17) has been removed and now unconditionally enabled. ([#96539](https://github.com/kubernetes/kubernetes/pull/96539), [@ialidzhikov](https://github.com/ialidzhikov))
- The `CSINodeInfo` feature gate that is GA since v1.17 is unconditionally enabled, and can no longer be specified via the `--feature-gates` argument. ([#96561](https://github.com/kubernetes/kubernetes/pull/96561), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Apps, Auth, Scheduling, Storage and Testing]
- The `apiserver_request_total` metric is promoted to stable status and no longer has a content-type dimensions, so any alerts/charts which presume the existence of this will fail. This is however, unlikely to be the case since it was effectively an unbounded dimension in the first place. ([#99788](https://github.com/kubernetes/kubernetes/pull/99788), [@logicalhan](https://github.com/logicalhan))
- The default delegating authorization options now allow unauthenticated access to healthz, readyz, and livez.  A system:masters user connecting to an authz delegator will not perform an authz check. ([#98325](https://github.com/kubernetes/kubernetes/pull/98325), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Auth, Cloud Provider and Scheduling]
- The deprecated feature gates `CSIDriverRegistry`, `BlockVolume` and `CSIBlockVolume` are now unconditionally enabled and can no longer be specified in component invocations. ([#98021](https://github.com/kubernetes/kubernetes/pull/98021), [@gavinfish](https://github.com/gavinfish)) [SIG Storage]
- The deprecated feature gates `RotateKubeletClientCertificate`, `AttachVolumeLimit`, `VolumePVCDataSource` and `EvenPodsSpread` are now unconditionally enabled and can no longer be specified in component invocations. ([#97306](https://github.com/kubernetes/kubernetes/pull/97306), [@gavinfish](https://github.com/gavinfish)) [SIG Node, Scheduling and Storage]
- The e2e suite can be instructed not to wait for pods in kube-system to be ready or for all nodes to be ready by passing `--allowed-not-ready-nodes=-1` when invoking the e2e.test program. This allows callers to run subsets of the e2e suite in scenarios other than perfectly healthy clusters. ([#98781](https://github.com/kubernetes/kubernetes/pull/98781), [@smarterclayton](https://github.com/smarterclayton)) [SIG Testing]
- The feature gates `WindowsGMSA` and `WindowsRunAsUserName` that are GA since v1.18 are now removed. ([#96531](https://github.com/kubernetes/kubernetes/pull/96531), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Node and Windows]
- The new `-gce-zones` flag on the `e2e.test` binary instructs tests that check for information about how the cluster interacts with the cloud to limit their queries to the provided zone list. If not specified, the current behavior of asking the cloud provider for all available zones in multi zone clusters is preserved. ([#98787](https://github.com/kubernetes/kubernetes/pull/98787), [@smarterclayton](https://github.com/smarterclayton)) [SIG API Machinery, Cluster Lifecycle and Testing]
- Update cri-tools to [v1.20.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.20.0) ([#97967](https://github.com/kubernetes/kubernetes/pull/97967), [@rajibmitra](https://github.com/rajibmitra)) [SIG Cloud Provider]
- Windows nodes on GCE will take longer to start due to dependencies installed at node creation time. ([#98284](https://github.com/kubernetes/kubernetes/pull/98284), [@pjh](https://github.com/pjh)) [SIG Cloud Provider]
- `apiserver_storage_objects` (a newer version of `etcd_object_counts`) is promoted and marked as stable. ([#100082](https://github.com/kubernetes/kubernetes/pull/100082), [@logicalhan](https://github.com/logicalhan))



### app-operator [5.2.0](https://github.com/giantswarm/app-operator/releases/tag/v5.2.0)

#### Changed
- Reject App CRs with version labels with the legacy `1.0.0` value.
- Validate `.spec.catalog` using Catalog CRs instead of AppCatalog CRs.

#### Fixed
- Fix creating `AppCatalog` CRs in appcatalogsync resource.



### aws-cni [1.9.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.9.0)

* Enhancement - [EC2 sdk model override](https://github.com/aws/amazon-vpc-cni-k8s/pull/1508) (#1508, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Prefix Delegation feature support](https://github.com/aws/amazon-vpc-cni-k8s/pull/1516) (#1516, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Header formatting for env variable](https://github.com/aws/amazon-vpc-cni-k8s/pull/1522) (#1522, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [non-nitro instances init issues](https://github.com/aws/amazon-vpc-cni-k8s/pull/1527) (#1527, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Add metrics for total prefix count and ips used per cidr](https://github.com/aws/amazon-vpc-cni-k8s/pull/1530) (#1530, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Update documentation for PD](https://github.com/aws/amazon-vpc-cni-k8s/pull/1540) (#1540, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Update SDK Go version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1544) (#1544, [@jayanthvn](https://github.com/jayanthvn))



### cluster-operator [3.10.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.10.0)

#### Changed
- Introducing `v1alpha3` CR's.



### external-dns [2.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.5.0)

#### Changed
- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.



### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

#### Added
- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.
#### Changed
- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.



### cert-manager [2.11.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.11.0)

#### Changed
- Upgrade to upstream `v1.5.4` ([#191](https://github.com/giantswarm/cert-manager-app/pull/191)).
- Add metadata to enable metrics scraping ([#181](https://github.com/giantswarm/cert-manager-app/pull/181)).
- Fix startubjob PSP ([#191](https://github.com/giantswarm/cert-manager-app/pull/191))
- Upgrade to upstream `v1.5.3` ([#184](https://github.com/giantswarm/cert-manager-app/pull/184)). This is the first version compatible with Kubernetes 1.22.
- Add metadata to enable metrics scraping ([#181](https://github.com/giantswarm/cert-manager-app/pull/181)).
- Update to upstream `v1.4.2` ([#174](https://github.com/giantswarm/cert-manager-app/pull/174)). This deprecates `v1alpha2`, `v1alpha3` and `v1beta1` versions of `cert-manager.io` and `acme.cert-manager.io` CRDs. Further information can be found in the [upstream release notes](https://cert-manager.io/docs/release-notes/release-notes-1.4/) of cert-manager.
- Increase resource requests for the ClusterIssuer and CRD installation Jobs ([#174](https://github.com/giantswarm/cert-manager-app/pull/174)) to prevent timeouts.



### chart-operator [2.19.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.19.0)

#### Removed
- Remove `tillermigration` resource now Helm 3 migration is complete.

#### Changed
- Increase memory limit for deploying large charts in workload clusters.



### cluster-autoscaler [1.21.0-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.0-gs2)

#### Changed
- Fix RBAC for cluster autoscaler 1.21.
- Updated cluster-autoscaler to version `1.21.0`.
- Use new node selector `node-role.kubernetes.io/master` in place of deprecated one `kubernetes.io/role`.
- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.

#### Added
- Add `VerticalPodAutoscaler` resource to adjust limits automatically.



### external-dns [2.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.5.0)

### Changed
- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.



### kube-state-metrics [1.4.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.4.0)

#### Changed
- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.



### metrics-server [1.5.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.5.0)

#### Changed
- Bumped API version for `RoleBinding` to `v1` as it was using a deprecated version (removed in 1.22).



### net-exporter [1.10.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.3)

#### Changed
- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.



### node-exporter [1.8.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.8.0)

#### Changed
- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.



### aws-ebs-csi-driver [2.3.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.3.1)

#### Fixed
- Enable permissions for ebs volume resizing by default.

#### Changed
- Bump aws-ebs-csi-driver version to `v1.2.0`
