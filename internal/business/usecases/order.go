package usecases

import (
	"applicationDesignTest/internal/business/domains"
	"context"
)

type orderUsecase struct {
	repo domains.OrderRepository
}

// Store implements domains.OrderUsecase.
func (o *orderUsecase) Store(ctx context.Context, order *domains.OrderDomain) error {
	return o.repo.Store(ctx, order)
}

func NewOrderUsecase(repo domains.OrderRepository) domains.OrderUsecase {
	return &orderUsecase{
		repo: repo,
	}
}
