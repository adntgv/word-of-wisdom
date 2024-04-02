package inmem

import (
	"applicationDesignTest/internal/business/domains"
	"applicationDesignTest/internal/datasources/records"
	"context"
	"fmt"
	"time"
)

type inmemRoomRepository struct {
	rooms    map[string]*records.Room
	bookings map[string][]*records.Booking
}

// BookRoomForDaysBetween implements domains.RoomRepository.
func (i *inmemRoomRepository) BookRoomForDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	record := records.FromRoomDomain(room)

	if err := i.exists(record); err != nil {
		return err
	}

	key := record.GetKey()

	bookings := i.bookings[key]
	if bookings == nil {
		bookings = make([]*records.Booking, 0)
	}

	bookings = append(bookings, &records.Booking{
		HotelID: record.HotelID,
		RoomID:  record.RoomID,
		From:    from,
		To:      to,
	})

	i.bookings[key] = bookings

	return nil
}

// GetRoomBookingDaysBetween implements domains.RoomRepository.
func (i *inmemRoomRepository) GetRoomBookingDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error) {
	record := records.FromRoomDomain(room)

	// Check if the room exists in the repository.
	if err := i.exists(record); err != nil {
		return nil, err
	}

	key := record.GetKey()
	bookedDays, ok := i.bookings[key]
	if !ok {
		// If there are no bookings for the room, return an empty slice.
		return []string{}, nil
	}

	result := make([]string, 0)

	// Iterate over each booking to check if it falls within the specified date range.
	for _, booking := range bookedDays {
		// Check if the booking overlaps with the given date range.
		if booking.From.Before(to) && booking.To.After(from) {
			// Generate all dates within the overlap range.
			current := booking.From
			for current.Before(booking.To) && current.Before(to) {
				if current.After(from) || current.Equal(from) {
					// Format and append the date to the result slice if it's within the range.
					result = append(result, current.Format(time.RFC822))
				}
				// Move to the next day.
				current = current.AddDate(0, 0, 1)
			}
		}
	}

	return result, nil
}

func (i *inmemRoomRepository) exists(room *records.Room) error {
	key := room.GetKey()

	if _, exists := i.rooms[key]; !exists {
		return fmt.Errorf("room with hotel id '%v' and room id '%v' does not exist", room.HotelID, room.RoomID)
	}

	return nil
}

func NewRoomRepository() domains.RoomRepository {
	rooms, bookings := initStorages()

	return &inmemRoomRepository{
		bookings: bookings,
		rooms:    rooms,
	}
}

func initStorages() (rooms map[string]*records.Room, bookings map[string][]*records.Booking) {
	rooms = make(map[string]*records.Room)
	bookings = make(map[string][]*records.Booking)

	// usually this is done in migrations

	for hotelId := 1; hotelId <= 5; hotelId++ {
		for roomId := 1; roomId <= 5; roomId++ {
			room := &records.Room{
				HotelID: fmt.Sprint(hotelId),
				RoomID:  fmt.Sprint(roomId),
			}

			key := room.GetKey()
			rooms[key] = room
		}
	}

	for _, room := range rooms {
		key := room.GetKey()

		today := time.Now()

		from := today.Add(time.Hour * 24)
		to := from.Add(time.Hour * 24)

		booking := &records.Booking{
			HotelID: room.HotelID,
			RoomID:  room.RoomID,
			From:    from,
			To:      to,
		}

		bookings[key] = append(bookings[key], booking)
	}

	return rooms, bookings
}
