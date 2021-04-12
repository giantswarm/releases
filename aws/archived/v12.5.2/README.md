# :zap: Tenant Cluster Release v12.5.2 for AWS :zap:

**If you are upgrading from 12.5.1, upgrading to this release will not roll your nodes. It will only update the apps.**

This release fixes an issue that causes app-operator to crash when handling cluster deletion.

Because we rarely delete clusters, the likelihood of this issue is low. But if we do delete a cluster (for example, a test cluster) with the broken app-operator, the operator will crash and stop reconciling app CRs in other clusters. Everything in that version breaks.

As this is a patch update that doesn't roll nodes, we highly recommended upgrading to it.

## Change details

### app-operator [2.3.5](https://github.com/giantswarm/app-operator/releases/tag/v2.3.5)

#### Fixed
- Skip removing finalizer for chart-operator chart CR if its not present.
- Skip deleting chart-operator in case of cluster deletion.
