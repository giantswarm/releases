# :zap: Giant Swarm Release v18.0.0 for AWS :zap:

<< Add description here >>

## Change details


### cluster-operator [4.5.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.5.0)

#### Added
- Add cilium app config map.



### containerlinux [3227.2.1](https://www.flatcar-linux.org/releases/#release-3227.2.1)

New Stable Release 3227.2.1

Changes since Stable 3227.2.0

## Security fixes:

- Linux ([CVE-2022-23816](https://nvd.nist.gov/vuln/detail/CVE-2022-23816), [CVE-2022-23825](https://nvd.nist.gov/vuln/detail/CVE-2022-23825), [CVE-2022-29900](https://nvd.nist.gov/vuln/detail/CVE-2022-29900), [CVE-2022-29901](https://nvd.nist.gov/vuln/detail/CVE-2022-29901))

## Bug fixes:

- Added support for Openstack for cloud-init activation ([flatcar-linux/init#76](https://github.com/flatcar-linux/init/pull/76))
- Excluded Wireguard interface from `systemd-networkd` default management ([Flatcar#808](https://github.com/flatcar-linux/Flatcar/issues/808))
- Fixed `/etc/resolv.conf` symlink by pointing it at `resolv.conf` instead of `stub-resolv.conf`. This bug was present since the update to systemd v250 ([coreos-overlay#2057](https://github.com/flatcar-linux/coreos-overlay/pull/2057))
- Fixed excluded interface type from default systemd-networkd configuration ([flatcar-linux/init#78](https://github.com/flatcar-linux/init/pull/78))
- Fixed space escaping in the `networkd` Ignition translation ([Flatcar#812](https://github.com/flatcar-linux/Flatcar/issues/812))

## Changes:


## Updates:

- Linux ([5.15.58](https://lwn.net/Articles/902917) (includes [5.15.57](https://lwn.net/Articles/902317), [5.15.56](https://lwn.net/Articles/902101)))
- ca-certificates ([3.81](https://firefox-source-docs.mozilla.org/security/nss/releases/nss_3_81.html))



### external-dns [2.15.1](https://github.com/giantswarm/external-dns-app/releases/tag/v2.15.1)

#### Changed
- Update alpine image to v3.15.5 ([#178](https://github.com/giantswarm/external-dns-app/pull/178))



### chart-operator [2.27.0](https://github.com/giantswarm/chart-operator/releases/tag/v2.27.0)

#### Added
- Ensure the `giantswarm` PriorityClass is created first on initial installation.



### etcd-kubernetes-resources-count-exporter [0.5.1](https://github.com/giantswarm/etcd-kubernetes-resources-count-exporter/releases/tag/v0.5.1)

#### Changed
- Bump build image to alpine 3.16.1.



### vertical-pod-autoscaler [2.5.0](https://github.com/giantswarm/vertical-pod-autoscaler-app/releases/tag/v2.5.0)

#### Changed
- Upgrade vertical-pod-autoscaler to [0.11.0](https://github.com/kubernetes/autoscaler/releases/tag/vertical-pod-autoscaler-0.11.0)
  Potentially breaking change:
  - Added validation - CPU values will be accepted only with resolution of 1 mCPU, memory with resolution of 1 b
  Other changes:
    - Switch to go 1.16
    - Admission controller now logs when it fails to start
    - Increase resolution of admission_latency_seconds metric
    - Reduce verbosity of some logs



