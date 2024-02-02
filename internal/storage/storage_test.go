package storage_test

import (
	"boosty/internal/storage"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorage(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		s, err := storage.New("settings.yaml")
		assert.NoError(t, err)
		s.Set("token", "42")

		assert.Equal(t, "42", s.Get("token"))
	})

	t.Run("just works", func(t *testing.T) {
		{
			s, err := storage.New("settings.yaml")
			assert.NoError(t, err)
			defer s.Close()
			s.Set("token", "42")
		}
		{
			s, err := storage.New("settings.yaml")
			assert.NoError(t, err)
			defer s.Close()
			assert.Equal(t, "42", s.Get("token"))

		}

	})
}
