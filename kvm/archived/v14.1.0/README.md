# :zap: Giant Swarm Release v14.1.0 for KVM :zap:

This release introduces a new feature that allows KVM clusters to run behind a proxy. It also adds support for host volumes.

## Change details


### cluster-operator [0.27.1](https://github.com/giantswarm/cluster-operator/releases/tag/v0.27.1)

#### Changed
- Dropped ensuring cluster CRDs from controllers.



### app-operator [4.4.0](https://github.com/giantswarm/app-operator/releases/tag/v4.4.0)

#### Added
- Add support for skip CRD flag when installing Helm releases.
- Emit events when config maps and secrets referenced in App CRs are updated.



### kvm-operator [3.17.2](https://github.com/giantswarm/kvm-operator/releases/tag/v3.17.2)

#### Fixed

- Remove reference from worker PVs on cluster deletion so they can be resued.

### Added

- Add flags for proxy settings and propagate them to ignition.

#### Changed

- Reconcile only deployments that are managed by kvm-operator.


### containerlinux [2765.2.4](https://kinvolk.io/flatcar-container-linux/releases/#release-2765.2.4)

**Security fixes**

*   Linux ([CVE-2021-3491](https://nvd.nist.gov/vuln/detail/CVE-2021-3491), [CVE-2021-31440](https://nvd.nist.gov/vuln/detail/CVE-2021-31440), [CVE-2021-31829](https://nvd.nist.gov/vuln/detail/CVE-2021-31829))
*   nvidia-drivers ([CVE-2021-1052](https://nvd.nist.gov/vuln/detail/CVE-2021-1052), [CVE-2021-1053](https://nvd.nist.gov/vuln/detail/CVE-2021-1053), [CVE-2021-1056](https://nvd.nist.gov/vuln/detail/CVE-2021-1056), [CVE-2021-1076](https://nvd.nist.gov/vuln/detail/CVE-2021-1076), [CVE-2021-1077](https://nvd.nist.gov/vuln/detail/CVE-2021-1077))
*   runc ([CVE-2021-30465](https://nvd.nist.gov/vuln/detail/CVE-2021-30465))

**Updates**

*   Linux ([5.10.37](https://lwn.net/Articles/856269/))
*   nvidia-drivers ([460.73.01](https://www.nvidia.com/Download/driverResults.aspx/172376/en-us))


### etcd [3.4.16](https://github.com/etcd-io/etcd/releases/tag/v3.4.16)

- Add [`--experimental-warning-apply-duration`](https://github.com/etcd-io/etcd/pull/12448) flag which allows apply duration threshold to be configurable.
- Fix [`--unsafe-no-fsync`](https://github.com/etcd-io/etcd/pull/12751) to still write-out data avoiding corruption (most of the time).
- Reduce [around 30% memory allocation by logging range response size without marshal](https://github.com/etcd-io/etcd/pull/12871).
- Add [exclude alarms from health check conditionally](https://github.com/etcd-io/etcd/pull/12880).


### kubernetes [1.19.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.11)

#### API Change
- We have added a new Priority & Fairness rule that exempts all probes (/readyz, /healthz, /livez) to prevent 
  restarting of "healthy" kube-apiserver instance(s) by kubelet. ([#101113](https://github.com/kubernetes/kubernetes/pull/101113), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
#### Feature
- Kubernetes is now built using go1.15.11 ([#101197](https://github.com/kubernetes/kubernetes/pull/101197), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built using go1.15.12 ([#101846](https://github.com/kubernetes/kubernetes/pull/101846), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Azurefile: Normalize share name to not include capital letters ([#100731](https://github.com/kubernetes/kubernetes/pull/100731), [@kassarl](https://github.com/kassarl)) [SIG Cloud Provider and Storage]
- EndpointSlice IP validation now matches Endpoints IP validation. ([#101084](https://github.com/kubernetes/kubernetes/pull/101084), [@robscott](https://github.com/robscott)) [SIG Apps and Network]
- EndpointSlice controllers are less likely to create duplicate EndpointSlices. ([#101764](https://github.com/kubernetes/kubernetes/pull/101764), [@aojea](https://github.com/aojea)) [SIG Apps and Network]
- Ensure service deleted when the Azure resource group has been deleted ([#100944](https://github.com/kubernetes/kubernetes/pull/100944), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix panic in JSON logging format caused by missing Duration encoder ([#101159](https://github.com/kubernetes/kubernetes/pull/101159), [@serathius](https://github.com/serathius)) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fix: azure file inline volume namespace issue in csi migration translation ([#101235](https://github.com/kubernetes/kubernetes/pull/101235), [@andyzhangx](https://github.com/andyzhangx)) [SIG Apps, Cloud Provider, Node and Storage]
- Fixed a bug where startupProbe stopped working after a container's first restart ([#101093](https://github.com/kubernetes/kubernetes/pull/101093), [@wzshiming](https://github.com/wzshiming)) [SIG Node]
- Fixed port-forward memory leak for long-running and heavily used connections. ([#99839](https://github.com/kubernetes/kubernetes/pull/99839), [@saschagrunert](https://github.com/saschagrunert)) [SIG API Machinery and Node]
- Kubelet: improve the performance when waiting for a synchronization of the node list with the kube-apiserver ([#99336](https://github.com/kubernetes/kubernetes/pull/99336), [@neolit123](https://github.com/neolit123)) [SIG Node]
- No support endpointslice in linux userpace mode ([#101502](https://github.com/kubernetes/kubernetes/pull/101502), [@JornShen](https://github.com/JornShen)) [SIG Network]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### cert-exporter [1.6.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.6.1)

#### Changed
- Set docker.io as the default registry



### chart-operator [2.15.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.15.0)

#### Added
- Proxy support in helm template.



### kube-state-metrics [1.3.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.1)

#### Changed
- Set docker.io as the default registry



### metrics-server [1.3.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.3.0)

#### Added
- Added new configuration value `extraArgs`.



### node-exporter [1.7.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.2)

#### Changed
- Set docker.io as the default registry



