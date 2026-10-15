# :zap: Giant Swarm Release v35.4.0 for AKS :zap:

This release adds Giant Swarm's Cilium for the BYO CNI opt-in (`global.connectivity.cni.managed: true` in cluster-aks). Cilium is only installed on clusters that opt in. Clusters on AKS-managed networking (the default) get no Cilium and no new dependencies. The BYO CNI mode is not supported yet: use it on test clusters only.

## Changes compared to v35.3.0


### Apps

- Added cilium v1.5.2

### cilium [v1.5.2](https://github.com/giantswarm/cilium-app/releases/tag/v1.5.2)

#### Changed

- Upgrade Cilium to [v1.19.8](https://github.com/cilium/cilium/releases/tag/v1.19.8).
