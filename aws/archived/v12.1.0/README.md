# :zap: Tenant Cluster Release v12.1.0 for AWS :zap:

**If you are upgrading from 12.0.0, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades Cert Manager from upstream v0.9.0 to v0.15.2.

It is part of a larger effort to maximize reliability of upgrades, by keeping 3rd party components within 30 days of upstream releases.

**This upgrade will stop Cert Manager from reconciling your existing resources** due to changes in its API - manual intervention is required to update affected resources. While the **negative impact to your workloads is low-to-none**, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer first**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

### cert-manager [2.0.2](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.0.2)

#### Changed

- Upgrade cert-manager from 0.9.0 to 0.15.2. ([#31](https://github.com/giantswarm/cert-manager-app/pull/31))
