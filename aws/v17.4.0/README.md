# :zap: Giant Swarm Release v17.4.0 for AWS :zap:

<< Add description here >>

## Change details


### aws-operator [cni-containerd-sock](https://github.com/giantswarm/aws-operator/releases/tag/vcni-containerd-sock)

Not found


### cluster-operator [4.3.0](https://github.com/giantswarm/cluster-operator/releases/tag/v4.3.0)

#### Changed
- Do not update "app-operator.giantswarm.io/version" label on app-operators when their value is 0.0.0 (aka they are reconciled by the management cluster app-operator). This is a use-case for App Bundles for example, because the App CRs they contain should be created in the MC so should be reconciled by the MC app-operator.



### external-dns [2.14.0](https://github.com/giantswarm/external-dns-app/releases/tag/v2.14.0)

#### Added
- VerticalPodAutoscaler for automatically setting requests and limits depending on usage. Fixes OOM kills on huge clusters.



