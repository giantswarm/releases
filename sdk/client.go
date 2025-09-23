package sdk

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path"
	"slices"
	"strings"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/microerror"
	"github.com/google/go-github/v75/github"
	"sigs.k8s.io/yaml"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
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

// GetNewReleasesForGitReference returns newly added releases for the specified provider and from the specified git
// reference.
//
// Currently, the supported providers are "aws" and "azure". Git reference can be any commit, branch or tag.
func (c *Client) GetNewReleasesForGitReference(ctx context.Context, provider Provider, gitReference string) ([]Release, error) {
	providerDirectory, err := getProviderDirectory(provider)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	_, directoryContentDefault, _, err := c.gitHubClient.Repositories.GetContents(ctx, GiantSwarmGitHubOrg, ReleasesRepo, providerDirectory, nil)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	directoryContains := func(items []*github.RepositoryContent, dirName string) bool {
		for _, item := range items {
			if item != nil && item.GetType() == "dir" && item.GetName() == dirName {
				return true
			}
		}
		return false
	}

	opts := &github.RepositoryContentGetOptions{
		Ref: gitReference,
	}
	_, directoryContentGitRef, _, err := c.gitHubClient.Repositories.GetContents(ctx, GiantSwarmGitHubOrg, ReleasesRepo, providerDirectory, opts)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var releases []Release
	for _, item := range directoryContentGitRef {
		if item == nil || item.GetType() != "dir" {
			continue
		}
		releaseDirectoryName := item.GetName()
		if releaseDirectoryName == "" || releaseDirectoryName == archivedDirectory {
			continue
		}
		if directoryContains(directoryContentDefault, releaseDirectoryName) {
			// this release is already included in the default branch
			continue
		}
		_, err = semver.NewVersion(releaseDirectoryName)
		if err != nil {
			// directory name is not a valid semver version, so treating as unknown directory and skipping
			continue
		}
		release, err := c.GetReleaseForGitReference(ctx, provider, releaseDirectoryName, gitReference)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		releases = append(releases, *release)
	}

	return releases, nil
}

// GetReleasesForGitReference returns all releases for the specified provider and from the specified git reference.
//
// Currently, the supported providers are "aws" and "azure". Git reference can be any commit, branch or tag.
func (c *Client) GetReleasesForGitReference(ctx context.Context, provider Provider, gitReference string) ([]Release, error) {
	providerDirectory, err := getProviderDirectory(provider)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	opts := &github.RepositoryContentGetOptions{
		Ref: gitReference,
	}
	_, directoryContent, _, err := c.gitHubClient.Repositories.GetContents(ctx, GiantSwarmGitHubOrg, ReleasesRepo, providerDirectory, opts)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var releases []Release
	for _, item := range directoryContent {
		if item == nil || item.GetType() != "dir" {
			continue
		}
		releaseDirectoryName := item.GetName()
		if releaseDirectoryName == "" || releaseDirectoryName == archivedDirectory {
			continue
		}
		_, err = semver.NewVersion(releaseDirectoryName)
		if err != nil {
			// directory name is not a valid semver version, so treating as unknown directory and skipping
			continue
		}
		release, err := c.GetReleaseForGitReference(ctx, provider, releaseDirectoryName, gitReference)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		releases = append(releases, *release)
	}

	return releases, nil
}

// GetReleaseForGitReference returns a release with the specified release version for the specified provider and from
// the specified git reference.
//
// Currently, the supported providers are "aws" and "azure". Release version can contain the 'v' prefix, but it doesn't have to.
// Git reference can be any commit, branch or tag.
func (c *Client) GetReleaseForGitReference(ctx context.Context, provider Provider, releaseVersion, gitReference string) (*Release, error) {
	// First we get GitHub release for the specified provider and release version.
	releaseVersion = strings.TrimPrefix(releaseVersion, "v")
	providerDirectory, err := getProviderDirectory(provider)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	releaseManifestPath := path.Join(providerDirectory, fmt.Sprintf("v%s", releaseVersion), ReleaseManifestFileName)
	opts := &github.RepositoryContentGetOptions{
		Ref: gitReference,
	}
	fileContentObject, _, _, err := c.gitHubClient.Repositories.GetContents(ctx, GiantSwarmGitHubOrg, ReleasesRepo, releaseManifestPath, opts)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	fileContent, err := fileContentObject.GetContent()
	if err != nil {
		return nil, microerror.Mask(err)
	}
	if fileContent == "" {
		return nil, microerror.Maskf(MissingReleaseManifestError, "release manifest contents not found at path '%s' for git reference '%s'", releaseManifestPath, gitReference)
	}

	release := &Release{}
	err = yaml.Unmarshal([]byte(fileContent), release)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return release, nil
}

// GetLatestRelease returns the latest release for the specified provider.
func (c *Client) GetLatestRelease(ctx context.Context, provider Provider) (*Release, error) {
	allActiveReleases, err := c.GetReleasesForGitReference(ctx, provider, ReleasesRepoDefaultBranch)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	if len(allActiveReleases) == 0 {
		return nil, microerror.Maskf(ReleaseNotFoundError, "Did not found any release for provider '%s'", provider)
	}

	var sortErr error
	slices.SortFunc[[]Release](allActiveReleases, func(r1, r2 Release) int {
		r1VersionString, err := r1.GetVersion()
		if err != nil {
			sortErr = err
			return 0
		}
		r1Version, err := semver.NewVersion(r1VersionString)
		if err != nil {
			sortErr = err
			return 0
		}
		r2VersionString, err := r2.GetVersion()
		if err != nil {
			sortErr = err
			return 0
		}
		r2Version, err := semver.NewVersion(r2VersionString)
		if err != nil {
			sortErr = err
			return 0
		}
		return r1Version.Compare(r2Version)
	})
	if sortErr != nil {
		return nil, microerror.Mask(sortErr)
	}

	latestRelease := allActiveReleases[len(allActiveReleases)-1]
	return &latestRelease, nil
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
