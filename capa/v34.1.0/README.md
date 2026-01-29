# :zap: Giant Swarm Release v34.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v34.0.0

### Components

- Flatcar from v4459.2.2 to [v4459.2.3](https://www.flatcar.org/releases/#release-4459.2.3)
- os-tooling from v1.26.3 to v1.26.4

### Apps

- aws-ebs-csi-driver from v3.4.1 to v4.1.0
- etcd-defrag from v1.2.3 to v1.2.4
- karpenter from v1.4.0 to v2.0.0
- observability-bundle from v2.5.0 to v2.6.0

### aws-ebs-csi-driver [v3.4.1...v4.1.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.4.1...v4.1.0)

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

- Change IAM role name for the ebs-csi-driver-controller, to differentiate it from the old one managed by the iam-operator.
- Remove dependency for the cloud-provider-aws in the aws-ebs-csi-driver HelmRelease. That dependency should be set in the bundle HelmRelease by the provider cluster chart
- Update CircleCI configuration to push both app and bundle charts
- Update README with bundle architecture documentation

#### Fixed

- Fix boolean type of the expansion
- Allow volume expansion by default on gp3

### etcd-defrag [v1.2.3...v1.2.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.3...v1.2.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.37.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/78))

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
