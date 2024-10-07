# :zap: Giant Swarm Release v29.3.0 for CAPA :zap:

## Changes compared to v29.2.0

This release does not contain any changes to components or apps, but makes use of an updated machine image, which includes a fix for accessing private Elastic Container Registries (ECR).

#### ⚠️ Breaking change introduced in v29.2.0 with `cluster-aws` version 2.1.0

- Do not allow additional properties in the following fields in order to avoid unnoticed typos:
  - global.connectivity.network
  - global.connectivity.network.pods
  - global.connectivity.network.services
  - global.connectivity.subnets[]
  - global.connectivity.topology
  - global.controlPlane
  - global.controlPlane.additionalSecurityGroups[]
  - global.controlPlane.machineHealthCheck
  - global.controlPlane.oidc
  - global.providerSpecific
  - global.providerSpecific.instanceMetadataOptions

If you were using values like `global.controlPlane.containerdVolumeSizeGB` and `global.controlPlane.kubeletVolumeSizeGB`, please move to the new `.global.controlPlane.libVolumeSizeGB` which defines the size of disk volume used for `/var/lib` mount point.
