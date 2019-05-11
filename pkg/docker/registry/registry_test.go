package registry

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type registrySuite struct {
	suite.Suite
	reg *MockedRegistry
}

func (s *registrySuite) SetupTest() {
	reg, err := SetupMock()
	s.NoError(err)
	s.reg = reg
}

func (s *registrySuite) TearDownTest() {
	ResetMock()
	s.reg = nil
}

func (s *registrySuite) TestTags() {
	s.reg.AddTag("library/redis", "latest", "1234")
	s.reg.AddTag("library/redis", "3.2", "1234")

	cases := []struct {
		input string
		tags  []string
		err   error
	}{
		{
			input: "redis",
			tags:  []string{"latest", "3.2"},
		},
		{
			input: "my/repo",
			err:   errors.New("path my/repo not found"),
		},
	}
	for _, tc := range cases {
		tags, err := TagsForImage(tc.input)
		if tc.err == nil {
			s.NoError(err)
		} else {
			s.Equal(tc.err.Error(), err.Error())
		}
		s.ElementsMatch(tc.tags, tags)
	}
}

func (s *registrySuite) TestDigest() {
	sha := "sha256:87856cc39862cec77541d68382e4867d7ccb29a85a17221446c857ddaebca916"
	s.reg.AddTag("library/redis", "3.2", sha)

	cases := []struct {
		input  string
		digest string
		err    error
	}{
		{
			input: "redis",
			err:   errors.New("tag latest not found for path library/redis"),
		},
		{
			input:  "redis:3.2",
			digest: sha,
		},
		{
			input: "redis:3.4",
			err:   errors.New("tag 3.4 not found for path library/redis"),
		},
		{
			input: "my/repo",
			err:   errors.New("path my/repo not found"),
		},
	}
	for i, tc := range cases {
		s.T().Run(fmt.Sprintf("%d: %s", i, tc.input), func(t *testing.T) {
			resolved, err := ResolveFromString(tc.input)
			if tc.err == nil {
				s.NoError(err)
			} else {
				s.Equal(tc.err.Error(), err.Error())
			}
			s.Equal(tc.digest, resolved.Digest.String())
		})
	}
}

func TestRegistrySuite(t *testing.T) {
	suite.Run(t, new(registrySuite))
}
