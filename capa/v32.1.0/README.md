# :zap: Giant Swarm Release v32.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.1.0
- cilium from v1.3.0 to v1.3.1

### aws-ebs-csi-driver [v3.0.5...v3.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.1.0)

#### Changed

- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).
