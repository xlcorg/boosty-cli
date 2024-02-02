package endpoint

import (
	"fmt"
	"log"
)

//go:generate stringer -linecomment -type=Endpoint
type Endpoint int

const (
	GetBlog  = Endpoint(iota)
	GetPosts // post/
)

type Config struct {
	dict map[Endpoint]string
}

func NewEndpointConfig(base string, blogName string) Config {
	return Config{
		dict: map[Endpoint]string{
			GetBlog:  fmt.Sprintf("%s/blog/%s", base, blogName),
			GetPosts: fmt.Sprintf("%s/blog/%s/%s", base, blogName, GetPosts.String()),
		},
	}
}

func (c Config) Get(e Endpoint) string {
	val, ok := c.dict[e]
	if ok {
		return val
	}

	log.Fatal("not found endpoint")
	return ""
}
