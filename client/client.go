// Package client is a client object for accessing Polygon-IO.
package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	addr = "https://api.polygon.io"
)

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

type TickersOptions struct {
	Market  string // stocks, crypto, fx
	Symbols []string
	Search  string
}

const tickerPageLimit = 50

func (c *Client) Tickers(ctx context.Context, opt *TickersOptions) ([]byte, error) {
	r, err := http.NewRequestWithContext(ctx, "GET",
		fmt.Sprintf("%s/v3/reference/tickers"+
			"?apiKey=%s"+
			"&market=%s"+
			"&limit=%d"+
			"&ticker=%s"+
			"&search=%s"+
			"&sort=ticker",
			addr,
			c.apiKey,
			opt.Market,
			tickerPageLimit,
			strings.Join(opt.Symbols, ","),
			opt.Search,
		), nil)
	if err != nil {
		return nil, err
	}
	return doRequest(r)
}

// Runs the request against the server and returns the raw response JSON.
func doRequest(r *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status (%d): %s", resp.StatusCode, body)
	}
	return body, nil
}
