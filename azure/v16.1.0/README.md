# :zap: Giant Swarm Release v16.1.0 for Azure :zap:

This is a security release featuring the latest Kubernetes 1.21 and Flatcar Linux versions as well as an updated version of the Giant Swarm applications. 

## Change details

### azure-operator [5.10.2](https://github.com/giantswarm/azure-operator/releases/tag/v5.10.2)

Upgraded from version 5.10.0

#### Changed
- When looking for the encryption secret, search on all namespaces (to support latest cluster-operator).

#### Fixed
- Consider case when API is down when checking if Master node is upgrading during node pool reconciliation.



### cluster-operator [3.12.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.12.0)

Switched from the legacy branch to the master branch. The [changelog](https://github.com/giantswarm/cluster-operator/blob/master/CHANGELOG.md) is huge, please check details in the project page.



### app-operator [5.3.1](https://github.com/giantswarm/app-operator/releases/tag/v5.3.1)

#### Added
- Support for App CRs with a v prefixed version. This enables Flux to automatically update the version based on its image tag.

#### Changed
- Use dynamic client instead of generated client for watching chart CRs.
- Validate `.spec.kubeConfig.secret.name` in validation resource.




### cert-operator [1.1.0](https://github.com/giantswarm/cert-operator/releases/tag/v1.1.0)

#### Changed
- Adjust helm chart to be used with `config-controller`.
- Replace `jwt-go` with `golang-jwt/jwt`.
- Manage Secrets in the same namespace in which CertConfigs are found.
- Make `expirationThreshold` configurable.



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



### containerlinux [2983.2.0](https://www.flatcar-linux.org/releases/#release-2983.2.0)

Upgraded from version 2905.2.5.

This upgrade provides the solution for a high number of security issues in the Linux Kernel, Containerd and Golang.
Please check details in the [upstream changelog page](https://www.flatcar-linux.org/releases/).


### calico [3.15.5](https://github.com/projectcalico/calico/releases/tag/v3.15.5)

#### Bug fixes
 - Fix that calico/node would fail to set NetworkUnavailable to false for etcd clusters with mismatched nodenames [node #949](https://github.com/projectcalico/node/pull/949) (@caseydavenport)
 - Fixes a bug where IPv6 networks were not handled properly by the failsafe rules [felix #2748](https://github.com/projectcalico/felix/pull/2748) (@mgleung)
 - Fix that, after a netlink read failure, Felix would tight loop reading from a closed channel.  Restart the event poll in that case. [felix #2713](https://github.com/projectcalico/felix/pull/2713) (@fasaxc)
#### Other changes
 - FailsafeInboundHostPorts & FailsafeOutboundHostPorts now support restricting to specific cidrs. New format <protocol>:<net>:<port> [felix #2721](https://github.com/projectcalico/felix/pull/2721) (@mgleung)



### etcd [3.4.18](https://github.com/etcd-io/etcd/releases/tag/v3.4.18)

Upgraded from version 3.4.16.

Please check details in the [upstream changelog page](https://github.com/etcd-io/etcd/blob/main/CHANGELOG-3.4.md).


### kube-state-metrics [1.5.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.5.1)

#### Fixed
- Fix permission to list and watch `leases.coordination.k8s.io`.



### azure-scheduled-events [0.5.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.5.0)

#### Added
- Added basic prometheus exporter.



