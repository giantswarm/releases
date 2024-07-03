# :zap: Giant Swarm Release v25.1.0 for CAPA :zap:

Introduces the new `giantswarm-cluster-suite` app and provides general improvements.

## Change details

### giantswarm-cluster-suite [0.1.1](https://github.com/giantswarm/giantswarm-cluster-suite/releases/tag/v0.1.17)

#### Added
- Create `giantswarm` Namespace when it does not already exist.
- Create `giantswarm-critical` PriorityClass when it does not already exists.

### security-bundle [1.7.1](https://github.com/giantswarm/security-bundle/releases/tag/v1.7.1)

#### Changed

- Bump `kyverno` app to the v0.17.14 version.
- Bump `starboard-exporter` app to the v0.7.11 version.

### prometheus-blackbox-exporter [0.4.2](https://github.com/giantswarm/prometheus-blackbox-exporter-app/releases/tag/v0.4.2)

#### Fix

- Remove duplicated team label.

### irsa-servicemonitors [0.1.0](https://github.com/giantswarm/irsa-servicemonitors-app/releases/tag/v0.1.0)

#### Fix

- Remove duplicated team label.
