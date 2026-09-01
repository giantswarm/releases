# :zap: Giant Swarm Release v36.0.0 for VMware Cloud Director :zap:

## Changes compared to v35.0.1

### Components

- Kubernetes from v1.35.8 to [v1.36.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.4)
- os-tooling from v1.34.0 to v1.34.1

### Apps

- cilium from v1.5.1 to v1.6.0
- etcd-defrag from v1.2.10 to v1.2.11
- teleport-kube-agent from v0.11.1 to v0.12.0

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

### etcd-defrag [v1.2.10...v1.2.11](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.10...v1.2.11)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.44.0. ([#129](https://github.com/giantswarm/etcd-defrag-app/pull/129))

### teleport-kube-agent [v0.11.1...v0.12.0](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.11.1...v0.12.0)

#### Changed

- Updated `teleport-kube-agent` to upstream version `v18.10.7`.
