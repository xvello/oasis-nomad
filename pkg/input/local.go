package input

import (
	"os"
	"path/filepath"

	"gopkg.in/src-d/go-billy.v4/osfs"
)

// Local creates a Target for files on the local filesystem
func Local(selectors []string) (*Targets, error) {
	var patched []string
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for _, s := range selectors {
		if !filepath.IsAbs(s) {
			s = filepath.Join(cwd, s)
		}
		patched = append(patched, s)
	}

	return &Targets{
		filesystem: osfs.New(string(filepath.Separator)),
		selectors:  patched,
	}, nil
}
