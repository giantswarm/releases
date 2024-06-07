package sdk

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/Masterminds/semver/v3"
	"github.com/giantswarm/microerror"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

const (
	preReleaseHashLength = 10
)

// Builder builds custom Releases by overriding a base release with custom cluster app and default apps.
type Builder struct {
	client      *Client
	provider    Provider
	baseRelease string

	// pre-release config
	preRelease preReleaseConfig

	// overrides
	clusterApp *ReleaseSpecComponent
	apps       []ReleaseSpecApp
}

// NewBuilder creates a new Builder with specified releases Client, provider and base release.
// If a base release is an empty string it will use the latest available release for the provider.
func NewBuilder(client *Client, provider Provider, baseRelease string) (*Builder, error) {
	if client == nil {
		return nil, microerror.Maskf(InvalidConfigError, "client must not be empty")
	}
	if !IsProviderSupported(provider) {
		return nil, microerror.Maskf(UnsupportedProviderError, "provider `%s` is not supported", provider)
	}

	var overridesDiff diff
	return &Builder{
		client:      client,
		provider:    provider,
		baseRelease: baseRelease,
		preRelease: preReleaseConfig{
			overridesDiff: &overridesDiff,
		},
	}, nil
}

// WithClusterApp overrides the cluster-<provider> app in the base release.
func (b *Builder) WithClusterApp(version, catalog string) *Builder {
	b.clusterApp = &ReleaseSpecComponent{
		Name:    fmt.Sprintf("cluster-%s", b.provider),
		Version: version,
		Catalog: catalog,
	}
	return b
}

// WithApp overrides the default app in the base release.
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

// WithPreReleasePrefix sets a custom prefix that is prepended to the pre-release of the custom release version.
func (b *Builder) WithPreReleasePrefix(prefix string) *Builder {
	b.preRelease.prefix = prefix
	return b
}

// WithRandomPreRelease specifies that custom release version will have a random pre-release with specified length. The
// specified length does not include the length of the optionally specified pre-release prefix.
func (b *Builder) WithRandomPreRelease(length int) *Builder {
	b.preRelease.isRandom = true
	b.preRelease.length = length
	return b
}

// Build a custom release.
func (b *Builder) Build(ctx context.Context) (*Release, error) {
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
	b.baseRelease, err = release.GetVersion()
	if err != nil {
		return nil, microerror.Mask(err)
	}

	versionBump := none

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
		b.preRelease.overridesDiff.addComponentDiff(*b.clusterApp)
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
		b.preRelease.overridesDiff.addAppDiff(appOverride)
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

	// Set pre-release is one is needed
	preRelease, err := b.buildPreReleaseString()
	if err != nil {
		return nil, microerror.Mask(err)
	}
	if preRelease != "" {
		*releaseVersion, err = releaseVersion.SetPrerelease(preRelease)
		if err != nil {
			return nil, microerror.Mask(err)
		}
	}

	// Finally, we update Release resource name
	release.Name = fmt.Sprintf("%s-%s", b.provider, releaseVersion.String())
	return release, nil
}

