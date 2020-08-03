## :zap:  Giant Swarm Release 12.1.0 for KVM :zap:

This release includes two **significant improvements to NGINX Ingress Controller**. It also includes a **fix for Quay being a single point of failure** by using Docker mirroring feature.

**The two NGINX Ingress Controller improvements:**

- Multiple NGINX Ingress Controllers per tenant cluster are now supported
- Management of NGINX IC NodePort Service is moved from kvm-operator to NGINX IC App itself

Along with kvm-operator, cluster-operator, and NGINX IC, release includes other upstream component upgrades.

You can find more details on components that were changed in [the release notes](https://github.com/giantswarm/releases/tree/master/kvm/v12.1.0).

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
