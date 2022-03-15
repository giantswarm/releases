# :zap: Giant Swarm Release v17.0.0 for Azure :zap:

<< Add description here >>

## Change details

### cluster-autoscaler [1.22.2-gs4](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.22.2-gs4)

***Fixed***

- Updated to correct cluster-autoscaler version
- Use GS-built 1.22 image to deliver upstream unreleased fix kubernetes/autoscaler#4600
 
***Added***

- Added support for specifying balance-similar-node-groups flag
          
### azure-operator [5.17.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.17.0)

***Fixed***

- Fix panic while checking for cgroups version during upgrade.

***Added***

- Add GiantSwarmCluster tag to Vnet.



