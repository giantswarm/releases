package sdk

import (
	"context"
	"net/http/httptest"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v62/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

var _ = Describe("Builder", func() {
	var server *httptest.Server
	var releasesBuilder *Builder

	BeforeEach(func() {
		server = createAndStartTestServer()

		// Now create Release client
		httpTestClient := server.Client()
		gitHubClient, err := github.NewClient(httpTestClient).WithEnterpriseURLs(gitHubApiBaseName, "")
		Expect(err).NotTo(HaveOccurred())

		releasesClient, err := NewClientWithGitHubClient(gitHubClient)
		Expect(err).NotTo(HaveOccurred())

		releasesBuilder, err = NewBuilder(releasesClient, ProviderAws, "")
		Expect(err).NotTo(HaveOccurred())
	})

	It("Overrides cluster app and creates a new patch release with pre-release suffix test", func() {
		// set a cluster app override where the pre-release is different
		const preReleasePrefix = "test"
		releasesBuilder = releasesBuilder.
			WithClusterApp("0.76.1-efdaba18e6e9866d62f3214f3d898b21c21e8b48", "").
			WithPreReleasePrefix(preReleasePrefix)

		// build the new release
		const expectedPreReleaseLength = len(preReleasePrefix) + 1 + preReleaseHashLength // <prefix>.<pre-release hash>
		newRelease, err := releasesBuilder.Build(context.Background())
		Expect(err).NotTo(HaveOccurred())

		// get the new release version string
		newReleaseVersionString, err := newRelease.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		newReleaseVersion, err := semver.NewVersion(newReleaseVersionString)
		Expect(err).NotTo(HaveOccurred())

		// check the new release version, expected is 25.1.1-<pre-release>
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Prerelease()).To(HavePrefix("%s.", preReleasePrefix))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(expectedPreReleaseLength))
	})

	It("Overrides cluster app and creates a new minor release with pre-release", func() {
		// set a cluster app override where the pre-release is different
		releasesBuilder = releasesBuilder.
			WithClusterApp("0.77.0", "")

		// build the new release
		newRelease, err := releasesBuilder.Build(context.Background())
		Expect(err).NotTo(HaveOccurred())

		// get the new release version string
		newReleaseVersionString, err := newRelease.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		newReleaseVersion, err := semver.NewVersion(newReleaseVersionString)
		Expect(err).NotTo(HaveOccurred())

		// Check the new release version.
		//
		// Since the builder will use the latest release, which is v25.1.0-demo.0, it is expected that the custom
		// release version has a minor bump (because we are building a newer release with a new cluster-aws minor
		// version), so the new release version is 25.2.0-<random 10-char pre-release suffix>.
		//
		// New release also has a pre-release because the base release is a pre-release.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(2)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(0)))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(10))
	})

	It("Overrides cluster app and creates a new minor release with random pre-release", func() {
		// set a cluster app override where the pre-release is different
		const preReleaseLength = 10
		releasesBuilder = releasesBuilder.
			WithClusterApp("0.77.0", "").
			WithRandomPreRelease(preReleaseLength)

		// build the new release
		newRelease, err := releasesBuilder.Build(context.Background())
		Expect(err).NotTo(HaveOccurred())

		// get the new release version string
		newReleaseVersionString, err := newRelease.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		newReleaseVersion, err := semver.NewVersion(newReleaseVersionString)
		Expect(err).NotTo(HaveOccurred())

		// Check the new release version.
		//
		// Since the builder will use the latest release, which is v25.1.0-demo.0, it is expected that the custom
		// release version has a minor bump (because we are building a newer release with a new cluster-aws minor
		// version), so the new release version is 25.2.0-<random 10-char pre-release suffix>.
		//
		// New release also has a pre-release because the base release is a pre-release.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(2)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(0)))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(preReleaseLength))
	})

	It("Creates a new release with random pre-release", func() {
		// set a cluster app override where the pre-release is different
		const preReleaseLength = 10
		releasesBuilder = releasesBuilder.WithRandomPreRelease(preReleaseLength)

		// build the new release
		newRelease, err := releasesBuilder.Build(context.Background())
		Expect(err).NotTo(HaveOccurred())

		// get the new release version string
		newReleaseVersionString, err := newRelease.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		newReleaseVersion, err := semver.NewVersion(newReleaseVersionString)
		Expect(err).NotTo(HaveOccurred())

		// Check the new release version. Since the builder will use the latest release, which is v25.1.0-demo.0,
		// expected is custom release version has a patch bump (because we are building a newer release),
		// so the new release version is 25.1.1-<random 10-char pre-release suffix>.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(10))
	})

	AfterEach(func() {
		server.Close()
	})
})
