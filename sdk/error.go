package sdk

import (
	"github.com/giantswarm/microerror"
)

var InvalidConfigError = &microerror.Error{
	Kind: "InvalidConfigError",
}

var ReleaseNotFoundError = &microerror.Error{
	Kind: "ReleaseNotFoundError",
}

var MissingReleaseManifestError = &microerror.Error{
	Kind: "MissingReleaseManifestError",
}
