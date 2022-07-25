# :zap: Giant Swarm Release v18.0.0-alpha1 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [v13.0.0-alpha1](https://github.com/giantswarm/aws-operator/releases/tag/vv13.0.0-alpha1)

Not found


### kubernetes [1.23.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.9)

#### Bug or Regression
- Fix a bug that caused the wrong result length when using --chunk-size and --selector together ([#110757](https://github.com/kubernetes/kubernetes/pull/110757), [@Abirdcfly](https://github.com/Abirdcfly)) [SIG API Machinery and Testing]
- Fix bug that prevented the job controller from enforcing activeDeadlineSeconds when set ([#110545](https://github.com/kubernetes/kubernetes/pull/110545), [@harshanarayana](https://github.com/harshanarayana)) [SIG Apps]
- Fix image pulling failure when IMDS is unavailable in kubelet startup ([#110523](https://github.com/kubernetes/kubernetes/pull/110523), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix printing resources with int64 fields ([#110602](https://github.com/kubernetes/kubernetes/pull/110602), [@sanchezl](https://github.com/sanchezl)) [SIG API Machinery]
- Fixed a regression introduced in 1.23.0 where Azure load balancers were not kept up to date with the state of cluster nodes. In particular, nodes that are not in the ready state and are not newly created (i.e. not having the `node.cloudprovider.kubernetes.io/uninitialized` taint) now get removed from Azure load balancers. ([#109932](https://github.com/kubernetes/kubernetes/pull/109932), [@ricky-rav](https://github.com/ricky-rav)) [SIG Cloud Provider]
- Fixed potential scheduler crash when scheduling with unsatisfied nodes in PodTopologySpread. ([#110853](https://github.com/kubernetes/kubernetes/pull/110853), [@kerthcet](https://github.com/kerthcet)) [SIG Scheduling]
- Kubeadm: fix the bug that configurable KubernetesVersion not respected during kubeadm join ([#111022](https://github.com/kubernetes/kubernetes/pull/111022), [@SataQiu](https://github.com/SataQiu)) [SIG Cluster Lifecycle]
- Reduced time taken to sync proxy rules on Windows kube-proxy with kernelspace mode ([#110702](https://github.com/kubernetes/kubernetes/pull/110702), [@daschott](https://github.com/daschott)) [SIG Network and Windows]
- Updated cAdvisor to v0.43.1 to pick up a kubelet fix where network metrics can be missing in some cases when used with containerd ([#111013](https://github.com/kubernetes/kubernetes/pull/111013), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/google/cadvisor: [v0.43.0 → v0.43.1](https://github.com/google/cadvisor/compare/v0.43.0...v0.43.1)
#### Removed
_Nothing has changed._



### etcd [3.5.4](https://github.com/etcd-io/etcd/releases/tag/v3.5.4)

Not found


### aws-cloud-controller-manager [1.23.1](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.23.1)

Not found


### app-operator [6.3.0](https://github.com/giantswarm/app-operator/releases/tag/v6.3.0)

#### Added
- If no userconfig configmap or secret reference is specified but one is found following the default naming convention (`*-user-values` / `*-user-secrets`) then the App resource is updated to reference the found configmap/secret.



### aws-ebs-csi-driver [2.16.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.16.1)

#### Fixed
- Changing controller `httpEndpoint` to `8610` because of overlapping ports.



### coredns [1.11.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.11.0)

#### Changed
- Update `coredns` to upstream version [1.9.3](https://coredns.io/2022/05/27/coredns-1.9.3-release/).



### cluster-autoscaler [1.23.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.23.1)

#### Changed
- Update cluster-autoscaler to version `1.23.1`.
#### [1.22.2-gs7] - 2022-06-27
#### Changed
- Enable balance similar nodepools by default
#### Fixed
- Ignore labels to consider nodepools similar groups
#### Added
- Support to add extra arguments
#### [1.22.2-gs6] - 2022-04-07
#### [1.22.2-gs5] - 2022-04-06
#### Added
- Support cloud provider alias names (GCP -> GCE)
#### [1.22.2-gs4] - 2022-03-08
#### Fixed
- Updated to correct cluster-autoscaler version
- Use GS-built 1.22 image to deliver upstream unreleased fix https://github.com/kubernetes/autoscaler/pull/4600
#### [1.22.2-gs3] - 2022-02-07
#### Added
- Added support for specifying `balance-similar-node-groups` flag
#### [1.22.2-gs2] - 2022-01-14
#### Fixed
- Fix RBAC for version 1.22.
#### [1.22.2-gs1] - 2022-01-11
#### Changed
- Updated cluster-autoscaler to version `1.22.2`.
#### [1.21.2-gs1] - 2021-12-21
#### Changed
- Updated cluster-autoscaler to version `1.21.2`.
#### [1.21.0-gs2] - 2021-09-29
- Fix RBAC for cluster autoscaler 1.21.
#### [1.21.0-gs1] - 2021-09-13
#### Changed
- Updated cluster-autoscaler to version `1.21.0`.



### metrics-server [1.8.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.8.0)

#### Changed
- Updated metrics-server version to 0.6.1.



### vertical-pod-autoscaler [2.4.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.1)

#### Fixed
- Correct selector in admission controller PDB



### cilium [0.2.5](https://github.com/giantswarm/cilium-app/releases/tag/v0.2.5)

#### Changed
- Use retagged images instead of upstream ones.
- Run the default policies creation job in hostNetwork.



### chart-operator [2.26.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.26.0)

#### Changed
- Use `127.0.0.1` as KUBERNETES_SERVICE_HOST when `bootstrapMode` is enabled.



