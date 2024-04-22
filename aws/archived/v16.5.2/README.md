# :zap: Giant Swarm Release v16.5.2 for AWS :zap:

This is a patch release to add missing Elastic File System (EFS) IAM permissions. It also allows EFS CSI Driver to get the Instance Identity from the Instance Metadata Service (IMDS).

## Change details


### aws-operator [10.14.2](https://github.com/giantswarm/aws-operator/releases/tag/v10.14.2)

#### Fixed
- Added EFS policy to the ec2 instance role to allow to use the EFS driver out of the box



### kiam [2.5.1](https://github.com/giantswarm/kiam-app/releases/tag/v2.5.1)

#### Fixed
- Allow `whiteListRouteRegexp` to default to `/latest/*`.



