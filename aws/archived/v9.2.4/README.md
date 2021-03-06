# :zap: Tenant Cluster Release 9.2.4 for AWS :zap:

**If you are upgrading from 9.2.3, upgrading to this release will not roll your nodes. It will only update the apps.**

This release improves the reliability of NGINX Ingress Controller. Most importantly, kernel and app settings have been tuned to increase out-of-the-box performance. The app's Helm chart was also adjusted to improve its availability when rolling out configuration changes.

Below, you can find more details on components that were changed with this release.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.9](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v169-2020-04-22))

- Restricted PodSecurityPolicy volumes to only those required (removes wildcard).
- Tuned `net.ipv4.ip_local_port_range` to `1024 65535` as a safe sysctl.
- Tuned `net.core.somaxconn` to `32768` via an initContainer with privilege escalation.
- Default `worker-processes` to `4`.
- Default `max-worker-connections` to upstream default (currently `16384`).
- Ignore NGINX IC Deployment replica count configuration when HorizontalPodAutoscaler is enabled.
- Dropped unnecessary Helm release revision annotation from NGINX IC Deployment.
- Adjusted README for display in the web interface context.
