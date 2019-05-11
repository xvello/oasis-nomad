package registry

import (
	"strings"

	"github.com/genuinetools/reg/registry"
)

const (
	dockerHubPrefix     = "docker.io/"
	dockerLibraryPrefix = dockerHubPrefix + "library/"
)

// TagsForImage returns all available tags for a
// given image string
func TagsForImage(image string) ([]string, error) {
	specs, err := registry.ParseImage(image)
	if err != nil {
		return nil, err
	}

	reg, err := connect(specs.Domain)
	if err != nil {
		return nil, err
	}

	return reg.Tags(specs.Path)
}

// ResolveDigest lookups the registry to update the Digest field
// in a given ImageSpecs struct
func ResolveDigest(i registry.Image) (registry.Image, error) {
	reg, err := connect(i.Domain)
	if err != nil {
		return i, err
	}

	digest, err := reg.Digest(i)
	if err != nil {
		return i, err
	}
	err = i.WithDigest(digest)
	if err != nil {
		return i, err
	}

	return i, nil
}

// ResolveFromString parses an image string as ImageSpec
// and resolves the latest available digest
func ResolveFromString(image string) (registry.Image, error) {
	spec, err := registry.ParseImage(image)
	if err != nil {
		return spec, err
	}
	return ResolveDigest(spec)
}

// ImageShortString returns a concise string representation
// of the passed image, removing the docker.io registry
// and library path if present, as these can be implicit
func ImageShortString(image registry.Image) string {
	long := image.String()
	for _, prefix := range []string{dockerLibraryPrefix, dockerHubPrefix} {
		if strings.HasPrefix(long, prefix) {
			return strings.TrimPrefix(long, prefix)
		}
	}
	return long
}
