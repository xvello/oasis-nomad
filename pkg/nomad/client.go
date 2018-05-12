package nomad

import "github.com/hashicorp/nomad/api"

// Connect creates a new connection to the nomad server
// If conf is nil, the default config from envvars will be used.
func Connect(conf *api.Config) (*Client, error) {
	if conf == nil {
		conf = api.DefaultConfig()
	}

	cli, err := api.NewClient(conf)
	if err != nil {
		return nil, err
	}

	return &Client{cli: cli}, nil
}
