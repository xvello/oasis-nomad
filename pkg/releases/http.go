package releases

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// get returns the body of the HTTP request if successful,
// caller needs to explicitly close the body if returned.
func get(urlParts ...string) (io.ReadCloser, error) {
	parts := []string{releaseBaseURL}
	parts = append(parts, urlParts...)
	url := strings.Join(parts, "/")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unexpected HTTP status: %s", resp.Status)
	}

	return resp.Body, nil
}
