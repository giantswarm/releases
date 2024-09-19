# :zap: Giant Swarm Release v26.1.1 for CAPA :zap:

This release introduces several changes that are required for Vintage to CAPA migration use-cases.

Most notable change is that now `auditd` is *disabled* by default. If you actively use this feature, please add the following field `global.components.auditd.enabled` set to `true` in the Cluster App user values before the upgrade.

## Changes compared to v26.1.0

### Components

- cluster-aws from v1.3.0 to v1.3.2

### cluster-aws [v1.3.0...v1.3.2](https://github.com/giantswarm/cluster-aws/compare/v1.3.0...v1.3.2)

#### Added

- Chart: Add `global.connectivity.network.pods.nodeCidrMaskSize` to schema.
- Chart: Allow to enable `auditd` through `global.components.auditd.enabled`.
- Chart: Support multiple service account issuers.

#### Changed

- Chart: Update `cluster` to v1.0.1.
  - Allow to enable `auditd` service through `global.components.auditd.enabled`.
  - Support multiple service account issuers.
  - Allow configuring kube-controller-manager `--node-cidr-mask-size` flag.
