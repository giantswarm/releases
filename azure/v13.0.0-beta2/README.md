# :zap: Giant Swarm Release v13.0.0-beta2 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.0.0-beta5](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0-beta5)

Not found


### app-operator [2.7.0](https://github.com/giantswarm/app-operator/releases/tag/v2.7.0)

#### Added
- Secure the webhook with token value from control plane catalog.



### chart-operator [2.3.5](https://github.com/giantswarm/chart-operator/releases/tag/v2.3.5)

#### Fixed
- Stop repeating helm upgrade for the failed helm release.



### kubernetes [1.18.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.12)




### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.18)




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



### cert-exporter [1.3.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.3.0)

#### Added
- Add Network Policy.
#### Changed
- Remove `hostNetwork` and `hostPID` capabilities.



### chart-operator [2.5.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.5.0)

#### Added
- Validate the cache in helmclient to avoid state requests when pulling tarballs.
- Call status webhook with token values.
#### Fixed
- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.



### external-dns [1.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v1.5.0)

#### Changed
- Upgrade upstream external-dns from v0.7.3 to [v0.7.4](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.4).



### kube-state-metrics [1.2.1](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.2.1)

#### Fixed
- Support deployment of `kube-state-metrics-app` on chinese installations.



### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.2)

#### Changed
- Updated backward incompatible Kubernetes dependencies to v1.18.5.
#### Fixed
- Fixed indentation problem with the daemonset template.



### node-exporter [1.6.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.6.0)

#### Changed
- Disable the diskstats collector on azure.



