package domains

import (
	"context"
	"time"
)

type OrderDomain struct {
	HotelID   string
	RoomID    string
	UserEmail string
	From      time.Time
	To        time.Time
}

type OrderUsecase interface {
	Store(ctx context.Context, order *OrderDomain) error
}

type OrderRepository interface {
	Store(ctx context.Context, order *OrderDomain) error
}

type RoomDomain struct {
	HotelID string
	RoomID  string
}

type RoomUsecase interface {
	Book(ctx context.Context, room *RoomDomain, from time.Time, to time.Time) error
}

type RoomRepository interface {
	GetRoomBookingDaysBetween(ctx context.Context, room *RoomDomain, from time.Time, to time.Time) ([]string, error)
	BookRoomForDaysBetween(ctx context.Context, room *RoomDomain, from time.Time, to time.Time) error
}
