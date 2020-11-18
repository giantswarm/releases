# :zap: Tenant Cluster Release v12.1.1 for AWS :zap:

This release provides new aws-operator version with reliability improvements and upgrades kiam-app from v1.3.1 to v1.4.0, to align with Cert Manager v0.16.1 API changes.

**This version of Cert Manager no longer reconciles your existing resources** due to changes in its API. Manual intervention is required to update affected resources. While the negative impact to your workloads is low-to-none, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

## Change details

### aws-operator [v8.8.0](https://github.com/giantswarm/aws-operator/blob/master/CHANGELOG.md#876---2020-08-11)

- Fixes a certain case where a release upgrade would have left cluster in an in-between state.
- aws-operator now emits events on the `AWSCluster`, `AWSControlPlane`, and `AWSMachineDeployment` resources to facilitate following the process of cluster creation, upgrades etc., and debugging issues.
- Fixed passing custom pod CIDR to k8scloudconfig.

### kiam [v3.6](https://github.com/uswitch/kiam/blob/master/CHANGELOG.md#v36) (Giant Swarm app [v1.4.0](https://github.com/giantswarm/kiam-app/releases/tag/v1.4.0))

- Updated cert-manager API groups. ([#36](https://github.com/giantswarm/kiam-app/pull/36))
