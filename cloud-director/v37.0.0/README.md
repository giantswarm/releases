# :zap: Giant Swarm Release v37.0.0 for VMware Cloud Director :zap:

## Changes compared to v36.2.0

### Components

- cluster-cloud-director from v9.4.0 to v9.5.0
- cluster from v9.4.0 to v9.5.0
- Flatcar from v4757.2.2 to [v4757.2.3](https://www.flatcar.org/releases/#release-4757.2.3)
- Kubernetes from v1.36.5 to [v1.37.2](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.37.md#v1.37.2)

### cluster-cloud-director [v9.4.0...v9.5.0](https://github.com/giantswarm/cluster-cloud-director/compare/v9.4.0...v9.5.0)

#### Added

- Chart: Add `internal.unsolicitedArp.enabled` flag (default `false`). When enabled, nodes send unsolicited ARP for the IPv4 addresses of the interface holding the default route before `kubeadm` runs, so gateways drop stale IP-to-MAC entries when a new node reuses the IP of a removed node. This is only needed in Giant Swarm's testing environment so far and most likely not required in customer environments.

### cluster [v9.4.0...v9.5.0](https://github.com/giantswarm/cluster/compare/v9.4.0...v9.5.0)

#### Added

- Control Plane: Add `internal.advancedConfiguration.controlPlane.apiServer.enableAggregatorRouting` to set `--enable-aggregator-routing` on KubeadmControlPlane (disabled by default).

#### Changed

- Apps: Nest net-exporter `DNSCheck` and `Hosts` default values under `NetExporter` to match the chart's values schema.
- Control Plane: Harden kube-scheduler on KubeadmControlPlane and KamajiControlPlane following the [Kubernetes scheduler hardening guide](https://kubernetes.io/docs/concepts/security/hardening-guide/scheduler/). Set `--profiling=false`, `--tls-cipher-suites` (same list as kube-apiserver), `--authentication-skip-lookup=false`, `--authentication-tolerate-lookup-failure=false`, `--permit-address-sharing=false` and `--permit-port-sharing=false`. The scheduler now fails to start if it cannot look up its authentication configuration from the API server. `--bind-address=0.0.0.0` and `/metrics` in `--authorization-always-allow-paths` are kept for metrics scraping.

### Apps

- cert-exporter from v2.12.1 to v2.13.0
- cert-manager from v4.1.2 to v4.1.3
- net-exporter from v1.24.1 to v1.24.2
- observability-bundle from v3.6.0 to v3.8.0

### cert-exporter [v2.12.1...v2.13.0](https://github.com/giantswarm/cert-exporter/compare/v2.12.1...v2.13.0)

#### Added

- `fingerprint` label (first 16 hex characters of the certificate's SHA-256 fingerprint) on `cert_exporter_not_after` and `cert_exporter_secret_not_after`.

#### Changed

- `/metrics` now logs gather and encoding errors (such as a dropped duplicate series) instead of discarding them silently.
- A truly identical certificate duplicated within one secret key or file is reported once, without producing a gather error on every scrape.

#### Fixed

- Certificates sharing a serial number in one secret key or file no longer collide into one series. Previously only the first, expired certificate was exported.
- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### cert-manager [v4.1.2...v4.1.3](https://github.com/giantswarm/cert-manager-app/compare/v4.1.2...v4.1.3)

#### Fixed

- Use server-side apply in the `cert-manager-giantswarm-clusterissuer` hook Job, so its memory no longer grows with the number of CRDs in the cluster and it stops being OOMKilled on clusters with many CRDs.

### net-exporter [v1.24.1...v1.24.2](https://github.com/giantswarm/net-exporter/compare/v1.24.1...v1.24.2)

#### Changed

- Move the image base to alpine 3.24.2.
- Prepare for devctl-generated CI: keep the `control-plane-catalog` publish in `.circleci/custom.yml` and the go deps grouping in `renovate-custom.json5`.

#### Fixed

- Bump `golang.org/x/net` to v0.61.0 to fix CVE-2026-78663.

### observability-bundle [v3.6.0...v3.8.0](https://github.com/giantswarm/observability-bundle/compare/v3.6.0...v3.8.0)

#### Changed

- Update `alloy` apps to 0.25.0 (Alloy v1.21.3).
- `mimir.rules.kubernetes` now recovers after a Mimir outage during a leader change.
- `loki_write_*` metrics no longer carry the `tenant` label.
- Override `CAdvisor` default scrape interval from 10s to 60s in kube-prometheus-stack.
