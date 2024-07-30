# :zap: Giant Swarm Release v29.0.0 for CAPA :zap:

## Changes compared to v28.1.0

### Components

- cluster-aws from v1.3.0 to v2.0.0
- Kubernetes from v1.28.11 to v1.29.7

### cluster-aws [v1.3.0...v2.0.0](https://github.com/giantswarm/cluster-aws/compare/v1.3.0...v2.0.0)

#### Added

- Add `global.metadata.labels` to values schema. This field is used to add labels to the cluster resources.
- Enable `observability-policies` app.

#### Changed

- Update cluster chart to v1.1.0.
  - This sets cilium `kubeProxyReplacement` config to `"true"` instead to `"strict"` (`"strict"` has been deprecated since cilium v1.14, see [this upstream cilium](https://github.com/cilium/cilium/issues/32711) issue for more details).
- Update `ami` named template to correctly render OS image name with the new format `flatcar-stable-<flatcar version>-kube-<kubernetes version>-tooling-<capi-image-builder version>-gs`.

### Apps

- cert-exporter from v2.9.0 to v2.9.1
- cert-manager from v3.7.9 to v3.8.1
- cloud-provider-aws from v1.28.6-gs1 to v1.29.3-gs1
- cluster-autoscaler from v1.28.5-gs1 to v1.29.3-gs1
- irsa-servicemonitors from v0.0.1 to v0.1.0
- k8s-audit-metrics from v0.9.0 to v0.10.0
- k8s-dns-node-cache from v2.6.2 to v2.8.1
- net-exporter from v1.19.0 to v1.21.0
- prometheus-blackbox-exporter from v0.4.1 to v0.4.2
- security-bundle from v1.7.0 to v1.8.0
- teleport-kube-agent from v0.9.0 to v0.9.2
- vertical-pod-autoscaler from v5.2.2 to v5.2.4

### cert-exporter [v2.9.0...v2.9.1](https://github.com/giantswarm/cert-exporter/compare/v2.9.0...v2.9.1)

#### Changed

- Chart: Update PolicyExceptions to v2beta1. ([#358](https://github.com/giantswarm/cert-exporter/pull/358))

### cert-manager [v3.7.9...v3.8.1](https://github.com/giantswarm/cert-manager-app/compare/v3.7.8...v3.8.1)

#### Changed

- Bump architect-orb@5.3.1 to fix CVE-2024-24790.
- Improves `cainjector`'s Vertical Pod Autoscaler

### cloud-provider-aws [v1.28.6-gs1...v1.29.3-gs1](https://github.com/giantswarm/aws-cloud-controller-manager-app/compare/v1.28.6-gs1...v1.29.3-gs1)

#### Changed

- Chart: Update to upstream v1.29.3. ([#62](https://github.com/giantswarm/aws-cloud-controller-manager-app/pull/62))

### cluster-autoscaler [v1.28.5-gs1...v1.29.3-gs1](https://github.com/giantswarm/cluster-autoscaler-app/compare/v1.28.5-gs1...v1.29.3-gs1)

#### Changed

- Chart: Update to upstream v1.29.3. ([#279](https://github.com/giantswarm/cluster-autoscaler-app/pull/279))

### irsa-servicemonitors [v0.0.1...v0.1.0](https://github.com/giantswarm/irsa-servicemonitors-app/compare/v0.0.1...v0.1.0)

#### Changed

- Removing duplicate label ([#5](https://github.com/giantswarm/irsa-servicemonitors-app/pull/5))

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

### security-bundle [v1.7.0...v1.8.0](https://github.com/giantswarm/security-bundle/compare/v1.7.0...v1.8.0)

#### Added

- Add `kyverno-crds` app to handle Kyverno CRD install.

#### Changed

- Update `kyverno` (app) to v0.17.15. This version disables the CRD install job in favor of `kyverno-crds` App.
- Update `starboard-exporter` (app) to v0.7.11.

### teleport-kube-agent [v0.9.0...v0.9.2](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.9.0...v0.9.2)

#### Changed

- Introduced `podAntiAffinity` so `teleport-kube-agent` pods run on different `control-plane` nodes also increased the number of replicas to 3 to maintain better high availability.
- Changed the way registry is being parsed in helm templates

### vertical-pod-autoscaler [v5.2.2...v5.2.4](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.2.2...v5.2.4)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.8.3. ([#301](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/301))
- Chart: Change `restartPolicy` to `OnFailure` for the CRD job. ([#298](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/298))
