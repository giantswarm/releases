# :zap: Giant Swarm Release v19.1.0 for AWS :zap:

This is a maintenance release featuring latest 1.24 Kubernetes versions as well as components upgrades. This release also introduces new features which are described in next sections.

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.3.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#330---2023-05-11) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Cilium AWS ENI mode

Following our work on changing the CNI of the Giant Swarm Workload Clusters towards Cilium, we have added a possibility to migrate to the Cilium AWS ENI mode instead of plain Cilium setup.

> **WARNING:** The Cilium AWS ENI mode can be *ONLY* enabled while upgrading from `18.4.2` to `19.1.0` release. From that point forward the Workload Clusters will be running in Cilium AWS ENI mode and cannot be switched back to our default Cilium that comes with `19.0.0`. Both the Cilium and Cilium AWS ENI mode will receive the same level of support going forward.

Feature can be enabled via an annotation `cilium.giantswarm.io/ipam-mode: eni` set on the Cluster CR while on `18.4.2` release prior to `19.1.0` upgrade. When the upgrade is triggered, the underlying infrastructure will choose to continue with the [Cilium AWS ENI mode](https://docs.cilium.io/en/latest/network/concepts/ipam/eni/). This is meant for the users that do not want to migrate any of the underlying network infrastructure that has been linked with the Giant Swarm Workload Clusters. The network setup after the upgrade will be the same as while running `aws-cni` with kube-proxy.

## Kyverno by default

This is the release preparing for the migration away from Pod Security Policies (PSP) in favor of Pod Security Standards (PSS) in Kubernetes 1.25. Our `security-bundle` is now installed by default, and will deploy `kyverno` and `restricted` level PSS policies in `audit` mode. These resources are provided in order to allow time to create exceptions for workloads which need them before the policies are changed to `enforce` in a future release. For more information about PSS please read our official [documentation](https://docs.giantswarm.io/advanced/security-policy-enforcement/). Please also take a look at the `kyverno` [documentation](https://docs.giantswarm.io/platform-overview/security/platform-security/#kyverno) to utilize fully its potential.

> **WARNING:** If you are already running `kyverno` as Giant Swarm Managed App, the installation of `security-bundle` will fail. However the already existing `kyverno` deployment and its configuration can be adopted by the bundle after the upgrade is finished. Please talk to your Account Engineer if you have any questions.

## AWSMachineDeployment CR's annotation to change the Flatcar Release Version

This feature allows customers to set an annotation on AWSMachineDeployment CR's to change the Flatcar Release Version. For now it only allows setting `alpha.giantswarm.io/flatcar-release-version: "3689.0.0"` or a higher version. We have added this feature to accommodate the issues with Cilium CNI high CPU usage on small clusters. This feature is solely to enable customers to run the Flatcar `alpha` channel which consists of the `kernel 6` version that fixes the issue, while waiting for a `stable` Flatcar release.

Please read the detailed description of the designed behaviour of the annotation:
- when setting the annotation, the TCNP CloudFormation Stack for the specific node pool is rolled and replaces the OS image
- when removing the annotation the node pool is updated and switches back to the default OS image which is coming from the AWS release
- when upgrading the cluster to a new AWS release, the node pool uses the specific flatcar release from the annotation as long as you don't change by either setting it to a higher version or removing the annotation.


## Change details


### app-operator [6.8.0](https://github.com/giantswarm/app-operator/releases/tag/v6.8.0)

#### Added

- Add Service Monitor by default to make it compliant with the latest monitoring improvements 



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
- Rename function for better readability.


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


### security-bundle [0.16.2](https://github.com/giantswarm/security-bundle/releases/tag/v0.16.2)

#### Changed
- Update to `kyverno` (app) version 0.14.7, introducing exception mechanisms for `chart-operator` and restricting wildcards for Kinds.
- Disabled the default apps `falco`, `trivy`, `trivy-operator` and `starboard-exporter`. This apps can be manually enabled.


