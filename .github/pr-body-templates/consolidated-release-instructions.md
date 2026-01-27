<details>
<summary>📦 Update Components and Apps</summary>


To update the release files in this PR, comment with `/update-release` and provide arguments directly to `devctl`. For consolidated releases, you **must** specify which provider you are updating.

Available providers: `aws`, `azure`, `cloud-director`, `eks`, `vsphere`.

**Examples:**

*   Update a single component for a specific provider:
    `/update-release --provider aws --component flatcar@4152.2.3`

*   Update multiple components and apps for a specific provider:
    `/update-release --provider azure --component cluster-azure@2.4.1 --app azuredisk-csi-driver@1.32.9`

*   Specify app dependencies using the `#` separator:
    `/update-release --provider aws --app my-app@1.2.3@@dependency1#dependency2`

**Add descriptions:**

*   Update README.md for specific provider:
    `/update-readme --provider aws "This release includes component updates and improvements."`

*   Update announcement.md for specific provider:
    `/update-announcement --provider aws "Workload cluster release includes important updates."`

*   Update all providers (auto-detected):
    `/update-readme "This release brings component updates across all supported providers."`

</details>

<details>
<summary>📌 Pin Component or App Versions</summary>


To pin a specific component or app version (preventing automatic bumps), add a comment with `/pin-version`. 

The workflow automatically detects which providers use the component/app. Use `--provider` to pin for a specific provider only (e.g., when a shared component has issues on one provider but not others).

**Pin duration:**
- **Default (no `--until` flag)**: Pins only for this release. Future releases will auto-update.
- **With `--until vX.Y.Z`**: Pins until the specified version. Example: `--until v32.0.0` pins for all releases < v32.0.0

**Examples:**

*   Pin provider-specific component (auto-detects AWS):
    `/pin-version --component cluster-aws@6.2.0`

*   Pin shared component for all providers that use it:
    `/pin-version --component flatcar@4152.2.3`

*   Pin shared component for one provider only:
    `/pin-version --provider azure --component flatcar@4152.2.3`

*   Pin until a specific version:
    `/pin-version --component cluster-aws@6.2.0 --until v33.0.0`

*   Add a reason:
    `/pin-version --component cluster-aws@6.2.0 --until v33.0.0 --reason "Version 6.4.0 has known issues"`

</details>

<details>
<summary>🧪 Trigger E2E Tests</summary>


To trigger the E2E test for each new Release added in this PR, add a comment with:

`/run releases-test-suites`

If your release is a new _patch_ or _minor_ release for an older major release, specify the previous release for upgrade tests:

`/run releases-test-suites PREVIOUS_RELEASE=25.1.2`

You can also limit which tests are run:

`/run releases-test-suites TARGET_SUITES=./providers/capa/standard`

For conformance tests:

`/run conformance-tests PROVIDER=capa RELEASE_VERSION=29.1.0`

</details>

<details>
<summary>🏗️ Trigger MC Creation Tests</summary>


To validate that this release doesn't break Management Cluster creation, run:

`/run generate-mc PROVIDER=capa`

A free installation from `gawstest1-10` is auto-selected. You can override with `INSTALLATION=gawstest1` if needed.

**Note:** Auto-selection only works for CAPA. For other providers, specify the installation explicitly.

This will:
- Create an ephemeral MC using the release from your PR branch (auto-detected)
- Post a GitHub check status to this PR
- Automatically clean up on completion

**Note:** This test is separate from the WC E2E tests and validates that MC creation works with the new release.

</details>
