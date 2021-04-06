# :zap: Giant Swarm Release v13.1.1 for AWS :zap:

This release provides a bug fix for the external-dns-app.

> **_Warning:_** The nginx app needs to be updated to `v1.14.0+` because a new version of `external-dns` is included in this release.

## Change details


### external-dns [2.3.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.3.0)

#### Changed
- Change default annotation filter to match the one we use for the nginx ingress controller.
#### Added
- Add sidecar container for `provider: aws` to periodically validate IAM credential acessibility ([#76](https://github.com/giantswarm/external-dns-app/pull/76))



