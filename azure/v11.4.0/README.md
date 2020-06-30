## :zap:  Giant Swarm Release 11.4.0 for Azure :zap:

In this release NGINX Ingress Controller LoadBalancer Service external traffic policy has been made configurable, and the policy default has been changed from `Cluster` to `Local`. These changes:

- Enable source IP preservation, needed for IP based security access control.
- Remove unnecessary ingress traffic intra-cluster network hops, increasing concurrent connection capacity and reducing latency.
- Allow configuration to be reverted, where these defaults and associated tradeoffs are found to be inappropriate.

**Management of NGINX IC LoadBalancer Service is moved from azure-operator to NGINX IC App itself** to:

- Enable external traffic policy configurability, in a way consistent with other NGINX IC configuration options.
- Lay the foundation for making NGINX IC App optional and not pre-installed in a future Azure platform release.
Along with azure-operator and NGINX IC, minor improvements were also made to coreDNS, cluster-operator, Flatcar, and Kubernetes.
When upgrading clusters, migration of LoadBalancer Service is automated but not fully automatic. Therefore, please delegate cluster upgrades to your SE.

**Note for SEs:** After cluster upgrade to 11.4.0, both old `ingress-loadbalancer` LoadBalancer Service managed by azure-operator and new one `nginx-ingress-controller` managed by NGINX IC App remain on the cluster. To switch the ingress traffic to the new LoadBalancer and remove old NGINX LoadBalancer Service without downtime please:

- Together with the customer have any firewall in front of NGINX reconfigured to allow both old and new LoadBalancer Service IPs.
- Next use [the migration script](https://github.com/giantswarm/azure-operator/blob/master/scripts/migrate-nginx-ingress-controller.sh) to switch DNS records to the new load balancer IP. The script ensures IP is assigned to the new LB, and also that the cluster DNS records resolve to it instead of old IP.
- Now delete the old NGINX IC LoadBalancer Service (the one called `ingress-loadbalancer`).

Below, you can find more details on components that were changed with this release.

### azure-operator [v4.1.0](https://github.com/giantswarm/azure-operator/blob/v4.1.0/CHANGELOG.md#410---2020-06-24)

- Moved NGINX LoadBalancer Service management from azure-operator to nginx-ingress-controller app.
- The default egress strategy for worker nodes VMSS is now a VNET Gateway rather than the Load Balancer.

### CoreDNS v1.6.5 ([Giant Swarm app v1.1.10](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v1110-2020-06-29))

- Made forward options optional.
- Made resources (requests/limits) configurable.

### cluster-operator [v0.23.10](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.10)

- Fixed a bug in user values migration logic for apps.
- Enabled NGINX LoadBalancer Service on Azure.

### Flatcar v2512.2.1
- Updated from v2345.3.1 - 
changelog since [v2345.3.1](https://www.flatcar-linux.org/releases/#release-2512.2.0) and [v2512.2.0](https://www.flatcar-linux.org/releases/#release-2512.2.1)
- Includes a fix for [CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543) regarding Intel Microcode vulnerabilities.

### Kubernetes v1.16.12
- Updated from v1.16.8 - 
changelog since [v1.16.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v11611-1),
since [v1.16.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v11610),
since [v1.16.9](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v1169) and
since [v1.16.8](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v1168)
- Includes a fix for CVE-2019-11253 related to json/yaml decoding where large or malformed documents could consume excessive server resources. Request bodies for normal API requests (create/delete/update/patch operations of regular resources) are now limited to 3MB.

### nginx-ingress-controller v0.33.0 ([Giant Swarm app v1.7.0](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v170-2020-06-23))

- Changed NGINX Service type from NodePort to LoadBalancer for Azure.
- Made NGINX LoadBalancer Service external traffic policy configurable.
- Use Local as NGINX LoadBalancer Service default external traffic policy.
- Upgraded to ingress-nginx [0.33.0](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0330).
