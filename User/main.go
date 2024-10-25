package main

import (
	"enhanced-notes/config"
	"enhanced-notes/src/api"
	"log"
)

func main() {
	cfg, err := config.SetupEnv()

	if err != nil {
		log.Fatalf("config file is not loaded %v", err)
	}

	api.StartServer(cfg)
}