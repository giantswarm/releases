# :zap: Giant Swarm Release v18.0.1 for Azure :zap:

This is a bugfix release that could lead to `azure-cloud-controller-manager` app to be configured with the wrong Pod CIDR.
Upgrading to this release from 18.0.0 will not require rolling of nodes.

## Change details

### cluster-operator [4.6.2](https://github.com/giantswarm/cluster-operator/releases/tag/v4.6.2)

#### Fixed
- Use `AzureConfig`'s `Spec.Azure.VirtualNetwork.CalicoSubnetCIDR` field for Calico CIDR rather than `Spec.Cluster.Calico.Subnet`.



