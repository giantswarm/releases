# :zap: Giant Swarm Release v35.1.0 for Azure :zap:

## Changes compared to v35.0.1

### Components

- cluster-azure from v9.0.0 to v9.2.0
- cluster from v8.0.0 to v8.2.0
- Flatcar from v4593.2.4 to [v4593.2.5](https://www.flatcar.org/releases/#release-4593.2.5)
- os-tooling from v1.34.0 to v1.34.1

### cluster-azure [v9.0.0...v9.2.0](https://github.com/giantswarm/cluster-azure/compare/v9.0.0...v9.2.0)

#### Added

- Add initial Bring-Your-Own-Network support.
- Add option to be able to disable the Private Link on private clusters (`.global.connectivity.network.enablePrivateLinkWithPrivateMode`). Useful for BYON scenarios.

### cluster [v8.0.0...v8.2.0](https://github.com/giantswarm/cluster/compare/v8.0.0...v8.2.0)

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

### Apps

- azure-cloud-controller-manager from v2.1.0 to v2.2.0
- azure-cloud-node-manager from v2.1.0 to v2.2.0
- azuredisk-csi-driver from v2.1.0 to v2.2.0
- azurefile-csi-driver from v2.0.0 to v2.1.0
- cilium from v1.5.1 to v1.6.0
- cluster-autoscaler from v2.0.4 to v2.1.0
- etcd-defrag from v1.2.10 to v1.2.12
- network-policies from v0.2.0 to v0.3.0
- observability-bundle from v3.3.1 to v3.4.0
- security-bundle from v2.3.0 to v2.4.0
- teleport-kube-agent from v0.11.1 to v0.12.0

### azure-cloud-controller-manager [v2.1.0...v2.2.0](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v2.1.0...v2.2.0)

#### Changed

- Chart: Update to upstream v1.36.5.

### azure-cloud-node-manager [v2.1.0...v2.2.0](https://github.com/giantswarm/azure-cloud-node-manager-app/compare/v2.1.0...v2.2.0)

#### Changed

- Chart: Update to upstream v1.36.5.

### azuredisk-csi-driver [v2.1.0...v2.2.0](https://github.com/giantswarm/azuredisk-csi-driver-app/compare/v2.1.0...v2.2.0)

#### Changed

- Migrate to App Build Suite (ABS).
- Chart: Update to upstream v1.34.5.

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### azurefile-csi-driver [v2.0.0...v2.1.0](https://github.com/giantswarm/azurefile-csi-driver-app/compare/v2.0.0...v2.1.0)

#### Changed

- Migrate to App Build Suite (ABS).
- Chart: Update to upstream v1.35.7.

#### Removed

- Removed `PodSecurityPolicy`.
- Removed `global.podSecurityStandards.enforced` helm value.

### cilium [v1.5.1...v1.6.0](https://github.com/giantswarm/cilium-app/compare/v1.5.1...v1.6.0)

#### Changed

- Upgrade Cilium to [v1.20.1](https://github.com/cilium/cilium/releases/tag/v1.20.1) from v1.19.7. Please review the upstream [1.20 upgrade notes](https://docs.cilium.io/en/v1.20/operations/upgrade/) before rolling this out.
- Serve the ztunnel image from `gsoci.azurecr.io/giantswarm/cilium-ztunnel` instead of pulling it from upstream. Cilium v1.20 moved this image from `docker.io/istio/ztunnel` to `quay.io/cilium/ztunnel:v1.0.0`, and our mirror carries exactly that digest, so it no longer has to be allow-listed in `sync/unmirrored-images.txt`. Only used by `encryption.type=ztunnel`, which we do not support.

#### Removed

- Upstream removed these long-deprecated Helm values in v1.20. None of them are set by this chart's defaults, and because the chart's `values.schema.json` does not reject unknown keys, leftovers in existing values are **silently ignored** rather than rejected — check your values before upgrading:
  - `encryption.strictMode.{enabled,cidr,allowRemoteNodeIdentities}` → use `encryption.strictMode.egress.*`
  - `encryption.ipsec.encryptedOverlay`
  - `clustermesh.enableMCSAPISupport` → use `clustermesh.mcsapi.enabled` (MCS-API is now stable upstream)
  - `clustermesh.apiserver.tls.{server,admin,remote}.{cert,key}` and `clustermesh.apiserver.tls.enableSecrets` → enable auto-generation or pre-create the secrets
  - `hubble.redact.kafka.apiKey` → Kafka-aware L7 policy support and proxylib were removed upstream
  - `preflight.tofqdnsPreCache` → the preflight FQDN poller was removed upstream
  - `hubble.ui.backend.{livenessProbe,readinessProbe}.enabled`
- `sync/patches/certgen/`. Cilium v1.20 ships the `certgen.enforceCAValidityThroughoutLeavesDuration` value and wires `--ca-enforce-validity-throughout-leaves-duration` into both certgen job specs itself, so the Giant Swarm patch that added them became a no-op (it detected this and skipped). The default stays `true` and the rendered job specs are unchanged, so two more patches drop out of `diffs/`.

### cluster-autoscaler [v2.0.4...v2.1.0](https://github.com/giantswarm/cluster-autoscaler-app/compare/v2.0.4...v2.1.0)

#### Changed

- Chart: Update to upstream v1.36.1.

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

### observability-bundle [v3.3.1...v3.4.0](https://github.com/giantswarm/observability-bundle/compare/v3.3.1...v3.4.0)

#### Changed

- Values: Update Prometheus Operator CRD and Kube Prometheus Stack to v23.0.0.

### security-bundle [v2.3.0...v2.4.0](https://github.com/giantswarm/security-bundle/compare/v2.3.0...v2.4.0)

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

#### Removed

- Remove `gel` (app).

#### Fixed

- Give `kubescape` a 15m install and upgrade timeout. It does not finish installing within Flux's 5m default, so it failed with `context deadline exceeded` and then retried indefinitely.
- Set `createNamespace` on every app in the bundle, so each one creates its target namespace instead of relying on another app to have created it first. Previously `kubescape` failed with `namespaces "kubescape" not found`, and the apps targeting `security-bundle` could only install after `kyverno-policy-operator` had created it.

### teleport-kube-agent [v0.11.1...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.11.1...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
