package sdk

import (
	"context"
	"net/http/httptest"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/google/go-github/v85/github"
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
			WithClusterApp("1.0.0-efdaba18e6e9866d62f3214f3d898b21c21e8b48", "").
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

		// check the new release version, expected is 25.0.1-<pre-release>
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(0)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Prerelease()).To(HavePrefix("%s.", preReleasePrefix))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(expectedPreReleaseLength))
	})

	It("Overrides cluster app and creates a new minor release with pre-release", func() {
		// set a cluster app override where the pre-release is different
		releasesBuilder = releasesBuilder.
			WithClusterApp("1.1.0", "")

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
		// Since the builder will use the latest release, which is v25.0.0, it is expected that the custom
		// release version has a minor bump (because we are building a newer release with a new cluster-aws minor
		// version), so the new release version is 25.1.0.
		//
		// New release also doesn't have a pre-release because the base release doesn't have it.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(0)))
		Expect(newReleaseVersion.Prerelease()).To(BeEmpty())
	})

	It("Overrides cluster app and creates a new minor release with random pre-release", func() {
		// set a cluster app override where the pre-release is different
		const preReleaseLength = 10
		releasesBuilder = releasesBuilder.
			WithClusterApp("1.1.0", "").
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
		// Since the builder will use the latest release, which is v25.0.0, it is expected that the custom
		// release version has a minor bump (because we are building a newer release with a new cluster-aws minor
		// version), so the new release version is 25.1.0-<random 10-char pre-release suffix>.
		//
		// New release also has a random pre-release with the specified length.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(1)))
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

		// Check the new release version. Since the builder will use the latest release, which is v25.0.0,
		// expected is custom release version has a patch bump (because we are building a newer release),
		// so the new release version is 25.0.1-<random 10-char pre-release suffix>.
		//
		// New release also has a random pre-release with the specified length.
		Expect(newReleaseVersion.Major()).To(Equal(uint64(25)))
		Expect(newReleaseVersion.Minor()).To(Equal(uint64(0)))
		Expect(newReleaseVersion.Patch()).To(Equal(uint64(1)))
		Expect(newReleaseVersion.Prerelease()).To(HaveLen(10))
	})

	It("Creates a new release with release date set to now", func() {
		// time before the release is created
		timeBefore := time.Now()

		// build the new release
		newRelease, err := releasesBuilder.Build(context.Background())
		Expect(err).NotTo(HaveOccurred())

		// time after the release is created
		timeAfter := time.Now()

		// Check the new release time
		Expect(newRelease.Spec.Date).NotTo(BeNil())
		Expect(newRelease.Spec.Date.Time).To(BeTemporally(">=", timeBefore))
		Expect(newRelease.Spec.Date.Time).To(BeTemporally("<=", timeAfter))
	})

	AfterEach(func() {
		server.Close()
	})
})

var _ = Describe("overrideApp", func() {
	It("preserves dependsOn when override has empty list", func() {
		release := &Release{
			Spec: ReleaseSpec{
				Apps: []ReleaseSpecApp{
					{
						Name:      "cert-exporter",
						Version:   "2.9.15",
						Catalog:   "default",
						DependsOn: []string{"kyverno-crds"},
					},
				},
			},
		}

		overrideApp(release, ReleaseSpecApp{
			Name:      "cert-exporter",
			Version:   "2.9.16",
			DependsOn: []string{},
		})

		Expect(release.Spec.Apps[0].Version).To(Equal("2.9.16"))
		Expect(release.Spec.Apps[0].DependsOn).To(Equal([]string{"kyverno-crds"}))
	})

	It("preserves dependsOn with multiple dependencies", func() {
		release := &Release{
			Spec: ReleaseSpec{
				Apps: []ReleaseSpecApp{
					{
						Name:      "security-bundle",
						Version:   "1.16.1",
						Catalog:   "giantswarm",
						DependsOn: []string{"prometheus-operator-crd", "kyverno-crds"},
					},
				},
			},
		}

		overrideApp(release, ReleaseSpecApp{
			Name:      "security-bundle",
			Version:   "1.17.0",
			DependsOn: []string{},
		})

		Expect(release.Spec.Apps[0].Version).To(Equal("1.17.0"))
		Expect(release.Spec.Apps[0].DependsOn).To(Equal([]string{"prometheus-operator-crd", "kyverno-crds"}))
	})

	It("allows overriding dependsOn when explicitly provided", func() {
		release := &Release{
			Spec: ReleaseSpec{
				Apps: []ReleaseSpecApp{
					{
						Name:      "cert-exporter",
						Version:   "2.9.15",
						Catalog:   "default",
						DependsOn: []string{"kyverno-crds"},
					},
				},
			},
		}

		overrideApp(release, ReleaseSpecApp{
			Name:      "cert-exporter",
			Version:   "2.9.16",
			DependsOn: []string{"new-dependency"},
		})

		Expect(release.Spec.Apps[0].Version).To(Equal("2.9.16"))
		Expect(release.Spec.Apps[0].DependsOn).To(Equal([]string{"new-dependency"}))
	})

	It("preserves catalog when override has empty catalog", func() {
		release := &Release{
			Spec: ReleaseSpec{
				Apps: []ReleaseSpecApp{
					{
						Name:    "coredns",
						Version: "1.29.1",
						Catalog: "default",
					},
				},
			},
		}

		overrideApp(release, ReleaseSpecApp{
			Name:    "coredns",
			Version: "1.29.2",
			Catalog: "",
		})

		Expect(release.Spec.Apps[0].Version).To(Equal("1.29.2"))
		Expect(release.Spec.Apps[0].Catalog).To(Equal("default"))
	})
})
