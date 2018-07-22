package releases

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVerifySignatureOK(t *testing.T) {
	signed, err := os.Open("testdata/nomad_0.8.4_SHA256SUMS")
	require.NoError(t, err)
	signature, err := os.Open("testdata/nomad_0.8.4_SHA256SUMS.sig")
	require.NoError(t, err)

	require.NoError(t, verifySignature(signed, signature))
}

func TestVerifySignatureFail(t *testing.T) {
	signed := strings.NewReader("dummy")
	signature, err := os.Open("testdata/nomad_0.8.4_SHA256SUMS.sig")
	require.NoError(t, err)

	require.Error(t, verifySignature(signed, signature))
}
