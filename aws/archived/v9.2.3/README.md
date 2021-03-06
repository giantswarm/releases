# :zap: Tenant Cluster Release 9.2.3 for AWS :zap:

**If you are upgrading from 9.2.2, upgrading to this release will not roll your nodes. It will only update the apps.**

This release improves the reliability of NGINX Ingress Controller. Most important, termination configuration was adjusted so that active connections now get drained gracefully.

Below, you can find more details on components that were changed with this release.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.8](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v168-2020-04-09))

- Aligned graceful termination configuration with changes done in upstream [ingress-nginx 0.26.0](https://github.com/kubernetes/ingress-nginx/releases/tag/nginx-0.26.0)
  - Use `wait-shutdown` as preStop hook
  - Make `terminationGracePeriodSeconds` configurable
  - Set `terminationGracePeriodSeconds` by default to 5min, to align with 270 second default `worker-shutdown-timeout`.
- Default `max-worker-connections` to `0`, making it same as `max-worker-open-files` i.e. `max open files (system's limit) / worker-processes - 1024`.
  This optimizes for high load conditions where it improves performance at the cost of increasing RAM utilization (even on idle).
- HorizontalPodAutoscaler was tuned to use `targetMemoryUtilizationPercentage` of `80` due to increased RAM utilization with new default for `max-worker-connections` of `0`.
- Changed default `error-log-level` from `error` to `notice`.
- Removed use of `enable-dynamic-certificates` CLI flag, it has been deprecated since [ingress-nginx 0.26.0](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0260) via [ingress-nginx PR #4356](https://github.com/kubernetes/ingress-nginx/pull/4356).
- Added a link to the README in the sources of Chart.yaml
