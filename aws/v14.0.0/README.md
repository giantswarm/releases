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



### app-operator [3.1.0](https://github.com/giantswarm/app-operator/releases/tag/v3.1.0)

#### Changed
- Enable mutating and validating webhooks in app-admission-controller for
tenant app CRs.
#### Added
- Make resync period configurable for use in integration tests.
- Pause App CR reconciliation when it has
-  `app-operator.giantswarm.io/paused=true` annotation.
- Print difference between the current chart and desired chart.



### aws-operator [10.0.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.0.0)

#### Added
- Add `cleanupiamroles` resource for detaching third party policies from our IAM
  roles.
- Update `k8scloudconfig` version to `v10.0.0` to include change for Kubernetes 1.19.
- Allow configuration of `MINIMUM_IP_TARGET` and `WARM_IP_TARGET` for AWS CNI via annotations on `AWSCluster`
#### Changed
- Include Account ID in the s3bucket for access logs. It is a breaking change, that will put access logs to a new s3 bucket.
- Change AWS CNI and AWS CNI k8s plugin log verbosity to `INFO`.
- Change AWS CNI log file to `stdout`.
- Add retry logic for decrypt units to avoid flapping.



### etcd [3.4.14](https://github.com/etcd-io/etcd/releases/tag/v3.4.14)

Not found


### aws-cni [1.7.8](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.7.8)

* Improvement - [Replace DescribeNetworkInterfaces with paginated version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1333) (#1333, @haouc)



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



### external-dns [1.6.0](https://github.com/giantswarm/external-dns-app/releases/tag/v1.6.0)

#### Changed
- Upgrade upstream external-dns from v0.7.4 to [v0.7.6](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.6).



