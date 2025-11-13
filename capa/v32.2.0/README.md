# :zap: Giant Swarm Release v32.2.0 for CAPA :zap:

<< Add description here >>

## Changes compared to v32.1.0

### Components

- cluster-aws from v5.3.0 to v6.2.0
- Flatcar from v4230.2.4 to [v4459.2.0](https://www.flatcar.org/releases/#release-4459.2.0)
- Kubernetes from v1.32.9 to [v1.32.10](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.32.md#v1.32.10)

### cluster-aws [v5.3.0...v6.2.0](https://github.com/giantswarm/cluster-aws/compare/v5.3.0...v6.2.0)

#### Added

- Add `capa-karpenter-taint-remover` to handle CAPA - Karpenter taint race condition.
- Add standard tags to IRSA infrastructure.
- Expose value to configure `terminationGracePeriod` in the karpenter node pools.

#### Changed

- Change default consolidation time to 6 hours to avoid constant node rolling.
- Rename `capa-karpenter-taint-remover` app.
- Set `terminationGracePeriod` default to 30m, to avoid having `karpenter` nodes stuck in `Deleting` state due to `Pods` blocking the deletion i.e. PDBs.
- Chart: Update `cluster` to v4.2.0.
- The container registry passed as value to default apps is set to `gsoci.azurecr.io`, regardless of the cluster region. The mirroring feature of `containerd` will make sure the right registry is used.
- Switch to HelmReleases to install `karpenter` and `karpenter-crossplane-resources` charts.
- Bump flux `HelmReleases` api version to v2.
- Reduce heartbeat timeout for ASG lifecycle hooks to from 30 minutes to 3 minutes since aws-node-termination-handler-app (NTH) can now send heartbeats
- Configure the following `startupTaints` to help `karpenter` ignore pending `Pods` due to these taints that will be removed after the node starts, avoiding unnecessary instance provisioning:
  - `node.cluster.x-k8s.io/uninitialized:NoSchedule`
  - `node.cilium.io/agent-not-ready:NoSchedule`
  - `ebs.csi.aws.com/agent-not-ready:NoExecute`
- Include `cilium` ENI mode pod CIDRs in the NodePort Services security group ingress rules.

#### Removed

- Removed `capi-node-labeler` app. From now on, the worker nodes won't have the `node-role.kubernetes.io/worker` or `node.kubernetes.io/worker` labels.

### Apps

- etcd-defrag from v1.2.2 to v1.2.3
- security-bundle from v1.14.0 to v1.15.0

### etcd-defrag [v1.2.2...v1.2.3](https://github.com/giantswarm/etcd-defrag-app/compare/v1.2.2...v1.2.3)

#### Changed

- Chart: Update dependency ahrtr/etcd-defrag to v0.36.0. ([#69](https://github.com/giantswarm/etcd-defrag-app/pull/69))

### security-bundle [v1.14.0...v1.15.0](https://github.com/giantswarm/security-bundle/compare/v1.14.0...v1.15.0)

#### Added

- Add `kubescape` (app) version v0.0.4.

#### Changed

- Update `kyverno` (app) to v0.21.1.
- Update `kyverno-crds` (app) to v1.15.0.
