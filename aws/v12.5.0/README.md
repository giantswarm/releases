# :zap: Giant Swarm Release v12.5.0 for AWS :zap:

This release offers the possibility to add additional Network Pools to the Control Plane and flexibly choose the IP range for new Tenant Clusters from these pool. It also upgrades Kubernetes to v1.17.13.

**Note for Solution Engineers:**

- Helm3:
  - Please use [Upgrading tenant clusters to Helm 3](https://intranet.giantswarm.io/docs/dev-and-releng/helm/helm3-tenant-cluster-upgrade/) as a guide on the upgrade process for the checks and monitoring steps.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.3.x** and above.
- cert-manager-app:
  - Please use this [upgrade script](https://github.com/giantswarm/cert-manager-app/blob/master/files/migrate-v090-to-v200.sh) to assist with the process. Due to changes in Cert Manager's API, associated Ingresses and Secrets must also be updated to ensure they are reconciled by Cert Manager.
  - **Note for future 12.x.x releases:**
    - Please persist this note and the above, until all customers are on AWS **v12.1.x** and above.

## Change details


### aws-operator [9.1.2](https://github.com/giantswarm/aws-operator/releases/tag/v9.1.2)

#### Added
- Add etcd client certificates for Prometheus.
- Add `--service.aws.hostaccesskey.role` flag.
- Add `api.<cluster ID>.k8s.<base domain>` and `*.<cluster ID>.k8s.<base domain>` records into CP internal hosted zone.
#### Fixes
- Fix `vpc`/`route-table` lookups.
#### Changed
- Access Control Plane AWS account using role assumption. This is to prepare
  running aws-operator inside a Tenant Cluster.
- Changed AWS CNI parameters to be more conservative with preallocated IPs while not hitting the AWS API too hard.
#### Changed
- Update `k8scloudconfig` to `v8.0.3`.



### cluster-operator [3.3.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.3.1)

#### Fixed
- Manage Tenant Cluster API errors gracefully.

### Kubernetes [1.17.13](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v11712)

* Prevent logging of docker config contents if file is malformed ([#95348](https://github.com/kubernetes/kubernetes/pull/95348))
* Do not fail sorting empty elements. ([#94666](https://github.com/kubernetes/kubernetes/pull/94666))
* Fix detach azure disk issue when vm not exist ([#95177](https://github.com/kubernetes/kubernetes/pull/95177))
* Fix etcd_object_counts metric reported by kube-apiserver ([#94817](https://github.com/kubernetes/kubernetes/pull/94817))
* Fix kubectl printer to correctly handle timestamps of events emitted using events.k8s.io API ([#90227](https://github.com/kubernetes/kubernetes/pull/90227))
* Fix the cloudprovider_azure_api_request_duration_seconds metric buckets to correctly capture the latency metrics. Previously, the majority of the calls would fall in the "+Inf" bucket. ([#95376](https://github.com/kubernetes/kubernetes/pull/95376))
* Fix: detach azure disk broken on Azure Stack ([#94885](https://github.com/kubernetes/kubernetes/pull/94885))
* Fixed a bug where improper storage and comparison of endpoints led to excessive API traffic from the endpoints controller ([#94935](https://github.com/kubernetes/kubernetes/pull/94935))
* Kubeadm: warn but do not error out on missing "ca.key" files for root CA, front-proxy CA and etcd CA, during "kubeadm join --control-plane" if the user has provided all certificates, keys and kubeconfig files which require signing with the given CA keys. ([#94988](https://github.com/kubernetes/kubernetes/pull/94988))

Check the [upstream changelog](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.17.md#changelog-since-v11712) for details on all changes.
