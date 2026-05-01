# :zap: Giant Swarm Release v34.3.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v34.2.0

### Components

- Flatcar from v4459.2.4 to [v4593.2.0](https://www.flatcar.org/releases/#release-4593.2.0)

### Apps

- aws-ebs-csi-driver from v4.1.2 to v4.2.0
- karpenter from v2.3.0 to v2.4.0

### aws-ebs-csi-driver [v4.1.2...v4.2.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.1.2...v4.2.0)

#### :warning: Breaking Changes

- **Workload chart renamed** from `aws-ebs-csi-driver-app` to `aws-ebs-csi-driver`. The OCI catalog artifact name changes accordingly.
- **Bundle values restructured**: upstream chart values are now under the `upstream:` key in the bundle `values.yaml`. The `giantswarm.workloadValues` helper handles the transformation automatically, so bundle users only need to place overrides under `upstream:` in their App CR ConfigMap.
- **Direct workload chart install**: if installing the workload chart directly (without the bundle), all upstream values must be under the `upstream:` key, and extras (`verticalPodAutoscaler`, `networkPolicy`, `global.podSecurityStandards`) are at the top level.

#### Changed

- Add `io.giantswarm.application.audience: all` annotation to publish the app to the customer Backstage catalog.
- Migrate chart metadata annotations to `io.giantswarm.application.*` format for both the app and bundle charts.

### karpenter [v2.3.0...v2.4.0](https://github.com/giantswarm/karpenter-app/compare/v2.3.0...v2.4.0)

#### Added

- Add `cluster.x-k8s.io/cluster-name` label to the karpenter HelmRelease.
- Add `iam:GetInstanceProfile` permission to Karpenter IAM role.
- Add karpenter CRDs.
- Set `helm.sh/resource-policy: keep` on the karpenter CRDs so they survive HelmRelease uninstall and prevent cascade-deleting `NodePool`/`NodeClaim`/`EC2NodeClass` resources.

#### Changed

- Switch e2e scale test from App CR to Flux HelmRelease for deploying hello-world, avoiding `values-schema-violation` errors caused by app-platform injected properties.
- Improve Crossplane ConfigMap fetching logic
