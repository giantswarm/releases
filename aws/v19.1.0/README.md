# :zap: Giant Swarm Release v19.1.0 for AWS :zap:

This is a maintenance release featuring latest 1.24 Kubernetes versions as well as components upgrades. This release also introduces new features which are described in next sections.

***IAM Permissions Requirements***
The minimal requirement for the IAM permissions is [Version 3.3.0](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/blob/master/CHANGELOG.md#330---2023-05-11) of [giantswarm-aws-account-prerequisites](https://github.com/giantswarm/giantswarm-aws-account-prerequisites/) repository.

## Cilium AWS ENI mode

Following our work on changing the CNI of the Giant Swarm Workload Clusters towards Cilium, we have added a possibility to migrate to the Cilium AWS ENI mode instead of plain Cilium setup.

> **WARNING:** The Cilium AWS ENI mode can *ONLY* be enabled while upgrading from `18.4.2` to `19.1.0` release. From that point forward the Workload Clusters will be running in Cilium AWS ENI mode and cannot be switched back to our default Cilium that comes with `19.0.0`. Both the Cilium and Cilium AWS ENI mode will receive the same level of support going forward.

This feature can be enabled via the annotation `cilium.giantswarm.io/ipam-mode: eni`, set on the Cluster CR, while on `18.4.2` release and prior to `19.1.0` upgrade. When the upgrade is triggered, the underlying infrastructure will choose to continue with the [Cilium AWS ENI mode](https://docs.cilium.io/en/latest/network/concepts/ipam/eni/). This is meant for the users that do not want to migrate any of the underlying network infrastructure that has been linked with the Giant Swarm Workload Clusters. The network setup after the upgrade will be the same as while running `aws-cni` with `kube-proxy`.

## Kyverno by default

This release prepares for the migration away from Pod Security Policies (PSP) in favor of Pod Security Standards (PSS) in Kubernetes 1.25. Our `security-bundle` is now installed by default, and will deploy `kyverno` and `restricted` level PSS policies in `audit` mode. These resources are provided in order to allow time to create exceptions for workloads which need them before the policies are changed to `enforce` in a future release. For more information about PSS please read our official [documentation](https://docs.giantswarm.io/advanced/security-policy-enforcement/). Please also take a look at the `kyverno` [documentation](https://docs.giantswarm.io/platform-overview/security/platform-security/#kyverno) to fully utilize its potential.

> **WARNING:** If you are already running `kyverno` as Giant Swarm Managed App, the installation of `security-bundle` will fail. However the already existing `kyverno` deployment and its configuration can be adopted by the bundle after the upgrade is finished. Please talk to your Account Engineer if you have any questions.

## AWSMachineDeployment CR's annotation to change the Flatcar Release Version

This feature allows customers to set an annotation on AWSMachineDeployment CR's to change the Flatcar Release Version. For now it only allows setting `alpha.giantswarm.io/flatcar-release-version: "3689.0.0"` or higher version. We have added this feature to accommodate the issues with Cilium CNI high CPU usage on small clusters. This feature is solely to enable customers to run the Flatcar `alpha` channel which consists of the `kernel 6` version that fixes the issue, while waiting for a `stable` Flatcar release.

The annotation behaves as follows:
- when setting the annotation, the TCNP CloudFormation Stack for the specific node pool is rolled and replaces the OS image
- when removing the annotation, the node pool is updated and switches back to the default OS image which is coming from the AWS release
- when upgrading the cluster to a new AWS release, the node pool uses the specific flatcar release from the annotation as long as you don't change by either setting it to a higher version or removing the annotation.

## Change details


### etcd [3.5.9](https://github.com/etcd-io/etcd/releases/tag/v3.5.9)

#### etcd server
- Fix [LeaseTimeToLive API may return keys to clients which have no read permission on the keys](https://github.com/etcd-io/etcd/pull/15815).
#### Dependencies
- Compile binaries using [go 1.19.9](https://github.com/etcd-io/etcd/pull/15822).



### kubernetes [1.24.17](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.17)

#### Feature
- Kubernetes is now built with Go 1.20.7 ([#119837](https://github.com/kubernetes/kubernetes/pull/119837), [@jeremyrickard](https://github.com/jeremyrickard)) [SIG Apps, Cloud Provider, Node, Release, Storage and Testing]
#### Bug or Regression
- Fixed a bug where clusters that use KMS v1 with skewed API servers on versions v1.24 and v1.25 would see internal errors when attempting to read encrypted data via the v1.24 API servers. ([#119387](https://github.com/kubernetes/kubernetes/pull/119387), [@enj](https://github.com/enj)) [SIG API Machinery and Auth]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### app-operator [6.8.0](https://github.com/giantswarm/app-operator/releases/tag/v6.8.0)

#### Added
- Add Service Monitor by default to make it complain with the latest monitoring improvements 



### aws-operator [14.22.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.22.0)

#### Changed

- Get AMI data from helm value rather than from hardcoded string in the code.



⚠️ Attention: Major release [3.0.0](#300---2023-07-26) contains breaking changes in user values! Please make sure you [read about the upgrade instructions](https://github.com/giantswarm/cert-manager-app/blob/main/docs/upgrading.md)! ⚠️
#### Added
- Add NetworkPolicies for controller and cainjector. ([#354](https://github.com/giantswarm/cert-manager-app/pull/354))

#### Fixed
- Fix rule names of PolicyException.



### cluster-operator [5.8.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.8.0)

#### Added
- Add ENI mode for Cilium on AWS.
- Consider new control-plane label.
#### Changed
- Propagate `global.podSecurityStandards.enforced` value set to `true` for PSS migration
- Rename function for better readbility.



### containerlinux [3510.2.7](https://www.flatcar-linux.org/releases/#release-3510.2.7)

 _Changes since **Stable 3510.2.6**_
 
 #### Security fixes:
 
 - Linux ([CVE-2022-40982](https://nvd.nist.gov/vuln/detail/CVE-2022-40982), [CVE-2022-41804](https://nvd.nist.gov/vuln/detail/CVE-2022-41804), [CVE-2023-1206](https://nvd.nist.gov/vuln/detail/CVE-2023-1206), [CVE-2023-20569](https://nvd.nist.gov/vuln/detail/CVE-2023-20569), [CVE-2023-4004](https://nvd.nist.gov/vuln/detail/CVE-2023-4004), [CVE-2023-4147](https://nvd.nist.gov/vuln/detail/CVE-2023-4147), [CVE-2023-20569](https://nvd.nist.gov/vuln/detail/CVE-2023-20569), [CVE-2023-23908](https://nvd.nist.gov/vuln/detail/CVE-2023-23908))
 
 #### Bug fixes:
 
 - Fixed the restart of Systemd services when the main process is being killed by a SIGHUP signal ([flatcar#1157](https://github.com/flatcar/Flatcar/issues/1157))
 
 #### Updates:
 
 - Linux ([5.15.125](https://lwn.net/Articles/940801) (includes [5.15.124](https://lwn.net/Articles/940339), [5.15.123](https://lwn.net/Articles/939424)))


### net-exporter [1.17.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.17.0)

#### Changed
- Add security context values to make chart comply to PodSecurityStandard restricted profile.



### node-exporter [1.17.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.17.1)

#### Changed
- fix apparmor annotation



### etcd-kubernetes-resources-count-exporter [1.4.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.4.0)

#### Changed
- Add Max memory (default 500Mi) for VPA.



### cilium-servicemonitors [0.1.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.2)

#### Changed
- Drop metrics with high cardinality.
- Increase scrape interval to 60s.



### cert-exporter [2.6.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.6.0)




### cilium [0.12.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.12.0)

#### Added
- Support creating `CiliumNetworkPolicy` manifests that allow egress requests to DNS and proxy hosts
#### Changed
- Add missing conditional for PSP rendering of default-policies installer job



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



### observability-bundle [0.8.2](https://github.com/giantswarm/observability-bundle/releases/tag/v0.8.)

- Upgrade promtail to 1.4.0.



### cilium-servicemonitors [0.1.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.1.2)

#### Changed
- Drop metrics with high cardinality.
- Increase scrape interval to 60s.



### aws-ebs-csi-driver [2.25.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.25.0)

#### Changed
- Updated ebs-csi-driver to `v1.21.0` and updated sidecar images.



### cluster-autoscaler [1.24.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3)

#### Changed
- Change ScaleDownUtilizationThreshold default from 0.5 to 0.7
- Update cluster-autoscaler to version `1.24.3`.



### cilium [0.12.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.12.0)

#### Added
- Support creating `CiliumNetworkPolicy` manifests that allow egress requests to DNS and proxy hosts
#### Changed
- Add missing conditional for PSP rendering of default-policies installer job



### external-dns [2.39.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.39.0)

#### Changed
- Replace monitoring labels with ServiceMonitor ([#296](https://github.com/giantswarm/external-dns-app/pull/296)).
- Update ATS to 0.4.1 and python deps ([#297](https://github.com/giantswarm/external-dns-app/pull/297)).



### etcd-kubernetes-resources-count-exporter [1.4.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.4.0)

#### Changed
- Add Max memory (default 500Mi) for VPA.



### cert-exporter [2.6.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.6.0)

### Changed

- Remove the `Exist` toleration from deployment. This allows the pod to be rescheduled on a drained node sometimes causing the drain of a node to fail and require a manual fix



### cert-manager [3.3.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.3.0)

⚠️ Attention: Major release [3.0.0](#300---2023-07-26) contains breaking changes in user values! Please make yourself familiar with its changelog! ⚠️
#### Added
- Add NetworkPolicies for controller and cainjector. ([#354](https://github.com/giantswarm/cert-manager-app/pull/354))


### security-bundle [0.17.0](https://github.com/giantswarm/security-bundle/releases/tag/v0.17.0)

#### Added

- Update to kyverno (app) upstream version 1.10.2. Note: This update includes breaking changes in the values structure, please check the migration docs before upgrading.
- Update to trivy (app) version 0.8.3.
- Update to falco (app) version 0.6.5.



### k8s-audit-metrics [0.7.0](https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.7.0)

#### Changed

- Switched to kube-system namespace by default
- Added Cilium Network Policy to scrape /metrics on port 8000
