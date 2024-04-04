package main

import (
	"log"
	"wordOfWisdom/internal/tcp/handlers"
	"wordOfWisdom/pkg/tcp/server"
)

func main() {
	handler := handlers.NewConnectionHandler().Handle

	server, err := server.NewServer(handler)
	if err != nil {
		log.Panicf("could not create  server app: %v", err)
	}

	if err := server.Run(); err != nil {
		log.Fatal("erro while running  server app", err)
	}
}