func (b *Builder) buildPreReleaseString() (string, error) {
	if b.preRelease.length == 0 {
		b.preRelease.length = 10
	}
	if b.preRelease.isRandom {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		const charsWithoutZero = "abcdefghijklmnopqrstuvwxyz123456789"
		const chars = "abcdefghijklmnopqrstuvwxyz0123456789"
		result := make([]byte, b.preRelease.length)
		// set first char to non-zero char
		result[0] = charsWithoutZero[r.Intn(len(charsWithoutZero))]
		// set remaining chars
		for i := 1; i < b.preRelease.length; i++ {
			result[i] = chars[r.Intn(len(chars))]
		}
		preReleaseString := string(result)
		if b.preRelease.prefix != "" {
			preReleaseString = fmt.Sprintf("%s.%s", b.preRelease.prefix, preReleaseString)
		}
		return preReleaseString, nil
	}

	// Since random pre-release is not needed, we now check if we need a pre-release or not.
	// A pre-release string is needed if one of the following conditions is met:
	// 1. Base release has a pre-release string, i.e. when you customise a base Release with a pre-release version
	//    you get a custom Release with a pre-release version.
	// 2. Cluster app override has a pre-release string, i.e. when you override a cluster app with a pre-release
	//    version you get a custom Release with a pre-release version.
	// 3. One of app overrides has a pre-release string, i.e. when you override an app with a pre-release version
	//    you get a custom Release with pre-release version.

	preReleaseBuildFunc := func() string {
		preReleaseString := b.preRelease.overridesDiff.hashAndTruncate(b.preRelease.length)
		if b.preRelease.prefix != "" {
			preReleaseString = fmt.Sprintf("%s.%s", b.preRelease.prefix, preReleaseString)
		}
		return preReleaseString
	}

	// check base release
	baseReleaseHasPreReleaseString, err := versionStringHasPreReleaseString(b.baseRelease)
	if err != nil {
		return "", microerror.Mask(err)
	}
	if baseReleaseHasPreReleaseString {
		return preReleaseBuildFunc(), nil
	}

	// check if cluster app has a pre-release
	clusterAppPreReleaseString, err := versionStringHasPreReleaseString(b.clusterApp.Version)
	if err != nil {
		return "", microerror.Mask(err)
	}
	if clusterAppPreReleaseString {
		return preReleaseBuildFunc(), nil
	}

	// check if apps have a pre-release
	for _, app := range b.apps {
		appPreReleaseString, err := versionStringHasPreReleaseString(app.Version)
		if err != nil {
			return "", microerror.Mask(err)
		}
		if appPreReleaseString {
			return preReleaseBuildFunc(), nil
		}
	}

	// Finally, we return only pre-release prefix it has been specified
	// check if pre-release prefix is specified
	if b.preRelease.prefix != "" {
		return b.preRelease.prefix, nil
	}

	return "", nil
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

// preReleaseConfig specifies how the pre-release is built.
type preReleaseConfig struct {
	// prefix of the pre-release string. e.g. v1.2.3-prefix.h97o23f46t
	prefix string

	// overridesDiff contains diff between apps and components from the base release and the apps and components that
	// are overriding them.
	overridesDiff *diff

	// isRandom is a flag that indicates if the pre-release suffix is random.
	isRandom bool

	// length of the pre-release string (without prefix).
	length int
}

type diff string

func (d *diff) addAppDiff(app ReleaseSpecApp) {
	newDiff := fmt.Sprintf("%s+%s+%s+%s", string(*d), app.Name, app.Catalog, app.Version)
	for _, dependsOn := range app.DependsOn {
		newDiff = fmt.Sprintf("%s+%s", newDiff, dependsOn)
	}
	*d = diff(newDiff)
}

func (d *diff) addComponentDiff(component ReleaseSpecComponent) {
	newDiff := fmt.Sprintf("%s+%s+%s+%s", string(*d), component.Name, component.Catalog, component.Version)
	*d = diff(newDiff)
}

// hashAndTruncate takes a string that will be hashed and the desired length to which the hash result should be
// truncated.
//
// It hashes the string using SHA-256 algorithm and returns the truncated hexadecimal representation
// of the hash. If the length input is greater than or equal to the length of the hash, the entire
// hash is returned.
func (d *diff) hashAndTruncate(length int) string {
	h := sha256.New()
	h.Write([]byte(*d))
	hashedString := hex.EncodeToString(h.Sum(nil))
	if len(hashedString) < length {
		length = len(hashedString)
	}

	return hashedString[:length]
}

func versionStringHasPreReleaseString(v string) (bool, error) {
	semverVersion, err := semver.NewVersion(v)
	if err != nil {
		return false, microerror.Mask(err)
	}

	return semverVersion.Prerelease() != "", nil
}
