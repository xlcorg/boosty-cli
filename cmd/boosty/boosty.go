package main

import (
	"fmt"
	"os"

	"boosty/internal/cmd"
)

func main() {
	command := cmd.NewDefaultBoostyCommand()
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
