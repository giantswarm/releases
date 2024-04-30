# :zap: Giant Swarm Release v20.1.0 for AWS :zap:

This release provides security updates for container linux and a fix for IMDSv2 only clusters.

## Change details


### aws-operator [16.1.1](https://github.com/giantswarm/aws-operator/releases/tag/v16.1.1)

#### Fixed
- Bump k8scc to fix issues with IMDS v2.



### cert-operator [3.4.0](https://github.com/giantswarm/cert-operator/releases/tag/v3.4.0)

#### Changed
- Avoid exiting with a failure at startup time if the PKI cleanup fails.



### cluster-operator [5.11.1](https://github.com/giantswarm/cluster-operator/releases/tag/v5.11.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.
#### Added
- Add team label in resources.
- Add `global.podSecurityStandards.enforced` value for PSS migration.
### Fixed
- Fix release version check for PSS enforcement.


### containerlinux [3815.2.2](https://www.flatcar-linux.org/releases/#release-3815.2.2)

 _Changes since **Stable 3815.2.1**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-28746](https://nvd.nist.gov/vuln/detail/CVE-2023-28746), [CVE-2023-47233](https://nvd.nist.gov/vuln/detail/CVE-2023-47233), [CVE-2023-52639](https://nvd.nist.gov/vuln/detail/CVE-2023-52639), [CVE-2023-6270](https://nvd.nist.gov/vuln/detail/CVE-2023-6270), [CVE-2023-7042](https://nvd.nist.gov/vuln/detail/CVE-2023-7042), [CVE-2024-22099](https://nvd.nist.gov/vuln/detail/CVE-2024-22099), [CVE-2024-23307](https://nvd.nist.gov/vuln/detail/CVE-2024-23307), [CVE-2024-24861](https://nvd.nist.gov/vuln/detail/CVE-2024-24861), [CVE-2024-26584](https://nvd.nist.gov/vuln/detail/CVE-2024-26584), [CVE-2024-26585](https://nvd.nist.gov/vuln/detail/CVE-2024-26585), [CVE-2024-26642](https://nvd.nist.gov/vuln/detail/CVE-2024-26642), [CVE-2024-26651](https://nvd.nist.gov/vuln/detail/CVE-2024-26651), [CVE-2024-26654](https://nvd.nist.gov/vuln/detail/CVE-2024-26654), [CVE-2024-26659](https://nvd.nist.gov/vuln/detail/CVE-2024-26659), [CVE-2024-26686](https://nvd.nist.gov/vuln/detail/CVE-2024-26686), [CVE-2024-26700](https://nvd.nist.gov/vuln/detail/CVE-2024-26700), [CVE-2024-26809](https://nvd.nist.gov/vuln/detail/CVE-2024-26809))
 - Downgraded xz-utils to 5.4.2 as precaution even though Flatcar is not affected of the SSH backdoor ([CVE-2024-3094](https://nvd.nist.gov/vuln/detail/CVE-2024-3094))
 - openssh ([CVE-2023-48795](https://nvd.nist.gov/vuln/detail/CVE-2023-48795), [CVE-2023-51384](https://nvd.nist.gov/vuln/detail/CVE-2023-51384), [CVE-2023-51385](https://nvd.nist.gov/vuln/detail/CVE-2023-51385))
 
 #### Bug fixes:
 
 - Disabled user-configdrive.service on OpenStack when config drive is used, which caused the hostname to be overwritten. The coreos-cloudinit.service unit already runs on OpenStack if the system is not configured via ignition. ([Flatcar#1385](https://github.com/flatcar/Flatcar/issues/1385))
 - Fixed `toolbox` to prevent mounted `ctr` snapshots from being garbage-collected ([toolbox#9](https://github.com/flatcar/toolbox/pull/9))
 
 #### Changes:
 
 - Disabled real-time priority for multipathd as it prevents the cgroups2 cpu controller from working. ([scripts#1771](https://github.com/flatcar/scripts/pull/1771))
 - SDK: Unified qemu image formats, so that the `qemu_uefi` build target provides the regular `qemu` and the `qemu_uefi_secure` artifacts ([scripts#1847](https://github.com/flatcar/scripts/pull/1847))
 
 #### Updates:
 
 - Linux ([6.1.85](https://lwn.net/Articles/969355) (includes [6.1.84](https://lwn.net/Articles/968254), [6.1.83](https://lwn.net/Articles/966759), [6.1.82](https://lwn.net/Articles/965607)))
 - ca-certificates ([3.99](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_99.html))
 - openssh ([9.6p1](https://www.openssh.com/releasenotes.html#9.6p1))



### etcd [3.5.13](https://github.com/etcd-io/etcd/releases/tag/v3.5.13)

#### etcd server
- Fix leases wrongly revoked by the leader by [ignoring old leader's leases revoking request](https://github.com/etcd-io/etcd/pull/17425).
- Fix [no progress notification being sent for watch that doesn't get any events](https://github.com/etcd-io/etcd/pull/17566).
- Fix [watch event loss after compaction](https://github.com/etcd-io/etcd/pull/17612).
#### Package `clientv3`
- Add [client backoff and retry config options](https://github.com/etcd-io/etcd/pull/17363).
- [Ignore SetKeepAlivePeriod errors on OpenBSD](https://github.com/etcd-io/etcd/pull/17387).
- [Support unix/unixs socket in client or peer URLs](https://github.com/etcd-io/etcd/pull/15940)
#### gRPC Proxy
- Add [three flags (see below) for grpc-proxy](https://github.com/etcd-io/etcd/pull/17447)
  - `--dial-keepalive-time`
  - `--dial-keepalive-timeout`
  - `--permit-without-stream`
#### Dependencies
- Upgrade [bbolt to v1.3.9](https://github.com/etcd-io/etcd/pull/17483).
- Compile binaries using [go 1.21.8](https://github.com/etcd-io/etcd/pull/17537).
- Upgrade [google.golang.org/protobuf to v1.33.0 to address CVE-2024-24786](https://github.com/etcd-io/etcd/pull/17553).
- Upgrade github.com/sirupsen/logrus to v1.9.3 to address [PRISMA-2023-0056](https://github.com/etcd-io/etcd/pull/17482).
#### Others
- [Make CGO_ENABLED configurable](https://github.com/etcd-io/etcd/pull/17421).



### app-operator [6.11.0](https://github.com/giantswarm/app-operator/releases/tag/v6.11.0)

#### Added
- Add support for App resources having a dependency on HelmReleases.



### vertical-pod-autoscaler [5.2.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v5.2.1)

#### Changed
- Chart: Update `appVersion` and `README.md`. ([#281](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/281))



### etcd-kubernetes-resources-count-exporter [1.10.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.10.0)

#### Changed 
- Set min VPA settings and adjust CPU and memory resources.
- Use PodMonitor instead of legacy labels for monitoring.



### vertical-pod-autoscaler-crd [3.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v3.1.0)

- Chart: Sync CRDs to VPA v1.1.0. ([#93](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/93))



### observability-bundle [1.3.4](https://github.com/giantswarm/observability-bundle/releases/tag/v1.3.4)

#### Changed
- Upgrade `kube-prometheus-stack` to 9.1.2.



### k8s-audit-metrics [0.9.0](https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.9.0)

#### Added
- Add team label in resources.
- Use ServiceMonitor for monitoring.
#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### cert-manager [3.7.4](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.4)

#### Added
- Added support for `AzureDNS` integration with a `Service Principal` on `clusterIssuer` helm chart .
#### Changed
- Changed `appVersion` to `v1.14.2`



### chart-operator [3.2.1](https://github.com/giantswarm/chart-operator/releases/tag/v3.2.1)

#### Fixed
- Use separate rest configs for different Kubernetes clients.



### cilium [0.22.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.22.0)

#### Added
- Add helm values schema.
#### Changed
- Add safe-to-evict annotations to Hubble Relay and UI pods.
- Enable deletion of extra network policies.
- Update team label to `cabbage`



### cluster-autoscaler [1.25.3-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.25.3-gs2)

#### Added
- Add possibility to use egress proxy.
#### Changed
- Chart: Improve proxy settings. ([#249](https://github.com/giantswarm/cluster-autoscaler-app/pull/249))



### external-dns [3.1.0](https://github.com/giantswarm/external-dns-app/releases/tag/v3.1.0)

#### Changed
- Remove default namespaceFilter configuration. ([#324](https://github.com/giantswarm/external-dns-app/pull/324)).



