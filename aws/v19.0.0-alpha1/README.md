# :zap: Giant Swarm Release v19.0.0-alpha1 for AWS :zap:

<< Add description here >>

## Change details


### cert-operator [3.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v3.0.1)

#### Fixed
- Allow running unique and non unique cert-operators in the same namespace.



### aws-operator [14.9.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.9.0)

#### Changed
- Bump k8s-api-healthz image to 0.2.0.
#### Fixed
- Don't mark master instance as unhealthy if local etcd instance is unresponsive but the whole etcd cluster is also down.
- Don't mark master instance as unhealthy if local API server instance is unresponsive but the whole API server is also down.



### containerlinux [3374.2.3](https://www.flatcar-linux.org/releases/#release-3374.2.3)

_Changes since **Stable 3374.2.2**_

#### Security fixes:

- Linux ([CVE-2022-3169](https://nvd.nist.gov/vuln/detail/CVE-2022-3169), [CVE-2022-3344](https://nvd.nist.gov/vuln/detail/CVE-2022-3344), [CVE-2022-3424](https://nvd.nist.gov/vuln/detail/CVE-2022-3424), [CVE-2022-3521](https://nvd.nist.gov/vuln/detail/CVE-2022-3521), [CVE-2022-3534](https://nvd.nist.gov/vuln/detail/CVE-2022-3534), [CVE-2022-3545](https://nvd.nist.gov/vuln/detail/CVE-2022-3545), [CVE-2022-3643](https://nvd.nist.gov/vuln/detail/CVE-2022-3643), [CVE-2022-4378](https://nvd.nist.gov/vuln/detail/CVE-2022-4378), [CVE-2022-45869](https://nvd.nist.gov/vuln/detail/CVE-2022-45869), [CVE-2022-45934](https://nvd.nist.gov/vuln/detail/CVE-2022-45934), [CVE-2022-47518](https://nvd.nist.gov/vuln/detail/CVE-2022-47518), [CVE-2022-47519](https://nvd.nist.gov/vuln/detail/CVE-2022-47519), [CVE-2022-47520](https://nvd.nist.gov/vuln/detail/CVE-2022-47520), [CVE-2022-47521](https://nvd.nist.gov/vuln/detail/CVE-2022-47521))
- git ([CVE-2022-23521](https://nvd.nist.gov/vuln/detail/CVE-2022-23521), [CVE-2022-41903](https://nvd.nist.gov/vuln/detail/CVE-2022-41903))

#### Bug fixes:

- Fix "ext4 deadlock under heavy I/O load" kernel issue. The patch for this is included provisionally while we stay with Kernel 5.15.86. ([Flatcar#847](https://github.com/flatcar/Flatcar/issues/847), [coreos-overlay#2402](https://github.com/flatcar/coreos-overlay/pull/2402))

#### Changes:


#### Updates:

- Linux ([5.15.86](https://lwn.net/Articles/918808) (includes [5.15.85](https://lwn.net/Articles/918329), [5.15.84](https://lwn.net/Articles/918206), [5.15.83](https://lwn.net/Articles/917896), [5.15.82](https://lwn.net/Articles/917400), [5.15.81](https://lwn.net/Articles/916763), [5.15.80](https://lwn.net/Articles/916003)))
- git ([2.37.5](https://github.com/git/git/blob/v2.37.5/Documentation/RelNotes/2.37.5.txt))



### kubernetes [1.24.11](https://github.com/kubernetes/kubernetes/releases/tag/v1.24.11)

#### Feature
- Kubelet TCP and HTTP probes are more effective using networking resources: conntrack entries, sockets, ... 
  This is achieved by reducing the TIME-WAIT state of the connection to 1 second, instead of the defaults 60 seconds. This allows kubelet to free the socket, and free conntrack entry and ephemeral port associated. ([#115143](https://github.com/kubernetes/kubernetes/pull/115143), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Kubernetes is now built with Go 1.19.6 ([#115831](https://github.com/kubernetes/kubernetes/pull/115831), [@cpanato](https://github.com/cpanato)) [SIG Release and Testing]
#### Bug or Regression
- Fix the regression that introduced 34s timeout for DELETECOLLECTION calls ([#115482](https://github.com/kubernetes/kubernetes/pull/115482), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fixed bug which caused the status of Indexed Jobs to only be updated when there are newly completed indexes. The completed indexes are now updated if the .status.completedIndexes has values outside of the [0, .spec.completions> range ([#115457](https://github.com/kubernetes/kubernetes/pull/115457), [@danielvegamyhre](https://github.com/danielvegamyhre)) [SIG Apps]
- Golang.org/x/net updates to v0.7.0 to fix CVE-2022-41723 ([#115789](https://github.com/kubernetes/kubernetes/pull/115789), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Architecture, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Security and Storage]
- The Kubernetes API server now correctly detects and closes existing TLS connections when its client certificate file for kubelet authentication has been rotated. ([#115580](https://github.com/kubernetes/kubernetes/pull/115580), [@enj](https://github.com/enj)) [SIG API Machinery, Node and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- golang.org/x/net: 1e63c2f → v0.7.0
- golang.org/x/sys: v0.3.0 → v0.5.0
- golang.org/x/term: v0.3.0 → v0.5.0
- golang.org/x/text: v0.5.0 → v0.7.0
#### Removed
_Nothing has changed._



### etcd [3.5.7](https://github.com/etcd-io/etcd/releases/tag/v3.5.7)

#### etcd server
- Fix [Remove memberID from data corrupt alarm](https://github.com/etcd-io/etcd/pull/14852).
- Fix [Allow non mutating requests pass through quotaKVServer when NOSPACE](https://github.com/etcd-io/etcd/pull/14884).
- Fix [nil pointer panic for readonly txn due to nil response](https://github.com/etcd-io/etcd/pull/14899).
- Fix [The last record which was partially synced to disk isn't automatically repaired](https://github.com/etcd-io/etcd/pull/15069).
- Fix [etcdserver might promote a non-started learner](https://github.com/etcd-io/etcd/pull/15096).
#### Package `clientv3`
- Reverted the fix to [auth invalid token and old revision errors in watch](https://github.com/etcd-io/etcd/pull/14995).
#### Security
- Use [distroless base image](https://github.com/etcd-io/etcd/pull/15016) to address critical Vulnerabilities.
- Updated [base image from base-debian11 to static-debian11 and removed dependency on busybox](https://github.com/etcd-io/etcd/pull/15037).
- Bumped [some dependencies](https://github.com/etcd-io/etcd/pull/15018) to address some HIGH Vulnerabilities.
#### Go
- Require [Go 1.17+](https://github.com/etcd-io/etcd/pull/15019).
- Compile with [Go 1.17+](https://go.dev/doc/devel/release#go1.17)



### aws-cloud-controller-manager [1.24.1-gs1](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.24.1-gs1)

#### Changed
- Bump to upstream version 1.24.1.



### cert-manager [2.20.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.20.0)

#### Added
- Adds support for DNS01 challenge via AWS Route53 ([#284](https://github.com/giantswarm/cert-manager-app/pull/292))



### chart-operator [2.34.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.34.0)

#### Changed
- Selecting private Helm client on demand for some operations.



### cilium [0.8.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.8.0)

#### Changed
- Bump all manifests to upstream version 1.13.
- Enable Hubble
- Enable Monitoring for Agent, Operator and Hubble



### cluster-autoscaler [1.24.0-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.0-gs1)

#### Changed
- Update cluster-autoscaler to version `1.24.0`.



### coredns [1.14.2](https://github.com/giantswarm/coredns-app/releases/tag/v1.14.2)

#### Changed
- ConfigMap: Add lameduck of 5 seconds to health check ([#191](https://github.com/giantswarm/coredns-app/pull/191)).



### external-dns [2.33.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.33.0)

#### Added
- Add support to run in `hostNetwork` (primary used in `CAPZ` based management clusters)



### vertical-pod-autoscaler [3.2.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.2.0)

#### Added
- Add VPA CR to VPA.



### vertical-pod-autoscaler-crd [2.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.0)

#### Changed
- Synced with new upstream repo
#### Changed
- Remove `push-to-app-collection` jobs for onprem providers since this app became a part of default apps bundle.
#### Added
- Add icon to Chart.yml for use in happa



### observability-bundle [0.2.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.2.0)

#### Changed
- Upgrade `prometheus-agent` from to 0.1.7 to 0.2.0.



### k8s-dns-node-cache-app [2.0.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.0.0)

#### Breaking change, application only compatible with Cilium from this release onwards.
#### Changed
- Adapt application to be deployed on Cilium.



### prometheus-blackbox-exporter [0.2.2](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.2.2)

#### Changed
- Add IRSA webhook module
- Fix k8s CA for API probe
- Use IPV4 by default



### cilium-servicemonitors [0.0.2](https://github.com/giantswarm/cilium-servicemonitors-app/releases/tag/v0.0.2)

#### Changed
- Add labels to servicemonitors



### irsa-servicemonitors [0.0.1](https://github.com/giantswarm/irsa-servicemonitors-app/releases/tag/v0.0.1)

#### Added

- First release


