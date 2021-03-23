### What is Cluster API?

Cluster API is a Kubernetes sub-project focused on providing declarative APIs and tooling to simplify provisioning, upgrading, and operating multiple Kubernetes clusters.

Started by the Kubernetes Special Interest Group (SIG) Cluster Lifecycle, the Cluster API project uses Kubernetes-style APIs and patterns to automate cluster lifecycle management for platform operators. The supporting infrastructure, like virtual machines, networks, load balancers, and VPCs, as well as the Kubernetes cluster configuration are all defined in the same way that application developers operate deploying and managing their workloads. This enables consistent and repeatable cluster deployments across a wide variety of infrastructure environments.

### Why is Giant Swarm moving to Cluster API?

In 2017, due to lack of better alternatives, Giant Swarm decided to build their product based on their own implementation of automation, tools, and operators. This allowed Giant Swarm to provide a stable and easy-to-use product which you are currently using to manager your Kubernetes clusters. Although this solution was ideal for a long time it has it's downsides:
- You have limited control over configuration of individual parts of the infrastructure. The product is not flexible enough and sometimes doesn't support all of your use cases.
- Adding new functionalities is time-consuming and as hyperscalers (Amazon, Microsoft) expand their offering it becomes increasingly difficult to keep up and make it all available to you in a timely manner.
- Your cluster management is limited to what Giant Swarm tooling enables and there are no available alternatives.
- Although Giant Swarm provides pure "vanilla" Kubernetes clusters, it is currently not possible to move the workload clusters and have them managed by another product without migration. 

Giant Swarm has been closely watching the development of the upstream Cluster API project and since it achieved a satisfying level of maturity, decided to adopt it in order to address all of the downsides mentioned above and bring more value to you:
- Significantly more flexibility and configurability over your infrastructure. Things like deploying into existing networks or ... will now be possible.
- A future-proof API that protects you from vendor lock-ins, developed by a dynamically growing community which will increase the speed of adding new functionalities over time.
- The possibility to use all tooling compatible with the Cluster API -- not only that created by Giant Swarm.
- Workload clusters can be managed by management clusters from other providers if they are compatible with Cluster API, like EKS or AKS.
- Adding new providers will become easier and Giant Swarm will be in a better position to expand their offering beyond AWS, Azure, and current On-Premises solutions.
- The possibility to have workload clusters from different providers managed by one management cluster to enable multi-cloud use cases.

### What will the move look like?

Giant Swarm is investing an enormous amount of effort to make the process safe and convenient for you. From your perspective, it will look like a regular update to a new major version -- no migration will be required. On the Giant Swarm side, a combination of automated and manual steps will be performed to replace Giant Swarm's components with upstream Cluster API. During a similar process for undertaken for Azure customers replacing CoreOS with Flatcar, the only inconvenience for the users was the need to schedule the upgrade ahead of time to make sure there's an engineer available to perform the manual parts of the transition.

These are the steps Giant Swarm plans to make:
1. ...


### What are the risks?

### What will change?
