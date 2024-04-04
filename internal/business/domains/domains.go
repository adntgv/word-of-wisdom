package domains

import (
	"context"
)

type ChallangeDomain struct {
	Challange string
	Nonce     string
	Prefix    string
}

type ChallangeUsecase interface {
	Generate(ctx context.Context) (*ChallangeDomain, error)
	Validate(ctx context.Context, challange *ChallangeDomain) error
}

type ChallangeRepository interface {
}

type QuoteDomain struct {
	ID   int
	Text string
}

type QuoteUsecase interface {
	GetRandomQuote(ctx context.Context) (*QuoteDomain, error)
}

type QuoteRepository interface {
	GetQuoteIds(ctx context.Context) ([]int, error)
	GetQuoteById(ctx context.Context, id int) (*QuoteDomain, error)
}
