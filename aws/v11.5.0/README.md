# :zap: Giant Swarm Release v11.5.0 for AWS :zap:

This release includes a number of fixes, improvements, and picks up various component and app upgrades.

**Security:** [CVE-2020-8557](https://github.com/kubernetes/kubernetes/issues/93032), CVE-2020-8558, and [CVE-2020-8559](https://github.com/kubernetes/kubernetes/issues/92914) are mitigated through an upgrade to Kubernetes v1.16.13.

**Registry mirrors:** Dockerd is now configured to use several registries (_registry mirrors_), to avoid a single point of failure. Workloads managed by Giant Swarm are configured to make use of this. Any containers that specify a registry explicitly (like `image: reg.acme.com/namespace/repo:tag`) will not be affected. However, when no registry is named, the image will first be looked up in the DockerHub registry (docker.io) and then in the configured mirror(s).
Note: Giant Swarm installations on AWS China (cn-north-1) do not yet provide any registry mirrors as of this release.

## Change details

### cluster-operator

- Fixed a problem where the conditions in the [Cluster](https://docs.giantswarm.io/reference/cp-k8s-api/clusters.cluster.x-k8s.io/) resource would have the wrong timestamp.

### aws-operator

- Fixed a problem where the operator would crashloop when there were no tenant clusters, or rather: no Elastic Load Balancers (ELB), in the installation.

### coreDNS vX (Giant Swarm app v1.2.0)

- TODO

### Kubernetes

- Uses debian-hyperkube-base@v1.1.1 based on debian-iptables image
- Init containers are now considered for the calculation of resource requests when scheduling