## :zap: Giant Swarm Release 11.3.1 for KVM :zap:

This release fixes a problem that prevented clusters with OIDC user and group prefix settings to work as expected in `11.3.0`.

### kvm-operator [v3.11.1](https://github.com/giantswarm/kvm-operator/releases/tag/v3.11.1)
- Use Release.Revision in Helm chart for Helm 3 support.
- Fix OIDC settings which are passed to masters.
