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
- [ ] Release uses latest supported version of all default apps
- [ ] MC creation test passed (see "Triggering MC Creation Tests" below)

### Triggering E2E tests

To trigger the E2E test for each new Release added in this PR, add a comment with the following:

`/run releases-test-suites`

If your release is a new _patch_ release for an older major release, you need to specify the previous release for use in upgrade tests, for example for `25.1.2` (exists) to `25.1.3` (added in your PR):

`/run releases-test-suites PREVIOUS_RELEASE=25.1.2`

If your release is a new _minor_ release for an older major release, e.g. `25.3.0` (exists) to `25.4.0` (added in your PR), it works the same way:

`/run releases-test-suites PREVIOUS_RELEASE=25.3.0`

You can also limit which tests are run:

`/run releases-test-suites TARGET_SUITES=./providers/capa/standard`

If you want to trigger conformance tests, you can do so by adding a comment similar to the following:

`/run conformance-tests PROVIDER=capa RELEASE_VERSION=29.1.0`

For more details see the [README.md](/README.md#running-tests-against-prs).

### Triggering MC Creation Tests

To prevent releases from breaking Management Cluster creation, run these tests before merging. The test recreates MCs using the release from your PR branch (auto-detected from the PR's commit SHA) to catch issues early.

**Test a specific provider:**

`/run generate-mc INSTALLATION=<installation> PROVIDER=<provider>`

Examples:
- `/run generate-mc INSTALLATION=goten PROVIDER=capa`
- `/run generate-mc INSTALLATION=goose PROVIDER=capz`
- `/run generate-mc INSTALLATION=goshawk PROVIDER=cloud-director`
- `/run generate-mc INSTALLATION=gmc PROVIDER=vsphere`

### Test all providers

`/run generate-mc-all`

**NOTE:** This command tests recreating `capa/goten`, `capz/goose`, `cloud-director/goshawk` and `vsphere/gmc` on head commit of PR.

**Optional parameters:**
- `MC_BOOTSTRAP_REF` - Git ref for mc-bootstrap repo (default: `main`). Use this to test with a specific mc-bootstrap branch.

This will:
- Recreate an MC using the release from your PR branch
- Post a GitHub check status to this PR
- Automatically clean up on completion

**Note:** This test is separate from the WC E2E tests and validates that MC creation works with the new release.

For more details see the [CAPI release drafting documentation](https://intranet.giantswarm.io/docs/product/releases/capi/capi-release-drafting/#mc-creation-tests).
