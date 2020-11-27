# :zap: This Week in Provider-Independent Announcements (November 27, 2020):zap:

For a comprehensive and always-updated view of all new features and changes for apps, tenant clusters, UI, and documentation, please check out [Changes and Releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed Apps

1. [Grafana v0.1.0-alpha1](https://github.com/giantswarm/grafana-app) is now available in the Giant Swarm Catalog. It contains the latest upstream Grafana v7.3.1.
2. [Cert Manager v2.3.3](https://docs.giantswarm.io/changes/managed-apps/cert-manager-app/v2.3.3/) schedules hook Jobs on master nodes. You can now create AWS clusters with no Node Pools and Cert Manager will start up correctly.

## Web UI

[Happa v1.1.1](https://github.com/giantswarm/happa/releases/tag/v1.1.1) simplified deleting an organization and [Happa v1.1.0](https://github.com/giantswarm/happa/releases/tag/v1.1.0) added support for upgrading to beta versions.

## Documentation

- We extended our guide on [automatic termination of unhealthy nodes](https://docs.giantswarm.io/basics/automatic-termination-of-bad-nodes/) for Azure, where this function is available as tenant cluster release v13.1.0.
- The article on the [operational layers](https://docs.giantswarm.io/basics/giant-swarm-operational-layers/) got an overhaul.
- We added explanation for the use of deprecated tenant cluster releases where it matters.
- The guide on [adding labels to tenant clusters](https://docs.giantswarm.io/guides/tenant-cluster-labelling/) has been updated to inform you that this is possible on Azure.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
