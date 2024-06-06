package sdk

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/microerror"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

const (
	preReleaseHashLength = 10
)

type Builder struct {
	client      *Client
	provider    Provider
	baseRelease string

	preReleasePrefix string
	clusterApp       *ReleaseSpecComponent
	apps             []ReleaseSpecApp
}

func NewBuilder(client *Client, provider Provider, baseRelease string) (*Builder, error) {
	if client == nil {
		return nil, microerror.Maskf(InvalidConfigError, "client must not be empty")
	}
	if !IsProviderSupported(provider) {
		return nil, microerror.Maskf(UnsupportedProviderError, "provider `%s` is not supported", provider)
	}

	return &Builder{
		client:      client,
		provider:    provider,
		baseRelease: baseRelease,
	}, nil
}

func (b *Builder) WithClusterApp(version, catalog string) *Builder {
	b.clusterApp = &ReleaseSpecComponent{
		Name:    fmt.Sprintf("cluster-%s", b.provider),
		Version: version,
		Catalog: catalog,
	}
	return b
}

func (b *Builder) WithApp(name, version, catalog string, dependsOn []string) *Builder {
	app := ReleaseSpecApp{
		Name:      name,
		Version:   version,
		Catalog:   catalog,
		DependsOn: dependsOn,
	}
	b.apps = append(b.apps, app)
	return b
}

func (b *Builder) WithPreReleasePrefix(prePrefix string) *Builder {
	b.preReleasePrefix = prePrefix
	return b
}

func (b *Builder) Build(ctx context.Context) (*Release, error) {
	var buildPreRelease bool
	if b.clusterApp != nil {
		// Check if cluster app version is pre-release
		clusterAppVersion, err := semver.NewVersion(b.clusterApp.Version)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		buildPreRelease = clusterAppVersion.Prerelease() != ""
	}
	for _, app := range b.apps {
		// Check if app version is pre-release
		appVersion, err := semver.NewVersion(app.Version)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		buildPreRelease = appVersion.Prerelease() != ""
	}

	release, err := b.build(ctx, buildPreRelease)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return release, nil
}

func (b *Builder) BuildPreRelease(ctx context.Context) (*Release, error) {
	release, err := b.build(ctx, true)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	return release, nil
}

