# :zap: Giant Swarm Release v33.0.0 for CAPA :zap:

> **WARNING:** This release enables the `OwnerReferencesPermissionEnforcement` admission controller by default. This means that only users with delete permission to an object can change its `metadata.ownerReferences`, and only users with update permission to the finalizers subresource of the referenced owner can change `metadata.ownerReferences[x].blockOwnerDeletion`. If you have workloads that need to modify these fields, please ensure that the necessary RBAC permissions are in place before upgrading to this release.

Example:
```yaml
- apiGroups: ["<group>"]
  resources: ["<resource>", "<resource>/finalizers"]
  verbs: ["delete", "..."]  # Add any additional verbs your use case requires
```

## Changes compared to v32.0.0

### Components

- cluster-aws from v5.0.0 to v6.0.0
- Flatcar from v4230.2.2 to [v4230.2.3](https://www.flatcar-linux.org/releases/#release-4230.2.3)
- Kubernetes from v1.32.9 to [v1.33.5](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.33.md#v1.33.5)

### cluster-aws [v5.0.0...v6.0.0](https://github.com/giantswarm/cluster-aws/compare/v5.0.0...v6.0.0)

#### Added

- Add standard tags to IRSA infrastructure.
- Expose value to configure `terminationGracePeriod` in the karpenter node pools.

#### Changed

- Chart: Update `cluster` to v4.2.0.
- The container registry passed as value to default apps is set to `gsoci.azurecr.io`, regardless of the cluster region. The mirroring feature of `containerd` will make sure the right registry is used.
- Switch to HelmReleases to install `karpenter` and `karpenter-crossplane-resources` charts.
- Bump flux `HelmReleases` api version to v2.
- Reduce heartbeat timeout for ASG lifecycle hooks to from 30 minutes to 3 minutes since aws-node-termination-handler-app (NTH) can now send heartbeats
- Configure the following `startupTaints` to help `karpenter` ignore pending `Pods` due to these taints that will be removed after the node starts, avoiding unnecessary instance provisioning:
  - `node.cluster.x-k8s.io/uninitialized:NoSchedule`
  - `node.cilium.io/agent-not-ready:NoSchedule`
  - `ebs.csi.aws.com/agent-not-ready:NoExecute`
- Include `cilium` ENI mode pod CIDRs in the NodePort Services security group ingress rules.

#### Removed

- Removed `capi-node-labeler` app. From now on, the worker nodes won't have the `node-role.kubernetes.io/worker` or `node.kubernetes.io/worker` labels.

### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.2.0
- aws-nth-bundle from v1.2.2 to v1.3.0
- cert-exporter from v2.9.9 to v2.9.12
- cert-manager from v3.9.2 to v3.9.3
- cilium from v1.3.0 to v1.3.1
- cloud-provider-aws from v1.32.3 to v1.33.2-1
- cluster-autoscaler from v1.32.2-gs1 to v1.33.1-2
- coredns from v1.27.0 to v1.28.2
- etcd-defrag from v1.0.8 to v1.2.1
- etcd-k8s-res-count-exporter from v1.10.7 to v1.10.9
- k8s-audit-metrics from v0.10.6 to v0.10.8
- Added karpenter v1.3.0
- Added karpenter-crossplane-resources v0.4.0
- node-exporter from v1.20.5 to v1.20.7
- observability-bundle from v2.2.2 to v2.3.2
- vertical-pod-autoscaler from v6.0.1 to v6.1.1
- vertical-pod-autoscaler-crd from v4.0.1 to v4.1.1

### aws-ebs-csi-driver [v3.0.5...v3.2.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.2.0)

#### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.
- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### aws-nth-bundle [v1.2.2...v1.3.0](https://github.com/giantswarm/aws-nth-bundle/compare/v1.2.2...v1.3.0)

#### Changed

- Upgrade aws-nth-crossplane-resources to v1.3.0, fixing support for multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters, and supporting heartbeat sending
- Upgrade aws-node-termination-handler-app to v1.23.0, enabling heartbeats by default and upgrading to upstream application version v1.25.2 which fixes a resource leak bug relevant to heartbeat sending
- Upgrade aws-nth-crossplane-resources to v1.1.1, supporting multiple OIDC providers in the NTH IAM role as required for cleanup of migrated vintage clusters

### cert-exporter [v2.9.9...v2.9.12](https://github.com/giantswarm/cert-exporter/compare/v2.9.9...v2.9.12)

#### Changed

- Go: Update dependencies.
- Chart: Add value to toggle creation of Daemonset resources.
- Go: Update dependencies.

### cert-manager [v3.9.2...v3.9.3](https://github.com/giantswarm/cert-manager-app/compare/v3.9.2...v3.9.3)

#### Changed

- Fix missing targetPort in `cainjector-service`

### cilium [v1.3.0...v1.3.1](https://github.com/giantswarm/cilium-app/compare/v1.3.0...v1.3.1)

#### Changed

- Upgrade Cilium to [v1.18.2](https://github.com/cilium/cilium/releases/tag/v1.18.2).

### cloud-provider-aws [v1.32.3...v1.33.2-1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.32.3...v1.33.2-1)

#### Changed

- Chart: Update to upstream v1.33.2.

### cluster-autoscaler [v1.32.2-gs1...v1.33.1-2](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.32.2-gs1...v1.33.1-2)

#### Changed

- Chart: Grant access to `VolumeAttachments`. ([#345](https://github.com/giantswarm/cluster-autoscaler-app/pull/345))
- Update Kyverno API to v2 for policy exceptions
- Chart: Update to upstream v1.33.1.

### coredns [v1.27.0...v1.28.2](https://github.com/giantswarm/coredns-app/compare/v1.27.0...v1.28.2)

#### Changed

- Update `coredns` image to [1.13.1](https://github.com/coredns/coredns/releases/tag/v1.13.1).
- Add value to toggle creation of controlplane deployment.
- Update `coredns` image to [1.13.0](https://github.com/coredns/coredns/releases/tag/v1.13.0).

### etcd-defrag [v1.0.8...v1.2.1](https://github.com/giantswarm/etcd-defrag-app/compare/v1.0.8...v1.2.1)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.34.0. ([#62](https://github.com/giantswarm/etcd-defrag-app/pull/62))
- Chart: Update dependency ahrtr/etcd-defrag to v0.33.0. ([#60](https://github.com/giantswarm/etcd-defrag-app/pull/60))
- Update Kyverno API to v2 for policy exceptions
- Chart: Update dependency ahrtr/etcd-defrag to v0.32.0. ([#57](https://github.com/giantswarm/etcd-defrag-app/pull/57))

### etcd-k8s-res-count-exporter [v1.10.7...v1.10.9](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.7...v1.10.9)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.6...v0.10.8](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.6...v0.10.8)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### karpenter [v1.3.0](https://github.com/giantswarm/karpenter-app/releases/tag/v1.3.0)

#### Changed

- Updated karpenter to 1.7.1

### karpenter-crossplane-resources [v0.4.0](https://github.com/giantswarm/karpenter-crossplane-resources/releases/tag/v0.4.0)

#### Changed

- Add `iam:ListInstanceProfiles` for release 1.7.1

### node-exporter [v1.20.5...v1.20.7](https://github.com/giantswarm/node-exporter-app/compare/v1.20.5...v1.20.7)

#### Changed

- Go: Update dependencies.
- Update Kyverno API to v2 for policy exceptions
- Go: Update dependencies.

### observability-bundle [v2.2.2...v2.3.2](https://github.com/giantswarm/observability-bundle/compare/v2.2.2...v2.3.2)

#### Added

- Add KSM metrics for cloudnative-pg Cluster objects

#### Changed

- Update alloy-app to 0.15.0
  - Bumps alloy to 1.11.0

#### Fixed

- Update alloy-app to 0.15.1
  - Bumps alloy to 1.11.2

### vertical-pod-autoscaler [v6.0.1...v6.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v6.0.1...v6.1.1)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v11.1.1. ([#375](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/375))
- Chart: Update Helm release vertical-pod-autoscaler to v11.1.0. ([#372](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/372))

### vertical-pod-autoscaler-crd [v4.0.1...v4.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/compare/v4.0.1...v4.1.1)

#### Changed

- Chart: Sync to upstream. ([#166](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/166))
- Chart: Sync to upstream. ([#164](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/164))
