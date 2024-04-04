package inmem

import (
	"context"
	"fmt"
	"wordOfWisdom/internal/business/domains"
	"wordOfWisdom/internal/datasources/records"
)

type inmemQuoteRepository struct {
	data map[int]*records.Quote
}

// GetQuoteById implements domains.QuoteRepository.
func (i *inmemQuoteRepository) GetQuoteById(ctx context.Context, id int) (*domains.QuoteDomain, error) {
	quite, ok := i.data[id]
	if !ok {
		return nil, fmt.Errorf("quote with id %v not found", id)
	}

	return quite.ToDomain(), nil
}

// GetQuoteIds implements domains.QuoteRepository.
func (i *inmemQuoteRepository) GetQuoteIds(ctx context.Context) ([]int, error) {
	ids := make([]int, 0)

	for id := range i.data {
		ids = append(ids, id)
	}

	return ids, nil
}

func NewQuoteRepository() domains.QuoteRepository {
	return &inmemQuoteRepository{
		data: initQuotes(),
	}
}

func initQuotes() map[int]*records.Quote {
	quotes := make(map[int]*records.Quote)

	return quotes
}
