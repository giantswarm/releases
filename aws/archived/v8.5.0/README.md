# :zap:  Tenant Cluster Release 8.5.0 for AWS :zap:

**IMPORTANT**: Due to cgroup restructure (change can be found [here](https://github.com/giantswarm/k8scloudconfig/pull/564)) and more resource reservation for core components, we are now switching to bigger master instance types (from m4.large to m4.xlarge). This adds some room for additional workload scheduling on master.
This release also contains multiple fixes to the private/public hosted zones which were altered in the previous release. In-Cluster communication between services and the Kubernetes API should now stay inside the VPC without issues.

On the feature side, the release introduces an internal-api endpoint, which allows communication with Kubernetes API via the private network. This allows you, for example, to restrict public Kubernetes API access to Giant Swarm VPN addresses only and use internal API via peered networks.

### aws-operator v5.4.0
- Duplicate etcd recordset into public hosted zone.
- Add ingress internal load-balancer in private hosted zone.
- Use private subnets for internal Kubernetes API loadbalancer.
- Refactor Kubernetes API whitelisting with public/private subnets.
- Add internal-api recordset into public hosted zone.

### cluster-operator v0.20.0
- Install chart-operator from default app catalog.
