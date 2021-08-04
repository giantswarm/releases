# :zap: Giant Swarm Release v15.2.0-eric for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [10.6.1-c8311b47e8686e8eb0b6eda1ed03641d4438f473](https://github.com/giantswarm/aws-operator/releases/tag/v10.6.1-c8311b47e8686e8eb0b6eda1ed03641d4438f473)

Not found


### kubernetes [1.21.3](https://github.com/kubernetes/kubernetes/releases/tag/v1.21.3)

#### Feature
- Kubernetes is now built with Golang 1.16.6 ([#103670](https://github.com/kubernetes/kubernetes/pull/103670), [@cpanato](https://github.com/cpanato)) [SIG Cloud Provider, Instrumentation, Release and Testing]
- Updates the following images to pick up CVE fixes:
  - `debian` to v1.8.0
  - `debian-iptables` to v1.6.5
  - `setcap` to v2.0.3 ([#103235](https://github.com/kubernetes/kubernetes/pull/103235), [@thejoycekung](https://github.com/thejoycekung)) [SIG API Machinery, Release and Testing]
#### Bug or Regression
- Fix scoring for NodeResourcesMostAllocated and NodeResourcesBalancedAllocation plugins when nodes have containers with no requests. This was leaving to under-utilization of small nodes. ([#102925](https://github.com/kubernetes/kubernetes/pull/102925), [@alculquicondor](https://github.com/alculquicondor)) [SIG Scheduling]
- ServiceOwnsFrontendIP shouldn't report error when the public IP doesn't match ([#102516](https://github.com/kubernetes/kubernetes/pull/102516), [@nilo19](https://github.com/nilo19)) [SIG Cloud Provider]
- Switch scheduler to generate the merge patch on pod status instead of the full pod ([#103133](https://github.com/kubernetes/kubernetes/pull/103133), [@marwanad](https://github.com/marwanad)) [SIG Scheduling]
- VSphere: Fix regression during attach disk if datastore is within a storage folder or datastore cluster. ([#102969](https://github.com/kubernetes/kubernetes/pull/102969), [@gnufied](https://github.com/gnufied)) [SIG Cloud Provider]
#### Dependencies
#### Added
_Nothing has changed._
#### Changed
- sigs.k8s.io/structured-merge-diff/v4: v4.1.0 → v4.1.2
#### Removed
_Nothing has changed._



### containerlinux [2765.2.6](https://www.flatcar-linux.org/releases/#release-2765.2.6)


**Security fixes**



*   Linux ([CVE-2020-26558](https://nvd.nist.gov/vuln/detail/CVE-2020-26558), [CVE-2021-0129](https://nvd.nist.gov/vuln/detail/CVE-2021-0129), [CVE-2020-24587](https://nvd.nist.gov/vuln/detail/CVE-2020-24587), [CVE-2020-24586](https://nvd.nist.gov/vuln/detail/CVE-2020-24586), [CVE-2020-24588](https://nvd.nist.gov/vuln/detail/CVE-2020-24588), [CVE-2020-26139](https://nvd.nist.gov/vuln/detail/CVE-2020-26139), [CVE-2020-26145](https://nvd.nist.gov/vuln/detail/CVE-2020-26145), [CVE-2020-26147](https://nvd.nist.gov/vuln/detail/CVE-2020-26147), [CVE-2020-26141](https://nvd.nist.gov/vuln/detail/CVE-2020-26141), [CVE-2021-3564](https://nvd.nist.gov/vuln/detail/CVE-2021-3564), [CVE-2021-28691](https://nvd.nist.gov/vuln/detail/CVE-2021-28691), [CVE-2021-3587](https://nvd.nist.gov/vuln/detail/CVE-2021-3587), [CVE-2021-3573](https://nvd.nist.gov/vuln/detail/CVE-2021-3573))

**Bug fixes**



*   Update-engine sent empty requests when restarted before a pending reboot ([Flatcar#388](https://github.com/kinvolk/Flatcar/issues/388))
*   motd login prompt list of failed services: The output of "systemctl list-units --state=failed --no-legend" contains a bullet point which is not expected and ended up being taken as the unit name of failed units which was previously on the start of the line. Filtered the bullet point out to stay compatible with the old behavior in case upstream would remove the bullet point again. ([coreos-overlay#1042](https://github.com/kinvolk/coreos-overlay/pull/1042))

**Updates**



*   Linux ([5.10.43](https://lwn.net/Articles/859022/))


