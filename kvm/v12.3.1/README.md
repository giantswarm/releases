# :zap: Tenant Cluster Release v12.3.1 for KVM :zap:

**If you are upgrading from 12.3.0, upgrading to this release will not roll your nodes.**

This patch release fixes a problem causing the accidental deletion and reinstallation of Preinstalled Apps (such as CoreDNS) in 12.x.x tenant clusters.

Please upgrade all older clusters to this version in order to prevent possible downtime. 

**Note before upgrade:**

Please contact your Solution Engineer before upgrading. The upgrade is automated. However, it includes a data migration from Helm 2 release configmaps to Helm 3 release secrets, there are some pre-upgrade checks and we recommend monitoring the upgrade to ensure safety.

**Note for Solution Engineers:**

Before upgrading, please ensure cluster is on [KVM 12.1.x](https://github.com/giantswarm/releases/tree/master/kvm/v12.1.0) platform release first.

Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.

**Note for future 12.x.x releases:**

Please ensure cluster is on [KVM 12.1.x](https://github.com/giantswarm/releases/tree/master/kvm/v12.1.0) platform release first before upgrading to 12.2.0+ Please persist this note until all customers are on KVM 12.1.x and above.

## Change details

### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02318---2020-10-21)
- Remove all chartconfig migration logic that caused accidental deletion and is no longer needed.

### app-operator [2.3.5](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#235---2020-10-20)
- Fix YAML comparison for chart configmaps and secrets.
