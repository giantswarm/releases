package v1alpha1

import (
	"github.com/giantswarm/microerror"
)

var InvalidReleaseResourceNameError = &microerror.Error{
	Kind: "InvalidReleaseResourceNameError",
}

var UnsupportedProviderError = &microerror.Error{
	Kind: "UnsupportedProviderError",
}

var ComponentNotFoundError = &microerror.Error{
	Kind: "ReleaseComponentNotFoundError",
}
