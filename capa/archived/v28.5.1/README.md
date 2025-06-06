# :zap: Giant Swarm Release v28.5.1 for CAPA :zap:

This release introduces improvements for ENI mode targetting the CAPA migration process.

## Changes compared to v28.5.0

### Components

- cluster-aws from v1.3.6 to v1.3.7

### cluster-aws [v1.3.6...v1.3.7](https://github.com/giantswarm/cluster-aws/compare/v1.3.6...v1.3.7)

#### Added

- Add ingress rule in nodes Security Group to allow access to the Kubelet API when using ENI mode. This is needed by the metrics server to gather metrics from the Kubelet

### Apps

- cilium-crossplane-resources from v0.1.0 to v0.2.0

### cilium-crossplane-resources [v0.1.0...v0.2.0](https://github.com/giantswarm/cilium-crossplane-resources/compare/v0.1.0...v0.2.0)

#### Added

- Add a Security Group rule for node to pod communication
