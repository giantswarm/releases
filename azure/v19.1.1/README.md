# :zap: Giant Swarm Release v19.1.1 for Azure :zap:

<< Add description here >>

## Change details


### containerlinux [3510.2.3](https://www.flatcar-linux.org/releases/#release-3510.2.3)

_Changes since **Stable 3510.2.2**_
 
#### Security fixes:
 
- Linux ([CVE-2022-48425](https://nvd.nist.gov/vuln/detail/CVE-2022-48425))
 
#### Updates:
 
- Linux ([5.15.113](https://lwn.net/Articles/932883) (includes [5.15.112](https://lwn.net/Articles/932134)))
- ca-certificates ([3.90](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_90.html))


### vertical-pod-autoscaler [3.5.3](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v3.5.3)

#### Added
- Add `cluster-autoscaler safe-to-evict` annotation to `recommender` and `updater`



### observability-bundle [0.7.1](https://github.com/giantswarm/observability-bundle/releases/tag/v0.7.1)

#### Changed
- Upgrade `promtail-app` to 1.1.1.
- Upgrade `prometheus-operator-app` to 5.0.6.


