# :zap: Giant Swarm Release v14.0.0 for AWS :zap:

<< Add description here >>

## Change details


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



### app-operator [3.0.0](https://github.com/giantswarm/app-operator/releases/tag/v3.0.0)

#### Changed
- Enable mutating and validating webhooks in app-admission-controller for
tenant app CRs.
#### Added
- Make resync period configurable for use in integration tests.
- Pause App CR reconciliation when it has
  `app-operator.giantswarm.io/paused=true` annotation.
- Print difference between the current chart and desired chart.



### chart-operator [2.6.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.6.0)

#### Added
- Print difference between current release and desired release.
#### Changed
- Updated Helm to v3.4.2.



### aws-operator [9.3.6-dev](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.6-dev)

Not found


### containerlinux [2605.11.0](https://www.flatcar-linux.org/releases/#release-2605.11.0)

**Security fixes**

 * Linux
   - [CVE-2020-27815](https://www.openwall.com/lists/oss-security/2020/11/30/5)
   - [CVE-2020-29568](https://nvd.nist.gov/vuln/detail/CVE-2020-29568)
   - [CVE-2020-29569](https://nvd.nist.gov/vuln/detail/CVE-2020-29569)

**Bug fixes**

*   networkd: avoid managing MAC addresses for veth devices ([kinvolk/init#33](https://github.com/kinvolk/init/pull/33))

**Updates**

*   Linux ([5.4.87](https://lwn.net/Articles/841900/))



### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

Not found


### cluster-autoscaler [1.19.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.1)

#### Changed
- Updated cluster-autoscaler to version `1.19.1`.



