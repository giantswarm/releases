# :zap: Giant Swarm Release v13.0.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.18.6](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.6)

#### API Change
- Fix bug in reflector that couldn't recover from "Too large resource version" errors ([#92537](https://github.com/kubernetes/kubernetes/pull/92537), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery]
#### Bug or Regression
- CVE-2020-8557 (Medium): Node-local denial of service via container /etc/hosts file. See https://github.com/kubernetes/kubernetes/issues/93032 for more details. ([#92916](https://github.com/kubernetes/kubernetes/pull/92916), [@joelsmith](https://github.com/joelsmith)) [SIG Node]
- Containers which specify a `startupProbe` but not a `readinessProbe` were previously considered "ready" before the `startupProbe` completed, but are now considered "not-ready". ([#92196](https://github.com/kubernetes/kubernetes/pull/92196), [@thockin](https://github.com/thockin)) [SIG Node]
- Extend kube-apiserver /readyz with new "informer-sync" check ensuring that internal informers are synced. ([#92644](https://github.com/kubernetes/kubernetes/pull/92644), [@wojtek-t](https://github.com/wojtek-t)) [SIG API Machinery and Testing]
- Fix throttling issues when Azure VM computer name prefix is different from VMSS name ([#92793](https://github.com/kubernetes/kubernetes/pull/92793), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix: GetLabelsForVolume panic issue for azure disk PV ([#92166](https://github.com/kubernetes/kubernetes/pull/92166), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: don't use docker config cache if it's empty ([#92330](https://github.com/kubernetes/kubernetes/pull/92330), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: use force detach for azure disk ([#91948](https://github.com/kubernetes/kubernetes/pull/91948), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixes a problem with 63-second or 1-second connection delays with some VXLAN-based
  network plugins which was first widely noticed in 1.16 (though some users saw it
  earlier than that, possibly only with specific network plugins). If you were previously
  using ethtool to disable checksum offload on your primary network interface, you should
  now be able to stop doing that. ([#92035](https://github.com/kubernetes/kubernetes/pull/92035), [@danwinship](https://github.com/danwinship)) [SIG Network and Node]
- Kubeadm: add the deprecated flag --port=0 to kube-controller-manager and kube-scheduler manifests to disable insecure serving. Without this flag the components by default serve (e.g. /metrics) insecurely on the default node interface (controlled by --address). Users that wish to override this behavior and enable insecure serving can pass a custom --port=X via kubeadm's "extraArgs" mechanic for these components. ([#92720](https://github.com/kubernetes/kubernetes/pull/92720), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: during "join", don't re-add an etcd member if it already exists in the cluster. ([#92118](https://github.com/kubernetes/kubernetes/pull/92118), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- hyperkube: Use debian-hyperkube-base@v1.1.1 image
  
    Includes iproute2 to fix a regression in hyperkube images
    when using hyperkube as a kubelet ([#92624](https://github.com/kubernetes/kubernetes/pull/92624), [@justaugustus](https://github.com/justaugustus)) [SIG Cluster Lifecycle, Network and Release]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



