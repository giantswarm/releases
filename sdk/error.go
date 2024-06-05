package sdk

import (
	"github.com/giantswarm/microerror"
)

var InvalidConfigError = &microerror.Error{
	Kind: "InvalidConfigError",
}

var MissingReleaseManifestError = &microerror.Error{
	Kind: "MissingReleaseManifestError",
}
