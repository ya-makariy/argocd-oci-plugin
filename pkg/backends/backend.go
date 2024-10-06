package backends

import (
	"context"

	oras "oras.land/oras-go/v2"
	"oras.land/oras-go/v2/content/file"
	"oras.land/oras-go/v2/registry/remote"
	"oras.land/oras-go/v2/registry/remote/auth"
)

type RegistryBackend struct {
	context context.Context
	client  auth.Client
}

func NewRegistryBackend(client auth.Client) *RegistryBackend {
	return &RegistryBackend{
		context: context.Background(),
		client:  client,
	}
}

func (rb *RegistryBackend) PullFiles(repo *remote.Repository, tag, fspath string) error {
	fs, err := file.New(fspath)
	if err != nil {
		return err
	}
	defer fs.Close()

	repo.Client = &rb.client
	_, err = oras.Copy(rb.context, repo, tag, fs, tag, oras.DefaultCopyOptions)
	if err != nil {
		return err
	}

	return nil
}
