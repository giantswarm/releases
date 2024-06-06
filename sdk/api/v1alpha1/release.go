package v1alpha1

import (
	"fmt"
	"strings"

	"github.com/giantswarm/microerror"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReleaseState string

const (
	StateActive     ReleaseState = "active"
	StateDeprecated ReleaseState = "deprecated"
	StateWIP        ReleaseState = "wip"
	StatePreview    ReleaseState = "preview"
)

func (r ReleaseState) String() string {
	return string(r)
}

// Release is a Kubernetes resource (CR) representing a Giant Swarm workload cluster release.
type Release struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              ReleaseSpec   `json:"spec"`
	Status            ReleaseStatus `json:"status"`
}

type ReleaseSpec struct {
	// Apps describes apps used in this release.
	Apps []ReleaseSpecApp `json:"apps"`

	// Components describes components used in this release.
	Components []ReleaseSpecComponent `json:"components"`

	// Date that the release became active.
	Date *metav1.Time `json:"date"`

	// EndOfLifeDate is the date and time when support for a workload cluster using
	// this release ends. This may not be set at the time of release creation
	// and can be specified later.
	EndOfLifeDate *metav1.Time `json:"endOfLifeDate,omitempty"`

	// State indicates the availability of the release: deprecated, active, or wip.
	State ReleaseState `json:"state"`

	// Notice outlines anything worth being aware of in this release.
	Notice string `json:"notice,omitempty"`
}

type ReleaseSpecComponent struct {
	// Catalog specifies the name of the app catalog that this component belongs to.
	Catalog string `json:"catalog,omitempty"`
	// Name of the component.
	Name string `json:"name"`
	// Reference is the component's version in the catalog (e.g. 1.2.3 or 1.2.3-abc8675309).
	Reference string `json:"reference,omitempty"`
	// ReleaseOperatorDeploy informs the release-operator that it should deploy the component.
	ReleaseOperatorDeploy bool `json:"releaseOperatorDeploy,omitempty"`
	// Version of the component.
	Version string `json:"version"`
}

type ReleaseSpecApp struct {
	// Catalog specifies the name of the app catalog that this app belongs to.
	Catalog string `json:"catalog,omitempty"`
	// Version of the upstream component used in the app.
	ComponentVersion string `json:"componentVersion,omitempty"`
	// DependsOn is the list of apps that should be installed before installation of this one is attempted.
	DependsOn []string `json:"dependsOn,omitempty"`
	// Name of the app.
	Name string `json:"name"`
	// Version of the app.
	Version string `json:"version"`
}

type ReleaseStatus struct {
	// Ready indicates if all components of the release have been deployed.
	Ready bool `json:"ready"`
	// InUse indicates whether a release is actually used by a cluster.
	InUse bool `json:"inUse"`
}

type ReleaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []Release `json:"items"`
}

func (r *Release) GetProvider() (Provider, error) {
	provider, _, err := r.getResourceNameParts()
	if err != nil {
		return "", microerror.Mask(err)
	}

	return provider, nil
}

func (r *Release) GetVersion() (string, error) {
	_, releaseVersion, err := r.getResourceNameParts()
	if err != nil {
		return "", microerror.Mask(err)
	}

	return releaseVersion, nil
}

func (r *Release) GetKubernetesVersion() (string, error) {
	const kubernetesComponentName = "kubernetes"
	kubernetesComponentSpec, ok := r.LookupComponentSpec(kubernetesComponentName)
	if !ok {
		return "", microerror.Maskf(ComponentNotFoundError, "Component '%s' not found in the Release resource %s.", kubernetesComponentName, r.Name)
	}
	return kubernetesComponentSpec.Version, nil
}

func (r *Release) GetFlatcarVersion() (string, error) {
	const flatcarComponentName = "flatcar"
	flatcarComponentSpec, ok := r.LookupComponentSpec(flatcarComponentName)
	if !ok {
		return "", microerror.Maskf(ComponentNotFoundError, "Component '%s' not found in the Release resource %s.", flatcarComponentName, r.Name)
	}
	return flatcarComponentSpec.Version, nil
}

func (r *Release) GetClusterAppVersion() (string, error) {
	provider, err := r.GetProvider()
	if err != nil {
		return "", microerror.Mask(err)
	}
	var clusterAppName = fmt.Sprintf("cluster-%s", provider)
	clusterAppComponentSpec, ok := r.LookupComponentSpec(clusterAppName)
	if !ok {
		return "", microerror.Maskf(ComponentNotFoundError, "Component '%s' not found in the Release resource %s.", clusterAppName, r.Name)
	}
	return clusterAppComponentSpec.Version, nil
}

func (r *Release) LookupAppSpec(appName string) (ReleaseSpecApp, bool) {
	for _, app := range r.Spec.Apps {
		if app.Name == appName {
			return app, true
		}
	}

	return ReleaseSpecApp{}, false
}

func (r *Release) LookupComponentSpec(appName string) (ReleaseSpecComponent, bool) {
	for _, component := range r.Spec.Components {
		if component.Name == appName {
			return component, true
		}
	}

	return ReleaseSpecComponent{}, false
}

func (r *Release) getResourceNameParts() (Provider, string, error) {
	providerString, releaseVersion, ok := strings.Cut(r.Name, "-")
	if !ok {
		return "", "", microerror.Maskf(InvalidReleaseResourceNameError, "Release resource must be in format '{provider}-{version}', but '-' is not found in '%s'.", r.Name)
	}
	provider := Provider(providerString)

	if !IsProviderSupported(provider) {
		return "", "", microerror.Maskf(UnsupportedProviderError, "Provider '%s' is not supported. Supported providers are: %s.", provider, SupportedProviders)
	}

	return provider, releaseVersion, nil
}
