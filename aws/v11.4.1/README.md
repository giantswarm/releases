# :zap: Giant Swarm Release v11.4.1 for AWS :zap:

<< Add description here >>

## Change details


### calico [3.10.4](https://github.com/projectcalico/calico/releases/tag/v3.10.4)

Not found


### etcd [3.4.9](https://github.com/etcd-io/etcd/releases/tag/v3.4.9)

See [code changes](https://github.com/etcd-io/etcd/compare/v3.4.8...v3.4.9) and [v3.4 upgrade guide](https://github.com/etcd-io/etcd/blob/master/Documentation/upgrades/upgrade_3_4.md) for any breaking changes.
#### Package `wal`
- Add [missing CRC checksum check in WAL validate method otherwise causes panic](https://github.com/etcd-io/etcd/pull/11924).
  - See https://github.com/etcd-io/etcd/issues/11918.
#### Go
- Compile with [*Go 1.12.17*](https://golang.org/doc/devel/release.html#go1.12).
<hr>



### kubernetes [1.16.9](https://github.com/kubernetes/kubernetes/releases/tag/v1.16.9)

#### Feature
- deps: Update to Golang 1.13.9
  - build: Remove kube-cross image building ([#89400](https://github.com/kubernetes/kubernetes/pull/89400), [@justaugustus](https://github.com/justaugustus)) [SIG Release and Testing]
#### Bug or Regression
- Client-go: resolves an issue with informers falling back to full list requests when timeouts are encountered, rather than re-establishing a watch. ([#89977](https://github.com/kubernetes/kubernetes/pull/89977), [@liggitt](https://github.com/liggitt)) [SIG API Machinery and Testing]
- Ensure Azure availability zone is always in lower cases. ([#89722](https://github.com/kubernetes/kubernetes/pull/89722), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix invalid VMSS updates due to incorrect cache ([#89002](https://github.com/kubernetes/kubernetes/pull/89002), [@ArchangelSDY](https://github.com/ArchangelSDY)) [SIG Cloud Provider]
- Fix panic in kubelet when running IPv4/IPv6 dual-stack mode with a CNI plugin ([#82508](https://github.com/kubernetes/kubernetes/pull/82508), [@aanm](https://github.com/aanm)) [SIG Network and Node]
- Fix the VMSS name and resource group name when updating Azure VMSS for LoadBalancer backendPools ([#89337](https://github.com/kubernetes/kubernetes/pull/89337), [@feiskyer](https://github.com/feiskyer)) [SIG Cloud Provider]
- Fix: check disk status before delete azure disk ([#88360](https://github.com/kubernetes/kubernetes/pull/88360), [@andyzhangx](https://github.com/andyzhangx)) [SIG Cloud Provider]
- Fixed a data race in kubelet image manager that can cause static pod workers to silently stop working. ([#88915](https://github.com/kubernetes/kubernetes/pull/88915), [@roycaihw](https://github.com/roycaihw)) [SIG Node]
- Fixed an issue that could cause the kubelet to incorrectly run concurrent pod reconciliation loops and crash. ([#89055](https://github.com/kubernetes/kubernetes/pull/89055), [@tedyu](https://github.com/tedyu)) [SIG Node]
- Fixes conversion error for HorizontalPodAutoscaler objects with invalid annotations ([#89968](https://github.com/kubernetes/kubernetes/pull/89968), [@liggitt](https://github.com/liggitt)) [SIG Autoscaling]
- Fixes conversion error in multi-version custom resources that could cause metadata.generation to increment on no-op patches or updates of a custom resource. ([#88995](https://github.com/kubernetes/kubernetes/pull/88995), [@liggitt](https://github.com/liggitt)) [SIG API Machinery]
- For GCE cluster provider, fix bug of not being able to create internal type load balancer for clusters with more than 1000 nodes in a single zone. ([#89902](https://github.com/kubernetes/kubernetes/pull/89902), [@wojtek-t](https://github.com/wojtek-t)) [SIG Cloud Provider, Network and Scalability]
- For volumes that allow attaches across multiple nodes, attach and detach operations across different nodes are now executed in parallel. ([#89240](https://github.com/kubernetes/kubernetes/pull/89240), [@verult](https://github.com/verult)) [SIG Apps, Node and Storage]
- Kubelet metrics gathered through metrics-server or prometheus should no longer timeout for Windows nodes running more than 3 pods. ([#87730](https://github.com/kubernetes/kubernetes/pull/87730), [@marosset](https://github.com/marosset)) [SIG Node, Testing and Windows]
- Restores priority of static control plane pods in the cluster/gce/manifests control-plane manifests ([#89970](https://github.com/kubernetes/kubernetes/pull/89970), [@liggitt](https://github.com/liggitt)) [SIG Cluster Lifecycle and Node]
#### Other (Cleanup or Flake)
- Reduce event spam during a volume operation error. ([#89794](https://github.com/kubernetes/kubernetes/pull/89794), [@msau42](https://github.com/msau42)) [SIG Storage]



### aws-operator [8.7.1-dev](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.1-dev)

Not found


### cluster-operator [2.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.0)

#### Added
- Add `deletecrs` handler for better CR cleanup.
- Add `controlPlaneStatus` handler for master nodes status.
#### Changed
- Remove controller context.
- Bump alpine version to 3.12



### containerlinux [2512.2.1](https://www.flatcar-linux.org/releases/#release-2512.2.1)

## Flatcar updates

Security fixes:
- Fix the Intel Microcode vulnerabilities ([CVE-2020-0543](https://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2020-0543))

Changes:
- A source code and licensing overview is available under `/usr/share/licenses/INFO`

Updates:
- Linux [4.19.128](https://lwn.net/Articles/822841/)
- intel-microcode [20200609](https://github.com/intel/Intel-Linux-Processor-Microcode-Data-Files/releases/tag/microcode-20200609)


### aws-cni [1.6.0](https://github.com/aws/amazon-vpc-cni-k8s/releases/tag/v1.6.0)

* Feature - [Add fallback to fetch limits from EC2 API](https://github.com/aws/amazon-vpc-cni-k8s/pull/782) (#782, @mogren)
* Feature - [Additional tags to ENI](https://github.com/aws/amazon-vpc-cni-k8s/pull/734) (#734, @nithu0115)
* Feature - [Add support for a 'no manage' tag](https://github.com/aws/amazon-vpc-cni-k8s/pull/726) (#726, @euank)
* Feature - [Use CRI to obtain pod sandbox IDs instead of Kubernetes API](https://github.com/aws/amazon-vpc-cni-k8s/pull/714) (#714, @drakedevel)
* Feature - [Add support for listening on unix socket for introspection endpoint](https://github.com/aws/amazon-vpc-cni-k8s/pull/713) (#713, @adammw)
* Feature - [Add MTU to the plugin config](https://github.com/aws/amazon-vpc-cni-k8s/pull/676) (#676, @mogren)
* Feature - [Clean up leaked ENIs on startup](https://github.com/aws/amazon-vpc-cni-k8s/pull/624) (#624, @mogren)
* Feature - [Introduce a minimum target for ENI IPs](https://github.com/aws/amazon-vpc-cni-k8s/pull/612) (#612, @asheldon)
* Feature - [Allow peered VPC CIDRs to be excluded from SNAT](https://github.com/aws/amazon-vpc-cni-k8s/pull/520) (#520, @totahuanocotl, @rewiko, @yorg1st)
* Feature - [Get container ID from kube rather than docker](https://github.com/aws/amazon-vpc-cni-k8s/pull/371) (#371, @rudoi)
* Improvement - [Make entrypoint script fail if any step fails](https://github.com/aws/amazon-vpc-cni-k8s/pull/839) (#839, @drakedevel)
* Improvement - [Place binaries in cmd/ and packages in pkg/](https://github.com/aws/amazon-vpc-cni-k8s/pull/815) (#815, @jaypipes)
* Improvement - [De-dupe calls to DescribeNetworkInterfaces](https://github.com/aws/amazon-vpc-cni-k8s/pull/808) (#808, @jaypipes)
* Improvement - [Update RollingUpdate strategy to allow 10% unavailable](https://github.com/aws/amazon-vpc-cni-k8s/pull/805) (#805, @gavinbunney)
* Improvement - [Bump github.com/vishvananda/netlink version from 1.0.0 to 1.1.0](https://github.com/aws/amazon-vpc-cni-k8s/pull/801) (#802, @ajayk)
* Improvement - [Adding node affinity for Fargate](https://github.com/aws/amazon-vpc-cni-k8s/pull/792) (#792, @nithu0115)
* Improvement - [Force ENI/IP reconciliation to delete from the datastore](https://github.com/aws/amazon-vpc-cni-k8s/pull/754) (#754, @tatatodd)
* Improvement - [Use dockershim.sock for CRI](https://github.com/aws/amazon-vpc-cni-k8s/pull/751) (#751, @mogren)
* Improvement - [Treating ErrUnknownPod from ipamd to be a noop and not returning error](https://github.com/aws/amazon-vpc-cni-k8s/pull/750) (#750, @uruddarraju)
* Improvement - [Copy CNI plugin and config in entrypoint not agent](https://github.com/aws/amazon-vpc-cni-k8s/pull/735) (#735, @jaypipes)
* Improvement - [Adding m6g instance types](https://github.com/aws/amazon-vpc-cni-k8s/pull/742) (#742, Srini Ramabadran)
* Improvement - [Remove deprecated session.New method](https://github.com/aws/amazon-vpc-cni-k8s/pull/729) (#729, @nithu0115)
* Improvement - [Scope watch on "pods" to only pods associated with the local node](https://github.com/aws/amazon-vpc-cni-k8s/pull/716) (#716, @jacksontj)
* Improvement - [Update ENI limits to match documentation](https://github.com/aws/amazon-vpc-cni-k8s/pull/710) (#710, @mogren)
* Improvement - [Reduce image layers and strip debug flags](https://github.com/aws/amazon-vpc-cni-k8s/pull/699) (#699, @mogren)
* Improvement - [Add run-integration-tests.sh script](https://github.com/aws/amazon-vpc-cni-k8s/pull/698) (#698, @nckturner)
* Improvement - [Return the error from ipamd to plugin](https://github.com/aws/amazon-vpc-cni-k8s/pull/688) (#688, @mogren)
* Improvement - [Bump aws-sdk-go to v1.23.13](https://github.com/aws/amazon-vpc-cni-k8s/pull/681) (#681, @mogren)
* Improvement - [Add support for m5n/m5dn/r5n/r5dn instances](https://github.com/aws/amazon-vpc-cni-k8s/pull/657) (#657, @Jeffwan)
* Improvement - [Add IPs to the first ENI on startup](https://github.com/aws/amazon-vpc-cni-k8s/pull/648) (#648, @mogren)
* Improvement - [Add shutdown listener](https://github.com/aws/amazon-vpc-cni-k8s/pull/645) (#645, @mogren)
* Improvement - [Made timeouts exponential](https://github.com/aws/amazon-vpc-cni-k8s/pull/640) (#640, @Zyqsempai)
* Improvement - [Remove vendor folder](https://github.com/aws/amazon-vpc-cni-k8s/pull/635) (#635, @mogren)
* Improvement - [Update protobuf to v1.3.2](https://github.com/aws/amazon-vpc-cni-k8s/pull/633) (#633, @mogren)
* Improvement - [Reduce log level to Trace for the most common Debug lines](https://github.com/aws/amazon-vpc-cni-k8s/pull/631) (#631, @mogren)
* Improvement - [Bump grpc version to v1.23.1](https://github.com/aws/amazon-vpc-cni-k8s/pull/629) (#629, @mogren)
* Improvement - [Add inCoolingPeriod for AddressInfo](https://github.com/aws/amazon-vpc-cni-k8s/pull/627) (#627, @chendotjs)
* Improvement - [Added retryNbackoff for tagENI method](https://github.com/aws/amazon-vpc-cni-k8s/pull/626) (#626, @nithu0115)
* Improvement - [Update backoff code from upstream and use when detaching ENIs](https://github.com/aws/amazon-vpc-cni-k8s/pull/623) (#623, @mogren)
* Improvement - [Update kubeconfig lookup with eksctl clusters](https://github.com/aws/amazon-vpc-cni-k8s/pull/513) (#513, @dkeightley)
* Improvement - [Fix introspection port in troubleshooting docs](https://github.com/aws/amazon-vpc-cni-k8s/pull/512) (#512, @drakedevel)
* Bug fix - [Log security groups correctly](https://github.com/aws/amazon-vpc-cni-k8s/pull/646) (#646, @mogren)
* Bug fix - [Fix WARM_ENI_TARGET=0](https://github.com/aws/amazon-vpc-cni-k8s/pull/587) (#587, @mogren)



### cert-exporter [1.2.3](https://github.com/giantswarm/cert-exporter/releases/tag/v1.2.3)

#### Changed
- Update prometheus/client_golang dependency
- Migrate from dep to go modules
- Move to App deployment



### cert-manager [1.0.8](https://github.com/giantswarm/cert-manager-app/releases/tag/v1.0.8)

#### Changed
- Allowed resource requests and limits to be configured with `values.yaml`. ([#24](https://github.com/giantswarm/cert-manager-app/pull/24))



### chart-operator [0.12.4](https://github.com/giantswarm/chart-operator/releases/tag/v0.12.4)

#### Changed
- Always set chart CR annotations so update state calculation is accurate.
- Only update failed Helm releases if the chart values or version has changed.



### cluster-autoscaler [1.16.0](https://github.com/giantswarm/cluster-autoscaler-app/releases/tag/v1.16.0)

Note that with this release we start to align the versioning scheme to the upstream project, so our v1.16.x represents upstream v1.16.x.
#### Changed
- Use cluster-autoscaler version v1.16.5.
  - This version introduces a new method to read AWS EC2 instance type details from an AWS API. Since this API is not reachable from the AWS China regions, the autoscaler is started with the `--aws-use-static-instance-list=true` flag.
- Set `scan-interval` to 30 seconds (from 10 seconds) to save resources.
- Set `scale-down-unneeded-time` to 5 minutes (from the default of 10 minutes) to release unneeded nodes earlier.
- Lower `scaleDownUtilizationThreshold` to 0.5.



### coredns [1.1.10](https://github.com/giantswarm/coredns-app/releases/tag/v1.1.10)

#### Changed
- Make resource requests/limits configurable.
- Applying Go modules.



### external-dns [1.2.2](https://github.com/giantswarm/external-dns-app/releases/tag/v1.2.2)

#### Changed
- Upgrade upstream external-dns from v0.5.18 to v0.7.2.



### kiam [1.2.2](https://github.com/giantswarm/kiam-app/releases/tag/v1.2.2)

#### Added
- Add configurable agent whitelist of proxy routes.



### kube-state-metrics [1.1.0](https://github.com/giantswarm/kube-state-metrics-app/releases/tag/v1.1.0)

#### Changed
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.
#### Fixed
- Fix invalid cluster role binding for Helm 3 compatibility.



### metrics-server [1.1.0](https://github.com/giantswarm/metrics-server-app/releases/tag/v1.1.0)

#### Changed
- Added 100.64.0.0/10 to the allowed egress subnets in NetworkPolicy.



### net-exporter [1.9.0](https://github.com/giantswarm/net-exporter/releases/tag/v1.9.0)

#### Added
- Add `ntp` collector.



### node-exporter [1.2.0](https://github.com/giantswarm/node-exporter-app/releases/tag/v1.2.0)

#### Changed
- Change Priority Class to `system-node-critical`



