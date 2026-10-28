# :zap: Giant Swarm Release v35.4.0 for CAPA :zap:

## Changes compared to v35.3.0

### Components

- cluster-aws from v10.5.0 to v10.6.0

### cluster-aws [v10.5.0...v10.6.0](https://github.com/giantswarm/cluster-aws/compare/v10.5.0...v10.6.0)

#### Added

- Add `global.nodePools.<pool>.replicas` for Karpenter node pools to keep a fixed number of nodes (static node pool). This is a tech preview of the alpha Karpenter `StaticCapacity` feature and requires a release with Karpenter v1.14.1 and the matching `aws-resolver-rules-operator`.
- Add `global.nodePools.<pool>.limits.nodes` to cap the number of nodes of a Karpenter node pool.
- Add the `Balanced` consolidation policy for Karpenter node pools.

#### Fixed

- Default `limits.cpu` and `limits.memory` of Karpenter node pools separately, so setting only one of them no longer renders the other one empty.

