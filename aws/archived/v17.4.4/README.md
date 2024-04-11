# :zap: Giant Swarm Release v17.4.4 for AWS :zap:

This is a patch release to add missing Elastic File System (EFS) IAM permissions. It also allows EFS CSI Driver to get the Instance Identity from the Instance Metadata Service (IMDS).

## Change details


### aws-operator [11.16.1](https://github.com/giantswarm/aws-operator/releases/tag/v11.16.1)

#### Fixed
- Added EFS policy to the ec2 instance role to allow to use the EFS driver out of the box



### kiam [2.5.1](https://github.com/giantswarm/kiam-app/releases/tag/v2.5.1)

#### Fixed
- Allow `whiteListRouteRegexp` to default to `/latest/*`.



### cert-manager [2.15.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.15.3)

#### Added
- Webhook: Add `PodDisruptionBudget` and pod anti-affinity.
- Startup  API check: Add `NetworkPolicy`.
#### Changed
- Webhook: Increase replica count to 2.



### external-dns [2.15.2](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.2)

#### Changed
- Update init container image to v3.16.2([#182](https://github.com/giantswarm/external-dns-app/pull/182))

### etcd-kubernetes-resources-count-exporter [1.8.0](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v1.8.0)

Bump to latest version to address security and performance issues. 
