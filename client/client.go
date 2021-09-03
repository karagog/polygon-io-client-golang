// Package client is a client object for accessing Polygon-IO.
package client

// Client is a client object. Create with New().
type Client struct {
	apiKey string
}

// New instantiates a new Client object.
func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}
