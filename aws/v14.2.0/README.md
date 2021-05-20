# :zap: Giant Swarm Release v14.2.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [10.3.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.3.0)

#### Fixed
- Updated OperatorKit to v4.3.1 for Kubernetes 1.20 support.
- Cancel update loop if source or target release is not found.
- Updated IPAM library to avoid IP conflicts.
#### Added
- Clean up VPC peerings from a cluster VPC when is cluster deleted.
- Clean up Application and Network loadbalancers created by Kubernetes when cluster is deleted.
- Add new flatcar AMIs.
#### Changed
- Fix issues with etcd initial cluster resolving into ELB and causing errors.
- Update `k8scloudconfig` to version `v10.5.0` to support kubernetes `v1.20`.
- Use `networkctl reload` for managing networking to avoid bug in `systemd`.



### cert-manager [2.7.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.7.0)

- Update to upstream `v1.3.1` ([#155](https://github.com/giantswarm/cert-manager-app/pull/155)). This mitigates failed cert-manager-app installations due to CRD conversion issues.



### cluster-autoscaler [1.19.3](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.19.3)

Not found


