package utils

import "fmt"

// ImageSpecs holds the specs for a container image reference
type ImageSpecs struct {
	Registry string
	Image    string
	Tag      string
	Digest   string
}

// String returns the docker image string representation
// FIXME: support custom registries
func (i ImageSpecs) String() string {
	out := fmt.Sprintf("%s:%s", i.Image, i.Tag)
	if i.Digest != "" {
		out = fmt.Sprintf("%s@%s", out, i.Digest)
	}

	return out
}
