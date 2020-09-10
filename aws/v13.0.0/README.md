# :zap: Giant Swarm Release v13.0.0 for AWS :zap:

This release upgrades external-dns app to v1.4.0 to improve observability.

**If you are upgrading from 12.1.3, upgrading to this release will not roll your nodes. It will only update the apps.**

**If you are upgrading from 12.1.0 or older platform release, this release uses Cert Manager 0.16.1 and will no longer reconcile existing resources** due to changes in its API. Manual intervention is required to update affected resources. While the negative impact to your workloads is low-to-none, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

## Change details

### external-dns v0.7.3 [Giant Swarm app 1.4.0](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#140---2020-08-21)

- Added monitoring headless Service.
- Added more Giant Swarm custom monitoring annotations.
- Explicitly exposed metrics container port.
- Use default external-dns metrics port 7979.
