## :zap:  Tenant Cluster Release 11.3.1 for Azure :zap:

This release brings improvements to the upgrade process from clusters using CoreOS image to Flatcar image.
The temporary firewall rule for workers was changed to be more focused in blocking traffic towards the API load balancer.

While upgrading to this release from CoreOS according to tests there will be 0% request loss on both ingress and egress traffic.

**Customers running clusters on 11.3.0**

This release should be solely used for clusters that are not yet running on 11.3.0 (Flatcar image) and its existance is to improve the upgrade process from CoreOS to Flatcar.
This release does not add any value if you upgrade from 11.3.0 to 11.3.1 thus you can discard this release/upgrade if you are already running on 11.3.0

**Customers upgrading from 11.2.x to 11.3.1**

Please also read carefully release notes for 11.3.0 including Flatcar OS migration [steps](https://github.com/giantswarm/releases/tree/master/azure/v11.3.0).
