# :zap: Giant Swarm Release v36.0.0 for CAPA :zap:

## Changes compared to v35.0.0

### Components

- cluster-aws from v9.0.1-dev.migrate-app--elmreleases.2026-08-19.15-46-00.h84cf66d to v9.0.0
- Kubernetes from v1.35.7 to [v1.36.3](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.36.md#v1.36.3)

### Apps

- aws-ebs-csi-driver from v4.3.1-dev.zyzju.2026-08-19.15-15-03.h758f4fc to v4.3.0
- karpenter from v2.4.1-dev.iptkq.2026-08-19.15-34-36.h90bb246 to v2.4.0

### aws-ebs-csi-driver [v4.3.1-dev.zyzju.2026-08-19.15-15-03.h758f4fc...v4.3.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/compare/v4.3.1-dev.zyzju.2026-08-19.15-15-03.h758f4fc...v4.3.0)

#### Added

- Add IRSA environment variables (`AWS_ROLE_ARN`, `AWS_WEB_IDENTITY_TOKEN_FILE`), projected ServiceAccountToken volume, and `AWS_REGION` to the EBS CSI controller, enabling IRSA authentication in CAPA clusters.
- Propagate proxy values from the bundle (`proxy.http`, `proxy.noProxy`) to the upstream chart (`proxy.http_proxy`, `proxy.no_proxy`) when set.

#### Fixed

- Re-enable metrics, force use of `ServiceMonitor` to avoid rendering without them if CRDs are not installed yet
- Fix VPA `updateMode` for `ebs-csi-node` DaemonSet from `Auto` to `Initial`. VPA cannot evict DaemonSet pods, so `Auto` mode silently produces recommendations without ever applying them. `Initial` correctly sets resources at pod creation time.

### karpenter [v2.4.1-dev.iptkq.2026-08-19.15-34-36.h90bb246...v2.4.0](https://github.com/giantswarm/karpenter-app/compare/v2.4.1-dev.iptkq.2026-08-19.15-34-36.h90bb246...v2.4.0)

#### Added

- Add `cluster.x-k8s.io/cluster-name` label to the karpenter HelmRelease.
- Add `iam:GetInstanceProfile` permission to Karpenter IAM role.
- Add karpenter CRDs.
- Set `helm.sh/resource-policy: keep` on the karpenter CRDs so they survive HelmRelease uninstall and prevent cascade-deleting `NodePool`/`NodeClaim`/`EC2NodeClass` resources.

#### Changed

- Switch e2e scale test from App CR to Flux HelmRelease for deploying hello-world, avoiding `values-schema-violation` errors caused by app-platform injected properties.
- Improve Crossplane ConfigMap fetching logic
