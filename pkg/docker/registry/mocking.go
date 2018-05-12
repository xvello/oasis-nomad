package registry

import (
	"errors"
	"fmt"
)

// MockedRegistry holds data for the mock
type MockedRegistry struct {
	entries map[string]mockedRepo
}

type mockedRepo map[string]string

// SetupMock returns a new MockedRegistry, all subsequent
// calls to the registry package will use this mock until
// ResetMock is caled
func SetupMock() (*MockedRegistry, error) {
	m := &MockedRegistry{
		entries: make(map[string]mockedRepo),
	}
	mockedSource = m

	return m, nil
}

// ResetMock removes the current mock, subsequent calls to the
// registry packakge will use the network
func ResetMock() error {
	mockedSource = nil
	return nil
}

// AddTag registers a new tag in the mocked registry
func (m *MockedRegistry) AddTag(repo, ref, digest string) error {
	if repo == "" {
		return errors.New("empty repo name")
	}
	if ref == "" {
		return errors.New("empty ref name")
	}
	if digest == "" {
		return errors.New("empty digest")
	}

	r, found := m.entries[repo]
	if !found {
		r = make(mockedRepo)
		m.entries[repo] = r
	}
	r[ref] = digest
	return nil
}

// Digest is part of the Source interface
func (m *MockedRegistry) Digest(repo, ref string) (string, error) {
	r, found := m.entries[repo]
	if !found {
		return "", fmt.Errorf("repo %s not found", repo)
	}
	digest, found := r[ref]
	if !found {
		return "", fmt.Errorf("ref %s not found for repo %s", ref, repo)
	}

	return digest, nil
}

// Tags is part of the Source interface
func (m *MockedRegistry) Tags(repo string) ([]string, error) {
	r, found := m.entries[repo]
	if !found {
		return nil, fmt.Errorf("repo %s not found", repo)
	}
	tags := make([]string, 0, len(r))
	for tag := range r {
		tags = append(tags, tag)
	}
	return tags, nil
}
