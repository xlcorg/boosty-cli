package cmd

import (
	"boosty/internal/cmd/get"
	"github.com/spf13/cobra"
)

func NewDefaultBoostyCommand() *cobra.Command {

	var rootCmd = &cobra.Command{Use: "boosty"}
	rootCmd.AddCommand(get.NewCommand())

	return rootCmd
}
