package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/go-cmd/cmd"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/xvello/oasis-nomad/pkg/releases"
)

const (
	nomadBaseURL = "http://localhost:44646"
)

type backgroundRun struct {
	cmd        *cmd.Cmd
	statusChan <-chan cmd.Status
}

func (c *backgroundRun) wait() cmd.Status {
	return <-c.statusChan
}

func (c *backgroundRun) stop() cmd.Status {
	c.cmd.Stop()
	return c.wait()
}

type executable struct {
	Command     string
	DefaultArgs []string
	DataDirPath string
	Env         []string
}

func (e *executable) run(args ...string) cmd.Status {
	c := e.runBackground(args...)
	result := <-c.statusChan
	return result
}

func (e *executable) runBackground(args ...string) *backgroundRun {
	c := cmd.NewCmdOptions(cmd.Options{Buffered: true}, e.Command, args...)
	if len(c.Args) == 0 {
		c.Args = e.DefaultArgs
	}
	c.Env = e.Env

	if e.DataDirPath != "" {
		os.RemoveAll(e.DataDirPath)
		os.MkdirAll(e.DataDirPath, 0700)
	}

	return &backgroundRun{
		cmd:        c,
		statusChan: c.Start(),
	}
}

func newOasis(t *testing.T) *executable {
	return &executable{
		Command: "../oasis",
		Env: []string{
			"NOMAD_ADDR=" + nomadBaseURL,
		},
	}
}

func newNomadServer(t *testing.T, folder, version string) *executable {
	nomadVersions, err := releases.GetReleases("nomad")
	require.NoError(t, err)
	rel, err := nomadVersions.Find(version)
	require.NoError(t, err)
	dest := filepath.Join(folder, "nomad")
	err = rel.Download(dest, "linux", "amd64")
	require.NoError(t, err)

	dataDir := filepath.Join(folder, "data", "nomad")
	return &executable{
		Command: dest,
		DefaultArgs: []string{
			"agent",
			"-config", buildAbs(t, "testdata/nomad.hcl"),
			"-data-dir", dataDir,
		},
		DataDirPath: dataDir,
		Env:         []string{},
	}
}

func newRegistry(t *testing.T, folder string) *executable {
	dataDir := filepath.Join(folder, "data", "registry")
	return &executable{
		Command:     "registry",
		DefaultArgs: []string{"serve", buildAbs(t, "testdata/nomad.hcl")},
		DataDirPath: dataDir,
		Env: []string{
			"REGISTRY_STORAGE_FILESYSTEM_ROOTDIRECTORY=" + dataDir,
		},
	}
}

func buildAbs(t *testing.T, rel string) string {
	path, err := filepath.Abs(rel)
	require.NoError(t, err)
	return path
}

func assertSuccess(t *testing.T, status cmd.Status) {
	t.Helper()
	output := strings.Join(status.Stderr, "\n")
	assert.True(t, status.Complete)
	assert.Equal(t, 0, status.Exit, output)
	assert.NoError(t, status.Error, output)
}
