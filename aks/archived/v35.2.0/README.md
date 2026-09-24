# :zap: Giant Swarm Release v35.2.0 for AKS :zap:

## Changes compared to v35.1.0

### Components

- cluster-aks from v0.5.0 to v0.6.0

### cluster-aks [v0.5.0...v0.6.0](https://github.com/giantswarm/cluster-aks/compare/v0.5.0...v0.6.0)

#### Changed

- Chart: Update `cluster` to v8.2.0.

#### Fixed

- Sanitize the `app.kubernetes.io/version` label value so it is always a valid Kubernetes label. New dev builds are longer and might trigger a validation failure in some cases.
- Configure `observability-bundle` and `security-bundle` HelmReleases dependencies to not include `cilium`, since it's not installed on AKS.

### cluster [v7.0.0...v8.2.0](https://github.com/giantswarm/cluster/compare/v7.0.0...v8.2.0)

#### Added

- Add `preKubeadmCommandsTemplateName` and `postKubeadmCommandsTemplateName` hooks under `providerIntegration.controlPlane.kubeadmConfig` and `providerIntegration.workers.kubeadmConfig`. They name a provider template that renders a YAML list of additional kubeadm commands, once for the control plane and once per node pool for workers.
- Add `internal.advancedConfiguration.kubelet.evictionHard` values. Providers need them to tell autoscalers such as Karpenter how much of a node's resources is allocatable.
- SELinux: Add `global.components.selinux.writablePolicyStore` value (default `true`) to allow loading additional SELinux policies.

#### Changed

- SELinux: Keep AVC audit logs (required for SELinux policy generation).
- SELinux: Relabel the whole filesystem except read-only `/usr` (previously only `/etc/kubernetes`).
- SELinux: Correctly label CA certificates in `/etc/ssl/certs` for mounting into containers.
- App to HR Migration: Skip v35.0.0 pre-releases and update `docker-kubectl` to v1.36.4.
- Chart: Rework HelmRelease clean-up job.
- Chart: Migrate Apps to HelmReleases.

### Apps

- network-policies from v0.2.0 to v0.3.0
- node-exporter from v1.20.11 to v1.20.13
- observability-bundle from v3.3.0 to v3.4.0
- security-bundle from v2.2.0 to v2.4.0
- teleport-kube-agent from v0.11.1 to v0.12.0

### network-policies [v0.2.0...v0.3.0](https://github.com/giantswarm/network-policies-app/compare/v0.2.0...v0.3.0)

#### Added

- Add `keywords` to `Chart.yaml`.
- Declare the `io.giantswarm.application.audience` (`all`) and
- Add optional `denyEgressToIMDS` policy denying pod egress to the instance metadata service. Disabled by default.
- Add apptest-framework e2e test suite.

#### Changed

- Move the team annotation from the legacy `application.giantswarm.io/team` key to

### node-exporter [v1.20.11...v1.20.13](https://github.com/giantswarm/node-exporter-app/compare/v1.20.11...v1.20.13)

#### Changed

- CircleCI: Do not override app version.
- Chart: Move PolicyException to `kube-system` namespace.

### observability-bundle [v3.3.0...v3.4.0](https://github.com/giantswarm/observability-bundle/compare/v3.3.0...v3.4.0)

#### Changed

- Values: Update Prometheus Operator CRD and Kube Prometheus Stack to v23.0.0.
- Values: Generate schema for Alloy PodLogs CRDs.
- Values: Add Cilium as dependency for Alloy apps & Kube Prometheus Stack.
- Values: Update Alloy apps to v0.21.2.

#### Removed

- Values: Remove unused catalog.

### security-bundle [v2.2.0...v2.4.0](https://github.com/giantswarm/security-bundle/compare/v2.2.0...v2.4.0)

#### Added

- Add e2e scenarios covering trivy-operator `VulnerabilityReport` creation, starboard-exporter metrics for that report, kyverno restricted PSS enforcement, and kyverno-policy-operator `PolicyException` translation.

#### Changed

- Update `exception-recommender` (app) to v0.3.0.
- Update `falco` (app) to v0.13.0.
- Update `jiralert` (app) to v0.1.4.
- Update `kubescape` (app) to v0.1.1.
- Update `kyverno-policies` (app) to v0.27.1.
- Update `policy-api` (app) to v0.0.12.
- Update `starboard-exporter` (app) to v1.2.15.
- Update `trivy` (app) to v0.18.0.
- Update `trivy-operator` (app) to v0.15.0.
- Update `kyverno-policy-operator` (app) to v0.2.3.
- Update `policy-api` (app) to v0.0.9.
- Update `starboard-exporter` (app) to v1.2.3.
- Update `trivy` (app) to v0.17.0.
- Update `trivy-operator` (app) to v0.13.3.
- Run the E2E test suites automatically on release PRs by adding `.github/release-pr-body.md`.
- Values: Make Kyverno and Kyverno Policy Operator depend on Cilium.

#### Removed

- Remove `gel` (app).

#### Fixed

- Give `kubescape` a 15m install and upgrade timeout. It does not finish installing within Flux's 5m default, so it failed with `context deadline exceeded` and then retried indefinitely.
- Set `createNamespace` on every app in the bundle, so each one creates its target namespace instead of relying on another app to have created it first. Previously `kubescape` failed with `namespaces "kubescape" not found`, and the apps targeting `security-bundle` could only install after `kyverno-policy-operator` had created it.

### teleport-kube-agent [v0.11.1...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.11.1...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
