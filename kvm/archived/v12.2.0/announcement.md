## :zap:  Tenant Cluster Release 12.2.0 for KVM :zap:

As of this release, **NGINX Ingress Controller App** is now an **optional (and not pre-installed)** component on KVM.

This enables use of alternative ingress controllers without wasting resources where NGINX is not the preferred option.

Now NGINX App installations can be managed and updated independent of the cluster, which is both a benefit and a new responsibility 😅

**Before upgrading, please ensure cluster is on KVM v12.1.x first.**

You can find more details on components that were changed in the [release notes](https://github.com/giantswarm/releases/tree/master/kvm/v12.2.0).

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.