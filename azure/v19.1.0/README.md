# :zap: Giant Swarm Release v19.1.0 for Azure :zap:

This is a maintainance and security release, featuring latest Flatcar Container Linux and Kubernetes versions. 

This is the release preparing for the migration away from Pod Security Policies (PSP) in favor of Pod Security Standards (PSS) in Kubernetes 1.25. Our `security-bundle` is now installed by default, and will deploy `kyverno` and `restricted` level PSS policies in `audit` mode. These resources are provided in order to allow time to create exceptions for workloads which need them before the policies are changed to `enforce` in a future release. For more information about PSS please read our official [documentation](https://docs.giantswarm.io/advanced/security-policy-enforcement/). Please also take a look at the `kyverno` [documentation](https://docs.giantswarm.io/platform-overview/security/platform-security/#kyverno) to utilize fully its potential.

> **WARNING:** If you are already running `kyverno` as Giant Swarm Managed App, the installation of `security-bundle` will fail. However the already existing `kyverno` deployment and its configuration can be adopted by the bundle after the upgrade is finished. Please talk to your Account Engineer if you have any questions.

## Change details


### app-operator [6.7.0](https://github.com/giantswarm/app-operator/releases/tag/v6.7.0)

#### Changed
- Only include PodSecurityPolicy on clusters with policy/v1beta1 api available.
- Only include PodMonitor on clusters with monitoring.coreos.com/v1 api available.
#### Removed
- Stop pushing to `openstack-app-collection`.



### azure-operator [7.3.0](https://github.com/giantswarm/azure-operator/releases/tag/v7.3.0)

#### Changed
- Bump k8scc to 16.2.0
- Remove logic that migrates CAPI CRDs from experimental group to new group.



### cluster-operator [5.6.1](https://github.com/giantswarm/cluster-operator/releases/tag/v5.6.1)

#### Fixed
- Don't enable Cilium network policies on Azure.



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



### containerlinux [3510.2.2](https://www.flatcar-linux.org/releases/#release-3510.2.2)

 _Changes since **Stable 3510.2.1**_
 
#### Security fixes:
 
 - Linux ([CVE-2023-1380](https://nvd.nist.gov/vuln/detail/CVE-2023-1380), [CVE-2023-1859](https://nvd.nist.gov/vuln/detail/CVE-2023-1859), [CVE-2023-2002](https://nvd.nist.gov/vuln/detail/CVE-2023-2002), [CVE-2023-2269](https://nvd.nist.gov/vuln/detail/CVE-2023-2269), [CVE-2023-31436](https://nvd.nist.gov/vuln/detail/CVE-2023-31436), [CVE-2023-32233](https://nvd.nist.gov/vuln/detail/CVE-2023-32233))
 
#### Bug fixes:
 
 
#### Changes:
 
 
#### Updates:
 
 - Linux ([5.15.111](https://lwn.net/Articles/931652) (includes [5.15.110](https://lwn.net/Articles/930600), [5.15.109](https://lwn.net/Articles/930263), [5.15.108](https://lwn.net/Articles/929679), [5.15.107](https://lwn.net/Articles/929015)))
 - ca-certificates ([3.89.1](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_89_1.html))



### etcd [3.5.9](https://github.com/etcd-io/etcd/releases/tag/v3.5.9)

#### etcd server
- Fix [LeaseTimeToLive API may return keys to clients which have no read permission on the keys](https://github.com/etcd-io/etcd/pull/15815).
#### Dependencies
- Compile binaries using [go 1.19.9](https://github.com/etcd-io/etcd/pull/15822).



### azure-cloud-controller-manager [1.24.18-gs4](https://github.com/giantswarm/azure-cloud-controller-manager-app/releases/tag/v1.24.18-gs4)

#### Changed
- Use `node-role.kubernetes.io/control-plane: ""` as master node selector.



### azure-cloud-node-manager [1.24.18-gs3](https://github.com/giantswarm/azure-cloud-node-manager-app/releases/tag/v1.24.18-gs3)

#### Changed
- Remove `capabitlities.apiversion.has` check for VPA to avoid race condition between this app being installed and the api-version providing app being installed
  - With this change the installation of the chart will fail until the `api-version` is available



### azuredisk-csi-driver [1.26.2-gs4](https://github.com/giantswarm/azuredisk-csi-driver-app/releases/tag/v1.26.2-gs4)

#### Changed
- Remove `capabitlities.apiversion.has` check for VPA to avoid race condition between this app being installed and the api-version providing app being installed
  - With this change the installation of the chart will fail until the `api-version` is available



### azurefile-csi-driver [1.26.0-gs2](https://github.com/giantswarm/azurefile-csi-driver-app/releases/tag/v1.26.0-gs2)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### cert-exporter [2.6.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.6.0)




### chart-operator [2.35.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.35.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### coredns [1.18.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.18.0)

#### Added
- Add scaling based on custom metrics ([#209](https://github.com/giantswarm/coredns-app/pull/209)).
- Add a new field additionalLocalZones which can be used to introduce more internal local zones, e.g. linkerd.
#### Changed
- Decouple PDB configuration from deployment updateStrategy ([#208](https://github.com/giantswarm/coredns-app/pull/208)).
- Disable IPV6.
- Create a coredns zone for each cluster domain.
- Adjust the settings for upscaling HPA when hitting 60% CPU.
- Adjust the settings for downscaling HPA to 30 minutes.
- Adjust the min and max memory settings per Pod.
- Enable cache inconditionaly for . and local zones.
- Adjust the settings for upscaling HPA when hitting 80% Memory.



### etcd-kubernetes-resources-count-exporter [1.2.1](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.2.1)

#### Removed
- Removed debug code that was dumping all events' contents to stdout.



### external-dns [2.37.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.37.1)

#### Changed
- Remove deprecated annotation from Pod.[#265](https://github.com/giantswarm/external-dns-app/pull/265).
#### Fixed
- Fix indentation in environment variables for secret injection [#272](https://github.com/giantswarm/external-dns-app/pull/272).



### metrics-server [2.2.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.
- Switch to `apiVersion: policy/v1` for PodDisruptionBudget.



### net-exporter [1.16.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.16.2)

#### Changed
- Reduce CPU and Mem requests.



### node-exporter [1.16.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.1)

#### Changed
- Enable service monitor.



### cluster-autoscaler [1.24.0-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.0-gs2)

#### Changed
- Add 'projected' volumes to the PSP.
- Add new-pod-scale-up-delay variable.
- Disable PSPs for k8s 1.25 and newer.



### azure-scheduled-events [0.8.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.8.0)

#### Changed
- Bumped dependencies.
- Switched to go 1.18. 



### vertical-pod-autoscaler [3.5.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.2)

#### Changed
- Raised limits for all components.



### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap



### observability-bundle [0.7.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.7.0)

#### Changed
- Upgrade `prometheus-operator-app` to 5.0.5.
- Upgrade `prometheus-operator-crd` to 5.0.0.



### security-bundle [0.16.0](https://github.com/giantswarm/security-bundle/releases/tag/v0.16.0)

#### Changed
- Update to `kyverno` (app) version 0.14.7, introducing exception mechanisms for `chart-operator` and restricting wildcards for Kinds.
- Disabled the default apps `falco`, `trivy`, `trivy-operator` and `starboard-exporter`. This apps can be manually enabled.



