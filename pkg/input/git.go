package input

import (
	"gopkg.in/src-d/go-billy.v4/memfs"
	git "gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"gopkg.in/src-d/go-git.v4/storage/memory"
)

// FromGit creates a Target for files in a git repository
func FromGit(url string, selectors []string) (*Targets, error) {
	fs := memfs.New()
	_, err := git.Clone(memory.NewStorage(), fs, &git.CloneOptions{
		Depth:         1,
		ReferenceName: plumbing.HEAD,
		SingleBranch:  true,
		URL:           url,
	})
	if err != nil {
		return nil, err
	}
	return &Targets{
		filesystem: fs,
		selectors:  selectors,
	}, nil
}
