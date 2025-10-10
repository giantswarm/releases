# :zap: Giant Swarm Release v32.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.0.0

### Components

- cluster-aws from v5.0.0 to v5.2.0
- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)

### cluster-aws [v5.0.0...v5.2.0](https://github.com/giantswarm/cluster-aws/compare/v5.0.0...v5.2.0)

#### Added

- Expose value to configure `terminationGracePeriod` in the karpenter node pools.

#### Changed

- Reduce heartbeat timeout for ASG lifecycle hooks to from 30 minutes to 3 minutes since aws-node-termination-handler-app (NTH) can now send heartbeats

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.2.0
- aws-nth-bundle from v1.2.2 to v1.3.0
- cert-exporter from v2.9.9 to v2.9.10
- cilium from v1.3.0 to v1.3.1
- Added cluster-aws v5.2.0
- etcd-defrag from v1.0.8 to v1.1.0
- etcd-k8s-res-count-exporter from v1.10.7 to v1.10.8
- k8s-audit-metrics from v0.10.6 to v0.10.7
- node-exporter from v1.20.5 to v1.20.6
- vertical-pod-autoscaler from v6.0.1 to v6.1.0
- vertical-pod-autoscaler-crd from v4.0.1 to v4.1.0

### aws-ebs-csi-driver [v3.0.5...v3.2.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.2.0)

#### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.
- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### aws-nth-bundle [v1.2.2...v1.3.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.2...v1.3.0)

#### Changed

- Upgrade aws-nth-crossplane-resources to v1.3.0, fixing support for multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters, and supporting heartbeat sending
- Upgrade aws-node-termination-handler-app to v1.23.0, enabling heartbeats by default and upgrading to upstream application version v1.25.2 which fixes a resource leak bug relevant to heartbeat sending
- Upgrade aws-nth-crossplane-resources to v1.1.1, supporting multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters

### cert-exporter [v2.9.9...v2.9.10](https://github.com/giantswarm/cert-exporter/compare/v2.9.9...v2.9.10)

#### Changed

- Go: Update dependencies.

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### cluster-aws [v5.2.0](https://github.com/giantswarm/cluster-aws/releases/tag/v5.2.0)

#### Changed

- Reduce heartbeat timeout for ASG lifecycle hooks to from 30 minutes to 3 minutes since aws-node-termination-handler-app (NTH) can now send heartbeats

### etcd-defrag [v1.0.8...v1.1.0](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.8...v1.1.0)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Chart: Update dependency ahrtr/etcd-defrag to v0.32.0. ([#57](https://github.com/giantswarm/etcd-defrag-app/pull/57))

### etcd-k8s-res-count-exporter [v1.10.7...v1.10.8](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.7...v1.10.8)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.6...v0.10.7](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.6...v0.10.7)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### node-exporter [v1.20.5...v1.20.6](https://github.com/giantswarm/node-exporter-app/compare/v1.20.5...v1.20.6)

#### Changed

- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### vertical-pod-autoscaler [v6.0.1...v6.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.0.1...v6.1.0)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.1.0. ([#372](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/372))

### vertical-pod-autoscaler-crd [v4.0.1...v4.1.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.0.1...v4.1.0)

#### Changed

- Chart: Sync to upstream. ([#164](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/164))
