package usecases

import (
	"context"
	"fmt"
	"math/rand"
	"wordOfWisdom/internal/business/domains"
)

type quoteUsecase struct {
	repo domains.QuoteRepository
}

// GetRandomQuote implements domains.QuoteUsecase.
func (r *quoteUsecase) GetRandomQuote(ctx context.Context) (*domains.QuoteDomain, error) {
	quoteIds, err := r.repo.GetQuoteIds(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get quote count: %v", err)
	}

	idx := rand.Intn(len(quoteIds))

	id := quoteIds[idx]

	return r.repo.GetQuoteById(ctx, id)
}

func NewQuoteUsecase(repo domains.QuoteRepository) domains.QuoteUsecase {
	return &quoteUsecase{
		repo: repo,
	}
}
