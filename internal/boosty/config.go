package boosty

import (
	"time"
)

const (
	defaultEndpoint     = "https://api.boosty.to/v1"
	defaultRetryTimeout = 1 * time.Second
	defaultRetryCount   = 5
)

type Config struct {
	endpoint     string
	retryTimeout time.Duration
	retryCount   int
	debug        bool
}

func NewConfig() Config {
	return Config{
		endpoint:     defaultEndpoint,
		retryTimeout: defaultRetryTimeout,
		retryCount:   defaultRetryCount,
		debug:        false,
	}
}

// WithEndpoint sets boosty web api baseAPI
//
// Default value is "https://api.boosty.to/v1"
func (c Config) WithEndpoint(addr string) Config {
	c.endpoint = addr
	return c
}

// WithRequestTimeout sets the maximum duration of the request
//
// Default value is 1 second
func (c Config) WithRequestTimeout(d time.Duration) Config {
	c.retryTimeout = d
	return c
}

func (c Config) WithRetryCount(val int) Config {
	c.retryCount = val
	return c
}

func (c Config) WithDebugEnable() Config {
	c.debug = true
	return c
}
