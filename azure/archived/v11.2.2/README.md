## ⚡️ Tenant Cluster Release 11.2.2 for Azure ⚡️

This release contains improvements in the upgrade procedure.
We have improved new nodes discovery at the stage where the cluster size is doubled with new versions of the machines.
We experienced issues with stuck upgrades where not enough quota was available, breaking the state of azure-operator at new nodes discovery. Thanks to improvements no manual action will be needed from Giant Swarm side to continue the upgrade when the quotas are met again.

The release also includes reliability improvements to these pre-installed apps: NGINX Ingress Controller and kube-state-metrics.

To upgrade, please contact your Solution Engineer.

Below, you can find more details on components that were changed with this release.

### azure-operator v3.0.5

- Improved the discovery of new nodes.

### kube-state-metrics v1.9.5 ([Giant Swarm app v1.0.5](https://github.com/giantswarm/kube-state-metrics-app/blob/master/CHANGELOG.md#v105))

- Upgraded to kube-state-metrics [new release 1.9.5](https://github.com/kubernetes/kube-state-metrics/releases/tag/v1.9.5).

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.7](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v167-2020-04-08))

- Align graceful termination configuration with changes done in upstream [ingress-nginx 0.26.0](https://github.com/kubernetes/ingress-nginx/releases/tag/nginx-0.26.0)
  - Use `wait-shutdown` as preStop hook,
  - Make `terminationGracePeriodSeconds` configurable,
  - Set `terminationGracePeriodSeconds` by default to 5min, to align with 270 second default `worker-shutdown-timeout`.
