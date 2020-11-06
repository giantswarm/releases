# :zap: Giant Swarm Release v12.3.2 for KVM :zap:

**Nodes will be rolled during upgrade to this version.**

This patch release adds registry credentials to prevent an issue with image pulling rate limits recently introduced by Docker Hub.

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

Before upgrading, please ensure cluster is on [KVM 12.1.x](https://github.com/giantswarm/releases/tree/master/kvm/v12.1.0) platform release first.

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.

**Note for future 12.x.x releases:**

Please ensure cluster is on [KVM 12.1.x](https://github.com/giantswarm/releases/tree/master/kvm/v12.1.0) platform release first before upgrading to 12.2.0+ Please persist this note until all customers are on KVM 12.1.x and above.

## Change details

### kvm-operator [v3.13.0](https://github.com/giantswarm/kvm-operator/blob/master/CHANGELOG.md#3130---2020-10-30)
- Update `k8scloudconfig` to v7.2.0, containing a fix for DockerHub QPS.
