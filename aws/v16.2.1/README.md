# :zap: Giant Swarm Release v16.2.1 for AWS :zap:

This release downgrades Flatcar to version `2905.2.6` to restore version 1 of the kernel cgroups feature.
It also bumps to latest Kubernetes 1.21 patch release and to latest 1.21 cluster autoscaler.

## Change details

### containerlinux [2905.2.6](https://www.flatcar-linux.org/releases/#release-2905.2.6)

Downgraded from 2983.2.0 to restore Cgroups v1.

### kubernetes [1.21.8](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.8)

#### Feature
- Kubernetes is now built with Golang 1.16.11 ([#106839](https://github.com/kubernetes/kubernetes/pull/106839), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Kubernetes is now built with Golang 1.16.12 ([#106983](https://github.com/kubernetes/kubernetes/pull/106983), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update golang.org/x/net to v0.0.0-20211209124913-491a49abca63 ([#106961](https://github.com/kubernetes/kubernetes/pull/106961), [@cpanato](https://github.com/cpanato)) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]
#### Bug or Regression
- Fix: skip instance not found when decoupling vmss from lb ([#105835](https://github.com/kubernetes/kubernetes/pull/105835), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fixed SELinux relabeling of CSI volumes after CSI driver failure. ([#106553](https://github.com/kubernetes/kubernetes/pull/106553), [@jsafrane](https://github.com/jsafrane)) [SIG Node and Storage]
- Kubeadm: allow the "certs check-expiration" command to not require the existence of the cluster CA key (ca.key file) when checking the expiration of managed certificates in kubeconfig files. ([#106929](https://github.com/kubernetes/kubernetes/pull/106929), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- Kubeadm: during execution of the "check expiration" command, treat the etcd CA as external if there is a missing etcd CA key file (etcd/ca.key) and perform the proper validation on certificates signed by the etcd CA. Additionally, make sure that the CA for all entries in the output table is included - for both certificates on disk and in kubeconfig files. ([#106924](https://github.com/kubernetes/kubernetes/pull/106924), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
- The scheduler's assumed pods have 2min instead of 30s to receive nodeName pod updates ([#106632](https://github.com/kubernetes/kubernetes/pull/106632), [@ahg-g](https://github.com/ahg-g)) [SIG Scheduling]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- golang.org/x/net: 3d97a24 → 491a49a
#### Removed
_Nothing has changed._



### cluster-autoscaler [1.21.2-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.2-gs1)

### Changed

- Upgraded to upstream version 1.21.2.


### cert-exporter [2.0.1](https://github.com/giantswarm/cert-exporter/releases/tag/v2.0.1)

#### Changed
- Equalise labels in the helm chart.



### external-dns [2.7.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.7.0)

#### Changed
- Upgrade upstream external-dns from v0.9.0 to [v0.10.2](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.10.2). The new release brings a lot of smaller improvements and bug fixes.
- Remove support for Kubernetes <= 1.18.
#### Fixed
- Fix dry-run option.



### cert-manager [2.12.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.12.0)

#### Changed
- Upgrade to upstream image [`v1.6.1`](https://github.com/jetstack/cert-manager/releases/tag/v1.6.1) ([#204](https://github.com/giantswarm/cert-manager-app/pull/204)). This version stops serving cert-manager API versions `v1alpha2, v1alpha3, and v1beta1`. If you need to upgrade your resources, [this document](https://cert-manager.io/docs/installation/upgrading/remove-deprecated-apis/#upgrading-existing-cert-manager-resources) explains the process.



