# :zap: Giant Swarm Release v18.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [14.13.0](https://github.com/giantswarm/aws-operator/releases/tag/v14.13.0)

#### Fixed
- Use `alpine` as image for aws-cni's `routes-fixer`. 
#### Changed
- Allow externalDNS role to be assumed by any SA containing "external-dns" to allow multiple app deployments.



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



### vertical-pod-autoscaler [3.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.4.0)

#### Added
- Add cilium network policies.



### vertical-pod-autoscaler-crd [2.0.0](https://github.com/giantswarm/vertical-pod-autoscaler-crd/releases/tag/v2.0.0)

#### Changed
- Synced with new upstream repo
#### Changed
- Remove `push-to-app-collection` jobs for onprem providers since this app became a part of default apps bundle.
#### Added
- Add icon to Chart.yml for use in happa



### kiam [3.0.0](https://github.com/giantswarm/kiam-app/releases/tag/v3.0.0)

#### Added
- Allow running without iptables mode
- Add cilium local redirect policy
- Move to Cilium Network Policies
- Move to main catalog



