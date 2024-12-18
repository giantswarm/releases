# :zap: Giant Swarm Release v19.3.1 for AWS :zap:

This patch release addresses an error in the calculation of max pods per node when using Cilium in ENI IPAM mode.

## Change details


### aws-operator [15.0.0](https://github.com/giantswarm/aws-operator/releases/tag/v15.0.0)

#### Fixed
- Bump k8scc to fix calculation of max pods per node when in ENI mode.
#### Changed
- [Breaking change] Removed code that allowed switching from AWS-CNI to Cilium. Releases using this AWS-operator can't be upgraded to from v18 releases.



### containerlinux [3602.2.3](https://www.flatcar-linux.org/releases/#release-3602.2.3)

 _Changes since **Stable 3602.2.2**_
 
 #### Security fixes:
 
 - Linux ([CVE-2023-46862](https://nvd.nist.gov/vuln/detail/CVE-2023-46862), [CVE-2023-6121](https://nvd.nist.gov/vuln/detail/CVE-2023-6121))
 
 #### Bug fixes:
 
 - Deleted files in `/etc` that have a tmpfiles rule that normally would recreate them will now show up again through the `/etc` lowerdir ([Flatcar#1265](https://github.com/flatcar/Flatcar/issues/1265), [bootengine#79](https://github.com/flatcar/bootengine/pull/79))
 
 #### Updates:
 
 - Linux ([5.15.142](https://lwn.net/Articles/954114) (includes [5.15.141](https://lwn.net/Articles/953649/), [5.15.140](https://lwn.net/Articles/953130), [5.15.139](https://lwn.net/Articles/952004)))
 - ca-certificates ([3.95](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_95.html))


### etcd [3.5.11](https://github.com/etcd-io/etcd/releases/tag/v3.5.11)

#### etcd server
- Fix distributed tracing by ensuring `--experimental-distributed-tracing-sampling-rate` configuration option is available to [set tracing sample rate](https://github.com/etcd-io/etcd/pull/16951).
- Fix [url redirects while checking peer urls during new member addition](https://github.com/etcd-io/etcd/pull/16986)
#### Dependencies
- Compile binaries using [go 1.20.12](https://github.com/etcd-io/etcd/pull/17077)
- Fix [CVE-2023-47108](https://github.com/advisories/GHSA-8pgv-569h-w5rw) by [bumping go.opentelemetry.io/otel to 1.20.0 and go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc to 0.46.0](https://github.com/etcd-io/etcd/pull/16946).



### app-operator [6.10.2](https://github.com/giantswarm/app-operator/releases/tag/v6.10.2)

#### Changed
- Set `gsoci.azurecr.io` as the default container registry for this app's image(s).



### external-dns [3.0.0](https://github.com/giantswarm/external-dns-app/releases/tag/v3.0.0)

#### Added
- Add vendir for upstream sync.
- Add namespaced feature to scope permissions to one namespace.
- Add support for Gateway API ([#305](https://github.com/giantswarm/external-dns-app/pull/305)).
#### Changed
- Deployment: Align to upstream ([#255](https://github.com/giantswarm/external-dns-app/pull/255)).
  - Use `crd.podSecurityContext` for crd job.
  - Rename `global.resources` as `resources`.
  - Rename `externalDNS.extraArgs` as `extraArgs`.
  - Rename `externalDNS.policy` as `policy`.
  - Rename `externalDNS.sources` as `sources` and adjust default value.
  - Rename `externalDNS.interval` as `interval`.
  - Rename `global.image` as `image` using helper for name composition.
  - Move `global.securityContext` to `podSecurityContext` and align names.
- Service: Align to upstream ([#243](https://github.com/giantswarm/external-dns-app/pull/243)).
  - Replace `global.metrics.port` value with `service.port`.
  - Add service annotations with GS defaults.
  - Set readinessProbe and livenessProbe from values.
  - Move podAnnotations to values.
- Update README and config docs ([#290](https://github.com/giantswarm/external-dns-app/pull/290)).
- Switch Registry to ACR ([#318](https://github.com/giantswarm/external-dns-app/pull/318)).
#### Removed
- Deployment: Align to upstream ([#255](https://github.com/giantswarm/external-dns-app/pull/255)).
  - Remove dedicated option for `min-event-sync-interval` and set it in extraArgs.
  - Remove `externalDNS.dryRun` option.
- Secrets: Remove deprecated values for AWS Route53 external authentication [#266](https://github.com/giantswarm/external-dns-app/pull/266).
- Remove support for KIAM ([#278](https://github.com/giantswarm/external-dns-app/pull/278)).
- Remove `aws.iam.customRoleName` value ([#278](https://github.com/giantswarm/external-dns-app/pull/278)).
- Remove `aws`, `gcpProject` and `externalDNS` values ([#284](https://github.com/giantswarm/external-dns-app/pull/284)).
- Remove Azure volume configuration ([#284](https://github.com/giantswarm/external-dns-app/pull/284)).
- Remove unused helpers ([#290](https://github.com/giantswarm/external-dns-app/pull/290)).
- Remove PSP ([#305](https://github.com/giantswarm/external-dns-app/pull/305)).



### net-exporter [1.18.2](https://github.com/giantswarm/net-exporter/releases/tag/v1.18.2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### security-bundle [1.4.2](https://github.com/giantswarm/security-bundle/releases/tag/v1.4.2)

#### Changed
- Update to `kyverno-app` (app) version v0.16.4.
- Update to `kyverno-policies` (app) version v0.20.2.
- Update to `exception-recommender` (app) to v0.0.6.
- Update to `starboard-exporter` (app) version v0.7.5. 
#### Added
- Add `options` to individual app settings to allow custom timeout values.



### chart-operator [3.1.1](https://github.com/giantswarm/chart-operator/releases/tag/v3.1.1)

#### Changed
- Configure gsoci.azurecr.io as the registry to use by default



### cluster-autoscaler [1.24.3-gs4](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.24.3-gs4)

#### Added
- Service monitor.



### coredns [1.20.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.20.0)

#### Added
- Add NET_BIND_SERVICE capability back to containers.
#### Changed
- Upgrade CoreDNS to [v1.11.1](https://github.com/coredns/coredns/releases/tag/v1.11.1).



