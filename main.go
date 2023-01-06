package main

import (
	"github.com/walletkita/wallet-core/app"
	"github.com/walletkita/wallet-core/app/router"
	"github.com/walletkita/wallet-core/internal/config"
	"github.com/walletkita/wallet-core/internal/server"
)

func main() {
	// load config
	cfg, err := config.NewConfig("./config.yml")
	if err != nil {
		app.LOG.Fatalf("Error load config : %v", err)
	}

	// run migrator

	// running instance server
	srv := server.NewServer(*cfg, router.NewRouter())
	srv.Run()
}
