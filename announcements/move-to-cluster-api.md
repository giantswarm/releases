### What is Cluster API?

Cluster API is a Kubernetes sub-project focused on providing declarative APIs and tooling to simplify provisioning, upgrading, and operating multiple Kubernetes clusters.

Started by the Kubernetes Special Interest Group (SIG) Cluster Lifecycle, the Cluster API project uses Kubernetes-style APIs and patterns to automate cluster lifecycle management for platform operators. The supporting infrastructure, like virtual machines, networks, load balancers, and VPCs, as well as the Kubernetes cluster configuration are all defined in the same way that application developers operate deploying and managing their workloads. This enables consistent and repeatable cluster deployments across a wide variety of infrastructure environments.

### Why is Giant Swarm moving to Cluster API?

In 2017, due to lack of better alternatives, Giant Swarm decided to build their product based on their own implementation of automation, tools, and operators. This allowed Giant Swarm to provide a stable and easy-to-use product which you are currently using to manager your Kubernetes clusters. Although this solution was ideal for a long time it has it's downsides:
- You have limited control over configuration of individual parts of the infrastructure. The Product is not flexible enough and sometimes doesn't support all of your use cases.
- Adding new functionalities is time-consuming and as Hyperscalers (Amazon, Microsoft) expand their offering it becomes increasingly difficult to keep up and make it all available to you in a timely manner.
- Your cluster management is limited to what Giant Swarm tooling enables and there are no available alternatives.
- Although Giant Swarm provides pure Vanilla Kubernetes clusters, it is currently not possible to move the workload clusters and have them managed by another product without migration. 

Giant Swarm has been watching closely the development of the upstream Cluster API project and since it achieved a satisfying level of maturity, decided to adopt it in order to address all of the downsides mentioned above and bring more value to you:
- A lot more flexibility and configurability over your infastructure. Things like deploying into existing networks or ... will now be possible.
- Future proof API that protects you from vendor lock-ins, developed by dynamicaly growing community which will increase the speed of adding new functionalities over time.
- Possibility to use all the tooling compatible with the Cluster API not only one created by Giant Swarm.
- Workload Cluster could be managed by management clusters from other providers if they are compatible with Cluster API like EKS or AKS.
- Adding new providers will become easier and Giant Swarm will be in a better position to expand their offer beyond AWS, Azure, and On-Premises.
- It will be possible to have workload clusters from different providers managed by one management cluster to enable Multi-cloud use cases.



