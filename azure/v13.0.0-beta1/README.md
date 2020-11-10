# :zap: Giant Swarm Release v13.0.0-beta1 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.0.0-beta1](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0-beta1)

Not found


### app-operator [2.3.3](https://github.com/giantswarm/app-operator/releases/tag/v2.3.3)

#### Added
- Delete chart-operator helm release and chart CR so it can be re-installed.



### chart-operator [2.3.5](https://github.com/giantswarm/chart-operator/releases/tag/v2.3.5)

#### Fixed
- Stop repeating helm upgrade for the failed helm release.



### kubernetes [1.18.10](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.10)

#### Design
- Prevent logging of docker config contents if file is malformed ([#95347](https://github.com/kubernetes/kubernetes/pull/95347), [@sfowl](https://github.com/sfowl)) [SIG Auth and Node]
#### Bug or Regression
- Do not fail sorting empty elements. ([#94666](https://github.com/kubernetes/kubernetes/pull/94666), [@soltysh](https://github.com/soltysh)) [SIG CLI]
- Ensure getPrimaryInterfaceID not panic when network interfaces for Azure VMSS are null ([#94801](https://github.com/kubernetes/kubernetes/pull/94801), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Fix bug where loadbalancer deletion gets stuck because of missing resource group &#35;75198 ([#93962](https://github.com/kubernetes/kubernetes/pull/93962), [@phiphi282](https://github.com/phiphi282)) [SIG Cloud Provider]
- Fix detach azure disk issue when vm not exist ([#95177](https://github.com/kubernetes/kubernetes/pull/95177), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix etcd_object_counts metric reported by kube-apiserver ([#94818](https://github.com/kubernetes/kubernetes/pull/94818), [@tkashem](https://github.com/tkashem)) [SIG API Machinery]
- Fix network_programming_latency metric reporting for Endpoints/EndpointSlice deletions, where we don't have correct timestamp ([#95363](https://github.com/kubernetes/kubernetes/pull/95363), [@wojtek-t](https://github.com/wojtek-t)) [SIG Network and Scalability]
- Fix scheduler cache snapshot when a Node is deleted before its Pods ([#95154](https://github.com/kubernetes/kubernetes/pull/95154), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- Fix the `cloudprovider_azure_api_request_duration_seconds` metric buckets to correctly capture the latency metrics. Previously, the majority of the calls would fall in the "+Inf" bucket. ([#95375](https://github.com/kubernetes/kubernetes/pull/95375), [@marwanad](https://github.com/marwanad)) [SIG Cloud Provider and Instrumentation]
- Fix: azure disk resize error if source does not exist ([#93011](https://github.com/kubernetes/kubernetes/pull/93011), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fix: detach azure disk broken on Azure Stack ([#94885](https://github.com/kubernetes/kubernetes/pull/94885), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a bug where improper storage and comparison of endpoints led to excessive API traffic from the endpoints controller ([#94934](https://github.com/kubernetes/kubernetes/pull/94934), [@damemi](https://github.com/damemi)) [SIG Apps, Network and Testing]
- Gracefully delete nodes when their parent scale set went missing ([#95289](https://github.com/kubernetes/kubernetes/pull/95289), [@bpineau](https://github.com/bpineau)) [SIG Cloud Provider]
- Kubeadm: warn but do not error out on missing "ca.key" files for root CA, front-proxy CA and etcd CA, during "kubeadm join --control-plane" if the user has provided all certificates, keys and kubeconfig files which require signing with the given CA keys. ([#94988](https://github.com/kubernetes/kubernetes/pull/94988), [@neolit123](https://github.com/neolit123)) [SIG Cluster Lifecycle]
#### Other (Cleanup or Flake)
- Masks ceph RBD adminSecrets in logs when logLevel >= 4 ([#95245](https://github.com/kubernetes/kubernetes/pull/95245), [@sfowl](https://github.com/sfowl)) [SIG Storage]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
_Nothing has changed._
#### Removed
_Nothing has changed._



### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.18)

Not found


### containerlinux [2605.7.0](https://www.flatcar-linux.org/releases/#release-2605.7.0)

Security fixes:

- Linux - [CVE-2020-25645](https://nvd.nist.gov/vuln/detail/CVE-2020-25645), [CVE-2020-25643](https://nvd.nist.gov/vuln/detail/CVE-2020-25643), [CVE-2020-25211](https://nvd.nist.gov/vuln/detail/CVE-2020-25211)

Bug fixes:

- Ensured that the `/etc/coreos` to `/etc/flatcar` symlink always exists, relevant for the Container Linux Config transpiler (ct) when specifying directives for `update:` or `locksmith:` while also reformatting the rootfs ([baselayout PR#7](https://github.com/flatcar-linux/baselayout/pull/7))

Updates:

- Linux [5.4.72](https://lwn.net/Articles/834537/)


### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.4.12...v3.4.13) and [v3.4 upgrade guide](https://github.com/etcd-io/etcd/blob/master/Documentation/upgrades/upgrade_3_4.md) for any breaking changes.
#### Security
- A [log warning](https://github.com/etcd-io/etcd/pull/12242) is added when etcd use any existing directory that has a permission different than 700 on Linux and 777 on Windows.



### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)

#### Other changes
 - Add FelixConfiguration parameters to explicitly allow encapsulated packets from workloads. [libcalico-go #1302](https://github.com/projectcalico/libcalico-go/pull/1302) (@doublek)
 - Respect explicit configuration for drop rules for encapsulated packets originating from workloads. [felix #2487](https://github.com/projectcalico/felix/pull/2487) (@doublek)



