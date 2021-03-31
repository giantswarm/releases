# :zap: Giant Swarm Release v13.2.0 for KVM :zap:

<< Add description here >>

## Change details


### kvm-operator [3.15.0](https://github.com/giantswarm/kvm-operator/releases/tag/v3.15.0)

#### Added
- Add vertical pod autoscaler for operator pods.
- Automatically delete TC node pods when NotReady for too long.
#### Changed
- Do not drain node pods when cluster is being deleted.
- Update for Kubernetes 1.19 compatibility.
- Update k8s-kvm to v0.4.1 with QEMU v5.2.0 and Flatcar DNS fix.
#### Fixed
- Use managed-by label to check node deployments are deleted before cluster namespace.
- Remove IPs from endpoints when the corresponding workload cluster node is not ready.



### app-operator [3.2.1](https://github.com/giantswarm/app-operator/releases/tag/v3.2.1)

Not found


### cluster-operator [0.24.2](https://github.com/giantswarm/cluster-operator/releases/tag/v0.24.2)

Not found


### cert-operator [1.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v1.0.1)

#### Fixed
- Add `list` permission for `cluster.x-k8s.io`.



### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Changed
- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.



