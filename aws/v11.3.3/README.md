## :zap: Tenant Cluster Release v11.3.3 for AWS :zap:

**If you are upgrading from 11.3.2, upgrading to this release will not roll your nodes. It will only update the apps.**

This release fixes an issue where `cert-manager` could be killed for exceeding its memory limit. This stops SSL certificates from being
automatically renewed, which would cause expired certificates to show your site as not secure.

**Note for SEs:** This release contains an external-dns fix introduced in [11.3.2](https://github.com/giantswarm/releases/tree/master/aws/v11.3.2). It requires manual intervention for cluster upgrades in China to work. When upgrading a cluster, existing ingress A+TXT record sets do not get replaced with CNAME+TXT record sets even when external-dns is configured with CNAMEs as preferred. After upgrading, delete the ingress A+TXT record sets. external-dns will then automatically create CNAME+TXT record sets.

**Note for future 11.3.x releases:** Until [upstream external-dns issue](https://github.com/kubernetes-sigs/external-dns/issues/1617) is fixed, please include this note and the one above in all future 11.3.x releases.

## cert-manager-app v0.9.0 ([Giant Swarm app v1.0.8](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#v108-2020-04-30))

- Remove deployment memory limit.
- Allow resource requests and limits to be configured with `values.yaml`.
