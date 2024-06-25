package sdk

const (
	// GiantSwarmGitHubOrg is the name of the Giant Swarm GitHub organization.
	GiantSwarmGitHubOrg = "giantswarm"

	// ReleasesRepo is the name of the Giant Swarm's releases repository where the releases are published.
	ReleasesRepo = "releases"

	// ReleasesRepoDefaultBranch is the name of the releases repository default branch.
	ReleasesRepoDefaultBranch = "master"

	// ReleaseManifestFileName is the name of the YAML file which contains a Release resource. This file can be found
	// in the releases repository in <provider>/<version>/release.yaml, and it is also attached in GitHub release in
	// the giantswarm/releases repository.
	ReleaseManifestFileName = "release.yaml"
)
