package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type NomadCommandsSuite struct {
	*oasisSuite
}

func (s *NomadCommandsSuite) SetupTest() {
	s.oasisSuite.SetupTest()
	status := s.oasis.run("wait", "--frequency", "50ms")
	assertSuccess(s.T(), status)
}

func (s *NomadCommandsSuite) TestWait() {
	status := s.oasis.run("wait", "--frequency", "50ms")
	assertSuccess(s.T(), status)
}

func (s *NomadCommandsSuite) TestRunAndReset() {
	jobs, err := listNomadJobs()
	require.NoError(s.T(), err)
	assert.Len(s.T(), jobs, 0)

	status := s.oasis.run("run", "testdata/job*.nomad")
	assertSuccess(s.T(), status)

	jobs, err = listNomadJobs()
	require.NoError(s.T(), err)
	assert.Len(s.T(), jobs, 2)

	status = s.oasis.run("reset")
	assertSuccess(s.T(), status)

	jobs, err = listNomadJobs()
	require.NoError(s.T(), err)
	assert.Len(s.T(), jobs, 0)
}

func TestNomadCommandsSuite(t *testing.T) {
	suite.Run(t, &NomadCommandsSuite{
		oasisSuite: newSuite(nomadServer),
	})
}
