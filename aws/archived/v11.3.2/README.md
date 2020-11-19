## :zap: Tenant Cluster Release v11.3.2 for AWS :zap:

**If you are upgrading from 11.3.1, upgrading to this release will not roll your nodes. It will only update the apps.**

This release fixes an external-dns issue affecting clusters in AWS China. We added an option for the app to use CNAME records instead of ALIAS, in order for DNS resolution to work in China.

Note that the fix works for new cluster creation only. Upgrades require a manual workaround. We're working with the upstream external-dns community to fix it. In the meantime, we recommend delegating cluster upgrades to your SE.```

**Note for SEs:** When upgrading a cluster, existing ingress A+TXT record sets do not get replaced with CNAME+TXT record sets even when external-dns is configured with CNAMEs as preferred. After upgrading, delete the ingress A+TXT record sets. external-dns will then automatically create CNAME+TXT record sets.

## external-dns v0.5.18 ([Giant Swarm app v1.2.1](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#v121-2020-05-29))

- Prefer CNAMEs record sets for AWS SDK configuration with explicit credentials.
