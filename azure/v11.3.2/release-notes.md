## :zap:  Giant Swarm Release 11.3.2 for Azure :zap:

**If you are upgrading from 11.3.1, upgrading to this release will not roll your nodes. It will only update the apps.**

This release improves the reliability of NGINX Ingress Controller.

Specifically, liveness probe is configured to be more fault tolerant than readiness probe. This helps shed load when it goes beyond replica capacity, speeding up recovery when NGINX gets overloaded.

Below, you can find more details on components that were changed with this release.

### nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.12](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v1612-2020-06-04))

- Made healthcheck probes configurable.
- Made liveness probe more resilient.

## external-dns v0.5.18 ([Giant Swarm app v1.2.1](https://github.com/giantswarm/external-dns-app/blob/master/CHANGELOG.md#v121-2020-05-29))

- Prefer CNAMEs record sets for AWS SDK configuration with explicit credentials.
