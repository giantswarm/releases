# :zap: Giant Swarm Release v15.1.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.20.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.20.9)

#### Feature
- Kubernetes 1.20.x is now built using Go 1.15.14 ([#103677](https://github.com/kubernetes/kubernetes/pull/103677), [@puerco](https://github.com/puerco)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
#### Bug or Regression
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102999](https://github.com/kubernetes/kubernetes/pull/102999), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/structured-merge-diff/v4: v4.0.3 → v4.1.2
#### Removed
_Nothing has changed._



### etcd [3.14.16](https://github.com/etcd-io/etcd/releases/tag/v3.14.16)

Not found


### calico [3.15.5](https://github.com/projectcalico/calico/releases/tag/v3.15.5)

#### Bug fixes
 - Fix that calico/node would fail to set NetworkUnavailable to false for etcd clusters with mismatched nodenames [node #949](https://github.com/projectcalico/node/pull/949) (@caseydavenport)
 - Fixes a bug where IPv6 networks were not handled properly by the failsafe rules [felix #2748](https://github.com/projectcalico/felix/pull/2748) (@mgleung)
 - Fix that, after a netlink read failure, Felix would tight loop reading from a closed channel.  Restart the event poll in that case. [felix #2713](https://github.com/projectcalico/felix/pull/2713) (@fasaxc)
#### Other changes
 - FailsafeInboundHostPorts & FailsafeOutboundHostPorts now support restricting to specific cidrs. New format <protocol>:<net>:<port> [felix #2721](https://github.com/projectcalico/felix/pull/2721) (@mgleung)



### kiam [1.7.1](https://github.com/giantswarm/kiam-app/releases/tag/v1.7.1)

#### Changed
- Set docker.io as the default registry



### cert-manager [1.4.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v1.4.1)

Not found


### coredns [1.6.9](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.9)

Not found


### kube-state-metrics [1.9.8](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.9.8)

Not found


### metrics-server [0.4.4](https://github.com/giantswarm/metrics-server-app/releases/tag/v0.4.4)

Not found


### net-exporter [1.10.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.2)

#### Changed
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.



### node-exporter [1.2.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.2.0)

#### Changed
- Change Priority Class to `system-node-critical`



