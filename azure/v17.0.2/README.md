# :zap: Giant Swarm Release v17.0.2 for Azure :zap:

This release provides quality of life improvements and bug fixes to operators and apps.

## Change details


### app-operator [5.12.0](https://github.com/giantswarm/app-operator/releases/tag/v5.12.0)

#### Added
- Add `initialBootstrapMode` flag to allow deploying CNI as managed apps.



### chart-operator [2.24.0--component](https://github.com/giantswarm/chart-operator/releases/tag/v2.24.0--component)

Not found


### azure-operator [5.20.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.20.0)

#### Changed
- Bumped k8scc to latest version to fix `localhost` node name problem.



### kube-state-metrics [1.10.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.10.0)

#### Changed
- Add `Node Pool` labels to the default allowed labels in `--metric-labels-allowlist`.



### vertical-pod-autoscaler [2.4.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.4.0)

#### Changed
- Use patched docker image tagged `0.10.0-oomfix` for `recommender` and updater (see https://github.com/giantswarm/roadmap/issues/923).



### node-exporter [1.12.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.12.0)

#### Changed
- Enabled `diskstats` collector.



