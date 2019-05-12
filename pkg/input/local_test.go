package input

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	relativeName = "relative"
	absoluteName = "absolute"
)

func TestLocalTargets(t *testing.T) {
	initialDir, err := os.Getwd()
	require.NoError(t, err)
	defer os.Chdir(initialDir)

	dir, err := ioutil.TempDir("", "oasis-test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)
	os.Chdir(dir)

	relativePath := filepath.Join(dir, relativeName)
	err = ioutil.WriteFile(relativePath, []byte(relativeName), 0600)
	require.NoError(t, err)
	absolutePath := filepath.Join(dir, absoluteName)
	err = ioutil.WriteFile(absolutePath, []byte(absoluteName), 0600)
	require.NoError(t, err)

	expectedNames := []string{
		relativePath, // relative file name
		absolutePath, // absolute file path
		absolutePath, // widcard selector
		relativePath,
	}

	targets, err := Local([]string{relativeName, absolutePath, "*a*"})
	require.NoError(t, err)
	files := targets.Files()

	for i, expected := range expectedNames {
		select {
		case f := <-files:
			require.NotNil(t, f, "expected more files")
			assert.Equal(t, strings.TrimPrefix(expected, string(filepath.Separator)), f.Name())
			f.Close()
		case <-time.After(time.Millisecond):
			assert.FailNow(t, "expected file number %d", i)
		}
	}
	select {
	case f := <-files:
		assert.Nil(t, f, "unexpected remaining file")
	case <-time.After(time.Millisecond):
		assert.FailNow(t, "expected channel to be closed")
	}
}
