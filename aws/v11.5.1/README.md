# :zap: Giant Swarm Release v11.5.1 for AWS :zap:

<< Add description here >>

## Change details


### cluster-operator [2.3.2](https://github.com/giantswarm/cluster-operator/releases/tag/v2.3.2)

- Handle error basedomain not found gracefully, so that G8sControlPlane CR and MachineDeployment CRs can be deleted



### aws-operator [8.7.5](https://github.com/giantswarm/aws-operator/releases/tag/v8.7.5)

#### Changed
- Adjust number of host network pods on worker node for aws-cni



### calico [3.15.1](https://github.com/projectcalico/calico/releases/tag/v3.15.1)

#### Bug fixes
 - Fix issue with service IP advertisement breaking host service connectivity [confd #337](https://github.com/projectcalico/confd/pull/337) (@caseydavenport)



