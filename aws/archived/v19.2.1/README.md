# :zap: Giant Swarm Release v19.2.1 for AWS :zap:

This patch release of the 19.2 series addresses an error in the calculation of max pods per node when using Cilium in ENI IPAM mode.

## Change details


### aws-operator [14.23.1](https://github.com/giantswarm/aws-operator/releases/tag/v14.23.1)

#### Fixed
- Bump k8scc to 16.6.1 to fix max pod calculation.



