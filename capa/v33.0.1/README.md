# :zap: Giant Swarm Release v33.0.1 for CAPA :zap:

This release improves the stability of Karpenter node pools.

## Changes compared to v33.0.0

### Components

- cluster-aws from v6.0.0 to v6.2.0

### cluster-aws [v6.0.0...v6.2.0](https://github.com/giantswarm/cluster-aws/compare/v6.0.0...v6.2.0)

#### Added

- Add `capa-karpenter-taint-remover` to handle CAPA - Karpenter taint race condition.

#### Changed

- Change default consolidation time to 6 hours to avoid constant node rolling.
- Rename `capa-karpenter-taint-remover` app.
- Set `terminationGracePeriod` default to 30m, to avoid having `karpenter` nodes stuck in `Deleting` state due to `Pods` blocking the deletion i.e. PDBs.

### Apps

- aws-pod-identity-webhook from v1.19.1 to v2.0.0
- karpenter from v1.3.0 to v1.4.0
- Added karpenter-taint-remover v1.0.1
- security-bundle from v1.12.0 to v1.13.1

### aws-pod-identity-webhook [v1.19.1...v2.0.0](https://github.com/giantswarm/aws-pod-identity-webhook/compare/v1.19.1...v2.0.0)

#### Changed

- Upgrade IRSA to latest v0.6.9

### karpenter [v1.3.0...v1.4.0](https://github.com/giantswarm/karpenter-app/compare/v1.3.0...v1.4.0)

#### Changed

- Updated karpenter to 1.8.1
- Fixes RBAC issues when OwnerReferencesPermissionEnforcement featuregate is enabled by allowing finalizers sub'resource modification.

### karpenter-taint-remover [v1.0.1](https://github.com/giantswarm/capa-karpenter-taint-remover/releases/tag/v1.0.1)

#### Changed

- Use default catalog

### security-bundle [v1.12.0...v1.13.1](https://github.com/giantswarm/security-bundle/compare/v1.12.0...v1.13.1)

#### Changed

- Revert previous `kyverno` update ([#536](https://github.com/giantswarm/security-bundle/pull/536), [#531](https://github.com/giantswarm/security-bundle/pull/531), [#538](https://github.com/giantswarm/security-bundle/pull/538)).
- Update `kyverno-policy-operator` (app) to v0.1.6.
- Update `kyverno` (app) to v0.20.0.
- Update `kyverno-crds` (app) to v1.14.0.
- Update `kyverno-policies` (app) to v0.24.0.
- Update `kyverno-policy-operator` (app) to v0.1.5.
- Update `trivy-operator` (app) to v0.12.1.
- Update `trivy` (app) to v0.14.1.
- Update `falco` (app) to v0.11.0.
