# :zap: Giant Swarm Release v20.0.0 for AWS :zap:

<< Add description here >>

## Change details


### kubernetes [1.25.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.25.9)

#### Feature
- Kubernetes is now built with Go 1.19.8 ([#117131](https://github.com/kubernetes/kubernetes/pull/117131), [@xmudrii](https://github.com/xmudrii)) [SIG Release and Testing]
#### Bug or Regression
- Fix missing delete events on informer re-lists to ensure all delete events are correctly emitted and using the latest known object state, so that all event handlers and stores always reflect the actual apiserver state as best as possible ([#115900](https://github.com/kubernetes/kubernetes/pull/115900), [@odinuge](https://github.com/odinuge)) [SIG API Machinery]
- Fix: Route controller should update routes with NodeIP changed ([#116361](https://github.com/kubernetes/kubernetes/pull/116361), [@lzhecheng](https://github.com/lzhecheng)) [SIG Cloud Provider]
- Fixes a regression in the pod binding subresource to honor the `metadata.uid` precondition.
  This allows kube-scheduler to ensure it is assigns node names to the same instances of pods it made scheduling decisions for. ([#116776](https://github.com/kubernetes/kubernetes/pull/116776), [@alculquicondor](https://github.com/alculquicondor)) [SIG API Machinery and Testing]
- Kubelet: Fix fs quota monitoring on volumes ([#116794](https://github.com/kubernetes/kubernetes/pull/116794), [@pacoxu](https://github.com/pacoxu)) [SIG Storage]
#### Other (Cleanup or Flake)
- Service session affinity timeout tests are no longer required for Kubernetes network plugin conformance due to variations in existing implementations. New conformance tests will be developed to better express conformance in future releases. ([#112806](https://github.com/kubernetes/kubernetes/pull/112806), [@dcbw](https://github.com/dcbw)) [SIG Architecture, Network and Testing]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/apiserver-network-proxy/konnectivity-client: v0.0.35 → v0.0.36
#### Removed
_Nothing has changed._



### cilium [0.18.0](https://github.com/giantswarm/cilium-app/releases/tag/v0.18.0)

#### Changed
- Upgrade cilium to `1.14.3`.



### cluster-autoscaler [1.25.1-gs2](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.25.1-gs2)

#### Fixed
- Adjusted minimum allowed CPU and memory



### metrics-server [2.4.2](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.4.2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### observability-bundle [1.0.0](https://github.com/giantswarm/observability-bundle/releases/tag/v1.0.0)

#### Changed
- *!Breaking change*: Simplify configuration for the bundled apps. See our [upgrade guide](./docs/upgrade.md)
  - Move all user configs from under `apps.<appName>.userConfig` from string to regular helm values to `userConfig.<appName>`
  - Rename `prometheus-operator-app` to `kube-prometheus-stack`
  - Rename `promtail-app` to `promtail`
- Enforce `Cilium Network Policy` by default.
- Enforce `Pod Security Standard` by default.
- Upgrade `kube-prometheus-stack` to 8.1.1 and `prometheus-operator-crd` to 8.0.0
- Upgrade `grafana-agent` to 0.3.2.



### security-bundle [1.5.0](https://github.com/giantswarm/security-bundle/releases/tag/v1.5.0)

#### Added
- Add a `global.namespace` value for automatically setting all app namespaces.
#### Changed
- Update to exception-recommender (app) to v0.0.7.
- Update to falco (app) to v0.7.1.
- Update to jiralert (app) version v0.1.3.
- Update to kyverno-policy-operator (app) version v0.0.6.
- Update to starboard-exporter (app) version v0.7.7.
- Update to trivy-operator (app) to v0.5.1.



### cert-exporter [2.9.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.9.0)

#### Added
- Add cert name to secret metric.



### k8s-dns-node-cache-app [2.6.0](https://github.com/giantswarm/k8s-dns-node-cache-app/releases/tag/v2.6.0)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### aws-ebs-csi-driver [2.28.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.28.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### prometheus-blackbox-exporter [0.4.1](https://github.com/giantswarm/prometheus-blackbox-exporter/releases/tag/v0.4.1)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### node-exporter [1.18.2](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.18.2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### chart-operator [3.1.2](https://github.com/giantswarm/chart-operator/releases/tag/v3.1.2)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### coredns [1.21.0](https://github.com/giantswarm/coredns-app/releases/tag/v1.21.0)

#### Changed
- Configure `gsoci.azurecr.io` as the default container image registry.



### cert-manager [2.25.3](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.25.3)




