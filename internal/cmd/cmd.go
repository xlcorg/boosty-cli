package cmd

import (
	"boosty/internal/cmd/get"
	"boosty/internal/cmd/info"
	"boosty/internal/cmd/posts"
	"boosty/internal/storage"
	"github.com/spf13/cobra"
)

var author string
var token string

func NewDefaultBoostyCommand(version string, store storage.Storage) *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:     "boosty",
		Version: version,
	}

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Bearer token")
	rootCmd.PersistentFlags().StringVarP(&author, "blog", "b", "", "Blog name")

	rootCmd.AddCommand(get.NewCommand(store))
	rootCmd.AddCommand(info.NewCommand(store))
	rootCmd.AddCommand(posts.NewCommand(store))

	return rootCmd
}
