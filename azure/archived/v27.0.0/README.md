# :zap: Giant Swarm Release v27.0.0 for Azure :zap:

## Changes compared to v26.0.0

### Components

- cluster-azure from v0.18.0 to v1.0.0
- Flatcar from v3815.2.4 to v3815.2.5
- Kubernetes from v1.26.15 to v1.27.16

### cluster-azure [v0.18.0...v1.0.0](https://github.com/giantswarm/cluster-azure/compare/v0.18.0...v1.0.0)

#### Changed

- Chart: Update `cluster` to v1.1.0. ([#325](https://github.com/giantswarm/cluster-azure/pull/325))
  - Machine Template: Adapt new image format.
  - Apps: Enable `observability-policies`.

### Apps

- azure-cloud-controller-manager from v1.26.22-gs2 to v1.27.18-gs1
- azure-cloud-node-manager from v1.26.22-gs2 to v1.27.18-gs1
- cert-exporter from v2.9.0 to v2.9.1
- cert-manager from v3.7.6 to v3.8.1
- k8s-audit-metrics from v0.9.0 to v0.10.0
- k8s-dns-node-cache from v2.6.2 to v2.8.1
- net-exporter from v1.19.0 to v1.21.0
- observability-bundle from v1.3.4 to v1.5.3
- observability-policies v0.0.1
- security-bundle from v1.7.1 to v1.8.0
- teleport-kube-agent from v0.9.0 to v0.9.2
- vertical-pod-autoscaler from v5.2.2 to v5.2.4

### azure-cloud-controller-manager [v1.26.22-gs2...v1.27.18-gs1](https://github.com/giantswarm/azure-cloud-controller-manager-app/compare/v1.26.22-gs2...v1.27.18-gs1)

#### Changed

- Chart: Update to upstream v1.27.18. ([#81](https://github.com/giantswarm/azure-cloud-controller-manager-app/pull/81))

### azure-cloud-node-manager [v1.26.22-gs2...v1.27.18-gs1](https://github.com/giantswarm/azure-cloud-node-manager-app/compare/v1.26.22-gs2...v1.27.18-gs1)

#### Changed

- Chart: Update to upstream v1.27.18. ([#70](https://github.com/giantswarm/azure-cloud-node-manager-app/pull/70))

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

### observability-bundle [v1.3.4...v1.5.3](https://github.com/giantswarm/observability-bundle/compare/v1.3.4...v1.5.3)

#### Added

- Add `alloy` v0.3.0 as `alloy-logs`

#### Changed

- Rename `alloy-logs` app to camel case `alloyLogs`.
- Fix CNP issues (allow traffic from pods in kube-system to nginx-ingress-controller)
  - Upgrade `grafana-agent` to 0.4.5.
  - Upgrade `alloy` to 0.3.1.
  - Upgrade `promtail` to 1.5.4.
- Upgrade `prometheus-operator-crd` to 11.0.1.
- prometheus-operator will not check promql syntax for prometheusRules that are labelled `application.giantswarm.io/prometheus-rule-kind: loki`
- Upgrade `kube-prometheus-stack` to 11.0.0 and `prometheus-operator-crd` to 11.0.0. This upgrade mainly consists in:
  - kube-prometheus-stack dependency chart upgraded from [56.21.2](https://github.com/prometheus-community/helm-charts/releases/tag/kube-prometheus-stack-56.21.2) to [61.0.0](https://github.com/prometheus-community/helm-charts/releases/tag/kube-prometheus-stack-61.0.0)
  - prometheus upgrade from 2.50.1 to [2.53.0](https://github.com/prometheus-community/helm-charts/releases/tag/prometheus-25.22.0)
  - thanos ruler upgrade from 0.34.1 to [0.35.1](https://github.com/thanos-io/thanos/releases/tag/v0.35.1)
  - kube-state-metrics from 2.10.0 to 2.12.0
  - prometheus-operator from 0.71.2 [0.75.0](https://github.com/prometheus-operator/prometheus-operator/releases/tag/v0.75.0) - adding remoteWrite.proxyFromEnvironment and Scrape Class support
  - prometheus-node-exporter upgraded from 1.8.0 to [1.8.1](https://github.com/prometheus/node_exporter/releases/tag/v1.8.1)
- Upgrade `grafana-agent` from 0.4.3 to 0.4.4
  - This version enables the override the grafana agent `CiliumNetworkPolicy` egress and ingress sections.

### observability-policies [v0.0.1](https://github.com/giantswarm/observability-policies-app/releases/v0.0.1)

#### Added

- Add a ClusterPolicy to prevent prometheus-operator CRDs deletion.
- Create `observability-policies` app to deploy Kyverno Observability Policies into clusters.

### security-bundle [v1.7.1...v1.8.0](https://github.com/giantswarm/security-bundle/compare/v1.7.1...v1.8.0)

#### Added

- Add `kyverno-crds` app to handle Kyverno CRD install.

#### Changed

- Update `kyverno` (app) to v0.17.15. This version disables the CRD install job in favor of `kyverno-crds` App.

### teleport-kube-agent [v0.9.0...v0.9.2](https://github.com/giantswarm/teleport-kube-agent-app/compare/v0.9.0...v0.9.2)

#### Changed

- Introduced `podAntiAffinity` so `teleport-kube-agent` pods run on different `control-plane` nodes also increased the number of replicas to 3 to maintain better high availability.
- Changed the way registry is being parsed in helm templates

### vertical-pod-autoscaler [v5.2.2...v5.2.4](https://github.com/giantswarm/vertical-pod-autoscaler-app/compare/v5.2.2...v5.2.4)

#### Changed

- Chart: Update Helm release vertical-pod-autoscaler to v9.8.3. ([#301](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/301))
- Chart: Change `restartPolicy` to `OnFailure` for the CRD job. ([#298](https://github.com/giantswarm/vertical-pod-autoscaler-app/pull/298))
