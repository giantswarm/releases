# :zap: This Week in Provider-Independent Announcements (December 04, 2020):zap:

For a comprehensive and always-updated view of all new features and changes for apps, tenant clusters, UI, and documentation, please check out [Changes and Releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed Apps

- [Kong v1.1.0](https://docs.giantswarm.io/changes/managed-apps/kong-app/v1.1.0/) upgrades to Kong app v2.2.
- Vertical Pod Autoscaler is now available as a Playground Catalog app. It automatically adjusts the amount of CPU and memory requested by pods running in the Kubernetes Cluster. 

## Web UI

- Improved the support of Azure regions without any availability zones.
- Several fixes, including the one that brings back the form for setting provider credentials in an organization.

## gsctl

v0.27.0 adds support for scaling node pools to zero, and for Azure regions without availability zones.

## Documentation

- Updated the guide on [ceating a Cluster via the Control Plane Kubernetes API on Azure](https://docs.giantswarm.io/guides/creating-clusters-via-crs-on-azure/) to work well with the tenant cluster release v13.0.x for Azure.
- Added information regarding [which apps are preinstalled vs. optional](https://docs.giantswarm.io/reference/tenant-cluster-release-versions/#apps) depending on providers and versions.
- Several improvements to the search function.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
