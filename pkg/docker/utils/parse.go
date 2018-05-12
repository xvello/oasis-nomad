package utils

import (
	"fmt"
	"strings"
)

const (
	dockerHubRegistry = "registry-1.docker.io"
	defaultTag        = "latest"
	dockerHubLibrary  = "library"
)

// ParseImageString parses a human-readable image name into
// Registry + Image + Tag
// TODO: support custom registries
func ParseImageString(image string) (ImageSpecs, error) {
	specs := ImageSpecs{
		Registry: dockerHubRegistry,
		Tag:      defaultTag,
	}

	// Parse digest
	parts := strings.SplitN(image, "@", 2)
	if len(parts) == 2 {
		specs.Digest = parts[1]
	}

	// Parse tag
	parts = strings.SplitN(parts[0], ":", 2)
	if len(parts) == 2 {
		specs.Tag = parts[1]
	}

	// Parse image
	specs.Image = parts[0]

	// Special case of the docker library images
	if specs.Registry == dockerHubRegistry {
		if !strings.Contains(specs.Image, "/") {
			specs.Image = fmt.Sprintf("%s/%s", dockerHubLibrary, specs.Image)
		}
	}

	return specs, nil
}
