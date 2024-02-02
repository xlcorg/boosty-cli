package main

import (
	"boosty/internal/logger"
	"boosty/internal/storage"
	"boosty/internal/util/env"
	"fmt"
	"log"
	"os"

	"boosty/internal/cmd"
)

var (
	version = "unknown"
)

func main() {
	environment := env.GetStringOrDefault("ENV", "Prod")
	logger.InitLocal(environment == "Dev")
	defer logger.Sync()

	store, err := storage.New("~/.config/boosty/settings.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer func(store storage.StorageCloser) {
		err := store.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(store)

	command := cmd.NewDefaultBoostyCommand(version, store)
	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
