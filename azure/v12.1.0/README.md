## :zap:  Giant Swarm Release 12.1.0 for Azure :zap:

**If you are upgrading from 12.0.2, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades all Helm releases managed by Giant Swarm to use Helm v3.3.0.

This lets us benefit from the improved security and keep up to date with the community. We also removed the Tiller deployment from tenant clusters, removing its gRPC endpoint, which reduces operational complexity. 

Below, you can find more details on components that were changed with this release.

### app-operator [v2.1.1](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#211---2020-08-26)
### chart-operator [v2.3.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#230---2020-08-24)
- Updated Helm to v3.3.0.
