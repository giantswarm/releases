package filesystem

import "github.com/giantswarm/microerror"

var invalidReleaseError = &microerror.Error{
	Kind: "invalidReleaseError",
}

// IsInvalidRelease asserts invalidReleaseError.
func IsInvalidRelease(err error) bool {
	return microerror.Cause(err) == invalidReleaseError
}

var releaseNotFoundError = &microerror.Error{
	Kind: "releaseNotFoundError",
}

// IsReleaseNotFound asserts releaseNotFoundError.
func IsReleaseNotFound(err error) bool {
	return microerror.Cause(err) == releaseNotFoundError
}
