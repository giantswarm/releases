# :zap: This week in provider-independent announcements (February 26, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

[Prometheus Operator v0.7.0](https://docs.giantswarm.io/changes/managed-apps/prometheus-operator-app/v0.7.0/) upgrades to upstream chart v13.10.0, which contains Prometheus v2.24.0.  
:warning: This is a breaking change for the app and requires Helm 3 (AWS: >= 12.3.0, Azure: >= 12.1.0, KVM: >= 12.3.0) to install.

[Fluent Logshipping v0.6.0](https://docs.giantswarm.io/changes/playground-apps/fluent-logshipping-app/v0.6.0/) drops fluentd to reduce memory usage as fluent-bit now supports AWS outputs. It adds toleration to run on any node with taints.

## kubectl gs

- The `template nodepool` command now supports scaling to zero, as it is possible on Azure.

## Documentation

- We added information on how to use [cloud provider resource tagging](https://docs.giantswarm.io/ui-api/management-api/creating-workload-clusters/aws/#cluster-provider-resource-tagging) on AWS with the Management API.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
