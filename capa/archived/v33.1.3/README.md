# :zap: Giant Swarm Release v33.1.3 for CAPA :zap:

Fix Karpenter schema for extraConfigs

## Changes compared to v33.1.2

### Components

- cluster-aws from v6.4.2 to v6.4.3

### cluster-aws [v6.4.2...v6.4.3](https://github.com/giantswarm/cluster-aws/compare/v6.4.2...v6.4.3)

#### Fixed

- Fix Karpenter schema: Use `helmRelease` schema instead of `app` schema. This corrects the `extraConfigs[].kind` field to accept `ConfigMap` and `Secret` (capitalized), and replaces the `priority` field with `optional` field, matching the HelmRelease resource structure.
