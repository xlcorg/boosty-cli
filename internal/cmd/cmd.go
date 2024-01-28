package cmd

import (
	"boosty/internal/cmd/get"
	"boosty/internal/cmd/info"
	"boosty/internal/cmd/posts"
	"github.com/spf13/cobra"
)

var author string

func NewDefaultBoostyCommand() *cobra.Command {
	var rootCmd = &cobra.Command{Use: "boosty"}

	// add config flag
	rootCmd.PersistentFlags().StringVarP(
		&author,
		"author",
		"a",
		"",
		"blog author")
	_ = rootCmd.MarkPersistentFlagRequired("author")

	//viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))

	rootCmd.AddCommand(get.NewCommand())
	rootCmd.AddCommand(info.NewCommand())
	rootCmd.AddCommand(posts.NewCommand())

	return rootCmd
}
