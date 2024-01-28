package boosty

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		c := NewConfig()
		assert.Equal(t, defaultEndpoint, c.endpoint)
		assert.Equal(t, defaultRetryTimeout, c.retryTimeout)
	})

	t.Run("custom baseAPI", func(t *testing.T) {
		c := NewConfig()
		expected := "https://custom-endpoint.to/v1/api"
		c = c.WithEndpoint(expected)

		assert.Equal(t, expected, c.endpoint)
	})

	t.Run("custom retryTimeout", func(t *testing.T) {
		c := NewConfig()
		expected := 42 * time.Second
		c = c.WithRequestTimeout(expected)

		assert.Equal(t, expected, c.retryTimeout)
	})
}
