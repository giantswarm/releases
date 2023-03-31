# :zap: Giant Swarm Release v18.3.0 for AWS :zap:

-This release contains changes that improves the stability of cluster upgrades

## Change details


### aws-operator [14.8.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.8.1)

Not found


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



