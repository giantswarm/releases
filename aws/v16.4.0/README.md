# :zap: Giant Swarm Release v16.4.0 for AWS :zap:

An release that is enabling an encryption key rotation for kubernetes secret encryption in etcd. Also updating Calico to latest version.

## Change details


### aws-operator [10.14.0](https://github.com/giantswarm/aws-operator/releases/tag/v10.14.0)

#### Changed
- Changes to EncryptionConfig in order to fully work with `encryption-provider-operator`.



### cluster-operator [3.13.0](https://github.com/giantswarm/cluster-operator/releases/tag/v3.13.0)

- Removed encryption key creation. Encryption keys will be managed by `encryption-provider-operator`.


### calico [3.21.3](https://github.com/projectcalico/calico/releases/tag/v3.21.3)
