# :zap: Giant Swarm Release v29.3.1 for CAPA :zap:

## Changes compared to v29.3.0

This release does not contain any changes to components or apps, but makes use of an cluster-aws chart which allows custom tags applied to EC2 instances only using the `global.providerSpecific.additionalNodeTags` property.

Expose the maxHealthyPercentage property to allow setting the maximum percentage of healthy machines in the Auto Scaling Group during upgrades.
