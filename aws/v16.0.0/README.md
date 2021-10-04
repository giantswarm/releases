# :zap: Giant Swarm Release v16.0.0 for AWS :zap:

<< Add description here >>

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

#### Important Security Information

This release contains changes that address the following vulnerabilities:

##### CVE-2021-25741: Symlink Exchange Can Allow Host Filesystem Access

A security issue was discovered in Kubernetes where a user may be able to
create a container with subpath volume mounts to access files &
directories outside of the volume, including on the host filesystem.

#### Feature
- Kubernetes is now built with Golang 1.16.8 ([#104906](https://github.com/kubernetes/kubernetes/pull/104906), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update the setcap image to buster-v2.0.1 ([#102377](https://github.com/kubernetes/kubernetes/pull/102377), [@xmudrii](https://github.com/xmudrii)) [SIG Release]

#### Failing Test
- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]

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



#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]
- Client-go: reduce verbosity of "Starting/Stopping reflector" messages to 3 again ([#102788](https://github.com/kubernetes/kubernetes/pull/102788), [@pohly](https://github.com/pohly)) [SIG API Machinery]
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102340](https://github.com/kubernetes/kubernetes/pull/102340), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]



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
