# :zap: This week in provider-independent announcements (April 30, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[EFK v0.5.1](https://docs.giantswarm.io/changes/managed-apps/efk-stack-app/v0.5.1/) upgrades to Open Distro for Elasticsearch v1.13.2.

[Prometheus Operator v0.8.0](https://github.com/giantswarm/prometheus-operator-app/blob/master/CHANGELOG.md#080---2021-04-26) upgrades to kube-prometheus-stack v15.2.0, which includes Prometheus v2.26.0.

[NGINX IC v1.16.1](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#1161---2021-04-20) is now more flexible, allowing you to change any annotation you want, for example, the AWS LoadBalancer Type.

[Cert Manager v2.6.0](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#260---2021-04-30) adds support for DNS01 challenge, allowing gigantic.io certificates to work behind a firewall. It simplifies creation of installations on-premises. Prior, [Cert Manager v2.5.0](https://github.com/giantswarm/cert-manager-app/blob/master/CHANGELOG.md#250---2021-04-19) upgrades to upstream v1.2.0.

[Aqua v5.3.4](https://github.com/giantswarm/aqua-app/blob/master/CHANGELOG.md#534---2021-04-28) upgrades to container image v5.3.21112.

## User interfaces

## Docs

The [search engine in docs](https://docs.giantswarm.io/) now searches the blog, making it quicker for you to find what you are looking for. Please try it and let us know what you think!

We updated these docs to reflect recent improvements: [Advanced ingress configuration](https://docs.giantswarm.io/advanced/ingress/configuration/), [The Giant Swarm App Platform](https://docs.giantswarm.io/app-platform/overview/), and [Prepare an Azure subscription to run Giant Swarm workload clusters](https://docs.giantswarm.io/getting-started/cloud-provider-accounts/azure/).

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
