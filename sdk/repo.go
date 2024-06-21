package sdk

import (
	"github.com/giantswarm/microerror"

	. "github.com/giantswarm/releases/sdk/api/v1alpha1"
)

func getProviderDirectory(provider Provider) (string, error) {
	switch provider {
	case ProviderAws:
		return "capa", nil
	}

	return "", microerror.Maskf(InvalidProviderError, "unknown provider: %s", provider)
}
