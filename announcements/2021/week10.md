# :zap: This week in provider-independent announcements (March 5, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[EFK Stack v0.5.0](https://docs.giantswarm.io/changes/managed-apps/efk-stack-app/v0.5.0/) is updated to Elasticsearch and Kibana v7.10.2.

[Aqua v5.3.2](https://docs.giantswarm.io/changes/managed-apps/aqua-app/v5.3.2/) and [v5.3.1](https://docs.giantswarm.io/changes/managed-apps/aqua-app/v5.3.1/) prior to it fix issues with data types, port configurations, secrets for DB credentials, and more. It also adds security context to relevant resources and guidance on advanced configuration to the README.

## Apps supported with best-effort

[Azure Ad Pod Identity v0.5.0](https://docs.giantswarm.io/changes/playground-apps/azure-ad-pod-identity-app/v0.5.0/) and prior releases from v0.3.0 to v0.4.2 upgrade the chart to v3.0.3. They also fix PSP capabilities to allow it to work on management clusters, as well as allow more configuration options to work by default with CAPZ (Cluster API Provider Azure)

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