func (b *Builder) build(ctx context.Context, buildPreRelease bool) (*Release, error) {
	var release *Release
	var err error

	// First we get the base release
	if b.baseRelease == "" {
		// use the latest release as a base
		release, err = b.client.GetLatestRelease(ctx, b.provider)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	} else {
		release, err = b.client.GetRelease(ctx, b.provider, b.baseRelease)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	versionBump := none
	changes := ""

	// Then we override cluster app
	if b.clusterApp != nil {
		currentClusterAppVersion, err := release.GetClusterAppVersion()
		if err != nil {
			return nil, microerror.Mask(err)
		}
		clusterAppVersionBump, err := getVersionPartToBump(b.clusterApp.Version, currentClusterAppVersion)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if clusterAppVersionBump > versionBump {
			versionBump = clusterAppVersionBump
		}
		overrideClusterApp(release, *b.clusterApp)
		changes += fmt.Sprintf("%s-%s-%s", fmt.Sprintf("cluster-%s", b.provider), b.clusterApp.Catalog, b.clusterApp.Version)
	}

	// And we override apps
	for _, appOverride := range b.apps {
		currentApp, ok := release.LookupAppSpec(appOverride.Name)
		if !ok {
			return nil, microerror.Maskf(AppNotFoundError, "app `%s` not found in release resource '%s'", appOverride.Name, release.Name)
		}
		appVersionBump, err := getVersionPartToBump(appOverride.Version, currentApp.Version)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if appVersionBump > versionBump {
			versionBump = appVersionBump
		}
		overrideApp(release, appOverride)
		changes += fmt.Sprintf(";%s-%s-%s", appOverride.Name, appOverride.Catalog, appOverride.Version)
		for _, appDependency := range appOverride.DependsOn {
			changes += fmt.Sprintf("-%s", appDependency)
		}
	}

	// Now we get the base release version
	releaseVersionString, err := release.GetVersion()
	if err != nil {
		return nil, microerror.Mask(err)
	}
	releaseVersion, err := semver.NewVersion(releaseVersionString)
	if err != nil {
		return nil, microerror.Mask(err)
	}
	// clear current pre-release so that patch gets incremented below if needed (because when pre-release is set,
	// calling releaseVersion.IncPatch() does not increment patch number)
	*releaseVersion, err = releaseVersion.SetPrerelease("")
	if err != nil {
		return nil, microerror.Mask(err)
	}

	// And then we bump it
	switch versionBump {
	case major:
		*releaseVersion = releaseVersion.IncMajor()
	case minor:
		*releaseVersion = releaseVersion.IncMinor()
	case patch:
		*releaseVersion = releaseVersion.IncPatch()
	case none:
		*releaseVersion = releaseVersion.IncPatch()
	}

	if buildPreRelease {
		preRelease := hashAndTruncate(changes, preReleaseHashLength)
		if b.preReleasePrefix != "" {
			preRelease = fmt.Sprintf("%s.%s", b.preReleasePrefix, preRelease)
		}
		*releaseVersion, err = releaseVersion.SetPrerelease(preRelease)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	// Finally, we update Release resource name
	release.Name = fmt.Sprintf("%s-%s", b.provider, releaseVersion.String())
	return release, nil
}

func overrideClusterApp(release *Release, clusterApp ReleaseSpecComponent) {
	for i, component := range release.Spec.Components {
		if component.Name != clusterApp.Name {
			continue
		}
		if clusterApp.Catalog == "" {
			// keep current catalog, if the new one is not set
			clusterApp.Catalog = release.Spec.Components[i].Catalog
		}
		release.Spec.Components[i] = clusterApp
		break
	}
}

func overrideApp(release *Release, appOverride ReleaseSpecApp) {
	for i, app := range release.Spec.Apps {
		if app.Name != appOverride.Name {
			continue
		}
		if appOverride.Catalog == "" {
			// keep current catalog, if the new one is not set
			appOverride.Catalog = app.Catalog
		}
		release.Spec.Apps[i] = appOverride
		break
	}
}

// hashAndTruncate takes a string that will be hashed and the desired length to which the hash result should be
// truncated.
//
// It hashes the string using SHA-256 algorithm and returns the truncated hexadecimal representation
// of the hash. If the length input is greater than or equal to the length of the hash, the entire
// hash is returned.
func hashAndTruncate(s string, length int) string {
	h := sha256.New()
	h.Write([]byte(s))
	hashedString := hex.EncodeToString(h.Sum(nil))
	if len(hashedString) < length {
		length = len(hashedString)
	}

	return hashedString[:length]
}

type versionPart int

const (
	major versionPart = 3
	minor versionPart = 2
	patch versionPart = 1
	none  versionPart = 0
)

func getVersionPartToBump(newVersion, oldVersion string) (versionPart, error) {
	newVersionSemver, err := semver.NewVersion(newVersion)
	if err != nil {
		return none, microerror.Mask(err)
	}
	oldVersionSemver, err := semver.NewVersion(oldVersion)
	if err != nil {
		return none, microerror.Mask(err)
	}

	switch {
	case newVersionSemver.Major() > oldVersionSemver.Major():
		return major, nil
	case newVersionSemver.Minor() > oldVersionSemver.Minor():
		return minor, nil
	case newVersionSemver.Patch() > oldVersionSemver.Patch():
		fallthrough
	case newVersionSemver.Prerelease() != oldVersionSemver.Prerelease():
		// in case of a different pre-release, the new version is most probably the latest app/component release with
		// some changes on top of it, so we want to bump at least a patch version.
		return patch, nil
	}

	return none, nil
}
