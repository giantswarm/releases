# :zap: Giant Swarm Release v35.4.0 for CAPA :zap:

## Changes compared to v35.3.0

### Components

- cluster-aws from v10.5.0 to v10.6.0
- Kubernetes from v1.35.9 to [v1.35.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.35.md#v1.35.10)

### cluster-aws [v10.5.0...v10.6.0](https://github.com/giantswarm/cluster-aws/compare/v10.5.0...v10.6.0)

#### Added

- Add `global.nodePools.<pool>.replicas` for Karpenter node pools to keep a fixed number of nodes (static node pool). This is a tech preview of the alpha Karpenter `StaticCapacity` feature and requires a release with Karpenter v1.14.1 and the matching `aws-resolver-rules-operator`.
- Add `global.nodePools.<pool>.limits.nodes` to cap the number of nodes of a Karpenter node pool.
- Add the `Balanced` consolidation policy for Karpenter node pools.

#### Fixed

- Default `limits.cpu` and `limits.memory` of Karpenter node pools separately, so setting only one of them no longer renders the other one empty.

### Apps

- cilium from v1.5.2 to v1.5.4

### cilium [v1.5.2...v1.5.4](https://github.com/giantswarm/cilium-app/compare/v1.5.2...v1.5.4)

#### Changed

- Upgrade Cilium to [v1.19.10](https://github.com/cilium/cilium/releases/tag/v1.19.10).
- Upgrade Cilium to [v1.19.9](https://github.com/cilium/cilium/releases/tag/v1.19.9).
- Exclude the Lens namespace label `k8slens-edit-resource-version` from identity computation.
