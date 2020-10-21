# :zap: Giant Swarm Release v12.1.1 for Azure :zap:

**If you are upgrading from 12.1.0, upgrading to this release will not roll your nodes.**

This patch release fixes a problem causing the accidental deletion and reinstallation of Preinstalled Apps (such as CoreDNS) in 12.x.x tenant clusters.

Please upgrade all older clusters to this version in order to prevent possible downtime. 

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on Azure v12.1.x and above.

## Change details

### cluster-operator [0.23.17](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02317---2020-10-19)
- Remove all chartconfig migration logic that caused accidental deletion and is no longer needed.

### app-operator [2.3.5](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#235---2020-10-20)
- Add resource version for chart configmaps and secrets to the chart CR to reduce latency of update events.
- Fix YAML comparison for chart configmaps and secrets.

### chart-operator [2.3.5](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#235---2020-10-13)
- Updated Helm to [v3.3.4](https://github.com/helm/helm/releases/tag/v3.3.4).
- Added event count metrics for delete, install, rollback and update of Helm releases.
