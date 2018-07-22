package releases

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

var stableVersionRegexp = regexp.MustCompile(`^[0-9]+\.[0-9]+\.[0-9]+$`)

// Download gets, verifies and uncompress a given release version
func (v *ReleaseVersion) Download(destination, buildOs, buildArch string) error {
	tmpDir, err := ioutil.TempDir("", "oasis-download-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tmpDir)

	zipPath := filepath.Join(tmpDir, "download.zip")
	err = v.downloadZip(zipPath, buildOs, buildArch)
	if err != nil {
		return err
	}

	return uncompressZip(zipPath, v.Name, destination)
}

func (v *ReleaseVersion) downloadZip(destination, buildOs, buildArch string) error {
	// Find the correct build
	var build *ReleaseBuild
	for _, b := range v.Builds {
		if b.OS == buildOs && b.Arch == buildArch {
			build = b
			break
		}
	}
	if build == nil {
		return fmt.Errorf("Could not find a build for %s - %s", buildOs, buildArch)
	}

	// Get file sha
	sha, err := v.getSha(build.Filename)
	if err != nil {
		return err
	}

	// Open download stream
	contents, err := v.get(build.Filename)
	defer contents.Close()
	if err != nil {
		return err
	}

	// Write to output file and compute sha
	output, err := os.Create(destination)
	if err != nil {
		return err
	}
	hasher := sha256.New()
	tee := io.TeeReader(contents, hasher)
	_, err = io.Copy(output, tee)
	output.Close()
	if err != nil {
		return err
	}

	// Verify checksum
	checksum := fmt.Sprintf("%x", hasher.Sum(nil))
	if checksum != string(sha) {
		return fmt.Errorf("Invalid checksum: %s, expected %s", checksum, string(sha))
	}
	return nil
}

func (v *ReleaseVersion) get(file string) (io.ReadCloser, error) {
	return get(v.Name, v.Version, file)
}

// getSha retrieves the sha256 of a given file after validating the gpg signature
func (v *ReleaseVersion) getSha(file string) ([]byte, error) {
	signature, err := v.get(v.Signature)
	defer signature.Close()
	if err != nil {
		return nil, err
	}
	sha, err := v.get(v.Shasums)
	defer sha.Close()
	if err != nil {
		return nil, err
	}
	var shaContents bytes.Buffer
	tee := io.TeeReader(sha, &shaContents)

	err = verifySignature(tee, signature)
	if err != nil {
		return nil, err
	}
	return extractSha(shaContents, file)
}

// extractSha extracts the sha checksum for a given bytes buffer
func extractSha(contents bytes.Buffer, file string) ([]byte, error) {
	newline := byte('\n')
	fileBytes := []byte(file)

	for {
		line, err := contents.ReadBytes(newline)
		if err != nil {
			return nil, err
		}
		if bytes.Contains(line, fileBytes) {
			return bytes.Fields(line)[0], nil
		}
	}
}
