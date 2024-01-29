package main

import (
	"fmt"
	"os"

	"boosty/internal/cmd"
)

var (
	version = "unknown"
)

func main() {
	command := cmd.NewDefaultBoostyCommand(version)
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
