# :zap: Giant Swarm Release v18.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [14.13.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.13.0)

#### Fixed
- Use `alpine` as image for aws-cni's `routes-fixer`. 
#### Changed
- Allow externalDNS role to be assumed by any SA containing "external-dns" to allow multiple app deployments.



### cert-operator [3.0.1](https://github.com/giantswarm/cert-operator/releases/tag/v3.0.1)

#### Fixed
- Allow running unique and non unique cert-operators in the same namespace.



### app-operator [6.6.4](https://github.com/giantswarm/app-operator/releases/tag/v6.6.4)

#### Changed
- Improved feedback when searching for an app in catalog.



### cilium-prerequisites [0.1.1](https://github.com/giantswarm/cilium-prerequisites/releases/tag/v0.1.1)

#### Fixed 
- Fixed kube-linter.



### observability-bundle [0.4.0](https://github.com/giantswarm/observability-bundle/releases/tag/v0.4.0)

#### Added
- Add extra configmap and secret to `promtail-app`.
#### Changed
- Upgrade `prometheus-operator-app` to 4.2.1.
#### Removed
- Moving prometheus-operator-app user-configs to the prometheus-operator-app https://github.com/giantswarm/prometheus-operator-app/pull/249



### vertical-pod-autoscaler [3.4.2](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.2)

#### Changed
- Remove circleci job for pushing to shared app collection



### vertical-pod-autoscaler-crd [2.0.1](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.1)

#### Changed
- in [#59](https://github.com/giantswarm/vertical-pod-autoscaler-crd/pull/59) removed duplicate resources for the CRDs definition causing errors during mc-bootstrap



### kiam [3.0.0](https://github.com/giantswarm/kiam-app/releases/tag/v3.0.0)

#### Added
- Allow running without iptables mode
- Add cilium local redirect policy
- Move to Cilium Network Policies
- Move to main catalog



### aws-ebs-csi-driver [2.21.0](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.21.0)

#### Added
- Add cilium network policies.



### cert-exporter [2.4.0](https://github.com/giantswarm/cert-exporter/releases/tag/v2.4.0)

#### Added
- Add cilium network policies.



### cert-manager [2.21.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.21.0)

#### Added
- Chart: Add `CiliumNetworkPolicy`. ([#301](https://github.com/giantswarm/cert-manager-app/pull/301))



### chart-operator [2.34.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.34.0)

#### Changed
- Selecting private Helm client on demand for some operations.



### coredns [1.15.2](https://github.com/giantswarm/coredns-app/releases/tag/v1.15.2)

#### Changed
- Add `http-metrics` port to the list of exposed ports so Prometheus can access container metadata (e.g. `__meta_kubernetes_pod_container_xxx`).



### external-dns [2.35.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.35.1)

#### Changed
- Create secret from `secretConfiguration.data` value without breaking AWS Credentials values compatibility.



### kube-state-metrics [1.15.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.15.0)

#### Added
- Add cilium network policies.



### metrics-server [2.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v2.1.0)

#### Added
- Add cilium network policies.



### net-exporter [1.14.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.14.0)

#### Added
- Add `Cilium Network Policy` to net-exporter.
#### Changed
- Don't push net-exporter to capa-app-collection because it's already a default app.
- Don't push net-exporter to cloud-director-app-collection and vsphere-app-collection because it's already in default app bundles.



