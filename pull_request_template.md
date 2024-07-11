
<!--
If this is a PR with details for new release please review [Workload Cluster Releases Board](https://github.com/orgs/giantswarm/projects/365)
- if there's an issue for this release open in "Planned" column without team assigned, please use it and try to include requested changes in your release (details of this process can be found [here](https://intranet.giantswarm.io/docs/product/releases/requesting-changes-in-next-platform-release/))
- otherwise create an appropriate ticket for your release in https://github.com/giantswarm/roadmap and add it to the releases board

Ping @sig-product for review of release notes.
--->

### Checklist
- [ ] Roadmap issue created
- [ ] Release uses latest stable Flatcar
- [ ] Release uses latest Kubernetes patch version

### Triggering e2e tests

To trigger the E2E test for each new Release added in this PR add a comment with the following:

`/run releases-test-suites`

If you want to trigger conformance tests you can do so by adding a comment similar to the following:

`/run conformance-tests RELEASE_VERSION=v28.0.0 PROVIDER=capa`

For more details see the [README.md](/README.md#running-tests-against-prs)
