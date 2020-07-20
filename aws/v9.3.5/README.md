# :zap: Giant Swarm Release v9.3.5 for AWS :zap:

**If you are upgrading from 9.3.4, upgrading to this release will not roll your nodes. It will only update the apps.**

This release updates NGINX Ingress Controller to the latest upstream release.
Most importantly update includes a fix for the regression introduced in the previous upstream release related to `use-regex` and `rewrite` annotations.

Below, you can find more details on components that were changed with this release.

### nginx-ingress-controller v0.34.1 ([Giant Swarm app v1.7.3](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v173-2020-07-16))

- Upgraded upstream ingress-nginx from v0.34.0 to v0.34.1 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0341).
