package registry

import (
	"github.com/xvello/oasis-nomad/pkg/docker/utils"
)

// TagsForImage returns all available tags for a
// given image string
func TagsForImage(image string) ([]string, error) {
	specs, err := utils.ParseImageString(image)
	if err != nil {
		return nil, err
	}

	reg, err := connect(specs.Registry)
	if err != nil {
		return nil, err
	}

	return reg.Tags(specs.Image)
}

// ResolveDigest lookups the registry to update the Digest field
// in a given ImageSpecs struct
func ResolveDigest(i utils.ImageSpecs) (utils.ImageSpecs, error) {
	reg, err := connect(i.Registry)
	if err != nil {
		return i, err
	}

	digest, err := reg.Digest(i.Image, i.Tag)
	if err != nil {
		return i, err
	}
	i.Digest = digest
	return i, nil
}

// ResolveFromString parses an image string as ImageSpec
// and resolves the latest available digest
func ResolveFromString(image string) (utils.ImageSpecs, error) {
	spec, err := utils.ParseImageString(image)
	if err != nil {
		return spec, err
	}

	return ResolveDigest(spec)
}
