# :zap: Giant Swarm Release v12.5.1 for AWS :zap:

This release fixes an issue that prevented upgrades of the Control Planes.

## Change details


### aws-operator [9.1.3](https://github.com/giantswarm/aws-operator/releases/tag/v9.1.3)

#### Fixed
- Ignore error when missing APIServerPublicLoadBalancer CF Stack output to allow upgrade.
