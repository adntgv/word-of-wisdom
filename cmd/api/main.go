package main

import (
	"applicationDesignTest/cmd/api/server"
	"log"
)

func main() {
	api, err := server.NewApi()
	if err != nil {
		log.Panicf("could not create api server app: %v", err)
	}

	if err := api.Run(); err != nil {
		log.Fatal("erro while running api server app", err)
	}
}
