# :zap: Giant Swarm Release 12.0.1 for KVM :zap:

This release updates NGINX Ingress Controller to the latest upstream release.
Most importantly, it includes a fix for a regression introduced in the previous upstream release related to `use-regex` and `rewrite` annotations.

Below, you can find more details on components that were changed with this release.

---

### nginx-ingress-controller [1.7.3](https://github.com/giantswarm/nginx-ingress-controller-app/releases/tag/v1.7.3)

- Upgraded upstream `ingress-nginx` from 0.34.0 to 0.34.1 - [changelog](https://github.com/kubernetes/ingress-nginx/blob/master/Changelog.md#0341).
