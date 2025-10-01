# :zap: Giant Swarm Release v33.0.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)
- Kubernetes from v1.32.9 to [v1.33.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.33.md#v1.33.5)

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.1.0
- cilium from v1.3.0 to v1.3.1
- cloud-provider-aws from v1.32.3 to v1.33.2-1
- cluster-autoscaler from v1.32.2-gs1 to v1.33.1-1

### aws-ebs-csi-driver [v3.0.5...v3.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.1.0)

#### Changed

- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### cloud-provider-aws [v1.32.3...v1.33.2-1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.32.3...v1.33.2-1)

#### Changed

- Chart: Update to upstream v1.33.2.

### cluster-autoscaler [v1.32.2-gs1...v1.33.1-1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.32.2-gs1...v1.33.1-1)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Chart: Update to upstream v1.33.1.
