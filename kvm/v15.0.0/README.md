# :zap: Giant Swarm Release v15.0.0 for KVM :zap:

This is the first release for KVM with Kubernetes 1.20 and Calico 3.19.

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



### calico [3.19.1](https://github.com/projectcalico/calico/releases/tag/v3.19.1)

#### Bug fixes
- Fix issue with serviceaccount names larger than 63 characters. [libcalico-go #1423](https://github.com/projectcalico/libcalico-go/pull/1423) (@caseydavenport)



### kvm-operator [update_calico](https://github.com/giantswarm/kvm-operator/releases/tag/vupdate_calico)

Not found


### cert-exporter [1.7.1](https://github.com/giantswarm/cert-exporter/releases/tag/v1.7.1)

#### Fixed
- Fix configuration version in `Chart.yaml`.



### chart-operator [2.18.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.18.0)

#### Added
- Add releasemaxhistory resource which ensures we retry at a reduced rate when
there are repeated failed upgrades.
#### Changed
- Upgrade Helm release when failed even if version or values have not changed
to handle situations like failed webhooks where we should retry.



### coredns [1.6.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.0)

- Make `targetCPUUtilizationPercentage` in HPA configurable.



### net-exporter [1.10.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.2)

#### Changed
- Allow to customize dns service.
- Only check pod existence on dial errors. Check pod deletion directly by IP instead of listing pods and searching.



