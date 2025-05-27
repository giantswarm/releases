# :zap: Giant Swarm Release v25.4.0 for CAPA :zap:

This release introduces `aws-node-termination-handler` for graceful draining of nodes during an upgrade or other type of replacement of worker nodes.

Details can be found in the [node pools documentation](https://docs.giantswarm.io/tutorials/fleet-management/cluster-management/node-pools/#what-happens-when-rolling-nodes).

## Changes compared to v25.3.0

### Components

- cluster-aws from v1.3.4 to v1.3.5

### cluster-aws [v1.3.4...v1.3.5](https://github.com/giantswarm/cluster-aws/compare/v1.3.4...v1.3.5)

#### Added

- Values: Add `global.providerSpecific.controlPlaneAmi` & `global.providerSpecific.nodePoolAmi`.
- Add aws-node-termination-handler bundle
- Make ASG lifecycle hook heartbeat timeout configurable

### Apps

- aws-nth-bundle v1.2.0
- cert-exporter from v2.9.0 to v2.9.3

### aws-nth-bundle [v1.2.0](https://github.com/giantswarm/aws-nth-bundle/releases/tag/v1.2.0)

#### Added

- Send spot instance interruption and instance state change events to SQS queue so that aws-node-termination-handler can react to them

### cert-exporter [v2.9.0...v2.9.3](https://github.com/giantswarm/cert-exporter/compare/v2.9.0...v2.9.3)

#### Added

- Chart: Add VPA and resources configuration for deployment and daemonset. ([#382](https://github.com/giantswarm/cert-exporter/pull/382))

#### Changed

- Chart: Enable `global.podSecurityStandards.enforced`. ([#420](https://github.com/giantswarm/cert-exporter/pull/420))
- Chart: Update PolicyExceptions to v2beta1. ([#358](https://github.com/giantswarm/cert-exporter/pull/358))
