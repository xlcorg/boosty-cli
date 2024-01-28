package info

import (
	"boosty/internal/boosty"
	"boosty/pkg/util"
	"context"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [blog name]",
		Short: "Display information about a blog.",
		Args:  cobra.NoArgs,
		Run:   executeCommand,
	}

	return cmd
}

func executeCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	blogName, _ := cmd.Flags().GetString("author")
	util.CheckError(util.VerifyName(blogName))

	client, err := boosty.NewClient(blogName)
	util.CheckError(err)

	fmt.Printf("Getting information about: %s\n---\n", blogName)
	blog, err := client.GetBlog(ctx)
	util.CheckError(err)

	fmt.Println(blog)
	fmt.Println("---")
}
