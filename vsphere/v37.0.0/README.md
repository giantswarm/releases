# :zap: Giant Swarm Release v37.0.0 for vSphere :zap:

## Changes compared to v36.2.0

### Components

- cluster-vsphere from v11.4.0 to v11.5.0
- cluster from v9.4.0 to v9.5.0
- Flatcar from v4757.2.2 to [v4757.2.3](https://www.flatcar.org/releases/#release-4757.2.3)
- Kubernetes from v1.36.5 to [v1.37.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.37.md#v1.37.2)

### cluster [v9.4.0...v9.5.0](https://github.com/giantswarm/cluster/compare/v9.4.0...v9.5.0)

#### Added

- Control Plane: Add `internal.advancedConfiguration.controlPlane.apiServer.enableAggregatorRouting` to set `--enable-aggregator-routing` on KubeadmControlPlane (disabled by default).

#### Changed

- Apps: Nest net-exporter `DNSCheck` and `Hosts` default values under `NetExporter` to match the chart's values schema.
- Control Plane: Harden kube-scheduler on KubeadmControlPlane and KamajiControlPlane following the [Kubernetes scheduler hardening guide](https://kubernetes.io/docs/concepts/security/hardening-guide/scheduler/). Set `--profiling=false`, `--tls-cipher-suites` (same list as kube-apiserver), `--authentication-skip-lookup=false`, `--authentication-tolerate-lookup-failure=false`, `--permit-address-sharing=false` and `--permit-port-sharing=false`. The scheduler now fails to start if it cannot look up its authentication configuration from the API server. `--bind-address=0.0.0.0` and `/metrics` in `--authorization-always-allow-paths` are kept for metrics scraping.

### Apps

- cert-exporter from v2.12.1 to v2.13.0
- cert-manager from v4.1.2 to v4.3.0
- cilium from v1.6.2 to v1.8.0
- cloud-provider-vsphere from v2.5.1 to v2.6.0
- etcd-defrag from v1.2.13 to v1.2.14
- net-exporter from v1.24.1 to v1.24.2
- network-policies from v0.5.0 to v0.6.0
- observability-bundle from v3.6.0 to v3.8.0
- prometheus-blackbox-exporter from v0.10.0 to v0.11.0
- security-bundle from v2.4.0 to v2.5.0
- vsphere-csi-driver from v3.4.6 to v3.5.0

### cert-exporter [v2.12.1...v2.13.0](https://github.com/giantswarm/cert-exporter/compare/v2.12.1...v2.13.0)

#### Added

- `fingerprint` label (first 16 hex characters of the certificate's SHA-256 fingerprint) on `cert_exporter_not_after` and `cert_exporter_secret_not_after`.

#### Changed

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

### cilium [v1.6.2...v1.8.0](https://github.com/giantswarm/cilium-app/compare/v1.6.2...v1.8.0)

#### Added

- Support `k8s.apiServerURLs: auto` to enable Cilium's API server failover mode with the API endpoint and the `kubernetes` Service ClusterIP.

#### Changed

- Default `operator.rollOutPods` to `true`, so `cilium-operator` restarts when the rendered `cilium-config` changes. The operator reads its config only at startup.
- Default `authentication.enabled` to `false`, following upstream. With mutual authentication enabled, `cilium-agent` can deadlock during endpoint deletion. Clusters that use SPIRE mutual authentication need to set `authentication.enabled: true`.

### cloud-provider-vsphere [v2.5.1...v2.6.0](https://github.com/giantswarm/cloud-provider-vsphere-app/compare/v2.5.1...v2.6.0)

#### Changed

- Update dependency kubernetes/cloud-provider-vsphere to v1.37.0.
- Update architect to v10.12.2 (giantswarm/cloud-provider-vsphere-app#381)
- Update architect to v10.13.1 (giantswarm/cloud-provider-vsphere-app#383)
- Update architect to v10.13.2 (giantswarm/cloud-provider-vsphere-app#384)
- Update architect to v10.14.18 (giantswarm/cloud-provider-vsphere-app#390)
- Update architect to v10.15.1 (giantswarm/cloud-provider-vsphere-app#393)
- Update architect to v10.16.3 (giantswarm/cloud-provider-vsphere-app#394)
- Update architect to v10.16.5 (giantswarm/cloud-provider-vsphere-app#404)
- Update architect to v10.16.6 (giantswarm/cloud-provider-vsphere-app#406)
- Update architect to v10.16.7 (giantswarm/cloud-provider-vsphere-app#407)
- Update architect to v10.17.4 (giantswarm/cloud-provider-vsphere-app#410)

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### etcd-defrag [v1.2.13...v1.2.14](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.13...v1.2.14)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.48.0. ([#193](https://github.com/giantswarm/etcd-defrag-app/pull/193))

### net-exporter [v1.24.1...v1.24.2](https://github.com/giantswarm/net-exporter/compare/v1.24.1...v1.24.2)

#### Changed

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

### observability-bundle [v3.6.0...v3.8.0](https://github.com/giantswarm/observability-bundle/compare/v3.6.0...v3.8.0)

#### Changed

- Update `alloy` apps to 0.25.0 (Alloy v1.21.3).
- `mimir.rules.kubernetes` now recovers after a Mimir outage during a leader change.
- `loki_write_*` metrics no longer carry the `tenant` label.
- Override `CAdvisor` default scrape interval from 10s to 60s in kube-prometheus-stack.

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
