# :zap: Giant Swarm Release v12.5.0 for AWS :zap:

Offers the possibility to add additional Network Pools to the Control Plane and flexibly choose the IP range for new Tenant Clusters from these pool.

**Note for Solution Engineers:**

- Helm3: 
  - Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.3.x** and above.
- cert-manager-app: 
  - Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.1.x** and above.

## Change details


### aws-operator [9.1.2](https://github.com/giantswarm/aws-operator/releases/tag/v9.1.2)

Not found


### cluster-operator [3.3.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.3.1)

Not found


