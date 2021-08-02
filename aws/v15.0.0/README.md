# :zap: Giant Swarm Release v15.0.0 for AWS :zap:

This release provides support for Kubernetes 1.20. It also enables the Container Storage Interface (CSI) and the automatic termination of unhealthy nodes by default.

**Highlights**
- Kubernetes 1.20 support;
- Container Storage Interface (CSI) enabled by default;
- `gp3` is now the default storage class instead of `gp2`;
- Automatic termination of unhealthy nodes enabled by default;
- Security fixes:
    * 49 Linux CVEs;
    * 5 openssl CVEs;
    * 5 nvidia-drivers CVEs;
    * 1 runc CVE;
    * 1 containerd CVE;
    * 1 Kubernetes CVE.

> **_Warning:_**  From AWS workload cluster release v15.0.0, the automatic termination of unhealthy nodes is enabled by default. For more information about the feature and information how to disable it, please follow the [official documentation](https://docs.giantswarm.io/advanced/automatic-node-termination/). 

## Change details


### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases;
- Emit events when config maps and secrets referenced in App CRs are updated;
- Cache k8sclient, helmclient for later use;
- Apply the namespaceConfig to the desired chart;
- Install apps in CAPI Workload Clusters;
- Apply `compatibleProvider`, `namespace` metadata validation based on the relevant `AppCatalogEntry` CR;
- Add annotations from Helm charts to AppCatalogEntry CRs;
- Enable Vertical Pod Autoscaler.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support;
- Restore chart-operator when it had been deleted;
- Use backoff in chart CR watcher to wait until kubeconfig secret exists;

#### Changed
- Updated Helm to v3.5.3;
- Replace status webhook with chart CR status watcher;
- Sort AppCatalogEntry CRs by version and created timestamp;
- Watch cluster namespace for per workload cluster instances of app-operator.



### cluster-operator [3.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.8.0)

#### Changed
- Adjust helm chart to be used with `config-controller`.

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support;
- Fix `clusterIPRange` value in configmap;
- Fix `kubeconfig` resource to search secrets in all namespaces.



### aws-operator [10.6.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.6.1)

#### Added
- S3 vpc endpoint to AWS CNI subnet;
- Enabled EBS CSI migration.

#### Changed
- Upgrade `k8scloudconfig` to v10.8.1 which includes a change to better determine if memory eviction thresholds are crossed;
- Update Flatcar AMI's to the latest stable releases;
- Enabled EBS CSI migration;
- Avoid TCCPN stack failure by checking if a control-plane tag exists before adding it;
- Look up cloud tags in all namespaces;
- Find certs in all namespaces;
- Enable `terminate unhealthy node` feature by default;
- Add node termination counter per cluster metric.

#### Removed
- Removed default storage-class annotation, EBS CSI driver is taking over.



### containerlinux [2765.2.6](https://www.flatcar-linux.org/releases/#release-2765.2.6)


**Security fixes**

* Linux
    - [CVE-2020-26558](https://nvd.nist.gov/vuln/detail/CVE-2020-26558)
    - [CVE-2021-0129](https://nvd.nist.gov/vuln/detail/CVE-2021-0129)
    - [CVE-2020-24587](https://nvd.nist.gov/vuln/detail/CVE-2020-24587)
    - [CVE-2020-24586](https://nvd.nist.gov/vuln/detail/CVE-2020-24586)
    - [CVE-2020-24588](https://nvd.nist.gov/vuln/detail/CVE-2020-24588)
    - [CVE-2020-26139](https://nvd.nist.gov/vuln/detail/CVE-2020-26139)
    - [CVE-2020-26145](https://nvd.nist.gov/vuln/detail/CVE-2020-26145)
    - [CVE-2020-26147](https://nvd.nist.gov/vuln/detail/CVE-2020-26147)
    - [CVE-2020-26141](https://nvd.nist.gov/vuln/detail/CVE-2020-26141)
    - [CVE-2021-3564](https://nvd.nist.gov/vuln/detail/CVE-2021-3564)
    - [CVE-2021-28691](https://nvd.nist.gov/vuln/detail/CVE-2021-28691)
    - [CVE-2021-3587](https://nvd.nist.gov/vuln/detail/CVE-2021-3587)
    - [CVE-2021-3573](https://nvd.nist.gov/vuln/detail/CVE-2021-3573)
    - [CVE-2021-3491](https://nvd.nist.gov/vuln/detail/CVE-2021-3491)
    - [CVE-2021-31440](https://nvd.nist.gov/vuln/detail/CVE-2021-31440)
    - [CVE-2021-31829](https://nvd.nist.gov/vuln/detail/CVE-2021-31829)
    - [CVE-2021-28964](https://nvd.nist.gov/vuln/detail/CVE-2021-28964)
    - [CVE-2021-28972](https://nvd.nist.gov/vuln/detail/CVE-2021-28972)
    - [CVE-2021-28971](https://nvd.nist.gov/vuln/detail/CVE-2021-28971)
    - [CVE-2021-28951](https://nvd.nist.gov/vuln/detail/CVE-2021-28951)
    - [CVE-2021-28952](https://nvd.nist.gov/vuln/detail/CVE-2021-28952)
    - [CVE-2021-29266](https://nvd.nist.gov/vuln/detail/CVE-2021-29266)
    - [CVE-2021-28688](https://nvd.nist.gov/vuln/detail/CVE-2021-28688)
    - [CVE-2021-29264](https://nvd.nist.gov/vuln/detail/CVE-2021-29264)
    - [CVE-2021-29649](https://nvd.nist.gov/vuln/detail/CVE-2021-29649)
    - [CVE-2021-29650](https://nvd.nist.gov/vuln/detail/CVE-2021-29650)
    - [CVE-2021-29646](https://nvd.nist.gov/vuln/detail/CVE-2021-29646)
    - [CVE-2021-29647](https://nvd.nist.gov/vuln/detail/CVE-2021-29647)
    - [CVE-2021-29154](https://nvd.nist.gov/vuln/detail/CVE-2021-29154)
    - [CVE-2021-29155](https://nvd.nist.gov/vuln/detail/CVE-2021-29155)
    - [CVE-2021-23133](https://nvd.nist.gov/vuln/detail/CVE-2021-23133)
    - [CVE-2021-27365](https://nvd.nist.gov/vuln/detail/CVE-2021-27365)
    - [CVE-2021-27364](https://nvd.nist.gov/vuln/detail/CVE-2021-27364)
    - [CVE-2021-27363](https://nvd.nist.gov/vuln/detail/CVE-2021-27363)
    - [CVE-2021-28038](https://nvd.nist.gov/vuln/detail/CVE-2021-28038)
    - [CVE-2021-28039](https://nvd.nist.gov/vuln/detail/CVE-2021-28039)
    - [CVE-2021-28375](https://nvd.nist.gov/vuln/detail/CVE-2021-28375)
    - [CVE-2021-28660](https://nvd.nist.gov/vuln/detail/CVE-2021-28660)
    - [CVE-2021-27218](https://nvd.nist.gov/vuln/detail/CVE-2021-27218)
    - [CVE-2021-27219](https://nvd.nist.gov/vuln/detail/CVE-2021-27219)
    - [CVE-2020-25639](https://nvd.nist.gov/vuln/detail/CVE-2020-25639)
    - [CVE-2021-27365](https://nvd.nist.gov/vuln/detail/CVE-2021-27365)
    - [CVE-2021-27364](https://nvd.nist.gov/vuln/detail/CVE-2021-27364)
    - [CVE-2021-27363](https://nvd.nist.gov/vuln/detail/CVE-2021-27363)
    - [CVE-2021-28038](https://nvd.nist.gov/vuln/detail/CVE-2021-28038)
    - [CVE-2021-28039](https://nvd.nist.gov/vuln/detail/CVE-2021-28039)
    - [CVE-2021-26931](https://nvd.nist.gov/vuln/detail/CVE-2021-26931)
    - [CVE-2021-26930](https://nvd.nist.gov/vuln/detail/CVE-2021-26930)
    - [CVE-2021-26932](https://nvd.nist.gov/vuln/detail/CVE-2021-26932)
* openssl
    - [CVE-2021-23840](https://nvd.nist.gov/vuln/detail/CVE-2021-23840)
    - [CVE-2021-23841](https://nvd.nist.gov/vuln/detail/CVE-2021-23841)
    - [CVE-2020-1971](https://nvd.nist.gov/vuln/detail/CVE-2020-1971)
    - [CVE-2021-3449](https://nvd.nist.gov/vuln/detail/CVE-2021-3449)
    - [CVE-2021-3450](https://nvd.nist.gov/vuln/detail/CVE-2021-3450)
* nvidia-drivers
    - [CVE-2021-1052](https://nvd.nist.gov/vuln/detail/CVE-2021-1052)
    - [CVE-2021-1053](https://nvd.nist.gov/vuln/detail/CVE-2021-1053)
    - [CVE-2021-1056](https://nvd.nist.gov/vuln/detail/CVE-2021-1056)
    - [CVE-2021-1076](https://nvd.nist.gov/vuln/detail/CVE-2021-1076)
    - [CVE-2021-1077](https://nvd.nist.gov/vuln/detail/CVE-2021-1077)
* runc
    - [CVE-2021-30465](https://nvd.nist.gov/vuln/detail/CVE-2021-30465)
* containerd
    - [GHSA-6g2q-w5j3-fwh4](https://github.com/containerd/containerd/security/advisories/GHSA-6g2q-w5j3-fwh4)


**Bug fixes**

* Update-engine sent empty requests when restarted before a pending reboot ([Flatcar#388](https://github.com/kinvolk/Flatcar/issues/388))
motd login prompt list of failed services: The output of “systemctl list-units –state=failed –no-legend” contains a bullet point which is not expected and ended up being taken as the unit name of failed units which was previously on the start of the line. Filtered the bullet point out to stay compatible with the old behavior in case upstream would remove the bullet point again. ([coreos-overlay#1042](https://github.com/kinvolk/coreos-overlay/pull/1042))
* The Linux kernel IOMMU-related crash introduced in the 5.10.37 update got fixed through the 5.10.38 update ([Flatcar#400](https://github.com/kinvolk/Flatcar/issues/400))
* Fix the patch to update DefaultTasksMax in systemd ([coreos-overlay#971](https://github.com/kinvolk/coreos-overlay/pull/971))
* GCE: The old interface name ens4v1 which was replaced by eth0 due to a broken udev rule was restored, but now as alternative interface name, and eth0 will stay the primary name for consistency across cloud environments. ([init#38](https://github.com/kinvolk/init/pull/38))

**Changes**

* The virtio network interfaces got predictable interface names as alternative interface names, and thus these names can also be used to match for a specific interface in case there is more than one and the eth0 and eth1 name assignment is not stable. ([init#38](https://github.com/kinvolk/init/pull/38))



**Updates**

* Linux ([5.10.43](https://lwn.net/Articles/859022/))
* nvidia-drivers ([460.73.01](https://www.nvidia.com/Download/driverResults.aspx/172376/en-us))
* openssl ([1.1.1k](https://mta.openssl.org/pipermail/openssl-announce/2021-March/000197.html))
* open-iscsi ([2.1.4](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.4))


### kubernetes [1.20.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.8)

#### What's New (Major Themes)

##### Dockershim deprecation

Docker as an underlying runtime is being deprecated. Docker-produced images will continue to work in your cluster with all runtimes, as they always have.
The Kubernetes community [has written a blog post about this in detail](https://blog.k8s.io/2020/12/02/dont-panic-kubernetes-and-docker/) with [a dedicated FAQ page for it](https://blog.k8s.io/2020/12/02/dockershim-faq/).

##### External credential provider for client-go

The client-go credential plugins can now be passed in the current cluster information via the `KUBERNETES_EXEC_INFO` environment variable. Learn more about this on [client-go credential plugins documentation](https://docs.k8s.io/reference/access-authn-authz/authentication/#client-go-credential-plugins/).

##### CronJob controller v2 is available through feature gate

An alternative implementation of the `CronJob` controller is now available as an alpha feature in this release, which has experimental performance improvement by using informers instead of polling. While this will be the default behavior in the future, you can [try them in this release through a feature gate](https://docs.k8s.io/concepts/workloads/controllers/cron-jobs/).

##### PID Limits graduates to General Availability

PID Limits features are now generally available on both `SupportNodePidsLimit` (node-to-pod PID isolation) and `SupportPodPidsLimit` (ability to limit PIDs per pod), after being enabled-by-default in beta stage for a year.

##### API Priority and Fairness graduates to Beta

Initially introduced in 1.18, Kubernetes 1.20 now enables API Priority and Fairness (APF) by default. This allows `kube-apiserver` to [categorize incoming requests by priority levels](https://docs.k8s.io/concepts/cluster-administration/flow-control/).

##### IPv4/IPv6 run

IPv4/IPv6 dual-stack has been reimplemented for 1.20 to support dual-stack Services, based on user and community feedback. If your cluster has dual-stack enabled, you can create Services which can use IPv4, IPv6, or both, and you can change this setting for existing Services. Details are available in updated [IPv4/IPv6 dual-stack docs](https://docs.k8s.io/concepts/services-networking/dual-stack/), which cover the nuanced array of options.

We expect this implementation to progress from alpha to beta and GA in coming releases, so we’re eager to have you comment about your dual-stack experiences in [#k8s-dual-stack](https://kubernetes.slack.com/messages/k8s-dual-stack) or in [enhancements #563](https://features.k8s.io/563). 

##### go1.15.5

go1.15.5 has been integrated into the Kubernetes project as of this release, [including other infrastructure related updates on this effort](https://github.com/kubernetes/kubernetes/pull/95776).

##### CSI Volume Snapshot graduates to General Availability

CSI Volume Snapshot moves to GA in the 1.20 release. This feature provides a standard way to trigger volume snapshot operations in Kubernetes and allows Kubernetes users to incorporate snapshot operations in a portable manner on any Kubernetes environment regardless of supporting underlying storage providers.
Additionally, these Kubernetes snapshot primitives act as basic building blocks that unlock the ability to develop advanced, enterprise-grade, storage administration features for Kubernetes: including application or cluster level backup solutions.
Note that snapshot support will require Kubernetes distributors to bundle the Snapshot controller, Snapshot CRDs, and validation webhook. In addition, a CSI driver supporting the snapshot functionality must also be deployed on the cluster.

##### Non-recursive Volume Ownership (FSGroup) graduates to Beta

By default, the `fsgroup` setting, if specified, recursively updates permissions for every file in a volume on every mount. This can make mount, and pod startup, very slow if the volume has many files. 
This setting enables a pod to specify a `PodFSGroupChangePolicy` that indicates that volume ownership and permissions will be changed only when permission and ownership of the root directory do not match with expected permissions on the volume.

##### CSIDriver policy for FSGroup graduates to Beta

The FSGroup's CSIDriver Policy is now beta in 1.20. This allows CSIDrivers to explicitly indicate if they want Kubernetes to manage permissions and ownership for their volumes via `fsgroup`. 

##### Security Improvements for CSI Drivers (Alpha)

In 1.20, we introduce a new alpha feature `CSIServiceAccountToken`. This feature allows CSI drivers to impersonate the pods that they mount the volumes for. This improves the security posture in the mounting process where the volumes are ACL’ed on the pods’ service account without handing out unnecessary permissions to the CSI drivers’ service account. This feature is especially important for secret-handling CSI drivers, such as the secrets-store-csi-driver. Since these tokens can be rotated and short-lived, this feature also provides a knob for CSI drivers to receive `NodePublishVolume` RPC calls periodically with the new token. This knob is also useful when volumes are short-lived, e.g. certificates.

##### Introducing Graceful Node Shutdown (Alpha)

The `GracefulNodeShutdown` feature is now in Alpha. This allows kubelet to be aware of node system shutdowns, enabling graceful termination of pods during a system shutdown. This feature can be [enabled through feature gate](https://docs.k8s.io/concepts/architecture/nodes/#graceful-node-shutdown).

##### Runtime log sanitation

Logs can now be configured to use runtime protection from leaking sensitive data. [Details for this experimental feature is available in documentation](https://docs.k8s.io/concepts/cluster-administration/system-logs/#log-sanitization).

##### Pod resource metrics

On-demand metrics calculation is now available through `/metrics/resources`. [When enabled](
https://docs.k8s.io/concepts/cluster-administration/system-metrics#kube-scheduler-metrics), the endpoint will report the requested resources and the desired limits of all running pods.

##### Introducing `RootCAConfigMap`

`RootCAConfigMap` graduates to Beta, separating from `BoundServiceAccountTokenVolume`. The `kube-root-ca.crt` ConfigMap is now available to every namespace, by default. It contains the Certificate Authority bundle for verify kube-apiserver connections.

##### `kubectl debug` graduates to Beta

`kubectl alpha debug` graduates from alpha to beta in 1.20, becoming `kubectl debug`.
`kubectl debug` provides support for common debugging workflows directly from kubectl. Troubleshooting scenarios supported in this release of `kubectl` include:
Troubleshoot workloads that crash on startup by creating a copy of the pod that uses a different container image or command.
Troubleshoot distroless containers by adding a new container with debugging tools, either in a new copy of the pod or using an ephemeral container. (Ephemeral containers are an alpha feature that are not enabled by default.)
Troubleshoot on a node by creating a container running in the host namespaces and with access to the host’s filesystem.
Note that as a new builtin command, `kubectl debug` takes priority over any `kubectl` plugin named “debug”. You will need to rename the affected plugin.
Invocations using `kubectl alpha debug` are now deprecated and will be removed in a subsequent release. Update your scripts to use `kubectl debug` instead of `kubectl alpha debug`!
For more information about kubectl debug, see [Debugging Running Pods](https://kubernetes.io/docs/tasks/debug-application-cluster/debug-running-pod/) on the Kubernetes website, kubectl help debug, or reach out to SIG CLI by visiting [#sig-cli](https://kubernetes.slack.com/messages/sig-cli) or commenting on [enhancement #1441](https://features.k8s.io/1441).

##### Removing deprecated flags in kubeadm

`kubeadm` applies a number of deprecations and removals of deprecated features in this release. More details are available in the Urgent Upgrade Notes and Kind / Deprecation sections.

##### Pod Hostname as FQDN graduates to Beta

Previously introduced in 1.19 behind a feature gate, `SetHostnameAsFQDN` is now enabled by default. More details on this behavior are available in [documentation for DNS for Services and Pods](https://docs.k8s.io/concepts/services-networking/dns-pod-service/#pod-sethostnameasfqdn-field)

##### `TokenRequest` / `TokenRequestProjection` graduates to General Availability

Service account tokens bound to a pod is now a stable feature. The feature gates will be removed in 1.21 release. For more information, refer to notes below on the changelogs.

##### RuntimeClass feature graduates to General Availability.

The `node.k8s.io` API groups are promoted from `v1beta1` to `v1`. `v1beta1` is now deprecated and will be removed in a future release, please start using `v1`. ([#95718](https://github.com/kubernetes/kubernetes/pull/95718), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Apps, Auth, Node, Scheduling and Testing]

##### Cloud Controller Manager now exclusively shipped by Cloud Provider

Kubernetes will no longer ship an instance of the Cloud Controller Manager binary. Each Cloud Provider is expected to ship their own instance of this binary. Details for a Cloud Provider to create an instance of such a binary can be found [here](https://github.com/kubernetes/kubernetes/tree/master/staging/src/k8s.io/cloud-provider/sample). Anyone with questions on building a Cloud Controller Manager should reach out to SIG Cloud Provider. Questions about the Cloud Controller Manager on a Managed Kubernetes solution should go to the relevant Cloud Provider. Questions about the Cloud Controller Manager on a non managed solution can be brought up with SIG Cloud Provider.



#### Important Security Information

This release contains changes that address the following vulnerabilities:

##### CVE-2021-25735: Validating Admission Webhook does not observe some previous fields

A security issue was discovered in kube-apiserver that could allow node
updates to bypass a Validating Admission Webhook. You are only affected
by this vulnerability if you run a Validating Admission Webhook for Nodes
that denies admission based at least partially on the old state of the
Node object.

**Note**: This only impacts validating admission plugins that rely on old
values in certain fields, and does not impact calls from kubelet that go
through the built-in NodeRestriction admission plugin.

**Affected Versions**:
  - kube-apiserver v1.20.0 - v1.20.5
  - kube-apiserver v1.19.0 - v1.19.9
  - kube-apiserver <= v1.18.17

**Fixed Versions**:
  - kube-apiserver v1.21.0
  - kube-apiserver v1.20.6
  - kube-apiserver v1.19.10
  - kube-apiserver v1.18.18

This vulnerability was reported by Rogerio Bastos & Ari Lima from RedHat


**CVSS Rating:** Medium (6.5) [CVSS:3.0/AV:N/AC:L/PR:H/UI:N/S:U/C:N/I:H/A:H](https://www.first.org/cvss/calculator/3.0#CVSS:3.0/AV:N/AC:L/PR:H/UI:N/S:U/C:N/I:H/A:H)



#### Deprecation

- Docker support in the kubelet is now deprecated and will be removed in a future release. The kubelet uses a module called "dockershim" which implements CRI support for Docker and it has seen maintenance issues in the Kubernetes community. We encourage you to evaluate moving to a container runtime that is a full-fledged implementation of CRI (v1alpha1 or v1 compliant) as they become available. ([#94624](https://github.com/kubernetes/kubernetes/pull/94624), [@dims](https://github.com/dims)) [SIG Node]
- Kubeadm: deprecate self-hosting support. The experimental command "kubeadm alpha self-hosting" is now deprecated and will be removed in a future release. ([#95125](https://github.com/kubernetes/kubernetes/pull/95125), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: graduate the "kubeadm alpha certs" command to a parent command "kubeadm certs". The command "kubeadm alpha certs" is deprecated and will be removed in a future release. Please migrate. ([#94938](https://github.com/kubernetes/kubernetes/pull/94938), [@yagonobre](https://github.com/yagonobre)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated "kubeadm alpha kubelet config enable-dynamic" command. To continue using the feature please defer to the guide for "Dynamic Kubelet Configuration" at k8s.io. This change also removes the parent command "kubeadm alpha kubelet" as there are no more sub-commands under it for the time being. ([#94668](https://github.com/kubernetes/kubernetes/pull/94668), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: remove the deprecated --kubelet-config flag for the command "kubeadm upgrade node" ([#94869](https://github.com/kubernetes/kubernetes/pull/94869), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubectl: deprecate --delete-local-data ([#95076](https://github.com/kubernetes/kubernetes/pull/95076), [@dougsland](https://github.com/dougsland)) [SIG CLI, Cloud Provider and Scalability]
- Kubelet's deprecated endpoint `metrics/resource/v1alpha1` has been removed, please adopt `metrics/resource`. ([#94272](https://github.com/kubernetes/kubernetes/pull/94272), [@RainbowMango](https://github.com/RainbowMango)) [SIG Instrumentation and Node]
- Removes deprecated scheduler metrics DeprecatedSchedulingDuration, DeprecatedSchedulingAlgorithmPredicateEvaluationSecondsDuration, DeprecatedSchedulingAlgorithmPriorityEvaluationSecondsDuration ([#94884](https://github.com/kubernetes/kubernetes/pull/94884), [@arghya88](https://github.com/arghya88)) [SIG Instrumentation and Scheduling]
- Scheduler alpha metrics binding_duration_seconds and scheduling_algorithm_preemption_evaluation_seconds are deprecated, Both of those metrics are now covered as part of framework_extension_point_duration_seconds, the former as a PostFilter the latter and a Bind plugin. The plan is to remove both in 1.21 ([#95001](https://github.com/kubernetes/kubernetes/pull/95001), [@arghya88](https://github.com/arghya88)) [SIG Instrumentation and Scheduling]
- Support `controlplane` as a valid EgressSelection type in the EgressSelectorConfiguration API. `Master` is deprecated and will be removed in v1.22. ([#95235](https://github.com/kubernetes/kubernetes/pull/95235), [@andrewsykim](https://github.com/andrewsykim)) [SIG API Machinery]
- The v1alpha1 PodPreset API and admission plugin has been removed with no built-in replacement. Admission webhooks can be used to modify pods on creation. ([#94090](https://github.com/kubernetes/kubernetes/pull/94090), [@deads2k](https://github.com/deads2k)) [SIG API Machinery, Apps, CLI, Cloud Provider, Scalability and Testing]



#### API Change
- We have added a new Priority & Fairness rule that exempts all probes (/readyz, /healthz, /livez) to prevent 
  restarting of "healthy" kube-apiserver instance(s) by kubelet. ([#101112](https://github.com/kubernetes/kubernetes/pull/101112), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fixes using server-side apply with APIService resources ([#100714](https://github.com/kubernetes/kubernetes/pull/100714), [@kevindelgado](https://github.com/kevindelgado)) [SIG API Machinery, Apps and Testing]
- Regenerate protobuf code to fix CVE-2021-3121 ([#100501](https://github.com/kubernetes/kubernetes/pull/100501), [@joelsmith](https://github.com/joelsmith)) [SIG API Machinery, Apps, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
- Kubernetes is now built using go1.15.8 ([#98962](https://github.com/kubernetes/kubernetes/pull/98962), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- `TokenRequest` and `TokenRequestProjection` features have been promoted to GA. This feature allows generating service account tokens that are not visible in Secret objects and are tied to the lifetime of a Pod object. See https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#service-account-token-volume-projection for details on configuring and using this feature. The `TokenRequest` and `TokenRequestProjection` feature gates will be removed in v1.21.
  - kubeadm's kube-apiserver Pod manifest now includes the following flags by default "--service-account-key-file", "--service-account-signing-key-file", "--service-account-issuer". ([#93258](https://github.com/kubernetes/kubernetes/pull/93258), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Auth, Cluster Lifecycle, Storage and Testing]
- A new `nofuzz` go build tag now disables gofuzz support. Release binaries enable this. ([#92491](https://github.com/kubernetes/kubernetes/pull/92491), [@BenTheElder](https://github.com/BenTheElder)) [SIG API Machinery]
- Add WindowsContainerResources and Annotations to CRI-API UpdateContainerResourcesRequest ([#95741](https://github.com/kubernetes/kubernetes/pull/95741), [@katiewasnothere](https://github.com/katiewasnothere)) [SIG Node]
- Add a `serving` and `terminating` condition to the EndpointSlice API.
  `serving` tracks the readiness of endpoints regardless of their terminating state. This is distinct from `ready` since `ready` is only true when pods are not terminating. 
  `terminating` is true when an endpoint is terminating. For pods this is any endpoint with a deletion timestamp. ([#92968](https://github.com/kubernetes/kubernetes/pull/92968), [@andrewsykim](https://github.com/andrewsykim)) [SIG Apps and Network]
- Add dual-stack Services (alpha).  This is a BREAKING CHANGE to an alpha API.
  It changes the dual-stack API wrt Service from a single ipFamily field to 3
  fields: ipFamilyPolicy (SingleStack, PreferDualStack, RequireDualStack),
  ipFamilies (a list of families assigned), and clusterIPs (inclusive of
  clusterIP).  Most users do not need to set anything at all, defaulting will
  handle it for them.  Services are single-stack unless the user asks for
  dual-stack.  This is all gated by the "IPv6DualStack" feature gate. ([#91824](https://github.com/kubernetes/kubernetes/pull/91824), [@khenidak](https://github.com/khenidak)) [SIG API Machinery, Apps, CLI, Network, Node, Scheduling and Testing]
- Add support for hugepages to downward API ([#86102](https://github.com/kubernetes/kubernetes/pull/86102), [@derekwaynecarr](https://github.com/derekwaynecarr)) [SIG API Machinery, Apps, CLI, Network, Node, Scheduling and Testing]
- Adds kubelet alpha feature, `GracefulNodeShutdown` which makes kubelet aware of node system shutdowns and result in graceful termination of pods during a system shutdown. ([#96129](https://github.com/kubernetes/kubernetes/pull/96129), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
- AppProtocol is now GA for Endpoints and Services. The ServiceAppProtocol feature gate will be deprecated in 1.21. ([#96327](https://github.com/kubernetes/kubernetes/pull/96327), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Automatic allocation of NodePorts for services with type LoadBalancer can now be disabled by setting the (new) parameter
  Service.spec.allocateLoadBalancerNodePorts=false. The default is to allocate NodePorts for services with type LoadBalancer which is the existing behavior. ([#92744](https://github.com/kubernetes/kubernetes/pull/92744), [@uablrek](https://github.com/uablrek)) [SIG Apps and Network]
- Certain fields on  Service objects will be automatically cleared when changing the service's `type` to a mode that does not need those fields.  For example, changing from type=LoadBalancer to type=ClusterIP will clear the NodePort assignments, rather than forcing the user to clear them. ([#95196](https://github.com/kubernetes/kubernetes/pull/95196), [@thockin](https://github.com/thockin)) [SIG API Machinery, Apps, Network and Testing]
- Document that ServiceTopology feature is required to use `service.spec.topologyKeys`. ([#96528](https://github.com/kubernetes/kubernetes/pull/96528), [@andrewsykim](https://github.com/andrewsykim)) [SIG Apps]
- EndpointSlice has a new NodeName field guarded by the EndpointSliceNodeName feature gate.
  - EndpointSlice topology field will be deprecated in an upcoming release.
  - EndpointSlice "IP" address type is formally removed after being deprecated in Kubernetes 1.17.
  - The discovery.k8s.io/v1alpha1 API is deprecated and will be removed in Kubernetes 1.21. ([#96440](https://github.com/kubernetes/kubernetes/pull/96440), [@robscott](https://github.com/robscott)) [SIG API Machinery, Apps and Network]
- External facing API podresources is now available under k8s.io/kubelet/pkg/apis/ ([#92632](https://github.com/kubernetes/kubernetes/pull/92632), [@RenaudWasTaken](https://github.com/RenaudWasTaken)) [SIG Node and Testing]
- Fewer candidates are enumerated for preemption to improve performance in large clusters. ([#94814](https://github.com/kubernetes/kubernetes/pull/94814), [@adtac](https://github.com/adtac))
- Fix conversions for custom metrics. ([#94481](https://github.com/kubernetes/kubernetes/pull/94481), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- GPU metrics provided by kubelet are now disabled by default. ([#95184](https://github.com/kubernetes/kubernetes/pull/95184), [@RenaudWasTaken](https://github.com/RenaudWasTaken))
- If BoundServiceAccountTokenVolume is enabled, cluster admins can use metric `serviceaccount_stale_tokens_total` to monitor workloads that are depending on the extended tokens. If there are no such workloads, turn off extended tokens by starting `kube-apiserver` with flag `--service-account-extend-token-expiration=false` ([#96273](https://github.com/kubernetes/kubernetes/pull/96273), [@zshihang](https://github.com/zshihang)) [SIG API Machinery and Auth]
- Introduce alpha support for exec-based container registry credential provider plugins in the kubelet. ([#94196](https://github.com/kubernetes/kubernetes/pull/94196), [@andrewsykim](https://github.com/andrewsykim)) [SIG Node and Release]
- Introduces a metric source for HPAs which allows scaling based on container resource usage. ([#90691](https://github.com/kubernetes/kubernetes/pull/90691), [@arjunrn](https://github.com/arjunrn)) [SIG API Machinery, Apps, Autoscaling and CLI]
- Kube-apiserver now deletes expired kube-apiserver Lease objects:
  - The feature is under feature gate `APIServerIdentity`.
  - A flag is added to kube-apiserver: `identity-lease-garbage-collection-check-period-seconds` ([#95895](https://github.com/kubernetes/kubernetes/pull/95895), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery, Apps, Auth and Testing]
- Kube-controller-manager: volume plugins can be restricted from contacting local and loopback addresses by setting `--volume-host-allow-local-loopback=false`, or from contacting specific CIDR ranges by setting `--volume-host-cidr-denylist` (for example, `--volume-host-cidr-denylist=127.0.0.1/28,feed::/16`) ([#91785](https://github.com/kubernetes/kubernetes/pull/91785), [@mattcary](https://github.com/mattcary)) [SIG API Machinery, Apps, Auth, CLI, Network, Node, Storage and Testing]
- Migrate scheduler, controller-manager and cloud-controller-manager to use LeaseLock ([#94603](https://github.com/kubernetes/kubernetes/pull/94603), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery, Apps, Cloud Provider and Scheduling]
- Modify DNS-1123 error messages to indicate that RFC 1123 is not followed exactly ([#94182](https://github.com/kubernetes/kubernetes/pull/94182), [@mattfenwick](https://github.com/mattfenwick)) [SIG API Machinery, Apps, Auth, Network and Node]
- Move configurable fsgroup change policy for pods to beta ([#96376](https://github.com/kubernetes/kubernetes/pull/96376), [@gnufied](https://github.com/gnufied)) [SIG Apps and Storage]
- New flag is introduced, i.e. --topology-manager-scope=container|pod. 
  The default value is the "container" scope. ([#92967](https://github.com/kubernetes/kubernetes/pull/92967), [@cezaryzukowski](https://github.com/cezaryzukowski)) [SIG Instrumentation, Node and Testing]
- New parameter `defaultingType` for `PodTopologySpread` plugin allows to use k8s defined or user provided default constraints ([#95048](https://github.com/kubernetes/kubernetes/pull/95048), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- NodeAffinity plugin can be configured with AddedAffinity. ([#96202](https://github.com/kubernetes/kubernetes/pull/96202), [@alculquicondor](https://github.com/alculquicondor)) [SIG Node, Scheduling and Testing]
- Promote RuntimeClass feature to GA.
  Promote node.k8s.io API groups from v1beta1 to v1. ([#95718](https://github.com/kubernetes/kubernetes/pull/95718), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Apps, Auth, Node, Scheduling and Testing]
- Reminder: The labels "failure-domain.beta.kubernetes.io/zone" and "failure-domain.beta.kubernetes.io/region" are deprecated in favor of "topology.kubernetes.io/zone" and "topology.kubernetes.io/region" respectively.  All users of the "failure-domain.beta..." labels should switch to the "topology..." equivalents. ([#96033](https://github.com/kubernetes/kubernetes/pull/96033), [@thockin](https://github.com/thockin)) [SIG API Machinery, Apps, CLI, Cloud Provider, Network, Node, Scheduling, Storage and Testing]
- Server Side Apply now treats LabelSelector fields as atomic (meaning the entire selector is managed by a single writer and updated together), since they contain interrelated and inseparable fields that do not merge in intuitive ways. ([#93901](https://github.com/kubernetes/kubernetes/pull/93901), [@jpbetz](https://github.com/jpbetz)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Network, Node, Storage and Testing]
- Services will now have a `clusterIPs` field to go with `clusterIP`.  `clusterIPs[0]` is a synonym for `clusterIP` and will be synchronized on create and update operations. ([#95894](https://github.com/kubernetes/kubernetes/pull/95894), [@thockin](https://github.com/thockin)) [SIG Network]
- The ServiceAccountIssuerDiscovery feature gate is now Beta and enabled by default. ([#91921](https://github.com/kubernetes/kubernetes/pull/91921), [@mtaufen](https://github.com/mtaufen)) [SIG Auth]
- The status of v1beta1 CRDs without "preserveUnknownFields:false" now shows a violation, "spec.preserveUnknownFields: Invalid value: true: must be false". ([#93078](https://github.com/kubernetes/kubernetes/pull/93078), [@vareti](https://github.com/vareti))
- The usage of mixed protocol values in the same LoadBalancer Service is possible if the new feature gate MixedProtocolLBService is enabled. The feature gate is disabled by default. The user has to enable it for the API Server. ([#94028](https://github.com/kubernetes/kubernetes/pull/94028), [@janosi](https://github.com/janosi)) [SIG API Machinery and Apps]
- This PR will introduce a feature gate CSIServiceAccountToken with two additional fields in `CSIDriverSpec`. ([#93130](https://github.com/kubernetes/kubernetes/pull/93130), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Apps, Auth, CLI, Network, Node, Storage and Testing]
- Users can try the CronJob controller v2 using the feature gate. This will be the default controller in future releases. ([#93370](https://github.com/kubernetes/kubernetes/pull/93370), [@alaypatel07](https://github.com/alaypatel07)) [SIG API Machinery, Apps, Auth and Testing]
- VolumeSnapshotDataSource moves to GA in 1.20 release ([#95282](https://github.com/kubernetes/kubernetes/pull/95282), [@xing-yang](https://github.com/xing-yang)) [SIG Apps]
- WinOverlay feature graduated to beta ([#94807](https://github.com/kubernetes/kubernetes/pull/94807), [@ksubrmnn](https://github.com/ksubrmnn)) [SIG Windows]



#### Feature

- Kubernetes is now built using Go 1.15.13 ([#102786](https://github.com/kubernetes/kubernetes/pull/102786), [@thejoycekung](https://github.com/thejoycekung)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built using go1.15.11 ([#101192](https://github.com/kubernetes/kubernetes/pull/101192), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built using go1.15.12 ([#101845](https://github.com/kubernetes/kubernetes/pull/101845), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- AWS cloudprovider supports auto-discovering subnets without any kubernetes.io/cluster/<clusterName> tags. It also supports additional service annotation service.beta.kubernetes.io/aws-load-balancer-subnets to manually configure the subnets. ([#97431](https://github.com/kubernetes/kubernetes/pull/97431), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- Kubernetes is now built using go1.15.10 ([#100375](https://github.com/kubernetes/kubernetes/pull/100375), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- A new metric `apiserver_request_filter_duration_seconds` has been introduced that 
  measures request filter latency in seconds. ([#95207](https://github.com/kubernetes/kubernetes/pull/95207), [@tkashem](https://github.com/tkashem)) [SIG API Machinery and Instrumentation]
- A new set of alpha metrics are reported by the Kubernetes scheduler under the `/metrics/resources` endpoint that allow administrators to easily see the resource consumption (requests and limits for all resources on the pods) and compare it to actual pod usage or node capacity. ([#94866](https://github.com/kubernetes/kubernetes/pull/94866), [@smarterclayton](https://github.com/smarterclayton)) [SIG API Machinery, Instrumentation, Node and Scheduling]
- Add --experimental-logging-sanitization flag enabling runtime protection from leaking sensitive data in logs ([#96370](https://github.com/kubernetes/kubernetes/pull/96370), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Add a StorageVersionAPI feature gate that makes API server update storageversions before serving certain write requests. 
  This feature allows the storage migrator to manage storage migration for built-in resources. 
  Enabling internal.apiserver.k8s.io/v1alpha1 API and APIServerIdentity feature gate are required to use this feature. ([#93873](https://github.com/kubernetes/kubernetes/pull/93873), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery, Auth and Testing]
- Add a metric for time taken to perform recursive permission change ([#95866](https://github.com/kubernetes/kubernetes/pull/95866), [@JornShen](https://github.com/JornShen)) [SIG Instrumentation and Storage]
- Add a new `vSphere` metric: `cloudprovider_vsphere_vcenter_versions`. Its content shows `vCenter` hostnames with the associated server version. ([#94526](https://github.com/kubernetes/kubernetes/pull/94526), [@Danil-Grigorev](https://github.com/Danil-Grigorev)) [SIG Cloud Provider and Instrumentation]
- Add a new flag to set priority for the kubelet on Windows nodes so that workloads cannot overwhelm the node thereby disrupting kubelet process. ([#96051](https://github.com/kubernetes/kubernetes/pull/96051), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node and Windows]
- Add feature to size memory backed volumes ([#94444](https://github.com/kubernetes/kubernetes/pull/94444), [@derekwaynecarr](https://github.com/derekwaynecarr)) [SIG Storage and Testing]
- Add foreground cascading deletion to kubectl with the new `kubectl delete foreground|background|orphan` option. ([#93384](https://github.com/kubernetes/kubernetes/pull/93384), [@zhouya0](https://github.com/zhouya0))
- Add metrics for azure service operations (route and loadbalancer). ([#94124](https://github.com/kubernetes/kubernetes/pull/94124), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider and Instrumentation]
- Add network rule support in Azure account creation. ([#94239](https://github.com/kubernetes/kubernetes/pull/94239), [@andyzhangx](https://github.com/andyzhangx))
- Add node_authorizer_actions_duration_seconds metric that can be used to estimate load to node authorizer. ([#92466](https://github.com/kubernetes/kubernetes/pull/92466), [@mborsz](https://github.com/mborsz)) [SIG API Machinery, Auth and Instrumentation]
- Add pod_ based CPU and memory metrics to Kubelet's  /metrics/resource endpoint ([#95839](https://github.com/kubernetes/kubernetes/pull/95839), [@egernst](https://github.com/egernst)) [SIG Instrumentation, Node and Testing]
- Added `get-users` and `delete-user` to the `kubectl config` subcommand ([#89840](https://github.com/kubernetes/kubernetes/pull/89840), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- Added counter metric "apiserver_request_self" to count API server self-requests with labels for verb, resource, and subresource. ([#94288](https://github.com/kubernetes/kubernetes/pull/94288), [@LogicalShark](https://github.com/LogicalShark)) [SIG API Machinery, Auth, Instrumentation and Scheduling]
- Added new k8s.io/component-helpers repository providing shared helper code for (core) components. ([#92507](https://github.com/kubernetes/kubernetes/pull/92507), [@ingvagabund](https://github.com/ingvagabund)) [SIG Apps, Node, Release and Scheduling]
- Adds `create ingress` command to `kubectl` ([#78153](https://github.com/kubernetes/kubernetes/pull/78153), [@amimof](https://github.com/amimof)) [SIG CLI and Network]
- Adds a headless service on node-local-cache addon. ([#88412](https://github.com/kubernetes/kubernetes/pull/88412), [@stafot](https://github.com/stafot)) [SIG Cloud Provider and Network]
- Allow cross-compilation of kubernetes on different platforms. ([#94403](https://github.com/kubernetes/kubernetes/pull/94403), [@bnrjee](https://github.com/bnrjee)) [SIG Release]
- Azure: Support multiple services sharing one IP address ([#94991](https://github.com/kubernetes/kubernetes/pull/94991), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- CRDs: For structural schemas, non-nullable null map fields will now be dropped and defaulted if a default is available. null items in the list will continue being preserved, and fail validation if not nullable. ([#95423](https://github.com/kubernetes/kubernetes/pull/95423), [@apelisse](https://github.com/apelisse)) [SIG API Machinery]
- Changed: default "Accept: */*" header added to HTTP probes. See https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#http-probes (https://github.com/kubernetes/website/pull/24756) ([#95641](https://github.com/kubernetes/kubernetes/pull/95641), [@fonsecas72](https://github.com/fonsecas72)) [SIG Network and Node]
- Client-go credential plugins can now be passed in the current cluster information via the KUBERNETES_EXEC_INFO environment variable. ([#95489](https://github.com/kubernetes/kubernetes/pull/95489), [@ankeesler](https://github.com/ankeesler)) [SIG API Machinery and Auth]
- Command to start network proxy changes from 'KUBE_ENABLE_EGRESS_VIA_KONNECTIVITY_SERVICE ./cluster/kube-up.sh' to 'KUBE_ENABLE_KONNECTIVITY_SERVICE=true ./hack/kube-up.sh' ([#92669](https://github.com/kubernetes/kubernetes/pull/92669), [@Jefftree](https://github.com/Jefftree)) [SIG Cloud Provider]
- Configure AWS LoadBalancer health check protocol via service annotations. ([#94546](https://github.com/kubernetes/kubernetes/pull/94546), [@kishorj](https://github.com/kishorj))
- DefaultPodTopologySpread graduated to Beta. The feature gate is enabled by default. ([#95631](https://github.com/kubernetes/kubernetes/pull/95631), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling and Testing]
- E2e test for PodFsGroupChangePolicy ([#96247](https://github.com/kubernetes/kubernetes/pull/96247), [@saikat-royc](https://github.com/saikat-royc)) [SIG Storage and Testing]
- Ephemeral containers now apply the same API defaults as initContainers and containers ([#94896](https://github.com/kubernetes/kubernetes/pull/94896), [@wawa0210](https://github.com/wawa0210)) [SIG Apps and CLI]
- Graduate the Pod Resources API to G.A
  Introduces the pod_resources_endpoint_requests_total metric which tracks the total number of requests to the pod resources API ([#92165](https://github.com/kubernetes/kubernetes/pull/92165), [@RenaudWasTaken](https://github.com/RenaudWasTaken)) [SIG Instrumentation, Node and Testing]
- In dual-stack bare-metal clusters, you can now pass dual-stack IPs to `kubelet --node-ip`.
  eg: `kubelet --node-ip 10.1.0.5,fd01::0005`. This is not yet supported for non-bare-metal
  clusters.
  
  In dual-stack clusters where nodes have dual-stack addresses, hostNetwork pods
  will now get dual-stack PodIPs. ([#95239](https://github.com/kubernetes/kubernetes/pull/95239), [@danwinship](https://github.com/danwinship)) [SIG Network and Node]
- Introduce api-extensions category which will return: mutating admission configs, validating admission configs, CRDs and APIServices when used in kubectl get, for example. ([#95603](https://github.com/kubernetes/kubernetes/pull/95603), [@soltysh](https://github.com/soltysh)) [SIG API Machinery]
- Introduces a new GCE specific cluster creation variable KUBE_PROXY_DISABLE. When set to true, this will skip over the creation of kube-proxy (whether the daemonset or static pod). This can be used to control the lifecycle of kube-proxy separately from the lifecycle of the nodes. ([#91977](https://github.com/kubernetes/kubernetes/pull/91977), [@varunmar](https://github.com/varunmar)) [SIG Cloud Provider]
- Kube-apiserver now maintains a Lease object to identify itself: 
  - The feature is under feature gate `APIServerIdentity`. 
  - Two flags are added to kube-apiserver: `identity-lease-duration-seconds`, `identity-lease-renew-interval-seconds` ([#95533](https://github.com/kubernetes/kubernetes/pull/95533), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Kube-apiserver: The timeout used when making health check calls to etcd can now be configured with `--etcd-healthcheck-timeout`. The default timeout is 2 seconds, matching the previous behavior. ([#93244](https://github.com/kubernetes/kubernetes/pull/93244), [@Sh4d1](https://github.com/Sh4d1)) [SIG API Machinery]
- Kube-apiserver: added support for compressing rotated audit log files with `--audit-log-compress` ([#94066](https://github.com/kubernetes/kubernetes/pull/94066), [@lojies](https://github.com/lojies)) [SIG API Machinery and Auth]
- Kubeadm now prints warnings instead of throwing errors if the current system time is outside of the NotBefore and NotAfter bounds of a loaded certificate.  ([#94504](https://github.com/kubernetes/kubernetes/pull/94504), [@neolit123](https://github.com/neolit123))
- Kubeadm: Add a preflight check that the control-plane node has at least 1700MB of RAM ([#93275](https://github.com/kubernetes/kubernetes/pull/93275), [@xlgao-zju](https://github.com/xlgao-zju)) [SIG Cluster Lifecycle]
- Kubeadm: add the "--cluster-name" flag to the "kubeadm alpha kubeconfig user" to allow configuring the cluster name in the generated kubeconfig file ([#93992](https://github.com/kubernetes/kubernetes/pull/93992), [@prabhu43](https://github.com/prabhu43)) [SIG Cluster Lifecycle]
- Kubeadm: add the "--kubeconfig" flag to the "kubeadm init phase upload-certs" command to allow users to pass a custom location for a kubeconfig file. ([#94765](https://github.com/kubernetes/kubernetes/pull/94765), [@zhanw15](https://github.com/zhanw15)) [SIG Cluster Lifecycle]
- Kubeadm: make etcd pod request 100m CPU, 100Mi memory and 100Mi ephemeral_storage by default ([#94479](https://github.com/kubernetes/kubernetes/pull/94479), [@knight42](https://github.com/knight42)) [SIG Cluster Lifecycle]
- Kubeadm: make the command "kubeadm alpha kubeconfig user" accept a "--config" flag and remove the following flags:
  - apiserver-advertise-address / apiserver-bind-port: use either localAPIEndpoint from InitConfiguration or controlPlaneEndpoint from ClusterConfiguration.
  - cluster-name: use clusterName from ClusterConfiguration
  - cert-dir: use certificatesDir from ClusterConfiguration ([#94879](https://github.com/kubernetes/kubernetes/pull/94879), [@knight42](https://github.com/knight42)) [SIG Cluster Lifecycle]
- Kubectl create now supports creating ingress objects. ([#94327](https://github.com/kubernetes/kubernetes/pull/94327), [@rikatz](https://github.com/rikatz)) [SIG CLI and Network]
- Kubectl rollout history sts/sts-name --revision=some-revision will start showing the detailed view of  the sts on that specified revision ([#86506](https://github.com/kubernetes/kubernetes/pull/86506), [@dineshba](https://github.com/dineshba)) [SIG CLI]
- Kubectl: Previously users cannot provide arguments to a external diff tool via KUBECTL_EXTERNAL_DIFF env. This release now allow users to specify args to KUBECTL_EXTERNAL_DIFF env. ([#95292](https://github.com/kubernetes/kubernetes/pull/95292), [@dougsland](https://github.com/dougsland)) [SIG CLI]
- Kubemark now supports both real and hollow nodes in a single cluster. ([#93201](https://github.com/kubernetes/kubernetes/pull/93201), [@ellistarn](https://github.com/ellistarn)) [SIG Scalability]
- Kubernetes E2E test image manifest lists now contain Windows images. ([#77398](https://github.com/kubernetes/kubernetes/pull/77398), [@claudiubelu](https://github.com/claudiubelu)) [SIG Testing and Windows]
- Kubernetes is now built using go1.15.2
  - build: Update to k/repo-infra@v0.1.1 (supports go1.15.2)
  - build: Use go-runner:buster-v2.0.1 (built using go1.15.1)
  - bazel: Replace --features with Starlark build settings flag
  - hack/lib/util.sh: some bash cleanups
    
    - switched one spot to use kube::logging
    - make kube::util::find-binary return an error when it doesn't find
      anything so that hack scripts fail fast instead of with "binary not
      found errors".
    - this required deleting some genfeddoc stuff. the binary no longer
      exists in k/k repo since we removed federation/, and I don't see it
      in https://github.com/kubernetes-sigs/kubefed/ either. I'm assuming
      that it's gone for good now.
  
  - bazel: output go_binary rule directly from go_binary_conditional_pure
    
    From: [@mikedanese](https://github.com/mikedanese):
    Instead of aliasing. Aliases are annoying in a number of ways. This is
    specifically bugging me now because they make the action graph harder to
    analyze programmatically. By using aliases here, we would need to handle
    potentially aliased go_binary targets and dereference to the effective
    target.
  
    The comment references an issue with `pure = select(...)` which appears
    to be resolved considering this now builds.
  
  - make kube::util::find-binary not dependent on bazel-out/ structure
  
    Implement an aspect that outputs go_build_mode metadata for go binaries,
    and use that during binary selection. ([#94449](https://github.com/kubernetes/kubernetes/pull/94449), [@justaugustus](https://github.com/justaugustus)) [SIG Architecture, CLI, Cluster Lifecycle, Node, Release and Testing]
- Kubernetes is now built using go1.15.5
  - build: Update to k/repo-infra@v0.1.2 (supports go1.15.5) ([#95776](https://github.com/kubernetes/kubernetes/pull/95776), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- New default scheduling plugins order reduces scheduling and preemption latency when taints and node affinity are used ([#95539](https://github.com/kubernetes/kubernetes/pull/95539), [@soulxu](https://github.com/soulxu)) [SIG Scheduling]
- Only update Azure data disks when attach/detach ([#94265](https://github.com/kubernetes/kubernetes/pull/94265), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Promote SupportNodePidsLimit to GA to provide node-to-pod PID isolation.
  Promote SupportPodPidsLimit to GA to provide the ability to limit PIDs per pod. ([#94140](https://github.com/kubernetes/kubernetes/pull/94140), [@derekwaynecarr](https://github.com/derekwaynecarr))
- SCTP support in API objects (Pod, Service, NetworkPolicy) is now GA.
  Note that this has no effect on whether SCTP is enabled on nodes at the kernel level,
  and note that some cloud platforms and network plugins do not support SCTP traffic. ([#95566](https://github.com/kubernetes/kubernetes/pull/95566), [@danwinship](https://github.com/danwinship)) [SIG Apps and Network]
- Scheduler now ignores Pod update events if the resourceVersion of old and new Pods are identical. ([#96071](https://github.com/kubernetes/kubernetes/pull/96071), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Scheduling Framework: expose Run[Pre]ScorePlugins functions to PreemptionHandle which can be used in PostFilter extension point. ([#93534](https://github.com/kubernetes/kubernetes/pull/93534), [@everpeace](https://github.com/everpeace)) [SIG Scheduling and Testing]
- SelectorSpreadPriority maps to PodTopologySpread plugin when DefaultPodTopologySpread feature is enabled ([#95448](https://github.com/kubernetes/kubernetes/pull/95448), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Send GCE node startup scripts' logs to console and journal. ([#95311](https://github.com/kubernetes/kubernetes/pull/95311), [@karan](https://github.com/karan))
- SetHostnameAsFQDN has been graduated to Beta and therefore it is enabled by default. ([#95267](https://github.com/kubernetes/kubernetes/pull/95267), [@javidiaz](https://github.com/javidiaz)) [SIG Node]
- Support [service.beta.kubernetes.io/azure-pip-ip-tags] annotations to allow customers to specify ip-tags to influence public-ip creation in Azure [Tag1=Value1, Tag2=Value2, etc.] ([#94114](https://github.com/kubernetes/kubernetes/pull/94114), [@MarcPow](https://github.com/MarcPow)) [SIG Cloud Provider]
- Support custom tags for cloud provider managed resources ([#96450](https://github.com/kubernetes/kubernetes/pull/96450), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Support customize load balancer health probe protocol and request path ([#96338](https://github.com/kubernetes/kubernetes/pull/96338), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Support for Windows container images (OS Versions: 1809, 1903, 1909, 2004) was added to the pause:3.4 image. ([#91452](https://github.com/kubernetes/kubernetes/pull/91452), [@claudiubelu](https://github.com/claudiubelu)) [SIG Node, Release and Windows]
- Support multiple standard load balancers in one cluster ([#96111](https://github.com/kubernetes/kubernetes/pull/96111), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- The beta `RootCAConfigMap` feature gate is enabled by default and causes kube-controller-manager to publish a "kube-root-ca.crt" ConfigMap to every namespace. This ConfigMap contains a CA bundle used for verifying connections to the kube-apiserver. ([#96197](https://github.com/kubernetes/kubernetes/pull/96197), [@zshihang](https://github.com/zshihang)) [SIG API Machinery, Apps, Auth and Testing]
- The kubelet_runtime_operations_duration_seconds metric buckets were set to 0.005 0.0125 0.03125 0.078125 0.1953125 0.48828125 1.220703125 3.0517578125 7.62939453125 19.073486328125 47.6837158203125 119.20928955078125 298.0232238769531 and 745.0580596923828 seconds ([#96054](https://github.com/kubernetes/kubernetes/pull/96054), [@alvaroaleman](https://github.com/alvaroaleman)) [SIG Instrumentation and Node]
- There is a new pv_collector_total_pv_count metric that counts persistent volumes by the volume plugin name and volume mode. ([#95719](https://github.com/kubernetes/kubernetes/pull/95719), [@tsmetana](https://github.com/tsmetana)) [SIG Apps, Instrumentation, Storage and Testing]
- Volume snapshot e2e test to validate PVC and VolumeSnapshotContent finalizer ([#95863](https://github.com/kubernetes/kubernetes/pull/95863), [@RaunakShah](https://github.com/RaunakShah)) [SIG Cloud Provider, Storage and Testing]
- Warns user when executing kubectl apply/diff to a resource currently being deleted. ([#95544](https://github.com/kubernetes/kubernetes/pull/95544), [@SaiHarshaK](https://github.com/SaiHarshaK)) [SIG CLI]
- `kubectl alpha debug` has graduated to beta and is now `kubectl debug`. ([#96138](https://github.com/kubernetes/kubernetes/pull/96138), [@verb](https://github.com/verb)) [SIG CLI and Testing]
- `kubectl debug` gains support for changing container images when copying a pod for debugging, similar to how `kubectl set image` works. See `kubectl help debug` for more information. ([#96058](https://github.com/kubernetes/kubernetes/pull/96058), [@verb](https://github.com/verb)) [SIG CLI]



#### Documentation

- Fake dynamic client: document that List does not preserve TypeMeta in UnstructuredList ([#95117](https://github.com/kubernetes/kubernetes/pull/95117), [@andrewsykim](https://github.com/andrewsykim)) [SIG API Machinery]
- Kubelet: remove alpha warnings for CNI flags. ([#94508](https://github.com/kubernetes/kubernetes/pull/94508), [@andrewsykim](https://github.com/andrewsykim)) [SIG Network and Node]
- Updates docs and guidance on cloud provider InstancesV2 and Zones interface for external cloud providers:
  - removes experimental warning for InstancesV2
  - document that implementation of InstancesV2 will disable calls to Zones
  - deprecate Zones in favor of InstancesV2 ([#96397](https://github.com/kubernetes/kubernetes/pull/96397), [@andrewsykim](https://github.com/andrewsykim)) [SIG Cloud Provider]



#### Failing Test

- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Fix handing special characters in the volume path on Windows ([#99008](https://github.com/kubernetes/kubernetes/pull/99008), [@yujuhong](https://github.com/yujuhong)) [SIG Storage]
- Kube-proxy: fix a bug on UDP NodePort Services where stale conntrack entries may blackhole the traffic directed to the NodePort. ([#98305](https://github.com/kubernetes/kubernetes/pull/98305), [@aojea](https://github.com/aojea)) [SIG Network]
- Kubelet: the HostPort implementation in dockershim was not taking into consideration the HostIP field, causing that the same HostPort can not be used with different IP addresses.
  This bug causes the conformance test "HostPort validates that there is no conflict between pods with same hostPort but different hostIP and protocol" to fail. ([#98838](https://github.com/kubernetes/kubernetes/pull/98838), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Resolves an issue running Ingress conformance tests on clusters which use finalizers on Ingress objects to manage releasing load balancer resources ([#96742](https://github.com/kubernetes/kubernetes/pull/96742), [@spencerhance](https://github.com/spencerhance)) [SIG Network and Testing]
- The Conformance test "validates that there is no conflict between pods with same hostPort but different hostIP and protocol" now validates the connectivity to each hostPort, in addition to the functionality. ([#96627](https://github.com/kubernetes/kubernetes/pull/96627), [@aojea](https://github.com/aojea)) [SIG Scheduling and Testing]



#### Bug or Regression

- Added jitter factor to lease controller that better smears load on kube-apiserver over time. ([#101652](https://github.com/kubernetes/kubernetes/pull/101652), [@marseel](https://github.com/marseel)) [SIG API Machinery and Scalability]
- Avoid caching the Azure VMSS instances whose network profile is nil ([#100948](https://github.com/kubernetes/kubernetes/pull/100948), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Azure: avoid setting cached Sku when updating VMSS and VMSS instances ([#102005](https://github.com/kubernetes/kubernetes/pull/102005), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix a bug on the endpoint slices mirroring controller where endpoint NotReadyAddresses were mirrored as Ready to the corresponding EndpointSlice ([#102683](https://github.com/kubernetes/kubernetes/pull/102683), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Fix a bug that a preemptor pod may exist as a phantom in the scheduler. ([#102498](https://github.com/kubernetes/kubernetes/pull/102498), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix errors when accessing Windows container stats for Dockershim ([#98510](https://github.com/kubernetes/kubernetes/pull/98510), [@jsturtevant](https://github.com/jsturtevant)) [SIG Node and Windows]
- Fix removing pods from podTopologyHints mapping ([#101896](https://github.com/kubernetes/kubernetes/pull/101896), [@aheng-ch](https://github.com/aheng-ch)) [SIG Node]
- Fix: avoid nil-pointer panic when checking the frontend IP configuration ([#101739](https://github.com/kubernetes/kubernetes/pull/101739), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: delete non existing disk issue ([#102083](https://github.com/kubernetes/kubernetes/pull/102083), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed false-positive uncertain volume attachments, which led to unexpected detachment of CSI migrated volumes ([#101737](https://github.com/kubernetes/kubernetes/pull/101737), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps and Storage]
- Fixed garbage collection of dangling VolumeAttachments for PersistentVolumes migrated to CSI on startup of kube-controller-manager. ([#102176](https://github.com/kubernetes/kubernetes/pull/102176), [@timebertt](https://github.com/timebertt)) [SIG Apps and Storage]
- Improve speed of vSphere PV provisioning and reduce number of API calls ([#102350](https://github.com/kubernetes/kubernetes/pull/102350), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- Kubeadm: remove the "ephemeral_storage" request from the etcd static pod that kubeadm deploys on stacked etcd control plane nodes. This request has caused sporadic failures on some setups due to a problem in the kubelet with cadvisor and the LocalStorageCapacityIsolation feature gate. See this issue for more details: https://github.com/kubernetes/kubernetes/issues/99305 ([#102673](https://github.com/kubernetes/kubernetes/pull/102673), [@jackfrancis](https://github.com/jackfrancis)) [SIG Cluster Lifecycle]
- Register/Deregister Targets in chunks for AWS TargetGroup ([#101592](https://github.com/kubernetes/kubernetes/pull/101592), [@M00nF1sh](https://github.com/M00nF1sh)) [SIG Cloud Provider]
- Respect annotation size limit for server-side apply updates to the client-side apply annotation. Also, fix opt-out of this behavior by setting the client-side apply annotation to the empty string. ([#102105](https://github.com/kubernetes/kubernetes/pull/102105), [@julianvmodesto](https://github.com/julianvmodesto)) [SIG API Machinery]
- Reverted the previous fix for portforward cleanup because it introduced a kubelet regression which can lead into segmentation faults. ([#102586](https://github.com/kubernetes/kubernetes/pull/102586), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Azurefile: Normalize share name to not include capital letters ([#100731](https://github.com/kubernetes/kubernetes/pull/100731), [@kassarl](https://github.com/kassarl)) [SIG Cloud Provider and Storage]
- EndpointSlice IP validation now matches Endpoints IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSlice controllers are less likely to create duplicate EndpointSlices. ([#101763](https://github.com/kubernetes/kubernetes/pull/101763), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Ensure service deleted when the Azure resource group has been deleted ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix panic in JSON logging format caused by missing Duration encoder ([#101158](https://github.com/kubernetes/kubernetes/pull/101158), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fix smb mount PermissionDenied issue on Windows ([#99550](https://github.com/kubernetes/kubernetes/pull/99550), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider, Storage and Windows]
- Fix: azure file inline volume namespace issue in csi migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx)) [SIG Apps, Cloud Provider, Node and Storage]
- Fix: not tagging static public IP ([#101752](https://github.com/kubernetes/kubernetes/pull/101752), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix: set "host is down" as corrupted mount ([#101398](https://github.com/kubernetes/kubernetes/pull/101398), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug where startupProbe stopped working after a container's first restart ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed port-forward memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Kubectl create service now respects namespace flag ([#101005](https://github.com/kubernetes/kubernetes/pull/101005), [@zxh326](https://github.com/zxh326)) [SIG CLI]
- Kubelet: improve the performance when waiting for a synchronization of the node list with the kube-apiserver ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123)) [SIG Node]
- No support endpointslice in linux userpace mode ([#101503](https://github.com/kubernetes/kubernetes/pull/101503), [@JornShen](https://github.com/JornShen)) [SIG Network]
- Renames the timeout field for the DelegatingAuthenticationOptions to TokenRequestTimeout and set the timeout only for the token review client. Previously the timeout was also applied to watches making them reconnecting every 10 seconds. ([#101103](https://github.com/kubernetes/kubernetes/pull/101103), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Auth and Cloud Provider]
- Respect ExecProbeTimeout=false for dockershim ([#101126](https://github.com/kubernetes/kubernetes/pull/101126), [@jackfrancis](https://github.com/jackfrancis)) [SIG Node and Testing]
- Fix priority expander falling back to a random choice even though there is a higher priority option to choose
- Clone `kubernetes/kubernetes` in `update-vendor.sh` shallowly, instead of fetching all revisions
- Speed up binpacking by reducing the number of PreFilter calls (call once per pod instead of `pods` * `nodes times`)
- Speed up finding unneeded nodes by 5x+ in very large clusters by reducing the number of PreFilter calls
- Expose `--max-nodes-total` as a metric
- Errors in `IncreaseSize` changed from type `apiError` to `cloudProviderError`
- Make `build-in-docker` and `test-in-docker` work on Linux systems with SELinux enabled
- Fix an error where existing nodes were not considered as destinations while finding place for pods in scale-down simulations
- Remove redundant log lines and reduce severity around parsing kubeEnv
- Don't treat nodes created by virtual kubelet as nodes from non-autoscaled node groups
- Remove redundant logging around calculating node utilization
- Add configurable `--network` and `--rm` flags for docker in `Makefile`
- Subtract DaemonSet pods' requests from node allocatable in the denominator while computing node utilization
- Include taints by condition when determining if a node is unready/still starting
- Fix `update-vendor.sh` to work on OSX and zsh
- Add best-effort eviction for DaemonSet pods while scaling down non-empty nodes
- Add build support for ARM64
- Regenerate list of EC2 instances
- Fix pricing endpoint in AWS China Region
- Avoid systemd-logind loading configuration warning ([#97950](https://github.com/kubernetes/kubernetes/pull/97950), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Count pod overhead against an entity's ResourceQuota ([#99600](https://github.com/kubernetes/kubernetes/pull/99600), [@gjkim42](https://github.com/gjkim42)) [SIG API Machinery and Node]
- EndpointSlice controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100113](https://github.com/kubernetes/kubernetes/pull/100113), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSliceMirroring controller is now less likely to emit FailedToUpdateEndpointSlices events. ([#100143](https://github.com/kubernetes/kubernetes/pull/100143), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- Ensure only one LoadBalancer rule is created when HA mode is enabled ([#99825](https://github.com/kubernetes/kubernetes/pull/99825), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix kubelet from panic after getting the wrong signal ([#98200](https://github.com/kubernetes/kubernetes/pull/98200), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fix repeatedly acquire the inhibit lock ([#98088](https://github.com/kubernetes/kubernetes/pull/98088), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed bug that caused cAdvisor to incorrectly detect single-socket multi-NUMA topology. ([#99207](https://github.com/kubernetes/kubernetes/pull/99207), [@iwankgb](https://github.com/iwankgb)) [SIG Node]
- Fixing a bug where a failed node may not have the NoExecute taint set correctly ([#98168](https://github.com/kubernetes/kubernetes/pull/98168), [@CKchen0726](https://github.com/CKchen0726)) [SIG Apps and Node]
- Kubelet now cleans up orphaned volume directories automatically ([#95301](https://github.com/kubernetes/kubernetes/pull/95301), [@lorenz](https://github.com/lorenz)) [SIG Node and Storage]
- Resolves spurious `Failed to list *v1.Secret` or `Failed to list *v1.ConfigMap` messages in kubelet logs. ([#99538](https://github.com/kubernetes/kubernetes/pull/99538), [@liggitt](https://github.com/liggitt)) [SIG Auth and Node]
- Sync node status during kubelet node shutdown.
  Adds an pod admission handler that rejects new pods when the node is in progress of shutting down. ([#98005](https://github.com/kubernetes/kubernetes/pull/98005), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- We will no longer automatically delete all data when a failure is detected during creation of the volume data file on a CSI volume. Now we will only remove the data file and volume path. ([#96021](https://github.com/kubernetes/kubernetes/pull/96021), [@huffmanca](https://github.com/huffmanca)) [SIG Storage]
- Aggregate errors when putting vmss ([#98350](https://github.com/kubernetes/kubernetes/pull/98350), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Avoid marking node as Ready until node has synced with API servers at least once ([#97995](https://github.com/kubernetes/kubernetes/pull/97995), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Cleanup subnet in frontend IP configs to prevent huge subnet request bodies in some scenarios. ([#98132](https://github.com/kubernetes/kubernetes/pull/98132), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix CSI-migrated inline EBS volumes failing to mount if their volumeID is prefixed by aws:// ([#96821](https://github.com/kubernetes/kubernetes/pull/96821), [@wongma7](https://github.com/wongma7)) [SIG Storage]
- Fix azure file migration issue ([#97877](https://github.com/kubernetes/kubernetes/pull/97877), [@andyzhangx](https://github.com/andyzhangx)) [SIG Auth, Cloud Provider and Storage]
- Fix kubectl-convert import known versions ([#97754](https://github.com/kubernetes/kubernetes/pull/97754), [@wzshiming](https://github.com/wzshiming)) [SIG CLI and Testing]
- Fix the description of command line flags that can override --config ([#98786](https://github.com/kubernetes/kubernetes/pull/98786), [@changshuchao](https://github.com/changshuchao)) [SIG Scheduling]
- Fix the panic when kubelet registers if a node object already exists with no Status.Capacity or Status.Allocatable ([#97803](https://github.com/kubernetes/kubernetes/pull/97803), [@TeddyAndrieux](https://github.com/TeddyAndrieux)) [SIG Node]
- Fix the regression with the slow pods termination. Before this fix pods may take an additional time to terminate - up to one minute. Reversing the change that ensured that CNI resources cleaned up when the pod is removed on API server. ([#97980](https://github.com/kubernetes/kubernetes/pull/97980), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev)) [SIG Node]
- Fix to recover CSI volumes from certain dangling attachments ([#96617](https://github.com/kubernetes/kubernetes/pull/96617), [@yuga711](https://github.com/yuga711)) [SIG Apps and Storage]
- Fixed a bug that the kubelet cannot start on BtrfS. ([#98014](https://github.com/kubernetes/kubernetes/pull/98014), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Fixed an issue with garbage collection failing to clean up namespaced children of an object also referenced incorrectly by cluster-scoped children ([#98068](https://github.com/kubernetes/kubernetes/pull/98068), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Apps]
- Fixed provisioning of Cinder volumes migrated to CSI when StorageClass with AllowedTopologies was used. ([#98311](https://github.com/kubernetes/kubernetes/pull/98311), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixes a panic in the disruption budget controller for PDB objects with invalid selectors ([#98775](https://github.com/kubernetes/kubernetes/pull/98775), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Apps]
- Fixes connection errors when using `--volume-host-cidr-denylist` or `--volume-host-allow-local-loopback` ([#98436](https://github.com/kubernetes/kubernetes/pull/98436), [@liggitt](https://github.com/liggitt)) [SIG Network and Storage]
- Kubeadm: get k8s CI version markers from k8s infra bucket ([#98836](https://github.com/kubernetes/kubernetes/pull/98836), [@hasheddan](https://github.com/hasheddan)) [SIG Cluster Lifecycle and Release]
- Kubelet should ignore cgroup driver check on Windows node. ([#98383](https://github.com/kubernetes/kubernetes/pull/98383), [@pacoxu](https://github.com/pacoxu)) [SIG Node]
- Make podTopologyHints protected by lock ([#95111](https://github.com/kubernetes/kubernetes/pull/95111), [@choury](https://github.com/choury)) [SIG Node]
- Static pods will be deleted gracefully. ([#98103](https://github.com/kubernetes/kubernetes/pull/98103), [@gjkim42](https://github.com/gjkim42)) [SIG Node]
- Truncates a message if it hits the NoteLengthLimit when the scheduler records an event for the pod that indicates the pod has failed to schedule. ([#98715](https://github.com/kubernetes/kubernetes/pull/98715), [@carlory](https://github.com/carlory)) [SIG Scheduling]
- Warning about using a deprecated volume plugin is logged only once. ([#96751](https://github.com/kubernetes/kubernetes/pull/96751), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fix Azure file share not deleted issue when the namespace is deleted ([#97417](https://github.com/kubernetes/kubernetes/pull/97417), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix counting error in service/nodeport/loadbalancer quota check ([#97826](https://github.com/kubernetes/kubernetes/pull/97826), [@pacoxu](https://github.com/pacoxu)) [SIG API Machinery and Network]
- Fix missing cadvisor machine metrics. ([#97006](https://github.com/kubernetes/kubernetes/pull/97006), [@lingsamuel](https://github.com/lingsamuel)) [SIG Node]
- Fix: azure file latency issue for metadata-heavy workloads ([#97082](https://github.com/kubernetes/kubernetes/pull/97082), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed bug in CPUManager with race on container map access ([#97427](https://github.com/kubernetes/kubernetes/pull/97427), [@klueska](https://github.com/klueska)) [SIG Node]
- GCE Internal LoadBalancer sync loop will now release the ILB IP address upon sync failure. An error in ILB forwarding rule creation will no longer leak IP addresses. ([#97740](https://github.com/kubernetes/kubernetes/pull/97740), [@prameshj](https://github.com/prameshj)) [SIG Cloud Provider and Network]
- Kubeadm: avoid detection of the container runtime for commands that do not need it ([#97847](https://github.com/kubernetes/kubernetes/pull/97847), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Performance regression [#97685](https://github.com/kubernetes/kubernetes/issues/97685) has been fixed. ([#97860](https://github.com/kubernetes/kubernetes/pull/97860), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery]
- Use network.Interface.VirtualMachine.ID to get the binded VM
  Skip standalone VM when reconciling LoadBalancer ([#97639](https://github.com/kubernetes/kubernetes/pull/97639), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- AcceleratorStats will be available in the Summary API of kubelet when cri_stats_provider is used. ([#97018](https://github.com/kubernetes/kubernetes/pull/97018), [@ruiwen-zhao](https://github.com/ruiwen-zhao)) [SIG Node]
- Fixed FibreChannel volume plugin corrupting filesystems on detach of multipath volumes. ([#97013](https://github.com/kubernetes/kubernetes/pull/97013), [@jsafrane](https://github.com/jsafrane)) [SIG Storage]
- Fixed a bug in kubelet that will saturate CPU utilization after containerd got restarted. ([#97175](https://github.com/kubernetes/kubernetes/pull/97175), [@hanlins](https://github.com/hanlins)) [SIG Node]
- Kubeadm now installs version 3.4.13 of etcd when creating a cluster with v1.19 ([#97284](https://github.com/kubernetes/kubernetes/pull/97284), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubeadm: Fixes a kubeadm upgrade bug that could cause a custom CoreDNS configuration to be replaced with the default. ([#97016](https://github.com/kubernetes/kubernetes/pull/97016), [@rajansandeep](https://github.com/rajansandeep)) [SIG Cluster Lifecycle]
- Add kubectl wait  --ignore-not-found flag ([#90969](https://github.com/kubernetes/kubernetes/pull/90969), [@zhouya0](https://github.com/zhouya0)) [SIG CLI]
- Added support to kube-proxy for externalTrafficPolicy=Local setting via Direct Server Return (DSR) load balancers on Windows. ([#93166](https://github.com/kubernetes/kubernetes/pull/93166), [@elweb9858](https://github.com/elweb9858)) [SIG Network]
- Alter wording to describe pods using a pvc ([#95635](https://github.com/kubernetes/kubernetes/pull/95635), [@RaunakShah](https://github.com/RaunakShah)) [SIG CLI]
- An issues preventing volume expand controller to annotate the PVC with `volume.kubernetes.io/storage-resizer` when the PVC StorageClass is already updated to the out-of-tree provisioner is now fixed. ([#94489](https://github.com/kubernetes/kubernetes/pull/94489), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG API Machinery, Apps and Storage]
- Azure ARM client: don't segfault on empty response and http error ([#94078](https://github.com/kubernetes/kubernetes/pull/94078), [@bpineau](https://github.com/bpineau)) [SIG Cloud Provider]
- Azure armclient backoff step defaults to 1 (no retry). ([#94180](https://github.com/kubernetes/kubernetes/pull/94180), [@feiskyer](https://github.com/feiskyer))
- Azure: fix a bug that kube-controller-manager would panic if wrong Azure VMSS name is configured ([#94306](https://github.com/kubernetes/kubernetes/pull/94306), [@knight42](https://github.com/knight42)) [SIG Cloud Provider]
- Both apiserver_request_duration_seconds metrics and RequestReceivedTimestamp fields of an audit event now take into account the time a request spends in the apiserver request filters. ([#94903](https://github.com/kubernetes/kubernetes/pull/94903), [@tkashem](https://github.com/tkashem))
- Build/lib/release: Explicitly use '--platform' in building server images
  
  When we switched to go-runner for building the apiserver,
  controller-manager, and scheduler server components, we no longer
  reference the individual architectures in the image names, specifically
  in the 'FROM' directive of the server image Dockerfiles.
  
  As a result, server images for non-amd64 images copy in the go-runner
  amd64 binary instead of the go-runner that matches that architecture.
  
  This commit explicitly sets the '--platform=linux/${arch}' to ensure
  we're pulling the correct go-runner arch from the manifest list.
  
  Before:
  `FROM ${base_image}`
  
  After:
  `FROM --platform=linux/${arch} ${base_image}` ([#94552](https://github.com/kubernetes/kubernetes/pull/94552), [@justaugustus](https://github.com/justaugustus)) [SIG Release]
- Bump node-problem-detector version to v0.8.5 to fix OOM detection in with Linux kernels 5.1+ ([#96716](https://github.com/kubernetes/kubernetes/pull/96716), [@tosi3k](https://github.com/tosi3k)) [SIG Cloud Provider, Scalability and Testing]
- CSIDriver object can be deployed during volume attachment. ([#93710](https://github.com/kubernetes/kubernetes/pull/93710), [@Jiawei0227](https://github.com/Jiawei0227)) [SIG Apps, Node, Storage and Testing]
- Ceph RBD volume expansion now works even when ceph.conf was not provided. ([#92027](https://github.com/kubernetes/kubernetes/pull/92027), [@juliantaylor](https://github.com/juliantaylor))
- Change plugin name in fsgroupapplymetrics of csi and flexvolume to distinguish different driver ([#95892](https://github.com/kubernetes/kubernetes/pull/95892), [@JornShen](https://github.com/JornShen)) [SIG Instrumentation, Storage and Testing]
- Change the calculation of pod UIDs so that static pods get a unique value - will cause all containers to be killed and recreated after in-place upgrade. ([#87461](https://github.com/kubernetes/kubernetes/pull/87461), [@bboreham](https://github.com/bboreham)) [SIG Node]
- Change the mount way from systemd to normal mount except ceph and glusterfs intree-volume. ([#94916](https://github.com/kubernetes/kubernetes/pull/94916), [@smileusd](https://github.com/smileusd)) [SIG Apps, Cloud Provider, Network, Node, Storage and Testing]
- Changes to timeout parameter handling in 1.20.0-beta.2 have been reverted to avoid breaking backwards compatibility with existing clients. ([#96727](https://github.com/kubernetes/kubernetes/pull/96727), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Clear UDP conntrack entry on endpoint changes when using nodeport ([#71573](https://github.com/kubernetes/kubernetes/pull/71573), [@JacobTanenbaum](https://github.com/JacobTanenbaum)) [SIG Network]
- Cloud node controller: handle empty providerID from getProviderID ([#95342](https://github.com/kubernetes/kubernetes/pull/95342), [@nicolehanjing](https://github.com/nicolehanjing)) [SIG Cloud Provider]
- Disable watchcache for events ([#96052](https://github.com/kubernetes/kubernetes/pull/96052), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- Disabled `LocalStorageCapacityIsolation` feature gate is honored during scheduling. ([#96092](https://github.com/kubernetes/kubernetes/pull/96092), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Do not fail sorting empty elements. ([#94666](https://github.com/kubernetes/kubernetes/pull/94666), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Dual-stack: make nodeipam compatible with existing single-stack clusters when dual-stack feature gate become enabled by default ([#90439](https://github.com/kubernetes/kubernetes/pull/90439), [@SataQiu](https://github.com/SataQiu)) [SIG API Machinery]
- Duplicate owner reference entries in create/update/patch requests now get deduplicated by the API server. The client sending the request now receives a warning header in the API response. Clients should stop sending requests with duplicate owner references. The API server may reject such requests as early as 1.24. ([#96185](https://github.com/kubernetes/kubernetes/pull/96185), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery and Testing]
- Endpoint slice controller now mirrors parent's service label to its corresponding endpoint slices. ([#94443](https://github.com/kubernetes/kubernetes/pull/94443), [@aojea](https://github.com/aojea))
- Ensure getPrimaryInterfaceID not panic when network interfaces for Azure VMSS are null ([#94355](https://github.com/kubernetes/kubernetes/pull/94355), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Exposes and sets a default timeout for the SubjectAccessReview client for DelegatingAuthorizationOptions ([#95725](https://github.com/kubernetes/kubernetes/pull/95725), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Exposes and sets a default timeout for the TokenReview client for DelegatingAuthenticationOptions ([#96217](https://github.com/kubernetes/kubernetes/pull/96217), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery and Cloud Provider]
- Fix CVE-2020-8555 for Quobyte client connections. ([#95206](https://github.com/kubernetes/kubernetes/pull/95206), [@misterikkit](https://github.com/misterikkit)) [SIG Storage]
- Fix IP fragmentation of UDP and TCP packets not supported issues on LoadBalancer rules ([#96464](https://github.com/kubernetes/kubernetes/pull/96464), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix a bug that DefaultPreemption plugin is disabled when using (legacy) scheduler policy. ([#96439](https://github.com/kubernetes/kubernetes/pull/96439), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling and Testing]
- Fix a bug where loadbalancer deletion gets stuck because of missing resource group. ([#93962](https://github.com/kubernetes/kubernetes/pull/93962), [@phiphi282](https://github.com/phiphi282))
- Fix a concurrent map writes error in kubelet ([#93773](https://github.com/kubernetes/kubernetes/pull/93773), [@knight42](https://github.com/knight42)) [SIG Node]
- Fix a panic in `kubectl debug` when a pod has multiple init or ephemeral containers. ([#94580](https://github.com/kubernetes/kubernetes/pull/94580), [@kiyoshim55](https://github.com/kiyoshim55))
- Fix a regression where kubeadm bails out with a fatal error when an optional version command line argument is supplied to the "kubeadm upgrade plan" command ([#94421](https://github.com/kubernetes/kubernetes/pull/94421), [@rosti](https://github.com/rosti)) [SIG Cluster Lifecycle]
- Fix azure disk attach failure for disk size bigger than 4TB ([#95463](https://github.com/kubernetes/kubernetes/pull/95463), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix azure disk data loss issue on Windows when unmount disk ([#95456](https://github.com/kubernetes/kubernetes/pull/95456), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fix azure file migration panic ([#94853](https://github.com/kubernetes/kubernetes/pull/94853), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix bug in JSON path parser where an error occurs when a range is empty ([#95933](https://github.com/kubernetes/kubernetes/pull/95933), [@brianpursley](https://github.com/brianpursley)) [SIG API Machinery]
- Fix client-go prometheus metrics to correctly present the API path accessed in some environments. ([#74363](https://github.com/kubernetes/kubernetes/pull/74363), [@aanm](https://github.com/aanm)) [SIG API Machinery]
- Fix detach azure disk issue when vm not exist ([#95177](https://github.com/kubernetes/kubernetes/pull/95177), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix etcd_object_counts metric reported by kube-apiserver ([#94773](https://github.com/kubernetes/kubernetes/pull/94773), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fix incorrectly reported verbs for kube-apiserver metrics for CRD objects ([#93523](https://github.com/kubernetes/kubernetes/pull/93523), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Instrumentation]
- Fix k8s.io/apimachinery/pkg/api/meta.SetStatusCondition to update ObservedGeneration ([#95961](https://github.com/kubernetes/kubernetes/pull/95961), [@KnicKnic](https://github.com/KnicKnic)) [SIG API Machinery]
- Fix kubectl SchemaError on CRDs with schema using x-kubernetes-preserve-unknown-fields on array types. ([#94888](https://github.com/kubernetes/kubernetes/pull/94888), [@sttts](https://github.com/sttts)) [SIG API Machinery]
- Fix memory leak in kube-apiserver when underlying time goes forth and back. ([#96266](https://github.com/kubernetes/kubernetes/pull/96266), [@chenyw1990](https://github.com/chenyw1990)) [SIG API Machinery]
- Fix missing csi annotations on node during parallel csinode update. ([#94389](https://github.com/kubernetes/kubernetes/pull/94389), [@pacoxu](https://github.com/pacoxu)) [SIG Storage]
- Fix network_programming_latency metric reporting for Endpoints/EndpointSlice deletions, where we don't have correct timestamp ([#95363](https://github.com/kubernetes/kubernetes/pull/95363), [@wojtek-t](https://github.com/wojtek-t)) [SIG Network and Scalability]
- Fix paging issues when Azure API returns empty values with non-empty nextLink ([#96211](https://github.com/kubernetes/kubernetes/pull/96211), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix pull image error from multiple ACRs using azure managed identity ([#96355](https://github.com/kubernetes/kubernetes/pull/96355), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix race condition on timeCache locks. ([#94751](https://github.com/kubernetes/kubernetes/pull/94751), [@auxten](https://github.com/auxten))
- Fix regression on `kubectl port-forward` when TCP and UCP services were configured on the same port. ([#94728](https://github.com/kubernetes/kubernetes/pull/94728), [@amorenoz](https://github.com/amorenoz))
- Fix scheduler cache snapshot when a Node is deleted before its Pods ([#95130](https://github.com/kubernetes/kubernetes/pull/95130), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Fix the `cloudprovider_azure_api_request_duration_seconds` metric buckets to correctly capture the latency metrics. Previously, the majority of the calls would fall in the "+Inf" bucket. ([#94873](https://github.com/kubernetes/kubernetes/pull/94873), [@marwanad](https://github.com/marwanad)) [SIG Cloud Provider and Instrumentation]
- Fix vSphere volumes that could be erroneously attached to wrong node ([#96224](https://github.com/kubernetes/kubernetes/pull/96224), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- Fix verb & scope reporting for kube-apiserver metrics (LIST reported instead of GET) ([#95562](https://github.com/kubernetes/kubernetes/pull/95562), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Testing]
- Fix vSphere detach failure for static PVs ([#95447](https://github.com/kubernetes/kubernetes/pull/95447), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider and Storage]
- Fix: azure disk resize error if source does not exist ([#93011](https://github.com/kubernetes/kubernetes/pull/93011), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: detach azure disk broken on Azure Stack ([#94885](https://github.com/kubernetes/kubernetes/pull/94885), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: resize Azure disk issue when it's in attached state ([#96705](https://github.com/kubernetes/kubernetes/pull/96705), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: smb valid path error ([#95583](https://github.com/kubernetes/kubernetes/pull/95583), [@andyzhangx](https://github.com/andyzhangx)) [SIG Storage]
- Fix: use sensitiveOptions on Windows mount ([#94126](https://github.com/kubernetes/kubernetes/pull/94126), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug causing incorrect formatting of `kubectl describe ingress`. ([#94985](https://github.com/kubernetes/kubernetes/pull/94985), [@howardjohn](https://github.com/howardjohn)) [SIG CLI and Network]
- Fixed a bug in client-go where new clients with customized `Dial`, `Proxy`, `GetCert` config may get stale HTTP transports. ([#95427](https://github.com/kubernetes/kubernetes/pull/95427), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixed a bug that prevents kubectl to validate CRDs with schema using x-kubernetes-preserve-unknown-fields on object fields. ([#96369](https://github.com/kubernetes/kubernetes/pull/96369), [@gautierdelorme](https://github.com/gautierdelorme)) [SIG API Machinery and Testing]
- Fixed a bug that prevents the use of ephemeral containers in the presence of a validating admission webhook. ([#94685](https://github.com/kubernetes/kubernetes/pull/94685), [@verb](https://github.com/verb)) [SIG Node and Testing]
- Fixed a bug where aggregator_unavailable_apiservice metrics were reported for deleted apiservices. ([#96421](https://github.com/kubernetes/kubernetes/pull/96421), [@dgrisonnet](https://github.com/dgrisonnet)) [SIG API Machinery and Instrumentation]
- Fixed a bug where improper storage and comparison of endpoints led to excessive API traffic from the endpoints controller ([#94112](https://github.com/kubernetes/kubernetes/pull/94112), [@damemi](https://github.com/damemi)) [SIG Apps, Network and Testing]
- Fixed a regression which prevented pods with `docker/default` seccomp annotations from being created in 1.19 if a PodSecurityPolicy was in place which did not allow `runtime/default` seccomp profiles. ([#95985](https://github.com/kubernetes/kubernetes/pull/95985), [@saschagrunert](https://github.com/saschagrunert)) [SIG Auth]
- Fixed bug in reflector that couldn't recover from "Too large resource version" errors with API servers 1.17.0-1.18.5 ([#94316](https://github.com/kubernetes/kubernetes/pull/94316), [@janeczku](https://github.com/janeczku)) [SIG API Machinery]
- Fixed bug where kubectl top pod output is not sorted when --sort-by and --containers flags are used together ([#93692](https://github.com/kubernetes/kubernetes/pull/93692), [@brianpursley](https://github.com/brianpursley)) [SIG CLI]
- Fixed kubelet creating extra sandbox for pods with RestartPolicyOnFailure after all containers succeeded ([#92614](https://github.com/kubernetes/kubernetes/pull/92614), [@tnqn](https://github.com/tnqn)) [SIG Node and Testing]
- Fixes an issue proxying to ipv6 pods without specifying a port ([#94834](https://github.com/kubernetes/kubernetes/pull/94834), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Network]
- Fixes code generation for non-namespaced create subresources fake client test. ([#96586](https://github.com/kubernetes/kubernetes/pull/96586), [@Doude](https://github.com/Doude)) [SIG API Machinery]
- Fixes high CPU usage in kubectl drain ([#95260](https://github.com/kubernetes/kubernetes/pull/95260), [@amandahla](https://github.com/amandahla)) [SIG CLI]
- For vSphere Cloud Provider, If VM of worker node is deleted, the node will also be deleted by node controller ([#92608](https://github.com/kubernetes/kubernetes/pull/92608), [@lubronzhan](https://github.com/lubronzhan)) [SIG Cloud Provider]
- Gracefully delete nodes when their parent scale set went missing ([#95289](https://github.com/kubernetes/kubernetes/pull/95289), [@bpineau](https://github.com/bpineau)) [SIG Cloud Provider]
- HTTP/2 connection health check is enabled by default in all Kubernetes clients. The feature should work out-of-the-box. If needed, users can tune the feature via the HTTP2_READ_IDLE_TIMEOUT_SECONDS and HTTP2_PING_TIMEOUT_SECONDS environment variables. The feature is disabled if HTTP2_READ_IDLE_TIMEOUT_SECONDS is set to 0. ([#95981](https://github.com/kubernetes/kubernetes/pull/95981), [@caesarxuchao](https://github.com/caesarxuchao)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Node]
- If the user specifies an invalid timeout in the request URL, the request will be aborted with an HTTP 400.
  - If the user specifies a timeout in the request URL that exceeds the maximum request deadline allowed by the apiserver, the request will be aborted with an HTTP 400. ([#96061](https://github.com/kubernetes/kubernetes/pull/96061), [@tkashem](https://github.com/tkashem)) [SIG API Machinery, Network and Testing]
- If we set SelectPolicy MinPolicySelect on scaleUp behavior or scaleDown behavior,Horizontal Pod Autoscaler doesn't automatically scale the number of pods correctly ([#95647](https://github.com/kubernetes/kubernetes/pull/95647), [@JoshuaAndrew](https://github.com/JoshuaAndrew)) [SIG Apps and Autoscaling]
- Ignore apparmor for non-linux operating systems ([#93220](https://github.com/kubernetes/kubernetes/pull/93220), [@wawa0210](https://github.com/wawa0210)) [SIG Node and Windows]
- Ignore root user check when windows pod starts ([#92355](https://github.com/kubernetes/kubernetes/pull/92355), [@wawa0210](https://github.com/wawa0210)) [SIG Node and Windows]
- Improve error messages related to nodePort endpoint changes conntrack entries cleanup. ([#96251](https://github.com/kubernetes/kubernetes/pull/96251), [@ravens](https://github.com/ravens)) [SIG Network]
- In dual-stack clusters, kubelet will now set up both IPv4 and IPv6 iptables rules, which may
  fix some problems, eg with HostPorts. ([#94474](https://github.com/kubernetes/kubernetes/pull/94474), [@danwinship](https://github.com/danwinship)) [SIG Network and Node]
- Increase maximum IOPS of AWS EBS io1 volume to current maximum (64,000). ([#90014](https://github.com/kubernetes/kubernetes/pull/90014), [@jacobmarble](https://github.com/jacobmarble))
- Ipvs: ensure selected scheduler kernel modules are loaded ([#93040](https://github.com/kubernetes/kubernetes/pull/93040), [@cmluciano](https://github.com/cmluciano)) [SIG Network]
- K8s.io/apimachinery: runtime.DefaultUnstructuredConverter.FromUnstructured now handles converting integer fields to typed float values ([#93250](https://github.com/kubernetes/kubernetes/pull/93250), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- Kube-proxy now trims extra spaces found in loadBalancerSourceRanges to match Service validation. ([#94107](https://github.com/kubernetes/kubernetes/pull/94107), [@robscott](https://github.com/robscott)) [SIG Network]
- Kubeadm ensures "kubeadm reset" does not unmount the root "/var/lib/kubelet" directory if it is mounted by the user. ([#93702](https://github.com/kubernetes/kubernetes/pull/93702), [@thtanaka](https://github.com/thtanaka))
- Kubeadm now makes sure the etcd manifest is regenerated upon upgrade even when no etcd version change takes place ([#94395](https://github.com/kubernetes/kubernetes/pull/94395), [@rosti](https://github.com/rosti)) [SIG Cluster Lifecycle]
- Kubeadm now warns (instead of error out) on missing "ca.key" files for root CA, front-proxy CA and etcd CA, during "kubeadm join --control-plane" if the user has provided all certificates, keys and kubeconfig files which require signing with the given CA keys. ([#94988](https://github.com/kubernetes/kubernetes/pull/94988), [@neolit123](https://github.com/neolit123))
- Kubeadm: add missing "--experimental-patches" flag to "kubeadm init phase control-plane" ([#95786](https://github.com/kubernetes/kubernetes/pull/95786), [@Sh4d1](https://github.com/Sh4d1)) [SIG Cluster Lifecycle]
- Kubeadm: avoid a panic when determining if the running version of CoreDNS is supported during upgrades ([#94299](https://github.com/kubernetes/kubernetes/pull/94299), [@zouyee](https://github.com/zouyee)) [SIG Cluster Lifecycle]
- Kubeadm: ensure the etcd data directory is created with 0700 permissions during control-plane init and join ([#94102](https://github.com/kubernetes/kubernetes/pull/94102), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: fix coredns migration should be triggered when there are newdefault configs during kubeadm upgrade ([#96907](https://github.com/kubernetes/kubernetes/pull/96907), [@pacoxu](https://github.com/pacoxu)) [SIG Cluster Lifecycle]
- Kubeadm: fix the bug that kubeadm tries to call 'docker info' even if the CRI socket was for another CR ([#94555](https://github.com/kubernetes/kubernetes/pull/94555), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Kubeadm: for Docker as the container runtime, make the "kubeadm reset" command stop containers before removing them ([#94586](https://github.com/kubernetes/kubernetes/pull/94586), [@BedivereZero](https://github.com/BedivereZero)) [SIG Cluster Lifecycle]
- Kubeadm: make the kubeconfig files for the kube-controller-manager and kube-scheduler use the LocalAPIEndpoint instead of the ControlPlaneEndpoint. This makes kubeadm clusters more reseliant to version skew problems during immutable upgrades: https://kubernetes.io/docs/setup/release/version-skew-policy/#kube-controller-manager-kube-scheduler-and-cloud-controller-manager ([#94398](https://github.com/kubernetes/kubernetes/pull/94398), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: relax the validation of kubeconfig server URLs. Allow the user to define custom kubeconfig server URLs without erroring out during validation of existing kubeconfig files (e.g. when using external CA mode). ([#94816](https://github.com/kubernetes/kubernetes/pull/94816), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubectl: print error if users place flags before plugin name ([#92343](https://github.com/kubernetes/kubernetes/pull/92343), [@knight42](https://github.com/knight42)) [SIG CLI]
- Kubelet: assume that swap is disabled when `/proc/swaps` does not exist ([#93931](https://github.com/kubernetes/kubernetes/pull/93931), [@SataQiu](https://github.com/SataQiu)) [SIG Node]
- New Azure instance types do now have correct max data disk count information. ([#94340](https://github.com/kubernetes/kubernetes/pull/94340), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Cloud Provider and Storage]
- Port mapping now allows the same `containerPort` of different containers to different `hostPort` without naming the mapping explicitly. ([#94494](https://github.com/kubernetes/kubernetes/pull/94494), [@SergeyKanzhelev](https://github.com/SergeyKanzhelev))
- Print go stack traces at -v=4 and not -v=2 ([#94663](https://github.com/kubernetes/kubernetes/pull/94663), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Recreate EndpointSlices on rapid Service creation. ([#94730](https://github.com/kubernetes/kubernetes/pull/94730), [@robscott](https://github.com/robscott))
- Reduce volume name length for vSphere volumes ([#96533](https://github.com/kubernetes/kubernetes/pull/96533), [@gnufied](https://github.com/gnufied)) [SIG Storage]
- Remove ready file and its directory (which is created during volume SetUp) during emptyDir volume TearDown. ([#95770](https://github.com/kubernetes/kubernetes/pull/95770), [@jingxu97](https://github.com/jingxu97)) [SIG Storage]
- Reorganized iptables rules to fix a performance issue ([#95252](https://github.com/kubernetes/kubernetes/pull/95252), [@tssurya](https://github.com/tssurya)) [SIG Network]
- Require feature flag CustomCPUCFSQuotaPeriod if setting a non-default cpuCFSQuotaPeriod in kubelet config. ([#94687](https://github.com/kubernetes/kubernetes/pull/94687), [@karan](https://github.com/karan)) [SIG Node]
- Resolves a regression in 1.19+ with workloads targeting deprecated beta os/arch labels getting stuck in NodeAffinity status on node startup. ([#96810](https://github.com/kubernetes/kubernetes/pull/96810), [@liggitt](https://github.com/liggitt)) [SIG Node]
- Resolves non-deterministic behavior of the garbage collection controller when ownerReferences with incorrect data are encountered. Events with a reason of `OwnerRefInvalidNamespace` are recorded when namespace mismatches between child and owner objects are detected. The [kubectl-check-ownerreferences](https://github.com/kubernetes-sigs/kubectl-check-ownerreferences) tool can be run prior to upgrading to locate existing objects with invalid ownerReferences.
  - A namespaced object with an ownerReference referencing a uid of a namespaced kind which does not exist in the same namespace is now consistently treated as though that owner does not exist, and the child object is deleted.
  - A cluster-scoped object with an ownerReference referencing a uid of a namespaced kind is now consistently treated as though that owner is not resolvable, and the child object is ignored by the garbage collector. ([#92743](https://github.com/kubernetes/kubernetes/pull/92743), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Apps and Testing]
- Skip [k8s.io/kubernetes@v1.19.0/test/e2e/storage/testsuites/base.go:162]: Driver azure-disk doesn't support snapshot type DynamicSnapshot -- skipping
  skip [k8s.io/kubernetes@v1.19.0/test/e2e/storage/testsuites/base.go:185]: Driver azure-disk doesn't support ntfs -- skipping ([#96144](https://github.com/kubernetes/kubernetes/pull/96144), [@qinpingli](https://github.com/qinpingli)) [SIG Storage and Testing]
- StatefulSet Controller now waits for PersistentVolumeClaim deletion before creating pods. ([#93457](https://github.com/kubernetes/kubernetes/pull/93457), [@ymmt2005](https://github.com/ymmt2005))
- StreamWatcher now calls HandleCrash at appropriate sequence. ([#93108](https://github.com/kubernetes/kubernetes/pull/93108), [@lixiaobing1](https://github.com/lixiaobing1))
- Support the node label `node.kubernetes.io/exclude-from-external-load-balancers` ([#95542](https://github.com/kubernetes/kubernetes/pull/95542), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- The AWS network load balancer attributes can now be specified during service creation ([#95247](https://github.com/kubernetes/kubernetes/pull/95247), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- The `/debug/api_priority_and_fairness/dump_requests` path at an apiserver will no longer return a phantom line for each exempt priority level. ([#93406](https://github.com/kubernetes/kubernetes/pull/93406), [@MikeSpreitzer](https://github.com/MikeSpreitzer)) [SIG API Machinery]
- The kube-apiserver will no longer serve APIs that should have been deleted in GA non-alpha levels.  Alpha levels will continue to serve the removed APIs so that CI doesn't immediately break. ([#96525](https://github.com/kubernetes/kubernetes/pull/96525), [@deads2k](https://github.com/deads2k)) [SIG API Machinery]
- The kubelet recognizes the --containerd-namespace flag to configure the namespace used by cadvisor. ([#87054](https://github.com/kubernetes/kubernetes/pull/87054), [@changyaowei](https://github.com/changyaowei)) [SIG Node]
- Unhealthy pods covered by PDBs can be successfully evicted if enough healthy pods are available. ([#94381](https://github.com/kubernetes/kubernetes/pull/94381), [@michaelgugino](https://github.com/michaelgugino)) [SIG Apps]
- Update Calico to v3.15.2 ([#94241](https://github.com/kubernetes/kubernetes/pull/94241), [@lmm](https://github.com/lmm)) [SIG Cloud Provider]
- Update default etcd server version to 3.4.13 ([#94287](https://github.com/kubernetes/kubernetes/pull/94287), [@jingyih](https://github.com/jingyih)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle and Testing]
- Update max azure data disk count map ([#96308](https://github.com/kubernetes/kubernetes/pull/96308), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Update the PIP when it is not in the Succeeded provisioning state during the LB update. ([#95748](https://github.com/kubernetes/kubernetes/pull/95748), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Update the frontend IP config when the service's `pipName` annotation is changed ([#95813](https://github.com/kubernetes/kubernetes/pull/95813), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Update the route table tag in the route reconcile loop ([#96545](https://github.com/kubernetes/kubernetes/pull/96545), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Use NLB Subnet CIDRs instead of VPC CIDRs in Health Check SG Rules ([#93515](https://github.com/kubernetes/kubernetes/pull/93515), [@t0rr3sp3dr0](https://github.com/t0rr3sp3dr0)) [SIG Cloud Provider]
- Users will see increase in time for deletion of pods and also guarantee that removal of pod from api server  would mean deletion of all the resources from container runtime. ([#92817](https://github.com/kubernetes/kubernetes/pull/92817), [@kmala](https://github.com/kmala)) [SIG Node]
- Very large patches may now be specified to `kubectl patch` with the `--patch-file` flag instead of including them directly on the command line. The `--patch` and `--patch-file` flags are mutually exclusive. ([#93548](https://github.com/kubernetes/kubernetes/pull/93548), [@smarterclayton](https://github.com/smarterclayton)) [SIG CLI]
- Volume binding: report UnschedulableAndUnresolvable status instead of an error when bound PVs not found ([#95541](https://github.com/kubernetes/kubernetes/pull/95541), [@cofyc](https://github.com/cofyc)) [SIG Apps, Scheduling and Storage]
- Warn instead of fail when creating Roles and ClusterRoles with custom verbs via kubectl ([#92492](https://github.com/kubernetes/kubernetes/pull/92492), [@eddiezane](https://github.com/eddiezane)) [SIG CLI]
- When creating a PVC with the volume.beta.kubernetes.io/storage-provisioner annotation already set, the PV controller might have incorrectly deleted the newly provisioned PV instead of binding it to the PVC, depending on timing and system load. ([#95909](https://github.com/kubernetes/kubernetes/pull/95909), [@pohly](https://github.com/pohly)) [SIG Apps and Storage]
- [kubectl] Fail when local source file doesn't exist ([#90333](https://github.com/kubernetes/kubernetes/pull/90333), [@bamarni](https://github.com/bamarni)) [SIG CLI]




#### Other (Cleanup or Flake)

- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102341](https://github.com/kubernetes/kubernetes/pull/102341), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]
- Kubeadm: change the default image repository for CI images from 'gcr.io/kubernetes-ci-images' to 'gcr.io/k8s-staging-ci-images' ([#97087](https://github.com/kubernetes/kubernetes/pull/97087), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Resolves flakes in the Ingress conformance tests due to conflicts with controllers updating the Ingress object ([#98430](https://github.com/kubernetes/kubernetes/pull/98430), [@liggitt](https://github.com/liggitt)) [SIG Network and Testing]
- Handle slow CronJob lister in CronJob controller v2 and improve memory footprint. ([#96443](https://github.com/kubernetes/kubernetes/pull/96443), [@alaypatel07](https://github.com/alaypatel07)) [SIG Apps]
- --redirect-container-streaming is no longer functional. The flag will be removed in v1.22 ([#95935](https://github.com/kubernetes/kubernetes/pull/95935), [@tallclair](https://github.com/tallclair)) [SIG Node]
- A new metric `requestAbortsTotal` has been introduced that counts aborted requests for each `group`, `version`, `verb`, `resource`, `subresource` and `scope`. ([#95002](https://github.com/kubernetes/kubernetes/pull/95002), [@p0lyn0mial](https://github.com/p0lyn0mial)) [SIG API Machinery, Cloud Provider, Instrumentation and Scheduling]
- API priority and fairness metrics use snake_case in label names ([#96236](https://github.com/kubernetes/kubernetes/pull/96236), [@adtac](https://github.com/adtac)) [SIG API Machinery, Cluster Lifecycle, Instrumentation and Testing]
- Add fine-grained debugging to intra-pod conformance test to troubleshoot networking issues for potentially unhealthy nodes when running conformance or sonobuoy tests. ([#93837](https://github.com/kubernetes/kubernetes/pull/93837), [@jayunit100](https://github.com/jayunit100))
- Add the following metrics:
  - network_plugin_operations_total
  - network_plugin_operations_errors_total ([#93066](https://github.com/kubernetes/kubernetes/pull/93066), [@AnishShah](https://github.com/AnishShah))
- Adds a bootstrapping ClusterRole, ClusterRoleBinding and group for /metrics, /livez/*, /readyz/*, & /healthz/- endpoints. ([#93311](https://github.com/kubernetes/kubernetes/pull/93311), [@logicalhan](https://github.com/logicalhan)) [SIG API Machinery, Auth, Cloud Provider and Instrumentation]
- AdmissionReview objects sent for the creation of Namespace API objects now populate the `namespace` attribute consistently (previously the `namespace` attribute was empty for Namespace creation via POST requests, and populated for Namespace creation via server-side-apply PATCH requests) ([#95012](https://github.com/kubernetes/kubernetes/pull/95012), [@nodo](https://github.com/nodo)) [SIG API Machinery and Testing]
- Applies translations on all command descriptions ([#95439](https://github.com/kubernetes/kubernetes/pull/95439), [@HerrNaN](https://github.com/HerrNaN)) [SIG CLI]
- Base-images: Update to debian-iptables:buster-v1.3.0
  - Uses iptables 1.8.5
  - base-images: Update to debian-base:buster-v1.2.0
  - cluster/images/etcd: Build etcd:3.4.13-1 image
    - Uses debian-base:buster-v1.2.0 ([#94733](https://github.com/kubernetes/kubernetes/pull/94733), [@justaugustus](https://github.com/justaugustus)) [SIG API Machinery, Release and Testing]
- Changed: default "Accept-Encoding" header removed from HTTP probes. See https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/#http-probes ([#96127](https://github.com/kubernetes/kubernetes/pull/96127), [@fonsecas72](https://github.com/fonsecas72)) [SIG Network and Node]
- Client-go header logging (at verbosity levels >= 9) now masks `Authorization` header contents ([#95316](https://github.com/kubernetes/kubernetes/pull/95316), [@sfowl](https://github.com/sfowl)) [SIG API Machinery]
- Decrease warning message frequency on setting volume ownership for configmap/secret. ([#92878](https://github.com/kubernetes/kubernetes/pull/92878), [@jvanz](https://github.com/jvanz))
- Enhance log information of verifyRunAsNonRoot, add pod, container information ([#94911](https://github.com/kubernetes/kubernetes/pull/94911), [@wawa0210](https://github.com/wawa0210)) [SIG Node]
- Fix func name NewCreateCreateDeploymentOptions ([#91931](https://github.com/kubernetes/kubernetes/pull/91931), [@lixiaobing1](https://github.com/lixiaobing1)) [SIG CLI]
- Fix kubelet to properly log when a container is started. Previously, kubelet may log that container is dead and was restarted when it was actually started for the first time. This behavior only happened on pods with initContainers and regular containers. ([#91469](https://github.com/kubernetes/kubernetes/pull/91469), [@rata](https://github.com/rata))
- Fixes the message about no auth for metrics in scheduler. ([#94035](https://github.com/kubernetes/kubernetes/pull/94035), [@zhouya0](https://github.com/zhouya0)) [SIG Scheduling]
- Generators for services are removed from kubectl ([#95256](https://github.com/kubernetes/kubernetes/pull/95256), [@Git-Jiro](https://github.com/Git-Jiro)) [SIG CLI]
- Introduce kubectl-convert plugin. ([#96190](https://github.com/kubernetes/kubernetes/pull/96190), [@soltysh](https://github.com/soltysh)) [SIG CLI and Testing]
- Kube-scheduler now logs processed component config at startup ([#96426](https://github.com/kubernetes/kubernetes/pull/96426), [@damemi](https://github.com/damemi)) [SIG Scheduling]
- Kubeadm: Separate argument key/value in log msg ([#94016](https://github.com/kubernetes/kubernetes/pull/94016), [@mrueg](https://github.com/mrueg)) [SIG Cluster Lifecycle]
- Kubeadm: remove the CoreDNS check for known image digests when applying the addon ([#94506](https://github.com/kubernetes/kubernetes/pull/94506), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: update the default pause image version to 1.4.0 on Windows. With this update the image supports Windows versions 1809 (2019LTS), 1903, 1909, 2004 ([#95419](https://github.com/kubernetes/kubernetes/pull/95419), [@jsturtevant](https://github.com/jsturtevant)) [SIG Cluster Lifecycle and Windows]
- Kubectl: the `generator` flag of `kubectl autoscale` has been deprecated and has no effect, it will be removed in a feature release ([#92998](https://github.com/kubernetes/kubernetes/pull/92998), [@SataQiu](https://github.com/SataQiu)) [SIG CLI]
- Lock ExternalPolicyForExternalIP to default, this feature gate will be removed in 1.22. ([#94581](https://github.com/kubernetes/kubernetes/pull/94581), [@knabben](https://github.com/knabben)) [SIG Network]
- Mask ceph RBD adminSecrets in logs when logLevel >= 4. ([#95245](https://github.com/kubernetes/kubernetes/pull/95245), [@sfowl](https://github.com/sfowl))
- Remove offensive words from kubectl cluster-info command. ([#95202](https://github.com/kubernetes/kubernetes/pull/95202), [@rikatz](https://github.com/rikatz))
- Remove support for "ci/k8s-master" version label in kubeadm, use "ci/latest" instead. See [kubernetes/test-infra#18517](https://github.com/kubernetes/test-infra/pull/18517). ([#93626](https://github.com/kubernetes/kubernetes/pull/93626), [@vikkyomkar](https://github.com/vikkyomkar))
- Remove the dependency of csi-translation-lib module on apiserver/cloud-provider/controller-manager ([#95543](https://github.com/kubernetes/kubernetes/pull/95543), [@wawa0210](https://github.com/wawa0210)) [SIG Release]
- Scheduler framework interface moved from pkg/scheduler/framework/v1alpha to pkg/scheduler/framework ([#95069](https://github.com/kubernetes/kubernetes/pull/95069), [@farah](https://github.com/farah)) [SIG Scheduling, Storage and Testing]
- Service.beta.kubernetes.io/azure-load-balancer-disable-tcp-reset is removed.  All Standard load balancers will always enable tcp resets. ([#94297](https://github.com/kubernetes/kubernetes/pull/94297), [@MarcPow](https://github.com/MarcPow)) [SIG Cloud Provider]
- Stop propagating SelfLink (deprecated in 1.16) in kube-apiserver ([#94397](https://github.com/kubernetes/kubernetes/pull/94397), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Testing]
- Strip unnecessary security contexts on Windows ([#93475](https://github.com/kubernetes/kubernetes/pull/93475), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG Node, Testing and Windows]
- To ensure the code be strong,  add unit test for GetAddressAndDialer ([#93180](https://github.com/kubernetes/kubernetes/pull/93180), [@FreeZhang61](https://github.com/FreeZhang61)) [SIG Node]
- UDP and SCTP protocols can left stale connections that need to be cleared to avoid services disruption, but they can cause problems that are hard to debug.
  Kubernetes components using a loglevel greater or equal than 4 will log the conntrack operations and its output, to show the entries that were deleted. ([#95694](https://github.com/kubernetes/kubernetes/pull/95694), [@aojea](https://github.com/aojea)) [SIG Network]
- Update CNI plugins to v0.8.7 ([#94367](https://github.com/kubernetes/kubernetes/pull/94367), [@justaugustus](https://github.com/justaugustus)) [SIG Cloud Provider, Network, Node, Release and Testing]
- Update cri-tools to [v1.19.0](https://github.com/kubernetes-sigs/cri-tools/releases/tag/v1.19.0) ([#94307](https://github.com/kubernetes/kubernetes/pull/94307), [@xmudrii](https://github.com/xmudrii)) [SIG Cloud Provider]
- Update etcd client side to v3.4.13 ([#94259](https://github.com/kubernetes/kubernetes/pull/94259), [@jingyih](https://github.com/jingyih)) [SIG API Machinery and Cloud Provider]
- Users will now be able to configure all supported values for AWS NLB health check interval and thresholds for new resources. ([#96312](https://github.com/kubernetes/kubernetes/pull/96312), [@kishorj](https://github.com/kishorj)) [SIG Cloud Provider]
- V1helpers.MatchNodeSelectorTerms now accepts just a Node and a list of Terms ([#95871](https://github.com/kubernetes/kubernetes/pull/95871), [@damemi](https://github.com/damemi)) [SIG Apps, Scheduling and Storage]
- vSphere: improve logging message on node cache refresh event ([#95236](https://github.com/kubernetes/kubernetes/pull/95236), [@andrewsykim](https://github.com/andrewsykim)) [SIG Cloud Provider]
- `MatchNodeSelectorTerms` function moved to `k8s.io/component-helpers` ([#95531](https://github.com/kubernetes/kubernetes/pull/95531), [@damemi](https://github.com/damemi)) [SIG Apps, Scheduling and Storage]
- `kubectl api-resources` now prints the API version (as 'API group/version', same as output of `kubectl api-versions`). The column APIGROUP is now APIVERSION ([#95253](https://github.com/kubernetes/kubernetes/pull/95253), [@sallyom](https://github.com/sallyom)) [SIG CLI]
- `kubectl get ingress` now prefers the `networking.k8s.io/v1` over `extensions/v1beta1` (deprecated since v1.14). To explicitly request the deprecated version, use `kubectl get ingress.v1beta1.extensions`. ([#94309](https://github.com/kubernetes/kubernetes/pull/94309), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and CLI]



#### Dependencies

##### Added
- cloud.google.com/go/firestore: v1.1.0
- github.com/Azure/go-autorest: [v14.2.0+incompatible](https://github.com/Azure/go-autorest/tree/v14.2.0)
- github.com/armon/go-metrics: [f0300d1](https://github.com/armon/go-metrics/tree/f0300d1)
- github.com/armon/go-radix: [7fddfc3](https://github.com/armon/go-radix/tree/7fddfc3)
- github.com/bketelsen/crypt: [5cbc8cc](https://github.com/bketelsen/crypt/tree/5cbc8cc)
- github.com/form3tech-oss/jwt-go: [v3.2.2+incompatible](https://github.com/form3tech-oss/jwt-go/tree/v3.2.2)
- github.com/fvbommel/sortorder: [v1.0.1](https://github.com/fvbommel/sortorder/tree/v1.0.1)
- github.com/hashicorp/consul/api: [v1.1.0](https://github.com/hashicorp/consul/api/tree/v1.1.0)
- github.com/hashicorp/consul/sdk: [v0.1.1](https://github.com/hashicorp/consul/sdk/tree/v0.1.1)
- github.com/hashicorp/errwrap: [v1.0.0](https://github.com/hashicorp/errwrap/tree/v1.0.0)
- github.com/hashicorp/go-cleanhttp: [v0.5.1](https://github.com/hashicorp/go-cleanhttp/tree/v0.5.1)
- github.com/hashicorp/go-immutable-radix: [v1.0.0](https://github.com/hashicorp/go-immutable-radix/tree/v1.0.0)
- github.com/hashicorp/go-msgpack: [v0.5.3](https://github.com/hashicorp/go-msgpack/tree/v0.5.3)
- github.com/hashicorp/go-multierror: [v1.0.0](https://github.com/hashicorp/go-multierror/tree/v1.0.0)
- github.com/hashicorp/go-rootcerts: [v1.0.0](https://github.com/hashicorp/go-rootcerts/tree/v1.0.0)
- github.com/hashicorp/go-sockaddr: [v1.0.0](https://github.com/hashicorp/go-sockaddr/tree/v1.0.0)
- github.com/hashicorp/go-uuid: [v1.0.1](https://github.com/hashicorp/go-uuid/tree/v1.0.1)
- github.com/hashicorp/go.net: [v0.0.1](https://github.com/hashicorp/go.net/tree/v0.0.1)
- github.com/hashicorp/logutils: [v1.0.0](https://github.com/hashicorp/logutils/tree/v1.0.0)
- github.com/hashicorp/mdns: [v1.0.0](https://github.com/hashicorp/mdns/tree/v1.0.0)
- github.com/hashicorp/memberlist: [v0.1.3](https://github.com/hashicorp/memberlist/tree/v0.1.3)
- github.com/hashicorp/serf: [v0.8.2](https://github.com/hashicorp/serf/tree/v0.8.2)
- github.com/jmespath/go-jmespath/internal/testify: [v1.5.1](https://github.com/jmespath/go-jmespath/releases/tag/internal%2Ftestify%2Fv1.5.1)
- github.com/mitchellh/cli: [v1.0.0](https://github.com/mitchellh/cli/tree/v1.0.0)
- github.com/mitchellh/go-testing-interface: [v1.0.0](https://github.com/mitchellh/go-testing-interface/tree/v1.0.0)
- github.com/mitchellh/gox: [v0.4.0](https://github.com/mitchellh/gox/tree/v0.4.0)
- github.com/mitchellh/iochan: [v1.0.0](https://github.com/mitchellh/iochan/tree/v1.0.0)
- github.com/pascaldekloe/goe: [57f6aae](https://github.com/pascaldekloe/goe/tree/57f6aae)
- github.com/posener/complete: [v1.1.1](https://github.com/posener/complete/tree/v1.1.1)
- github.com/ryanuber/columnize: [9b3edd6](https://github.com/ryanuber/columnize/tree/9b3edd6)
- github.com/sean-/seed: [e2103e2](https://github.com/sean-/seed/tree/e2103e2)
- github.com/subosito/gotenv: [v1.2.0](https://github.com/subosito/gotenv/tree/v1.2.0)
- github.com/willf/bitset: [d5bec33](https://github.com/willf/bitset/tree/d5bec33)
- gopkg.in/ini.v1: v1.51.0
- gopkg.in/yaml.v3: 9f266ea
- rsc.io/quote/v3: v3.1.0
- rsc.io/sampler: v1.3.0

##### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.9 → v0.0.19
- github.com/gogo/protobuf: [v1.3.1 → v1.3.2](https://github.com/gogo/protobuf/compare/v1.3.1...v1.3.2)
- github.com/kisielk/errcheck: [v1.2.0 → v1.5.0](https://github.com/kisielk/errcheck/compare/v1.2.0...v1.5.0)
- github.com/yuin/goldmark: [v1.1.27 → v1.2.1](https://github.com/yuin/goldmark/compare/v1.1.27...v1.2.1)
- golang.org/x/sync: cd5d95a → 67f06af
- golang.org/x/tools: c1934b7 → 113979e
- sigs.k8s.io/structured-merge-diff/v4: v4.0.1 → v4.0.3
- github.com/google/cadvisor: [v0.37.0 → v0.38.8](https://github.com/google/cadvisor/compare/v0.37.0...v0.38.8)
- cloud.google.com/go/bigquery: v1.0.1 → v1.4.0
- cloud.google.com/go/datastore: v1.0.0 → v1.1.0
- cloud.google.com/go/pubsub: v1.0.1 → v1.2.0
- cloud.google.com/go/storage: v1.0.0 → v1.6.0
- cloud.google.com/go: v0.51.0 → v0.54.0
- github.com/Azure/go-autorest/autorest/adal: [v0.8.2 → v0.9.5](https://github.com/Azure/go-autorest/compare/autorest/adal/v0.8.2...autorest/adal/v0.9.5)
- github.com/Azure/go-autorest/autorest/date: [v0.2.0 → v0.3.0](https://github.com/Azure/go-autorest/compare/autorest/date/v0.2.0...autorest/date/v0.3.0)
- github.com/Azure/go-autorest/autorest/mocks: [v0.3.0 → v0.4.1](https://github.com/Azure/go-autorest/compare/autorest/mocks/v0.3.0...autorest/mocks/v0.4.1)
- github.com/Azure/go-autorest/autorest: [v0.9.6 → v0.11.1](https://github.com/Azure/go-autorest/compare/autorest/v0.9.6...autorest/v0.11.1)
- github.com/Azure/go-autorest/logger: [v0.1.0 → v0.2.0](https://github.com/Azure/go-autorest/compare/logger/v0.1.0...logger/v0.2.0)
- github.com/Azure/go-autorest/tracing: [v0.5.0 → v0.6.0](https://github.com/Azure/go-autorest/compare/tracing/v0.5.0...tracing/v0.6.0)
- github.com/Microsoft/go-winio: [fc70bd9 → v0.4.15](https://github.com/Microsoft/go-winio/compare/fc70bd9...v0.4.15)
- github.com/aws/aws-sdk-go: [v1.28.2 → v1.35.24](https://github.com/aws/aws-sdk-go/compare/v1.28.2...v1.35.24)
- github.com/blang/semver: [v3.5.0+incompatible → v3.5.1+incompatible](https://github.com/blang/semver/compare/v3.5.0...v3.5.1)
- github.com/checkpoint-restore/go-criu/v4: [v4.0.2 → v4.1.0](https://github.com/checkpoint-restore/go-criu/v4/compare/v4.0.2...v4.1.0)
- github.com/containerd/containerd: [v1.3.3 → v1.4.1](https://github.com/containerd/containerd/compare/v1.3.3...v1.4.1)
- github.com/containerd/ttrpc: [v1.0.0 → v1.0.2](https://github.com/containerd/ttrpc/compare/v1.0.0...v1.0.2)
- github.com/containerd/typeurl: [v1.0.0 → v1.0.1](https://github.com/containerd/typeurl/compare/v1.0.0...v1.0.1)
- github.com/coreos/etcd: [v3.3.10+incompatible → v3.3.13+incompatible](https://github.com/coreos/etcd/compare/v3.3.10...v3.3.13)
- github.com/docker/docker: [aa6a989 → bd33bbf](https://github.com/docker/docker/compare/aa6a989...bd33bbf)
- github.com/go-gl/glfw/v3.3/glfw: [12ad95a → 6f7a984](https://github.com/go-gl/glfw/v3.3/glfw/compare/12ad95a...6f7a984)
- github.com/golang/groupcache: [215e871 → 8c9f03a](https://github.com/golang/groupcache/compare/215e871...8c9f03a)
- github.com/golang/mock: [v1.3.1 → v1.4.1](https://github.com/golang/mock/compare/v1.3.1...v1.4.1)
- github.com/golang/protobuf: [v1.4.2 → v1.4.3](https://github.com/golang/protobuf/compare/v1.4.2...v1.4.3)
- github.com/google/go-cmp: [v0.4.0 → v0.5.2](https://github.com/google/go-cmp/compare/v0.4.0...v0.5.2)
- github.com/google/pprof: [d4f498a → 1ebb73c](https://github.com/google/pprof/compare/d4f498a...1ebb73c)
- github.com/google/uuid: [v1.1.1 → v1.1.2](https://github.com/google/uuid/compare/v1.1.1...v1.1.2)
- github.com/gorilla/mux: [v1.7.3 → v1.8.0](https://github.com/gorilla/mux/compare/v1.7.3...v1.8.0)
- github.com/gorilla/websocket: [v1.4.0 → v1.4.2](https://github.com/gorilla/websocket/compare/v1.4.0...v1.4.2)
- github.com/jmespath/go-jmespath: [c2b33e8 → v0.4.0](https://github.com/jmespath/go-jmespath/compare/c2b33e8...v0.4.0)
- github.com/karrick/godirwalk: [v1.7.5 → v1.16.1](https://github.com/karrick/godirwalk/compare/v1.7.5...v1.16.1)
- github.com/opencontainers/go-digest: [v1.0.0-rc1 → v1.0.0](https://github.com/opencontainers/go-digest/compare/v1.0.0-rc1...v1.0.0)
- github.com/opencontainers/runc: [819fcc6 → v1.0.0-rc92](https://github.com/opencontainers/runc/compare/819fcc6...v1.0.0-rc92)
- github.com/opencontainers/runtime-spec: [237cc4f → 4d89ac9](https://github.com/opencontainers/runtime-spec/compare/237cc4f...4d89ac9)
- github.com/opencontainers/selinux: [v1.5.2 → v1.6.0](https://github.com/opencontainers/selinux/compare/v1.5.2...v1.6.0)
- github.com/prometheus/procfs: [v0.1.3 → v0.2.0](https://github.com/prometheus/procfs/compare/v0.1.3...v0.2.0)
- github.com/quobyte/api: [v0.1.2 → v0.1.8](https://github.com/quobyte/api/compare/v0.1.2...v0.1.8)
- github.com/spf13/cobra: [v1.0.0 → v1.1.1](https://github.com/spf13/cobra/compare/v1.0.0...v1.1.1)
- github.com/spf13/viper: [v1.4.0 → v1.7.0](https://github.com/spf13/viper/compare/v1.4.0...v1.7.0)
- github.com/storageos/go-api: [343b3ef → v2.2.0+incompatible](https://github.com/storageos/go-api/compare/343b3ef...v2.2.0)
- github.com/stretchr/testify: [v1.4.0 → v1.6.1](https://github.com/stretchr/testify/compare/v1.4.0...v1.6.1)
- github.com/vishvananda/netns: [52d707b → db3c7e5](https://github.com/vishvananda/netns/compare/52d707b...db3c7e5)
- go.etcd.io/etcd: 17cef6e → dd1b699
- go.opencensus.io: v0.22.2 → v0.22.3
- golang.org/x/crypto: 75b2880 → 7f63de1
- golang.org/x/exp: da58074 → 6cc2880
- golang.org/x/lint: fdd1cda → 738671d
- golang.org/x/net: ab34263 → 69a7880
- golang.org/x/oauth2: 858c2ad → bf48bf1
- golang.org/x/sys: ed371f2 → 5cba982
- golang.org/x/text: v0.3.3 → v0.3.4
- golang.org/x/time: 555d28b → 3af7569
- golang.org/x/xerrors: 9bdfabe → 5ec99f8
- google.golang.org/api: v0.15.1 → v0.20.0
- google.golang.org/genproto: cb27e3a → 8816d57
- google.golang.org/grpc: v1.27.0 → v1.27.1
- google.golang.org/protobuf: v1.24.0 → v1.25.0
- honnef.co/go/tools: v0.0.1-2019.2.3 → v0.0.1-2020.1.3
- k8s.io/gengo: 8167cfd → 83324d8
- k8s.io/klog/v2: v2.2.0 → v2.4.0
- k8s.io/kube-openapi: 6aeccd4 → d219536
- k8s.io/system-validators: v1.1.2 → v1.2.0
- k8s.io/utils: d5654de → 67b214c

##### Removed
- github.com/armon/consul-api: [eb2c6b5](https://github.com/armon/consul-api/tree/eb2c6b5)
- github.com/go-ini/ini: [v1.9.0](https://github.com/go-ini/ini/tree/v1.9.0)
- github.com/ugorji/go: [v1.1.4](https://github.com/ugorji/go/tree/v1.1.4)
- github.com/xlab/handysort: [fb3537e](https://github.com/xlab/handysort/tree/fb3537e)
- github.com/xordataexchange/crypt: [b2862e3](https://github.com/xordataexchange/crypt/tree/b2862e3)
- vbom.ml/util: db5cfe1


### aws-ebs-csi-driver [2.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.1.0)

#### Changed

- Update aws-ebs-csi-driver to v1.1.0.



### aws-cni [1.8.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.8.0)

Changes since v1.7.10:

* Bug - Use symmetric return path for non-VPC traffic - alternate solution (#1475, @kishorj)
* Bug - Gracefully handle failed ENI SG update (#1341, @jayanthvn)
* Bug - Fix CNI crashing when there is no available IP addresses (#1499, @M00nF1sh)
* Bug - Use primary ENI SGs if SG is null for Custom networking (#1259, @jayanthvn)
* Bug - Don't cache dynamic VPC IPv4 CIDR info (#1113, @anguslees)
* Improvement - Address Excessive API Server calls from CNI Pods (#1419, @achevuru)
* Improvement - refine ENI tagging logic (#1482, @M00nF1sh)
* Improvement - Change tryAssignIPs to assign up to configured WARM_IP_TARGET (#1279, @jacksontj)
* Improvement - Use regional STS endpoint (#1332, @nithu0115)
* Improvement - Update containernetworking dependencies (#1200, @mogren)
* Improvement - Split Calico manifest into two (#1410, @caseydavenport)
* Improvement - Update Calico manifest to support ARM & AMD (#1282, @jayanthvn)
* Improvement - Auto gen of AWS CNI, metrics helper and calico artifacts through helm (#1271, @jayanthvn)
* Improvement - Refactor EC2 Metadata IMDS code (#1225, @anguslees)
* Improvement - Unnecessary logging for each CNI invocation (#1469, @jayanthvn)
* Improvement - New instance types (#1463, @jayanthvn)
* Improvement - Use 'exec' ENTRYPOINTs (#1432, @anguslees)
* Improvement - Fix logging texts for ENI cleanup (#1209, @mogren)
* Improvement - Remove Duplicated vlan IPTable rules (#1208, @mogren)
* Improvement - Minor code cleanup (#1198, @mogren)
* HelmChart - Adding flags to support overriding container runtime endpoint. (#1443, @haouc)
* HelmChart - Add podLabels to amazon-vpc-cni chart (#1440, @haouc)
* HelmChart - Add workflow to sync aws-vpc-cni helm chart to eks-charts (#1430, @fawadkhaliq)
* Testing - Remove validation of VPC CIDRs from ip rules (#1476, @kishorj)
* Testing - Updated agent version (#1474, @cgchinmay)
* Testing - Fix for CI failure (#1470, @achevuru)
* Testing - Binary for mtu and veth prefix check (#1458, @cgchinmay)
* Testing - add test to verify cni-metrics-helper puts metrics to CW (#1461, @abhipth)
* Testing - add e2e test for security group for pods (#1459, @abhipth)
* Testing - Added Test cases for EnvVars check on CNI daemonset (#1431, @cgchinmay)
* Testing - add test to verify host networking setup & cleanup (#1457, @abhipth)
* Testing - Runners failing because of docker permissions (#1456, @jayanthvn)
* Testing - decouple test helper input struct from netlink library (#1455, @abhipth)
* Testing - add custom networking e2e test suite (#1445, @abhipth)
* Testing - add integration test for ipamd env variables (#1453, @abhipth)
* Testing - add agent for testing pod networking (#1448, @abhipth)
* Testing - fix format of commited code to fix unit test step (#1449, @abhipth)
* Testing - Unblocks Github Action Integration Tests (#1435, @couralex6)
* Testing - add warm ENI/IP target integration tests (#1438, @abhipth)
* Testing - add service connectivity test (#1436, @abhipth)
* Testing - add network connectivity test (#1424, @abhipth)
* Testing - add ginkgo automation framework (#1416, @abhipth)
* Testing - Add some test coverage to allocating ENIs (#1234, @mogren)
* Testing - Add some minimal tests to metrics (#1228, @mogren)

Changes since v1.7.9:

* Improvement - Multi card support - Prevent route override for primary ENI across multi-cards ENAs (#1396 , [@jayanthvn](https://github.com/Jayanthvn))

Changes since v1.7.8:
* Improvement - Adds http timeout to aws sessions (#1370 by couralex6)
* Improvement - Switch calico to be deployed with the Tigera operator (#1297 by tmjd)
* Improvement - Update calico to v3.17.1 (#1328 by lwr20)
* Improvement - update plugins to v0.9.0 (#1362 by fr0stbyte)
* Improvement - update github.com/containernetworking/plugins to v0.9.0 (#1350 by fr0stbyte)
* Bug - Fix regex match for getting primary interface (#1311 by Jayanthvn)
* Bug - Output to stderr when no log file path is passed (#1275 by couralex6)
* Bug - Fix deletion of hostVeth rule for pods using security group (#1376 by SaranBalaji90)



### cluster-autoscaler [1.20.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.3)

#### Changed
- Allow users to set container resources;
- Update cluster-autoscaler to version `1.20.0`.
