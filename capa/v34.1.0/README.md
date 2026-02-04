# :zap: Giant Swarm Release v34.1.0 for CAPA :zap:

## Changes compared to v34.0.0

### Components

- cluster-aws from v7.2.5 to v7.3.0
- Flatcar from v4459.2.2 to [v4459.2.3](https://www.flatcar.org/releases/#release-4459.2.3)
- Kubernetes from v1.34.3 to [v1.34.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.4)
- os-tooling from v1.26.3 to v1.26.4

### cluster-aws [v7.2.5...v7.3.0](https://github.com/giantswarm/cluster-aws/compare/v7.2.5...v7.3.0)

#### Added

- Add JSON schema validation patterns for `global.providerSpecific.region`.
- Add JSON schema validation patterns for `global.providerSpecific.awsAccountId`.
- Add JSON schema validation patterns for `global.controlPlane.instanceType` and node pool `instanceType`.
- Add JSON schema `maxLength: 20` constraint for `global.metadata.name`, aligning with the constraint enforced by [our kyverno policies](https://github.com/giantswarm/kyverno-policies-ux/blob/main/policies/ux/cluster-names.yaml).

#### Changed

- Install the `aws-ebs-csi-driver-bundle` that contains the `aws-ebs-csi-driver` app, together with the crossplane resources to manage the AWS IAM Roles required by the app.
- Install the `karpenter-bundle` that contains the `karpenter` app, together with the crossplane custom resources to manage the AWS resources required by `karpenter`.
- Use `cluster` chart values for Karpenter kubelet `systemReserved` and `kubeReserved` configuration instead of hardcoded values.
- Set correct `maxPods` value for karpenter node pools, based on the configured `nodeCidrMaskSize`, but capped at 110 pods.
- Always install the `karpenter-bundle`, regardless of whether karpenter node pools are configured. This is useful when deleting karpenter node pools, because otherwise the karpenter app was being removed and karpenter did not have time to clean up the node pools.

#### Fixed

- Install node-termination-handler bundle even if falling back to default node pools. No workers could come up without NTH, so `nodePools: {}` (= use default node pools) did not create a working cluster.

### Apps

- aws-ebs-csi-driver from v3.4.1 to v4.1.1
- cert-exporter from v2.9.15 to v2.9.16
- cilium from v1.3.4 to v1.4.0
- cluster-autoscaler from v1.34.1-1 to v1.34.2-1
- etcd-defrag from v1.2.3 to v1.2.4
- etcd-k8s-res-count-exporter from v1.10.12 to v1.10.13
- k8s-audit-metrics from v0.10.11 to v0.10.12
- k8s-dns-node-cache from v2.9.1 to v2.9.2
- karpenter from v1.4.0 to v2.0.0
- observability-bundle from v2.5.0 to v2.6.0
- security-bundle from v1.16.1 to v1.17.0

### aws-ebs-csi-driver [v3.4.1...v4.1.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.4.1...v4.1.1)

#### Added

- Introduce bundle chart architecture with Crossplane IAM resources.
  - Add `aws-ebs-csi-driver-app-bundle` chart that includes:
- Crossplane IAM Role with EBS CSI driver permissions
- Flux HelmRelease to deploy the workload cluster chart
- ConfigMap for values passthrough
  - Bundle chart is installed on the management cluster and deploys the app chart to the workload cluster
  - IAM role uses OIDC federation (IRSA) and reads configuration from `<clusterID>-crossplane-config` ConfigMap
  - Both charts share the same version and are released together

#### Changed

- Refactor crossplane config data retrieval. Fail installation if the ConfigMap can't be found, otherwise the chart was creating invalid IAM roles.
- Change IAM role name for the ebs-csi-driver-controller, to differentiate it from the old one managed by the iam-operator.
- Remove dependency for the cloud-provider-aws in the aws-ebs-csi-driver HelmRelease. That dependency should be set in the bundle HelmRelease by the provider cluster chart
- Update CircleCI configuration to push both app and bundle charts
- Update README with bundle architecture documentation

#### Fixed

- Fix boolean type of the expansion
- Allow volume expansion by default on gp3

### cert-exporter [v2.9.15...v2.9.16](https://github.com/giantswarm/cert-exporter/compare/v2.9.15...v2.9.16)

#### Changed

- Go: Update dependencies.

### cilium [v1.3.4...v1.4.0](https://github.com/giantswarm/cilium-app/compare/v1.3.4...v1.4.0)

#### Changed

- Upgrade Cilium to [v1.19.0](https://github.com/cilium/cilium/releases/tag/v1.19.0).
- Update chart icon to use Giant Swarm-hosted Cilium icon.

### cluster-autoscaler [v1.34.1-1...v1.34.2-1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.34.1-1...v1.34.2-1)

#### Changed

- Chart: Update to upstream v1.34.2.

### etcd-defrag [v1.2.3...v1.2.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.3...v1.2.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.37.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/78))

### etcd-k8s-res-count-exporter [v1.10.12...v1.10.13](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/compare/v1.10.12...v1.10.13)

#### Changed

- Migrate to App Build Suite (ABS) for Helm chart building.
- Go: Update dependencies.

### k8s-audit-metrics [v0.10.11...v0.10.12](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.10.11...v0.10.12)

#### Changed

- Migrate to App Build Suite (ABS) for Helm chart building.
- Go: Update dependencies.

### k8s-dns-node-cache [v2.9.1...v2.9.2](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.9.1...v2.9.2)

#### Changed

- Upgrade application to version 1.26.7 (includes coredns 1.13.1)

### karpenter [v1.4.0...v2.0.0](https://github.com/giantswarm/karpenter-app/compare/v1.4.0...v2.0.0)

#### Added

- Add e2e tests for this app.
- Add `karpenter-bundle` chart that consolidates `karpenter-app` and `karpenter-crossplane-resources` into a single deployable bundle. The bundle includes:
  - HelmRelease and OCIRepository for deploying karpenter to workload clusters
  - IAM roles for karpenter and nodeclassgenerator via Crossplane
  - SQS queue and CloudWatch event rules for interruption handling

#### Fixed

- Use only `clustertest` v3 instead of v2 and v3. We also upgraded to `apptest-framework` v3 due to this.

### observability-bundle [v2.5.0...v2.6.0](https://github.com/giantswarm/observability-bundle/compare/v2.5.0...v2.6.0)

#### Added

- Add KSM metrics for Gateway API resources

### security-bundle [v1.16.1...v1.17.0](https://github.com/giantswarm/security-bundle/compare/v1.16.1...v1.17.0)

#### Changed

- Update `kyverno` (app) to v0.23.0.
- Update `kyverno-crds` (app) to v1.16.0.
- Update `reports-server` (app) to v0.1.0.
- Update `cloudnative-pg` (app) to v0.0.13.
- Update `kubescape` (app) to v0.0.5.
- Update `starboard-exporter` (app) to v1.0.2.
