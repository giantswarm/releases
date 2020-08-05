# :zap: Giant Swarm Release v9.3.5 for AWS :zap:

This release updates NGINX Ingress Controller to the latest upstream release.
Most importantly, it includes a fix for a regression introduced in the previous upstream release related to `use-regex` and `rewrite` annotations.

Release also includes Kubernetes upgrade to 1.16.12, among other things to address a kube-proxy crash looping issue.

Below, you can find more details on components that were changed with this release.

### Kubernetes 1.16.12

- Updated from Kubernetes 1.16.11 -
changelog since [v1.16.11](https://github.com/kubernetes/kubernetes/blob/master/CHANGELOG/CHANGELOG-1.16.md#changelog-since-v11611)
- Added a wrapper script for iptables in the recently upgraded hyperkube image to avoid crash looping of kube-proxy.

### nginx-ingress-controller v0.34.1 ([Giant Swarm app v1.7.3](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v173-2020-07-16))

- Upgraded upstream ingress-nginx from v0.34.0 to v0.34.1 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0341).
