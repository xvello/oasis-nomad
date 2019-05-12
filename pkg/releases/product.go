package releases

import (
	"encoding/json"
	"fmt"
	"io"
	"sort"

	version "github.com/hashicorp/go-version"
)

// GetReleases retrieves the product information from releases.hashicorp.com
func GetReleases(product string) (*ReleaseProduct, error) {
	body, err := get(product, "index.json")
	defer body.Close()
	if err != nil {
		return nil, err
	}
	return parseReleases(body)
}

func parseReleases(r io.Reader) (*ReleaseProduct, error) {
	p := &ReleaseProduct{}
	decoder := json.NewDecoder(r)
	err := decoder.Decode(p)
	if err != nil && err != io.EOF {
		return nil, err
	}
	return p, nil
}

// Latest returns the latest (stable or not depending on parameter) available version
func (p *ReleaseProduct) Latest(stable bool) (*ReleaseVersion, error) {
	versions, err := p.ListVersions(stable)
	if err != nil {
		return nil, err
	}

	latest := versions[len(versions)-1].Original()
	v, ok := p.Versions[latest]
	if !ok {
		return nil, fmt.Errorf("No details for version %s of %s", latest, p.Name)
	}
	return v, nil
}

// Find returns the version matching the given string
func (p *ReleaseProduct) Find(version string) (*ReleaseVersion, error) {
	v, found := p.Versions[version]
	if !found {
		return nil, fmt.Errorf("Cannot find version %s", version)
	}
	return v, nil
}

// ListVersions returns all the available versions
func (p *ReleaseProduct) ListVersions(stable bool) ([]*version.Version, error) {
	if p.Versions == nil {
		return nil, fmt.Errorf("Invalid object for %s", p.Name)
	}

	var versions []*version.Version
	for _, raw := range p.Versions {
		v, err := version.NewVersion(raw.Version)
		if err != nil {
			continue
		}
		if stable && v.Prerelease() != "" {
			continue
		}
		versions = append(versions, v)
	}

	if len(versions) == 0 {
		if stable {
			return nil, fmt.Errorf("No stable version for %s", p.Name)
		}
		return nil, fmt.Errorf("No version for %s", p.Name)
	}

	sort.Sort(version.Collection(versions))
	return versions, nil
}
