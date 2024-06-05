package v1alpha1

import (
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
