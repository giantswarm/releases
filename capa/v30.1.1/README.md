# :zap: Giant Swarm Release v30.1.1 for CAPA :zap:

<< Add description here >>

## Changes compared to v30.1.0

### Components

### Apps

* security-bundle from 1.10.0 to [1.10.1](https://github.com/giantswarm/security-bundle/compare/v1.10.0...v1.10.1)

### security-bundle [1.10.1](https://github.com/giantswarm/security-bundle/compare/v1.10.0...v1.10.1)

#### Important

**Note:** Kyverno `PolicyExceptions` (API group `kyverno.io`) versions `v2alpha1` and `v2beta1` are deprecated and will be removed in the next Kyverno minor release (v1.14). Please update all Kyverno PolicyExceptions to `v2`. No action is required for Giant Swarm Policy API `PolicyExceptions` (API group `policy.giantswarm.io`), which are handled automatically.

#### Changed

- Update `kyverno-crds` (app) to v1.13.1.
