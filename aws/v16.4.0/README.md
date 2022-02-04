# :zap: Giant Swarm Release v16.4.0 for AWS :zap:

This release provides a new feature to automatically rotate the Kubernetes API key used to encrypt secret data in etcd. It also includes the latest version of Calico.

**Highlights**
- Automation of the Kubernetes API key used to encrypt secret data in etcd;
- Calico `v3.21.3`;
- Fixed IAM policy to select ARNs of EBS snapshots for AWS China;
- `kiam-watchdog` now probes STS instead of Route53 by default.

**How does the Kubernetes API key rotation work?**

* The rotation is disabled by default and has to be enabled by setting the `encryption.giantswarm.io/enable-rotation` annotation on the `${CLUSTER-ID}-encryption-provider-config` secret;
* The key rotation happens if the key is at least 180 days old (counting from creation timestamp on `${CLUSTER-ID}-encryption-provider-config` secret or from last key rotation). It can also be forced by setting the `encryption.giantswarm.io/force-rotation` annotation to start the rotation process immediately;
* A new config is generated containing the new and old keys as soon as the process starts;
* The next step requires a roll of the control plane nodes (either manually or during an update);
* After the control plane nodes have been rolled and are using the new encryption configuration, the `encryption-provider-operator` will rewrite all secrets. This leads to the re-encryption of all secrets with the new key;
* The operator will remove the old encryption key after all the secrets are rewritten.


## Change details


### aws-ebs-csi-driver [2.8.1](https://github.com/giantswarm/aws-ebs-csi-driver-app/releases/tag/v2.8.1)

#### Fixed
- Use node selector according to control-plane and nodepool labels.


### aws-operator [10.14.1](https://github.com/giantswarm/aws-operator/releases/tag/v10.14.1)

#### Fixed
- Autoselect region ARN for ebs snapshots.

#### Changed
- Changes to EncryptionConfig in order to fully work with `encryption=provider-operator`.


### cluster-operator [3.13.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.13.0)

#### Changed
- Removed encryption key creation. Encryption keys will be managed by `encryption-provider-operator`.


### calico [3.21.3](https://github.com/projectcalico/calico/releases/tag/v3.21.3)

