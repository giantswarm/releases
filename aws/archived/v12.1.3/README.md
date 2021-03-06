# :zap: Tenant Cluster Release v12.1.3 for AWS :zap:

This release provides a new cluster-operator which fixes the issue that prevented release upgrade when reference id does not align with the G8sControlPlane id or MachineDeployment id and cluster status condition not being changed during cluster upgrade. 

**If you are upgrading from 12.1.0 or older platform release, this release uses Cert Manager 0.16.1 and will no longer reconcile existing resources** due to changes in its API. Manual intervention is required to update affected resources. While the negative impact to your workloads is low-to-none, to minimize disruption, **we recommend discussing this upgrade with your Solution Engineer**.

**Note for Solution Engineers:**

Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.

**Note for future 12.x.x releases:**

Please persist this note and the above, until all customers are on AWS v12.1.x and above.

## Change details

### cluster-operator [3.1.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.1.1)

- Fix cluster status is not updated during cluster upgrade.
- Add NetworkPolicy to restrict communication for cluster-operator. 
- Fix condition where reference id does not match with G8sControlplane or MachineDeployment.
