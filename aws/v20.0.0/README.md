# :zap: Giant Swarm Release v20.0.0 for AWS :zap:

The new major version is the first release supporting Kubernetes 1.25. 

Please remember that the Pod Security Policies (PSP) are no longer supported in this version and they have to be removed *BEFORE* the upgrade. From now on the Pod Security Standards (PSS) are the default for securing the clusters. Our docs offer additional information about [Pod Security Standards](https://docs.giantswarm.io/advanced/security-policy-enforcement/) as well as a [PSS migration guide](https://docs.giantswarm.io/advanced/security/security-policy-enforcement/cluster-admin-guide/) for cluster administrators.

This release will also be used as a base for the migration from `Giant Swarm Vintage` to `Cluster API for AWS` (CAPA). The migration will be scheduled as any other upgrade. We will be reaching out separately to each customer to present the plan as well as the CAPA benefits.

> **WARNING:** After upgrading to `19.3.0`, it is highly advised to begin removal of all customer-managed PSPs from the cluster. Kubernetes `v1.25` removes the Pod Security Policy resource from the API, meaning workloads (like Helm charts) which still contain PSPs will fail to install after the Giant Swarm `v20` upgrade.

> **WARNING:** `observability-bundle` will be upgraded to `v1.3.0`, which contains breaking changes to the configuration for the bundled apps. Please check our [upgrade guide](https://github.com/giantswarm/observability-bundle/blob/main/docs/upgrade.md) or reach out to your Account Engineer for more details.

## Change details


### app-operator [6.10.3](https://github.com/giantswarm/app-operator/releases/tag/v6.10.3)

#### Fixed
- Move pss values under the global property



### aws-operator [16..0](https://github.com/giantswarm/aws-operator/releases/tag/v16.1.0)

#### Changed
- Bump k8scc to v18 to enable k8s 1.25 support.
- Bump k8scc to avoid running etcd defrag on all masters at the same time.

### Fixed

- Handle karpenter nodes in node-termination-handler.

### etcd [3.5.12](https://github.com/etcd-io/etcd/releases/tag/v3.5.12)

#### etcd server
- Add [livez/readyz HTTP endpoints](https://github.com/etcd-io/etcd/pull/17039)
- Fix [not validating database consistent index, and panicking on nil backend](https://github.com/etcd-io/etcd/pull/17151)
- Document [`experimental-enable-lease-checkpoint-persist` flag in etcd help](https://github.com/etcd-io/etcd/pull/17190)
- Fix [needlessly flocking snapshot files when deleting](https://github.com/etcd-io/etcd/pull/17206)
- Add [digest for etcd base image](https://github.com/etcd-io/etcd/pull/17205)
- Fix [delete inconsistencies in read buffer](https://github.com/etcd-io/etcd/pull/17230)
#### Dependencies
- Compile binaries using [go 1.20.13](https://github.com/etcd-io/etcd/pull/17275)
- Upgrade [golang.org/x/crypto to v0.17+ to address CVE-2023-48795](https://github.com/etcd-io/etcd/pull/17346)



### chart-operator-extensions [1.1.2](https://github.com/giantswarm/chart-operator-extensions/releases/tag/v1.1.2)

#### Fixed
- Move pss values under the global property



### node-exporter [1.19.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.19.0)

#### Added
- Add VPA configuration to `node-exporter` app.



### k8s-dns-node-cache-app [2.6.1](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.6.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.
- Revert force_tcp option from external DNS zone (#67).



### aws-ebs-csi-driver [2.28.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.28.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### cluster-autoscaler [1.25.1-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.25.1-gs2)

#### Fixed
- Adjusted minimum allowed CPU and memory



### prometheus-blackbox-exporter [0.4.1](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.4.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### security-bundle [1.6.2](https://github.com/giantswarm/security-bundle/releases/tag/v1.6.2)

#### Changed
- Update to exception-recommender (app) to v0.1.1.
- Update to falco (app) to v0.8.0.
- Update to kyverno-policy-operator (app) version v0.0.7.
- Update to kyverno (app) version v0.17.6.
- Update to starboard-exporter (app) version v0.7.8.
- Update to trivy-operator (app) to v0.7.2.
- Update to trivy (app) to v0.10.0.
- Update to kyverno (app) to v0.17.5.


### aws-cloud-controller-manager [1.25.14-gs2](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.25.14-gs2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### chart-operator [3.1.3](https://github.com/giantswarm/chart-operator/releases/tag/v3.1.3)

#### Fixed
- Move pss values under the global property
#### Changed
- Use base images from `gsoci.azurecr.io`



### cilium [0.19.2](https://github.com/giantswarm/cilium-app/releases/tag/v0.19.2)

#### Fixed
- Replace `ToServices`/`ToPorts` combination in CiliumNetworkPolicy because of breakage in Cilium v1.14

#### Changed
- Upgrade cilium to 1.14.5.


### metrics-server [2.4.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.4.2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### vertical-pod-autoscaler-crd [3.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v3.0.0)

#### Changed
- Synced VPA CRD for v1.0.0



### etcd-kubernetes-resources-count-exporter [1.9.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.9.0)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### observability-bundle [1.3.0](https://github.com/giantswarm/observability-bundle/releases/tag/v1.3.0)

#### Changed
- *Breaking change*: Simplify configuration for the bundled apps. See our [upgrade guide](https://github.com/giantswarm/observability-bundle/blob/main/docs/upgrade.md)
  - Move all user configs from under `apps.<appName>.userConfig` from string to regular helm values to `userConfig.<appName>`
  - Rename `prometheus-operator-app` to `kube-prometheus-stack`
  - Rename `promtail-app` to `promtail`

- Upgrade `kube-prometheus-stack` to 9.1.0.
- Upgrade `prometheus-operator-crd` to 9.0.0.
- Add the global.podSecurityStandards.enforced value back to be able to work on CAPI WCs.
- Add dependency on prometheus-operator-crd to all apps.
- Upgrade `promtail` to 1.5.1.
- Upgrade `grafana-agent` to 0.4.1.
- upgrade `prometheus-agent` to 0.6.9.
- Enforce `Cilium Network Policy` by default.
- Enforce `Pod Security Standard` by default.

### cert-exporter [2.9.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.9.0)

#### Added
- Add cert name to secret metric.


### cert-manager [3.7.1](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.1)

#### Added
- Added `acme-solvers-networkpolicy` `NetworkPolicy` namespace to `kube-system`


### coredns [1.21.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.21.0)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.


### vertical-pod-autoscaler [5.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v5.0.0)

#### Changed
- Change ImageRegistry to `gsoci.azurecr.io`.
- Upgrade dependency chart to 9.6.0

### containerlinux [3815.2.0](https://www.flatcar-linux.org/releases/#release-3815.2.0)

⚠️ From Alpha 3794.0.0 Torcx has been removed - please assert that you don’t rely on specific Torcx mechanism but now use systemd-sysext. See [here](https://www.flatcar.org/docs/latest/provisioning/sysext/) for more information.

**Changes since Stable-3602.2.3**

#### Security fixes
- Linux ([CVE-2023-7192](https://nvd.nist.gov/vuln/detail/CVE-2023-7192) (includes [CVE-2023-6932](https://nvd.nist.gov/vuln/detail/CVE-2023-6932), [CVE-2023-6931](https://nvd.nist.gov/vuln/detail/CVE-2023-6931), [CVE-2023-6817](https://nvd.nist.gov/vuln/detail/CVE-2023-6817), [CVE-2023-6622](https://nvd.nist.gov/vuln/detail/CVE-2023-6622), [CVE-2023-6606](https://nvd.nist.gov/vuln/detail/CVE-2023-6606), [CVE-2023-6546](https://nvd.nist.gov/vuln/detail/CVE-2023-6546), [CVE-2023-6531](https://nvd.nist.gov/vuln/detail/CVE-2023-6531), [CVE-2023-6176](https://nvd.nist.gov/vuln/detail/CVE-2023-6176), [CVE-2023-6121](https://nvd.nist.gov/vuln/detail/CVE-2023-6121), [CVE-2023-5717](https://nvd.nist.gov/vuln/detail/CVE-2023-5717), [CVE-2023-5345](https://nvd.nist.gov/vuln/detail/CVE-2023-5345), [CVE-2023-5197](https://nvd.nist.gov/vuln/detail/CVE-2023-5197), [CVE-2023-51782](https://nvd.nist.gov/vuln/detail/CVE-2023-51782), [CVE-2023-51781](https://nvd.nist.gov/vuln/detail/CVE-2023-51781), [CVE-2023-51780](https://nvd.nist.gov/vuln/detail/CVE-2023-51780), [CVE-2023-51779](https://nvd.nist.gov/vuln/detail/CVE-2023-51779), [CVE-2023-5158](https://nvd.nist.gov/vuln/detail/CVE-2023-5158), [CVE-2023-5090](https://nvd.nist.gov/vuln/detail/CVE-2023-5090), [CVE-2023-4921](https://nvd.nist.gov/vuln/detail/CVE-2023-4921), [CVE-2023-46862](https://nvd.nist.gov/vuln/detail/CVE-2023-46862), [CVE-2023-46813](https://nvd.nist.gov/vuln/detail/CVE-2023-46813), [CVE-2023-4623](https://nvd.nist.gov/vuln/detail/CVE-2023-4623), [CVE-2023-45871](https://nvd.nist.gov/vuln/detail/CVE-2023-45871), [CVE-2023-45863](https://nvd.nist.gov/vuln/detail/CVE-2023-45863), [CVE-2023-45862](https://nvd.nist.gov/vuln/detail/CVE-2023-45862), [CVE-2023-4569](https://nvd.nist.gov/vuln/detail/CVE-2023-4569), [CVE-2023-4459](https://nvd.nist.gov/vuln/detail/CVE-2023-4459), [CVE-2023-44466](https://nvd.nist.gov/vuln/detail/CVE-2023-44466), [CVE-2023-4394](https://nvd.nist.gov/vuln/detail/CVE-2023-4394), [CVE-2023-4389](https://nvd.nist.gov/vuln/detail/CVE-2023-4389), [CVE-2023-4387](https://nvd.nist.gov/vuln/detail/CVE-2023-4387), [CVE-2023-4385](https://nvd.nist.gov/vuln/detail/CVE-2023-4385), [CVE-2023-42755](https://nvd.nist.gov/vuln/detail/CVE-2023-42755), [CVE-2023-42754](https://nvd.nist.gov/vuln/detail/CVE-2023-42754), [CVE-2023-42753](https://nvd.nist.gov/vuln/detail/CVE-2023-42753), [CVE-2023-42752](https://nvd.nist.gov/vuln/detail/CVE-2023-42752), [CVE-2023-4273](https://nvd.nist.gov/vuln/detail/CVE-2023-4273), [CVE-2023-4244](https://nvd.nist.gov/vuln/detail/CVE-2023-4244), [CVE-2023-4208](https://nvd.nist.gov/vuln/detail/CVE-2023-4208), [CVE-2023-4207](https://nvd.nist.gov/vuln/detail/CVE-2023-4207), [CVE-2023-4206](https://nvd.nist.gov/vuln/detail/CVE-2023-4206), [CVE-2023-4155](https://nvd.nist.gov/vuln/detail/CVE-2023-4155), [CVE-2023-4147](https://nvd.nist.gov/vuln/detail/CVE-2023-4147), [CVE-2023-4132](https://nvd.nist.gov/vuln/detail/CVE-2023-4132), [CVE-2023-40283](https://nvd.nist.gov/vuln/detail/CVE-2023-40283), [CVE-2023-4015](https://nvd.nist.gov/vuln/detail/CVE-2023-4015), [CVE-2023-4004](https://nvd.nist.gov/vuln/detail/CVE-2023-4004), [CVE-2023-39198](https://nvd.nist.gov/vuln/detail/CVE-2023-39198), [CVE-2023-39197](https://nvd.nist.gov/vuln/detail/CVE-2023-39197), [CVE-2023-39194](https://nvd.nist.gov/vuln/detail/CVE-2023-39194), [CVE-2023-39193](https://nvd.nist.gov/vuln/detail/CVE-2023-39193), [CVE-2023-39192](https://nvd.nist.gov/vuln/detail/CVE-2023-39192), [CVE-2023-39189](https://nvd.nist.gov/vuln/detail/CVE-2023-39189), [CVE-2023-3867](https://nvd.nist.gov/vuln/detail/CVE-2023-3867), [CVE-2023-3866](https://nvd.nist.gov/vuln/detail/CVE-2023-3866), [CVE-2023-3865](https://nvd.nist.gov/vuln/detail/CVE-2023-3865), [CVE-2023-3863](https://nvd.nist.gov/vuln/detail/CVE-2023-3863), [CVE-2023-38432](https://nvd.nist.gov/vuln/detail/CVE-2023-38432), [CVE-2023-38431](https://nvd.nist.gov/vuln/detail/CVE-2023-38431), [CVE-2023-38430](https://nvd.nist.gov/vuln/detail/CVE-2023-38430), [CVE-2023-38429](https://nvd.nist.gov/vuln/detail/CVE-2023-38429), [CVE-2023-38428](https://nvd.nist.gov/vuln/detail/CVE-2023-38428), [CVE-2023-38427](https://nvd.nist.gov/vuln/detail/CVE-2023-38427), [CVE-2023-38426](https://nvd.nist.gov/vuln/detail/CVE-2023-38426), [CVE-2023-38409](https://nvd.nist.gov/vuln/detail/CVE-2023-38409), [CVE-2023-3812](https://nvd.nist.gov/vuln/detail/CVE-2023-3812), [CVE-2023-3777](https://nvd.nist.gov/vuln/detail/CVE-2023-3777), [CVE-2023-3776](https://nvd.nist.gov/vuln/detail/CVE-2023-3776), [CVE-2023-3773](https://nvd.nist.gov/vuln/detail/CVE-2023-3773), [CVE-2023-3772](https://nvd.nist.gov/vuln/detail/CVE-2023-3772), [CVE-2023-3611](https://nvd.nist.gov/vuln/detail/CVE-2023-3611), [CVE-2023-3610](https://nvd.nist.gov/vuln/detail/CVE-2023-3610), [CVE-2023-3609](https://nvd.nist.gov/vuln/detail/CVE-2023-3609), [CVE-2023-35829](https://nvd.nist.gov/vuln/detail/CVE-2023-35829), [CVE-2023-35828](https://nvd.nist.gov/vuln/detail/CVE-2023-35828), [CVE-2023-35827](https://nvd.nist.gov/vuln/detail/CVE-2023-35827), [CVE-2023-35826](https://nvd.nist.gov/vuln/detail/CVE-2023-35826), [CVE-2023-35824](https://nvd.nist.gov/vuln/detail/CVE-2023-35824), [CVE-2023-35823](https://nvd.nist.gov/vuln/detail/CVE-2023-35823), [CVE-2023-35788](https://nvd.nist.gov/vuln/detail/CVE-2023-35788), [CVE-2023-3567](https://nvd.nist.gov/vuln/detail/CVE-2023-3567), [CVE-2023-35001](https://nvd.nist.gov/vuln/detail/CVE-2023-35001), [CVE-2023-3439](https://nvd.nist.gov/vuln/detail/CVE-2023-3439), [CVE-2023-34324](https://nvd.nist.gov/vuln/detail/CVE-2023-34324), [CVE-2023-34319](https://nvd.nist.gov/vuln/detail/CVE-2023-34319), [CVE-2023-34256](https://nvd.nist.gov/vuln/detail/CVE-2023-34256), [CVE-2023-33952](https://nvd.nist.gov/vuln/detail/CVE-2023-33952), [CVE-2023-33951](https://nvd.nist.gov/vuln/detail/CVE-2023-33951), [CVE-2023-3390](https://nvd.nist.gov/vuln/detail/CVE-2023-3390), [CVE-2023-3359](https://nvd.nist.gov/vuln/detail/CVE-2023-3359), [CVE-2023-3358](https://nvd.nist.gov/vuln/detail/CVE-2023-3358), [CVE-2023-3357](https://nvd.nist.gov/vuln/detail/CVE-2023-3357), [CVE-2023-3355](https://nvd.nist.gov/vuln/detail/CVE-2023-3355), [CVE-2023-33288](https://nvd.nist.gov/vuln/detail/CVE-2023-33288), [CVE-2023-33203](https://nvd.nist.gov/vuln/detail/CVE-2023-33203), [CVE-2023-3269](https://nvd.nist.gov/vuln/detail/CVE-2023-3269), [CVE-2023-3268](https://nvd.nist.gov/vuln/detail/CVE-2023-3268), [CVE-2023-32269](https://nvd.nist.gov/vuln/detail/CVE-2023-32269), [CVE-2023-32258](https://nvd.nist.gov/vuln/detail/CVE-2023-32258), [CVE-2023-32257](https://nvd.nist.gov/vuln/detail/CVE-2023-32257), [CVE-2023-32254](https://nvd.nist.gov/vuln/detail/CVE-2023-32254), [CVE-2023-32252](https://nvd.nist.gov/vuln/detail/CVE-2023-32252), [CVE-2023-32250](https://nvd.nist.gov/vuln/detail/CVE-2023-32250), [CVE-2023-32248](https://nvd.nist.gov/vuln/detail/CVE-2023-32248), [CVE-2023-32247](https://nvd.nist.gov/vuln/detail/CVE-2023-32247), [CVE-2023-32233](https://nvd.nist.gov/vuln/detail/CVE-2023-32233), [CVE-2023-3220](https://nvd.nist.gov/vuln/detail/CVE-2023-3220), [CVE-2023-3212](https://nvd.nist.gov/vuln/detail/CVE-2023-3212), [CVE-2023-3161](https://nvd.nist.gov/vuln/detail/CVE-2023-3161), [CVE-2023-3159](https://nvd.nist.gov/vuln/detail/CVE-2023-3159), [CVE-2023-31436](https://nvd.nist.gov/vuln/detail/CVE-2023-31436), [CVE-2023-3141](https://nvd.nist.gov/vuln/detail/CVE-2023-3141), [CVE-2023-31248](https://nvd.nist.gov/vuln/detail/CVE-2023-31248), [CVE-2023-3111](https://nvd.nist.gov/vuln/detail/CVE-2023-3111), [CVE-2023-31085](https://nvd.nist.gov/vuln/detail/CVE-2023-31085), [CVE-2023-3090](https://nvd.nist.gov/vuln/detail/CVE-2023-3090), [CVE-2023-30772](https://nvd.nist.gov/vuln/detail/CVE-2023-30772), [CVE-2023-30456](https://nvd.nist.gov/vuln/detail/CVE-2023-30456), [CVE-2023-3006](https://nvd.nist.gov/vuln/detail/CVE-2023-3006), [CVE-2023-2985](https://nvd.nist.gov/vuln/detail/CVE-2023-2985), [CVE-2023-2898](https://nvd.nist.gov/vuln/detail/CVE-2023-2898), [CVE-2023-28866](https://nvd.nist.gov/vuln/detail/CVE-2023-28866), [CVE-2023-28466](https://nvd.nist.gov/vuln/detail/CVE-2023-28466), [CVE-2023-28410](https://nvd.nist.gov/vuln/detail/CVE-2023-28410), [CVE-2023-28328](https://nvd.nist.gov/vuln/detail/CVE-2023-28328), [CVE-2023-28327](https://nvd.nist.gov/vuln/detail/CVE-2023-28327), [CVE-2023-26607](https://nvd.nist.gov/vuln/detail/CVE-2023-26607), [CVE-2023-26606](https://nvd.nist.gov/vuln/detail/CVE-2023-26606), [CVE-2023-26545](https://nvd.nist.gov/vuln/detail/CVE-2023-26545), [CVE-2023-26544](https://nvd.nist.gov/vuln/detail/CVE-2023-26544), [CVE-2023-25775](https://nvd.nist.gov/vuln/detail/CVE-2023-25775), [CVE-2023-2513](https://nvd.nist.gov/vuln/detail/CVE-2023-2513), [CVE-2023-25012](https://nvd.nist.gov/vuln/detail/CVE-2023-25012), [CVE-2023-2430](https://nvd.nist.gov/vuln/detail/CVE-2023-2430), [CVE-2023-23559](https://nvd.nist.gov/vuln/detail/CVE-2023-23559), [CVE-2023-23455](https://nvd.nist.gov/vuln/detail/CVE-2023-23455), [CVE-2023-23454](https://nvd.nist.gov/vuln/detail/CVE-2023-23454), [CVE-2023-23002](https://nvd.nist.gov/vuln/detail/CVE-2023-23002), [CVE-2023-23001](https://nvd.nist.gov/vuln/detail/CVE-2023-23001), [CVE-2023-22999](https://nvd.nist.gov/vuln/detail/CVE-2023-22999), [CVE-2023-22998](https://nvd.nist.gov/vuln/detail/CVE-2023-22998), [CVE-2023-22997](https://nvd.nist.gov/vuln/detail/CVE-2023-22997), [CVE-2023-22996](https://nvd.nist.gov/vuln/detail/CVE-2023-22996), [CVE-2023-2269](https://nvd.nist.gov/vuln/detail/CVE-2023-2269), [CVE-2023-2236](https://nvd.nist.gov/vuln/detail/CVE-2023-2236), [CVE-2023-2235](https://nvd.nist.gov/vuln/detail/CVE-2023-2235), [CVE-2023-2194](https://nvd.nist.gov/vuln/detail/CVE-2023-2194), [CVE-2023-2177](https://nvd.nist.gov/vuln/detail/CVE-2023-2177), [CVE-2023-2166](https://nvd.nist.gov/vuln/detail/CVE-2023-2166), [CVE-2023-2163](https://nvd.nist.gov/vuln/detail/CVE-2023-2163), [CVE-2023-2162](https://nvd.nist.gov/vuln/detail/CVE-2023-2162), [CVE-2023-2156](https://nvd.nist.gov/vuln/detail/CVE-2023-2156), [CVE-2023-21255](https://nvd.nist.gov/vuln/detail/CVE-2023-21255), [CVE-2023-2124](https://nvd.nist.gov/vuln/detail/CVE-2023-2124), [CVE-2023-21106](https://nvd.nist.gov/vuln/detail/CVE-2023-21106), [CVE-2023-21102](https://nvd.nist.gov/vuln/detail/CVE-2023-21102), [CVE-2023-20938](https://nvd.nist.gov/vuln/detail/CVE-2023-20938), [CVE-2023-20928](https://nvd.nist.gov/vuln/detail/CVE-2023-20928), [CVE-2023-20593](https://nvd.nist.gov/vuln/detail/CVE-2023-20593), [CVE-2023-20588](https://nvd.nist.gov/vuln/detail/CVE-2023-20588), [CVE-2023-20569](https://nvd.nist.gov/vuln/detail/CVE-2023-20569), [CVE-2023-2019](https://nvd.nist.gov/vuln/detail/CVE-2023-2019), [CVE-2023-2008](https://nvd.nist.gov/vuln/detail/CVE-2023-2008), [CVE-2023-2006](https://nvd.nist.gov/vuln/detail/CVE-2023-2006), [CVE-2023-2002](https://nvd.nist.gov/vuln/detail/CVE-2023-2002), [CVE-2023-1998](https://nvd.nist.gov/vuln/detail/CVE-2023-1998), [CVE-2023-1990](https://nvd.nist.gov/vuln/detail/CVE-2023-1990), [CVE-2023-1989](https://nvd.nist.gov/vuln/detail/CVE-2023-1989), [CVE-2023-1872](https://nvd.nist.gov/vuln/detail/CVE-2023-1872), [CVE-2023-1859](https://nvd.nist.gov/vuln/detail/CVE-2023-1859), [CVE-2023-1855](https://nvd.nist.gov/vuln/detail/CVE-2023-1855), [CVE-2023-1838](https://nvd.nist.gov/vuln/detail/CVE-2023-1838), [CVE-2023-1829](https://nvd.nist.gov/vuln/detail/CVE-2023-1829), [CVE-2023-1670](https://nvd.nist.gov/vuln/detail/CVE-2023-1670), [CVE-2023-1652](https://nvd.nist.gov/vuln/detail/CVE-2023-1652), [CVE-2023-1637](https://nvd.nist.gov/vuln/detail/CVE-2023-1637), [CVE-2023-1611](https://nvd.nist.gov/vuln/detail/CVE-2023-1611), [CVE-2023-1583](https://nvd.nist.gov/vuln/detail/CVE-2023-1583), [CVE-2023-1582](https://nvd.nist.gov/vuln/detail/CVE-2023-1582), [CVE-2023-1513](https://nvd.nist.gov/vuln/detail/CVE-2023-1513), [CVE-2023-1382](https://nvd.nist.gov/vuln/detail/CVE-2023-1382), [CVE-2023-1380](https://nvd.nist.gov/vuln/detail/CVE-2023-1380), [CVE-2023-1281](https://nvd.nist.gov/vuln/detail/CVE-2023-1281), [CVE-2023-1249](https://nvd.nist.gov/vuln/detail/CVE-2023-1249), [CVE-2023-1206](https://nvd.nist.gov/vuln/detail/CVE-2023-1206), [CVE-2023-1194](https://nvd.nist.gov/vuln/detail/CVE-2023-1194), [CVE-2023-1193](https://nvd.nist.gov/vuln/detail/CVE-2023-1193), [CVE-2023-1192](https://nvd.nist.gov/vuln/detail/CVE-2023-1192), [CVE-2023-1118](https://nvd.nist.gov/vuln/detail/CVE-2023-1118), [CVE-2023-1095](https://nvd.nist.gov/vuln/detail/CVE-2023-1095), [CVE-2023-1079](https://nvd.nist.gov/vuln/detail/CVE-2023-1079), [CVE-2023-1078](https://nvd.nist.gov/vuln/detail/CVE-2023-1078), [CVE-2023-1077](https://nvd.nist.gov/vuln/detail/CVE-2023-1077), [CVE-2023-1076](https://nvd.nist.gov/vuln/detail/CVE-2023-1076), [CVE-2023-1075](https://nvd.nist.gov/vuln/detail/CVE-2023-1075), [CVE-2023-1074](https://nvd.nist.gov/vuln/detail/CVE-2023-1074), [CVE-2023-1073](https://nvd.nist.gov/vuln/detail/CVE-2023-1073), [CVE-2023-1032](https://nvd.nist.gov/vuln/detail/CVE-2023-1032), [CVE-2023-0615](https://nvd.nist.gov/vuln/detail/CVE-2023-0615), [CVE-2023-0590](https://nvd.nist.gov/vuln/detail/CVE-2023-0590), [CVE-2023-0469](https://nvd.nist.gov/vuln/detail/CVE-2023-0469), [CVE-2023-0468](https://nvd.nist.gov/vuln/detail/CVE-2023-0468), [CVE-2023-0461](https://nvd.nist.gov/vuln/detail/CVE-2023-0461), [CVE-2023-0459](https://nvd.nist.gov/vuln/detail/CVE-2023-0459), [CVE-2023-0458](https://nvd.nist.gov/vuln/detail/CVE-2023-0458), [CVE-2023-0394](https://nvd.nist.gov/vuln/detail/CVE-2023-0394), [CVE-2023-0386](https://nvd.nist.gov/vuln/detail/CVE-2023-0386), [CVE-2023-0266](https://nvd.nist.gov/vuln/detail/CVE-2023-0266), [CVE-2023-0210](https://nvd.nist.gov/vuln/detail/CVE-2023-0210), [CVE-2023-0179](https://nvd.nist.gov/vuln/detail/CVE-2023-0179), [CVE-2023-0160](https://nvd.nist.gov/vuln/detail/CVE-2023-0160), [CVE-2023-0045](https://nvd.nist.gov/vuln/detail/CVE-2023-0045), [CVE-2022-48619](https://nvd.nist.gov/vuln/detail/CVE-2022-48619), [CVE-2022-48502](https://nvd.nist.gov/vuln/detail/CVE-2022-48502), [CVE-2022-48425](https://nvd.nist.gov/vuln/detail/CVE-2022-48425), [CVE-2022-48424](https://nvd.nist.gov/vuln/detail/CVE-2022-48424), [CVE-2022-48423](https://nvd.nist.gov/vuln/detail/CVE-2022-48423), [CVE-2022-4842](https://nvd.nist.gov/vuln/detail/CVE-2022-4842), [CVE-2022-47943](https://nvd.nist.gov/vuln/detail/CVE-2022-47943), [CVE-2022-47942](https://nvd.nist.gov/vuln/detail/CVE-2022-47942), [CVE-2022-47941](https://nvd.nist.gov/vuln/detail/CVE-2022-47941), [CVE-2022-47940](https://nvd.nist.gov/vuln/detail/CVE-2022-47940), [CVE-2022-47939](https://nvd.nist.gov/vuln/detail/CVE-2022-47939), [CVE-2022-47938](https://nvd.nist.gov/vuln/detail/CVE-2022-47938), [CVE-2022-47929](https://nvd.nist.gov/vuln/detail/CVE-2022-47929), [CVE-2022-47521](https://nvd.nist.gov/vuln/detail/CVE-2022-47521), [CVE-2022-47520](https://nvd.nist.gov/vuln/detail/CVE-2022-47520), [CVE-2022-47519](https://nvd.nist.gov/vuln/detail/CVE-2022-47519), [CVE-2022-47518](https://nvd.nist.gov/vuln/detail/CVE-2022-47518), [CVE-2022-4662](https://nvd.nist.gov/vuln/detail/CVE-2022-4662), [CVE-2022-45934](https://nvd.nist.gov/vuln/detail/CVE-2022-45934), [CVE-2022-45919](https://nvd.nist.gov/vuln/detail/CVE-2022-45919), [CVE-2022-45887](https://nvd.nist.gov/vuln/detail/CVE-2022-45887), [CVE-2022-45886](https://nvd.nist.gov/vuln/detail/CVE-2022-45886), [CVE-2022-45869](https://nvd.nist.gov/vuln/detail/CVE-2022-45869), [CVE-2022-43945](https://nvd.nist.gov/vuln/detail/CVE-2022-43945), [CVE-2022-4382](https://nvd.nist.gov/vuln/detail/CVE-2022-4382), [CVE-2022-4379](https://nvd.nist.gov/vuln/detail/CVE-2022-4379), [CVE-2022-4378](https://nvd.nist.gov/vuln/detail/CVE-2022-4378), [CVE-2022-43750](https://nvd.nist.gov/vuln/detail/CVE-2022-43750), [CVE-2022-42896](https://nvd.nist.gov/vuln/detail/CVE-2022-42896), [CVE-2022-42895](https://nvd.nist.gov/vuln/detail/CVE-2022-42895), [CVE-2022-42722](https://nvd.nist.gov/vuln/detail/CVE-2022-42722), [CVE-2022-42721](https://nvd.nist.gov/vuln/detail/CVE-2022-42721), [CVE-2022-42720](https://nvd.nist.gov/vuln/detail/CVE-2022-42720), [CVE-2022-42719](https://nvd.nist.gov/vuln/detail/CVE-2022-42719), [CVE-2022-42703](https://nvd.nist.gov/vuln/detail/CVE-2022-42703), [CVE-2022-4269](https://nvd.nist.gov/vuln/detail/CVE-2022-4269), [CVE-2022-42432](https://nvd.nist.gov/vuln/detail/CVE-2022-42432), [CVE-2022-42329](https://nvd.nist.gov/vuln/detail/CVE-2022-42329), [CVE-2022-42328](https://nvd.nist.gov/vuln/detail/CVE-2022-42328), [CVE-2022-41858](https://nvd.nist.gov/vuln/detail/CVE-2022-41858), [CVE-2022-41850](https://nvd.nist.gov/vuln/detail/CVE-2022-41850), [CVE-2022-41849](https://nvd.nist.gov/vuln/detail/CVE-2022-41849), [CVE-2022-41674](https://nvd.nist.gov/vuln/detail/CVE-2022-41674), [CVE-2022-4139](https://nvd.nist.gov/vuln/detail/CVE-2022-4139), [CVE-2022-4128](https://nvd.nist.gov/vuln/detail/CVE-2022-4128), [CVE-2022-41218](https://nvd.nist.gov/vuln/detail/CVE-2022-41218), [CVE-2022-40982](https://nvd.nist.gov/vuln/detail/CVE-2022-40982), [CVE-2022-4095](https://nvd.nist.gov/vuln/detail/CVE-2022-4095), [CVE-2022-40768](https://nvd.nist.gov/vuln/detail/CVE-2022-40768), [CVE-2022-40307](https://nvd.nist.gov/vuln/detail/CVE-2022-40307), [CVE-2022-40133](https://nvd.nist.gov/vuln/detail/CVE-2022-40133), [CVE-2022-3977](https://nvd.nist.gov/vuln/detail/CVE-2022-3977), [CVE-2022-39190](https://nvd.nist.gov/vuln/detail/CVE-2022-39190), [CVE-2022-39189](https://nvd.nist.gov/vuln/detail/CVE-2022-39189), [CVE-2022-3910](https://nvd.nist.gov/vuln/detail/CVE-2022-3910), [CVE-2022-38457](https://nvd.nist.gov/vuln/detail/CVE-2022-38457), [CVE-2022-3707](https://nvd.nist.gov/vuln/detail/CVE-2022-3707), [CVE-2022-36946](https://nvd.nist.gov/vuln/detail/CVE-2022-36946), [CVE-2022-36879](https://nvd.nist.gov/vuln/detail/CVE-2022-36879), [CVE-2022-3649](https://nvd.nist.gov/vuln/detail/CVE-2022-3649), [CVE-2022-3646](https://nvd.nist.gov/vuln/detail/CVE-2022-3646), [CVE-2022-3643](https://nvd.nist.gov/vuln/detail/CVE-2022-3643), [CVE-2022-3640](https://nvd.nist.gov/vuln/detail/CVE-2022-3640), [CVE-2022-3635](https://nvd.nist.gov/vuln/detail/CVE-2022-3635), [CVE-2022-3630](https://nvd.nist.gov/vuln/detail/CVE-2022-3630), [CVE-2022-3629](https://nvd.nist.gov/vuln/detail/CVE-2022-3629), [CVE-2022-36280](https://nvd.nist.gov/vuln/detail/CVE-2022-36280), [CVE-2022-3628](https://nvd.nist.gov/vuln/detail/CVE-2022-3628), [CVE-2022-3625](https://nvd.nist.gov/vuln/detail/CVE-2022-3625), [CVE-2022-3623](https://nvd.nist.gov/vuln/detail/CVE-2022-3623), [CVE-2022-3621](https://nvd.nist.gov/vuln/detail/CVE-2022-3621), [CVE-2022-3619](https://nvd.nist.gov/vuln/detail/CVE-2022-3619), [CVE-2022-36123](https://nvd.nist.gov/vuln/detail/CVE-2022-36123), [CVE-2022-3595](https://nvd.nist.gov/vuln/detail/CVE-2022-3595), [CVE-2022-3594](https://nvd.nist.gov/vuln/detail/CVE-2022-3594), [CVE-2022-3586](https://nvd.nist.gov/vuln/detail/CVE-2022-3586), [CVE-2022-3577](https://nvd.nist.gov/vuln/detail/CVE-2022-3577), [CVE-2022-3565](https://nvd.nist.gov/vuln/detail/CVE-2022-3565), [CVE-2022-3564](https://nvd.nist.gov/vuln/detail/CVE-2022-3564), [CVE-2022-3543](https://nvd.nist.gov/vuln/detail/CVE-2022-3543), [CVE-2022-3541](https://nvd.nist.gov/vuln/detail/CVE-2022-3541), [CVE-2022-3534](https://nvd.nist.gov/vuln/detail/CVE-2022-3534), [CVE-2022-3526](https://nvd.nist.gov/vuln/detail/CVE-2022-3526), [CVE-2022-3524](https://nvd.nist.gov/vuln/detail/CVE-2022-3524), [CVE-2022-3521](https://nvd.nist.gov/vuln/detail/CVE-2022-3521), [CVE-2022-34918](https://nvd.nist.gov/vuln/detail/CVE-2022-34918), [CVE-2022-34495](https://nvd.nist.gov/vuln/detail/CVE-2022-34495), [CVE-2022-34494](https://nvd.nist.gov/vuln/detail/CVE-2022-34494), [CVE-2022-3435](https://nvd.nist.gov/vuln/detail/CVE-2022-3435), [CVE-2022-3424](https://nvd.nist.gov/vuln/detail/CVE-2022-3424), [CVE-2022-33981](https://nvd.nist.gov/vuln/detail/CVE-2022-33981), [CVE-2022-33744](https://nvd.nist.gov/vuln/detail/CVE-2022-33744), [CVE-2022-33743](https://nvd.nist.gov/vuln/detail/CVE-2022-33743), [CVE-2022-33742](https://nvd.nist.gov/vuln/detail/CVE-2022-33742), [CVE-2022-33741](https://nvd.nist.gov/vuln/detail/CVE-2022-33741), [CVE-2022-33740](https://nvd.nist.gov/vuln/detail/CVE-2022-33740), [CVE-2022-3344](https://nvd.nist.gov/vuln/detail/CVE-2022-3344), [CVE-2022-3303](https://nvd.nist.gov/vuln/detail/CVE-2022-3303), [CVE-2022-32981](https://nvd.nist.gov/vuln/detail/CVE-2022-32981), [CVE-2022-3239](https://nvd.nist.gov/vuln/detail/CVE-2022-3239), [CVE-2022-32296](https://nvd.nist.gov/vuln/detail/CVE-2022-32296), [CVE-2022-32250](https://nvd.nist.gov/vuln/detail/CVE-2022-32250), [CVE-2022-3202](https://nvd.nist.gov/vuln/detail/CVE-2022-3202), [CVE-2022-3169](https://nvd.nist.gov/vuln/detail/CVE-2022-3169), [CVE-2022-3115](https://nvd.nist.gov/vuln/detail/CVE-2022-3115), [CVE-2022-3113](https://nvd.nist.gov/vuln/detail/CVE-2022-3113), [CVE-2022-3112](https://nvd.nist.gov/vuln/detail/CVE-2022-3112), [CVE-2022-3111](https://nvd.nist.gov/vuln/detail/CVE-2022-3111), [CVE-2022-3110](https://nvd.nist.gov/vuln/detail/CVE-2022-3110), [CVE-2022-3108](https://nvd.nist.gov/vuln/detail/CVE-2022-3108), [CVE-2022-3107](https://nvd.nist.gov/vuln/detail/CVE-2022-3107), [CVE-2022-3105](https://nvd.nist.gov/vuln/detail/CVE-2022-3105), [CVE-2022-3104](https://nvd.nist.gov/vuln/detail/CVE-2022-3104), [CVE-2022-3078](https://nvd.nist.gov/vuln/detail/CVE-2022-3078), [CVE-2022-3077](https://nvd.nist.gov/vuln/detail/CVE-2022-3077), [CVE-2022-30594](https://nvd.nist.gov/vuln/detail/CVE-2022-30594), [CVE-2022-3028](https://nvd.nist.gov/vuln/detail/CVE-2022-3028), [CVE-2022-29968](https://nvd.nist.gov/vuln/detail/CVE-2022-29968), [CVE-2022-29901](https://nvd.nist.gov/vuln/detail/CVE-2022-29901), [CVE-2022-29900](https://nvd.nist.gov/vuln/detail/CVE-2022-29900), [CVE-2022-2978](https://nvd.nist.gov/vuln/detail/CVE-2022-2978), [CVE-2022-2977](https://nvd.nist.gov/vuln/detail/CVE-2022-2977), [CVE-2022-2964](https://nvd.nist.gov/vuln/detail/CVE-2022-2964), [CVE-2022-2959](https://nvd.nist.gov/vuln/detail/CVE-2022-2959), [CVE-2022-29582](https://nvd.nist.gov/vuln/detail/CVE-2022-29582), [CVE-2022-29581](https://nvd.nist.gov/vuln/detail/CVE-2022-29581), [CVE-2022-2938](https://nvd.nist.gov/vuln/detail/CVE-2022-2938), [CVE-2022-29156](https://nvd.nist.gov/vuln/detail/CVE-2022-29156), [CVE-2022-2905](https://nvd.nist.gov/vuln/detail/CVE-2022-2905), [CVE-2022-28893](https://nvd.nist.gov/vuln/detail/CVE-2022-28893), [CVE-2022-28796](https://nvd.nist.gov/vuln/detail/CVE-2022-28796), [CVE-2022-2873](https://nvd.nist.gov/vuln/detail/CVE-2022-2873), [CVE-2022-28390](https://nvd.nist.gov/vuln/detail/CVE-2022-28390), [CVE-2022-28389](https://nvd.nist.gov/vuln/detail/CVE-2022-28389), [CVE-2022-28388](https://nvd.nist.gov/vuln/detail/CVE-2022-28388), [CVE-2022-28356](https://nvd.nist.gov/vuln/detail/CVE-2022-28356), [CVE-2022-27950](https://nvd.nist.gov/vuln/detail/CVE-2022-27950), [CVE-2022-2785](https://nvd.nist.gov/vuln/detail/CVE-2022-2785), [CVE-2022-27672](https://nvd.nist.gov/vuln/detail/CVE-2022-27672), [CVE-2022-27666](https://nvd.nist.gov/vuln/detail/CVE-2022-27666), [CVE-2022-27223](https://nvd.nist.gov/vuln/detail/CVE-2022-27223), [CVE-2022-26966](https://nvd.nist.gov/vuln/detail/CVE-2022-26966), [CVE-2022-2663](https://nvd.nist.gov/vuln/detail/CVE-2022-2663), [CVE-2022-26490](https://nvd.nist.gov/vuln/detail/CVE-2022-26490), [CVE-2022-2639](https://nvd.nist.gov/vuln/detail/CVE-2022-2639), [CVE-2022-26373](https://nvd.nist.gov/vuln/detail/CVE-2022-26373), [CVE-2022-26365](https://nvd.nist.gov/vuln/detail/CVE-2022-26365), [CVE-2022-2602](https://nvd.nist.gov/vuln/detail/CVE-2022-2602), [CVE-2022-2590](https://nvd.nist.gov/vuln/detail/CVE-2022-2590), [CVE-2022-2588](https://nvd.nist.gov/vuln/detail/CVE-2022-2588), [CVE-2022-2586](https://nvd.nist.gov/vuln/detail/CVE-2022-2586), [CVE-2022-2585](https://nvd.nist.gov/vuln/detail/CVE-2022-2585), [CVE-2022-25636](https://nvd.nist.gov/vuln/detail/CVE-2022-25636), [CVE-2022-25375](https://nvd.nist.gov/vuln/detail/CVE-2022-25375), [CVE-2022-25258](https://nvd.nist.gov/vuln/detail/CVE-2022-25258), [CVE-2022-2503](https://nvd.nist.gov/vuln/detail/CVE-2022-2503), [CVE-2022-24959](https://nvd.nist.gov/vuln/detail/CVE-2022-24959), [CVE-2022-24958](https://nvd.nist.gov/vuln/detail/CVE-2022-24958), [CVE-2022-24448](https://nvd.nist.gov/vuln/detail/CVE-2022-24448), [CVE-2022-23960](https://nvd.nist.gov/vuln/detail/CVE-2022-23960), [CVE-2022-2380](https://nvd.nist.gov/vuln/detail/CVE-2022-2380), [CVE-2022-23222](https://nvd.nist.gov/vuln/detail/CVE-2022-23222), [CVE-2022-2318](https://nvd.nist.gov/vuln/detail/CVE-2022-2318), [CVE-2022-2308](https://nvd.nist.gov/vuln/detail/CVE-2022-2308), [CVE-2022-23042](https://nvd.nist.gov/vuln/detail/CVE-2022-23042), [CVE-2022-23041](https://nvd.nist.gov/vuln/detail/CVE-2022-23041), [CVE-2022-23040](https://nvd.nist.gov/vuln/detail/CVE-2022-23040), [CVE-2022-23039](https://nvd.nist.gov/vuln/detail/CVE-2022-23039), [CVE-2022-23038](https://nvd.nist.gov/vuln/detail/CVE-2022-23038), [CVE-2022-23037](https://nvd.nist.gov/vuln/detail/CVE-2022-23037), [CVE-2022-23036](https://nvd.nist.gov/vuln/detail/CVE-2022-23036), [CVE-2022-22942](https://nvd.nist.gov/vuln/detail/CVE-2022-22942), [CVE-2022-2196](https://nvd.nist.gov/vuln/detail/CVE-2022-2196), [CVE-2022-2153](https://nvd.nist.gov/vuln/detail/CVE-2022-2153), [CVE-2022-21505](https://nvd.nist.gov/vuln/detail/CVE-2022-21505), [CVE-2022-21499](https://nvd.nist.gov/vuln/detail/CVE-2022-21499), [CVE-2022-21166](https://nvd.nist.gov/vuln/detail/CVE-2022-21166), [CVE-2022-21125](https://nvd.nist.gov/vuln/detail/CVE-2022-21125), [CVE-2022-21123](https://nvd.nist.gov/vuln/detail/CVE-2022-21123), [CVE-2022-2078](https://nvd.nist.gov/vuln/detail/CVE-2022-2078), [CVE-2022-20572](https://nvd.nist.gov/vuln/detail/CVE-2022-20572), [CVE-2022-20566](https://nvd.nist.gov/vuln/detail/CVE-2022-20566), [CVE-2022-20423](https://nvd.nist.gov/vuln/detail/CVE-2022-20423), [CVE-2022-20422](https://nvd.nist.gov/vuln/detail/CVE-2022-20422), [CVE-2022-20421](https://nvd.nist.gov/vuln/detail/CVE-2022-20421), [CVE-2022-20369](https://nvd.nist.gov/vuln/detail/CVE-2022-20369), [CVE-2022-20368](https://nvd.nist.gov/vuln/detail/CVE-2022-20368), [CVE-2022-20158](https://nvd.nist.gov/vuln/detail/CVE-2022-20158), [CVE-2022-20008](https://nvd.nist.gov/vuln/detail/CVE-2022-20008), [CVE-2022-1998](https://nvd.nist.gov/vuln/detail/CVE-2022-1998), [CVE-2022-1976](https://nvd.nist.gov/vuln/detail/CVE-2022-1976), [CVE-2022-1975](https://nvd.nist.gov/vuln/detail/CVE-2022-1975), [CVE-2022-1974](https://nvd.nist.gov/vuln/detail/CVE-2022-1974), [CVE-2022-1973](https://nvd.nist.gov/vuln/detail/CVE-2022-1973), [CVE-2022-1943](https://nvd.nist.gov/vuln/detail/CVE-2022-1943), [CVE-2022-1882](https://nvd.nist.gov/vuln/detail/CVE-2022-1882), [CVE-2022-1852](https://nvd.nist.gov/vuln/detail/CVE-2022-1852), [CVE-2022-1789](https://nvd.nist.gov/vuln/detail/CVE-2022-1789), [CVE-2022-1734](https://nvd.nist.gov/vuln/detail/CVE-2022-1734), [CVE-2022-1729](https://nvd.nist.gov/vuln/detail/CVE-2022-1729), [CVE-2022-1679](https://nvd.nist.gov/vuln/detail/CVE-2022-1679), [CVE-2022-1671](https://nvd.nist.gov/vuln/detail/CVE-2022-1671), [CVE-2022-1652](https://nvd.nist.gov/vuln/detail/CVE-2022-1652), [CVE-2022-1651](https://nvd.nist.gov/vuln/detail/CVE-2022-1651), [CVE-2022-1516](https://nvd.nist.gov/vuln/detail/CVE-2022-1516), [CVE-2022-1462](https://nvd.nist.gov/vuln/detail/CVE-2022-1462), [CVE-2022-1353](https://nvd.nist.gov/vuln/detail/CVE-2022-1353), [CVE-2022-1263](https://nvd.nist.gov/vuln/detail/CVE-2022-1263), [CVE-2022-1205](https://nvd.nist.gov/vuln/detail/CVE-2022-1205), [CVE-2022-1204](https://nvd.nist.gov/vuln/detail/CVE-2022-1204), [CVE-2022-1199](https://nvd.nist.gov/vuln/detail/CVE-2022-1199), [CVE-2022-1198](https://nvd.nist.gov/vuln/detail/CVE-2022-1198), [CVE-2022-1184](https://nvd.nist.gov/vuln/detail/CVE-2022-1184), [CVE-2022-1158](https://nvd.nist.gov/vuln/detail/CVE-2022-1158), [CVE-2022-1055](https://nvd.nist.gov/vuln/detail/CVE-2022-1055), [CVE-2022-1048](https://nvd.nist.gov/vuln/detail/CVE-2022-1048), [CVE-2022-1016](https://nvd.nist.gov/vuln/detail/CVE-2022-1016), [CVE-2022-1015](https://nvd.nist.gov/vuln/detail/CVE-2022-1015), [CVE-2022-1012](https://nvd.nist.gov/vuln/detail/CVE-2022-1012), [CVE-2022-1011](https://nvd.nist.gov/vuln/detail/CVE-2022-1011), [CVE-2022-0995](https://nvd.nist.gov/vuln/detail/CVE-2022-0995), [CVE-2022-0847](https://nvd.nist.gov/vuln/detail/CVE-2022-0847), [CVE-2022-0742](https://nvd.nist.gov/vuln/detail/CVE-2022-0742), [CVE-2022-0617](https://nvd.nist.gov/vuln/detail/CVE-2022-0617), [CVE-2022-0516](https://nvd.nist.gov/vuln/detail/CVE-2022-0516), [CVE-2022-0500](https://nvd.nist.gov/vuln/detail/CVE-2022-0500), [CVE-2022-0494](https://nvd.nist.gov/vuln/detail/CVE-2022-0494), [CVE-2022-0492](https://nvd.nist.gov/vuln/detail/CVE-2022-0492), [CVE-2022-0487](https://nvd.nist.gov/vuln/detail/CVE-2022-0487), [CVE-2022-0435](https://nvd.nist.gov/vuln/detail/CVE-2022-0435), [CVE-2022-0433](https://nvd.nist.gov/vuln/detail/CVE-2022-0433), [CVE-2022-0382](https://nvd.nist.gov/vuln/detail/CVE-2022-0382), [CVE-2022-0330](https://nvd.nist.gov/vuln/detail/CVE-2022-0330), [CVE-2022-0185](https://nvd.nist.gov/vuln/detail/CVE-2022-0185), [CVE-2022-0168](https://nvd.nist.gov/vuln/detail/CVE-2022-0168), [CVE-2022-0002](https://nvd.nist.gov/vuln/detail/CVE-2022-0002), [CVE-2022-0001](https://nvd.nist.gov/vuln/detail/CVE-2022-0001), [CVE-2021-45469](https://nvd.nist.gov/vuln/detail/CVE-2021-45469), [CVE-2021-44879](https://nvd.nist.gov/vuln/detail/CVE-2021-44879), [CVE-2021-43976](https://nvd.nist.gov/vuln/detail/CVE-2021-43976), [CVE-2021-4197](https://nvd.nist.gov/vuln/detail/CVE-2021-4197), [CVE-2021-4155](https://nvd.nist.gov/vuln/detail/CVE-2021-4155), [CVE-2021-3923](https://nvd.nist.gov/vuln/detail/CVE-2021-3923), [CVE-2021-33655](https://nvd.nist.gov/vuln/detail/CVE-2021-33655), [CVE-2021-33135](https://nvd.nist.gov/vuln/detail/CVE-2021-33135), [CVE-2021-26401](https://nvd.nist.gov/vuln/detail/CVE-2021-26401), [CVE-2020-36516](https://nvd.nist.gov/vuln/detail/CVE-2020-36516)))
- Go ([CVE-2023-39323](https://nvd.nist.gov/vuln/detail/CVE-2023-39323), [CVE-2023-39322](https://nvd.nist.gov/vuln/detail/CVE-2023-39322), [CVE-2023-39321](https://nvd.nist.gov/vuln/detail/CVE-2023-39321), [CVE-2023-39320](https://nvd.nist.gov/vuln/detail/CVE-2023-39320), [CVE-2023-39319](https://nvd.nist.gov/vuln/detail/CVE-2023-39319), [CVE-2023-39318](https://nvd.nist.gov/vuln/detail/CVE-2023-39318), [CVE-2023-29409](https://nvd.nist.gov/vuln/detail/CVE-2023-29409), [CVE-2023-29406](https://nvd.nist.gov/vuln/detail/CVE-2023-29406), [CVE-2023-29405](https://nvd.nist.gov/vuln/detail/CVE-2023-29405), [CVE-2023-29404](https://nvd.nist.gov/vuln/detail/CVE-2023-29404), [CVE-2023-29403](https://nvd.nist.gov/vuln/detail/CVE-2023-29403), [CVE-2023-29402](https://nvd.nist.gov/vuln/detail/CVE-2023-29402))
- OpenSSL ([CVE-2023-3446](https://nvd.nist.gov/vuln/detail/CVE-2023-3446), [CVE-2023-2975](https://nvd.nist.gov/vuln/detail/CVE-2023-2975), [CVE-2023-2650](https://nvd.nist.gov/vuln/detail/CVE-2023-2650))
- Python ([CVE-2023-41105](https://nvd.nist.gov/vuln/detail/CVE-2023-41105), [CVE-2023-40217](https://nvd.nist.gov/vuln/detail/CVE-2023-40217))
- SDK: Rust ([CVE-2023-38497](https://nvd.nist.gov/vuln/detail/CVE-2023-38497))
- VMware: open-vm-tools ([CVE-2023-20900](https://nvd.nist.gov/vuln/detail/CVE-2023-20900), [CVE-2023-20867](https://nvd.nist.gov/vuln/detail/CVE-2023-20867))
- binutils ([CVE-2023-1579](https://nvd.nist.gov/vuln/detail/CVE-2023-1579), [CVE-2022-4285](https://nvd.nist.gov/vuln/detail/CVE-2022-4285), [CVE-2022-38533](https://nvd.nist.gov/vuln/detail/CVE-2022-38533))
- c-ares ([CVE-2023-32067](https://nvd.nist.gov/vuln/detail/CVE-2023-32067), [CVE-2023-31147](https://nvd.nist.gov/vuln/detail/CVE-2023-31147), [CVE-2023-31130](https://nvd.nist.gov/vuln/detail/CVE-2023-31130), [CVE-2023-31124](https://nvd.nist.gov/vuln/detail/CVE-2023-31124))
- curl ([CVE-2023-38546](https://nvd.nist.gov/vuln/detail/CVE-2023-38546), [CVE-2023-38545](https://nvd.nist.gov/vuln/detail/CVE-2023-38545), [CVE-2023-38039](https://nvd.nist.gov/vuln/detail/CVE-2023-38039), [CVE-2023-28322](https://nvd.nist.gov/vuln/detail/CVE-2023-28322), [CVE-2023-28321](https://nvd.nist.gov/vuln/detail/CVE-2023-28321), [CVE-2023-28320](https://nvd.nist.gov/vuln/detail/CVE-2023-28320), [CVE-2023-28319](https://nvd.nist.gov/vuln/detail/CVE-2023-28319))
- git ([CVE-2023-29007](https://nvd.nist.gov/vuln/detail/CVE-2023-29007), [CVE-2023-25815](https://nvd.nist.gov/vuln/detail/CVE-2023-25815), [CVE-2023-25652](https://nvd.nist.gov/vuln/detail/CVE-2023-25652))
- glibc ([CVE-2023-4911](https://nvd.nist.gov/vuln/detail/CVE-2023-4911), [CVE-2023-4806](https://nvd.nist.gov/vuln/detail/CVE-2023-4806), [CVE-2023-4527](https://nvd.nist.gov/vuln/detail/CVE-2023-4527))
- go ([CVE-2023-39325](https://nvd.nist.gov/vuln/detail/CVE-2023-39325))
- grub ([CVE-2023-4693](https://nvd.nist.gov/vuln/detail/CVE-2023-4693), [CVE-2023-4692](https://nvd.nist.gov/vuln/detail/CVE-2023-4692), [CVE-2022-3775](https://nvd.nist.gov/vuln/detail/CVE-2022-3775), [CVE-2022-28737](https://nvd.nist.gov/vuln/detail/CVE-2022-28737), [CVE-2022-28736](https://nvd.nist.gov/vuln/detail/CVE-2022-28736), [CVE-2022-28735](https://nvd.nist.gov/vuln/detail/CVE-2022-28735), [CVE-2022-28734](https://nvd.nist.gov/vuln/detail/CVE-2022-28734), [CVE-2022-28733](https://nvd.nist.gov/vuln/detail/CVE-2022-28733), [CVE-2022-2601](https://nvd.nist.gov/vuln/detail/CVE-2022-2601), [CVE-2021-3981](https://nvd.nist.gov/vuln/detail/CVE-2021-3981), [CVE-2021-3697](https://nvd.nist.gov/vuln/detail/CVE-2021-3697), [CVE-2021-3696](https://nvd.nist.gov/vuln/detail/CVE-2021-3696), [CVE-2021-3695](https://nvd.nist.gov/vuln/detail/CVE-2021-3695), [CVE-2021-20233](https://nvd.nist.gov/vuln/detail/CVE-2021-20233), [CVE-2021-20225](https://nvd.nist.gov/vuln/detail/CVE-2021-20225), [CVE-2020-27779](https://nvd.nist.gov/vuln/detail/CVE-2020-27779), [CVE-2020-27749](https://nvd.nist.gov/vuln/detail/CVE-2020-27749), [CVE-2020-25647](https://nvd.nist.gov/vuln/detail/CVE-2020-25647), [CVE-2020-25632](https://nvd.nist.gov/vuln/detail/CVE-2020-25632), [CVE-2020-14372](https://nvd.nist.gov/vuln/detail/CVE-2020-14372), [CVE-2020-10713](https://nvd.nist.gov/vuln/detail/CVE-2020-10713))
- intel-microcode ([CVE-2023-23908](https://nvd.nist.gov/vuln/detail/CVE-2023-23908), [CVE-2022-41804](https://nvd.nist.gov/vuln/detail/CVE-2022-41804), [CVE-2022-40982](https://nvd.nist.gov/vuln/detail/CVE-2022-40982))
- libarchive ([libarchive-20230729](https://github.com/libarchive/libarchive/releases/tag/v3.7.1))
- libcap ([CVE-2023-2603](https://nvd.nist.gov/vuln/detail/CVE-2023-2603), [CVE-2023-2602](https://nvd.nist.gov/vuln/detail/CVE-2023-2602))
- libmicrohttpd ([CVE-2023-27371](https://nvd.nist.gov/vuln/detail/CVE-2023-27371))
- libtirpc ([libtirpc-rhbg-2224666](http://git.linux-nfs.org/?p=steved/libtirpc.git;a=commit;h=1d2e10afb2ffc35cb3623f57a15f712359f18e75), [libtirpc-rhbg-2150611](http://git.linux-nfs.org/?p=steved/libtirpc.git;a=commit;h=f7f0abdf267698de3f74a0285405b1b01f40893b), [libtirpc-rhbg-2138317](http://git.linux-nfs.org/?p=steved/libtirpc.git;a=commit;h=4a2d85c64110ee9e21a8c4f9dafd6b0ae621506d))
- libxml2 ([libxml2-20230428](https://gitlab.gnome.org/GNOME/libxml2/-/releases/v2.11.4))
- lua ([CVE-2022-33099](https://nvd.nist.gov/vuln/detail/CVE-2022-33099))
- mit-krb5 ([CVE-2023-36054](https://nvd.nist.gov/vuln/detail/CVE-2023-36054))
- ncurses ([CVE-2023-29491](https://nvd.nist.gov/vuln/detail/CVE-2023-29491))
- nvidia-drivers ([CVE-2023-25516](https://nvd.nist.gov/vuln/detail/CVE-2023-25516), [CVE-2023-25515](https://nvd.nist.gov/vuln/detail/CVE-2023-25515))
- openldap ([CVE-2023-2953](https://nvd.nist.gov/vuln/detail/CVE-2023-2953))
- procps ([CVE-2023-4016](https://nvd.nist.gov/vuln/detail/CVE-2023-4016))
- protobuf ([CVE-2022-1941](https://nvd.nist.gov/vuln/detail/CVE-2022-1941))
- qemu ([CVE-2023-2861](https://nvd.nist.gov/vuln/detail/CVE-2023-2861), [CVE-2023-0330](https://nvd.nist.gov/vuln/detail/CVE-2023-0330))
- samba ([CVE-2022-1615](https://nvd.nist.gov/vuln/detail/CVE-2022-1615), [CVE-2021-44142](https://nvd.nist.gov/vuln/detail/CVE-2021-44142))
- shadow ([CVE-2023-29383](https://nvd.nist.gov/vuln/detail/CVE-2023-29383))
- sudo ([CVE-2023-28487](https://nvd.nist.gov/vuln/detail/CVE-2023-28487), [CVE-2023-28486](https://nvd.nist.gov/vuln/detail/CVE-2023-28486), [CVE-2023-27320](https://nvd.nist.gov/vuln/detail/CVE-2023-27320))
- torcx ([CVE-2022-28948](https://nvd.nist.gov/vuln/detail/CVE-2022-28948))
- vim ([CVE-2023-2610](https://nvd.nist.gov/vuln/detail/CVE-2023-2610), [CVE-2023-2609](https://nvd.nist.gov/vuln/detail/CVE-2023-2609), [CVE-2023-2426](https://nvd.nist.gov/vuln/detail/CVE-2023-2426))


#### Bug fixes
 - AWS: Fixed the Amazon SSM agent that was crashing. ([Flatcar#1307](https://github.com/flatcar/Flatcar/issues/1307))
 - Added AWS EKS support for versions 1.24-1.28. Fixed `/usr/share/amazon/eks/download-kubelet.sh` to include download paths for these versions. ([scripts#1210](https://github.com/flatcar/scripts/pull/1210))
 - Fix the RemainAfterExit clause in nvidia.service ([Flatcar#1169](https://github.com/flatcar/Flatcar/issues/1169))
 - Fixed a bug resulting in coreos-cloudinit resetting the instance hostname to 'localhost' if no metadata could be found ([coreos-cloudinit#25](https://github.com/flatcar/coreos-cloudinit/pull/25), [Flatcar#1262](https://github.com/flatcar/Flatcar/issues/1262)), with contributions from [MichaelEischer](https://github.com/MichaelEischer)
 - Fixed bug in handling renamed network interfaces when generating login issue ([init#102](https://github.com/flatcar/init/pull/102))
 - Fixed iterating over the OEM update payload signatures which prevented the AWS OEM update to 3745.x.y ([update-engine#31](https://github.com/flatcar/update_engine/pull/31))
 - Fixed quotes handling for update-engine ([Flatcar#1209](https://github.com/flatcar/Flatcar/issues/1209))
 - Fixed supplying extension update payloads with a custom base URL in Nebraska ([Flatcar#1281](https://github.com/flatcar/Flatcar/issues/1281))
 - Fixed the missing `/etc/extensions/` symlinks for the inbuilt Docker/containerd systemd-sysext images on update from Beta 3760.1.0 ([update_engine#32](https://github.com/flatcar/update_engine/pull/32))
 - Fixed the postinstall hook failure when updating from Azure instances without OEM systemd-sysext images to Flatcar Alpha 3745.x.y ([update_engine#29](https://github.com/flatcar/update_engine/pull/29))
 - GCP: Fixed OS Login enabling ([scripts#1445](https://github.com/flatcar/scripts/pull/1445))
 - Made `sshkeys.service` more robust to only run `coreos-metadata-sshkeys@core.service` when not masked and also retry on failure ([init#112](https://github.com/flatcar/init/pull/112))

#### Changes
 - :warning: Dropped support for niftycloud and interoute. For interoute we haven't been generating the images for some time already. ([scripts#971](https://github.com/flatcar/scripts/pull/971)) :warning:
 - AWS OEM images now use a systemd-sysext image for layering additional platform-specific software on top of `/usr`
 - Added TLS Kernel module ([scripts#865](https://github.com/flatcar/scripts/pull/865))
 - Added support for multipart MIME userdata in coreos-cloudinit. Ignition now detects multipart userdata and delegates execution to coreos-cloudinit. (scripts#873)
 - Azure and QEMU OEM images now use systemd-sysext images for layering additional platform-specific software on top of `/usr`. For Azure images this also means that the image has a normal Python installation available through the sysext image. The OEM software is still not updated but this will be added soon.
 - Change nvidia.service to type oneshot (from the default "simple") so the subsequent services (configured with "Requires/After") are executed after the driver installation is successfully finished (flatcar/Flatcar#1136)
 - Enabled the virtio GPU driver ([scripts#830](https://github.com/flatcar/scripts/pull/830))
 - Migrate to Type=notify in containerd.service. Changed the unit to Type=notify, utilizing the existing containerd support for sd_notify call after socket setup.
 - Migrated the NVIDIA installer from the Azure/AWS OEM partition to `/usr` to make it available on all platforms ([scripts#932](https://github.com/flatcar/scripts/pull/932/), [Flatcar#1077](https://github.com/flatcar/Flatcar/issues/1077))
 - Moved a mountpoint of the OEM partition from `/usr/share/oem` to `/oem`. `/usr/share/oem` became a symlink to `/oem` for backward compatibility. Despite the move, the initrd images providing files through `/usr/share/oem` should keep using `/usr/share/oem`. The move was done to enable activating the OEM sysext images that are placed in the OEM partition.
 - OEM vendor tools are now A/B updated if they are shipped as systemd-sysext images, the migration happens when both partitions require a systemd-sysext OEM image - note that this will delete the `nvidia.service` from `/etc` on Azure because it's now part of `/usr` ([Flatcar#60](https://github.com/flatcar/Flatcar/issues/60))
 - Reworked the VMware OEM software to be shipped as A/B updated systemd-sysext image
 - SDK: Experimental support for [prefix builds](https://github.com/flatcar/scripts/blob/main/PREFIX.md) to create distro independent, portable, self-contained applications w/ all dependencies included. With contributions from [chewi](https://github.com/chewi) and [HappyTobi](https://github.com/HappyTobi).
 - Started shipping default ssh client and ssh daemon configs in `/etc/ssh/ssh_config` and `/etc/ssh/sshd_config` which include config snippets in `/etc/ssh/ssh_config.d` and `/etc/ssh/sshd_config.d`, respectively.
 - The open-vm-tools package in VMware OEM now comes with vmhgfs-fuse, udev rules, pam and vgauth
 - Updated locksmith to use non-deprecated resource control options in the systemd unit ([Locksmith#20](https://github.com/flatcar/locksmith/pull/20))

#### Updates
- Linux ([6.1.73](https://lwn.net/Articles/958343) (includes [6.1.72](https://lwn.net/Articles/957376), [6.1.71](https://lwn.net/Articles/957009), [6.1.70](https://lwn.net/Articles/956526), [6.1.69](https://lwn.net/Articles/955814), [6.1.68](https://lwn.net/Articles/954989), [6.1.67](https://lwn.net/Articles/954455), [6.1.66](https://lwn.net/Articles/954112), [6.1.65](https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tag/?h=v6.1.65), [6.1.64](https://lwn.net/Articles/953132), [6.1.63](https://lwn.net/Articles/952003), [6.1.62](https://lwn.net/Articles/950700), [6.1.61](https://lwn.net/Articles/949826), [6.1.60](https://lwn.net/Articles/948817), [6.1.59](https://lwn.net/Articles/948299), [6.1.58](https://lwn.net/Articles/947820), [6.1.57](https://lwn.net/Articles/947298), [6.1.56](https://lwn.net/Articles/946854), [6.1.55](https://lwn.net/Articles/945379), [6.1.54](https://git.kernel.org/pub/scm/linux/kernel/git/stable/linux.git/tag/?h=v6.1.54), [6.1.53](https://lwn.net/Articles/944358), [6.1.52](https://lwn.net/Articles/943754), [6.1.51](https://lwn.net/Articles/943403), [6.1.50](https://lwn.net/Articles/943112), [6.1.49](https://lwn.net/Articles/942880), [6.1.48](https://lwn.net/Articles/942865), [6.1.47](https://lwn.net/Articles/942531), [6.1.46](https://lwn.net/Articles/941774), [6.1.45](https://lwn.net/Articles/941273), [6.1.44](https://lwn.net/Articles/940800), [6.1.43](https://lwn.net/Articles/940338), [6.1.42](https://lwn.net/Articles/939423), [6.1.41](https://lwn.net/Articles/939103), [6.1.40](https://lwn.net/Articles/939015), [6.1.39](https://lwn.net/Articles/938619), [6.1.38](https://lwn.net/Articles/937403), [6.1.37](https://lwn.net/Articles/937082), [6.1.36](https://lwn.net/Articles/936674), [6.1.35](https://lwn.net/Articles/935588), [6.1.34](https://lwn.net/Articles/934623), [6.1.33](https://lwn.net/Articles/934319), [6.1.32](https://lwn.net/Articles/933908), [6.1.31](https://lwn.net/Articles/933281), [6.1.30](https://lwn.net/Articles/932882), [6.1.29](https://lwn.net/Articles/932133), [6.1.28](https://lwn.net/Articles/931651), [6.1.27](https://lwn.net/Articles/930597/), [6.1](https://kernelnewbies.org/Linux_6.1)))
- Linux Firmware ([20230919](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230919) (includes [20230804](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230804), [20230625](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230625), [20230515](https://git.kernel.org/pub/scm/linux/kernel/git/firmware/linux-firmware.git/tag/?h=20230515)))
- AWS: amazon-ssm-agent ([3.2.985.0](https://github.com/aws/amazon-ssm-agent/releases/tag/3.2.985.0))
- Go ([1.20.9](https://go.dev/doc/devel/release#go1.20.9) (includes [1.20.8](https://go.dev/doc/devel/release#go1.20.8), [1.20.7](https://go.dev/doc/devel/release#go1.20.7), [1.20.6](https://go.dev/doc/devel/release#go1.20.6), [1.20.5](https://go.dev/doc/devel/release#go1.20.5), [1.20.4](https://go.dev/doc/devel/release#go1.20.4), [1.20.10](https://go.dev/doc/devel/release#go1.20.10), [1.19.13](https://go.dev/doc/devel/release#go1.19.13), [1.19.12](https://go.dev/doc/devel/release#go1.19.12), [1.19.11](https://go.dev/doc/devel/release#go1.19.11), [1.19.10](https://go.dev/doc/devel/release#go1.19.10)))
- OpenSSL ([3.0.9](https://github.com/openssl/openssl/blob/openssl-3.0.9/NEWS.md#major-changes-between-openssl-308-and-openssl-309-30-may-2023))
- SDK: Rust ([1.72.1](https://github.com/rust-lang/rust/releases/tag/1.72.1) (includes [1.72.0](https://github.com/rust-lang/rust/releases/tag/1.72.0), [1.71.1](https://github.com/rust-lang/rust/releases/tag/1.71.1), [1.71.0](https://github.com/rust-lang/rust/releases/tag/1.71.0), [1.70.0](https://github.com/rust-lang/rust/releases/tag/1.70.0)))
- SDK: file ([5.45](https://github.com/file/file/blob/FILE5_45/ChangeLog))
- SDK: gnuconfig ([20230731](https://git.savannah.gnu.org/cgit/config.git/log/?id=d4e37b5868ef910e3e52744c34408084bb13051c))
- SDK: libxslt ([1.1.38](https://gitlab.gnome.org/GNOME/libxslt/-/releases/v1.1.38))
- SDK: man-db ([2.11.2](https://gitlab.com/man-db/man-db/-/tags/2.11.2))
- SDK: man-pages ([6.03](https://lore.kernel.org/lkml/d56662b2-538c-7252-9052-8afbf325f843@gmail.com/T/))
- SDK: pahole ([1.25](https://github.com/acmel/dwarves/blob/master/changes-v1.25))
- SDK: perf ([6.3](https://kernelnewbies.org/LinuxChanges#Linux_6.3.Tracing.2C_perf_and_BPF))
- SDK: perl ([5.36.1](https://perldoc.perl.org/perl5361delta))
- SDK: portage ([3.0.49](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.49) (includes [3.0.49](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.49), [3.0.46](https://gitweb.gentoo.org/proj/portage.git/tree/NEWS?h=portage-3.0.46)))
- SDK: python ([3.11.5](https://www.python.org/downloads/release/python-3115/) (includes [3.11.3](https://www.python.org/downloads/release/python-3113/), [3.10.12](https://www.python.org/downloads/release/python-31012/), [3.10.11](https://www.python.org/downloads/release/python-31011/)))
- SDK: qemu ([8.0.4](https://wiki.qemu.org/ChangeLog/8.0) (includes [8.0.3](https://wiki.qemu.org/ChangeLog/8.0), [7.2.3](https://wiki.qemu.org/ChangeLog/7.2)))
- SDK: qemu-guest-agent ([8.0.3](https://wiki.qemu.org/ChangeLog/8.0#Guest_agent) (includes [8.0.0](https://wiki.qemu.org/ChangeLog/8.0#Guest_agent)))
- VMWARE: libdnet ([1.16.2](https://github.com/ofalk/libdnet/releases/tag/libdnet-1.16.2) (includes [1.16](https://github.com/ofalk/libdnet/releases/tag/libdnet-1.16)))
- VMware: open-vm-tools ([12.3.0](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.3.0) (includes [12.2.5](https://github.com/vmware/open-vm-tools/releases/tag/stable-12.2.5)))
- XZ Utils ([5.4.3](https://git.tukaani.org/?p=xz.git;a=blob;f=NEWS;h=2f4d35adca6198671434d2988803cc9316ad1ec8;hb=dbb3a536ed9873ffa0870321f6873e564c6a9da8))
- afterburn ([5.5.0](https://github.com/coreos/afterburn/releases/tag/v5.5.0))
- bind-tools ([9.16.42](https://bind9.readthedocs.io/en/v9.16.42/notes.html#notes-for-bind-9-16-42) (includes [9.16.41](https://bind9.readthedocs.io/en/v9.16.41/notes.html#notes-for-bind-9-16-41)))
- binutils ([2.40](https://lists.gnu.org/archive/html/info-gnu/2023-01/msg00003.html))
- bpftool ([6.3](https://git.kernel.org/pub/scm/linux/kernel/git/bpf/bpf-next.git/log/tools/bpf/bpftool?h=v6.3))
- c-ares ([1.19.1](https://github.com/c-ares/c-ares/releases/tag/cares-1_19_1))
- cJSON ([1.7.16](https://github.com/DaveGamble/cJSON/releases/tag/v1.7.16))
- ca-certificates ([3.96.1](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_96_1.html) (includes [3.96](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_96.html)))
- checkpolicy ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- cifs-utils ([7.0](https://lists.samba.org/archive/samba-technical/2022-August/137528.html))
- containerd ([1.7.7](https://github.com/containerd/containerd/releases/tag/v1.7.7) (includes [1.7.6](https://github.com/containerd/containerd/releases/tag/v1.7.6), [1.7.5](https://github.com/containerd/containerd/releases/tag/v1.7.5), [1.7.4](https://github.com/containerd/containerd/releases/tag/v1.7.4), [1.7.3](https://github.com/containerd/containerd/releases/tag/v1.7.3), [1.7.2](https://github.com/containerd/containerd/releases/tag/v1.7.2)))
- coreutils ([9.3](https://lists.gnu.org/archive/html/info-gnu/2023-04/msg00006.html) (includes [9.1](https://git.savannah.gnu.org/cgit/coreutils.git/tree/NEWS?h=v9.1)))
- cryptsetup ([2.6.1](https://gitlab.com/cryptsetup/cryptsetup/-/blob/v2.6.1/docs/v2.6.1-ReleaseNotes) (includes [2.6.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/v2.6.0/docs/v2.6.0-ReleaseNotes), [2.5.0](https://gitlab.com/cryptsetup/cryptsetup/-/blob/v2.5.0/docs/v2.5.0-ReleaseNotes)))
- curl ([8.4.0](https://curl.se/changes.html#8_4_0) (includes [8.3.0](https://curl.se/changes.html#8_3_0), [8.2.1](https://curl.se/changes.html#8_2_1), [8.2.0](https://curl.se/changes.html#8_2_0), [8.1.2](https://curl.se/changes.html#8_1_2), [8.1.0](https://curl.se/changes.html#8_1_0)))
- debianutils ([5.7](https://metadata.ftp-master.debian.org/changelogs//main/d/debianutils/debianutils_5.7-0.4_changelog))
- diffutils ([3.10](https://lists.gnu.org/archive/html/info-gnu/2023-05/msg00009.html))
- elfutils ([0.189](https://sourceware.org/pipermail/elfutils-devel/2023q1/006023.html))
- ethtool ([6.4](https://git.kernel.org/pub/scm/network/ethtool/ethtool.git/tree/NEWS?h=v6.4) (includes [6.3](https://git.kernel.org/pub/scm/network/ethtool/ethtool.git/commit/?id=7bdf78f0d2a9ae1571fe9444e552490130e573fd), [6.2](https://git.kernel.org/pub/scm/network/ethtool/ethtool.git/tree/NEWS?h=v6.2)))
- gawk ([5.2.2](https://lists.gnu.org/archive/html/info-gnu/2023-05/msg00008.html))
- gcc ([13.2](https://gcc.gnu.org/gcc-13/changes.html))
- gdb ([13.2](https://lists.gnu.org/archive/html/info-gnu/2023-05/msg00011.html))
- gdbm ([1.23](https://lists.gnu.org/archive/html/info-gnu/2022-02/msg00004.html))
- git ([2.41.0](https://lore.kernel.org/git/xmqqleh3a3wm.fsf@gitster.g/) (includes [2.39.3](https://github.com/git/git/blob/v2.39.3/Documentation/RelNotes/2.39.3.txt)))
- glib ([2.76.4](https://gitlab.gnome.org/GNOME/glib/-/releases/2.76.4) (includes [2.76.3](https://gitlab.gnome.org/GNOME/glib/-/releases/2.76.3), [2.76.2](https://gitlab.gnome.org/GNOME/glib/-/releases/2.76.2)))
- glibc ([2.37](https://sourceware.org/git/?p=glibc.git;a=tag;h=refs/tags/glibc-2.37))
- gmp ([6.3.0](https://gmplib.org/gmp6.3))
- gptfdisk ([1.0.9](https://sourceforge.net/p/gptfdisk/code/ci/1d46f3723bc25f5598266f7d9a3548af3cee0c77/tree/NEWS))
- grep ([3.8](http://savannah.gnu.org/forum/forum.php?forum_id=10227) (includes [3.11](https://lists.gnu.org/archive/html/info-gnu/2023-05/msg00004.html)))
- grub ([2.06](https://lists.gnu.org/archive/html/grub-devel/2021-06/msg00022.html))
- gzip ([1.13](https://savannah.gnu.org/news/?id=10501))
- hwdata ([0.373](https://github.com/vcrhonek/hwdata/commits/v0.373) (includes [0.372](https://github.com/vcrhonek/hwdata/commits/v0.372), [0.371](https://github.com/vcrhonek/hwdata/commits/v0.371), [0.367](https://github.com/vcrhonek/hwdata/releases/tag/v0.367)))
- inih ([57](https://github.com/benhoyt/inih/releases/tag/r57) (includes [56](https://github.com/benhoyt/inih/releases/tag/r56)))
- intel-microcode ([20230808](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20230808) (includes [20230613](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20230613), [20230512](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20230512)))
- iperf ([3.14](https://github.com/esnet/iperf/blob/master/RELNOTES.md#iperf-314-2023-07-07))
- iproute2 ([6.4.0](https://git.kernel.org/pub/scm/network/iproute2/iproute2.git/log/?h=v6.4.0) (includes [6.3.0](https://lwn.net/Articles/930473/), [6.2](https://lwn.net/Articles/923952/)))
- ipset ([7.17](https://git.netfilter.org/ipset/tree/ChangeLog?id=186f9b57c60bb53aae5f6633eff1e9d5e9095c3e))
- kbd ([2.6.1](https://github.com/legionus/kbd/releases/tag/v2.6.1) (includes [2.6.0](https://github.com/legionus/kbd/releases/tag/v2.6.0), [2.5.1](https://github.com/legionus/kbd/releases/tag/v2.5.1)))
- kexec-tools ([2.0.24](https://github.com/horms/kexec-tools/releases/tag/v2.0.24))
- kmod ([30](https://lwn.net/Articles/899526/))
- ldb ([2.4.4](https://gitlab.com/samba-team/samba/-/commit/b686ef00da46d4a0c0aba0c61b1866cbc9b462b6) (includes [2.4.3](https://gitlab.com/samba-team/samba/-/commit/604f94704f30e90ef960aa2be62a14d2e614a002), [2.4.2](https://gitlab.com/samba-team/samba/-/commit/d93892d2e8ed69758c15ab18bc03bba09e715bc6)))
- less ([633](http://www.greenwoodsoftware.com/less/news.633.html) (includes [632](http://www.greenwoodsoftware.com/less/news.632.html)))
- libarchive ([3.7.1](https://github.com/libarchive/libarchive/releases/tag/v3.7.1) (includes [3.7.0](https://github.com/libarchive/libarchive/releases/tag/v3.7.0)))
- libassuan ([2.5.6](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=libassuan.git;a=blob;f=NEWS;h=e52bb5dd36ac93ea227e53e89f82af9ccf38f339;hb=6b50ee6bcdd6aa81bd7cc3fb2379864c3ed479b8))
- libbsd ([0.11.7](https://lists.freedesktop.org/archives/libbsd/2022-October/000337.html))
- libcap ([2.69](https://sites.google.com/site/fullycapable/release-notes-for-libcap#h.iuvg7sbjg8pe))
- libgcrypt ([1.10.2](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=libgcrypt.git;a=blob;f=NEWS;h=c9a239615f8070427a96688b1be40a81e59e9b8a;hb=1c5cbacf3d88dded5063e959ee68678ff7d0fa56) (includes [1.10.1](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=libgcrypt.git;a=blob;f=NEWS;h=03132c2a115e35783a782c64777cf5f5b1a2825f;hb=ae0e567820c37f9640440b3cff77d7c185aa6742)))
- libgpg-error ([1.47](https://dev.gnupg.org/T6231) (includes [1.46](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=libgpg-error.git;a=blob;f=NEWS;h=14b0ba97d6ba2b10b3178f2e4a3e24bfc2355bb3;hb=ea031873aa9642831017937fd33e9009d514ee07)))
- libksba ([1.6.4](https://git.gnupg.org/cgi-bin/gitweb.cgi?p=libksba.git;a=blob;f=NEWS;h=f640523209c1c9ce9855040e53914a79d24d6a67;hb=557999424ebd13e70d6fc17e648a5dd2a06f440b))
- libmd ([1.1.0](https://git.hadrons.org/cgit/libmd.git/log/?h=1.1.0))
- libmicrohttpd ([0.9.77](https://gitlab.com/libmicrohttpd/libmicrohttpd/-/releases/v0.9.77) (includes [0.9.76](https://lists.gnu.org/archive/html/libmicrohttpd/2023-02/msg00000.html)))
- libnftnl ([1.2.6](https://git.netfilter.org/libnftnl/log/?h=libnftnl-1.2.6) (includes [1.2.5](https://git.netfilter.org/libnftnl/log/?h=libnftnl-1.2.5)))
- libnl ([3.8.0](https://github.com/thom311/libnl/compare/libnl3_7_0...libnl3_8_0))
- libnvme ([1.5](https://github.com/linux-nvme/libnvme/releases/tag/v1.5))
- libpcap ([1.10.4](https://github.com/the-tcpdump-group/libpcap/blob/24832dd2728bd95ed9b9464ef27b47a943c38003/CHANGES#L51))
- libpcre ([8.45](https://www.pcre.org/original/changelog.txt))
- libpipeline ([1.5.7](https://gitlab.com/libpipeline/libpipeline/-/tags/1.5.7))
- libselinux ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- libsemanage ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- libsepol ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- libtirpc ([1.3.4](https://marc.info/?l=linux-nfs&m=169667640909830&w=2))
- libusb ([1.0.26](https://github.com/libusb/libusb/blob/v1.0.26/ChangeLog))
- libuv ([1.46.0](https://github.com/libuv/libuv/releases/tag/v1.46.0) (includes [1.45.0](https://github.com/libuv/libuv/releases/tag/v1.45.0)))
- libxml2 ([2.11.5](https://gitlab.gnome.org/GNOME/libxml2/-/releases/v2.11.5) (includes [2.11.4](https://gitlab.gnome.org/GNOME/libxml2/-/releases/v2.11.4)))
- lsof ([4.98.0](https://github.com/lsof-org/lsof/blob/4.98.0/00DIST#L5471))
- lua ([5.4.6](https://www.lua.org/manual/5.4/readme.html#changes) (includes [5.4.4](https://www.lua.org/manual/5.4/readme.html#changes)))
- mit-krb5 ([1.21.2](http://web.mit.edu/kerberos/krb5-1.21/))
- multipath-tools ([0.9.5](https://github.com/opensvc/multipath-tools/commits/0.9.5))
- ncurses ([6.4](https://invisible-island.net/ncurses/announce.html#h2-release-notes))
- nettle ([3.9.1](https://git.lysator.liu.se/nettle/nettle/-/blob/nettle_3.9.1_release_20230601/ChangeLog))
- nmap ([7.94](https://nmap.org/changelog.html#7.94))
- nvidia-drivers ([535.104.05](https://docs.nvidia.com/datacenter/tesla/tesla-release-notes-535-104-05/index.html))
- nvme-cli ([2.5](https://github.com/linux-nvme/nvme-cli/releases/tag/v2.5) (includes [2.3](https://github.com/linux-nvme/nvme-cli/releases/tag/v2.3)))
- open-isns ([0.102](https://github.com/open-iscsi/open-isns/blob/v0.102/ChangeLog))
- openldap ([2.6.4](https://git.openldap.org/openldap/openldap/-/blob/OPENLDAP_REL_ENG_2_6_4/CHANGES) (includes [2.6.3](https://lists.openldap.org/hyperkitty/list/openldap-announce@openldap.org/thread/FQJM2JSSSOMLQH7XC7Q5IZJYOGCTV2LK/), [2.6](https://lists.openldap.org/hyperkitty/list/openldap-announce@openldap.org/thread/IHS5V46H6NFNFUERMC6AWMPHTWRVNLFA/), [2.5.14](https://lists.openldap.org/hyperkitty/list/openldap-announce@openldap.org/thread/TZQHR4SIWUA5BZTKDAKSFDOOGDVU4TU7/), [2.5](https://lists.openldap.org/hyperkitty/list/openldap-announce@openldap.org/thread/BH3VDPG6IYYF5L5U6LZGHHKMJY5HFA3L/)))
- openssh ([9.5p1](https://www.openssh.com/releasenotes.html#9.5p1) (includes [9.4p1](https://www.openssh.com/releasenotes.html#9.4p1)))
- parted ([3.6](https://git.savannah.gnu.org/gitweb/?p=parted.git;a=blob;f=NEWS;h=52bb11697039f70e55120c571750f9ee761a75aa;hb=3b5f327b213d21e9adb9ba933c78dd898fee5b1d))
- pax-utils ([1.3.7](https://gitweb.gentoo.org/proj/pax-utils.git/log/?h=v1.3.7))
- pciutils ([3.9.0](https://github.com/pciutils/pciutils/releases/tag/v3.9.0) (includes [3.10.0](https://github.com/pciutils/pciutils/blob/v3.10.0/ChangeLog)))
- pigz ([2.8](https://zlib.net/pipermail/pigz-announce_zlib.net/2023-August/000018.html))
- policycoreutils ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- popt ([1.19](https://github.com/rpm-software-management/popt/releases/tag/popt-1.19-release))
- procps ([4.0.4](https://gitlab.com/procps-ng/procps/-/releases/v4.0.4) (includes [4.0.3](https://gitlab.com/procps-ng/procps/-/releases/v4.0.3), [4.0.0](https://gitlab.com/procps-ng/procps/-/releases/v4.0.0)))
- protobuf ([21.9](https://github.com/protocolbuffers/protobuf/releases/tag/v21.9))
- psmisc ([23.6](https://gitlab.com/psmisc/psmisc/-/blob/v23.6/ChangeLog))
- quota ([4.09](https://sourceforge.net/p/linuxquota/code/ci/87d2fd7635e4bca54fa2a00b8d5b073ba9ca521b/tree/Changelog))
- rpcsvc-proto ([1.4.4](https://github.com/thkukuk/rpcsvc-proto/releases/tag/v1.4.4))
- runc ([1.1.9](https://github.com/opencontainers/runc/releases/tag/v1.1.9) (includes [1.1.8](https://github.com/opencontainers/runc/releases/tag/v1.1.8)))
- samba ([4.18.4](https://wiki.samba.org/index.php/Samba_4.18_Features_added/changed#Samba_4.18.4))
- sed ([4.9](https://lists.gnu.org/archive/html/info-gnu/2022-11/msg00001.html))
- selinux-base ([2.20221101](https://github.com/SELinuxProject/refpolicy/releases/tag/RELEASE_2_20221101))
- selinux-base-policy ([2.20221101](https://github.com/SELinuxProject/refpolicy/releases/tag/RELEASE_2_20221101))
- selinux-container ([2.20221101](https://github.com/SELinuxProject/refpolicy/releases/tag/RELEASE_2_20221101))
- selinux-sssd ([2.20221101](https://github.com/SELinuxProject/refpolicy/releases/tag/RELEASE_2_20221101))
- selinux-unconfined ([2.20221101](https://github.com/SELinuxProject/refpolicy/releases/tag/RELEASE_2_20221101))
- semodule-utils ([3.5](https://github.com/SELinuxProject/selinux/releases/tag/3.5))
- smartmontools ([7.3](https://github.com/smartmontools/smartmontools/releases/tag/RELEASE_7_3))
- sqlite ([3.42.0](https://sqlite.org/releaselog/3_42_0.html))
- strace ([6.4](https://github.com/strace/strace/releases/tag/v6.4) (includes [6.3](https://github.com/strace/strace/releases/tag/v6.3), [6.2](https://github.com/strace/strace/releases/tag/v6.2)))
- sudo ([1.9.13p3](https://www.sudo.ws/releases/stable/#1.9.13p3))
- talloc ([2.4.0](https://gitlab.com/samba-team/samba/-/commit/5224ed98eeba43f22b5f5f87de5947fbb1c1c7c1) (includes [2.3.4](https://gitlab.com/samba-team/samba/-/commit/0189ccf9fc3d2a77cc83cffe180e307bcdccebb4)))
- tar ([1.35](https://lists.gnu.org/archive/html/info-gnu/2023-07/msg00005.html))
- tdb ([1.4.8](https://gitlab.com/samba-team/samba/-/commit/eab796a4f9172e602dc262f3c99ead35b35929e7) (includes [1.4.7](https://gitlab.com/samba-team/samba/-/commit/27ceb1c3ad786386e746a5e2968780d791393b9e), [1.4.6](https://gitlab.com/samba-team/samba/-/commit/1c776e54cf33b46b2ed73263f093d596a0cdbb2f)))
- tevent ([0.14.1](https://gitlab.com/samba-team/samba/-/commit/d80f28b081e515e32a480daf80b42cf782447a9c) (includes [0.14.0](https://gitlab.com/samba-team/samba/-/commit/3c6d28ebae27dba8e40558ae37ae8138ea0b4bdc), [0.13.0](https://gitlab.com/samba-team/samba/-/commit/63d4db63feda920c8020f8484a8b31065b7f1380), [0.12.1](https://gitlab.com/samba-team/samba/-/commit/53692735c733d01acbd953641f831a1f5e0cf6c5), [0.12.0](https://gitlab.com/samba-team/samba/-/tags/tevent-0.12.0)))
- usbutils ([015](https://github.com/gregkh/usbutils/blob/79b796f945ea7d5c2b0e2a74f9b8819cb7948680/NEWS))
- userspace-rcu ([0.14.0](https://github.com/urcu/userspace-rcu/blob/v0.13.2/ChangeLog))
- util-linux ([2.38.1](https://github.com/util-linux/util-linux/releases/tag/v2.38.1))
- vim ([9.0.1678](https://github.com/vim/vim/commits/v9.0.1678) (includes [9.0.1677](https://github.com/vim/vim/commits/v9.0.1677), [9.0.1503](https://github.com/vim/vim/commits/v9.0.1503)))
- wget ([1.21.4](https://lists.gnu.org/archive/html/info-gnu/2023-05/msg00003.html))
- whois ([5.5.18](https://github.com/rfc1036/whois/blob/v5.5.18/debian/changelog) (includes [5.5.17](https://github.com/rfc1036/whois/commit/bac7108b01cfd54c517444efa1239e10e6edd5a4)))
- xfsprogs ([6.4.0](https://git.kernel.org/pub/scm/fs/xfs/xfsprogs-dev.git/tree/doc/CHANGES?h=v6.4.0) (includes [6.3.0](https://git.kernel.org/pub/scm/fs/xfs/xfsprogs-dev.git/tree/doc/CHANGES?h=v6.3.0)))
- zstd ([1.5.5](https://github.com/facebook/zstd/releases/tag/v1.5.5))


**Changes since Beta-3760.1.1**

 #### Security fixes:
 
 - Linux ([CVE-2023-1193](https://nvd.nist.gov/vuln/detail/CVE-2023-1193), [CVE-2023-51779](https://nvd.nist.gov/vuln/detail/CVE-2023-51779), [CVE-2023-51780](https://nvd.nist.gov/vuln/detail/CVE-2023-51780), [CVE-2023-51781](https://nvd.nist.gov/vuln/detail/CVE-2023-51781), [CVE-2023-51782](https://nvd.nist.gov/vuln/detail/CVE-2023-51782), [CVE-2023-6531](https://nvd.nist.gov/vuln/detail/CVE-2023-6531), [CVE-2023-6606](https://nvd.nist.gov/vuln/detail/CVE-2023-6606), [CVE-2023-6622](https://nvd.nist.gov/vuln/detail/CVE-2023-6622), [CVE-2023-6817](https://nvd.nist.gov/vuln/detail/CVE-2023-6817), [CVE-2023-6931](https://nvd.nist.gov/vuln/detail/CVE-2023-6931))
 
 #### Bug fixes:
 
 - AWS: Fixed the Amazon SSM agent that was crashing. ([Flatcar#1307](https://github.com/flatcar/Flatcar/issues/1307))
 - Fixed a bug resulting in coreos-cloudinit resetting the instance hostname to 'localhost' if no metadata could be found ([coreos-cloudinit#25](https://github.com/flatcar/coreos-cloudinit/pull/25), [Flatcar#1262](https://github.com/flatcar/Flatcar/issues/1262)), with contributions from [MichaelEischer](https://github.com/MichaelEischer)
 - Fixed supplying extension update payloads with a custom base URL in Nebraska ([Flatcar#1281](https://github.com/flatcar/Flatcar/issues/1281))
 

#### Updates
- Linux ([6.1.73](https://lwn.net/Articles/958343) (includes [6.1.72](https://lwn.net/Articles/957376), [6.1.71](https://lwn.net/Articles/957009), [6.1.70](https://lwn.net/Articles/956526), [6.1.69](https://lwn.net/Articles/955814), [6.1.68](https://lwn.net/Articles/954989), [6.1.67](https://lwn.net/Articles/954455)))
- ca-certificates ([3.96.1](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_96_1.html) (includes [3.96](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_96.html)))

- Upgrade VPA components to 1.0.0


### kubernetes [1.25.16](https://github.com/kubernetes/kubernetes/releases/tag/v1.25.16)

> **DISCLAIMER:** Summary of upstream Kubernetes release notes generated with the help of ChatGPT

Kubernetes version 1.25 introduces numerous enhancements and features to improve the stability, scalability, and security of the platform. Some highlights include:

1. *Security Deep Dive*
- *Pod Security Admission (PSA):* PSA offers significant control and customization for security (Pod Security Polices are removed from the API):
  - *Hierarchical Namespaces:* You can structure how policies are applied across your cluster's namespaces.
  - *Three Levels of Enforcement:* Offers "enforce", "audit", and "warn" modes, giving you flexibility during migration and testing.
  - *Flexible Policy Definition:* Policies are defined within Kubernetes itself, providing easier integration and management.
  - *Migration Notes:* If you heavily rely on PSPs, investigate built-in replacement policies or consider developing custom PSA policies for complex requirements.
2. *Storage Deep Dive*
  - *CSI Ephemeral Volume Support:Pod-level Temporary Storage:* Ideal for scratch space needed within the Pod's lifecycle (caches, build artifacts, etc.).
  - *Volume Data Stays Local:* Data is written to the host node's filesystem, not a persistent storage backend.
  - *CSI Drivers:* The specific implementation may vary depending on your CSI driver's capabilities.
  - *Volume Expansion:Online Resizing: You can change volume size while Pods are still using them.
  - *Restrictions:* Not all storage providers may support this, and you may need to update your CSI drivers or storage classes.
3. *Networking Deep Dive*
  - *Local Ephemeral Storage Capacity Isolation:Protects Node Resources:* Helps prevent Pods with ephemeral storage from overwhelming a node's storage capacity.
  - *Improved Node Stability:* Ensures critical system components and other workloads aren't starved for storage resources.
4 *Additional Details*
  - *cgroups v2 Support:Modern cgroups:* Improved resource accounting and control for container workloads, especially important in resource-constrained environments.
  - *Important Note:* You may need to make changes to your cluster configuration if you were customizing cgroup settings for version 1.
  - *CRD Validation:Enforce Data Quality:* Ensures the data stored in Custom Resources adheres to your defined rules (structure, ranges, etc.).
  - *Beta Status:* While useful, some limitations might exist, so carefully test and observe during use.

*Important to Remember:*
Experimental Features: Version 1.25 likely contains other experimental features. Proceed with caution if you choose to implement them in production.
Tooling Impact Certain tools and integrations you use with your Kubernetes-as-a-Service might require updates to be compatible with the newer APIs and removals introduced in this version.

These are just a few of the key highlights of Kubernetes 1.25. The release includes many other enhancements, bug fixes, and updates to existing features. 

Kubernetes version 1.25 to 1.25.16 includes several enhancements, bug fixes, and stability improvements. Some notable changes in this release range include:

1. Stability improvements: Various enhancements have been made to improve the stability and reliability of the Kubernetes platform, addressing issues reported by users and community contributors.
2. Bug fixes: Numerous bugs have been fixed across different components of Kubernetes, including issues related to networking, storage, scheduling, and authentication.
3. Performance optimizations: Efforts have been made to optimize the performance of Kubernetes, resulting in improved resource utilization and responsiveness.
4. Security enhancements: Several security-related enhancements have been implemented to strengthen the overall security posture of Kubernetes, addressing vulnerabilities and improving security features.
5. API updates: Updates to the Kubernetes API have been introduced, including new features, improvements, and changes to existing resources and endpoints.
6. Deprecations and removals: Certain features, APIs, or components may have been deprecated or removed in this release, with guidance provided for users to adapt their configurations accordingly.
7. Documentation improvements: The documentation has been updated and expanded to provide more comprehensive guidance for users, administrators, and developers working with Kubernetes.

Overall, Kubernetes 1.25 to 1.25.16 brings a range of improvements to enhance the stability, performance, security, and usability of the platform for users and administrators.

Users are encouraged to review the [detailed release notes](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.25.md) for a comprehensive overview of the changes and to ensure compatibility with their existing deployments.
