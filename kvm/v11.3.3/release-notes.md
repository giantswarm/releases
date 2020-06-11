## :zap: Giant Swarm Release 11.3.3 for KVM :zap:

**If you are upgrading from 11.3.1 or 11,3,2 upgrading to this release will not roll your nodes. It will only update the apps.**

This release includes a security fix from helm 2.16.8.

Below, you can find more details on components that were changed with this release.

### chart-operator v0.13.2 ([Giant Swarm app](https://github.com/giantswarm/chart-operator/releases/tag/v0.13.2))

- Upgrade helm `tiller` pod to v2.16.8 version which including [CVE issue](https://github.com/helm/helm/security/advisories/GHSA-cjjc-xp8v-855w) patch  
