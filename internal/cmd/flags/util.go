package flags

import (
	"boosty/internal/storage"
	"github.com/spf13/cobra"
)

func GetValue(key string, cmd *cobra.Command, store storage.Storage) string {
	value, _ := cmd.Flags().GetString(key)
	if value == "" {
		value = store.Get(key)
	} else {
		store.Set(key, value)
	}
	return value
}
