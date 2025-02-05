# :zap: Giant Swarm Release v29.6.1 for CAPA :zap:

## Changes compared to v29.6.0

### Components

- cluster-aws from v2.6.0 to v2.6.1

### cluster-aws [v2.6.0...v2.6.1](https://github.com/giantswarm/cluster-aws/compare/v2.6.0...v2.6.1)

#### Added

- Add ingress rule in nodes Security Group to allow access to the Kubelet API when using ENI mode. This is needed by the metrics server to gather metrics from the Kubele

### Apps

- cilium-crossplane-resources from v0.1.0 to v0.2.0

### cilium-crossplane-resources [v0.1.0...v0.2.0](https://github.com/giantswarm/cilium-crossplane-resources/compare/v0.1.0...v0.2.0)

#### Added

- Add a Security Group rule for node to pod communication
