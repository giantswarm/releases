# :zap: Giant Swarm Release v14.0.0 for KVM :zap:

This release upgrades Kubernetes to 1.19. A summary of relevant changes is included in these release notes. The release also includes other minor component updates summarized below the list of Kubernetes changes.


## Change details

### kubernetes [1.19.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.19.9)

#### Expanded CLI support for debugging workloads and nodes

SIG CLI expanded on debugging with `kubectl` to support two new debugging workflows: debugging workloads by creating a copy, and debugging nodes by creating a container in host namespaces. These can be convenient to:
 - Insert a debug container in clusters that don’t have ephemeral containers enabled
 - Modify a crashing container for easier debugging by changing its image, for example to busybox, or its command, for example, to `sleep 1d` so you have time to `kubectl exec`.
 - Inspect configuration files on a node's host filesystem

#### EndpointSlices are now enabled by default

EndpointSlices are an exciting new API that provides a scalable and extensible alternative to the Endpoints API. EndpointSlices track IP addresses, ports, readiness, and topology information for Pods backing a Service.

In Kubernetes 1.19 this feature will be enabled by default with kube-proxy reading from EndpointSlices instead of Endpoints. Although this will mostly be an invisible change, it should result in noticeable scalability improvements in large clusters. It will also enable significant new features in future Kubernetes releases like Topology Aware Routing.

#### Ingress graduates to General Availability

SIG Network has graduated the widely used [Ingress API](https://kubernetes.io/docs/concepts/services-networking/ingress/) to general availability in Kubernetes 1.19. This change recognises years of hard work by Kubernetes contributors, and paves the way for further work on future networking APIs in Kubernetes.

#### seccomp graduates to General Availability

The seccomp (secure computing mode) support for Kubernetes has graduated to General Availability (GA). This feature can be used to increase the workload security by restricting the system calls for a Pod (applies to all containers) or single containers.

#### KubeSchedulerConfiguration graduates to Beta

SIG Scheduling graduates `KubeSchedulerConfiguration` to Beta. The [KubeSchedulerConfiguration](https://kubernetes.io/docs/reference/scheduling/config) feature allows you to tune the algorithms and other settings of the kube-scheduler. You can easily enable or disable specific functionality (contained in plugins) in selected scheduling phases without having to rewrite the rest of the configuration. Furthermore, a single kube-scheduler instance can serve different configurations, called profiles. Pods can select the profile they want to be scheduled under via the `.spec.schedulerName` field.

#### General ephemeral volumes

Kubernetes provides volume plugins whose lifecycle is tied to a pod and can be used as scratch space (e.g. the builtin “empty dir” volume type) or to load some data in to a pod (e.g. the builtin ConfigMap and Secret volume types or “CSI inline volumes”). The new [generic ephemeral volumes](https://github.com/kubernetes/enhancements/tree/master/keps/sig-storage/1698-generic-ephemeral-volumes) alpha feature allows any existing storage driver that supports dynamic provisioning to be used as an ephemeral volume with the volume’s lifecycle bound to the Pod.
 - It can be used to provide scratch storage that is different from the root disk, for example persistent memory, or a separate local disk on that node.
 - All StorageClass parameters for volume provisioning are supported.
 - All features supported with PersistentVolumeClaims are supported, such as storage capacity tracking, snapshots and restore, and volume resizing.

#### Immutable Secrets and ConfigMaps (beta)

Secret and ConfigMap volumes can be marked as immutable, which significantly reduces load on the API server if there are many Secret and ConfigMap volumes in the cluster.
See [ConfigMap](https://kubernetes.io/docs/concepts/configuration/configmap/) and [Secret](https://kubernetes.io/docs/concepts/configuration/secret/) for more information.

#### Increase the Kubernetes support window to one year

As of Kubernetes 1.19, bugfix support via patch releases for a Kubernetes minor release has increased from 9 months to 1 year.


### kvm-operator [3.16.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.16.0)

#### Added

- Add vertical pod autoscaler configuration.
- Automatically delete WC node pods when NotReady for too long (per-cluster opt-in only).

#### Changed

- Do not drain node pods when cluster is being deleted to improve deletion time and deadlocks.
- Update for Kubernetes 1.19 compatibility.
- Update `k8s-kvm` to v0.4.1 with QEMU v5.2.0 and Flatcar DNS fix.
- Update `k8scloudconfig` to use `calico-crd-installer`.

#### Fixed

- Use `managed-by` label to check node deployments are deleted before cluster namespace.
- Remove IPs from endpoints when the corresponding workload cluster node is not ready.


### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

#### Security

- Restrict ingress to only expose the status endpoint.


### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Added

- Pause Chart CR reconciliation when it has chart-operator.giantswarm.io/paused=true annotation.

#### Changed
- Set docker.io as the default registry.
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.
- Updating namespace metadata using namespaceConfig in `Chart` CRs.
- Deploy `giantswarm-critical` PriorityClass when it's not found.

### coredns [1.4.1](https://github.com/giantswarm/coredns-app/releases/tag/v1.4.1)

#### Changed

- Set docker.io as the default registry
- Update `coredns` to upstream version [1.8.0](https://coredns.io/2020/10/22/coredns-1.8.0-release/).
- Added monitoring annotations and common labels.

### net-exporter [1.10.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.10.0)

#### Changed

- Add label selector for pods to help lower memory usage.
