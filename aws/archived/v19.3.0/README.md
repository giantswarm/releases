# :zap: Giant Swarm Release v19.3.0 for AWS :zap:

This release marks the final step for migration away from Pod Security Policies (PSP) in favor of Pod Security Standards (PSS). The `security-bundle` that was added as a default application in the `19.1.0` release with `kyverno` in PSS `audit` mode will now switch to `enforce` mode, meaning PSS policies will be actively enforced and any configured PolicyExceptions will be evaluated and applied.

The `PodSecurityPolicy` type still exists in this release. However, PSPs have been removed from all Giant Swarm applications deployed with this release in preparation for the upcoming upgrade to the Kubernetes v1.25, where PSPs will no longer exist. Security policy for Giant Swarm applications is now enforced by PSS in order to ensure the security of these applications.

Our docs offer additional information about [Pod Security Standards](https://docs.giantswarm.io/advanced/security-policy-enforcement/) as well as a [PSS migration guide](https://docs.giantswarm.io/advanced/security/security-policy-enforcement/cluster-admin-guide/) for cluster administrators.

> **WARNING:** Before upgrading to this release, it is required to upgrade ALL Managed Applications provided by Giant Swarm to the latest version supporting PSS. Please reach out to your Account Engineer to validate the versions of applications that are required. 

> **WARNING:** After upgrading to `19.3.0`, it is highly advised to begin removal of all PSPs from the cluster. Kubernetes `v1.25` removes the Pod Security Policy resource from the API, meaning workloads (like Helm charts) which still contain PSPs will fail to install after the upcoming Giant Swarm `v20` release.

## Change details


### aws-operator [14.24.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.24.1)

#### Added
- Add `global.podSecurityStandards.enforced` value for PSS migration.
- Emit event when an unhealthy node is terminated.
- Bump `badnodedetector` to be able to use `node-problem-detector` app for unhealthy node termination.
- Add a additional IAM permission for `cluster-autoscaler`.
#### Changed
- Bump k8scc to disable PSPs in preparation for switch to PSS.
- Disable cluster autoscaler during rollouts of node pool ASGs.
- Bump etcd-cluster-migrator to v1.2.0 to disable PSPs in preparation for switch to PSS



### cluster-operator [5.10.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.10.0)

#### Changed
- Automatically set `global.podSecurityStandards.enforced` to `true` for clusters >= 19.3.0.



### containerlinux [3602.2.2](https://www.flatcar-linux.org/releases/#release-3602.2.2)

:warning: From Alpha 3794.0.0 Torcx has been removed - please assert that you don't rely on specific Torcx mechanism but now use systemd-sysext. See [here](https://www.flatcar.org/docs/latest/provisioning/sysext/) for more information.


 _Changes since **Stable 3602.2.1**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-46813](https://nvd.nist.gov/vuln/detail/CVE-2023-46813), [CVE-2023-5178](https://nvd.nist.gov/vuln/detail/CVE-2023-5178), [CVE-2023-5717](https://nvd.nist.gov/vuln/detail/CVE-2023-5717))
 

 #### Changes:
 
 - Brightbox: The regular OpenStack image should now be used, it includes Afterburn for instance metadata attributes
 - OpenStack: An uncompressed image is provided for simpler import (since the images use qcow2 inline compression, there is no benefit in using the `.gz` or `.bz2` images)
 - linux kernel: added zstd support for squashfs kernel module ([scripts#1297](https://github.com/flatcar/scripts/pull/1297))
 
 #### Updates:
 
 - Linux ([5.15.138](https://lwn.net/Articles/950714) (includes [5.15.137](https://lwn.net/Articles/948818)))


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



### app-operator [6.10.1](https://github.com/giantswarm/app-operator/releases/tag/v6.10.1)

#### Fixed
- Add policy exception so that controller can be deployed in bootstrap mode (uses host network)



### chart-operator [3.1.0](https://github.com/giantswarm/chart-operator/releases/tag/v3.1.0)

#### Changed
- Force-disable PSP-related resources when `global.podSecurityStandards.enforced` value is true.



### coredns [1.19.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.19.1)

#### Changed
- Build App with ABS.
- Add basic tests with ATS.
- ATS: Rework tests. ([#248](https://github.com/giantswarm/coredns-app/pull/248))
- Chart: Fix usage of `name` & `namespace`. ([#249](https://github.com/giantswarm/coredns-app/pull/249))



### node-exporter [1.18.1](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.18.1)

#### Changed
- Make PolicyException namespace configurable.



### cert-exporter [2.8.4](https://github.com/giantswarm/cert-exporter/releases/tag/v2.8.4)

#### Fixed
- Fix daemonset and deployment Kyverno PolicyException.



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



### cilium [0.17.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.17.0)

#### Changed
- Generate cilium chart using our fork and `vendir`.



### cluster-autoscaler [1.24.3-gs3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3-gs3)

#### Added

- Kyverno policy exceptions for cluster-autoscaler.


### vertical-pod-autoscaler [4.6.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v4.6.0)

#### Changed
- Revert docker image to `0.14.0`



### prometheus-blackbox-exporter [0.4.0](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.4.0)

#### Added
- Add `global.podSecurityStandards.enforced` value for PSS migration.



### aws-cloud-controller-manager [1.24.1-gs10](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.24.1-gs10)

#### Fixed
- Add required values for pss policies.
#### Added
- Add toggle mechanism for `PSPs`.



### aws-ebs-csi-driver [2.28.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.28.0)

#### Added
- Add a job that removes a gp2 storage class for EKS.



### cert-manager [2.25.2](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.25.2)

### Added
- Added support for PSS resolving issue on upgrade to newer v3+ releases

### Changed
- Changed conditional for PSPs to `{{- if not .global.podSecurityStandards.enforced }}`

### observability-bundle [0.10.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.10.1)

#### Fixed
- Extend `prometheus-operator-app` timeout to avoid race condition with VPA causing the app to be stuck in `pending-install` state.



### k8s-audit-metrics [0.8.0](https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.8.0)

#### Changed
- Replace condition for PSP CR installation.



