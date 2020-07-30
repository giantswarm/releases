## :zap:  Giant Swarm Release 12.2.0 for KVM :zap:

**If you are upgrading from 12.1.0, upgrading to this release will not roll your nodes.**

As of this release **NGINX Ingress Controller App** is **optional and not pre-installed** component on KVM. This allows NGINX App installations to be managed independently from the base platform lifecycle. It is both benefit but also new responsibility to keep NGINX App installations up-to-date separately from rest of the cluster. Making NGINX optional enables use of other ingress controller alternatives without wasting resources where NGINX is not the preferred option. Upgrading existing tenant clusters with pre-installed NGINX will leave NGINX unchanged. Existing NGINX App custom resources will still have `giantswarm.io/managed-by: cluster-operator` label, but it will be ignored. The label will be cleaned up at a later point after all tenant clusters have been upgraded and KVM platform releases older than v12.2.0 archived.

**Note for cluster upgrades:**

Please ensure cluster is on KVM 12.1.x platform release first before upgrading the cluster to 12.2.0+

Below, you can find more details on components that were changed with this release.

### cluster-operator [0.23.14](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.14)

- Support for making NGINX IC App optional and not pre-installed.
