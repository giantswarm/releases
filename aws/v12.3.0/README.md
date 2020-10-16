## :zap:  Giant Swarm Release 12.3.0 for AWS :zap:

**If you are upgrading from 12.2.0, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades all Helm releases managed by Giant Swarm to use [Helm v3.3.4](https://github.com/helm/helm/releases/tag/v3.3.4).

This lets us benefit from the improved security model and keep up to date with the community. We also remove the Tiller deployment from the giantswarm namespace, removing its gRPC endpoint, which reduces operational complexity.

If you are still using Helm 2 then these Helm releases will not be affected. However we encourage you to upgrade to Helm 3. As Helm 2 support ends on November 13th 2020. https://helm.sh/blog/helm-v2-deprecation-timeline/

Below, you can find more details on components that were changed with this release.

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

- Helm3: 
  - Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.3.x** and above.
- cert-manager-app: 
  - Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.1.x** and above.

## Change details

### app-operator [v2.3.2](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#232---2020-09-29)
### chart-operator [v2.3.3](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#233---2020-09-29)
- Updated Helm to [v3.3.4](https://github.com/helm/helm/releases/tag/v3.3.4).
