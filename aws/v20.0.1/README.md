# :zap: Giant Swarm Release v20.0.1 for AWS :zap:

<< Add description here >>

## Change details


### cert-operator [3.4.0](https://github.com/giantswarm/cert-operator/releases/tag/v3.4.0)

#### Changed
- Avoid exiting with a failure at startup time if the PKI cleanup fails.



### containerlinux [3815.2.1](https://www.flatcar-linux.org/releases/#release-3815.2.1)

_Changes since **Stable 3815.2.0**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-52429](https://nvd.nist.gov/vuln/detail/CVE-2023-52429), [CVE-2023-52434](https://nvd.nist.gov/vuln/detail/CVE-2023-52434), [CVE-2023-52435](https://nvd.nist.gov/vuln/detail/CVE-2023-52435), [CVE-2024-0340](https://nvd.nist.gov/vuln/detail/CVE-2024-0340), [CVE-2024-1151](https://nvd.nist.gov/vuln/detail/CVE-2024-1151), [CVE-2024-23850](https://nvd.nist.gov/vuln/detail/CVE-2024-23850), [CVE-2024-23851](https://nvd.nist.gov/vuln/detail/CVE-2024-23851), [CVE-2024-26582](https://nvd.nist.gov/vuln/detail/CVE-2024-26582), [CVE-2024-26583](https://nvd.nist.gov/vuln/detail/CVE-2024-26583), [CVE-2024-26586](https://nvd.nist.gov/vuln/detail/CVE-2024-26586), [CVE-2024-26593](https://nvd.nist.gov/vuln/detail/CVE-2024-26593))
 
 #### Bug fixes:
 
 - Fixed that systemd-sysext images can extend directories where Flatcar extensions are also shipping files, e.g., that the sysext-bakery Kubernetes extension works when OEM extensions are present ([sysext-bakery#50](https://github.com/flatcar/sysext-bakery/issues/50))
 - Fixed the handling of OEM update payloads in a Nebraska response with self-hosted packages in an airgapped environment ([update_engine#39](https://github.com/flatcar/update_engine/pull/39))
 - Restored support for custom OEMs supplied in the PXE boot where `/usr/share/oem` brings the OEM partition contents ([Flatcar#1376](https://github.com/flatcar/Flatcar/issues/1376))
 
 #### Changes:
 
 
 #### Updates:
 
 - Linux ([6.1.81](https://lwn.net/Articles/964562) (includes [6.1.80](https://lwn.net/Articles/964174), [6.1.79](https://lwn.net/Articles/963358), [6.1.78](https://lwn.net/Articles/962559)))
 - ca-certificates ([3.98](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_98.html))
 - keyutils ([1.6.3](https://git.kernel.org/pub/scm/linux/kernel/git/dhowells/keyutils.git/commit/?id=cb3bb194cca88211cbfcdde2f10c0f43c3fb8ec3) (includes [1.6.2](https://git.kernel.org/pub/scm/linux/kernel/git/dhowells/keyutils.git/commit/?id=454f80f537e5d1aad506599b6776e4cc1cf5f0f2)))


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
#### Others
- [Make CGO_ENABLED configurable](https://github.com/etcd-io/etcd/pull/17421).



### cluster-autoscaler [1.25.3-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.25.3-gs2)

#### Added
- Add possibility to use egress proxy.
#### Changed
- Chart: Improve proxy settings. ([#249](https://github.com/giantswarm/cluster-autoscaler-app/pull/249))



### aws-ebs-csi-driver [2.30.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.30.0)

#### Removed
- Remove unused service for controller.
- Add `Port` definition for metrics port in controller.



### cert-manager [3.7.4](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.4)

#### Added
- Added support for `AzureDNS` integration with a `Service Principal` on `clusterIssuer` helm chart .
#### Changed
- Changed `appVersion` to `v1.14.2`



### cilium [0.22.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.22.0)

#### Added
- Add helm values schema.
#### Changed
- Add safe-to-evict annotations to Hubble Relay and UI pods.
- Enable deletion of extra network policies.
- Update team label to `cabbage`



### etcd-kubernetes-resources-count-exporter [1.10.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.10.0)

#### Changed 
- Set min VPA settings and adjust CPU and memory resources.
- Use PodMonitor instead of legacy labels for monitoring.



### k8s-audit-metrics [0.9.0](https://github.com/giantswarm/k8s-audit-metrics/releases/tag/v0.9.0)

#### Added
- Add team label in resources.
- Use ServiceMonitor for monitoring.
#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### chart-operator [3.2.1](https://github.com/giantswarm/chart-operator/releases/tag/v3.2.1)

#### Fixed
- Use separate rest configs for different Kubernetes clients.



### external-dns [3.1.0](https://github.com/giantswarm/external-dns-app/releases/tag/v3.1.0)

#### Changed
- Remove default namespaceFilter configuration. ([#324](https://github.com/giantswarm/external-dns-app/pull/324)).



### vertical-pod-autoscaler [5.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v5.1.0)

#### Added
- Repository: Add ATS. ([#267](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/267))
#### Changed
- Chart: Improve readability of image tag. ([#263](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/263))
- Repository: Chores. ([#266](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/266))
  - Repository: Move `.kube-linter.yaml`.
  - Repository: Rework ABS.
  - Repository: Rework CircleCI.
  - Repository: Rework README.
  - Chart: Regenerate values schema JSON.
  - Chart: Rework `Chart.yaml`.
  - Chart: Rework `README.md.gotmpl`.
- Chart: Rework chart. ([#269](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/269))
  - Chart: Rework helpers.
  - Chart: Rework vertical pod autoscalers.
  - Chart: Rework policy exceptions.
  - Chart: Rework network policies.
  - Chart: Rework CRD patch.
- Chart: PSS compliance. ([#270](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/270))
#### Removed
- Repository: Chores. ([#266](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/266))
  - Repository: Remove unused script.
  - Repository: Remove `.nancy-ignore*`.
  - Repository: Remove images.
  - Repository: Remove config.
  - Repository: Remove `.gitignore`.
  - Chart: Remove `.helmignore`.
  - Chart: Remove useless CI values.
- Chart: Rework chart. ([#269](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/269))
  - Chart: Remove global network policies.



