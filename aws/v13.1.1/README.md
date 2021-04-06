# :zap: Giant Swarm Release v13.1.1 for AWS :zap:

<< Add description here >>

## Change details


### external-dns [2.3.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.3.0)

#### Changed
- Change default annotation filter to match the one we use for the nginx ingress controller.
#### Added
- Add sidecar container for `provider: aws` to periodically validate IAM credential acessibility ([#76](https://github.com/giantswarm/external-dns-app/pull/76))



