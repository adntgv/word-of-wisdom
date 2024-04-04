package main

import (
	"log"
	"wordOfWisdom/config"
	"wordOfWisdom/pkg/tcp/server"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Panicf("could initiate config: %v", err)
	}

	server, err := server.NewServer(cfg)
	if err != nil {
		log.Panicf("could not create server app: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatal("erro while running server app", err)
	}
}
