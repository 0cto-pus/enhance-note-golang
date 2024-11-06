package main

import (
	"enhance-notes-suggestion/config"
	"enhance-notes-suggestion/src/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded %v", err)
	}

	api.StartServer(cfg)
}  