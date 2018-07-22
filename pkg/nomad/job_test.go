package nomad

import (
	"testing"

	"github.com/hashicorp/nomad/jobspec"
	"github.com/stretchr/testify/suite"

	"github.com/xvello/oasis-nomad/pkg/docker/registry"
)

type injectSuite struct {
	suite.Suite
	reg *registry.MockedRegistry
}

func (s *injectSuite) SetupTest() {
	reg, err := registry.SetupMock()
	s.NoError(err)
	s.reg = reg
}

func (s *injectSuite) TearDownTest() {
	registry.ResetMock()
	s.reg = nil
}

func (s *injectSuite) TestInject() {
	s.reg.AddTag("library/redis", "3.2", "sha256:6ff2a3a2ddb62378e778180ead0acaf5b44f6e719e42a1ae8c261dd969a16f2a")

	input, err := jobspec.ParseFile("testdata/inject_input.nomad")
	s.NoError(err)
	expected, err := jobspec.ParseFile("testdata/inject_output.nomad")
	s.NoError(err)
	s.NotNil(expected)

	err = addDigests(input)
	s.NoError(err)

	s.Equal(expected, input)
}

func TestInjectSuite(t *testing.T) {
	suite.Run(t, new(injectSuite))
}
