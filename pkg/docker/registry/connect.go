package registry

import (
	"time"

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
	opt := registry.Opt{
		SkipPing: true,
		Timeout:  5 * time.Second,
	}
	return registry.New(auth, opt)
}
