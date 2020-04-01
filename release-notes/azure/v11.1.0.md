## :zap: Giant Swarm Release 11.1.0 for Azure :zap:

This is the first Giant Swarm release which includes Multiple Availability Zones for Azure. This release also introduces faster and less disruptive upgrade procedure that respects PDBs, as well as improved worker labeling.


Due to Azure limitations, upgraded clusters do not support multiple Availability Zones.  In order for a cluster to be able to scale across multiple Availability Zones, it must be created using the v11.1.0 release.


**You have 2 ways to proceed with this release:**
1. Create a new 11.1.0 cluster with Availability Zones and migrate your workloads to use the Multi AZ features.
2. Upgrade to 11.1.0 to keep up to date with kubernetes and other features like new graceful upgrades without the Multi AZ.

---

### azure-operator v2.9.0

#### Minor changes

- Cluster scaling is less disruptive.

#### Improved worker nodes' upgrade process

With this release the worker nodes' upgrade process got notable changes and it is now faster and less disruptive. It works by mirroring the existing nodes with new ones and then draining the old nodes. Once all old nodes are drained, they are terminated. This way kubernetes controller manager and scheduler can move the running workload to new nodes while respecting configured scheduling directives for undisrupted service.

In order to ensure continuous service during upgrade, the workload must be configured with multiple replicas and appropriate [pod disruption budget](https://kubernetes.io/docs/tasks/run-application/configure-pdb/).
