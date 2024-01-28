package get

import (
	"context"
	"fmt"
	"time"

	"boosty/internal/clients/boosty"
	"github.com/spf13/cobra"
)

var postsLimit int

func newCmdGetPosts() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "posts [blogName]",
		Short: "Display posts",
		Args:  cobra.MinimumNArgs(0),
		Run:   executePostsCommand,
	}

	cmd.Flags().IntVarP(&postsLimit, "limit", "l", 3, "limit of posts")

	return cmd
}

func executePostsCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	checkError(verify(cmd, args))

	blogName := args[0]
	client, err := boosty.NewClient(blogName)
	checkError(err)

	fmt.Printf("Getting %d posts: %s\n---\n", postsLimit, blogName)

	posts, err := client.GetPosts(ctx, postsLimit)
	checkError(err)

	for _, post := range posts {
		fmt.Println(post)
		fmt.Println("---")
	}
}
