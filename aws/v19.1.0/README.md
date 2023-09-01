# :zap: Giant Swarm Release v19.1.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [6.8.0](https://github.com/giantswarm/app-operator/releases/tag/v6.8.0)

#### Added
- Add Service Monitor by default to make it complain with the latest monitoring improvements 



### aws-operator [14.21.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.21.0)

#### Added
- Allow newer flatcar releases for node pools as provided by AWS release.
- Add sigs.k8s.io/cluster-api-provider-aws/role tag to all subnets as preparation for migration to CAPI.

#### Changed
- Unmanage interfaces for CNI eth[1-9] on workers eth[2-9] on masters
- [cilium eni mode] Only run aws-node, calico and kube-proxy on old nodes during migration to cilium.


### cert-operator [3.2.1](https://github.com/giantswarm/cert-operator/releases/tag/v3.2.1)

#### Fixed
- Fix rule names of PolicyException.



### cluster-operator [5.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.8.0)

#### Added
- Add ENI mode for Cilium on AWS.
- Consider new control-plane label.
#### Changed
- Propagate `global.podSecurityStandards.enforced` value set to `true` for PSS migration
- Rename function for better readbility.



### containerlinux [3510.2.6](https://www.flatcar-linux.org/releases/#release-3510.2.6)

 _Changes since **Stable 3510.2.5**_
 
 #### Security fixes:
 
 - Linux ([CVE-2022-48502](https://nvd.nist.gov/vuln/detail/CVE-2022-48502), [CVE-2023-20593](https://nvd.nist.gov/vuln/detail/CVE-2023-20593), [CVE-2023-2898](https://nvd.nist.gov/vuln/detail/CVE-2023-2898), [CVE-2023-31248](https://nvd.nist.gov/vuln/detail/CVE-2023-31248), [CVE-2023-35001](https://nvd.nist.gov/vuln/detail/CVE-2023-35001), [CVE-2023-3611](https://nvd.nist.gov/vuln/detail/CVE-2023-3611), [CVE-2023-3776](https://nvd.nist.gov/vuln/detail/CVE-2023-3776), [CVE-2023-38432](https://nvd.nist.gov/vuln/detail/CVE-2023-38432), [CVE-2023-3863](https://nvd.nist.gov/vuln/detail/CVE-2023-3863))
 - linux-firmware ([CVE-2023-20593](https://nvd.nist.gov/vuln/detail/CVE-2023-20593))
 
 #### Updates:
 
 - Linux ([5.15.122](https://lwn.net/Articles/939104) (includes [5.15.121](https://lwn.net/Articles/939016), [5.15.120](https://lwn.net/Articles/937404)))
 - ca-certificates ([3.92](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_92.html))
 - linux-firmware ([20230625](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230625))



### etcd [3.5.9](https://github.com/etcd-io/etcd/releases/tag/v3.5.9)

#### etcd server
- Fix [LeaseTimeToLive API may return keys to clients which have no read permission on the keys](https://github.com/etcd-io/etcd/pull/15815).
#### Dependencies
- Compile binaries using [go 1.19.9](https://github.com/etcd-io/etcd/pull/15822).



### kubernetes [1.24.15](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.15)

#### Feature
- Kubernetes 1.24.x is now built with Go 1.19.10 ([#118557](https://github.com/kubernetes/kubernetes/pull/118557), [@puerco](https://github.com/puerco)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Release, Storage and Testing]
#### Bug or Regression
- Fixes a bug at kube-apiserver start where APIService objects for custom resources could be deleted and recreated. ([#118104](https://github.com/kubernetes/kubernetes/pull/118104), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- If `kubeadm reset` finds no etcd member ID for the peer it removes during the `remove-etcd-member` phase, it continues immediately to other phases, instead of retrying the phase for up to 3 minutes before continuing. ([#118192](https://github.com/kubernetes/kubernetes/pull/118192), [@dlipovetsky](https://github.com/dlipovetsky)) [SIG Cluster Lifecycle]
- Kubeadm: fix a bug where the static pod changes detection logic is inconsistent with kubelet ([#118069](https://github.com/kubernetes/kubernetes/pull/118069), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
#### Dependencies
#### Added
- github.com/a8m/tree: [10a5fd5](https://github.com/a8m/tree/tree/10a5fd5)
- github.com/dougm/pretty: [2ee9d74](https://github.com/dougm/pretty/tree/2ee9d74)
- github.com/rasky/go-xdr: [4930550](https://github.com/rasky/go-xdr/tree/4930550)
- github.com/vmware/vmw-guestinfo: [25eff15](https://github.com/vmware/vmw-guestinfo/tree/25eff15)
#### Changed
- github.com/google/uuid: [v1.1.2 → v1.3.0](https://github.com/google/uuid/compare/v1.1.2...v1.3.0)
- github.com/kr/pretty: [v0.2.1 → v0.3.0](https://github.com/kr/pretty/compare/v0.2.1...v0.3.0)
- github.com/rogpeppe/go-internal: [v1.3.0 → v1.6.1](https://github.com/rogpeppe/go-internal/compare/v1.3.0...v1.6.1)
- github.com/vmware/govmomi: [v0.20.3 → v0.30.0](https://github.com/vmware/govmomi/compare/v0.20.3...v0.30.0)
#### Removed
_Nothing has changed._



### cert-exporter [2.6.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.6.0)

#### Changed
- Remove the Exist toleration from deployment. This allows the pod to be rescheduled on a drained node sometimes causing the drain of a node to fail and require a manual fix



### coredns [1.18.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.18.1)

#### Fixed
- Remove `fallthrough` for reverse zones from kubernetes plugin.



### external-dns [2.39.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.39.0)

#### Changed
- Replace monitoring labels with ServiceMonitor ([#296](https://github.com/giantswarm/external-dns-app/pull/296)).
- Update ATS to 0.4.1 and python deps ([#297](https://github.com/giantswarm/external-dns-app/pull/297)).



### vertical-pod-autoscaler [4.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v4.0.0)

#### Changed
WARNING: this version requires Cilium to run because of the dependency on the CiliumNetworkPolicy CRD
- Upgrade dependency chart to 9.2.0.
- Adjusted the resource and limits to accomodate larger clusters by default
- Adjusted the admission controller to give it more QPS against the API
- Adjusted the updater to give it more QPS against the API
- Adjusted the recommender to give it
  - more QPS against the API
  - doubling the memory in case of an OOMKilled event
  - Using the 95% percentile for the calculation of the CPU usage: should allow to scale up more precisely to account for spikes in CPU consumption of the workload
  - Adjusted the resource and limits to accomodate larger clusters by default
  - Calculating recommendations only for workloads which do have a VPA custom resource, instead of all workloads
  - Removed standard network policies to decrease maintenance burden
  - Fixed Cilium Network Policy to allow CRD jobs execution
  - Added Cilium Network Policy weight for an early execution
  - Disabled VPA for the updater pod otherwise it keeps on getting re-scheduled because the memory consumption varies a lot between reconsiling resources and idle
  - Disabled VPA for the recommender pod otherwise it keeps on getting re-scheduled because the memory consumption varies a lot between reconsiling resources and idle



### observability-bundle [0.7.3](https://github.com/giantswarm/observability-bundle/releases/tag/v0.7.3)

#### Changed
- Add extra config for `prometheus-operator-app` to be able to enable cilium.
#### Changed
- Upgrade `prometheus-operator-app` and `prometheus-operator-crd` to 5.1.0.



### cilium-servicemonitors [0.1.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.2)

#### Changed
- Drop metrics with high cardinality.
- Increase scrape interval to 60s.



### aws-ebs-csi-driver [2.25.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.25.0)

#### Changed
- Updated ebs-csi-driver to `v1.21.0` and updated sidecar images.



### cilium [0.11.1](https://github.com/giantswarm/cilium-app/releases/tag/v0.11.1)

#### Changed
- Create custom CNI config depending on provider to allow bigger customization.
- Bump all manifests to upstream version 1.13.6.



### cluster-autoscaler [1.24.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3)

#### Changed
- Change ScaleDownUtilizationThreshold default from 0.5 to 0.7
- Update cluster-autoscaler to version `1.24.3`.



### net-exporter [1.17.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.17.0)

#### Changed
- Add security context values to make chart comply to PodSecurityStandard restricted profile.



### node-exporter [1.16.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.1)

#### Changed
- Enable service monitor.



### etcd-kubernetes-resources-count-exporter [1.4.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.4.0)

#### Changed
- Add Max memory (default 500Mi) for VPA.



