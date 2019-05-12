package tests

import (
	"context"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type fixture int

const (
	scratchDir fixture = iota
	nomadServer
	registry
)

type oasisSuite struct {
	suite.Suite
	enabled   map[fixture]bool
	suiteCtx  context.Context
	ctxCancel context.CancelFunc
	tempDir   string
}

func newSuite(want ...fixture) *oasisSuite {
	s := &oasisSuite{
		enabled: make(map[fixture]bool),
	}
	for _, f := range want {
		s.enabled[f] = true
	}
	return s
}

func (s *oasisSuite) Wants(f fixture) bool {
	return s.enabled[f]
}

func (s *oasisSuite) ScratchPath(parts ...string) string {
	pathParts := append([]string{s.tempDir, "scratch"}, parts...)
	return filepath.Join(pathParts...)
}

func (s *oasisSuite) SetupSuite() {
	var err error
	s.tempDir, err = ioutil.TempDir("", "oasis-testing-")
	require.NoError(s.T(), err)

	s.suiteCtx, s.ctxCancel = context.WithCancel(context.Background())
}

func (s *oasisSuite) TearDownSuite() {
	err := os.RemoveAll(s.tempDir)
	require.NoError(s.T(), err)
}

func (s *oasisSuite) SetupTest() {
	if s.Wants(scratchDir) {
		err := os.MkdirAll(s.ScratchPath(), 0700)
		require.NoError(s.T(), err)
	}
}

func (s *oasisSuite) TearDownTest() {
	if s.Wants(scratchDir) {
		err := os.RemoveAll(s.ScratchPath())
		require.NoError(s.T(), err)
	}
}
