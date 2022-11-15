# :zap: Giant Swarm Release v18.1.0 for AWS :zap:

<< Add description here >>

## Change details


### app-operator [6.4.1](https://github.com/giantswarm/app-operator/releases/tag/v6.4.1)

#### Fixed
- Fix a bug that skips adding the chart-values or chart-secrets entries to the Chart CR when they are only created via extra configs of the App CR



### cluster-operator [5.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v5.3.0)

#### Changed
- Enable IRSA by default on v19+ clusters.



### kubernetes [1.23.14](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.14)

#### API Change
- Make STS available replicas optional again, ([#109241](https://github.com/kubernetes/kubernetes/pull/109241), [@ravisantoshgudimetla](https://github.com/ravisantoshgudimetla)) [SIG API Machinery and Apps]
- Make STS available replicas optional again. ([#113122](https://github.com/kubernetes/kubernetes/pull/113122), [@ashrayjain](https://github.com/ashrayjain)) [SIG Apps]
- Protobuf serialization of metav1.MicroTime timestamps (used in `Lease` and `Event` API objects) has been corrected to truncate to microsecond precision, to match the documented behavior and JSON/YAML serialization. Any existing persisted data is truncated to microsecond when read from etcd. ([#111936](https://github.com/kubernetes/kubernetes/pull/111936), [@haoruan](https://github.com/haoruan)) [SIG API Machinery]
#### Bug or Regression
- Consider only plugin directory and not entire kubelet root when cleaning up mounts ([#112921](https://github.com/kubernetes/kubernetes/pull/112921), [@mattcary](https://github.com/mattcary)) [SIG Storage]
- Etcd: Update to v3.5.5 ([#113100](https://github.com/kubernetes/kubernetes/pull/113100), [@mk46](https://github.com/mk46)) [SIG API Machinery, Cloud Provider, Cluster Lifecycle and Testing]
- Fixed a bug where a change in the `appProtocol` for a Service did not trigger a load balancer update. ([#113033](https://github.com/kubernetes/kubernetes/pull/113033), [@MartinForReal](https://github.com/MartinForReal)) [SIG Cloud Provider and Network]
- Kube-proxy, will restart in case it detects that the Node assigned pod.Spec.PodCIDRs have changed ([#113258](https://github.com/kubernetes/kubernetes/pull/113258), [@code-elinka](https://github.com/code-elinka)) [SIG Network]
- Kubelet no longer reports terminated container metrics from cAdvisor ([#112964](https://github.com/kubernetes/kubernetes/pull/112964), [@bobbypage](https://github.com/bobbypage)) [SIG Node]
- Kubelet: fix GetAllocatableCPUs method in cpumanager ([#113422](https://github.com/kubernetes/kubernetes/pull/113422), [@Garrybest](https://github.com/Garrybest)) [SIG Node]
- Pod logs using --timestamps are not broken up with timestamps anymore. ([#113517](https://github.com/kubernetes/kubernetes/pull/113517), [@rphillips](https://github.com/rphillips)) [SIG Node]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- github.com/stretchr/objx: [v0.2.0 → v0.4.0](https://github.com/stretchr/objx/compare/v0.2.0...v0.4.0)
- github.com/stretchr/testify: [v1.7.0 → v1.8.0](https://github.com/stretchr/testify/compare/v1.7.0...v1.8.0)
- go.uber.org/goleak: v1.1.10 → v1.2.0
- gopkg.in/yaml.v3: 496545a → v3.0.1
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.30 → v0.0.33
#### Removed
_Nothing has changed._



### etcd [3.5.5](https://github.com/etcd-io/etcd/releases/tag/v3.5.5)

#### Deprecations
- Deprecated [SetKeepAlive and SetKeepAlivePeriod in limitListenerConn](https://github.com/etcd-io/etcd/pull/14366).

#### Package `clientv3`
- Fix [do not overwrite authTokenBundle on dial](https://github.com/etcd-io/etcd/pull/14132).
- Fix [IsOptsWithPrefix returns false even if WithPrefix() is included](https://github.com/etcd-io/etcd/pull/14187).

#### etcd server
- [Build official darwin/arm64 artifacts](https://github.com/etcd-io/etcd/pull/14436).
- Add [`etcd --max-concurrent-streams`](https://github.com/etcd-io/etcd/pull/14219) flag to configure the max concurrent streams each client can open at a time, and defaults to math.MaxUint32.
- Add [`etcd --experimental-compact-hash-check-enabled --experimental-compact-hash-check-time`](https://github.com/etcd-io/etcd/issues/14039) flags to support enabling reliable corruption detection on compacted revisions.
- Fix [unexpected error during txn](https://github.com/etcd-io/etcd/issues/14110).
- Fix [lease leak issue due to tokenProvider isn't enabled when restoring auth store from a snapshot](https://github.com/etcd-io/etcd/pull/13205).
- Fix [the race condition between goroutine and channel on the same leases to be revoked](https://github.com/etcd-io/etcd/pull/14087).
- Fix [lessor may continue to schedule checkpoint after stepping down leader role](https://github.com/etcd-io/etcd/pull/14087).
- Fix [Restrict the max size of each WAL entry to the remaining size of the WAL file](https://github.com/etcd-io/etcd/pull/14127).
- Fix [Protect rangePermCache with a RW lock correctly](https://github.com/etcd-io/etcd/pull/14227)
- Fix [memberID equals zero in corruption alarm](https://github.com/etcd-io/etcd/pull/14272)
- Fix [Durability API guarantee broken in single node cluster](https://github.com/etcd-io/etcd/pull/14424)
- Fix [etcd fails to start after performing alarm list operation and then power off/on](https://github.com/etcd-io/etcd/pull/14429)
- Fix [authentication data not loaded on member startup](https://github.com/etcd-io/etcd/pull/14409)

#### etcdctl v3

- Fix [etcdctl move-leader may fail for multiple endpoints](https://github.com/etcd-io/etcd/pull/14434)


#### Other
- [Bump golang.org/x/crypto to latest version](https://github.com/etcd-io/etcd/pull/13996) to address [CVE-2022-27191](https://github.com/advisories/GHSA-8c26-wmh5-6g9v).
- [Bump OpenTelemetry to 1.0.1 and gRPC to 1.41.0](https://github.com/etcd-io/etcd/pull/14312).


### kiam [2.5.1](https://github.com/giantswarm/kiam-app/releases/tag/v2.5.1)

#### Fixed
- Allow `whiteListRouteRegexp` to default to `/latest/*`.



### aws-ebs-csi-driver [2.17.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.17.0)

#### Changed
- Add short names for Volume Snapshot CRDs.
- Update aws-ebs-csi-driver version to `v1.12.1`.
- Update csi-snapshotter version to `v6.0.1`.
- Update csi-resizer version to `v1.4.0`.
- Update csi-node-driver-registrar version to `v2.5.1`.
- Update snapshot-controller version to `v6.1.1`.



### cert-exporter [2.3.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.3.0)

#### Changed
- Update base container image to quay.io/giantswarm/alpine:3.16.2-giantswarm.
- Update go to 1.18.
- Update github.com/giantswarm/k8sclient to v7.0.1.
- Update github.com/hashicorp/vault/api to v1.7.2.
- Update github.com/prometheus/client_golang to v1.13.0.
- Update github.com/spf13/afero to v1.9.2.
- Update k8s.io/api to v0.23.10.
- Update k8s.io/apimachinery to v0.23.10.
- Update k8s.io/client-go to v0.23.10.
#### Added
- Add /etc/kubernetes/pki to --cert-paths flag in DaemonSet deployment.



### cert-manager [2.18.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.18.0)

#### Added
- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in `deployment/cert-manager-cainjector`, `deployment/cert-manager-controller` and `deployment/cert-manager-webhook` if defined in `values.yaml`.
- Support for using `cluster-apps-operator` specific `cluster.proxy` values.



### chart-operator [2.32.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.32.0)

#### Added
- Support for running behind a proxy.
  - `HTTP_PROXY`,`HTTPS_PROXY` and `NO_PROXY` are set as environment variables in `deployment/chart-operator` if defined in `values.yaml`.
- Support for using `cluster-apps-operator` generated `cluster.proxy` values.



### external-dns [2.17.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.17.1)

#### Changed
- Allow using AWS Route53 from any provider [#200](https://github.com/giantswarm/external-dns-app/pull/200)



### kube-state-metrics [1.14.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.14.0)

#### Changed
- Upgrade kube-state-metrics to 2.6.0



