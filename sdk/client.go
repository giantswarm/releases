package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sigs.k8s.io/yaml"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"

	"github.com/giantswarm/microerror"
	"github.com/google/go-github/v62/github"
)

type Client struct {
	gitHubClient *github.Client
}

func NewClient() *Client {
	return &Client{
		gitHubClient: github.NewClient(nil),
	}
}

// NewClientWithHttpClient creates a new instance of Client with a specified http.Client.
// It returns a pointer to Client and an error. If the httpClient is nil, it returns
// an error of type InvalidConfigError indicating that gitHubClient must be specified.
// The function initializes gitHubClient using the github.NewClient method and returns it
// along with the error as a pointer to Client.
func NewClientWithHttpClient(httpClient *http.Client) (*Client, error) {
	if httpClient == nil {
		return nil, microerror.Maskf(InvalidConfigError, "gitHubClient must be specified")
	}
	gitHubClient := github.NewClient(httpClient)

	return &Client{
		gitHubClient,
	}, nil
}

func NewClientWithGitHubClient(gitHubClient *github.Client) (*Client, error) {
	if gitHubClient == nil {
		return nil, microerror.Maskf(InvalidConfigError, "gitHubClient must be specified")
	}
	return &Client{
		gitHubClient: gitHubClient,
	}, nil
}

func NewClientWithGitHubToken(gitHubToken string) *Client {
	gitHubClient := github.NewClient(nil)
	if gitHubToken != "" {
		gitHubClient = gitHubClient.WithAuthToken(gitHubToken)
	}

	return &Client{
		gitHubClient,
	}
}

func (c *Client) GetPublishedRelease(ctx context.Context, provider Provider, releaseVersion string) (*Release, error) {
	// First we get GitHub release for the specified provider and release version.
	releaseVersion = strings.TrimPrefix(releaseVersion, "v")
	releaseTag := fmt.Sprintf("%s/v%s", provider, releaseVersion)
	gitHubRelease, _, err := c.gitHubClient.Repositories.GetReleaseByTag(ctx, GiantSwarmGitHubOrg, ReleasesRepo, releaseTag)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	// Then we get the GitHub release asset that contains release.yaml manifest file.
	var releaseManifestUrl string
	for _, asset := range gitHubRelease.Assets {
		if asset.GetName() == ReleaseManifestFileName {
			releaseManifestUrl = asset.GetBrowserDownloadURL()
			break
		}
	}
	if releaseManifestUrl == "" {
		return nil, microerror.Maskf(MissingReleaseManifestError, "Release asset with release.yaml manifest file not found")
	}

	// Now we download the release manifest by using the GitHub asset URL.
	var response *http.Response
	response, err = c.gitHubClient.Client().Get(releaseManifestUrl)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	defer func() {
		closeErr := response.Body.Close()
		if err == nil && closeErr != nil {
			err = closeErr
		}
	}()
	var releaseManifestYamlContent []byte
	releaseManifestYamlContent, err = io.ReadAll(response.Body)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	// Finally, we deserialize the downloaded release manifest.
	release := &Release{}
	err = yaml.Unmarshal(releaseManifestYamlContent, release)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return release, nil
}
