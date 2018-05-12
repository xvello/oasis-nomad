package nomad

import "github.com/hashicorp/nomad/api"

// Client wraps around the nomad api client and adds new methods
type Client struct {
	cli *api.Client
}
