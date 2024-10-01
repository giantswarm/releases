# :zap: Giant Swarm Release v20.1.5 for AWS :zap:

<< Add description here >>

## Change details


### cilium [0.25.1](https://github.com/giantswarm/cilium-app/releases/tag/v0.25.1)

#### Changed
- Upgrade cilium to v1.15.6.
- Fix regression setting Policy BPF Max map `policyMapMax` back to 65536 from 16384.

#### Added

Cilium ENI mode for CAPA becomes usable with these changes:
- Add security group tag filter for pod network
- Select subnets from secondary VPC CIDRs


