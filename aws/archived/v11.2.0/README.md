# :zap: Tenant Cluster Release v11.2.0 for AWS :zap:

This release adds support for [EC2 Spot Instances](https://aws.amazon.com/ec2/spot/) in your node pools to allow reduction of your compute resource cost. Find details on how to use spot instances [in our node pools documentation](https://docs.giantswarm.io/basics/nodepools/#instance-distribution).

Our web UI has been updated to support the new features of this release. Users of  [`gsctl`](https://github.com/giantswarm/gsctl), please update to the latest v0.20.0.

If you intend to use spot instances, be aware that an AWS service limit called `EC2 Spot Instances` is effective. In order to prepare for growth, you should request an increase for this limit for all instance types you would like to use.

## aws-operator [v8.4.0](https://github.com/giantswarm/aws-operator/releases/tag/v8.4.0)

- Add support for instance distribution between spot and on-demand in worker node Auto Scaling Groups (ASG).
- Add mixed instance support for worker node Auto Scaling Groups (ASG).
- Place master node in Auto Scaling Group (ASG) as preparation to run multiple master nodes in the future.

## cluster-operator [v2.1.9](https://github.com/giantswarm/cluster-operator/blob/master/CHANGELOG.md#219-2020-04-23)

- Once a cluster creation has been completed, we set `.status.infrastructureReady` in the `Cluster` resource to `true`.
- Fixed RBAC rules for the reconciliation of the Control Plane resources.

## kube-state-metrics v1.9.5 (Giant Swarm app [v1.0.5](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.0.5))

- Change upstream version to 1.9.5
