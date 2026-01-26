# :zap: Giant Swarm Release v34.1.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v34.0.0

### Components

- os-tooling from v1.26.3 to v1.26.4

### Apps

- aws-ebs-csi-driver from v3.4.1 to v4.0.3
- etcd-defrag from v1.2.3 to v1.2.4

### aws-ebs-csi-driver [v3.4.1...v4.0.3](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v3.4.1...v4.0.3)

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

- Remove dependency for the cloud-provider-aws in the aws-ebs-csi-driver HelmRelease. That dependency should be set in the bundle HelmRelease by the provider cluster chart
- Update CircleCI configuration to push both app and bundle charts
- Update README with bundle architecture documentation

#### Fixed

- Fix boolean type of the expansion
- Allow volume expansion by default on gp3

### etcd-defrag [v1.2.3...v1.2.4](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.3...v1.2.4)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.37.0. ([#78](https://github.com/giantswarm/etcd-defrag-app/pull/78))
