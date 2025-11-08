package sdk

import (
	"context"
	_ "embed"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/google/go-github/v78/github"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

const (
	gitHubApiHostNameAndPort = "localhost:8081"
	gitHubApiBaseName        = "http://" + gitHubApiHostNameAndPort + "/"
)

var (
	//go:embed testdata/capa/github_list_api_response.json
	listReleasesGitHubResponse string

	//go:embed testdata/capa/v25.0.0-demo.0/github_get_api_response.json
	v2500GitHubReleaseResponse string

	//go:embed testdata/capa/v25.0.0-demo.0/release.yaml
	v2500ReleaseManifest string

	//go:embed testdata/capa/v25.1.0-demo.0/release.yaml
	v2510ReleaseManifest string

	//go:embed testdata/capa/github_get_capa_for_git_default.json
	capaDirectoryGitDefaultResponse string

	//go:embed testdata/capa/github_get_capa_for_git_ref.json
	capaDirectoryGitReferenceResponse string

	//go:embed testdata/capa/v25.0.0/release_yaml_content_response.json
	v25ReleaseManifestContentResponse string

	//go:embed testdata/capa/v25.0.0-alpha.3/release_yaml_content_response.json
	v25Alpha3ReleaseManifestContentResponse string
)

var _ = Describe("Client", func() {
	var server *httptest.Server
	var releasesClient *Client

	BeforeEach(func() {
		server = createAndStartTestServer()

		// Now create Release client
		httpTestClient := server.Client()
		gitHubClient, err := github.NewClient(httpTestClient).WithEnterpriseURLs(gitHubApiBaseName, "")
		Expect(err).NotTo(HaveOccurred())

		releasesClient, err = NewClientWithGitHubClient(gitHubClient)
		Expect(err).NotTo(HaveOccurred())
	})

	It("Gets release for the specified provider and release version", func() {
		ctx := context.Background()
		const releaseVersion = "25.0.0-demo.0"
		release, err := releasesClient.GetRelease(ctx, ProviderAws, releaseVersion)
		Expect(err).NotTo(HaveOccurred())

		// Check resource name
		expectedReleaseResourceName := fmt.Sprintf("%s-%s", ProviderAws, releaseVersion)
		Expect(release.Name).To(Equal(expectedReleaseResourceName))

		// Check provider name
		resultProvider, err := release.GetProvider()
		Expect(err).NotTo(HaveOccurred())
		Expect(resultProvider).To(Equal(ProviderAws))

		// Check release version
		resultReleaseVersion, err := release.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		Expect(resultReleaseVersion).To(Equal(releaseVersion))

		// Check app name and version
		appName := "coredns"
		appVersion := "1.21.0"
		appSpec, ok := release.LookupAppSpec(appName)
		Expect(ok).To(BeTrue())
		Expect(appSpec.Name).To(Equal(appName))
		Expect(appSpec.Version).To(Equal(appVersion))

		// Check component version
		const componentName = "flatcar-variant"
		const componentVersion = "1.0.0"
		componentSpec, ok := release.LookupComponentSpec(componentName)
		Expect(ok).To(BeTrue())
		Expect(componentSpec.Name).To(Equal(componentName))
		Expect(componentSpec.Version).To(Equal(componentVersion))

		// Check cluster app version
		clusterAppVersion, err := release.GetClusterAppVersion()
		Expect(err).NotTo(HaveOccurred())
		const expectedClusterAppVersion = "0.76.1-b76af2c26f4224ffb0d718e940e232fac05c89a0"
		Expect(clusterAppVersion).To(Equal(expectedClusterAppVersion))

		// Check Kubernetes version
		kubernetesVersion, err := release.GetKubernetesVersion()
		Expect(err).NotTo(HaveOccurred())
		const expectedKubernetesVersion = "1.25.16"
		Expect(kubernetesVersion).To(Equal(expectedKubernetesVersion))

		// Check Flatcar version
		flatcarVersion, err := release.GetFlatcarVersion()
		Expect(err).NotTo(HaveOccurred())
		const expectedFlatcarVersion = "3815.2.2"
		Expect(flatcarVersion).To(Equal(expectedFlatcarVersion))
	})

	It("Gets the latest release for the specified provider", func() {
		ctx := context.Background()
		release, err := releasesClient.GetLatestRelease(ctx, ProviderAws)
		Expect(err).NotTo(HaveOccurred())

		// Check release version
		const expectedLatestReleaseVersion = "25.0.0"
		resultReleaseVersion, err := release.GetVersion()
		Expect(err).NotTo(HaveOccurred())
		Expect(resultReleaseVersion).To(Equal(expectedLatestReleaseVersion))
	})

	It("Gets all releases for the specified provider and git reference", func() {
		ctx := context.Background()
		releases, err := releasesClient.GetReleasesForGitReference(ctx, ProviderAws, "add-capa-v25.0.0")
		Expect(err).NotTo(HaveOccurred())

		expectedReleases := []string{
			"25.0.0-alpha.3",
			"25.0.0",
		}

		for i, release := range releases {
			resultReleaseVersion, err := release.GetVersion()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultReleaseVersion).To(Equal(expectedReleases[i]))
		}
	})

	It("Gets new releases for the specified provider and git reference", func() {
		ctx := context.Background()
		releases, err := releasesClient.GetNewReleasesForGitReference(ctx, ProviderAws, "add-capa-v25.0.0")
		Expect(err).NotTo(HaveOccurred())

		expectedReleases := []string{
			"25.0.0",
		}

		for i, release := range releases {
			resultReleaseVersion, err := release.GetVersion()
			Expect(err).NotTo(HaveOccurred())
			Expect(resultReleaseVersion).To(Equal(expectedReleases[i]))
		}
	})

	AfterEach(func() {
		server.Close()
	})
})

func createAndStartTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/v3/repos/giantswarm/releases/releases", func(rw http.ResponseWriter, req *http.Request) {
		// return GitHub response JSON
		_, err := rw.Write([]byte(listReleasesGitHubResponse))
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/api/v3/repos/giantswarm/releases/releases/tags/aws/v25.0.0-demo.0", func(rw http.ResponseWriter, req *http.Request) {
		// return GitHub response JSON
		_, err := rw.Write([]byte(v2500GitHubReleaseResponse))
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/giantswarm/releases/releases/download/aws/v25.0.0-demo.0/release.yaml", func(rw http.ResponseWriter, req *http.Request) {
		// return release.yaml manifest
		_, err := rw.Write([]byte(v2500ReleaseManifest))
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/giantswarm/releases/releases/download/aws/v25.1.0-demo.0/release.yaml", func(rw http.ResponseWriter, req *http.Request) {
		// return release.yaml manifest
		_, err := rw.Write([]byte(v2510ReleaseManifest))
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/api/v3/repos/giantswarm/releases/contents/capa", func(rw http.ResponseWriter, req *http.Request) {
		// return GitHub response JSON
		// Get the query parameters
		queryParams := req.URL.Query()
		// Retrieve specific query parameter
		ref := queryParams.Get("ref")
		var err error

		switch ref {
		case "add-capa-v25.0.0":
			fallthrough
		case ReleasesRepoDefaultBranch:
			_, err = rw.Write([]byte(capaDirectoryGitReferenceResponse))
		default:
			_, err = rw.Write([]byte(capaDirectoryGitDefaultResponse))
		}
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/api/v3/repos/giantswarm/releases/contents/capa/v25.0.0/release.yaml", func(rw http.ResponseWriter, req *http.Request) {
		// return GitHub response JSON
		_, err := rw.Write([]byte(v25ReleaseManifestContentResponse))
		Expect(err).NotTo(HaveOccurred())
	})
	mux.HandleFunc("/api/v3/repos/giantswarm/releases/contents/capa/v25.0.0-alpha.3/release.yaml", func(rw http.ResponseWriter, req *http.Request) {
		// return GitHub response JSON
		_, err := rw.Write([]byte(v25Alpha3ReleaseManifestContentResponse))
		Expect(err).NotTo(HaveOccurred())
	})

	// create test server
	server := httptest.NewUnstartedServer(mux)

	// Add listeners on custom hostname and port
	listener, err := net.Listen("tcp", gitHubApiHostNameAndPort)
	Expect(err).NotTo(HaveOccurred())
	err = server.Listener.Close()
	Expect(err).NotTo(HaveOccurred())
	server.Listener = listener
	server.Start()
	return server
}
