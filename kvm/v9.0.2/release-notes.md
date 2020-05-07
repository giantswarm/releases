## :zap: Giant Swarm Release 9.0.2 for KVM :zap:

This release fixes a problem that prevented clusters with OIDC user and group prefix settings to work as expected in `9.0.1`.

### kvm-operator [v3.9.2](https://github.com/giantswarm/kvm-operator/releases/tag/v3.9.2)
- Use Release.Revision in Helm chart for Helm 3 support.
- Fix OIDC settings which are passed to masters.
