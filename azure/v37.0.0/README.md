# :zap: Giant Swarm Release v37.0.0 for Azure :zap:

## Changes compared to v36.0.0

### Components

- cluster-azure from v9.3.1 to v10.0.0
- cluster from v8.3.1 to v9.0.0
- Flatcar from v4593.2.5 to [v4757.2.1](https://www.flatcar.org/releases/#release-4757.2.1)
- Kubernetes from v1.36.5 to [v1.37.1](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.37.md#v1.37.1)
- os-tooling from v1.35.0 to v1.36.0

### cluster [v8.3.1...v9.0.0](https://github.com/giantswarm/cluster/compare/v8.3.1...v9.0.0)

#### Added

- Add `filesTemplateName` hook under `providerIntegration.workers.kubeadmConfig`. It names a provider template that renders a YAML list of files, once per node pool, so that a provider can deploy files to selected node pools only. A node pool for which the template renders nothing keeps its `KubeadmConfig` spec, and therefore its spec hash, unchanged.
- Enable `mergeDefaultEvictionSettings` to keep defaults for eviction like `nodefs.available` and `nodefs.inodesFree` which would otherwise be set to 0. This rolls all nodes.

#### Changed

- Stop deleting and recreating `/etc/ssl/certs` on nodes for fixing SELinux labeling, so any custom certificates placed there directly are preserved.

#### Removed

- Chart: Remove App to HelmRelease migration.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### Apps

- cilium from v1.6.0 to v1.6.1
- coredns from v1.33.0 to v1.34.0
- external-dns from v3.5.0 to v3.6.0
- net-exporter from v1.24.0 to v1.24.1
- network-policies from v0.3.1 to v0.3.2

### cilium [v1.6.0...v1.6.1](https://github.com/giantswarm/cilium-app/compare/v1.6.0...v1.6.1)

#### Added

- Declare the `io.giantswarm.application.audience` (`all`) and

#### Changed

- Upgrade Cilium to [v1.20.2](https://github.com/cilium/cilium/releases/tag/v1.20.2).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Point `home` at this repository instead of `https://cilium.io/`, as the standard requires.

### coredns [v1.33.0...v1.34.0](https://github.com/giantswarm/coredns-app/compare/v1.33.0...v1.34.0)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Changed

- Update `coredns` image to [1.14.7](https://github.com/coredns/coredns/releases/tag/v1.14.7).

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### external-dns [v3.5.0...v3.6.0](https://github.com/giantswarm/external-dns-app/compare/v3.5.0...v3.6.0)

#### Added

- Chart metadata: add `io.giantswarm.application.managed` annotation (`"true"`).
- Chart metadata: add `keywords`.

#### Changed

- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Upgrade external-dns to [v0.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.22.0).
- Sync to upstream helm chart [1.22.0](https://github.com/kubernetes-sigs/external-dns/releases/tag/external-dns-helm-chart-1.22.0).
  - Add `replicaCount` value to scale the deployment down to `0` or back to `1`.
  - Add `service.enabled` value to skip creating the `Service`.
  - Add `hostAliases` value to inject entries into the pod's `/etc/hosts`.
  - Add `crd` as a valid `registry` value, with the matching RBAC on `dnsrecords`.
  - Install the new `DNSRecord` CRD.
  - Narrow the RBAC verbs on `dnsendpoints/status` from `*` to `update`.
  - `policy` is now a required value. Our default of `sync` is unchanged.
  - Pin `annotationPrefix` to `external-dns.alpha.kubernetes.io/`, keeping the previous annotation prefix after upstream changed the default.

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### net-exporter [v1.24.0...v1.24.1](https://github.com/giantswarm/net-exporter/compare/v1.24.0...v1.24.1)

#### Added

- Add `keywords` to `Chart.yaml`.
- Add the `io.giantswarm.application.audience` (`all`) and `io.giantswarm.application.managed`

#### Changed

- Replace `interface{}` with `any` and use for-range over integers (Go modernization).
- Move the team annotation from the legacy `application.giantswarm.io/team` key to
- Set `Chart.yaml` `apiVersion` to `v2`. The chart declares no dependencies and had no

#### Fixed

- The `helm.sh/chart` label is valid for long chart versions: the 63-character cut trims the whole trailing run of `-`, `.` and `_`.

### network-policies [v0.3.1...v0.3.2](https://github.com/giantswarm/network-policies-app/compare/v0.3.1...v0.3.2)

#### Changed

- Convert the `allow-ingress-from-konnectivity` `CiliumNetworkPolicy` into a `CiliumClusterwideNetworkPolicy` named
