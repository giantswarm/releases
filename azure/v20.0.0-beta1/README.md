# :zap: Giant Swarm Release v20.0.0-beta1 for Azure :zap:

This is Giant Swarm release preparing towards Kubernetes 1.25 with migration PSPs. This is unsecure release by default disabling all PSPs from all components. Please do not upgrade to that release directly and talk to your Account Engineer for more details.

## Change details


### app-operator [6.7.0](https://github.com/giantswarm/app-operator/releases/tag/v6.7.0)

#### Changed
- Only include PodSecurityPolicy on clusters with policy/v1beta1 api available.
- Only include PodMonitor on clusters with monitoring.coreos.com/v1 api available.
#### Removed
- Stop pushing to `openstack-app-collection`.



### azure-operator [8.0.0](https://github.com/giantswarm/azure-operator/releases/tag/v8.0.0)

#### Added
- Add new control-plane label to detect master nodes.
#### Removed
- Remove logic that migrates CAPI CRDs from experimental group to new group.
#### Changed
- Kubernetes 1.25 support.



### cluster-operator [5.6.1](https://github.com/giantswarm/cluster-operator/releases/tag/v5.6.1)

#### Fixed
- Don't enable Cilium network policies on Azure.



### kubernetes [1.25.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.25.9)

#### Feature
- Kubernetes is now built with Go 1.19.8 ([#117131](https://github.com/kubernetes/kubernetes/pull/117131), [@xmudrii](https://github.com/xmudrii)) [SIG Release and Testing]
#### Bug or Regression
- Fix missing delete events on informer re-lists to ensure all delete events are correctly emitted and using the latest known object state, so that all event handlers and stores always reflect the actual apiserver state as best as possible ([#115900](https://github.com/kubernetes/kubernetes/pull/115900), [@odinuge](https://github.com/odinuge)) [SIG API Machinery]
- Fix: Route controller should update routes with NodeIP changed ([#116361](https://github.com/kubernetes/kubernetes/pull/116361), [@lzhecheng](https://github.com/lzhecheng)) [SIG Cloud Provider]
- Fixes a regression in the pod binding subresource to honor the `metadata.uid` precondition.
  This allows kube-scheduler to ensure it is assigns node names to the same instances of pods it made scheduling decisions for. ([#116776](https://github.com/kubernetes/kubernetes/pull/116776), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery and Testing]
- Kubelet: Fix fs quota monitoring on volumes ([#116794](https://github.com/kubernetes/kubernetes/pull/116794), [@pacoxu](https://github.com/pacoxu)) [SIG Storage]
#### Other (Cleanup or Flake)
- Service session affinity timeout tests are no longer required for Kubernetes network plugin conformance due to variations in existing implementations. New conformance tests will be developed to better express conformance in future releases. ([#112806](https://github.com/kubernetes/kubernetes/pull/112806), [@dcbw](https://github.com/dcbw)) [SIG Architecture, Network and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.35 → v0.0.36
#### Removed
_Nothing has changed._



### containerlinux [3510.2.1](https://www.flatcar-linux.org/releases/#release-3510.2.1)

_Changes since **Stable 3510.2.0**_

#### Security fixes:

- Linux ([CVE-2022-4269](https://nvd.nist.gov/vuln/detail/CVE-2022-4269), [CVE-2022-4379](https://nvd.nist.gov/vuln/detail/CVE-2022-4379), [CVE-2023-1076](https://nvd.nist.gov/vuln/detail/CVE-2023-1076), [CVE-2023-1077](https://nvd.nist.gov/vuln/detail/CVE-2023-1077), [CVE-2023-1079](https://nvd.nist.gov/vuln/detail/CVE-2023-1079), [CVE-2023-1118](https://nvd.nist.gov/vuln/detail/CVE-2023-1118), [CVE-2023-1611](https://nvd.nist.gov/vuln/detail/CVE-2023-1611), [CVE-2023-1670](https://nvd.nist.gov/vuln/detail/CVE-2023-1670), [CVE-2023-1829](https://nvd.nist.gov/vuln/detail/CVE-2023-1829), [CVE-2023-1855](https://nvd.nist.gov/vuln/detail/CVE-2023-1855), [CVE-2023-1989](https://nvd.nist.gov/vuln/detail/CVE-2023-1989), [CVE-2023-1990](https://nvd.nist.gov/vuln/detail/CVE-2023-1990), [CVE-2023-23004](https://nvd.nist.gov/vuln/detail/CVE-2023-23004), [CVE-2023-25012](https://nvd.nist.gov/vuln/detail/CVE-2023-25012), [CVE-2023-28466](https://nvd.nist.gov/vuln/detail/CVE-2023-28466), [CVE-2023-30456](https://nvd.nist.gov/vuln/detail/CVE-2023-30456), [CVE-2023-30772](https://nvd.nist.gov/vuln/detail/CVE-2023-30772))
- nvidia-drivers ([CVE-2022-31607](https://nvd.nist.gov/vuln/detail/CVE-2022-31607), [CVE-2022-31608](https://nvd.nist.gov/vuln/detail/CVE-2022-31608), [CVE-2022-31615](https://nvd.nist.gov/vuln/detail/CVE-2022-31615), [CVE-2022-34665](https://nvd.nist.gov/vuln/detail/CVE-2022-34665), [CVE-2022-34666](https://nvd.nist.gov/vuln/detail/CVE-2022-34666), [CVE-2022-34670](https://nvd.nist.gov/vuln/detail/CVE-2022-34670), [CVE-2022-34673](https://nvd.nist.gov/vuln/detail/CVE-2022-34673), [CVE-2022-34674](https://nvd.nist.gov/vuln/detail/CVE-2022-34674), [CVE-2022-34676](https://nvd.nist.gov/vuln/detail/CVE-2022-34676), [CVE-2022-34677](https://nvd.nist.gov/vuln/detail/CVE-2022-34677), [CVE-2022-34678](https://nvd.nist.gov/vuln/detail/CVE-2022-34678), [CVE-2022-34679](https://nvd.nist.gov/vuln/detail/CVE-2022-34679), [CVE-2022-34680](https://nvd.nist.gov/vuln/detail/CVE-2022-34680), [CVE-2022-34682](https://nvd.nist.gov/vuln/detail/CVE-2022-34682), [CVE-2022-34684](https://nvd.nist.gov/vuln/detail/CVE-2022-34684), [CVE-2022-42254](https://nvd.nist.gov/vuln/detail/CVE-2022-42254), [CVE-2022-42255](https://nvd.nist.gov/vuln/detail/CVE-2022-42255), [CVE-2022-42256](https://nvd.nist.gov/vuln/detail/CVE-2022-42256), [CVE-2022-42257](https://nvd.nist.gov/vuln/detail/CVE-2022-42257), [CVE-2022-42258](https://nvd.nist.gov/vuln/detail/CVE-2022-42258), [CVE-2022-42259](https://nvd.nist.gov/vuln/detail/CVE-2022-42259), [CVE-2022-42260](https://nvd.nist.gov/vuln/detail/CVE-2022-42260), [CVE-2022-42261](https://nvd.nist.gov/vuln/detail/CVE-2022-42261), [CVE-2022-42263](https://nvd.nist.gov/vuln/detail/CVE-2022-42263), [CVE-2022-42264](https://nvd.nist.gov/vuln/detail/CVE-2022-42264), [CVE-2022-42265](https://nvd.nist.gov/vuln/detail/CVE-2022-42265))

#### Bug fixes:
- Fixed the broken emerge-gitclone in the dev-container owing to the missing migration action around the unification of the Flatcar core repositories

#### Changes:
- The package upgrade for nvidia-drivers might result in not supporting a few of the older NVIDIA Tesla GPUs. If you are facing issues, set `NVIDIA_DRIVER_VERSION=460.106.00` in `/etc/flatcar/nvidia-metadata`

#### Updates:

- Linux ([5.15.106](https://lwn.net/Articles/928343) (includes [5.15.105](https://lwn.net/Articles/927860), [5.15.104](https://lwn.net/Articles/926873), [5.15.103](https://lwn.net/Articles/926415), [5.15.102](https://lwn.net/Articles/925991), [5.15.101](https://lwn.net/Articles/925939), [5.15.100](https://lwn.net/Articles/925913), [5.15.99](https://lwn.net/Articles/925844)))
- nvidia-drivers ([525.105.17](https://docs.nvidia.com/datacenter/tesla/tesla-release-notes-525-105-17/index.html))



### calico [3.25.1](https://github.com/projectcalico/calico/releases/tag/v3.25.1)

Not found


### azure-cloud-controller-manager [1.24.18-gs3](https://github.com/giantswarm/azure-cloud-controller-manager-app/releases/tag/v1.24.18-gs3)

#### Changed
- Remove `capabitlities.apiversion.has` check for VPA to avoid race condition between this app being installed and the api-version providing app being installed
  - With this change the installation of the chart will fail until the `api-version` is available



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



### cert-exporter [2.5.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.5.1)




### chart-operator [2.35.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.35.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### coredns [1.17.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.17.0)

#### Added
- Add scaling based on custom metrics ([#209](https://github.com/giantswarm/coredns-app/pull/209)).
#### Changed
- Decouple PDB configuration from deployment updateStrategy ([#208](https://github.com/giantswarm/coredns-app/pull/208)).



### etcd-kubernetes-resources-count-exporter [1.2.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### external-dns [2.37.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.37.1)

Not found


### metrics-server [2.2.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.2.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.
- Switch to `apiVersion: policy/v1` for PodDisruptionBudget.



### net-exporter [1.15.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.15.0)

#### Changed
- Allow requests from the api-server.
- Disable PSPs for k8s 1.25 and newer.



### node-exporter [1.16.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.0)

#### Changed
- Disable PSPs for k8s 1.25 and newer.



### vertical-pod-autoscaler [3.4.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.2)

#### Changed
- Remove circleci job for pushing to shared app collection



### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap



### observability-bundle [0.5.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.5.1)

#### Changed
- Remove cluster prefix to app name in _helpers.tpl



### k8s-dns-node-cache [2.1.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.1.0)

#### Changed
- Switch to ServiceMonitors for metrics scraping.



