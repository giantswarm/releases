<!--
If this is a PR with details for a new release, please review the [Workload Cluster Releases Board](https://github.com/orgs/giantswarm/projects/365):

- If there's an issue for this release open in the "Planned" column without a team assigned, please use it and try to include requested changes in your release (details of this process can be found [here](https://intranet.giantswarm.io/docs/product/releases/requesting-changes-in-next-platform-release)).
- Otherwise create an appropriate issue for your release in https://github.com/giantswarm/roadmap and add it to the releases board.

Ping @sig-product for review of release notes.
--->

### Checklist

- [ ] Roadmap issue created
- [ ] Release uses latest stable Flatcar
- [ ] Release uses latest Kubernetes patch version

### Triggering E2E tests

To trigger the E2E test for each new Release added in this PR, add a comment with the following:

`/run releases-test-suites`

If you want to trigger conformance tests, you can do so by adding a comment similar to the following:

`/run conformance-tests PROVIDER=capa RELEASE_VERSION=29.1.0`

For more details see the [README.md](/README.md#running-tests-against-prs).
