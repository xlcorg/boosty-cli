package boosty

import (
	"fmt"
	"net/url"

	"boosty/internal/clients/boosty/endpoint"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	blogName  string
	baseAPI   *url.URL
	endpoints endpoint.Config
	http      *resty.Client
	debug     bool
}

func NewClient(blogName string) (*Client, error) {
	return NewClientWithConfig(blogName, NewConfig())
}

func NewClientWithConfig(blogName string, config Config) (*Client, error) {
	baseURL, err := url.Parse(config.endpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse endpoing: %w", err)
	}

	return &Client{
		debug:     config.debug,
		blogName:  blogName,
		baseAPI:   baseURL,
		endpoints: endpoint.NewEndpointConfig(baseURL.String(), blogName),
		http: resty.New().
			SetDebug(false).
			SetRetryCount(config.retryCount).
			SetRetryMaxWaitTime(config.retryTimeout).
			SetDisableWarn(true),
	}, nil
}

func (c *Client) getEndpoint(path endpoint.Endpoint) string {
	return c.endpoints.Get(path)
}
