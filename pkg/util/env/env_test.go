package env_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"storage-api/internal/util/env"
)

func TestTryGetBool(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		val, err := env.TryGetBool("NONE")
		assert.Equal(t, false, val)

		expected := `Environment variable not found: "NONE"`
		assert.Equal(t, expected, err.Error())
	})
}
