package releases

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExtractSha(t *testing.T) {
	sumsFile, err := os.Open("testdata/nomad_0.8.4_SHA256SUMS")
	require.NoError(t, err)

	var sumsBuf bytes.Buffer
	io.Copy(&sumsBuf, sumsFile)

	sum, err := extractSha(sumsBuf, "nomad_0.8.4_linux_amd64.zip")
	require.NoError(t, err)
	require.Equal(t, []byte("42fc455d09ea0085cc79d7be4fb2089a9ab7db3cc2e8047e8437855a81b090e9"), sum)
}

/*func TestIsStable(t *testing.T) {
	cases := []struct {
		version *ReleaseVersion
		stable  bool
	}{
		{
			version: &ReleaseVersion{Version: "0.8.4"},
			stable:  true,
		},
		{
			version: &ReleaseVersion{Version: "1.2.12"},
			stable:  true,
		},
		{
			version: &ReleaseVersion{Version: "0.8.4-rc1"},
			stable:  false,
		},
		{
			version: &ReleaseVersion{Version: "0.7.0-beta1"},
			stable:  false,
		},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("%d: %s", i, tc.version.Version), func(t *testing.T) {
			assert.Equal(t, tc.stable, tc.version.IsStable())
		})
	}
}
*/
