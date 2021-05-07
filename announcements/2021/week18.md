# :zap: This week in provider-independent announcements (May 7, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[Cert Manager v2.7.0](https://docs.giantswarm.io/changes/managed-apps/cert-manager-app/v2.7.0/) upgrades to upstream v1.3.1, which fixes an issue preventing app upgrades.

[Kong v2.0.0](https://docs.giantswarm.io/changes/managed-apps/kong-app/v2.0.0/) upgrades to upstream stable v2.0.0. Make sure to review the [upstream changelog](https://github.com/Kong/charts/blob/main/charts/kong/CHANGELOG.md#200) and the [upgrade documentation](https://github.com/giantswarm/kong-app/blob/master/helm/kong-app/UPGRADE.md), as support for Helm 2 as well as other 1.x features are dropped. Customers runnning 1.2.0 have no disruptions.

## User interfaces

[kubectl gs](https://docs.giantswarm.io/ui-api/kubectl-gs/), our plug-in for the Giant Swarm Management API, is now available as binary for the ARM achitecture (e. g. Apple M1), both for Linux and macOS.

## Documentation

We now provide a curated list of [cost optimization](https://docs.giantswarm.io/advanced/cost-optimization/) tools for Kubernetes.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
