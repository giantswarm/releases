# :zap: Giant Swarm Release v25.1.0 for CAPA :zap:

This release updates the components, keeping them up to date with Vintage AWS v20.1.x series. Several improvements for Vintage to CAPA migration have also been included.

## Change details compared to CAPA 25.0.0

### cluster-aws [1.1.0](https://github.com/giantswarm/cluster-aws/releases/tag/v1.1.0)

### Fixed
- Fixed China IRSA suffix

#### Added
- Add the Management Cluster name as a tag to the AWS resources created by CAPA.
- Add the node pool name as a tag to the AWS resources associated with the node pool.

#### Changed
- Update cluster chart to 0.35.0


### cert-manager [3.7.9](https://github.com/giantswarm/cert-manager-app/releases/tag/v3.7.9)

#### Fix
- Remove quotes from acme-http01-solver-image argument. The quotes are used when looking up the image which causes an error.

#### Update
- Improves container security by setting `runAsGroup` and `runAsUser` greater than zero for all deployments.

### containerlinux [3815.2.5](https://www.flatcar-linux.org/releases/#release-3815.2.5)

 _Changes since **Stable 3815.2.4**_
 
#### Security fixes:
 
 - openssh ([CVE-2024-6387](https://nvd.nist.gov/vuln/detail/CVE-2024-6387))

#### Updates:
 
 - Linux ([6.1.96](https://lwn.net/Articles/979851))
 - openssh ([9.7_p1](https://www.openssh.com/txt/release-9.7))

### cilium [0.25.1](https://github.com/giantswarm/cilium-app/releases/tag/v0.25.1)

#### Changed
- Fix regression setting Policy BPF Max map policyMapMax back to 65536 from 16384.
- Upgrade cilium to v1.15.6.

