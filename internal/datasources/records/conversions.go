package records

import "applicationDesignTest/internal/business/domains"

func (o *Order) ToDomain() *domains.OrderDomain {
	return &domains.OrderDomain{
		HotelID:   o.HotelID,
		RoomID:    o.RoomID,
		UserEmail: o.UserEmail,
		From:      o.From,
		To:        o.To,
	}
}

func FromOrderDomain(o *domains.OrderDomain) *Order {
	return &Order{
		HotelID:   o.HotelID,
		RoomID:    o.RoomID,
		UserEmail: o.UserEmail,
		From:      o.From,
		To:        o.To,
	}
}

func (r *Room) ToDomain() *domains.RoomDomain {
	return &domains.RoomDomain{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
	}
}

func FromRoomDomain(r *domains.RoomDomain) *Room {
	return &Room{
		HotelID: r.HotelID,
		RoomID:  r.RoomID,
	}
}
