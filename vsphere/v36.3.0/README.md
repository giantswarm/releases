# :zap: Giant Swarm Release v36.3.0 for vSphere :zap:

## Changes compared to v36.2.0

### Components

- cluster-vsphere from v11.4.0 to v11.6.0
- cluster from v9.4.0 to v9.6.0
- Flatcar from v4757.2.2 to [v4757.2.3](https://www.flatcar.org/releases/#release-4757.2.3)
- Kubernetes from v1.36.5 to [v1.36.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.6)

### cluster-vsphere [v11.4.0...v11.6.0](https://github.com/giantswarm/cluster-vsphere/compare/v11.4.0...v11.6.0)

#### Changed

- Update architect to v10.17.4 (giantswarm/cluster-vsphere#632)

### cluster [v9.4.0...v9.6.0](https://github.com/giantswarm/cluster/compare/v9.4.0...v9.6.0)

#### Added

- Apps: Add `cluster.app.version.satisfies` helper to gate app values on a semver constraint of the deployed app version, including release candidates, dev builds and version overrides.
- Control Plane: Add `internal.advancedConfiguration.controlPlane.apiServer.enableAggregatorRouting` to set `--enable-aggregator-routing` on KubeadmControlPlane (disabled by default).

#### Changed

- Apps: Nest net-exporter `DNSCheck` and `Hosts` default values under `NetExporter` to match the chart's values schema.
- Control Plane: Harden kube-scheduler on KubeadmControlPlane and KamajiControlPlane following the [Kubernetes scheduler hardening guide](https://kubernetes.io/docs/concepts/security/hardening-guide/scheduler/). Set `--profiling=false`, `--tls-cipher-suites` (same list as kube-apiserver), `--authentication-skip-lookup=false`, `--authentication-tolerate-lookup-failure=false`, `--permit-address-sharing=false` and `--permit-port-sharing=false`. The scheduler now fails to start if it cannot look up its authentication configuration from the API server. `--bind-address=0.0.0.0` and `/metrics` in `--authorization-always-allow-paths` are kept for metrics scraping.

#### Fixed

- Apps: cert-manager release candidates and dev builds of v4 and newer get the v4 values.

### Apps

- cert-exporter from v2.12.1 to v2.14.0
- cert-manager from v4.1.2 to v4.3.0
- cilium from v1.6.2 to v1.8.1
- etcd-defrag from v1.2.13 to v1.2.14
- k8s-dns-node-cache from v2.11.1 to v2.11.3
- net-exporter from v1.24.1 to v1.25.1
- network-policies from v0.5.0 to v0.6.0
- node-exporter from v1.21.0 to v1.21.1
- observability-bundle from v3.6.0 to v3.8.1
- prometheus-blackbox-exporter from v0.10.0 to v0.11.0
- security-bundle from v2.4.0 to v2.5.0
- teleport-kube-agent from v0.12.0 to v0.13.0
- vsphere-csi-driver from v3.4.6 to v3.5.0

### cert-exporter [v2.12.1...v2.14.0](https://github.com/giantswarm/cert-exporter/compare/v2.12.1...v2.14.0)

#### Added

- `daemonset.updateStrategy` Helm value to configure the DaemonSet update strategy.
- `fingerprint` label (first 16 hex characters of the certificate's SHA-256 fingerprint) on `cert_exporter_not_after` and `cert_exporter_secret_not_after`.

#### Changed

- The DaemonSet rolling update replaces up to 10% of the pods at a time by default (previously one pod at a time), so the rollout finishes within the Helm timeout on large workload clusters.
- `/metrics` now logs gather and encoding errors (such as a dropped duplicate series) instead of discarding them silently.
- A truly identical certificate duplicated within one secret key or file is reported once, without producing a gather error on every scrape.

#### Fixed

- Certificates sharing a serial number in one secret key or file no longer collide into one series. Previously only the first, expired certificate was exported.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### cert-manager [v4.1.2...v4.3.0](https://github.com/giantswarm/cert-manager-app/compare/v4.1.2...v4.3.0)

#### Changed

- The ACME HTTP-01 solver of the `letsencrypt-giantswarm` ClusterIssuer now uses Gateway API (`gatewayHTTPRoute`) instead of an Ingress, and targets Envoy Gateway.

#### Removed

- `giantSwarmClusterIssuer.acme.http01.enabled` and `giantSwarmClusterIssuer.acme.http01.ingressClassName` (which defaulted to `nginx`). Setting either now fails the Helm render with a message pointing at the replacement values.

#### Fixed

- Re-release `cert-manager` v4.2.0: CI didn't run properly on the release, and chart versions drifted since then.
- Use server-side apply in the `cert-manager-giantswarm-clusterissuer` hook Job, so its memory no longer grows with the number of CRDs in the cluster and it stops being OOMKilled on clusters with many CRDs.

### cilium [v1.6.2...v1.8.1](https://github.com/giantswarm/cilium-app/compare/v1.6.2...v1.8.1)

#### Added

- Support `k8s.apiServerURLs: auto` to enable Cilium's API server failover mode with the API endpoint and the `kubernetes` Service ClusterIP.

#### Changed

- Upgrade Cilium to [v1.20.5](https://github.com/cilium/cilium/releases/tag/v1.20.5).
- Tolerate the `node-role.kubernetes.io/etcd` taint in `cilium-operator` and the SPIRE agent, following upstream.
- Default `operator.rollOutPods` to `true`, so `cilium-operator` restarts when the rendered `cilium-config` changes. The operator reads its config only at startup.
- Default `authentication.enabled` to `false`, following upstream. With mutual authentication enabled, `cilium-agent` can deadlock during endpoint deletion. Clusters that use SPIRE mutual authentication need to set `authentication.enabled: true`.

### etcd-defrag [v1.2.13...v1.2.14](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.13...v1.2.14)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.48.0. ([#193](https://github.com/giantswarm/etcd-defrag-app/pull/193))

### k8s-dns-node-cache [v2.11.1...v2.11.3](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.11.1...v2.11.3)

#### Changed

- Prepare the repository for generated CI: publish only to `default-catalog` (no longer to `giantswarm-catalog`), make the ATS tests independent of the App CR and move the kind configuration to `.ats/kind-config.yaml`, and mark `resources` and `updateStrategy` as free-form for the generated values schema.

#### Fixed

- Scrape node-cache metrics once instead of twice.
- Remove `readOnlyRootFilesystem` from the pod security context, where it is not a valid field.

### net-exporter [v1.24.1...v1.25.1](https://github.com/giantswarm/net-exporter/compare/v1.24.1...v1.25.1)

#### Added

- Add the `daemonset.updateStrategy` value to configure the update strategy of the DaemonSet.

#### Changed

- Stop restarting all DaemonSet pods on every upgrade, only when the pod template changes.
- Roll out DaemonSet updates to 10% of the pods at a time by default, so that the rollout finishes within the Helm timeout on large workload clusters.
- Move the image base to alpine 3.24.2.
- Prepare for devctl-generated CI: keep the `control-plane-catalog` publish in `.circleci/custom.yml` and the go deps grouping in `renovate-custom.json5`.

#### Fixed

- Bump `golang.org/x/net` to v0.61.0 to fix CVE-2026-78663.

### network-policies [v0.5.0...v0.6.0](https://github.com/giantswarm/network-policies-app/compare/v0.5.0...v0.6.0)

#### Added

- Add `denyByDefault.clusterWide` to replace the per-namespace default deny with a single clusterwide policy, with the namespace opt-out label `policy.giantswarm.io/default-deny: "false"`.
- Add `denyByDefault.clusterWide.allowDNS` to allow DNS egress for every namespace the clusterwide deny covers.
- Add five standard allows, always rendered and purely additive, each selecting the namespaces labelled `policy.giantswarm.io/allow-<dns|kube-apiserver|webhook-ingress|ingress-controller|metrics-scraping>: "true"`, and `standardAllows.webhookIngress.fromEntities` for the webhook ingress source.
- Add the `clusterwide-deny` E2E test suite for the clusterwide default deny, its opt-out label, the Same-namespace policy and the standard allow labels.

### node-exporter [v1.21.0...v1.21.1](https://github.com/giantswarm/node-exporter-app/compare/v1.21.0...v1.21.1)

#### Changed

- The DaemonSet rolls out 10% of the pods at a time by default (`updateStrategy.rollingUpdate.maxUnavailable: "10%"`, previously `1`). Upgrades on large workload clusters now finish within the Helm timeout. `maxUnavailable` accepts an integer or a percentage.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### observability-bundle [v3.6.0...v3.8.1](https://github.com/giantswarm/observability-bundle/compare/v3.6.0...v3.8.1)

#### Changed

- Update `alloy` apps to 0.25.0 (Alloy v1.21.3).
- `mimir.rules.kubernetes` now recovers after a Mimir outage during a leader change.
- `loki_write_*` metrics no longer carry the `tenant` label.
- Override `CAdvisor` default scrape interval from 10s to 60s in kube-prometheus-stack.

#### Fixed

- Update `kube-prometheus-stack` to 24.0.2 (Prometheus Operator `update` on its custom resources' finalizers).

### prometheus-blackbox-exporter [v0.10.0...v0.11.0](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.10.0...v0.11.0)

#### Added

- Add the `daemonsetUpdateStrategy` value to set the update strategy of the DaemonSet.

#### Changed

- The DaemonSet replaces up to 10% of its pods at a time during an update (`rollingUpdate.maxUnavailable: 10%`). An update replaced one pod at a time, so it exceeded the Helm release timeout on large workload clusters.

#### Fixed

- The DaemonSet omits `updateStrategy` when `daemonsetUpdateStrategy` is unset, for example overridden to `null`, instead of rendering `updateStrategy: null`.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### security-bundle [v2.4.0...v2.5.0](https://github.com/giantswarm/security-bundle/compare/v2.4.0...v2.5.0)

#### Changed

- *!Breaking change*: Turn `apps.<appName>.dependsOn` from a list into a map from dependency name to `true`/`false`, so that a provider can turn off a single dependency (for example `cilium` on clusters where the provider manages the CNI) instead of overriding the whole list.
- Update `trivy` (app) to v0.18.1.
- Update `falco` (app) to v0.14.1.
- Update `policy-api` (app) to v0.1.0.
- Update `kyverno-crds` (app) to v1.19.2.
- Update `kyverno` (app) to v0.26.1.
  - This release includes a new Kyverno minor version. Please refer to the upstream release notes for the latest changes:
  - https://github.com/kyverno/kyverno/releases/tag/v1.19.0
  - https://github.com/kyverno/kyverno/releases/tag/v1.19.1
  - https://github.com/kyverno/kyverno/releases/tag/v1.19.2
- Update `kyverno-policies` (app) to v0.28.19.
- Update `starboard-exporter` (app) to v1.2.43.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### teleport-kube-agent [v0.12.0...v0.13.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.12.0...v0.13.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.11.10`.
- Updated `teleport-kube-agent` to upstream version `v18.11.12`.

### vsphere-csi-driver [v3.4.6...v3.5.0](https://github.com/giantswarm/vsphere-csi-driver-app/compare/v3.4.6...v3.5.0)

#### Changed

- Create the vendir update issue as the `giantswarm-project-automation` GitHub App instead of architectbot's `ISSUE_AUTOMATION` token, and take the board item when the board's auto-add adds it first (giantswarm/giantswarm#38800).
- Update upstream vsphere-csi-driver chart to v3.8.0.
- Use the `giantswarm-release` GitHub App instead of the taylorbot personal access token to push sync commits.
- Use `vendir` to vendor the upstream manifest, and move all patches into the `sync` directory.
- Pin the upstream version in `vendir.yml` instead of `Chart.yaml`. The chart version is no longer overwritten with the upstream version during a sync.
- Update architect to v10.10.0 (giantswarm/vsphere-csi-driver-app#221)
- Update architect to v10.11.1 (giantswarm/vsphere-csi-driver-app#222)
- Update architect to v10.12.0 (giantswarm/vsphere-csi-driver-app#225)
- Update architect to v10.12.1 (giantswarm/vsphere-csi-driver-app#226)
- Update architect to v10.12.2 (giantswarm/vsphere-csi-driver-app#228)
- Add Github action to create an issue in `giantswarm/giantswarm` for Team Rocket on vendir updates.
- Update architect to v10.13.0 (giantswarm/vsphere-csi-driver-app#232)
- Update architect to v10.13.1 (giantswarm/vsphere-csi-driver-app#233)
- Update architect to v10.13.5 (giantswarm/vsphere-csi-driver-app#236)
- Use Linux tools to manipulate chart label instead of an inscrutable Perl command.
- Update architect to v10.14.7 (giantswarm/vsphere-csi-driver-app#239)
- Update actions/create-github-app-token to v3.3.0 (giantswarm/vsphere-csi-driver-app#250)
- Update architect to v10.14.18 (giantswarm/vsphere-csi-driver-app#251)
- Improve changelog generation script.
- Update architect to v10.16.3 (giantswarm/vsphere-csi-driver-app#254)
- Update architect to v10.16.6 (giantswarm/vsphere-csi-driver-app#258)
- Update architect to v10.16.7 (giantswarm/vsphere-csi-driver-app#259)
- Update architect to v10.17.4 (giantswarm/vsphere-csi-driver-app#266)
- Update mikefarah/yq to v4.54.3 (giantswarm/vsphere-csi-driver-app#264)

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
- Fix changelog generation script.
