# :zap: Giant Swarm Release v12.6.0 for AWS :zap:

This patch release prevents an issue with QPS (Queries per Second) limits introduced by Docker Hub.

## Change details


### aws-operator [9.2.0](https://github.com/giantswarm/aws-operator/releases/tag/v9.2.0)

#### Added
- Add `terminate-unhealthy-node` feature to automatically terminate bad and
  unhealthy nodes in a Cluster.
- Add `alpha.giantswarm.io/aws-metadata-v2` annotation to enable AWS Metadata
  API v2.
#### Fixed
- Fix dockerhub QPS by using paid user token for pulls.
- Remove dependency on `var-lib-etcd.automount` to avoid dependency cycle on
  new systemd.



