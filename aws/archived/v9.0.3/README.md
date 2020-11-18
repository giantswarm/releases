# :zap: Tenant Cluster Release 9.0.3 for AWS :zap:

**If you are upgrading from 9.0.2, upgrading to this release will not roll your nodes. It will only update the apps.**

---

**Note** If you are upgrading from 8.5.0 or 9.0.0:

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

This release improves the reliability of NGINX Ingress Controller. Most importantly, kernel and app settings have been tuned to increase out-of-the-box performance. The app's Helm chart was also adjusted to improve its availability when rolling out configuration changes.

This version also includes improvements to other components (chart-operator) as detailed in the changelog.

Below, you can find more details on components that were changed with this release.

### chart-operator [v0.13.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#v0130-2020-04-21)

- Fix update state calculation and status resource for long running deployments.
- Handle 503 responses when GitHub Pages is unavailable.
- Make HTTP client timeout configurable for pulling chart tarballs in AWS China.
- Switch from dep to go modules.
- Fix problem pushing chart to default app catalog.
- Always set chart CR annotations so update state calculation is accurate.
- Only update failed Helm releases if the chart values or version has changed.
- Deploy as a unique app in app collection in control plane clusters.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.9](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v169-2020-04-22))

- Restricted PodSecurityPolicy volumes to only those required (removes wildcard).
- Tuned `net.ipv4.ip_local_port_range` to `1024 65535` as a safe sysctl.
- Tuned `net.core.somaxconn` to `32768` via an initContainer with privilege escalation.
- Default `worker-processes` to `4`.
- Default `max-worker-connections` to upstream default (currently `16384`).
- Ignore NGINX IC Deployment replica count configuration when HorizontalPodAutoscaler is enabled.
- Dropped unnecessary Helm release revision annotation from NGINX IC Deployment.
- Adjusted README for display in the web interface context.
