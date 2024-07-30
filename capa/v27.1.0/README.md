# :zap: Giant Swarm Release v27.1.0 for CAPA :zap:

This release updates the apps and components, keeping them up to date with the latest v26 release. It also brings improvements for the container registry usage.

## Change details compared to CAPA 27.0.0

### cluster-aws [1.3.0](https://github.com/giantswarm/cluster-aws/releases/tag/v1.3.0)

#### Changed

- All workload clusters will by default use Zot registry as a pull-through cache of Azure Container Registry.

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
