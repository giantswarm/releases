# :zap: Tenant Cluster Release v12.6.0 for AWS :zap:

This patch release prevents an issue with QPS (Queries per Second) limits introduced by Docker Hub. Also, it solves a corner case scenario during ETCD mouting time.

This minor release also contains two alpha features to terminate unhealthy nodes and to use the new AWS metadata API v2. Both only works when the cluster CR is annotated properly.

## Change details

### aws-operator [9.2.0](https://github.com/giantswarm/aws-operator/releases/tag/v9.2.0)

#### Fixed

- Fix dockerhub QPS by using paid user token for pulls.
- Remove dependency on `var-lib-etcd.automount` to avoid dependency cycle on
  new systemd.

#### Added

- Add `terminate-unhealthy-node` alpha feature to automatically terminate bad and
  unhealthy nodes in a Cluster.
- Add `alpha.giantswarm.io/aws-metadata-v2` annotation to enable AWS Metadata API v2.

