# :zap: Giant Swarm Release v11.5.0 for AWS :zap:

This release includes a number of fixes, improvements, and picks up various component and app upgrades.

**Security:** [CVE-2020-8557](https://github.com/kubernetes/kubernetes/issues/93032), [CVE-2020-8558](https://github.com/kubernetes/kubernetes/issues/92315), and [CVE-2020-8559](https://github.com/kubernetes/kubernetes/issues/92914) are mitigated through an upgrade to Kubernetes v1.16.13.

**Registry mirrors:** Dockerd is now configured to use several registries (_registry mirrors_), to avoid a single point of failure. Workloads managed by Giant Swarm are configured to make use of this. Any containers that specify a registry explicitly (like `image: reg.acme.com/namespace/repo:tag`) will not be affected. However, when no registry is named, the image will first be looked up in the DockerHub registry (docker.io) and then in the configured mirror(s).
Note: Giant Swarm installations on AWS China (cn-north-1) do not yet provide any registry mirrors as of this release.

## Change details

### cluster-operator [v2.3.1](https://github.com/giantswarm/cluster-operator/blob/master/CHANGELOG.md#231---2020-07-14)

- Fixed a problem where the conditions in the [Cluster](https://docs.giantswarm.io/reference/cp-k8s-api/clusters.cluster.x-k8s.io/) resource would have the wrong timestamp.

### aws-operator [v8.7.3](https://github.com/giantswarm/aws-operator/blob/master/CHANGELOG.md#873---2020-07-15)

- Added support for registry mirrors.

### coreDNS v1.6.5 (Giant Swarm app [v1.2.0](https://github.com/giantswarm/coredns-app/blob/master/CHANGELOG.md#v120-2020-07-13))

- Added readiness probe, increased liveness probe tolerance to increase stability.

### Kubernetes v1.16.13

- Mitigation of [CVE-2020-8557](https://github.com/kubernetes/kubernetes/issues/93032): Node-local denial of service via container `/etc/hosts` file.
- Mitigation of [CVE-2020-8558](https://github.com/kubernetes/kubernetes/issues/92315): Node setting allows for neighboring hosts to bypass localhost boundary.
- Mitigation of [CVE-2020-8559](https://github.com/kubernetes/kubernetes/issues/92914): Privilege escalation from compromised node to cluster. The API Server will no longer proxy non-101 responses for upgrade requests. This could break proxied backends (such as an extension API server) that respond to upgrade requests with a non-101 response code.
- Init containers are now considered for the calculation of resource requests when scheduling.

Check the [upstream changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#v11613) for details on all changes.
