<details>
<summary>📦 Update Components and Apps</summary>


To update the release files in this PR, comment with `/update-release` and provide arguments directly to `devctl`.

**Examples:**

*   Update a single component:
    `/update-release --component flatcar@4152.2.3`

*   Update a single application:
    `/update-release --app aws-ebs-csi-driver@3.0.5`

*   Update multiple components and apps at once:
    `/update-release --component cluster-aws@4.0.2 --app karpenter-bundle@2.2.0 --app coredns@1.27.0`

*   Specify app dependencies using the `#` separator:
    `/update-release --app my-app@1.2.3@@dependency1#dependency2`

**Add descriptions:**

*   Update README.md with detailed description:
    `/update-readme "This release includes component updates and improvements."`

*   Update announcement.md with brief summary:
    `/update-announcement "Workload cluster release includes important updates."`

</details>

<details>
<summary>📌 Pin Component or App Versions</summary>


To pin a specific component or app version (preventing automatic bumps), add a comment with `/pin-version`.

**Pin duration:**
- **Default (no `--until` flag)**: Pins only for this release. Future releases will auto-update.
- **With `--until vX.Y.Z`**: Pins until the specified version. Example: `--until v32.0.0` pins for all releases < v32.0.0

**Examples:**

*   Pin component for this release only:
    `/pin-version --component flatcar@4152.2.3`

*   Pin app for this release only:
    `/pin-version --app cilium@1.2.2`

*   Pin component until a specific version:
    `/pin-version --component flatcar@4152.2.3 --until v32.0.0`

*   Pin app for all patch releases in this minor series:
    `/pin-version --app security-bundle@1.0.0 --until v31.2.0`

*   Add a reason for the pin:
    `/pin-version --app security-bundle@1.0.0 --until v32.0.0 --reason "Version 1.1.0 has known issues"`

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
