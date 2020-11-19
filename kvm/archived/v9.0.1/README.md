## :zap: Tenant Cluster Release 9.0.1 for KVM :zap:

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

This release includes Kubernetes v1.15.11 as well as some reliability and user experience improvements.
We highly recommend you to upgrade to this release if you want to continue running on Kubernetes 1.15 for now.
This is also the first release which is internally represented by our [Release CRD](https://docs.giantswarm.io/reference/cp-k8s-api/releases.release.giantswarm.io/). This is done in preparation of opening up the control plane to you.

In addition, this release replaces CoreOS with Flatcar Container Linux.
CoreOS has gone [end-of-life](https://coreos.com/os/eol/) and is being rapidly phased out.
Flatcar is a compatible fork of CoreOS which receives ongoing support.
To continue receiving security updates and to minimize the effort needed to migrate in the future, we recommend upgrading to this release.

This release includes multiple improvements to the NGINX Ingress Controller app:
1. It upgrades to upstream ingress-nginx `v0.30.0`.
2. Optional metrics service for prometheus-operator support was added. This allows prometheus-operator to scrape metrics to monitor the app.
3. NGINX IC SSL support for old browsers and clients (e.g. Safari 9) is dropped. This is due to default SSL ciphers no longer including AES-CBC based ciphers since they are considered weak. At your own risk, weak ciphers can still be enabled on demand independently for each cluster.

Below, you can find more details on components that were changed with this release.

### Kubernetes [v1.15.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.15.md#changelog-since-v11510)
- Updated from v1.15.5.
- Includes fixes for CVE-2020-8551 and CVE-2020-8552.
- Fixed a regression where the kubelet would fail to update the ready status of pods.
- Various other improvements to the management of cloud and cluster resources.

### kvm-operator [v3.9.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.9.1)
- Use Flatcar linux instead of CoreOS.
- Streamlined image templating for core components for quicker and easier releases in the future.
- Support setting OIDC username and groups prefix.
- Enabled per-cluster configuration of kube-proxy's `conntrackMaxPerCore` parameter.

### cert-exporter (GS [v1.2.2](https://github.com/giantswarm/cert-exporter/releases/tag/v1.2.2))

- Removed CPU limits.
- Change daemonset to use release revision not time for Helm 3 support.

### chart-operator [v0.13.0](https://github.com/giantswarm/chart-operator/releases/tag/v0.13.0)

- Deploy as a unique app in app collection in control plane clusters.
- Always set chart CR annotations so update state calculation is accurate.
- Only update failed Helm releases if the chart values or version has changed.
- Fix update state calculation and status resource for long running deployments.
- Handle 503 responses when GitHub Pages is unavailable.
- Make HTTP client timeout configurable for pulling chart tarballs in AWS China.
- Switch from dep to go modules
- Removed usage of legacy chartconfig CRs in Tiller metrics.
- Added chartmigration resource for migrating from chartconfig to chart CRs.
- Updated release resource.
  - Do not wait when installing or updating long running Helm releases.
  - Use version field from chart CR to reduce number of HTTP requests to pull chart tarballs.
  - Wait for deleted Helm release before removing finalizer.
- Updated status resource.
  - Improve reason field in CR status when installing a chart fails.
- Removed legacy chartconfig controller.
- Allowed usage of custom kubernetes domain.
- Handled timeouts pulling chart tarballs.
- Removed CPU limits.
- Updated for Kubernetes 1.16 compatibility.
- Cancelled the resource reconciliation in case of tarball errors.

### cluster-operator [v0.23.8](https://github.com/giantswarm/cluster-operator/releases/tag/v0.23.8)

- Stop setting IC replicas count.
- Fix cluster deletion by gracefully handling Tenant Cluster API errors.
- Initial cluster profile detection support.

### coredns v1.6.5 ([Giant Swarm app v1.1.8](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v118-2020-03-20))

- Updated coredns to upstream version 1.6.5.
- Removed CPU limits.
- Updated for Kubernetes 1.16 compatibility.
- Make `autopath` plugin configurable, optional and disabled by default.
- Allow custom forward configuration destination and options.
- Add Pod Disruption Budget.
- Use cluster.kubernetes.clusterDomain instead of cluster.kubernetes.domain for custom DNS suffix.

### metrics-server v0.3.3 ([Giant Swarm app v1.0.0](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#v100-2020-01-03))

- Updated for Kubernetes 1.16 compatibility.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.9](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v169-2020-04-22))

- Restrict PodSecurityPolicy volumes to only those required (removes wildcard).
- Tune net.ipv4.ip_local_port_range to 1024 65535 as a safe sysctl.
- Tune net.core.somaxconn to 32768 via an initContainer with privilege escalation.
- Use 4 worker processes by default.
- Ignore NGINX IC Deployment replica count configuration when HorizontalPodAutoscaler is enabled.
- Drop unnecessary Helm release revision annotation from NGINX IC Deployment.
- Adjusted README for display in the web interface context.
- HorizontalPodAutoscaler was tuned to use targetMemoryUtilizationPercentage of 80
- Removed use of enable-dynamic-certificates CLI flag, it has been deprecated since ingress-nginx 0.26.0
- Changed default error-log-level from error to notice.
- Align graceful termination configuration with changes made in upstream ingress-nginx 0.26.0
- Make controller.minReadySeconds configurable.
- Enabled HorizontalPodAutoscaler by default for selected cluster profiles.
- Upgrade to nginx-ingress-controller 0.30.0. - for details see the [changelog](https://github.com/kubernetes/ingress-nginx/releases/tag/nginx-0.30.0).
- Configured app icon.
- Added optional metrics Service (disabled by default) for prometheus-operator support.
- Added support for overriding all nginx [configmap settings](https://github.com/kubernetes/ingress-nginx/blob/master/docs/user-guide/nginx-configuration/configmap.md#configuration-options).

### node-exporter v0.18.1 ([Giant Swarm app v1.2.0](https://github.com/giantswarm/node-exporter-app/blob/master/CHANGELOG.md#120-2020-01-08))

- Updated to upstream version 0.18.1.
- Changed priority class to `system-node-critical`.

### Flatcar Linux [2345.3.1](https://www.flatcar-linux.org/releases/#release-2345.3.1)

- Updated from CoreOS 2191.5.0 - [changelog](https://www.flatcar-linux.org/releases/#release-2345.3.1)
