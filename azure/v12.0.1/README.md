## :zap:  Giant Swarm Release 12.0.1 for Azure :zap:

**If you are upgrading from 12.0.0, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades external-dns app to v0.7.3.

**Important Warning**
During master upgrade from 11.4.0 to 12.0.x, within the time frame of 30 seconds we had noticed a spike in requests failures. This is most likely caused by Azure CNI upgrade and despite our great efforts, we had not found a solution to maintain upgrade path and avoid this disturbance. Please keep this in mind when scheduling an upgrade window, and contact your SE if you have further questions.

**Notes for future 12.x.x releases:**
In order to proceed with upgrade, clusters must be using 11.4.0 or newer release version, otherwise the upgrade request will fail.
Pay attention to the increased failure rate of requests during upgrade of the master node and include this information in release notes if the problem persists.

Below, you can find more details on components that were changed with this release.

### external-dns v0.7.3 [Giant Swarm app 1.3.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#130---2020-08-18)

- Updated from v0.7.2. Check the [upstream changelog](https://github.com/kubernetes-sigs/external-dns/releases/tag/v0.7.3) for details on all changes.
