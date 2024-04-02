package requests

import (
	"applicationDesignTest/internal/business/domains"
	"fmt"
	"time"
)

type CreateOrderRequest struct {
	HotelID   string `json:"hotel_id"`
	RoomID    string `json:"room_id"`
	UserEmail string `json:"email"`
	From      string `json:"from"`
	To        string `json:"to"`
}

func (r *CreateOrderRequest) Validate() error {
	if r.HotelID == "" {
		return fmt.Errorf("hotel_id is required")
	}
	if r.RoomID == "" {
		return fmt.Errorf("room_id is required")
	}
	if r.UserEmail == "" {
		return fmt.Errorf("email is required")
	}

	now := time.Now()
	from, err := stringToTime(r.From)
	if err != nil {
		return fmt.Errorf("'from' field error: %v", err)
	}

	if from.Before(now) {
		return fmt.Errorf("'from' field cannot be in the past: %v", r.From)
	}

	to, err := stringToTime(r.To)
	if err != nil {
		return fmt.Errorf("'to' field error: %v", err)
	}

	if !to.After(from) {
		return fmt.Errorf("'to' field cannot before 'from': from: %v, to: %v", r.From, r.To)
	}

	return nil
}

func (r *CreateOrderRequest) ToRoom() *domains.RoomDomain {
	return &domains.RoomDomain{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
	}
}

func (r *CreateOrderRequest) ToOrder() *domains.OrderDomain {
	return &domains.OrderDomain{
		HotelID:   r.HotelID,
		RoomID:    r.RoomID,
		UserEmail: r.UserEmail,
		From:      r.GetFrom(),
		To:        r.GetTo(),
	}
}

func (r *CreateOrderRequest) GetFrom() time.Time {
	ts, _ := stringToTime(r.From)

	return ts
}

func (r *CreateOrderRequest) GetTo() time.Time {
	ts, _ := stringToTime(r.To)

	return ts
}

func stringToTime(t string) (time.Time, error) {
	ts, err := time.Parse(time.RFC3339, t)

	if err != nil {
		return ts, fmt.Errorf("could not parse time: %v", err)
	}

	return ts, nil
}
