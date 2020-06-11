## :zap: Giant Swarm Release 11.3.3 for KVM :zap:

**If you are upgrading from 11.3.1 or 11,3,2 upgrading to this release will not roll your nodes. It will only update the apps.**

This release upgrades the chart-operator component to Helm 2.16.8, addressing a high severity security vulnerability identified in Go's `crypto` package.

Below, you can find more details on components that were changed with this release.

### chart-operator v0.13.2 ([Giant Swarm app](https://github.com/giantswarm/chart-operator/releases/tag/v0.13.2))

- Getting md5 hash from user-value more reliable way. 
- Upgrade helm `tiller` pod to v2.16.8 version which including [CVE issue](https://github.com/helm/helm/security/advisories/GHSA-cjjc-xp8v-855w) patch  
