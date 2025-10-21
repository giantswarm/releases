# :zap: Giant Swarm Release v33.0.1 for CAPA :zap:

This releases improves the stability of Karpenter node pools.

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

- karpenter from v1.3.0 to v1.4.0

### karpenter [v1.3.0...v1.4.0](https://github.com/giantswarm/karpenter-app/compare/v1.3.0...v1.4.0)

#### Changed

- Updated karpenter to 1.8.1
- Fixes RBAC issues when OwnerReferencesPermissionEnforcement featuregate is enabled by allowing finalizers sub'resource modification.
