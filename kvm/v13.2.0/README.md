# :zap: Giant Swarm Release v13.2.0 for KVM :zap:

<< Add description here >>

## Change details


### kvm-operator [3.15.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.15.0)

#### Added
- Add vertical pod autoscaler for operator pods.
- Automatically delete TC node pods when NotReady for too long.
#### Changed
- Do not drain node pods when cluster is being deleted.
- Update for Kubernetes 1.19 compatibility.
- Update k8s-kvm to v0.4.1 with QEMU v5.2.0 and Flatcar DNS fix.
#### Fixed
- Use managed-by label to check node deployments are deleted before cluster namespace.
- Remove IPs from endpoints when the corresponding workload cluster node is not ready.



### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

Not found


### cluster-operator [0.24.2](https://github.com/giantswarm/cluster-operator/releases/tag/v0.24.2)

Not found


### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### kubernetes [1.19.4](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.4)

#### Bug or Regression
- An issues preventing volume expand controller to annotate the PVC with `volume.kubernetes.io/storage-resizer` when the PVC StorageClass is already updated to the out-of-tree provisioner is now fixed. ([#94489](https://github.com/kubernetes/kubernetes/pull/94489), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG API Machinery, Apps and Storage]
- Cloud node controller: handle empty providerID from getProviderID ([#95452](https://github.com/kubernetes/kubernetes/pull/95452), [@nicolehanjing](https://github.com/nicolehanjing)) [SIG Cloud Provider]
- Disable watchcache for events ([#96052](https://github.com/kubernetes/kubernetes/pull/96052), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
- Disabled `LocalStorageCapacityIsolation` feature gate is honored during scheduling. ([#96140](https://github.com/kubernetes/kubernetes/pull/96140), [@Huang-Wei](https://github.com/Huang-Wei)) [SIG Scheduling]
- Fix a bug that Pods with topologySpreadConstraints get scheduled to nodes without required labels. ([#95880](https://github.com/kubernetes/kubernetes/pull/95880), [@ialidzhikov](https://github.com/ialidzhikov)) [SIG Scheduling]
- Fix azure disk attach failure for disk size bigger than 4TB ([#95463](https://github.com/kubernetes/kubernetes/pull/95463), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix azure disk data loss issue on Windows when unmount disk ([#95456](https://github.com/kubernetes/kubernetes/pull/95456), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider and Storage]
- Fixed a bug causing incorrect formatting of `kubectl describe ingress`. ([#94985](https://github.com/kubernetes/kubernetes/pull/94985), [@howardjohn](https://github.com/howardjohn)) [SIG CLI and Network]
- Fixed a bug in client-go where new clients with customized `Dial`, `Proxy`, `GetCert` config may get stale HTTP transports. ([#95427](https://github.com/kubernetes/kubernetes/pull/95427), [@roycaihw](https://github.com/roycaihw)) [SIG API Machinery]
- Fixed a regression which prevented pods with `docker/default` seccomp annotations from being created in 1.19 if a PodSecurityPolicy was in place which did not allow `runtime/default` seccomp profiles. ([#95990](https://github.com/kubernetes/kubernetes/pull/95990), [@saschagrunert](https://github.com/saschagrunert)) [SIG Auth]
- Fixed kubelet creating extra sandbox for pods with RestartPolicyOnFailure after all containers succeeded ([#92614](https://github.com/kubernetes/kubernetes/pull/92614), [@tnqn](https://github.com/tnqn)) [SIG Node and Testing]
- Fixes high CPU usage in kubectl drain ([#95260](https://github.com/kubernetes/kubernetes/pull/95260), [@amandahla](https://github.com/amandahla)) [SIG CLI]
- If we set SelectPolicy MinPolicySelect on scaleUp behavior or scaleDown behavior,Horizontal Pod Autoscaler doesn't automatically scale the number of pods correctly ([#95647](https://github.com/kubernetes/kubernetes/pull/95647), [@JoshuaAndrew](https://github.com/JoshuaAndrew)) [SIG Apps and Autoscaling]
- Kube-proxy now trims extra spaces found in loadBalancerSourceRanges to match Service validation. ([#94107](https://github.com/kubernetes/kubernetes/pull/94107), [@robscott](https://github.com/robscott)) [SIG Network]
- Kubeadm: add missing "--experimental-patches" flag to "kubeadm init phase control-plane" ([#95786](https://github.com/kubernetes/kubernetes/pull/95786), [@Sh4d1](https://github.com/Sh4d1)) [SIG Cluster Lifecycle]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Changed
- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.



