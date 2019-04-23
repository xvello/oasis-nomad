package utils

import (
	"fmt"
	"strings"
)

const (
	dockerHubRegistry = "registry-1.docker.io"
	dockerHubShortReg = "docker.io"

	defaultTag       = "latest"
	dockerHubLibrary = "library"
)

// ParseImageString parses a human-readable image name into
// Registry + Image + Tag
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

	// Parse registry out of the image name
	if strings.Count(parts[0], "/") == 2 {
		parts = strings.SplitN(parts[0], "/", 2)
		specs.Registry = parts[0]
		specs.Image = parts[1]
	} else {
		specs.Image = parts[0]
	}

	// Special case of the docker registry
	if specs.Registry == dockerHubShortReg {
		specs.Registry = dockerHubRegistry
	}

	// Special case of the docker library images
	if specs.Registry == dockerHubRegistry {
		if !strings.Contains(specs.Image, "/") {
			specs.Image = fmt.Sprintf("%s/%s", dockerHubLibrary, specs.Image)
		}
	}

	return specs, nil
}
