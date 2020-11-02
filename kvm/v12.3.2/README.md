# :zap: Giant Swarm Release v12.3.2 for KVM :zap:

**If you are upgrading from 12.3.0, upgrading to this release will not roll your nodes.**

This patch release prevents an issue with QPS limits introduced by DockerHub.

## Change details

### kvm-operator [v3.13.0](https://github.com/giantswarm/kvm-operator/blob/master/CHANGELOG.md#3130---2020-10-30)
- Update `k8scloudconfig` to v7.2.0, containing a fix for DockerHub QPS.

