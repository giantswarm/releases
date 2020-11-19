## :zap:  Tenant Cluster Release 12.4.0 for AWS :zap:

This release upgrades cert-manager-app to `v2.3.0` which brings new patch version `v1.0.2` of the upstream project. 

## Change details

### cert-manager [2.3.0](https://github.com/giantswarm/cert-manager-app/releases/tag/v2.3.0)

* New Cert-Manager `v1.0.2` fixes the errors from `kubectl` invocation.
* Fix `hook-delete-policy` to delete hook resources to make upgrades reliable
* This version add support for Route53 using Kiam annotation.

### external-dns [1.5.0](https://github.com/giantswarm/external-dns-app/releases/tag/v1.5.0)

* New External-DNS upstream version `v0.7.4` with several bugfixes.
