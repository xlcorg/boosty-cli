package cmd

import (
	"boosty/internal/cmd/get"
	"boosty/internal/cmd/info"
	"boosty/internal/cmd/posts"
	"github.com/spf13/cobra"
)

var author string
var token string

func NewDefaultBoostyCommand() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "boosty"}

	rootCmd.PersistentFlags().StringVarP(&token, "token", "t", "", "Bearer token")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "", "Blog name")
	_ = rootCmd.MarkPersistentFlagRequired("author")

	rootCmd.AddCommand(get.NewCommand())
	rootCmd.AddCommand(info.NewCommand())
	rootCmd.AddCommand(posts.NewCommand())

	return rootCmd
}
