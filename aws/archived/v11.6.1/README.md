# :zap: Tenant Cluster Release v11.6.1 for AWS :zap:

**Nodes will be rolled during upgrade to this version.**

This patch release prevents an issue with QPS (Queries per Second) limits introduced by Docker Hub.

## Change details


### aws-operator [8.7.10](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.10)

#### Fixed
- Fix DockerHub QPS limits by using paid user token for pulls.



