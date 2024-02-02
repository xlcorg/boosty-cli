package posts

import (
	"boosty/internal/cmd/flags"
	"boosty/internal/storage"
	"boosty/internal/util"
	"context"
	"fmt"
	"time"

	"boosty/internal/boosty"
	"github.com/spf13/cobra"
)

var (
	store      storage.Storage
	postsLimit int
)

const (
	blogNameKey = "blog"
	tokenKey    = "token"
)

func NewCommand(s storage.Storage) *cobra.Command {
	store = s
	var cmd = &cobra.Command{
		Use:     "posts [blogName]",
		Short:   "Display posts",
		Args:    cobra.NoArgs,
		Run:     executePostsCommand,
		Aliases: []string{"ls", "last"},
	}

	cmd.Flags().IntVarP(&postsLimit, "limit", "l", 3, "limit of posts")

	return cmd
}

func executePostsCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if cmd.CalledAs() == "last" {
		postsLimit = 1
	}

	blogName := flags.GetValue(blogNameKey, cmd, store)
	util.CheckError(util.VerifyName(blogName))

	token := flags.GetValue(tokenKey, cmd, store)
	config := boosty.NewConfig()
	if token != "" {
		config = config.WithToken(token)
	}

	client, err := boosty.NewClientWithConfig(blogName, config)
	util.CheckError(err)

	fmt.Printf("Getting %d posts: %s\n---\n", postsLimit, blogName)

	posts, err := client.GetPosts(ctx, postsLimit)
	util.CheckError(err)

	for _, post := range posts {
		fmt.Println(post)
		fmt.Println("---")
	}
}
