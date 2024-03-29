# :zap: Giant Swarm Release v18.0.2 for AWS :zap:

This release fixes incompability with Cgroups V1 that was introduced in v18.0.0. The issue came from dropping dockershim and switching to containerd.

***IRSA highlights***
- Please read the [updated documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)
- Prior to upgrades please reach out to your Account Engineer and GiantSwarm team will help you in seemless migration if you have already enabled IRSA.
- Please remember to adapt the IAM policies prior to upgrade as specified in the [documentation](https://docs.giantswarm.io/advanced/iam-roles-for-service-accounts/)

***Features removed from v18.0.0***

- Cilium CNI was rolled back, and this release once again uses AWS CNI. This was done because of an upstream bug that prevented some advanced networking features to work properly.
- Out-of-tree cloud controller manager was removed, and in-tree alternative was enabled again because of an upstream bug regarding AWS CNI. Bug is already fixed in 1.24 version of upstream repo. 

## Change details


### aws-operator [13.2.1](https://github.com/giantswarm/aws-operator/releases/tag/v13.2.1)

### Fixed

- Fixed regression on cgroups v1 support after switching to containerd.


