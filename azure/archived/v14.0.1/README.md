# :zap: Giant Swarm Release v14.0.1 for Azure :zap:

This releases increases the Azure Events Termination timeout from 5 to 15 minutes for better upgrade experience while workloads are moved to new nodes.

## Change details


### azure-operator [5.3.1](https://github.com/giantswarm/azure-operator/releases/tag/v5.3.1)

### Changed
- Increase VMSS termination events timeout to 15 minutes.

### azure-scheduled-events [0.2.2](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.2.2)

### Added
- Remove the `Node` from Kubernetes API server right before approving the termination event.

### Fixed
- Keep looping on events if one loop errors out.
- Disable helm hook for new installations of the chart.
