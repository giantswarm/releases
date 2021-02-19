# :zap: This week in provider-independent announcements (February 19, 2021) :zap:

For a comprehensive and always-updated view of all new features and changes for apps, workload clusters, UI, and documentation, please check out [changes and releases](https://docs.giantswarm.io/changes/). Below are the highlights.

## Managed apps

- [KEDA](https://github.com/giantswarm/keda-app) (Kubernetes Event-Driven Autoscaling) is now available as an app in the Playground Catalog. Some customers report significant resource optimizations and cost savings from using KEDA.

- [Fluent log shipping v0.5.5](https://docs.giantswarm.io/changes/playground-apps/fluent-logshipping-app/v0.5.5/) was fixed to run using AWS Cloudwatch and S3. We also added the ability to use KIAM for Cloudwatch and updated the fluentd image.

## Web interface

- The Apps section has been re-designed completely to provide access to all apps quicker, independent of the catalog they belong to.

## Documentation

- We added documentation for [spot VMs on Azure](https://docs.giantswarm.io/advanced/spot-instances/azure/), which is supported as of Azure release v14.1.0.

- Also new is a general documentation page about [management clusters](https://docs.giantswarm.io/general/management-clusters/) which also details the update cycles.

- Our Management API CRD docs now includes annotations. See [AWSMachineDeployment](https://docs.giantswarm.io/ui-api/management-api/crd/awsmachinedeployments.infrastructure.giantswarm.io/#annotation-details-v1alpha2) for an example.

---
Please let <!subteam^S0GSG846L|Product Owners> know if you have any feedback or questions by replying to this announcement in a thread.
