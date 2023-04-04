# :zap: Giant Swarm Release v18.3.0 for AWS :zap:

This release contains changes that address several vulnerabilities as well as improvements.

## Change details


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



### aws-operator [14.12.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.12.0)

#### Changed
- Set ENV for nftables in `aws-cni`.
- Improved etcd resiliency and allow customization of --quota-backend-bytes.
- Bump k8s-api-healthz image to 0.2.0.
- When creating a cluster, create the master ASGs in parallel.

#### Fixed
- Don't mark master instance as unhealthy if local etcd instance is unresponsive but the whole etcd cluster is also down.
- Don't mark master instance as unhealthy if local API server instance is unresponsive but the whole API server is also down.



### kubernetes [1.23.17](https://github.com/kubernetes/kubernetes/releases/tag/v1.23.17)

#### Feature
- Kubelet TCP and HTTP probes are more effective using networking resources: conntrack entries, sockets, ... 
  This is achieved by reducing the TIME-WAIT state of the connection to 1 second, instead of the defaults 60 seconds. This allows kubelet to free the socket, and free conntrack entry and ephemeral port associated. ([#115143](https://github.com/kubernetes/kubernetes/pull/115143), [@aojea](https://github.com/aojea)) [SIG Network and Node]
- Kubernetes is now built with Go 1.19.6 ([#115830](https://github.com/kubernetes/kubernetes/pull/115830), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Golang.org/x/net updates to v0.7.0 to fix CVE-2022-41723 ([#115790](https://github.com/kubernetes/kubernetes/pull/115790), [@liggitt](https://github.com/liggitt)) [SIG API Machinery, Auth, CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Security and Storage]
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



### aws-cni [1.12.6](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.12.6)

* Bug - [Fix MTU parameter in egress-v4-cni plugin](https://github.com/aws/amazon-vpc-cni-k8s/pull/2295) (@jdn5126 )
* Documentation - [Fixing the log message to be meaningful](https://github.com/aws/amazon-vpc-cni-k8s/pull/2260) (@rajeeshckr )
* Improvement - [Add bmn-sf1.metal instance support](https://github.com/aws/amazon-vpc-cni-k8s/pull/2286) (@vpineda1996 )
* Improvement - [Support routing to external IPs behind service](https://github.com/aws/amazon-vpc-cni-k8s/pull/2243) (@jdn5126 )
* Improvement - [Use Go 1.19; fix egress-v4-cni MTU parsing, update containerd](https://github.com/aws/amazon-vpc-cni-k8s/pull/2303) (@jdn5126 )
* Improvement - [Added enviroment variable to allow ipamd to manage the ENIs on a non schedulable node](https://github.com/aws/amazon-vpc-cni-k8s/pull/2296) (@rajeeshckr )
* Improvement - [Use GET for IAM Permissions event; update controller-runtime from 0.13.1 to 0.14.4 and client-go from v0.25.5 to v0.26.1](https://github.com/aws/amazon-vpc-cni-k8s/pull/2304) (@jdn5126 )
* Improvement - [Remove old checkpoint migration logic; update containerd version](https://github.com/aws/amazon-vpc-cni-k8s/pull/2307) (@jdn5126 )



### cluster-autoscaler [1.23.1-gs1](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.23.1-gs1)

#### Changed
- Add 'projected' volumes to the PSP.
- Add new-pod-scale-up-delay variable.


### observability-bundle [0.3.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.3.0)

#### Changed
- Add new app dependency mechanism (`app-operator.giantswarm.io/depends-on`) to the prometheus-operator-app and agent so they are not installed until the CRD app is deployed.
- prometheus-operator: drop `apiserver_request_slo_duration_seconds_bucket` metrics from apiserver
- upgrade `prometheus-operator-app` to 4.0.1 and `prometheus-operator-crd` to 4.0.0
- upgrade `prometheus-agent` to 0.3.0 to support chinese registry
#### Added
- Add `promtail-app` v1.0.1 disabled by default.



### coredns [1.14.2](https://github.com/giantswarm/coredns-app/releases/tag/v1.14.2)

#### Changed
- ConfigMap: Add lameduck of 5 seconds to health check ([#191](https://github.com/giantswarm/coredns-app/pull/191)).



### cert-manager [2.20.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.20.3)

#### Added
- Add `node-role.kubernetes.io/control-plane` key to list of tolerations



