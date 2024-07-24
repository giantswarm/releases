# :zap: Giant Swarm Release v25.0.0 for CAPZ :zap:

We are happy to announce our first release for Cluster API for Azure (CAPZ) that uses the new releases framework.

## Migration to new releases flow

In order to consume the new flow, the following two fields need to be manually adapted:

* In ConfigMap `<cluster name>-userconfig` set `.Values.global.release` to the release version, e.g. `25.0.0`.
* In App `<cluster name>` set the `version` to an empty string.
