## :zap:  Giant Swarm Release 11.3.3 for Azure :zap:

**If you are upgrading from 11.3.1 or 11.3.2, upgrading to this release will not roll your nodes. It will only update the apps.**

This release adds support for using a different CIDR for the control plane and the tenant clusters which makes it more flexible when users need to connect the tenant cluster network with their own virtual or on-premise networks.

**Note when upgrading from v11.2.x to v11.3.x:**

This release contains the replacement of CoreOS with Flatcar introduced in v11.3.0. Please carefully read release notes for 11.3.0 including Flatcar OS migration [steps](https://github.com/giantswarm/releases/tree/master/azure/v11.3.0).

**Note for future v11.3.x releases:**

Please persist this and the above note until all customers are in v11.3.0 and above.

Below, you can find more details on components that were changed with this release.

### metrics-server v0.3.3 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/metrics-server-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Add the installation's IPAM cidr to the allowed egress subnets.

### net-exporter ([v1.8.1](https://github.com/giantswarm/net-exporter/blob/master/CHANGELOG.md#181---2020-06-17))

- Add the installation's IPAM cidr to the allowed egress subnets.

### kube-state-metrics v1.9.5 ([Giant Swarm app v1.1.0](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#110---2020-06-17))

- Add the installation's IPAM cidr to the allowed egress subnets.
