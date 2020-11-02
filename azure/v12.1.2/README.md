# :zap: Giant Swarm Release v12.1.2 for Azure :zap:

**Nodes will be rolled during upgrade to this version.**

This patch release prevents an issue with QPS limits introduced by DockerHub.

## Change details

### azure-operator [4.3.0](https://github.com/giantswarm/azure-operator/compare/v4.2.0...v4.3.0)
- Pass dockerhub token for kubelet authorized image pulling.
- Update `k8scloudconfig` to v7.2.0, containing a fix for DockerHub QPS.

