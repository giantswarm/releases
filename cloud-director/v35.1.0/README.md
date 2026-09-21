# :zap: Giant Swarm Release v35.1.0 for VMware Cloud Director :zap:

## Changes compared to v35.0.1

### Components

- cluster-cloud-director from v7.0.0 to v7.3.0
- cluster from v8.0.0 to v8.3.0
- Flatcar from v4593.2.5 to [v4757.2.0](https://www.flatcar.org/releases/#release-4757.2.0)
- os-tooling from v1.34.0 to v1.34.1

### cluster [v8.0.0...v8.3.0](https://github.com/giantswarm/cluster/compare/v8.0.0...v8.3.0)

#### Added

- Add `preKubeadmCommandsTemplateName` and `postKubeadmCommandsTemplateName` hooks under `providerIntegration.controlPlane.kubeadmConfig` and `providerIntegration.workers.kubeadmConfig`. They name a provider template that renders a YAML list of additional kubeadm commands, once for the control plane and once per node pool for workers.
- Add `internal.advancedConfiguration.kubelet.evictionHard` values. Providers need them to tell autoscalers such as Karpenter how much of a node's resources is allocatable.
- SELinux: Add `global.components.selinux.writablePolicyStore` value (default `true`) to allow loading additional SELinux policies.

#### Changed

- Enable the `ClusterTrustBundle` and `ClusterTrustBundleProjection` feature gates (Kubernetes 1.33+) and the `PodCertificateRequest` feature gate (Kubernetes 1.35+) by default on kube-apiserver, kube-controller-manager and kubelet.
- SELinux: Keep AVC audit logs (required for SELinux policy generation).
- SELinux: Relabel the whole filesystem except read-only `/usr` (previously only `/etc/kubernetes`).
- SELinux: Correctly label CA certificates in `/etc/ssl/certs` for mounting into containers.
- App to HR Migration: Skip v35.0.0 pre-releases and update `docker-kubectl` to v1.36.4.
- Chart: Rework HelmRelease clean-up job.

### Apps

- cert-exporter from v2.12.0 to v2.12.1
- coredns from v1.32.0 to v1.33.0
- etcd-defrag from v1.2.10 to v1.2.12
- network-policies from v0.2.0 to v0.3.0
- node-exporter from v1.20.13 to v1.21.0
- observability-bundle from v3.3.1 to v3.5.0
- teleport-kube-agent from v0.11.1 to v0.12.0

### cert-exporter [v2.12.0...v2.12.1](https://github.com/giantswarm/cert-exporter/compare/v2.12.0...v2.12.1)

#### Fixed

- A cert file that cannot be read no longer aborts the scan of its whole cert path. Previously one unreadable file (such as a root-only `0600` `ca.crt` on an AKS node, where the exporter runs as an unprivileged user) stopped the walk, silently dropping every file sorting after it from the metrics. Unreadable files are now logged and skipped individually.

### coredns [v1.32.0...v1.33.0](https://github.com/giantswarm/coredns-app/compare/v1.32.0...v1.33.0)

#### Changed

- Update `coredns` image to [1.14.6](https://github.com/coredns/coredns/releases/tag/v1.14.6).
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.

#### Fixed

- Honor the deprecated `configmap.log`, `loadbalancePolicy` and `configmap.cache` again. Since 1.31.0 `coredns.<zone>.log`, `coredns.<zone>.loadbalance` and `coredns.<zone>.cache.success.ttl` shipped defaults that shadowed them, so the old keys were silently ignored. They are now unset by default, restoring the documented fallback chain. Rendering with default values is unchanged.

### etcd-defrag [v1.2.10...v1.2.12](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.10...v1.2.12)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.45.0. ([#136](https://github.com/giantswarm/etcd-defrag-app/pull/136))
- Chart: Update dependency ahrtr/etcd-defrag to v0.44.0. ([#129](https://github.com/giantswarm/etcd-defrag-app/pull/129))

### network-policies [v0.2.0...v0.3.0](https://github.com/giantswarm/network-policies-app/compare/v0.2.0...v0.3.0)

#### Added

- Add `keywords` to `Chart.yaml`.
- Declare the `io.giantswarm.application.audience` (`all`) and
- Add optional `denyEgressToIMDS` policy denying pod egress to the instance metadata service. Disabled by default.
- Add apptest-framework e2e test suite.

#### Changed

- Move the team annotation from the legacy `application.giantswarm.io/team` key to

### node-exporter [v1.20.13...v1.21.0](https://github.com/giantswarm/node-exporter-app/compare/v1.20.13...v1.21.0)

#### Added

- `disableSystemdCollector` value to turn the systemd collector off, mirroring the existing `disableConntrackCollector` and `disableNvmeCollector` toggles. The collector needs a D-Bus connection to the host, which is refused on nodes where AppArmor mediates D-Bus (such as AKS Ubuntu nodes running under the default containerd profile), making it fail on every scrape. Defaults to `false`, so behaviour is unchanged.

### observability-bundle [v3.3.1...v3.5.0](https://github.com/giantswarm/observability-bundle/compare/v3.3.1...v3.5.0)

#### Added

- Point kube-prometheus-stack's control-plane ServiceMonitors at the alloy-metrics token Secret.

#### Changed

- Update `alloy` apps to 0.23.1 (Alloy v1.19.2).
- Update `prometheus-operator-crd` to 24.0.0 (Prometheus Operator CRDs v0.94.0).
- Update `kube-prometheus-stack` to 24.0.0 (chart 91.2.3, Prometheus Operator v0.94.0).
- Values: Update Prometheus Operator CRD and Kube Prometheus Stack to v23.0.0.

#### Fixed

- KSM custom resource state: Set the Gateway API `TCPRoute` and `UDPRoute` collectors to `v1`, which is the version the API server serves.

### teleport-kube-agent [v0.11.1...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.11.1...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
