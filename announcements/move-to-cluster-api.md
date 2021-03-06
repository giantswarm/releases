### What is Cluster API?

Cluster API is a Kubernetes sub-project focused on providing declarative APIs and tooling to simplify provisioning, upgrading, and operating multiple Kubernetes clusters.

Started by the Kubernetes Special Interest Group (SIG) Cluster Lifecycle, the Cluster API project uses Kubernetes-style APIs and patterns to automate cluster lifecycle management for platform operators. The supporting infrastructure, like virtual machines, networks, load balancers, and VPCs, as well as the Kubernetes cluster configuration are all defined in the same way that application developers operate deploying and managing their workloads. This enables consistent and repeatable cluster deployments across a wide variety of infrastructure environments. The project was started over 2 years ago, and is growing and improving rapidly. You can find more details [here](https://cluster-api.sigs.k8s.io/).

### Why is Giant Swarm moving to Cluster API?

In 2017, due to a lack of better alternatives, we decided to build our product based on our own implementation of automation, tools, and operators. This allowed us to provide a stable and easy-to-use product that you are currently using to manage your Kubernetes clusters. Although this solution was ideal for a long time it has its downsides:
- You have limited control over configuration of individual parts of the infrastructure. The product is not flexible enough and sometimes doesn't support all of your use cases.
- Adding new functionality is time-consuming and as hyperscalers (Amazon, Microsoft) expand their offering it becomes increasingly difficult to keep up and make it all available to you in a timely manner.
- Your cluster management is limited to what our tooling enables and there are no available alternatives.
- Although we provide pure "vanilla" Kubernetes clusters, it is currently not possible to move the workload clusters and have them managed by another product without migration. 

We have been closely watching the development of the upstream Cluster API project and since it achieved a satisfying level of maturity, we decided to adopt it in order to address all of the downsides mentioned above and bring more value to you:
- Significantly more flexibility and configurability over your infrastructure. Things like deploying into existing networks or customizing api-server configuration directly will now be possible.
- A future-proof API that protects you from vendor lock-ins, developed by a dynamically growing community which will increase the speed of adding new functionality over time.
- The possibility to use all tooling compatible with the Cluster API -- not only that created by us.
- Ability to create and manage workload clusters with managed control planes like EKS or AKS.
- Adding new providers will become easier and we will be in a better position to expand our offering beyond AWS, Azure, and current On-Premises solutions.
- The possibility to have workload clusters from different providers managed by one management cluster to enable multi-cloud use cases.

### What will the move look like?

We are investing an enormous amount of effort to make the process safe and convenient for you. From your perspective, it will look like a regular update to a new major version -- no workload migration will be required. On our side, a series of well-tested, automated actions will be performed to transform your current clusters into Cluster API ones. We want to provide a transition path for the clusters in all versions and transform them into Cluster API clusters with little to no downtime in the Kubernetes control plane and no downtime in your workloads. This will be an in-place upgrade to preserve your current IP addresses.


In the first phase, starting the transition will be enabled for Giant Swarm personnel only, but eventually, you will be able to perform it on your workload clusters autonomously.

### What are the risks?

The infrastructure behind Giant Swarm clusters differs in many ways from the one backing a Cluster API cluster. Our engineers worked hard to find an automated, reliable, and reproducible transition path between the two, but there is always the risk that something goes wrong. We will ensure a well-defined, thoroughly tested transition plan to mitigate the risks and have a recovery plan in place if something unexpected happens. The transition will also be done in phases, leaving your mission-critical cluster(s) to the end to adapt the process to all possible custom configurations and make it bullet-proof.


### What will change?

There will be very little to no differences between the Giant Swarm and the Cluster API clusters from the functional point of view.

The most noticeable change will be the deprecation of the 'gsctl' tool that will no longer work for creating or managing Cluster API clusters. You will be able to continue using our Web UI and 'kubectl gs', which support login directly through your company's Single Sign-On with all of its security features provided including multi-factor authentication.

The concept of organizations as a means to isolate different teams, business units, or projects within one installation will remain. Organizations will become more flexible and be mapped to native Kubernetes namespaces and isolation concepts, allowing you to assign users or groups from your identity provider with granular permissions known from [Kubernetes role-based access control (RBAC)](https://kubernetes.io/docs/reference/access-authn-authz/rbac/).


### What is the ETA?

We have made a big step forward by having 6 of 8 engineering teams within Giant Swarm form a spearhead towards CAPI for the whole month of March 2021. We made good progress, but there's still a lot to do to bring you a production-ready setup on par with the features you already enjoy today. Our next goal is to provide you with an early test version in the coming months and then work in parallel on getting that version ready for production and enable you with a non-breaking upgrade path to it. We will keep you updated as things progress. In the meantime feel free to reach out to us with questions, concerns, or general comments.
