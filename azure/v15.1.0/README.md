# :zap: Giant Swarm Release v15.1.0 for Azure :zap:

This is the first Giant Swarm Azure stable release featuring Kubernetes [1.20.8](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.20.md#v1208).

This release introduces a new feature allowing configuration of the public IP address for the NAT gateway of worker nodes for existing and new clusters. For more details please talk to your Account Engineer or follow [Giant Swarm Documentation](https://docs.giantswarm.io/advanced/egress-ip-address-azure/).


## Change details


### azure-operator [5.8.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.8.0)

#### Added
- Allow using an existing public IP for the NAT gateway of worker nodes.
#### Fixed
- Fix udev rules that caused `/boot` automount to fail.
#### Changed
- Upgrade `k8scloudconfig` to `v10.8.1` from `v10.5.0`.



### kubernetes [1.20.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.8)

#### Feature
- Kubernetes is now built using Go 1.15.13 ([#102786](https://github.com/kubernetes/kubernetes/pull/102786), [@thejoycekung](https://github.com/thejoycekung)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Failing Test
- Fixes the `should receive events on concurrent watches in same order` conformance test to work properly on clusters that auto-create additional configmaps in namespaces ([#101950](https://github.com/kubernetes/kubernetes/pull/101950), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
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
#### Other (Cleanup or Flake)
- Update the Debian images to pick up CVE fixes in the base images:
  - Update the `debian-base` image to v1.7.0
  - Update the `debian-iptables` image to v1.6.1 ([#102341](https://github.com/kubernetes/kubernetes/pull/102341), [@cpanato](https://github.com/cpanato)) [SIG API Machinery and Testing]_
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.15 → v0.0.19


### containerlinux [2765.2.6](https://www.flatcar-linux.org/releases/#release-2765.2.6)

#### Security fixes
- Linux ([CVE-2020-26558](https://nvd.nist.gov/vuln/detail/CVE-2020-26558), [CVE-2021-0129](https://nvd.nist.gov/vuln/detail/CVE-2021-0129), [CVE-2020-24587](https://nvd.nist.gov/vuln/detail/CVE-2020-24587), [CVE-2020-24586](https://nvd.nist.gov/vuln/detail/CVE-2020-24586), [CVE-2020-24588](https://nvd.nist.gov/vuln/detail/CVE-2020-24588), [CVE-2020-26139](https://nvd.nist.gov/vuln/detail/CVE-2020-26139), [CVE-2020-26145](https://nvd.nist.gov/vuln/detail/CVE-2020-26145), [CVE-2020-26147](https://nvd.nist.gov/vuln/detail/CVE-2020-26147), [CVE-2020-26141](https://nvd.nist.gov/vuln/detail/CVE-2020-26141), [CVE-2021-3564](https://nvd.nist.gov/vuln/detail/CVE-2021-3564), [CVE-2021-28691](https://nvd.nist.gov/vuln/detail/CVE-2021-28691), [CVE-2021-3587](https://nvd.nist.gov/vuln/detail/CVE-2021-3587), [CVE-2021-3573](https://nvd.nist.gov/vuln/detail/CVE-2021-3573))
#### Bug fixes
- Update-engine sent empty requests when restarted before a pending reboot ([Flatcar#388](https://github.com/kinvolk/Flatcar/issues/388))
- motd login prompt list of failed services: The output of "systemctl list-units --state=failed --no-legend" contains a bullet point which is not expected and ended up being taken as the unit name of failed units which was previously on the start of the line. Filtered the bullet point out to stay compatible with the old behavior in case upstream would remove the bullet point again. ([coreos-overlay#1042](https://github.com/kinvolk/coreos-overlay/pull/1042))
#### Updates
- Linux ([5.10.43](https://lwn.net/Articles/859022/))


### cert-exporter [1.7.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.7.1)

#### Fixed
- Fix configuration version in `Chart.yaml`.


### chart-operator [2.18.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.18.0)

#### Added
- Add releasemaxhistory resource which ensures we retry at a reduced rate when
there are repeated failed upgrades.
#### Changed
- Upgrade Helm release when failed even if version or values have not changed
to handle situations like failed webhooks where we should retry.



### external-dns [2.4.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.4.0)

#### Changed
- Upgrade upstream external-dns from v0.7.6 to [v0.8.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.8.0).
- Allow to configure the minimum interval between two consecutive synchronizations triggered from kubernetes events through `externalDNS.minEventSyncInterval`.



### net-exporter [1.10.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.2)

#### Changed
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.



### cluster-autoscaler [1.20.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.20.3)

- Allow users to set container resources
