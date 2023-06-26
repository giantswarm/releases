# :zap: Giant Swarm Release v17.4.3 for AWS :zap:

This release brings improvements to the aws-ebs-csi-driver app. It is especially important for customers running AWS clusters in China due to image pull issues.

## Change details


### aws-ebs-csi-driver [2.16.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.16.1)

#### Fixed
- Changing controller `httpEndpoint` to `8610` because of overlapping ports.
- Use global.image.registry as registry domain for all images.
- Bump aws-ebs-csi-driver version to v1.8.0.



