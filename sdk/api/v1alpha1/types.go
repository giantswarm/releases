package v1alpha1

import (
	"slices"
)

const (
	// ProviderAws represents the new AWS provider which is implemented with Cluster API.
	ProviderAws Provider = "aws"

	// ProviderAzure represents the new Azure provider which is implemented with Cluster API.
	ProviderAzure Provider = "azure"

	// ProviderVsphere represents the new Vsphere provider which is implemented with Cluster API.
	ProviderVsphere Provider = "vsphere"

	// ProviderCloudDirector represents the new VCD provider which is implemented with Cluster API.
	ProviderCloudDirector Provider = "cloud-director"

	// ProviderEKS represents the new EKS provider which is implemented with Cluster API.
	ProviderEKS Provider = "eks3"
)

var SupportedProviders = []Provider{
	ProviderAws,
	ProviderAzure,
	ProviderVsphere,
	ProviderCloudDirector,
	ProviderEKS,
}

type Provider string

// IsProviderSupported returns a flag that indicates if the specified provider is supported in Releases SDK.
func IsProviderSupported(provider Provider) bool {
	return slices.Contains(SupportedProviders, provider)
}
