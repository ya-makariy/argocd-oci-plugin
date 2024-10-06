package types

import "oras.land/oras-go/v2/registry/remote"

// Backend is an interface for the types of OCI-based registries that are supported
type Backend interface {
	// PullFiles pull the file at `repo` with specified `tag` based on configuration given in `annotations`
	PullFiles(repo *remote.Repository, tag, fspath string) error

	// // PullIndividualFiles pull the specific secret from `repo` with specified `tag` based on configuration given in `annotations`
	// PullIndividualFiles(repo, tag, filename string, annotations map[string]string) (interface{}, error)
}
