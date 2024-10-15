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
	case "":
		return "", microerror.Maskf(InvalidProviderError, "no provider supplied")

	// Providers that use a different path name to their provider name
	case ProviderAws:
		return "capa", nil

	// By default, the value of the provider is used as the path
	default:
		return string(provider), nil
	}
}
