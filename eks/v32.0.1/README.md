# :zap: Giant Swarm Release v32.0.1 for  :zap:

<< Add description here >>

## Changes compared to v32.0.0


### Apps

- aws-ebs-csi-driver from v3.0.5 to v3.3.0
- Added vertical-pod-autoscaler-crd v4.1.1

### aws-ebs-csi-driver [v3.0.5...v3.3.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.0.5...v3.3.0)

#### Changed

- Chart: Sync to upstream. ([#338](https://github.com/giantswarm/aws-ebs-csi-driver-app/pull/338))
  - Chart: Update AWS EBS CSI Driver from v1.41.0 to v1.51.0.
  - Chart: ⚠️ URGENT: XFS Compatibility Issue - Newly formatted XFS volumes may fail to mount on nodes with older kernels (Amazon Linux 2). Use `node.legacyXFS: true` as workaround.
  - Chart: ⚠️ URGENT: Controller Health Checks - Controller now performs AWS API dry-run checks. Ensure proper IAM permissions and network connectivity.
  - Chart: ⚠️ URGENT: StorageClass Parameter Deprecation* - `blockExpress` parameter is deprecated for `io2` volumes (now always uses 256,000 IOPS cap).
  - Chart: Add support for creating instant, point-in-time copies of EBS volumes within the same Availability Zone.
  - Chart: Add `debugLogs` parameter for maximum verbosity logging and debugging.
  - Chart: Add `metadataSources` configuration option for node metadata handling.
  - Chart: Add `disableMutation` parameter for service account mutation control.
  - Chart: Add support for updating node's max attachable volume count via `MutableCSINodeAllocatableCount` feature gate (Kubernetes 1.33+).
  - Chart: Update dependencies including AWS SDK, Prometheus, and various Go modules.
  - Chart: Add missing `enablePrometheusAnnotations` values for controller and node components.
  - Chart: Update sidecar container versions:
- csi-provisioner: v5.2.0 → v5.3.0
- csi-attacher: v4.8.1 → v4.9.0
- csi-snapshotter: v8.2.1 → v8.3.0
- livenessprobe: v2.14.0 → v2.16.0
- csi-resizer: v1.13.2 → v1.14.0
- csi-node-driver-registrar: v2.13.0 → v2.14.0
- volume-modifier-for-k8s: v0.5.1 → v0.8.0
- Configure `gsoci.azurecr.io` as the default container image registry.
- Set default `updateStrategy.rollingUpdate.maxUnavailable` to 25% in `DaemonSet` to speed up rolling update.

### vertical-pod-autoscaler-crd [v4.1.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v4.1.1)

#### Changed

- Chart: Sync to upstream. ([#166](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/166))
