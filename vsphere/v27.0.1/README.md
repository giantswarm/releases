# :zap: Giant Swarm Release v27.0.1 for vSphere :zap:

We are happy to announce the first release for vSphere that uses the new release framework.

## Migration to new releases flow

In order to consume the new flow, the following two fields need to be manually adapted:

* In ConfigMap `<cluster name>-userconfig` set `.Values.global.release` to the release version, e.g. `27.0.1`.
* In App `<cluster name>` remove the `spec.version` field. In case of GitOps, Flux might complain that the app manifest is invalid as the `spec.version` field is mandatory. In that case, edit the live App CR and set `spec.version` to an empty string. That will unblock Flux and allow it reconcile successfully.

And if you want to use `kubectl-gs` to create a cluster, you'd need to now specify the release version, e.g.:

```bash
kubectl-gs template cluster --provider vsphere --organization my_org --name cluster_name -vsphere-network-name network_name --release v27.0.1
```
