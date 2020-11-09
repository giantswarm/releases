# :zap: Giant Swarm Release v12.7.0 for AWS :zap:

This minor release adds an alpha feature that allows to configure values for AWS CF ASG UpdatePolicy values. Check guide [Fine-tuning upgrade disruption on AWS](https://docs.giantswarm.io/guides/fine-tuning-upgrade-disruption-on-aws/) for more information on this feature.

## Change details

### aws-operator [9.3.0](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.0)

#### Added

- Annotation `alpha.aws.giantswarm.io/update-max-batch-size` to configure max batch size in ASG update policy on cluster or machine deployment CR.
- Annotation `alpha.aws.giantswarm.io/update-pause-time` to configure pause between batches in ASG update on cluster or machine deployment CR.
