# :zap: Giant Swarm Release v16.1.0 for AWS :zap:

This release provides stability improvements, bug fixes and security fixes for various components. It also adds AWS CNI prefix delegation support.

**Highlights**
- Kubernetes 1.21.7 support;
- [AWS CNI prefix delegation](https://docs.giantswarm.io/ui-api/management-api/crd/awsclusters.infrastructure.giantswarm.io/#v1alpha2-alpha.cni.aws.giantswarm.io/prefix-delegation);
- AWS EBS CSI driver with volume snapshot support;
- Security fixes:
    * 10 Linux CVEs;
    * 1 containerd CVEs.

## Change details


### aws-cni [1.9.3](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.9.3)

* Improvement - [Update golang](https://github.com/aws/amazon-vpc-cni-k8s/pull/1665) (#1665, [@jayanthvn](https://github.com/jayanthvn))
* Improvement - [Pod startup latency with Calico and EKS](https://github.com/aws/amazon-vpc-cni-k8s/pull/1629) (#1629, [@jayanthvn](https://github.com/jayanthvn))
* Bug - [Make error count granular](https://github.com/aws/amazon-vpc-cni-k8s/pull/1651) (#1651, [@jayanthvn](https://github.com/jayanthvn))
* Bug - [ServiceAccount should precede DaemonSet in yaml aws](https://github.com/aws/amazon-vpc-cni-k8s/pull/1637) (#1637, [@sramabad1](https://github.com/sramabad1))
* Bug - [Fallback for get hypervisor type and eni ipv4 limits](https://github.com/aws/amazon-vpc-cni-k8s/pull/1616) (#1616, [@jayanthvn](https://github.com/jayanthvn))
* Bug - [fix typo and regenerate limits file ](https://github.com/aws/amazon-vpc-cni-k8s/pull/1597) (#1597, [@jayanthvn](https://github.com/jayanthvn))
* Testing - [Enable unit tests upon PR to release branch](https://github.com/aws/amazon-vpc-cni-k8s/pull/1684) (#1684, [@vikasmb](https://github.com/vikasmb))
* Testing - [Upgrade EKS cluster version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1680) (#1680, [@vikasmb](https://github.com/vikasmb)) 
* Testing - [UTs for no_manage=false](https://github.com/aws/amazon-vpc-cni-k8s/pull/1612) (#1612, [@jayanthvn](https://github.com/jayanthvn))
* Testing - [Run integration test on release branch](https://github.com/aws/amazon-vpc-cni-k8s/pull/1615) (#1615, [@vikasmb](https://github.com/vikasmb))
* Enhancement - [Support DISABLE_NETWORK_RESOURCE_PROVISIONING](https://github.com/aws/amazon-vpc-cni-k8s/pull/1586) (#1586, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Allow reconciler retry for InsufficientCIDR EC2 error](https://github.com/aws/amazon-vpc-cni-k8s/pull/1585) (#1585, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Support for setting no_manage=false](https://github.com/aws/amazon-vpc-cni-k8s/pull/1607) (#1607, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Support for m6i instances](https://github.com/aws/amazon-vpc-cni-k8s/pull/1601) (#1601, [@causton81](https://github.com/causton81))



### aws-ebs-csi-driver [2.7.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.7.1)

#### Changed
- Bump aws-ebs-csi-driver version to v1.4.0.
- Pre-Hook for snapshot CRDs.
- Use deployment for external-snapshot-controller.
- Change VolumeSnapshotter CRDs storage version from v1beta1 to v1.

#### Fixed
- Move CRD's to template.

#### Added
- Add common labels to allow helm operating on CRD's.



### aws-operator [10.10.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.10.1)

#### Added

- Adding latest flatcar images.
- Introduce AWS CNI Prefix delegation.

#### Changed

- Use k8smetadata for annotations.

#### Fixed

- Setting `kubernetes.io/replace/internal-elb` tag on private subnet TCNP stack.


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



### containerlinux [2905.2.6](https://www.flatcar-linux.org/releases/#release-2905.2.6)

**Security fixes**

* Linux
    - [CVE-2021-3764](https://nvd.nist.gov/vuln/detail/CVE-2021-3764)
    - [CVE-2021-3744](https://nvd.nist.gov/vuln/detail/CVE-2021-3744)
    - [CVE-2021-38300](https://nvd.nist.gov/vuln/detail/CVE-2021-38300)
    - [CVE-2021-20321](https://nvd.nist.gov/vuln/detail/CVE-2021-20321)
    - [CVE-2021-41864](https://nvd.nist.gov/vuln/detail/CVE-2021-41864)
    - [CVE-2021-41073](https://nvd.nist.gov/vuln/detail/CVE-2021-41073)
    - [CVE-2020-16119](https://nvd.nist.gov/vuln/detail/CVE-2020-16119)
    - [CVE-2021-3753](https://nvd.nist.gov/vuln/detail/CVE-2021-3753)
    - [CVE-2021-3739](https://nvd.nist.gov/vuln/detail/CVE-2021-3739)
    - [CVE-2021-40490](https://nvd.nist.gov/vuln/detail/CVE-2021-40490)
* containerd
    - [CVE-2021-41103](https://nvd.nist.gov/vuln/detail/CVE-2021-41103)


**Bux fixes**

* The tcsd service for TPM 1 is not started on machines with TPM 2 anymore where it fails and isn’t necessary ([flatcar-linux/coreos-overlay#1364](https://github.com/flatcar-linux/coreos-overlay/pull/1364))
* The Mellanox NIC Linux driver issue introduced in the previous release was fixed ([Flatcar#520](https://github.com/flatcar-linux/Flatcar/issues/520))



**Updates**

Linux ([5.10.75](https://lwn.net/Articles/873465/))
ca-certificates ([3.69.1](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_69_1.html#nss-3-69-1-release-notes))
containerd ([1.5.7](https://github.com/containerd))



### kiam-watchdog [0.4.0](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.4.0)

#### Added

- Add node-selector and tolerations.



### kube-state-metrics [1.5.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.5.1)

#### Changed

- Update architect-orb to v4.8.0 to stop pushing to app collection helm chart.
- Bumped to upstream version v2.2.4.

#### Fixed

- Fix permission to list and watch `leases.coordination.k8s.io`.



### kubernetes [1.21.7](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.7)

#### API Change
- Kube-apiserver: Fixes handling of CRD schemas containing literal null values in enums (#104989, @liggitt) [SIG API Machinery, Apps and Network]

#### Feature
- Kubernetes is now built with Golang 1.16.10 (#106224, @cpanato) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Update debian-base, debian-iptables, setcap images to pick up CVE fixes
  - Debian-base to v1.9.0
  - Debian-iptables to v1.6.7
  - setcap to v2.0.4 (#106147, @cpanato) [SIG Release and Testing]
- Kubernetes is now built with Golang 1.16.9 (#105672, @cpanato) [SIG Cloud Provider, Instrumentation, Release and Testing]

#### Failing Test
- Fixes hostpath storage e2e tests within SELinux enabled env (#105787, @Elbehery) [SIG Testing]
- Backport: e2e.test now uses the secure port to retrieve metrics data to ensure compatibility with 1.21 and 1.22 (for upgrade tests) (#104328, @pohly) [SIG API Machinery, Instrumentation, Storage and Testing]

#### Bug or Regression
- EndpointSlice Mirroring controller now cleans up managed EndpointSlices when a Service selector is added (#106135, @robscott) [SIG Apps, Network and Testing]
- Fix a bug that `--disabled-metrics` doesn't function well. (#106391, @Huang-Wei) [SIG API Machinery, Cluster Lifecycle and Instrumentation]
- Fix a panic in kubectl when creating secrets with an improper output type (#106354, @lauchokyip) [SIG CLI]
- Fix concurrent map access causing panics when logging timed-out API calls. (#106113, @marseel) [SIG API Machinery]
- Fixed very rare volume corruption when a pod is deleted while kubelet is offline.
  Retry FibreChannel devices cleanup after error to ensure FC device is detached before it can be used on another node. (#102656, @jsafrane) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Support more than 100 disk mounts on Windows (#105673, @andyzhangx) [SIG Storage and Windows]
- Aggregate errors when putting vmss
  fix: consolidate logs for instance not found error (#105365, @nilo19) [SIG Cloud Provider]
- Allow CSI drivers to just run offline expansion tests (#102740, @gnufied) [SIG Storage and Testing]
- Bump klog to v2.9.0, fixing log lines that render as byte arrays (#105407, @ehashman) [SIG Architecture, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation and Storage]
- Fix overriding logs files on reboot. (#105614, @rphillips) [SIG Node]
- Fix winkernel kube-proxy to only use dual stack when host and networking supports it (#101047, @jsturtevant) [SIG Network and Windows]
- Fix: ignore not a VMSS error for VMAS nodes in EnsureBackendPoolDeleted. (#105402, @ialidzhikov) [SIG Cloud Provider]
- Revert PR #102925 which introduced unexpected scheduling behavior based on balanced resource allocation (#105238, @damemi) [SIG Scheduling]
- Updates golang.org/x/text to v0.3.6 to fix CVE-2020-28852 (#102601, @jonesbr17) [SIG API Machinery, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node and Storage]

#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- k8s.io/kube-openapi: 591a79e → 3cc51fd
- k8s.io/utils: 67b214c → da69540
- k8s.io/klog/v2: v2.8.0 → v2.9.0
- golang.org/x/text: v0.3.4 → v0.3.6
#### Removed
_Nothing has changed._


