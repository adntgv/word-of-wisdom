package handlers

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"strings"
	"wordOfWisdom/internal/business/domains"
	"wordOfWisdom/internal/business/usecases"
	"wordOfWisdom/pkg/challanger"
)

type connectionHandler struct {
	quotes     domains.QuoteUsecase
	challanges domains.ChallangeUsecase
}

func NewConnectionHandler(repo domains.QuoteRepository, chal challanger.Challanger) *connectionHandler {
	return &connectionHandler{
		quotes:     usecases.NewQuoteUsecase(repo),
		challanges: usecases.NewChallangeUsecase(chal),
	}
}

func (h *connectionHandler) Handle(conn net.Conn) {
	defer conn.Close()
	ctx := context.Background()

	challenge, err := h.challanges.Generate(ctx)
	if err != nil {
		log.Printf("failed to generate challange: %v\n", err)
		if err1 := h.send(conn, "failed to generate challange\n"); err1 != nil {
			log.Printf("failed to send error response %v: %v\n", conn.RemoteAddr(), err1)
		}
		return
	}

	if err := h.send(conn, challenge.Challange); err != nil {
		log.Printf("failed to send challange to connection %v: %v\n", conn.RemoteAddr(), err)
		if err1 := h.send(conn, "failed to send challange\n"); err1 != nil {
			log.Printf("failed to send error response %v: %v\n", conn.RemoteAddr(), err1)
		}
		return
	}

	challenge.Nonce, err = h.getNonce(conn)
	if err != nil {
		log.Printf("failed to fetch nonce from %v: %v\n", conn.RemoteAddr(), err)
		if err1 := h.send(conn, "failed to fetch nonce\n"); err1 != nil {
			log.Printf("failed to send error response %v: %v\n", conn.RemoteAddr(), err1)
		}

		return
	}

	// Validate the PoW
	if err := h.challanges.Validate(ctx, challenge); err != nil {
		log.Printf("error for challenge %v validation: %v\n", challenge, err)
		fmt.Fprintf(conn, "did not solve the challenge\n")
		return
	}

	// Send the quote to the client
	quote, err := h.quotes.GetRandomQuote(ctx)
	if err != nil {
		log.Printf("failed to fetch a quote: %v\n", err)
		if err1 := h.send(conn, "failed to fetch a quote\n"); err1 != nil {
			log.Printf("failed to send error response %v: %v\n", conn.RemoteAddr(), err1)
		}
		return
	}

	if err := h.send(conn, quote.Text); err != nil {
		log.Printf("failed to send a quote to %v: %v\n", conn.RemoteAddr(), err)
		fmt.Fprintf(conn, "failed to send a quote\n")
		return
	}
}

func (h *connectionHandler) getNonce(conn net.Conn) (string, error) {
	// Read the response from the client
	nonce, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("could not read from connection %v: %v", conn, err)
	}

	return strings.TrimSpace(nonce), nil
}

func (h *connectionHandler) send(conn net.Conn, data string) error {
	// Send the data to the client
	_, err := fmt.Fprintf(conn, "%s\n", data)

	return err
}
