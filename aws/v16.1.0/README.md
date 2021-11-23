# :zap: Giant Swarm Release v16.1.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-cni [1.9.3](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.9.3)

* Improvement - [Update golang](https://github.com/aws/amazon-vpc-cni-k8s/pull/1665) (#1665, [@jayanthvn](https://github.com/jayanthvn))
* Improvement - [Pod startup latency with Calico and EKS](https://github.com/aws/amazon-vpc-cni-k8s/pull/1629) (#1629, [@jayanthvn](https://github.com/jayanthvn))
* Bug - [Make error count granular](https://github.com/aws/amazon-vpc-cni-k8s/pull/1651) (#1651, [@jayanthvn](https://github.com/jayanthvn))
* Bug - [ServiceAccount should precede DaemonSet in yaml aws](https://github.com/aws/amazon-vpc-cni-k8s/pull/1637) (#1637, [@sramabad1](https://github.com/sramabad1))
* Testing - [Enable unit tests upon PR to release branch](https://github.com/aws/amazon-vpc-cni-k8s/pull/1684) (#1684, [@vikasmb](https://github.com/vikasmb))
* Testing - [Upgrade EKS cluster version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1680) (#1680, [@vikasmb](https://github.com/vikasmb)) 



### aws-ebs-csi-driver [2.7.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.7.1)

#### Changed
- Bump aws-ebs-csi-driver version to v1.4.0.
- Pre-Hook for snapshot CRDs.
- Use deployment for external-snapshot-controller.
- Change VolumeSnapshotter CRDs storage version from v1beta1 to v1.



### aws-operator [10.10.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.10.0)

#### Added

- Adding latest flatcar images.
- Introduce AWS CNI Prefix delegation.

#### Changed

- Use k8smetadata for annotations.



### cert-exporter [2.0.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.0.0)

#### Changed
- Export presence of `giantswarm.io/service-type: managed` label in cert-manager `Issuer` and `ClusterIssuer` CR referenced by `Certificate` CR `issuerRef` spec field to `cert_exporter_certificate_cr_not_after` metric as `managed_issuer` label.
- Add `--monitor-files` and `--monitor-secrets` flags.
- Add Deployment to helm chart to avoid exporting secrets and certificate metrics from DaemonSets.
- Build container image using retagged giantswarm alpine.
- Run as non-root inside container.



### cluster-operator [3.11.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.11.0)

#### Added
- Check if `kiam-watchdog` app has to be enabled.
- Add cluster CA to cluster values configmap for apps like dex that need to
verify it.



### kiam-watchdog [0.4.0](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.4.0)

#### Added

- Add node-selector and tolerations.



### kubernetes [1.21.7](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.7)

#### Feature
- Kubernetes is now built with Golang 1.16.10 (#106224, @cpanato) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update debian-base, debian-iptables, setcap images to pick up CVE fixes
  - Debian-base to v1.9.0
  - Debian-iptables to v1.6.7
  - setcap to v2.0.4 (#106147, @cpanato) [SIG Release and Testing]
#### Failing Test
- Fixes hostpath storage e2e tests within SELinux enabled env (#105787, @Elbehery) [SIG Testing]
#### Bug or Regression
- EndpointSlice Mirroring controller now cleans up managed EndpointSlices when a Service selector is added (#106135, @robscott) [SIG Apps, Network and Testing]
- Fix a bug that `--disabled-metrics` doesn't function well. (#106391, @Huang-Wei) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fix a panic in kubectl when creating secrets with an improper output type (#106354, @lauchokyip) [SIG CLI]
- Fix concurrent map access causing panics when logging timed-out API calls. (#106113, @marseel) [SIG API Machinery]
- Fixed very rare volume corruption when a pod is deleted while kubelet is offline.
  Retry FibreChannel devices cleanup after error to ensure FC device is detached before it can be used on another node. (#102656, @jsafrane) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Support more than 100 disk mounts on Windows (#105673, @andyzhangx) [SIG Storage and Windows]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- k8s.io/kube-openapi: 591a79e → 3cc51fd
- k8s.io/utils: 67b214c → da69540
#### Removed
_Nothing has changed._


