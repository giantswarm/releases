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

// NewClientWithGitHubClient creates a releases Client with the specified GitHub client.
func NewClientWithGitHubClient(gitHubClient *github.Client) (*Client, error) {
	if gitHubClient == nil {
		return nil, microerror.Maskf(InvalidConfigError, "gitHubClient must be specified")
	}
	return &Client{
		gitHubClient: gitHubClient,
	}, nil
}

// NewClientWithGitHubToken creates a releases Client that is internally using a GitHub client which is created with the specified
// GitHub token.
func NewClientWithGitHubToken(gitHubToken string) *Client {
	gitHubClient := github.NewClient(nil)
	if gitHubToken != "" {
		gitHubClient = gitHubClient.WithAuthToken(gitHubToken)
	}

	return &Client{
		gitHubClient,
	}
}

// GetRelease returns a release with the specified release version for the specified provider. Specified release version
// can contain the 'v' prefix, but it doesn't have to.
func (c *Client) GetRelease(ctx context.Context, provider Provider, releaseVersion string) (*Release, error) {
	// First we get GitHub release for the specified provider and release version.
	releaseVersion = strings.TrimPrefix(releaseVersion, "v")
	releaseTag := fmt.Sprintf("%s/v%s", provider, releaseVersion)
	gitHubRelease, _, err := c.gitHubClient.Repositories.GetReleaseByTag(ctx, GiantSwarmGitHubOrg, ReleasesRepo, releaseTag)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	release, err := c.getReleaseResourceFromGitHubRelease(gitHubRelease)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return release, nil
}

// GetLatestRelease returns the latest release for the specified provider.
func (c *Client) GetLatestRelease(ctx context.Context, provider Provider) (*Release, error) {
	const releasesPerRequest = 30
	providerTagPrefix := fmt.Sprintf("%s/", provider)

	var releaseList []*github.RepositoryRelease
	var response *github.Response
	var err error
	hasMorePages := func(page int) bool {
		return page != 0
	}

	for page := 1; hasMorePages(page); page = response.NextPage {
		releaseList, response, err = c.gitHubClient.Repositories.ListReleases(ctx, GiantSwarmGitHubOrg, ReleasesRepo, &github.ListOptions{
			Page:    page,
			PerPage: releasesPerRequest,
		})
		if err != nil {
			return nil, microerror.Mask(err)
		}
		for _, gitHubRelease := range releaseList {
			releaseTag := gitHubRelease.GetTagName()
			if strings.HasPrefix(releaseTag, providerTagPrefix) {
				release, err := c.getReleaseResourceFromGitHubRelease(gitHubRelease)
				if err != nil {
					return nil, microerror.Mask(err)
				}
				return release, nil
			}
		}
	}

	return nil, microerror.Maskf(ReleaseNotFoundError, "Did not found any release for provider '%s'", provider)
}

func (c *Client) getReleaseResourceFromGitHubRelease(gitHubRelease *github.RepositoryRelease) (*Release, error) {
	// We get the GitHub release asset that contains release.yaml manifest file.
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
	var err error
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