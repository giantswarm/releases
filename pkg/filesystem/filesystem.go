package filesystem

import (
	"io/ioutil"
	"path/filepath"

	"github.com/giantswarm/apiextensions/pkg/apis/release/v1alpha1"
	"github.com/giantswarm/microerror"
	"sigs.k8s.io/yaml"

	"github.com/giantswarm/releases/pkg/key"
)

type Filesystem struct {
	root string
}

func New(root string) Filesystem {
	return Filesystem{
		root: root,
	}
}

func (f Filesystem) ReadFile(path string) ([]byte, error) {
	content, err := ioutil.ReadFile(filepath.Join(f.root, path))
	if err != nil {
		return nil, microerror.Mask(err)
	}
	return content, nil
}

func (f Filesystem) FindRelease(provider string, name string, archived bool) (v1alpha1.Release, error) {
	releases, err := f.FindReleases(provider, archived)
	if err != nil {
		return v1alpha1.Release{}, microerror.Mask(err)
	}

	for _, release := range releases {
		if release.Name == name {
			return release, nil
		}
	}

	return v1alpha1.Release{}, microerror.Mask(releaseNotFoundError)
}

func (f Filesystem) FindReleases(provider string, archived bool) ([]v1alpha1.Release, error) {
	path := filepath.Join(f.root, provider)
	if archived {
		path = filepath.Join(path, "archived")
	}

	releaseDirectories, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, microerror.Mask(err)
	}

	var releases []v1alpha1.Release
	for _, releaseDirectory := range releaseDirectories {
		if !releaseDirectory.IsDir() || releaseDirectory.Name() == "archived" {
			continue
		}

		releaseFile := filepath.Join(path, releaseDirectory.Name(), key.ReleaseFilename)
		data, err := ioutil.ReadFile(releaseFile)
		if err != nil {
			return nil, microerror.Mask(err)
		}

		var release v1alpha1.Release
		err = yaml.Unmarshal(data, &release)
		if err != nil {
			return nil, microerror.Mask(err)
		}
		if releaseDirectory.Name() != release.Name {
			return nil, microerror.Maskf(invalidReleaseError, "%s release %s is in directory %s which doesn't match its name", provider, release.Name, releaseDirectory)
		}

		releases = append(releases, release)
	}

	return releases, nil
}
