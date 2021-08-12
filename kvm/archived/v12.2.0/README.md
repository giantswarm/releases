## :zap:  Tenant Cluster Release 12.2.0 for KVM :zap:

**If you are upgrading from 12.1.0, upgrading to this release will not roll your nodes.**

As of this release, **NGINX Ingress Controller App** is now an **optional (and not pre-installed)** component on KVM.

This enables use of alternative ingress controllers without wasting resources where NGINX is not the preferred option.

Now NGINX App installations can be managed and updated independent of the cluster, which is both a benefit and a new responsibility 😅

Upgrading tenant clusters with pre-installed NGINX will leave NGINX unchanged. Existing NGINX App custom resources will still have `giantswarm.io/managed-by: cluster-operator` label, but it will be ignored. The label will be cleaned up later after all tenant clusters have been upgraded and KVM platform releases older than v12.2.0 archived.

**Note for cluster upgrades:**

Please ensure cluster is on KVM 12.1.x platform release first before upgrading the cluster to 12.2.0+

Below, you can find more details on components that were changed with this release.

### cluster-operator [0.23.14](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.14)

- Support for making NGINX IC App optional and not pre-installed.
