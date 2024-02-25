package info

import (
	"context"
	"errors"
	"fmt"
	"time"

	"boosty/internal/cmd/flags"
	"boosty/internal/storage"
	"boosty/internal/util"

	"boosty/internal/boosty"
	"github.com/spf13/cobra"
)

var store storage.Storage

const (
	blogNameKey = "blog"
	tokenKey    = "token"
)

func NewCommand(s storage.Storage) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info",
		Short: "Display information about blog.",
		Args:  cobra.NoArgs,
		Run:   executeCommand,
	}

	store = s

	return cmd
}

func executeCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	blogName := flags.GetValue(blogNameKey, cmd, store)
	util.CheckError(util.VerifyName(blogName))

	token := flags.GetValue(tokenKey, cmd, store)
	config := boosty.NewConfig()
	if token != "" {
		config = config.WithToken(token)
	}

	client, err := boosty.NewClientWithConfig(blogName, config)
	if err != nil {
		if errors.Is(err, boosty.ErrUserUnauthorized) {
			store.Delete(tokenKey)
		}
	}
	util.CheckError(err)

	fmt.Printf("Getting information about: %s\n---\n", blogName)
	blog, err := client.GetBlog(ctx)
	if err != nil {
		if errors.Is(err, boosty.ErrUserUnauthorized) {
			fmt.Println("Unauthorized. Token has been expired")
			store.Delete(tokenKey)
			fmt.Println("Token has been removed. Try again")
			return
		}
		util.CheckError(err)
	}

	fmt.Println(blog)
	fmt.Println("---")
}
