package releases

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func parseIndexFixture() (*ReleaseProduct, error) {
	file, err := os.Open("testdata/index.json")
	if err != nil {
		return nil, err
	}
	return parseReleases(file)
}

func TestParseReleases(t *testing.T) {
	_, err := parseIndexFixture()
	require.NoError(t, err)
}

func TestLatestStable(t *testing.T) {
	r, _ := parseIndexFixture()
	l, err := r.Latest(true)
	require.NoError(t, err)
	require.NotNil(t, l)
	assert.Equal(t, "0.8.4", l.Version)
}

func TestLatestUnStable(t *testing.T) {
	// 0.8.4-rc1 should be sorted before 0.8.4
	r, _ := parseIndexFixture()
	l, err := r.Latest(false)
	require.NoError(t, err)
	require.NotNil(t, l)
	assert.Equal(t, "0.8.4", l.Version)
}

func TestListVersionsStable(t *testing.T) {
	r, _ := parseIndexFixture()
	versions, err := r.ListVersions(true)
	require.NoError(t, err)

	var found bool
	for _, v := range versions {
		assert.NotEqual(t, "0.8.4-rc1", v.Original())
		if v.Original() == "0.8.4" {
			found = true
		}
	}

	assert.True(t, found, "version 0.8.4 not found")
}

func TestListVersionsUnStable(t *testing.T) {
	r, _ := parseIndexFixture()
	versions, err := r.ListVersions(false)
	require.NoError(t, err)

	var foundStable, foundUnstable bool
	for _, v := range versions {
		if v.Original() == "0.8.4-rc1" {
			foundUnstable = true
		}
		if v.Original() == "0.8.4" {
			foundStable = true
		}
	}
	assert.True(t, foundStable, "version 0.8.4 not found")
	assert.True(t, foundUnstable, "version 0.8.4-rc1 not found")
}
