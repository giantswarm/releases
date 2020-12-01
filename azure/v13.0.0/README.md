# :zap: Giant Swarm Release v13.0.0 for Azure :zap:

<< Add description here >>

## Change details


### azure-operator [5.0.0](https://github.com/giantswarm/aws-operator/releases/tag/v5.0.0)

Not found


### app-operator [2.7.0](https://github.com/giantswarm/app-operator/releases/tag/v2.7.0)

#### Added
- Secure the webhook with token value from control plane catalog.



### cert-operator [0.1.0](https://github.com/giantswarm/cert-operator/releases/tag/v0.1.0)

#### Changed
- No longer ensure CertConfig CRD.
- Use architect-orb to release cert-operator.
#### Added
- First release.



### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.18)

Not found


### kubernetes [1.18.12](https://github.com/kubernetes/kubernetes/releases/tag/v1.18.12)




### containerlinux [2605.8.0](https://www.flatcar-linux.org/releases/#release-2605.8.0)

Security fixes:

*   Linux -  [CVE-2020-27673](https://nvd.nist.gov/vuln/detail/CVE-2020-27673), [CVE-2020-27675](https://nvd.nist.gov/vuln/detail/CVE-2020-27675)

Bug fixes:

*   network: Restore KeepConfiguration=dhcp-on-stop ([kinvolk/init#30](https://github.com/kinvolk/init/pull/30))
*   systemd-stable-245.8: ingest latest fixes on top of upstream release ([#1](https://github.com/kinvolk/systemd/commit/261680bc0ea61777ac22ea1c42b0d728ec52ae14), [#2](https://github.com/kinvolk/systemd/commit/b2b382820bcfc166d048b85aadd90f5cf71c7a4a), [#3](https://github.com/kinvolk/systemd/commit/711ca814c9f2e81d3d25ebbed0b837b7d4fbbeda))

Updates:

*   Linux ([5.4.77](https://lwn.net/Articles/836795/))
*   systemd ([245.8](https://github.com/systemd/systemd-stable/releases/tag/v245.8))


### coredns [1.6.5](https://github.com/giantswarm/coredns-app/releases/tag/v1.6.5)

Not found


### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)

#### Other changes
 - Add FelixConfiguration parameters to explicitly allow encapsulated packets from workloads. [libcalico-go #1302](https://github.com/projectcalico/libcalico-go/pull/1302) (@doublek)
 - Respect explicit configuration for drop rules for encapsulated packets originating from workloads. [felix #2487](https://github.com/projectcalico/felix/pull/2487) (@doublek)



### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.4.12...v3.4.13) and [v3.4 upgrade guide](https://github.com/etcd-io/etcd/blob/master/Documentation/upgrades/upgrade_3_4.md) for any breaking changes.
#### Security
- A [log warning](https://github.com/etcd-io/etcd/pull/12242) is added when etcd use any existing directory that has a permission different than 700 on Linux and 777 on Windows.
#### Go
- Compile with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).
<hr>



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



### coredns [1.2.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.2.0)

#### Changed
- Apply a readiness probe
- Increase the liveness probe failure threshold from 5 failures to 7 failures



### external-dns [1.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v1.5.0)

#### Changed
- Upgrade upstream external-dns from v0.7.3 to [v0.7.4](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.4).



### kube-state-metrics [1.3.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.3.0)

#### Changed
- Change the Kubernetes Deployment name to include the app version.



### app-operator [2.7.0](https://github.com/giantswarm/app-operator/releases/tag/v2.7.0)

#### Added
- Secure the webhook with token value from control plane catalog.



### node-exporter [1.7.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.7.0)

#### Changed
- Change the Kubernetes Daemonset name to include the app version.



### metrics-server [1.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.1.0)

#### Changed
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.



### net-exporter [1.9.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.2)

#### Changed
- Updated backward incompatible Kubernetes dependencies to v1.18.5.
#### Fixed
- Fixed indentation problem with the daemonset template.



