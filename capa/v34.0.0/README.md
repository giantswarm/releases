# :zap: Giant Swarm Release v34.0.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v33.0.1

### Components

- cluster-aws from v6.2.0 to v6.4.0
- Flatcar from v4230.2.3 to [v4230.2.4](https://www.flatcar.org/releases/#release-4230.2.4)
- Kubernetes from v1.33.5 to [v1.34.1](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.1)
- os-tooling from v1.26.1 to v1.26.2

### cluster-aws [v6.2.0...v6.4.0](https://github.com/giantswarm/cluster-aws/compare/v6.2.0...v6.4.0)

#### Added

- Add node-problem-detector-app, disabled by default.

#### Changed

- Tidy up dependencies on `azs-getter`.
- Make `global.baseDomain` and `global.managementCluster` required values. These values will be passed to the chart when deploying it from the `cluster-app-installation-values` ConfigMap in the default namespace.
- Extract required values to its own central file to avoid repeating the `required` keyword and error messages. This is normally done automatically by a Kyverno policy.
- Change the default root disk size for Karpenter node pools. Karpenter will choose the cheapest instances, and certain instances, like `g6f.xlarge` come with some drivers that require a larger disk.
- Chart: Update `cluster` to v4.3.0.

### Apps

- aws-ebs-csi-driver from v3.2.0 to v3.3.0
- cert-exporter from v2.9.12 to v2.9.13
- cert-manager from v3.9.3 to v3.9.4
- etcd-defrag from v1.2.1 to v1.2.2
- etcd-k8s-res-count-exporter from v1.10.9 to v1.10.10
- k8s-audit-metrics from v0.10.8 to v0.10.9
- karpenter-crossplane-resources from v0.4.0 to v0.5.0
- node-exporter from v1.20.7 to v1.20.8
- security-bundle from v1.13.1 to v1.14.0

### aws-ebs-csi-driver [v3.2.0...v3.3.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.2.0...v3.3.0)

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

### cert-exporter [v2.9.12...v2.9.13](https://github.com/giantswarm/cert-exporter/compare/v2.9.12...v2.9.13)

#### Changed

- Go: Update dependencies.

### cert-manager [v3.9.3...v3.9.4](https://github.com/giantswarm/cert-manager-app/compare/v3.9.3...v3.9.4)

#### Added

- Add E2E tests using apptest-framework for automated PR testing across multiple providers (CAPA, CAPV, CAPZ, CAPVCD).
  - **Basic test suite**: Validates fresh installations
  - **Upgrade test suite**: Tests upgrade scenarios and certificate reconciliation
- Add certificate issuance integration test to cluster-test-suites.

#### Changed

- Upgrade cert-manager to v1.18.2.

### etcd-defrag [v1.2.1...v1.2.2](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.1...v1.2.2)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.35.0. ([#64](https://github.com/giantswarm/etcd-defrag-app/pull/64))

### etcd-k8s-res-count-exporter [v1.10.9...v1.10.10](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.9...v1.10.10)

#### Changed

- Go: Update dependencies.

### k8s-audit-metrics [v0.10.8...v0.10.9](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.8...v0.10.9)

#### Changed

- Go: Update dependencies.

### karpenter-crossplane-resources [v0.4.0...v0.5.0](https://github.com/giantswarm/karpenter-crossplane-resources/compare/v0.4.0...v0.5.0)

#### Added

- Add new Helm value to configure the workers IAM role. When Karpenter launches worker instances, it will attach the worker instance profile.

### node-exporter [v1.20.7...v1.20.8](https://github.com/giantswarm/node-exporter-app/compare/v1.20.7...v1.20.8)

#### Changed

- Go: Update dependencies.

### security-bundle [v1.13.1...v1.14.0](https://github.com/giantswarm/security-bundle/compare/v1.13.1...v1.14.0)

#### Changed

- Update `kyverno` (app) to v0.20.1.
- Update `kyverno-crds` (app) to v1.14.0.
- Update `kyverno-policies` (app) to v0.24.0.
- Update `reports-server` (app) to v0.0.3.
