## :zap:  Giant Swarm Release 12.1.0 for Azure :zap:

**If you are upgrading from 12.0.2, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades all Helm releases managed by Giant Swarm to use [Helm v3.3.0](https://github.com/helm/helm/releases/tag/v3.3.0).

This lets us benefit from the improved security model and keep up to date with the community. We also remove the Tiller deployment from the giantswarm namespace, removing its gRPC endpoint, which reduces operational complexity.

If you are still using Helm 2 then these Helm releases will not be affected. However we encourage you to upgrade to Helm 3. As Helm 2 support ends on November 13th 2020. https://helm.sh/blog/helm-v2-deprecation-timeline/

Below, you can find more details on components that were changed with this release.

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, which we recommend to monitor to ensure safety.

**Note for Solution Engineers:**

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/index.md) as a guide to an upgrade to Azure v12.1.0.

## Change details

### app-operator [v2.1.1](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#211---2020-08-26)
### chart-operator [v2.3.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#230---2020-08-24)
- Updated Helm to [v3.3.0](https://github.com/helm/helm/releases/tag/v3.3.0).

### cert-exporter [v1.2.4](https://github.com/giantswarm/cert-exporter/blob/master/CHANGELOG.md#124---2020-08-13)
- Adjusted vault token format check for base62 tokens.
