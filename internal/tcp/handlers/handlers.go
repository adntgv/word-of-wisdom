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
)

type connectionHandler struct {
	quotes domains.QuoteUsecase
}

func NewConnectionHandler(repo domains.QuoteRepository) *connectionHandler {
	return &connectionHandler{
		quotes: usecases.NewQuoteUsecase(repo),
	}
}

func (h *connectionHandler) Handle(conn net.Conn) {
	defer conn.Close()
	ctx := context.Background()

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
