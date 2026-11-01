# :zap: Giant Swarm Release v36.3.0 for CAPA :zap:

## Changes compared to v36.2.0

### Components

- Flatcar from v4757.2.2 to [v4757.2.3](https://www.flatcar.org/releases/#release-4757.2.3)
- Kubernetes from v1.36.5 to [v1.36.6](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.6)

### Apps

- cert-exporter from v2.12.1 to v2.12.2
- external-dns from v3.7.1 to v3.7.2
- k8s-audit-metrics from v0.10.13 to v0.10.14
- net-exporter from v1.24.1 to v1.24.2

### cert-exporter [v2.12.1...v2.12.2](https://github.com/giantswarm/cert-exporter/compare/v2.12.1...v2.12.2)

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### external-dns [v3.7.1...v3.7.2](https://github.com/giantswarm/external-dns-app/compare/v3.7.1...v3.7.2)

#### Fixed

- Values schema: accept the values the Giant Swarm platform injects that the stricter schema from 3.7.0 rejected with `values-schema-violation`: `managementCluster`, `global.podSecurityStandards`, and the `externalDNSIP` and `helm` cluster values cluster-apps-operator sets for private clusters and the management cluster. `global` now allows additional properties.

### k8s-audit-metrics [v0.10.13...v0.10.14](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.13...v0.10.14)

#### Changed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.
- Docs: Update links.
- Go: Update dependencies.

#### Fixed

- Fix ARM64 images.

### net-exporter [v1.24.1...v1.24.2](https://github.com/giantswarm/net-exporter/compare/v1.24.1...v1.24.2)

#### Changed

- Move the image base to alpine 3.24.2.
- Prepare for devctl-generated CI: keep the `control-plane-catalog` publish in `.circleci/custom.yml` and the go deps grouping in `renovate-custom.json5`.

#### Fixed

- Bump `golang.org/x/net` to v0.61.0 to fix CVE-2026-78663.
