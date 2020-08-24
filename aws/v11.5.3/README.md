# :zap: Giant Swarm Release v11.5.3 for AWS :zap:

This release provides a new cluster-operator which fixes preventing release upgrade when reference id does not align with the G8sControlPlane id or MachineDeployment id.

## Change details

### cluster-operator [2.3.3](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.3)

- Fixed condition where reference id does not match with G8sControlplane or MachineDeployment.
