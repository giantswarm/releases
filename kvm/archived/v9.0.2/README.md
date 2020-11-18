## :zap: Tenant Cluster Release 9.0.2 for KVM :zap:

**Note** If you are upgrading from 8.4.0 or 9.0.0:

Configuration that used to be on the Tenant Cluster has been moved to the Control Plane.

Services like the `nginx-ingress-controller`, `coredns`, and `cluster-autoscaler`
are no longer configurable through ConfigMaps on the tenant cluster.

Existing ConfigMaps will be automatically migrated when you upgrade, however if
you do not have access to the Control Plane API you will not be able to
manually configure these services.

e.g: If you had a cluster with id `e05c8`, then you'll find your
`nginx-ingress-controller-user-values` configmap now in a namespace called `e05c8`
on the Control Plane.

Please contact your Solution Engineer for more information.

---

This release fixes a problem that prevented clusters with OIDC user and group prefix settings to work as expected in `9.0.1`.

### kvm-operator [v3.9.2](https://github.com/giantswarm/kvm-operator/releases/tag/v3.9.2)
- Use Release.Revision in Helm chart for Helm 3 support.
- Fix OIDC settings which are passed to masters.
