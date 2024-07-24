# :zap: Giant Swarm Release v25.0.0 for Azure :zap:

We are happy to announce the first release for Azure that uses the new release framework.

## Migration to new releases flow

In order to consume the new flow, the following two fields need to be manually adapted:

* In ConfigMap `<cluster name>-userconfig` set `.Values.global.release` to the release version, e.g. `25.0.0`.
* In App `<cluster name>` remove the `spec.version` field. In case of GitOps, Flux might complain that the app manifest is invalid as the `spec.version` field is mandatory. In that case, edit the live App CR and set `spec.version` to an empty string. That will unblock Flux and allow it reconcile successfully.

And if you want to use `kubectl-gs` to create a cluster, you'd need to now specify the release version, e.g.:

```bash
kubectl-gs template cluster --provider capz --organization my-org --name cluster_name --region westeurope --azure-subscription-id AZURE_ID --release 25.0.0
```
