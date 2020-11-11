# :zap: Giant Swarm Release v13.0.0-beta1 for Azure :zap:

This is the first beta release to support **Kubernetes 1.18** on Azure.

**Important Warning** 
During master upgrade from 12.x to 13.0.0, within the time frame of 30 seconds we had noticed a spike in requests failures. This is most likely caused by Azure CNI upgrade and despite our great efforts, we had not found a solution to maintain upgrade path and avoid this disturbance. Please keep this in mind when scheduling an upgrade window, and contact your SE if you have further questions.

This release marks the first official support for Node Pools and GPU instances on Azure.

Apart from that, the release contains many changes in other components, including important security fixes in Kubernetes and Calico.

Below, you can find more details on components that were changed with this release.

## Change details

### kubernetes [1.18.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.18.md#v11810)
### app-operator [2.7.0](https://github.com/giantswarm/app-operator/blob/master/CHANGELOG.md#270---2020-11-09)
### azure-operator [5.0.0-beta1](https://github.com/giantswarm/azure-operator/blob/master/CHANGELOG.md#500-beta1---2020-11-12)
### calico [3.15.3](https://github.com/projectcalico/calico/releases/tag/v3.15.3)
### chart-operator [2.5.0](https://github.com/giantswarm/chart-operator/blob/master/CHANGELOG.md#250---2020-11-09)
### cluster-operator [0.23.18](https://github.com/giantswarm/cluster-operator/blob/legacy/CHANGELOG.md#02318---2020-10-21)
### containerlinux [2605.7.0](https://www.flatcar-linux.org/releases/#release-2605.7.0)
### etcd [3.4.13](https://github.com/etcd-io/etcd/releases/tag/v3.4.13)
