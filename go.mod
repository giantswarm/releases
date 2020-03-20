module github.com/giantswarm/releases

go 1.13

require (
	github.com/fatih/color v1.7.0
	github.com/giantswarm/apiextensions v0.1.1
	github.com/giantswarm/microerror v0.2.0
	github.com/giantswarm/micrologger v0.3.0 // indirect
	github.com/giantswarm/versionbundle v0.1.0
	github.com/go-openapi/errors v0.19.4 // indirect
	github.com/juju/errgo v0.0.0-20140925100237-08cceb5d0b53 // indirect
	k8s.io/apiextensions-apiserver v0.17.4
	sigs.k8s.io/yaml v1.2.0
)

replace github.com/giantswarm/microerror => github.com/giantswarm/microerror v0.1.0
