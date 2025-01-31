# :zap: Giant Swarm Release v29.2.1 for vSphere :zap:

## Changes compared to v29.1.0

### Components

- cluster-vsphere from v0.66.0 to v0.68.1 [v0.68.0...v0.68.1](https://github.com/giantswarm/cluster-vsphere/compare/v0.68.0...v0.68.1)

#### Added

- Add `components.containerd` to the schema and values.

#### Changed

- `kube-vip`:
  - Update to `v0.8.9`.
  - Add IP CIDR to static pod manifest.
- Make cloud-provider-vsphere HelmRelease catalog configurable.

### Apps

- cloud-provider-vsphere-app from v1.11.0 to v1.12.0

#### Changed

- Update `kube-vip` helm chart to `v0.6.2`.
- Update `kube-vip` to `v0.8.4`.
- Add initContainer to CSI node registrar DaemonSet to wait for CRD to be registered.
