package main

import (
	"fmt"
	"log"
	"wordOfWisdom/config"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Panicf("could initiate config: %v", err)
	}

	address := fmt.Sprintf("%v:%v", cfg.Host, cfg.Port)
	client, err := newClient(address)
	if err != nil {
		log.Fatalln(err)
	}

	challenge, prefix, err := client.getChallange()
	if err != nil {
		log.Fatalln(err)
	}

	nonce := client.solveChallenge(challenge, prefix)

	quote, err := client.getQuote(nonce)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Received quote:", quote)
}
