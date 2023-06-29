# :zap: Giant Swarm Release v19.1.0 for AWS :zap:

<< Add description here >>

## Change details


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



### cert-manager [2.24.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.24.1)

#### Added
- Add `cluster-autoscaler safe-to-evict` annotation to `controller` and `cainjector` through newly introduced `controller.podAnnotations` and `cainjector.podAnnotations` values. ([#330](https://github.com/giantswarm/cert-manager-app/pull/330))



### net-exporter [1.16.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.16.2)

#### Changed
- Reduce CPU and Mem requests.



### node-exporter [1.16.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.16.1)

#### Changed
- Enable service monitor.



### observability-bundle [0.7.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.7.1)

#### Changed
- Upgrade `promtail-app` to 1.1.1.
- Upgrade `prometheus-operator-app` to 5.0.6.



### vertical-pod-autoscaler [3.5.3](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.3)

#### Added
- Add `cluster-autoscaler safe-to-evict` annotation to `recommender` and `updater`



