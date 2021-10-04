# :zap: Giant Swarm Release v16.0.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [10.9.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.9.1)

Not found


### containerlinux [2905.2.3](https://www.flatcar-linux.org/releases/#release-2905.2.3)

New **Stable** release **2905.2.3**

_Changes since **Stable 2905.2.2**_

**Security fixes**



* Linux ([CVE-2021-3653](https://nvd.nist.gov/vuln/detail/CVE-2021-3653), [CVE-2021-3656](https://nvd.nist.gov/vuln/detail/CVE-2021-3656), [CVE-2021-38166](https://nvd.nist.gov/vuln/detail/CVE-2021-38166)) 
* openssl ([CVE-2021-3711](https://nvd.nist.gov/vuln/detail/CVE-2021-3711), [CVE-2021-3712](https://nvd.nist.gov/vuln/detail/CVE-2021-3712))

**Bug Fixes**



* Re-enabled kernel config FS_ENCRYPTION ([coreos-overlay#1212](https://github.com/kinvolk/coreos-overlay/pull/1212/))
* Fixed Perl in dev-container ([coreos-overlay#1238](https://github.com/kinvolk/coreos-overlay/pull/1238))

**Updates**



* Linux ([5.10.61](https://lwn.net/Articles/867497/))
* openssl ([1.1.1l](https://mta.openssl.org/pipermail/openssl-announce/2021-August/000206.html))


### kubernetes [1.21.5](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.5)

#### Feature
- Kubernetes is now built with Golang 1.16.8 ([#104906](https://github.com/kubernetes/kubernetes/pull/104906), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
#### Bug or Regression
- Fix NodeAuthenticator tests in dualstack ([#104840](https://github.com/kubernetes/kubernetes/pull/104840), [@ardaguclu](https://github.com/ardaguclu)) [SIG Auth and Testing]
- Fix: skip case sensitivity when checking Azure NSG rules
  fix: ensure InstanceShutdownByProviderID return false for creating Azure VMs ([#104447](https://github.com/kubernetes/kubernetes/pull/104447), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fixed occasional pod cgroup freeze when using cgroup v1 and systemd driver.
  Fixed "failed to create container ... unit already exists" when using cgroup v1 and systemd driver. ([#104530](https://github.com/kubernetes/kubernetes/pull/104530), [@kolyshkin](https://github.com/kolyshkin)) [SIG CLI, Cloud Provider, Cluster Lifecycle, Instrumentation, Node, Storage and Testing]
- Kube-proxy: delete stale conntrack UDP entries for loadbalancer ingress IP. ([#104151](https://github.com/kubernetes/kubernetes/pull/104151), [@aojea](https://github.com/aojea)) [SIG Network]
- Metrics changes: Fix exposed buckets of `scheduler_volume_scheduling_duration_seconds_bucket` metric ([#100720](https://github.com/kubernetes/kubernetes/pull/100720), [@dntosas](https://github.com/dntosas)) [SIG Apps, Instrumentation, Scheduling and Storage]
- Pass additional flags to subpath mount to avoid flakes in certain conditions ([#104347](https://github.com/kubernetes/kubernetes/pull/104347), [@mauriciopoppe](https://github.com/mauriciopoppe)) [SIG Storage]
- When using `kubectl replace` (or the equivalent API call) on a Service, the caller no longer needs to do a read-modify-write cycle to fetch the allocated values for `.spec.clusterIP` and `.spec.ports[].nodePort`.  Instead the API server will automatically carry these forward from the original object when the new object does not specify them. ([#104673](https://github.com/kubernetes/kubernetes/pull/104673), [@thockin](https://github.com/thockin)) [SIG Network]
#### Other (Cleanup or Flake)
- Kube-apiserver: sets an upper-bound on the lifetime of idle keep-alive connections and time to read the headers of incoming requests ([#103958](https://github.com/kubernetes/kubernetes/pull/103958), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Node]
#### Dependencies
#### Added
- github.com/bits-and-blooms/bitset: [v1.2.0](https://github.com/bits-and-blooms/bitset/tree/v1.2.0)
#### Changed
- github.com/cilium/ebpf: [v0.5.0 → v0.6.2](https://github.com/cilium/ebpf/compare/v0.5.0...v0.6.2)
- github.com/coreos/go-systemd/v22: [v22.3.1 → v22.3.2](https://github.com/coreos/go-systemd/v22/compare/v22.3.1...v22.3.2)
- github.com/golang/protobuf: [v1.4.3 → v1.5.0](https://github.com/golang/protobuf/compare/v1.4.3...v1.5.0)
- github.com/google/go-cmp: [v0.5.4 → v0.5.5](https://github.com/google/go-cmp/compare/v0.5.4...v0.5.5)
- github.com/opencontainers/runc: [v1.0.0-rc95 → v1.0.2](https://github.com/opencontainers/runc/compare/v1.0.0-rc95...v1.0.2)
- github.com/opencontainers/selinux: [v1.8.0 → v1.8.2](https://github.com/opencontainers/selinux/compare/v1.8.0...v1.8.2)
- github.com/sirupsen/logrus: [v1.7.0 → v1.8.1](https://github.com/sirupsen/logrus/compare/v1.7.0...v1.8.1)
- google.golang.org/protobuf: v1.25.0 → v1.26.0
#### Removed
- github.com/willf/bitset: [v1.1.11](https://github.com/willf/bitset/tree/v1.1.11)



### app-operator [5.2.0](https://github.com/giantswarm/app-operator/releases/tag/v5.2.0)

#### Changed
- Reject App CRs with version labels with the legacy `1.0.0` value.
- Validate `.spec.catalog` using Catalog CRs instead of AppCatalog CRs.



### aws-cni [1.9.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.9.0)

* Enhancement - [EC2 sdk model override](https://github.com/aws/amazon-vpc-cni-k8s/pull/1508) (#1508, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Prefix Delegation feature support](https://github.com/aws/amazon-vpc-cni-k8s/pull/1516) (#1516, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Header formatting for env variable](https://github.com/aws/amazon-vpc-cni-k8s/pull/1522) (#1522, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [non-nitro instances init issues](https://github.com/aws/amazon-vpc-cni-k8s/pull/1527) (#1527, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Add metrics for total prefix count and ips used per cidr](https://github.com/aws/amazon-vpc-cni-k8s/pull/1530) (#1530, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Update documentation for PD](https://github.com/aws/amazon-vpc-cni-k8s/pull/1540) (#1540, [@jayanthvn](https://github.com/jayanthvn))
* Enhancement - [Update SDK Go version](https://github.com/aws/amazon-vpc-cni-k8s/pull/1544) (#1544, [@jayanthvn](https://github.com/jayanthvn))



### cluster-operator [3.10.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.10.0)

#### Changed
- Introducing `v1alpha3` CR's.



### external-dns [2.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.5.0)

#### Changed
- Upgrade upstream external-dns from v0.8.0 to [v0.9.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.9.0). The new release brings a lot of smaller improvements and bug fixes.



### cert-exporter [1.8.0](https://github.com/giantswarm/cert-exporter/releases/tag/v1.8.0)

#### Added
- Add new `cert_exporter_certificate_cr_not_after` metric. This metric exports the `status.notAfter` field of cert-manager `Certificate` CR.
#### Changed
- Remove static certificate source label from `cert_exporter_secret_not_after` (static value `secret`) and `cert_exporter_not_after` (static value `file`) metrics.



### cert-manager [2.11.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.11.0)

#### Changed
- Upgrade to upstream `v1.5.4` ([#191](https://github.com/giantswarm/cert-manager-app/pull/191)).
- Add metadata to enable metrics scraping ([#181](https://github.com/giantswarm/cert-manager-app/pull/181)).
- Fix startubjob PSP ([#191](https://github.com/giantswarm/cert-manager-app/pull/191))


### chart-operator [2.19.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.19.0)

#### Removed
- Remove `tillermigration` resource now Helm 3 migration is complete.



### cluster-autoscaler [1.21.0-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.21.0-gs2)

Not found


### external-dns [2.6.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.6.0)

#### Added
- Add support for CAPZ clusters by detecting the Azure configuration file location.



### kube-state-metrics [1.4.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.4.0)

#### Changed
- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.



### metrics-server [1.5.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.5.0)

#### Changed
- Bumped API version for `RoleBinding` to `v1` as it was using a deprecated version (removed in 1.22).



### net-exporter [1.10.3](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.3)

#### Changed
- Prepare helm values to configuration management.
- Update architect-orb to v4.0.0.



### node-exporter [1.8.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.8.0)

#### Changed
- Migrate to configuration management.
- Update `architect-orb` to v4.0.0.



