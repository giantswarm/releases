# :zap: Giant Swarm Release v16.1.1 for Azure :zap:

This is a bug fix release to handle a problem during upgrades we found in release 16.1.0.

> **_Warning:_** The `flatcar` image included in this release enabled `cgroups v2` by default. Java applications can crash if they are using older JDK versions because they are unable to identify the memory limits from `cgroups v2`. This has been fixed on newer JDK versions. Please make sure that your applications are using the latest JDK 15, JDK 16 or JDK 17 before upgrading to this release. Upstream issue: https://bugs.openjdk.java.net/browse/JDK-8230305 .

## Change details

### azure-operator [5.12.0](https://github.com/giantswarm/azure-operator/releases/tag/v5.12.0)

#### Added

- Deal with AzureClusterConfig CR to avoid cluster operator conflict.


