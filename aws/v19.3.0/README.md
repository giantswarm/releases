# :zap: Giant Swarm Release v19.3.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [6.10.0](https://github.com/giantswarm/app-operator/releases/tag/v6.10.0)

#### Added
- Add option to disable k8s client cache.



### aws-operator [14.24.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.24.0)

#### Added
- Add `global.podSecurityStandards.enforced` value for PSS migration.
- Emit event when an unhealthy node is terminated.
- Bump `badnodedetector` to be able to use `node-problem-detector` app for unhealthy node termination.
- Add a additional IAM permission for `cluster-autoscaler`.
#### Changed
- Bump k8scc to disable PSPs in preparation for switch to PSS.
- Disable cluster autoscaler during rollouts of node pool ASGs.



### cluster-operator [5.10.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.10.0)

#### Changed
- Automatically set `global.podSecurityStandards.enforced` to `true` for clusters >= 19.3.0.



### containerlinux [3602.2.1](https://www.flatcar-linux.org/releases/#release-3602.2.1)

_Changes since **Stable 3602.2.0**_
 
 #### Security fixes:
 
- Linux ([CVE-2023-31085](https://nvd.nist.gov/vuln/detail/CVE-2023-31085), [CVE-2023-34324](https://nvd.nist.gov/vuln/detail/CVE-2023-34324), [CVE-2023-4244](https://nvd.nist.gov/vuln/detail/CVE-2023-4244), [CVE-2023-42754](https://nvd.nist.gov/vuln/detail/CVE-2023-42754),  [CVE-2023-5197](https://nvd.nist.gov/vuln/detail/CVE-2023-5197))
 - curl ([CVE-2023-38545](https://nvd.nist.gov/vuln/detail/CVE-2023-38545), [CVE-2023-38546](https://nvd.nist.gov/vuln/detail/CVE-2023-38546))
 
 #### Bug fixes:
 
 - Disabled systemd-networkd's RoutesToDNS setting by default to fix provisioning failures observed in VMs with multiple network interfaces on Azure ([scripts#1206](https://github.com/flatcar/scripts/pull/1206))
 - Fixed a regression in Docker resulting in file permissions being dropped from exported container images. ([scripts#1231](https://github.com/flatcar/scripts/pull/1231))
 
 #### Changes:
 
 - To make Kubernetes work by default, `/usr/libexec/kubernetes/kubelet-plugins/volume/exec` is now a symlink to the writable folder `/var/kubernetes/kubelet-plugins/volume/exec` ([Flatcar#1193](https://github.com/flatcar/Flatcar/issues/1193))
 
 #### Updates:
 
 - Linux ([5.15.136](https://lwn.net/Articles/948297) (includes [5.15.135](https://lwn.net/Articles/947299), [5.15.134](https://lwn.net/Articles/946855)))
 - ca-certificates ([3.94](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_94.html))


### etcd [3.5.10](https://github.com/etcd-io/etcd/releases/tag/v3.5.10)

#### etcd server
- Fix [`--socket-reuse-port` and `--socket-reuse-address` not able to be set in configuration file](https://github.com/etcd-io/etcd/pull/16435).
- Fix [corruption check may get a `ErrCompacted` error when server has just been compacted](https://github.com/etcd-io/etcd/pull/16048)
- Improve [Lease put performance for the case that auth is disabled or the user is admin](https://github.com/etcd-io/etcd/pull/16019)
- Improve [Skip getting authInfo from incoming context when auth is disabled](https://github.com/etcd-io/etcd/pull/16241)
- Fix [Hash and HashKV have duplicated RESTful API](https://github.com/etcd-io/etcd/pull/16490)
#### etcdutl v3
- Add [optional --bump-revision and --mark-compacted flag to etcdutl snapshot restore operation](https://github.com/etcd-io/etcd/pull/16165).
#### etcdctl v3
- Add [optional --bump-revision and --mark-compacted flag to etcdctl snapshot restore operation](https://github.com/etcd-io/etcd/pull/16165).
#### etcd grpc-proxy
- Fix [Memberlist results not updated when proxy node down](https://github.com/etcd-io/etcd/pull/15907).
#### Package `clientv3`
- Fix [Multiple endpoints with same prefix got mixed up](https://github.com/etcd-io/etcd/pull/15939)
- Fix [Unexpected blocking when barrier waits on a nonexistent key](https://github.com/etcd-io/etcd/pull/16188)
- Fix [Reset auth token when failing to authenticate due to auth being disabled](https://github.com/etcd-io/etcd/pull/16241)
- Fix [panic in etcd validate secure endpoints](https://github.com/etcd-io/etcd/pull/16565)
#### Dependencies
- Compile binaries using [go 1.20.10](https://github.com/etcd-io/etcd/pull/16745).
- Upgrade gRPC to 1.58.3 in https://github.com/etcd-io/etcd/pull/16625, https://github.com/etcd-io/etcd/pull/16781 and https://github.com/etcd-io/etcd/pull/16790. Note that gRPC server will reject requests with connection header (refer to https://github.com/grpc/grpc-go/pull/4803).
- Upgrade [bbolt to v1.3.8](https://github.com/etcd-io/etcd/pull/16833)



### aws-cloud-controller-manager [1.24.1-gs10](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.24.1-gs10)

#### Fixed
- Add required values for pss policies.
#### Added
- Add toggle mechanism for `PSPs`.



### cert-exporter [2.8.4](https://github.com/giantswarm/cert-exporter/releases/tag/v2.8.4)

#### Fixed
- Fix daemonset and deployment Kyverno PolicyException.



### cert-manager [3.5.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.5.3)

#### Adds
- adds extra `helm chart` for the `ciliumNetworkPolicies`
#### Changed
- changes the previous `netpols` `helm chart` to be used only for `networkPolicies`
- disables the `startup-api-check` job that waits for the webhookendpoints to become available



### cilium [0.18.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.18.0)

#### Changed
- Upgrade cilium to `1.14.3`.



### observability-bundle [0.10.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.10.1)

#### Fixed
- Extend `prometheus-operator-app` timeout to avoid race condition with VPA causing the app to be stuck in `pending-install` state.



### k8s-dns-node-cache-app [2.5.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.5.0)

#### Changed
- Install PSP resource based on values for PSP deprecation.
- Add PolicyExceptions for PSS compliance.
- Replace build pipeline with ABS.
- Replace testing pipeline with ATS basic test.



### security-bundle [1.4.0](https://github.com/giantswarm/security-bundle/releases/tag/v1.4.0)

#### Changed
- Revert namespace change of `exception-recommender` and `kyverno-policy-operator`.
- Update to `kyverno` (app) version 0.16.2.
- Update to `kyverno-policy-operator` (app) v0.0.5.
- Update to `exception-recommender` (app) v0.0.3.
- Update to `falco` (app) to v0.7.0.



### node-exporter [1.18.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.18.1)

#### Changed
- Make PolicyException namespace configurable.



### chart-operator [3.1.0](https://github.com/giantswarm/chart-operator/releases/tag/v3.1.0)

#### Changed
- Force-disable PSP-related resources when `global.podSecurityStandards.enforced` value is true.



### vertical-pod-autoscaler [4.6.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v4.6.0)

#### Changed
- Revert docker image to `0.14.0`



### aws-ebs-csi-driver [2.28.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.28.0)

#### Added
- Add a job that removes a gp2 storage class for EKS.



### cluster-autoscaler [1.24.3-gs3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3-gs3)

Not found


### coredns [1.19.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.19.1)

#### Changed
- Build App with ABS.
- Add basic tests with ATS.
- ATS: Rework tests. ([#248](https://github.com/giantswarm/coredns-app/pull/248))
- Chart: Fix usage of `name` & `namespace`. ([#249](https://github.com/giantswarm/coredns-app/pull/249))



### prometheus-blackbox-exporter [0.4.0](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.4.0)

#### Added
- Add `global.podSecurityStandards.enforced` value for PSS migration.



### k8s-audit-metrics [0.8.0](https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.8.0)

#### Changed
- Replace condition for PSP CR installation.



