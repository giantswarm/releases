# :zap: Giant Swarm Release v34.4.1 for CAPA :zap:

## Changes compared to v34.4.0

### Components

- cluster-aws from v7.7.2 to v7.7.3

### cluster-aws [v7.7.2...v7.7.3](https://github.com/giantswarm/cluster-aws/compare/v7.7.2...v7.7.3)

#### Added

- Add `global.connectivity.certManager.createIamRole` toggle (default `true`) to let customers opt out of provisioning the cert-manager IAM role via crossplane and bring their own role.

