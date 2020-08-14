# :zap: Giant Swarm Release v12.1.1 for AWS :zap:

**If you are upgrading from 12.0.0, upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades kiam App from v1.3.1 to v1.4.0, to align with Cert Manager v0.15.2 API changes.

It is part of a larger effort to maximize reliability of upgrades, by keeping 3rd party components within 30 days of upstream releases.

**This upgrade will stop Cert Manager from reconciling your existing resources** due to changes in its API - manual intervention is required to update affected resources. While the **negative impact to your workloads is low-to-none**, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer first**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

### aws-operator[8.8.0](https://github.com/giantswarm/aws-operator/releases/tag/v8.8.0)

#### Changed
- Add release version tag for ec2 instances.
- Update Cloudformation Stack when components version differ.
- Fix IAM policy on Tenant Clusters to manages IAM Role tags.
- Fixed passing custom pod CIDR to k8scloudconfig.

### kiam [1.4.0](https://github.com/giantswarm/kiam-app/releases/tag/v1.4.0)

#### Changed

- Updated cert-manager API groups. ([#36](https://github.com/giantswarm/kiam-app/pull/36))
