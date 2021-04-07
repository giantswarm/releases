# :zap: Tenant Cluster Release v9.3.2 for AWS :zap:

This release fixes an issue where upgrading from an earlier platform release to platform release v9.3.0 or v9.3.1 would result in the `nginx-ingress-controller` service having no endpoints, causing the NGINX IC app to not work.

Solution Engineers have applied a manual fix to affected clusters. This release provides a working upgrade path for clusters using AWS 9.x.x and older.

## nginx-ingress-controller v0.30.0 ([Giant Swarm app v1.6.11](https://github.com/giantswarm/nginx-ingress-controller-app/blob/master/CHANGELOG.md#v1611-2020-05-26))

- Align labels, using `app.kubernetes.io/name` instead of `k8s-app` where possible. `k8s-app` remains to be used for compatibility reasons.
