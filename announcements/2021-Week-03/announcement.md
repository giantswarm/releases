# :zap: This Week in Provider-Independent Announcements (January 22, 2021):zap:

Please note that we are aligning some key terms in our product communication with the Kubernetes project's terminology:

- The term *management cluster* replaces the term *control plane*. The Kubernetes project uses the term control plane when dealing with the master nodes (as they were called formerly) of a cluster, so we are removing an ambiguity that confused people.
- The API formerly called *Control Plane Kubernetes API* is now called *Management API*.
- What we called *tenant cluster* is now called a *workload cluster*.

For a comprehensive and always-updated view of all new features and changes for apps, tenant clusters, UI, and documentation, please check out [Changes and Releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed Apps

[Loki v0.1.0-alpha](https://github.com/giantswarm/loki-app/blob/master/CHANGELOG.md#010-alpha---2021-01-21), our preferred app for log management is now available. Promtail, the logging client, is coming soon. Interested in Loki? We'd like to hear from you.

[Prometheus Operator v0.5.2](https://docs.giantswarm.io/changes/managed-apps/prometheus-operator-app/v0.5.2/) introduces a new AlertmanagerConfig CRD and upgrades to Prometheus v2.22.1.

[EFK Stack v0.4.1](https://docs.giantswarm.io/changes/managed-apps/efk-stack-app/v0.4.1/) came quickly after v0.4.0 to fix a small oversight. With v0.4.x Elasticsearch and Kibana are running in version v7.10.0.

[Cert Manager v2.4.1](https://docs.giantswarm.io/changes/managed-apps/cert-manager-app/v2.4.1/) fixes an issue causing Cert Manager to fail (makes backoffLimit for clusterissuer job configurable).

[Aqua v5.3.0](https://docs.giantswarm.io/changes/managed-apps/aqua-app/v5.3.0/) upgrades all charts to upstream 5.3 and introduces schema for validation of user values.

## Playground Apps

[Jaeger Operator v0.2.2](https://docs.giantswarm.io/changes/playground-apps/jaeger-operator-app/v0.2.2/) upgrades to the upstream chart v1.21.2 and jaeger-all-in-one to v1.21.0.

## Web UI

## kubectl gs

- Added support for getting nodepools and removed mention of the unexistent ‘create cluster’ command.

## Documentation



---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
