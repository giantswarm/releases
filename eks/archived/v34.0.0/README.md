# :zap: Giant Swarm Release v34.0.0 for EKS :zap:

This release upgrades Kubernetes to v1.34.4 and adds several apps that can be enabled on demand (they come disabled by default).

## Changes compared to v33.0.0

### Components

- Kubernetes from v1.33.8 to [v1.34.4](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.34.md#v1.34.4)

### Apps

- Added aws-ebs-csi-driver-servicemonitors v0.1.0
- Added aws-nth-bundle v1.3.0
- Added cert-exporter v2.9.15
- Added cert-manager v3.9.4
- Added cert-manager-crossplane-resources v0.1.0
- Added chart-operator-extensions v1.1.2
- Added cilium v1.3.4
- Added cilium-crossplane-resources v0.2.1
- Added cilium-servicemonitors v0.1.3
- Added cluster-autoscaler v1.34.1-1
- Added coredns v1.29.1
- Added coredns-extensions v0.1.2
- Added external-dns v3.4.0
- Added k8s-dns-node-cache v2.9.1
- Added karpenter v1.4.0
- Added karpenter-crossplane-resources v0.5.1
- Added karpenter-taint-remover v1.0.1
- Added metrics-server v2.7.0
- Added net-exporter v1.23.0
- Added network-policies v0.1.3
- Added node-exporter v1.20.10
- Added node-problem-detector v0.5.2
- Added observability-bundle v2.5.0
- Added observability-policies v0.0.3
- Added priority-classes v0.3.0
- Added prometheus-blackbox-exporter v0.5.0
- Added security-bundle v1.16.1
- Added vertical-pod-autoscaler v6.1.1
- Added vertical-pod-autoscaler-crd v4.1.1
- teleport-kube-agent from v0.10.7 to v0.10.8

### priority-classes [v0.3.0](https://github.com/giantswarm/priority-classes/releases/tag/v0.3.0)

#### Changed

- Label now uses chart version instead of app version.

#### Removed

- Removed appVersion (only version is used now).

### teleport-kube-agent [v0.10.7...v0.10.8](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.10.7...v0.10.8)

#### Added

- Add `io.giantswarm.application.audience` and `io.giantswarm.application.managed` chart annotations for Backstage visibility.

#### Changed

- Migrate chart metadata annotations to OCI-compatible format.
