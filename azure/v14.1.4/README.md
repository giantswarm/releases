# :zap: Giant Swarm Release v14.1.4 for Azure :zap:

This is a bug fix release that involves the `external-dns` and `chart-operator` apps.

Upgrade from 14.1.3 to 14.1.4 will only roll the apps.

## Change details


### external-dns [2.3.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.3.0)

#### Changed
- Change default annotation filter to match the one we use for the nginx ingress controller.
#### Added
- Add sidecar container for `provider: aws` to periodically validate IAM credential acessibility ([#76](https://github.com/giantswarm/external-dns-app/pull/76))


### chart-operator [2.12.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.12.0)

#### Changed

- Set docker.io as the default registry
- Pass RESTMapper to helmclient to reduce the number of REST API calls.
- Updated Helm to v3.5.3.
