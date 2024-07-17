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
- prometheus-blackbox-exporter from v0.4.1 to v0.4.2
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
