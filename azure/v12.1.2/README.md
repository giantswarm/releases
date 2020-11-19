# :zap: Tenant Cluster Release v12.1.2 for Azure :zap:

**Nodes will be rolled during upgrade to this version.**

This patch release prevents an issue with QPS (Queries per Second) limits introduced by Docker Hub.

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on Azure v12.1.x and above.

## Change details

### azure-operator [4.3.0](https://github.com/giantswarm/azure-operator/compare/v4.2.0...v4.3.0)
- Pass dockerhub token for kubelet authorized image pulling.
- Update `k8scloudconfig` to v7.2.0, containing a fix for Docker Hub QPS.
