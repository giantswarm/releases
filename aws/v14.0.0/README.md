# :zap: Giant Swarm Release v14.0.0 for AWS :zap:

This release provides support for Kubernetes 1.19 on AWS.

Please note that with version `1.19` there a few breaking changes in the `Kubernetes` APIs.
Please refer to the [upstream documentation](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.19.md#urgent-upgrade-notes)
and feel free to get in touch with your solutions engineer for any concern you might have.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

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

### app-operator [3.1.0](https://github.com/giantswarm/app-operator/releases/tag/v3.1.0)

#### Changed
- Enable mutating and validating webhooks in app-admission-controller for
tenant app CRs.

#### Added
- Make resync period configurable for use in integration tests.
- Pause App CR reconciliation when it has
-  `app-operator.giantswarm.io/paused=true` annotation.
- Print difference between the current chart and desired chart.


### aws-operator [10.2.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.2.0)

#### Added
- Allow incoming NFS traffic on node pools for Elastic File System.
- Add support for tagging AWS resources, managed by the operator, based on the custom resource labels.
- Add `cleanupiamroles` resource for detaching third party policies from our IAM roles.
- Update `k8scloudconfig` version to `v10.0.0` to include change for Kubernetes 1.19.
- Allow configuration of `MINIMUM_IP_TARGET` and `WARM_IP_TARGET` for AWS CNI via annotations on `AWSCluster`

#### Changed
- Include Account ID in the s3bucket for access logs. It is a breaking change, that will put access logs to a new s3 bucket.
- Change AWS CNI and AWS CNI k8s plugin log verbosity to `INFO`.
- Change AWS CNI log file to `stdout`.
- Add retry logic for decrypt units to avoid flapping.
- Use values generated by `config-controller` to deploy `aws-operator` instead of catalog values.
- Use `giantswarm/config` versions matching `v1.x.x` major.
- Start updating `tcnp` CF stack only when `tccpn` CF stack is already updated. This ensure that master nodes are updated before worker nodes.


### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

- Fix server panic when force-new-cluster flag is enabled in a cluster which had learner node.

### aws-cni [1.7.8](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.7.8)

* Improvement - [Replace DescribeNetworkInterfaces with paginated version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1333) (#1333, @haouc)


### containerlinux [2605.12.0](https://www.flatcar-linux.org/releases/#release-2605.12.0)

**Security fixes**
* Linux
   - [CVE-2020-27815](https://www.openwall.com/lists/oss-security/2020/11/30/5)
   - [CVE-2020-29568](https://nvd.nist.gov/vuln/detail/CVE-2020-29568)
   - [CVE-2020-29569](https://nvd.nist.gov/vuln/detail/CVE-2020-29569)
   - [CVE-2020-28374](https://nvd.nist.gov/vuln/detail/CVE-2020-28374)
   - [CVE-2020-36158](https://nvd.nist.gov/vuln/detail/CVE-2020-36158)
*   go - [CVE-2021-3114](https://github.com/golang/go/issues/43786)
*   sudo - [CVE-2021-3156](https://nvd.nist.gov/vuln/detail/CVE-2021-3156), [CVE-2021-23239](https://nvd.nist.gov/vuln/detail/CVE-2021-23239)

**Bug fixes**

*   `/etc/iscsi/initiatorname.iscsi` is generated by the iscsi-init service ([#321](https://github.com/kinvolk/Flatcar/issues/321))
*   Prevent iscsiadm buffer overflow ([#318](https://github.com/kinvolk/Flatcar/issues/318))

**Changes**

*   Revert to building docker and containerd with go1.13 instead of go1.15. This reduces the SIGURG log spam ([Issue #315](https://github.com/kinvolk/Flatcar/issues/315) [PR #774](https://github.com/kinvolk/coreos-overlay/pull/774))
*   The containerd socket is now available in the default location (`/run/containerd/containerd.sock`) and also as a symlink in the previous location (`/run/docker/libcontainerd/docker-containerd.sock`) ([#771](https://github.com/kinvolk/coreos-overlay/pull/771))
*   With the iscsi update, the service unit has changed from iscsid to iscsi ([#791](https://github.com/kinvolk/coreos-overlay/pull/791))
*   AWS Pro: include scripts to facilitate setup of EKS workers ([#794](https://github.com/kinvolk/coreos-overlay/pull/794)).
*   Missed from earlier notes: with the previous open-iscsi update to 2.1.2, the service unit name changed from iscsid to iscsi ([#682](https://github.com/kinvolk/coreos-overlay/pull/682))

**Updates**

*   linux ([5.4.92](https://lwn.net/Articles/843687/))
*   open-iscsi ([2.1.3](https://github.com/open-iscsi/open-iscsi/releases/tag/2.1.3))
*   go ([1.15.7](https://go.googlesource.com/go/+/refs/tags/go1.15.7))
*   sudo ([1.9.5p2](https://github.com/sudo-project/sudo/releases/tag/SUDO_1_9_5p2))


### chart-operator [2.6.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.6.0)

#### Added
- Print difference between current release and desired release.

#### Changed
- Updated Helm to v3.4.2.


### cluster-autoscaler [1.19.1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.1)

#### Changed
- Updated cluster-autoscaler to version `1.19.1`.


### cert-manager [2.4.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.4.1)

#### Changed
- Made backoffLimit for clusterissuer job configurable. ([#125](https://github.com/giantswarm/cert-manager-app/pull/125))
- Updated clusterissuer subchart API groups to `cert-manager.io/v1`. ([#124](https://github.com/giantswarm/cert-manager-app/pull/124))


### cert-exporter [1.5.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.5.0)

#### Changed
- Check ca.crt expiries in TLS secrets. ([#109](https://github.com/giantswarm/cert-exporter/pull/109))


### chart-operator [2.7.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.7.1)

#### Fixed
- Only create VPA if autoscaling API group is present.


### kiam [1.7.0](https://github.com/giantswarm/kiam-app/releases/tag/v1.7.0)

- Add taint tolerations for kiam agent and kiam server.


### metrics-server [1.2.1](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.2.1)

- Push app to control plane catalogs


### node-exporter [1.7.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.1)

#### Changed
- Use the domain registry from installation values if it is present.


### external-dns [2.1.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.1.1)

#### Added
- Allow the sync policy to be configured. 
- Supports customisation of the txt-owner-id (whilst still defaulting for default apps).
- Supports dry-run mode and warns the user if enabled.

#### Changed
- Reworked the App to prepare it for customer use. ([#49](https://github.com/giantswarm/external-dns-app/pull/49))
  - General:
    - Pushes the app to the giantswarm app catalog.
    - Uses Helm release namespace.
    - Uses the release name for resource naming to avoid conflicts.
    - Added a values schema to catch incorrect values.
    - Generally makes the chart easier to use (fully documented values file).
  - external-dns options:
    - Allows customisation of the txt registry prefix.
    - Allows configuration of synchronisation interval.
    - Filter resources to reconcile via annotations.
  - AWS-specifc:
    - Allows the user to provide an IAM role to use.
    - Allows the user to provide the list of domains for external-dns to manage.
    - Allows configuration of batch size.
    - Allows configuration of CNAME instead of ALIAS records.
    - Allows configuration of the AWS zone type to update.
