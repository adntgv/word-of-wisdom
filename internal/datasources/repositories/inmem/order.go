package inmem

import (
	"applicationDesignTest/internal/business/domains"
	"applicationDesignTest/internal/datasources/records"
	"context"
	"fmt"
)

type inmemOrderRepository struct {
	data map[string]*records.Order
}

// Store implements domains.OrderRepository.
func (i *inmemOrderRepository) Store(ctx context.Context, order *domains.OrderDomain) error {
	record := records.FromOrderDomain(order)

	key := fmt.Sprintf("%v:%v", record.HotelID, record.RoomID)

	i.data[key] = record

	return nil
}

func NewOrderRepository() domains.OrderRepository {
	return &inmemOrderRepository{
		data: make(map[string]*records.Order),
	}
}
