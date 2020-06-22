## :zap:  Giant Swarm Release 11.4.0 for Azure :zap:

**If you are upgrading from 11.3.3, upgrading to this release will not roll your nodes. It will only update the apps.**

In this release NGINX Ingress Controller LoadBalancer Service external traffic policy has been made configurable, and the policy default has been changed from `Cluster` to `Local`. These changes:

- Enable source IP preservation, needed for IP based security access control.
- Remove unnecessary ingress traffic intra-cluster network hops, increasing concurrent connection capacity and reducing latency.
- Allow configuration to be reverted, where these defaults and associated tradeoffs are found not to be appropriate.

Management of NGINX IC LoadBalancer Service is moved from azure-operator to NGINX IC App itself to:

- Enable external traffic policy configurability, in a way consistent with other NGINX IC configuration options.
- Lay the foundation for making NGINX IC App optional and not pre-installed in a future Azure platform release.

When upgrading existing clusters from older Azure platform releases, migration of LoadBalancer Service is automated but not fully automatic. Therefore we ask of you to delegate cluster upgrades to your SE.

**Note for SEs:** After cluster upgrade to 11.4.0, both old `ingress-loadbalancer` LoadBalancer Service managed by azure-operator and new one `nginx-ingress-controller` managed by NGINX IC App remain on the cluster. To switch the ingress traffic to the new LoadBalancer and remove old NGINX LoadBalancer Service without downtime please:

- Together with the customer have any firewall in front of NGINX reconfigured to allow both old and new LoadBalancer Service IPs.
- Next use [the migration script](https://github.com/giantswarm/azure-operator/blob/master/scripts/migrate-nginx-ingress-controller.sh) to switch DNS records to the new load balancer IP. Now delete the old NGINX IC LoadBalancer Service.
- Finally, ensure that the tenant cluster's
  - Azure `kubernetes` load balancer backend pool includes the worker VM scale set, and that
  - All the worker VM scale set instances have latest model applied.

Below, you can find more details on components that were changed with this release.

### azure-operator v4.1.0

- Move NGINX LoadBalancer Service management from azure-operator to nginx-ingress-controller app.

### nginx-ingress-controller v1.7.0

- Change NGINX Service type from NodePort to LoadBalancer for Azure.
- Make NGINX LoadBalancer Service external traffic policy configurable.
- Use Local as NGINX LoadBalancer Service default external traffic policy.

### cluster-operator 0.23.10

- Fix a bug in user values migration logic for apps.
- Enable NGINX LoadBalancer Service on Azure.
