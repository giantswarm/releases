# :zap: Giant Swarm Release v32.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.1.0
- aws-nth-bundle from v1.2.2 to v1.2.3
- cilium from v1.3.0 to v1.3.1
- vertical-pod-autoscaler from v6.0.1 to v6.1.0
- vertical-pod-autoscaler-crd from v4.0.1 to v4.1.0

### aws-ebs-csi-driver [v3.0.5...v3.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.1.0)

#### Changed

- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### aws-nth-bundle [v1.2.2...v1.2.3](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.2...v1.2.3)

#### Changed

- Upgrade aws-nth-crossplane-resources to v1.1.1, supporting multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### vertical-pod-autoscaler [v6.0.1...v6.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.0.1...v6.1.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.1.0. ([#372](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/372))

### vertical-pod-autoscaler-crd [v4.0.1...v4.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.0.1...v4.1.0)

#### Changed

- Chart: Sync to upstream. ([#164](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/164))
