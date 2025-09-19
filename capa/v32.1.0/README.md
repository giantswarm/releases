# :zap: Giant Swarm Release v32.1.0 for CAPA :zap:

This major release brings significant platform enhancements and stability improvements across all supported providers.

## Changes compared to v32.0.0

### Components

- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.1.0

### aws-ebs-csi-driver [v3.0.5...v3.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.1.0)

#### Changed

- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.
