package get

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"boosty/internal/clients/boosty"
	"github.com/spf13/cobra"
)

func newCmdGetInfo() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "info [blog name]",
		Short: "Display information about a blog.",
		Args:  cobra.MinimumNArgs(1),
		Run:   executeCommand,
	}

	return cmd
}

func executeCommand(cmd *cobra.Command, args []string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	checkError(verify(cmd, args))

	blogName := args[0]
	client, err := boosty.NewClientWithConfig(blogName, boosty.NewConfig().WithDebugEnable())
	checkError(err)

	fmt.Printf("Getting information about: %s\n---\n", blogName)
	blog, err := client.GetBlog(ctx)
	checkError(err)

	fmt.Println(blog)
	fmt.Println("---")
}

func verify(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("invalid arguments")
	}
	_, err := verifyName(args[0])
	if err != nil {
		return err
	}

	return nil
}

func verifyName(name string) (string, error) {
	blogName := strings.TrimSpace(name)
	if len(blogName) == 0 {
		return "", fmt.Errorf("name must be specified")
	}

	return blogName, nil
}

func checkError(err error) {
	if err == nil {
		return
	}

	msg := fmt.Sprintf("Error: %s", err.Error())
	fatal(msg, 1)
}

func fatal(msg string, code int) {
	if len(msg) > 0 {
		// add newline if needed
		if !strings.HasSuffix(msg, "\n") {
			msg += "\n"
		}

		fmt.Fprint(os.Stderr, msg)
	}
	os.Exit(code)
}
