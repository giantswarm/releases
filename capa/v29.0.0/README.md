# :zap: Giant Swarm Release v29.0.0 for CAPA :zap:

## Changes compared to v28.0.0

### Components

- cluster-aws from 1.0.1 to 1.3.0
- Flatcar from 3815.2.2 to 3815.2.5
- Kubernetes from 1.28.11 to 1.29.7

### cluster-aws [v1.0.1...v1.3.0]

### Changed

- Update cluster chart version to v1.0.0. This update adds MC Zot deployment as a registry mirror for `gsoci.azurecr.io` registry. This is the new default behavior.
- Apps: Fix service monitor dependencies.
- Fixed China IRSA suffix

### Removed

- Feature Gates: Remove `CronJobTimeZone`. ([#699](https://github.com/giantswarm/cluster-aws/pull/699))

### Apps

- cert-exporter from v2.9.0 to v2.9.1
- cert-manager from v3.7.6 to v3.8.1
- cilium from v0.24.0 to v0.25.1
- cloud-provider-aws from v1.28.6-gs1 to v1.29.3-gs1
- cluster-autoscaler from v1.28.5-gs1 to v1.29.3-gs1
- k8s-audit-metrics from v0.9.0 to v0.10.0
- k8s-dns-node-cache from v2.6.2 to v2.8.1
- net-exporter from v1.19.0 to v1.21.0
- prometheus-blackbox-exporter from v0.4.1 to v0.4.2
- teleport-kube-agent from v0.9.0 to v0.9.2
- vertical-pod-autoscaler from v5.2.2 to v5.2.4

### cert-exporter [v2.9.0...v2.9.1](https://github.com/giantswarm/cert-exporter/compare/v2.9.0...v2.9.1)

#### Changed

- Chart: Update PolicyExceptions to v2beta1. ([#358](https://github.com/giantswarm/cert-exporter/pull/358))

### cert-manager [v3.7.6...v3.8.1](https://github.com/giantswarm/cert-manager-app/compare/v3.7.6...v3.8.1)

#### Added

- Improves container security by setting `runAsGroup` and `runAsUser` greater than zero for all deployments.

#### Changed

- Bump architect-orb@5.3.1 to fix CVE-2024-24790.
- Improves `cainjector`'s Vertical Pod Autoscaler
- Remove quotes from acme-http01-solver-image argument. The quotes are used when looking up the image which causes an error.
- Changed the way registry is being parsed in helm templates
- Enable VPA by default

### cilium [v0.24.0...v0.25.1](https://github.com/giantswarm/cilium-app/compare/v0.24.0...v0.25.1)

#### Changed

- Fix regression setting Policy BPF Max map `policyMapMax` back to 65536 from 16384.
- Upgrade cilium to [v1.15.6](https://github.com/cilium/cilium/releases/tag/v1.15.6).

### cloud-provider-aws [v1.28.6-gs1...v1.29.3-gs1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.28.6-gs1...v1.29.3-gs1)

#### Changed

- Chart: Update to upstream v1.29.3. ([#62](https://github.com/giantswarm/aws-cloud-controller-manager-app/pull/62))

### cluster-autoscaler [v1.28.5-gs1...v1.29.3-gs1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.28.5-gs1...v1.29.3-gs1)

#### Changed

- Chart: Update to upstream v1.29.3. ([#279](https://github.com/giantswarm/cluster-autoscaler-app/pull/279))

### k8s-audit-metrics [v0.9.0...v0.10.0](https://github.com/giantswarm/k8s-audit-metrics/compare/v0.9.0...v0.10.0)

#### Changed

- Add `securityContext.readOnlyRootFilesystem` helm value (default true).

### k8s-dns-node-cache [v2.6.2...v2.8.1](https://github.com/giantswarm/k8s-dns-node-cache-app/compare/v2.6.2...v2.8.1)

#### Changed

- Make the app visible for all providers.
- Reduce security exceptions [#89](https://github.com/giantswarm/k8s-dns-node-cache-app/pull/89).
  - Enable readOnly FS moving config to emptyDir volume.
  - Remove `NET_ADMIN` and drop `ALL` capabilities.
  - Add `NET_BIND_SERVICE` capability.
  - Add policy exception for `require-non-root-groups/autogen-check-runasgroup`.
  - Remove disallow-capabilities-* policy exceptions.
- Update PolicyException CR version to v2beta1.

### net-exporter [v1.19.0...v1.21.0](https://github.com/giantswarm/net-exporter/compare/v1.19.0...v1.21.0)

#### Changed

- Enable readOnlyRootFilesystem in securityContext (#376)[https://github.com/giantswarm/net-exporter/pull/376].
- Update module google.golang.org/grpc to v1.65.0 (#373).
- Update k8s modules to v0.30.2 (#375).
- Update quay.io/giantswarm/alpine Docker tag to v3.20.1 (#372).
- Add `node` and `app` labels in ServiceMonitor.

### prometheus-blackbox-exporter [v0.4.1...v0.4.2](https://github.com/giantswarm/prometheus-blackbox-exporter-app/compare/v0.4.1...v0.4.2)

#### Changed

- Remove duplicated team label.

### teleport-kube-agent [v0.9.0...v0.9.2](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.9.0...v0.9.2)

#### Changed

- Introduced `podAntiAffinity` so `teleport-kube-agent` pods run on different `control-plane` nodes also increased the number of replicas to 3 to maintain better high availability.
- Changed the way registry is being parsed in helm templates

### vertical-pod-autoscaler [v5.2.2...v5.2.4](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.2.2...v5.2.4)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.8.3. ([#301](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/301))
- Chart: Change `restartPolicy` to `OnFailure` for the CRD job. ([#298](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/298))
