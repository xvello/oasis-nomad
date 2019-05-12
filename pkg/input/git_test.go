package input

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/codeclysm/extract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGitTargets(t *testing.T) {
	dir, err := ioutil.TempDir("", "oasis-test")
	require.NoError(t, err)
	defer os.RemoveAll(dir)

	f, err := os.Open("testdata/bare-git.tar")
	require.NoError(t, err)
	err = extract.Tar(context.Background(), f, dir, nil)
	require.NoError(t, err)
	f.Close()

	expectedNames := []string{"b", "a1", "a2"}
	gitURL := fmt.Sprintf("%s/bare-git", dir)
	targets, err := FromGit(gitURL, []string{"c", "b", "a*"})
	require.NoError(t, err)
	files := targets.Files()

	for i, expected := range expectedNames {
		select {
		case f := <-files:
			require.NotNil(t, f, "expected more files")
			assert.Equal(t, expected, f.Name())
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
