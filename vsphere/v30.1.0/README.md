# :zap: Giant Swarm Release v30.1.0 for vSphere :zap:

## Changes compared to v30.0.0

### Components

- cluster-vsphere from v0.69.0 to v1.0.0

### cluster-vsphere [v0.69.0...v1.0.0](https://github.com/giantswarm/cluster-vsphere/compare/v0.69.0...v1.0.0)

#### Added

- Add `global.providerSpecific.templateSuffix` to set a suffix on the VM template to use.

#### Changed

- Split cloud provider app into separate HelmReleases.

### Apps

- cloud-provider-vsphere from v1.12.0 to v2.0.1
- kube-vip added at v0.2.0
- kube-vip-cloud-provider added at v0.3.0
- vsphere-csi-driver added at v3.4.2

### cloud-provider-vsphere [v1.12.0...v2.0.1](https://github.com/giantswarm/cloud-provider-vsphere-app/compare/v1.12.0...v2.0.1)

#### Changed

- Remove subcharts in order to deploy only the vSphere CPI (at upstream version `v1.30.0`).

### kube-vip [v0.2.0](https://github.com/giantswarm/kube-vip-app/blob/main/CHANGELOG.md#020---2025-02-25)

#### Added

- Initial release which tracks upstream version `0.8.4`.

### kube-vip-cloud-provider [v0.3.0](https://github.com/giantswarm/kube-vip-cloud-provider-app/blob/main/CHANGELOG.md#030---2025-03-14)

#### Added

- Initial release which tracks upstream version `0.0.10`.

#### Changed

- Run container with a read-only filesystem.

### vsphere-csi-driver [v3.4.2](https://github.com/giantswarm/vsphere-csi-driver-app/blob/main/CHANGELOG.md#342---2025-03-17)

#### Added

- Add upstream chart at `v3.3.0`.

#### Changed

- Correct kubectl image tag.

#### Removed

- Remove superfluous update script.
