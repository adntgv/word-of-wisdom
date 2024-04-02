package inmem

import (
	"applicationDesignTest/internal/business/domains"
	"applicationDesignTest/internal/datasources/records"
	"context"
	"reflect"
	"testing"
	"time"
)

func Test_inmemRoomRepository_BookRoomForDaysBetween(t *testing.T) {
	type fields struct {
		rooms    map[string]*records.Room
		bookings map[string][]*records.Booking
	}
	type args struct {
		ctx  context.Context
		room *domains.RoomDomain
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				rooms: map[string]*records.Room{
					"1:1": {
						HotelID: "1",
						RoomID:  "1",
					},
				},
				bookings: map[string][]*records.Booking{},
			},
			args: args{
				ctx: context.Background(),
				room: &domains.RoomDomain{
					HotelID: "1",
					RoomID:  "1",
				},
				from: time.Now(),
				to:   time.Now().Add(time.Hour),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inmemRoomRepository{
				rooms:    tt.fields.rooms,
				bookings: tt.fields.bookings,
			}
			if err := i.BookRoomForDaysBetween(tt.args.ctx, tt.args.room, tt.args.from, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("inmemRoomRepository.BookRoomForDaysBetween() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_inmemRoomRepository_GetRoomBookingDaysBetween(t *testing.T) {

	now := time.Now()

	type fields struct {
		rooms    map[string]*records.Room
		bookings map[string][]*records.Booking
	}
	type args struct {
		ctx  context.Context
		room *domains.RoomDomain
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				rooms: map[string]*records.Room{
					"1:1": {
						HotelID: "1",
						RoomID:  "1",
					},
				},
				bookings: map[string][]*records.Booking{
					"1:1": {
						{HotelID: "1", RoomID: "1", From: now.Add(time.Second), To: now.Add(time.Hour)},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				room: &domains.RoomDomain{
					HotelID: "1",
					RoomID:  "1",
				},
				from: now,
				to:   now.Add(time.Hour).Add(time.Second),
			},
			want: []string{
				now.Format(time.RFC822),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &inmemRoomRepository{
				rooms:    tt.fields.rooms,
				bookings: tt.fields.bookings,
			}
			got, err := i.GetRoomBookingDaysBetween(tt.args.ctx, tt.args.room, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("inmemRoomRepository.GetRoomBookingDaysBetween() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inmemRoomRepository.GetRoomBookingDaysBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}
