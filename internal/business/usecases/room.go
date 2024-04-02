package usecases

import (
	"applicationDesignTest/internal/business/domains"
	"context"
	"fmt"
	"time"
)

type roomUsecase struct {
	repo domains.RoomRepository
}

// Book implements domains.RoomUsecase.
func (r *roomUsecase) Book(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	unavailableDates, err := r.repo.GetRoomBookingDaysBetween(ctx, room, from, to)
	if err != nil {
		return fmt.Errorf("room with hotel id \"%v\" and room id \"%v\" cannot be booked: %v", room.HotelID, room.RoomID, err)
	}

	if len(unavailableDates) > 0 {
		return fmt.Errorf("room with hotel id \"%v\" and room id \"%v\" cannot be booked, following dates are unavailable: %v", room.HotelID, room.RoomID, unavailableDates)
	}

	return r.repo.BookRoomForDaysBetween(ctx, room, from, to)
}

func NewRoomUsecase(repo domains.RoomRepository) domains.RoomUsecase {
	return &roomUsecase{
		repo: repo,
	}
}
