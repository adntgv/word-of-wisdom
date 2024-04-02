package records

import (
	"fmt"
	"time"
)

type Room struct {
	HotelID string
	RoomID  string
}

func (r Room) GetKey() string {
	return fmt.Sprintf("%v:%v", r.HotelID, r.RoomID)
}

type Booking struct {
	HotelID string
	RoomID  string
	From    time.Time
	To      time.Time
}
