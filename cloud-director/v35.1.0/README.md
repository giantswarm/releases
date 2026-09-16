# :zap: Giant Swarm Release v35.1.0 for VMware Cloud Director :zap:

## Changes compared to v35.0.1

### Components

- cluster-cloud-director from v7.0.0 to v7.3.0
- cluster from v8.0.0 to v8.3.0

### cluster [v8.0.0...v8.3.0](https://github.com/giantswarm/cluster/compare/v8.0.0...v8.3.0)

#### Added

- Add `preKubeadmCommandsTemplateName` and `postKubeadmCommandsTemplateName` hooks under `providerIntegration.controlPlane.kubeadmConfig` and `providerIntegration.workers.kubeadmConfig`. They name a provider template that renders a YAML list of additional kubeadm commands, once for the control plane and once per node pool for workers.
- Add `internal.advancedConfiguration.kubelet.evictionHard` values. Providers need them to tell autoscalers such as Karpenter how much of a node's resources is allocatable.
- SELinux: Add `global.components.selinux.writablePolicyStore` value (default `true`) to allow loading additional SELinux policies.

#### Changed

- Enable the `ClusterTrustBundle` and `ClusterTrustBundleProjection` feature gates (Kubernetes 1.33+) and the `PodCertificateRequest` feature gate (Kubernetes 1.35+) by default on kube-apiserver, kube-controller-manager and kubelet.
- SELinux: Keep AVC audit logs (required for SELinux policy generation).
- SELinux: Relabel the whole filesystem except read-only `/usr` (previously only `/etc/kubernetes`).
- SELinux: Correctly label CA certificates in `/etc/ssl/certs` for mounting into containers.
- App to HR Migration: Skip v35.0.0 pre-releases and update `docker-kubectl` to v1.36.4.
- Chart: Rework HelmRelease clean-up job.
