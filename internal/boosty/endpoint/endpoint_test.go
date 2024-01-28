package endpoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEndpointConfig(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		c := NewEndpointConfig("https://api.boosty.to/v1", "foo")

		assert.Equal(t, "https://api.boosty.to/v1/blog/foo", c.dict[GetBlog])
		assert.Equal(t, "https://api.boosty.to/v1/blog/foo/post", c.dict[GetPosts])
	})
}

func TestConfig_GetEndpoint(t *testing.T) {
	t.Run("just works", func(t *testing.T) {
		c := NewEndpointConfig("https://api.boosty.to/v1", "foo")

		assert.Equal(t, "https://api.boosty.to/v1/blog/foo", c.Get(GetBlog))
		assert.Equal(t, "https://api.boosty.to/v1/blog/foo/post", c.Get(GetPosts))
	})

	t.Run("unknown endpoint", func(t *testing.T) {
		c := NewEndpointConfig("https://api.boosty.to/v1", "foo")
		c.Get(6)
	})

}
