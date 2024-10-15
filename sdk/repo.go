package sdk

import (
	"github.com/giantswarm/microerror"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

const (
	archivedDirectory = "archived"
)

func getProviderDirectory(provider Provider) (string, error) {
	switch provider {
	case ProviderAws:
		return "capa", nil
	case ProviderAzure:
		return "azure", nil
	case ProviderVsphere:
		return "vsphere", nil
	}

	return "", microerror.Maskf(InvalidProviderError, "unknown provider: %s", provider)
}
