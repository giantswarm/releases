# :zap: Giant Swarm Release v14.1.5 for Azure :zap:

This Workload Cluster release contains fixes for azure-operator handling multi-tenant service principal secrets as a prerequisite for migration from VPN Gateway to VNET peering. Moreover we have extended azure-scheduled-events app with additional Azure VMSS events handling. 

Please contact your Solution Architect in order to validate if this release is necessary for your use-case.

## Change details


### azure-operator [5.5.3-dev](https://github.com/giantswarm/azure-operator/releases/tag/v5.5.3-dev)

Not found


### azure-scheduled-events [0.4.0](https://github.com/giantswarm/azure-scheduled-events/releases/tag/v0.4.0)

#### Added
- React to `Preempt`, `Reboot` and `Redeploy` events to drain nodes properly.
#### Change
- Use the `NotBefore` field from the event to calculate drain timeout.



