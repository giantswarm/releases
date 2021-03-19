# :zap: This week in provider-independent announcements (March 19, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[Loki v0.1.2-beta](https://docs.giantswarm.io/changes/managed-apps/loki-app/v0.1.2-beta/) adds the option to install loki-multi-tenant-proxy to ease multi-tenant authentication and authorization. Prior, [Loki v0.1.1-alpha2](https://docs.giantswarm.io/changes/managed-apps/loki-app/v0.1.1-alpha2/) upgrades to Loki v2.2.0.

[Grafana v0.2.0](https://docs.giantswarm.io/changes/managed-apps/grafana-app/v0.2.0/) upgrades to upstream Grafana image v7.4.3.

[Fluent Logshipping v0.6.4](https://docs.giantswarm.io/changes/playground-apps/fluent-logshipping-app/v0.6.4/) breaks up the stream names for Cloudwatch Output, to help prevent from hitting API limits.

## Apps supported with best effort

[Goldilocks v0.1.0]() is now available. Some customers report significant cost savings from automatically setting requests and limits for clusters.

[Cloudflared v0.0.5]() is now available as an app supported with best effort. This app allows you to launch Cloudflare Argo Tunnels and then route to services inside your cluster. Best for customers whose ingress options are constained and / or clusters run on-premises.

[Strimzi Kafka Operator v0.2.2](https://docs.giantswarm.io/changes/playground-apps/strimzi-kafka-operator-app/v0.2.2/) fixes and removes the duplicated registry in image entry after templating.

## Documentation

- For consistent terminology and to avoid confusion, we now use the term `Giant Swarm REST API` where we simply talked about the API or the Giant Swarm API in the past.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
