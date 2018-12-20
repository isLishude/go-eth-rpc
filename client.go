package ethrpc

import (
	"net/http"

	"github.com/islishude/go-jsonrpc2"
)

// Denomination Units
const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)

// Ethereum is RPC object
type Ethereum struct {
	*jr2.Client
}

// SetSerName set service name
func (c *Ethereum) SetSerName(newName string) *Ethereum {
	c.SerName = newName
	return c
}

// NewEthereum is
func NewEthereum(url, authUser, authPass string, client *http.Client) *Ethereum {
	if client == nil {
		client = jr2.DefaultClient
	}
	return &Ethereum{
		&jr2.Client{
			SerName:    "Ethereum",
			BaseURL:    url,
			HTTPClient: client,
			Username:   authUser,
			Password:   authPass,
		},
	}
}

// DefaultClient is
var DefaultClient = &Ethereum{
	&jr2.Client{
		SerName:    "Ethereum",
		BaseURL:    "http://127.0.0.1:8545",
		HTTPClient: http.DefaultClient,
	},
}