#### BGP Improvements
For users of BGP you can now view the status of your BGP routers, including session status, RIB / FIB contents, and agent health via the new CalicoNodeStatus API. See the API documentation for more details.
In addition, you can control BGP advertisement of certain prefixes using the new disableBGPExport option on each IP pool, allowing greater control of your route sharing scheme.
Pull requests:
- Added Calico node status resource (CalicoNodeStatus) which represents a collection of status information for a node that Calico reports back to the user for use during troubleshooting. libcalico-go #1502 (@song-jiang)
- Report node BGP status from calico/node. node #1234 (@song-jiang)
- Add new syncer for BGP status API. typha #662 (@song-jiang)
- Don’t export BGP routes for IP pools that have disableBGPExport==true confd #647 (@coutinhop)
#### Service-based network policy improvements
In v3.20, we introduced egress policy rules that can match on Kubernetes services. In v3.21, we improved upon that in two ways. First, you can now use service matches in Calico NetworkPolicy and GlobalNetworkPolicy ingress rules. Second, you can now use service-based network policy rules on Windows nodes.
Pull requests:
- Policy ingress rules now support service selectors. felix #3024 (@mgleung)
- Windows data plane support for Service-based network policy rules felix #2917 (@caseydavenport)
- Allow services to be specified in the Source field of Ingress rules libcalico-go #1517 (@mgleung)
#### Option to run Calico as non-privileged and non-root
Calico can now optionally run in non-privileged and non-root mode, with some limitations. See the documentation for more information.
Pull requests:
- Change node and supporting binary permissions so that they can be run as a non-root user node #1224 (@mgleung)
- CNI plugin now sets route_localnet=1 for container interfaces cni-plugin #1168 (@mgleung)
- CNI plugins now have SUID bit set in order to run as non-root cni-plugin #1168 (@mgleung)
#### IPReservations API
You can use the new IPReservations API to reserve certain IP addresses so that they will not be used by Calico IPAM. This allows for fine-grained control of the IP space in your cluster.
Pull requests:
- Add support for IPReservations libcalico-go #1509 (@fasaxc)
#### Bug fixes
- Fix a serious regression introduced in v3.21.0 where the datastore watcher could get stuck and report stale information in clusters with >500 policies/pods/etc. The bug was triggered by needing to do a resync (for example after an etcd compaction) when there were enough resources to trigger the list pager. [calico #5332](https://github.com/projectcalico/calico/pull/5332) (@robbrockbank)
- Pass ExceptUpgradeService param to stop-calico.ps1 as well node #1372 (@lmm)
- Restrict Typha server to FIPS compliant cipher suites. typha #696 (@caseydavenport)
- Fix log spam from Calico upgrade service for Windows node #1343 (@song-jiang)
- Increase timeout for setting NetworkUnavailable on shutdown node #1341 (@caseydavenport)
- Fix potential panic and memory leak in kube-controllers caused by adding and subsequently deleting IPAM blocks kube-controllers #912 (@caseydavenport)
- IPAM GC correctly handles multiple IP addresses allocated with the same handle ID. kube-controllers #903 (@caseydavenport)
- Fix bug where invalid port structures were being sent to Felix, preventing pods with hostPorts specified from working. libcalico-go #1545 (@caseydavenport)
- Downgrade repetitive info level logging in calico/node autodetection code node #1237 (@caseydavenport)
- Updated ubi base images and CentOS repos to stop CVE false positives from being reported. node #1136 (@coutinhop)
- Fixed typo in umount command pod2daemon #64 (@ScheererJ)
- Fixes this bug which caused WireGuard stats to be collected even when WireGuard was disabled. Additionally, the version of the wgctrl dependency has been updated as the previous version caused thread leaks. felix #3057 (@mikestephen)
- Fix blackhole route table interface matches to handle empty interface regexes. felix #3007 (@robbrockbank)
- Fix slow performance when updating a Kubernetes namespace when there are many Pods (and in turn, slow startup performance when there are many namespaces). felix #2964 (@fasaxc)
- Close race condition that could result in an extra IPAM block being allocated to a node. libcalico-go #1488 (@caseydavenport)
- Fix that podIP annotation could be incorrectly clobbered for stateful set pods: https://github.com/projectcalico/calico/issues/4710 libcalico-go #1472 (@fasaxc)
- Fix removal of old CNI configuration on name-change cni-plugin #1153 (@caseydavenport)
- Readiness depends on all syncers typha #613 (@robbrockbank)
- Exclude RR nodes from BGP full mesh confd #619 (@coutinhop)
- Fixed a bug in ExternalTrafficPolicy=Local that lead to connection stalling. felix #3015 (@tomastigera)
- Fixed broken connections when client used the same port to connect to the same backed via a nodeport on different nodes. felix #2983 (@tomastigera)
- The eBPF mode implementation of DoNotTrack policy was incorrectly allowing an inbound connection through a HostEndpoint, when the HostEndpoint had DoNotTrack policy for the ingress direction but not for egress. For precise compatibility with Calico’s established DoNotTrack semantics, that connection should be disallowed, and now is. (Because of the lack of connection tracking, successful use of DoNotTrack policy to allow flows requires configuring the DoNotTrack policy symmetrically in both directions.) felix #2982 (@neiljerram)
#### Other changes
- Replace github.com/dgrijalva/jwt-go with active fork github.com/golang-jwt/jwt that resolves vulnerability flagged by scanners. libcalico-go #1554 (@lmm)
- calico/node logs write to /var/log/calico within the container by default, in addition to stdout node #1133 (@song-jiang)
- Read pod IP information from Amazon VPC CNI annotation, if present on the pod. libcalico-go #1523 (@caseydavenport)
- Update etcd client version to v3.5.0 libcalico-go #1495 (@Aceralon)
- Optimize lists and watches made against the Kubernetes API libcalico-go #1484 (@caseydavenport)
- WorkloadEndpoints now support hostPorts libcalico-go #1471 (@AloysAugustin)
- Include CNI plugin release v1.0.0 cni-plugin #1141 (@caseydavenport)
- Allow configuration of num_queues for Calico created veth interfaces cni-plugin #1116 (@arikachen)
- Typha now gives newly connected clients an extra grace period to catch up after sending the snapshot to reduce the possibility of cyclic disconnects. typha #614 (@fasaxc)
- Add calico-node upgrade service for upgrades on Windows node #1254 (@lmm)
- eBPF arm64/aarch64 node #1044 (@frozenprocess)
- BPF: Endpoints in EndpointsSlices that are not ready are excluded from NAT felix #3017 (@tomastigera)
- Calico’s eBPF dataplane now fully implements DoNotTrack policy felix #2910 (@neiljerram)
- Add HostPort support in the gRPC dataplane cni-plugin #1119 (@AloysAugustin)

### external-dns [2.9.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.9.0)

This release contains some changes to mitigate rate limiting on AWS clusters. Please take note of the defaults
for values `aws.batchChangeInterval`, `aws.zonesCacheDuration`, `externalDNS.interval`
and `externalDNS.minEventSyncInterval`.
If you already specify `--aws-batch-change-interval` or `--aws-zones-cache-duration`, please migrate to the new values `aws.batchChangeInterval` and `aws.zonesCacheDuration`.
#### Added
- Allow to set `--aws-batch-change-interval` through `aws.batchChangeInterval` value. Default `10s`.
- Allow to set `--aws-zones-cache-duration` through `aws.zonesCacheDuration` value. Default `3h`.
#### Changed
- Set default `externalDNS.interval` to `5m`.
- Set default `externalDNS.minEventSyncInterval` to `30s`.
- Allow setting Route53 credentials (`externalDNS.aws_access_key_id` and `externalDNS.aws_secret_access_key`) indepentent from `aws.access` value.
- Allow setting the AWS default region (`aws.region`) indepentent from `aws.access` value.
- Allow to omit the `--domain-filter` flag completely by setting `externalDNS.domainFilterList` to `null`.
- 
### kiam-watchdog [0.5.1](https://github.com/giantswarm/kiam-watchdog/releases/tag/v0.5.1)

### Added

- Added --probe-mode flag to allow using either 'route53' or 'sts' to probe AWS API. 'sts' is the default.
