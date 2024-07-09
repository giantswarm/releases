# :zap: Giant Swarm Release v20.1.4 for AWS :zap:

Fixes an issue with quotes in cert-manager acme-http01-solver-image argument.

## Change details


### cert-manager [3.7.9](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.9)

#### Fix
- Remove quotes from acme-http01-solver-image argument. The quotes are used when looking up the image which causes an error.



