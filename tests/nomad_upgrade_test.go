// +build !nonetwork

package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type NomadUpgradeSuite struct {
	*oasisSuite
	oasis *executable
}

func (s *NomadUpgradeSuite) SetupTest() {
	s.oasisSuite.SetupTest()
	err := os.MkdirAll(s.ScratchPath("bin"), 0700)
	require.NoError(s.T(), err)
}

func (s *NomadUpgradeSuite) TestSetup() {
	assert.True(s.T(), s.Wants(scratchDir))
	assert.False(s.T(), s.Wants(nomadServer))
	assert.NotEqual(s.T(), "", s.ScratchPath("nomad"))
}

func (s *NomadUpgradeSuite) TestDownloadAndUpgrade() {
	// Initial download
	status := s.oasis.run("nomad", "upgrade", "-p", s.ScratchPath(), "-v", "0.8.6")
	assertSuccess(s.T(), status)

	getVersion := &executable{
		Command:     s.ScratchPath("bin", "nomad"),
		DefaultArgs: []string{"version"},
	}
	status = getVersion.run()
	assertSuccess(s.T(), status)
	require.Len(s.T(), status.Stdout, 1)
	assert.Contains(s.T(), status.Stdout[0], "Nomad v0.8.6")

	// Upgrade and replace
	status = s.oasis.run("nomad", "upgrade", "-p", s.ScratchPath(), "-v", "0.9.1")
	assertSuccess(s.T(), status)
	status = getVersion.run()
	assertSuccess(s.T(), status)
	require.Len(s.T(), status.Stdout, 1)
	assert.Contains(s.T(), status.Stdout[0], "Nomad v0.9.1")

	// Both versions exist
	_, err := os.Stat(s.ScratchPath("bin", "nomad-0.8.6"))
	assert.NoError(s.T(), err)
	_, err = os.Stat(s.ScratchPath("bin", "nomad-0.9.1"))
	assert.NoError(s.T(), err)
}

func TestNomadUpgradeSuite(t *testing.T) {
	suite.Run(t, &NomadUpgradeSuite{
		oasisSuite: newSuite(scratchDir),
		oasis:      newOasis(t),
	})
}
