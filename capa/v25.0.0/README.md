# :zap: Giant Swarm Release v25.0.0 for CAPA :zap:

We are happy to announce our first `Cluster API for AWS` (CAPA) release v25. 

This is the first Giant Swarm supported CAPA release. It is available on CAPA Management Clusters and will be used as a first release to be upgraded to from `Vintage` workload clusters.

## CAPA benefits

Each existing customer using the Giant Swarm Vintage AWS product has been given a presentation about CAPA benefits. We have gathered the most crucial high-level advantages in the list below to refresh your memory:

- Easier cluster management, complete production ready clusters packaged as a Helm chart, more GitOps oriented approach
- Better visibility of subresources, clear API
- More transparent upgrades
- Improved cluster configuration validation
- Management clusters capable of deploying clusters in different regions
- Flexible network configuration
- Automatic certificate renewal in place for worker nodes
- Exposing more features based on the upstream implementations and contributions
- [Karpenter](https://karpenter.sh/) solution integrated into workload clusters out of the box
- More configurable registries with credentials and mirrors

## Important changes

Besides the benefits listed above, we have also presented changes that are introduced with CAPA. Here is the summary of most important points:

- Flat DNS structure - Bringing more flexibility into workload cluster management across regions. Old and new DNS are both available during and after migration, while new DNS needs to be used after migration. We will synchronize with customers for the old DNS phase out after the migrations per management cluster.
- Only public DNS zone is supported - Private hosted zones are no longer created. However, both zones (public and private) will be available during and after the migration. We will eventually work together to remove the private zones with each customer. Ingresses that depend on the cluster base domain must be updated after the migration.
- Cluster apps moved to `org-` namespaces - This change will allow to simplify RBAC and enable GitOps managed clusters to be created with a pre-defined set of applications.
- One set of controllers - On contrary to Vintage, where each release had a set of independent operators, the CAPA solution will manage all releases with a single set of controllers. This will result in fewer components running on the management clusters, hence lowering costs of operations as well as speeding up any required hotfixes.
- `GP2` volumes not supported - The majority of customers is already using `gp3` volumes, as we are refreshing the infrastructure, deprecated `kubernetes.io/aws-ebs` provisioner creating `gp2` volumes will also be removed
- Nodepools defaults are changed - With CAPA by default, if not explicitly configured, nodepools will now share the single subnet. Migrated Vintage clusters will not be affected. With that change the node limit will be determined by cluster VPC size, where additional VPC CIDRs can be added.
- [Teleport](https://docs.giantswarm.io/vintage/platform-overview/security/cluster-security/cluster-access/#admin-access-via-teleport) - Giant Swarm will now utilize `Teleport` application for the direct node access. It helps to strengthen security and simplifies compliance with regulations and network topology.


## Vintage to CAPA Migration

The migration itself is a fully automated process ran by Giant Swarm engineers using a `migration-cli` that handles all infrastructure as well as workload migrations. 
*The experience in the migration process should be the same as in usual upgrade of the WC itself.*

Prior to running the tool, a new Management Cluster based on the CAPA solution has to be created in order to fully make use of the CAPI lifecycle management as well as infrastructure.

As Giant Swarm manages the cloud infrastructure, there are no actions needed from customers for the migration itself. 
We have aimed to match the Vintage features as close as possible, introducing improvements where needed. 

One of the many improvements is actually deprecation of the `k8s-initiator` application, which was allowing to customize some parts of the kubernetes environment, catering for customer needs. 
This tool however was bringing a lot of freedom in terms of bash implementation that was ran in the tool itself. We have reviewed the use-cases that customers have implemented, exposed certain settings in CAPA and prepared a migration plan for those features as well as allowing any future customization. 
The most important part for each customer is to prepare the `{cluster_name}-migration-configuration` yaml file, representing the `k8s-initiator` app features used, which will be consumed then by the `migration-cli` and be populated in Cluster charts for future usage.

Your Account Engineer will provide you with a detailed checklist to go over prior to migration.

## New components with CAPA

### (NEW) capi-node-labeler [0.5.0](https://github.com/giantswarm/capi-node-labeler-app/releases/tag/v0.5.0)

### (NEW) aws-pod-identity-webhook [1.16.0](https://github.com/giantswarm/aws-pod-identity-webhook/releases/tag/v1.16.0)

### (NEW) teleport-kube-agent [0.9.0](https://github.com/giantswarm/teleport-kube-agent-app/releases/tag/v0.9.0)

### (NEW) cluster-aws [1.0.0](https://github.com/giantswarm/cluster-aws/releases/tag/v1.0.0)

### (NEW) cilium-crossplane-resources [0.1.0](https://github.com/giantswarm/cilium-crossplane-resources/releases/tag/v0.1.0)
    

## Change details compared to Vintage AWS 20.1.1

### cloud-provider-aws (formerly aws-cloud-controller-manager-app) [1.25.14-gs3](https://github.com/giantswarm/aws-cloud-controller-manager-app/releases/tag/v1.25.14-gs3)

#### Changed

- Reduce minimum CPU and memory requests.

### cert-manager [3.7.6](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.5)

#### Added

- Added Vertical Pod Autoscaler support for cainjector pods.

### cilium [0.24.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.24.0)

#### Added

- Cilium ENI mode for CAPA becomes usable with these changes.
    - Add security group tag filter for pod network.
    - Select subnets from secondary VPC CIDRs.
- Upgrade cilium to v1.15.4.

### cluster-autoscaler [1.27.3-gs9](https://github.com/giantswarm/cilium-app/releases/tag/v0.24.0)

#### Added

- Node Group Auto Discovery for CAPA MachinePools.
- Add service account annotations as value.
- Added service monitor.
- Add configurable node.nodeSelector in values.
- Add configurable node.caBundlePath in values.
- Repository: Add ABS & ATS.
- Helpers: Add fullname template.
- Helpers: Add application.giantswarm.io/team label.
- Deployment: Tolerate control-plane.
- Add possibility to use egress proxy.

#### Changed

- Update cluster-autoscaler to version 1.27.3.
- Change ScaleDownUtilizationThreshold default from 0.5 to 0.7
- Replace condition for PSP CR installation.
- Configure gsoci.azurecr.io as the default container image registry.
- Change helm value managementCluster from object to string.
- Chart: Make PSS compliant.
- Chart: Respect .Release.Name & .Release.Namespace.
- Values: Move balancingIgnoreLabels to configmap.
- Values: Rename clusterAutoscalerResources to resources.
- Values: Rename verticalPodAutoscaler to autoscaling.
- Helpers: Rename templates.
- Helpers: Keep -app suffix in name template.
- Deployment: Improve nodeSelector compatibility.
- Improve proxy settings
- Reduce minimum CPU and memory requests.


#### Removed

- Values: Remove port & serviceType.
- Helpers: Remove k8s-addon label.
- RBAC: Remove PSP role & binding.
- Deployment: Remove replicas.
- Deployment: Remove imagePullPolicy.
- Deployment: Remove aws-use-static-instance-list workaround.
- Policy Exception: Remove -exceptions suffix.
- Service: Replace by pod monitor.

### k8s-dns-node-cache [2.6.2](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.6.2)

#### Changed

- Reduce CPU requests

### net-exporter [1.19.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.19.0)

#### Added

- Add /blackbox endpoint.

### security-bundle [1.7.0](https://github.com/giantswarm/security-bundle/releases/tag/v1.7.0)

#### Added

- Add cloudnative-pg, edgedb, and reports-server apps (disabled).


#### Changed

- Update starboard-exporter (app) to v0.7.10.
- Update kyverno (app) to v0.17.13.
- Update trivy (app) to v0.12.0.
- Update trivy-operator (app) to v0.9.0.
- Update cloudnative-pg (app) to v0.0.5.

### vertical-pod-autoscaler [5.2.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v5.2.2)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.8.0.
- Chart: Update appVersion and README.md.
- Chart: Update appVersion and README.md, VPA v1.1.2.

