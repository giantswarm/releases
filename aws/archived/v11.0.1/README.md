# :zap: Tenant Cluster Release 11.0.1 for AWS :zap:

This release fixes an issue with network traffic between node pools in 10.x and
11.x clusters reported on the 5th of February 2020. **To prevent encountering
network traffic problems between node pools, please upgrade to this release.**



### aws-operator [v8.1.1](https://github.com/giantswarm/aws-operator/releases/tag/v8.1.1)
- Fix AWS resource tags.
- Allow network traffic between Node Pools.
