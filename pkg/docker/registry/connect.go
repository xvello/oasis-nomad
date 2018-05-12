package registry

import (
	"github.com/docker/docker/api/types"
	"github.com/genuinetools/reg/registry"
)

var mockedSource Source

// Source in the interface registry providers must implement.
// It allows to mock the registry connection for tests.
type Source interface {
	Digest(repository, ref string) (string, error)
	Tags(repository string) ([]string, error)
}

func connect(url string) (Source, error) {
	if mockedSource != nil {
		return mockedSource, nil
	}

	auth := types.AuthConfig{
		ServerAddress: url,
	}
	return registry.New(auth, false)
}
