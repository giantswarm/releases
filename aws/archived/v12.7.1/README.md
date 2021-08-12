# :zap: Giant Swarm Release v12.7.1 for AWS :zap:

This release fixes an issue that causes `ImagePullBackOff` errors when new nodes are becoming ready.

## Change details

### app-operator [2.7.0](https://github.com/giantswarm/app-operator/releases/tag/v2.7.0)

#### Added
- Secure the webhook with token value from control plane catalog.
- Adding webhook URL as annotation into chart CRs.
- Added Status update endpoint.
- Watch secrets referenced in app CRs to reduce latency when applying config
changes.
- Create appcatalogentry CRs for public app catalogs.
- Watch configmaps referenced in app CRs to reduce latency when applying config
changes.

#### Changed
- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.

#### Fixed
- Use resourceVersion of configmap for comparison instead of listing option.

### aws-operator [9.3.1-fix](https://github.com/giantswarm/aws-operator/releases/tag/v9.3.1-fix)

#### Changed
- Remove explicit registry pull limits defaulting to less restrictive upstream settings.

### chart-operator [2.5.1](https://github.com/giantswarm/chart-operator/releases/tag/v2.5.1)

#### Added
- Validate the cache in helmclient to avoid state requests when pulling tarballs.
- Call status webhook with token values.
- Call status webhook when webhook annotation is present.

#### Fixed
- Fix comparison of last deployed and revision optional fields in status resource.
- Set memory limit and reduce requests.
- Update apiextensions to v3 and replace CAPI with Giant Swarm fork.

#### Removed
- Remove chartmigration resource as migration from chartconfig to chart CRs is
complete.

### cluster-operator [3.4.1](https://github.com/giantswarm/cluster-operator/releases/tag/v3.4.1)

#### Added
- Add functionality to template `catalog` into `apps` depending on `release` CR.

#### Changed
- Update `apiextensions`, `k8sclient`, and `operatorkit` dependencies.
- Update github workflows.

#### Fixed
-  Allow annotations from current app CR to remain.
