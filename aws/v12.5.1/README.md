# :zap: Giant Swarm Release v12.5.1 for AWS :zap:

This release fixes an issue that prevented upgrades of the Control Planes.

## Change details


### aws-operator [9.1.3](https://github.com/giantswarm/aws-operator/releases/tag/v9.1.3)

#### Fixed
- Ignore error when missing APIServerPublicLoadBalancer CF Stack output to allow upgrade.



### app-operator [2.3.3](https://github.com/giantswarm/app-operator/releases/tag/v2.3.3)

#### Added
- Delete chart-operator helm release and chart CR so it can be re-installed.



### chart-operator [2.3.5](https://github.com/giantswarm/chart-operator/releases/tag/v2.3.5)

#### Fixed
- Stop repeating helm upgrade for the failed helm release.
