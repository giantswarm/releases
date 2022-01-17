# :zap: Giant Swarm Release v17.0.0-alpha1 for Azure :zap:

<< Add description here >>

## Change details


### app-operator [5.4.0](https://github.com/giantswarm/app-operator/releases/tag/v5.4.0)

#### Changed
- Update Helm to `v3.6.3`.
- Use controller-runtime client to remove CAPI dependency.
- Use `apptestctl` to install CRDs in integration tests to avoid hitting GitHub rate limits.
#### Removed
- Remove `releasemigration` resource now migration to Helm 3 is complete.



### cert-operator [1.3.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.3.0)

#### Changed
- Use `RenewSelf` instead of `LookupSelf` to prevent expiration of `Vault token`.



### kubernetes [1.22.5](https://github.com/kubernetes/kubernetes/releases/tag/v1.22.5)

#### Feature
- Kubernetes is now built with Golang 1.16.11 ([#106837](https://github.com/kubernetes/kubernetes/pull/106837), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built with Golang 1.16.12 ([#106982](https://github.com/kubernetes/kubernetes/pull/106982), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update golang.org/x/net to v0.0.0-20211209124913-491a49abca63 ([#106960](https://github.com/kubernetes/kubernetes/pull/106960), [@cpanato](https://github.com/cpanato)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
#### Bug or Regression
- A pod that the Kubelet rejects was still considered as being accepted for a brief period of time after rejection, which might cause some pods to be rejected briefly that could fit on the node.  A pod that is still terminating (but has status indicating it has failed) may also still be consuming resources and so should also be considered. ([#104918](https://github.com/kubernetes/kubernetes/pull/104918), [@ehashman](https://github.com/ehashman)) [SIG Node]
- Fix: skip instance not found when decoupling vmss from lb ([#105836](https://github.com/kubernetes/kubernetes/pull/105836), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Kubeadm: allow the "certs check-expiration" command to not require the existence of the cluster CA key (ca.key file) when checking the expiration of managed certificates in kubeconfig files. ([#106930](https://github.com/kubernetes/kubernetes/pull/106930), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: during execution of the "check expiration" command, treat the etcd CA as external if there is a missing etcd CA key file (etcd/ca.key) and perform the proper validation on certificates signed by the etcd CA. Additionally, make sure that the CA for all entries in the output table is included - for both certificates on disk and in kubeconfig files. ([#106925](https://github.com/kubernetes/kubernetes/pull/106925), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Respect grace period when updating static pods. ([#106394](https://github.com/kubernetes/kubernetes/pull/106394), [@gjkim42](https://github.com/gjkim42)) [SIG Node and Testing]
- Reverts graceful node shutdown to match 1.21 behavior of setting pods that have not yet successfully completed to "Failed" phase if the GracefulNodeShutdown feature is enabled in kubelet. The GracefulNodeShutdown feature is beta and must be explicitly configured via kubelet config to be enabled in 1.21+. This changes 1.22 and 1.23 behavior on node shutdown to match 1.21. If you do not want pods to be marked terminated on node shutdown in 1.22 and 1.23, disable the GracefulNodeShutdown feature. ([#106899](https://github.com/kubernetes/kubernetes/pull/106899), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
- Scheduler's assumed pods have 2min instead of 30s to receive nodeName pod updates ([#106633](https://github.com/kubernetes/kubernetes/pull/106633), [@ahg-g](https://github.com/ahg-g)) [SIG Scheduling]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- golang.org/x/net: 37e1c6a → 491a49a
#### Removed
_Nothing has changed._



### containerlinux [3033.2.0](https://www.flatcar-linux.org/releases/#release-3033.2.0)

Containerlinux release "3033.2.0" was not found in the changelog


### azure-operator [5.13.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.13.0)

#### Changed
- Bumped k8scc to latest version to support Calico 3.21.



### calico [3.21.3](https://github.com/projectcalico/calico/releases/tag/v3.21.3)

Not found


### chart-operator [2.20.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.20.0)

#### Changed
- Update Helm to v3.6.3.
- Use controller-runtime client to remove CAPI dependency.
#### Removed
- Remove unused helm 2 release collector.



### coredns [1.7.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.7.0)

#### Changed
- Update `coredns` to upstream version [1.8.6](https://coredns.io/2021/10/07/coredns-1.8.6-release/).



### kube-state-metrics [1.6.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.6.0)

#### Changed
- Bumped to upstream version v2.3.0.



### metrics-server [1.6.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.6.0)

#### Changed
- Updated metrics-server version to 0.5.2.



### node-exporter [1.9.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.9.0)

#### Updated
- Updated node-exporter version to 1.3.1.



### cluster-autoscaler [1.22.2-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs1)

Not found


