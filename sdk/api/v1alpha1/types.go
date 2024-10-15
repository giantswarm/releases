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
)

var SupportedProviders = []Provider{
	ProviderAws,
	ProviderAzure,
	ProviderVsphere,
}

type Provider string

// IsProviderSupported returns a flag that indicates if the specified provider is supported in Releases SDK.
func IsProviderSupported(provider Provider) bool {
	return slices.Contains(SupportedProviders, provider)
}
