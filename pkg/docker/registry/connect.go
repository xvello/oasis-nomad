package registry

import (
	"time"

	"github.com/docker/docker/api/types"
	"github.com/genuinetools/reg/registry"
	digest "github.com/opencontainers/go-digest"
)

var mockedSource Source

// Source in the interface registry providers must implement.
// It allows to mock the registry connection for tests.
type Source interface {
	Digest(image registry.Image) (digest.Digest, error)
	Tags(repository string) ([]string, error)
}

func connect(url string) (Source, error) {
	if mockedSource != nil {
		return mockedSource, nil
	}

	// Because as usual Docker likes to be a special case
	if url == "docker.io" {
		url = "https://registry.hub.docker.com"
	}

	opt := registry.Opt{
		Domain:   url,
		SkipPing: true,
		Timeout:  5 * time.Second,
	}
	return registry.New(types.AuthConfig{}, opt)
}
