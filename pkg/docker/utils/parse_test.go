package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseImageString(t *testing.T) {
	cases := []struct {
		input string
		specs ImageSpecs
		err   error
	}{
		{
			input: "user/test",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "user/test",
				Tag:      "latest",
				Digest:   "",
			},
			err: nil,
		},
		{
			input: "user/test:tagged",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "user/test",
				Tag:      "tagged",
				Digest:   "",
			},
			err: nil,
		},
		{
			input: "redis",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "library/redis",
				Tag:      "latest",
				Digest:   "",
			},
			err: nil,
		},
		{
			input: "user/test:mytag@sha:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "user/test",
				Tag:      "mytag",
				Digest:   "sha:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a",
			},
			err: nil,
		},
		{
			input: "user/test@sha:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "user/test",
				Tag:      "latest",
				Digest:   "sha:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a",
			},
			err: nil,
		},
		{
			input: "docker.io/user/test",
			specs: ImageSpecs{
				Registry: dockerHubRegistry,
				Image:    "user/test",
				Tag:      "latest",
				Digest:   "",
			},
			err: nil,
		},
		{
			input: "quay.io/user/test:version",
			specs: ImageSpecs{
				Registry: "quay.io",
				Image:    "user/test",
				Tag:      "version",
				Digest:   "",
			},
			err: nil,
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.input), func(t *testing.T) {
			got, err := ParseImageString(tc.input)
			if tc.err == nil {
				assert.NoError(t, err)
			} else {
				assert.Equal(t, tc.err.Error(), err.Error())
			}
			assert.Equal(t, tc.specs.Registry, got.Registry)
			assert.Equal(t, tc.specs.Image, got.Image)
			assert.Equal(t, tc.specs.Tag, got.Tag)
			assert.Equal(t, tc.specs.Digest, got.Digest)

			// Actually testing I,ageSpecs.String()
			assert.Equal(t, tc.specs, got)
		})
	}
}
