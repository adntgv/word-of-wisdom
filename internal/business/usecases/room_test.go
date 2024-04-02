package usecases

import (
	"applicationDesignTest/internal/business/domains"
	"context"
	"errors"
	"testing"
	"time"
)

// mockRoomRepository is a mock implementation of the domains.RoomRepository interface
type mockRoomRepository struct {
	getRoomBookingDaysBetweenFunc func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error)
	bookRoomForDaysBetweenFunc    func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error
}

func (m *mockRoomRepository) GetRoomBookingDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error) {
	return m.getRoomBookingDaysBetweenFunc(ctx, room, from, to)
}

func (m *mockRoomRepository) BookRoomForDaysBetween(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
	return m.bookRoomForDaysBetweenFunc(ctx, room, from, to)
}

func Test_roomUsecase_Book(t *testing.T) {
	ctx := context.Background()
	room := &domains.RoomDomain{HotelID: "1", RoomID: "101"}

	type fields struct {
		repo domains.RoomRepository
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
		setup   func(repo *mockRoomRepository)
		wantErr bool
	}{
		{
			name: "Successful booking",
			fields: fields{
				repo: &mockRoomRepository{},
			},
			args: args{
				ctx:  ctx,
				room: room,
				from: time.Now(),
				to:   time.Now().Add(48 * time.Hour),
			},
			setup: func(repo *mockRoomRepository) {
				repo.getRoomBookingDaysBetweenFunc = func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error) {
					return []string{}, nil
				}
				repo.bookRoomForDaysBetweenFunc = func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) error {
					return nil
				}
			},
			wantErr: false,
		},
		{
			name: "Booking with unavailable dates",
			fields: fields{
				repo: &mockRoomRepository{},
			},
			args: args{
				ctx:  ctx,
				room: room,
				// Assuming you want to test a specific date range, replace YYYY-MM-DD with actual dates
				from: func(s string) time.Time { ts, _ := time.Parse("2006-01-02", s); return ts }("2023-04-09"),
				to:   func(s string) time.Time { ts, _ := time.Parse("2006-01-02", s); return ts }("2023-04-11"),
			},
			setup: func(repo *mockRoomRepository) {
				repo.getRoomBookingDaysBetweenFunc = func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error) {
					// This simulates that the room is unavailable on 2023-04-10
					return []string{"2023-04-10"}, nil
				}
			},
			wantErr: true,
		},
		{
			name: "Repository error on checking availability",
			fields: fields{
				repo: &mockRoomRepository{},
			},
			args: args{
				ctx:  ctx,
				room: room,
				from: time.Now(),
				to:   time.Now().Add(48 * time.Hour),
			},
			setup: func(repo *mockRoomRepository) {
				repo.getRoomBookingDaysBetweenFunc = func(ctx context.Context, room *domains.RoomDomain, from time.Time, to time.Time) ([]string, error) {
					return nil, errors.New("database error")
				}
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup(tt.fields.repo.(*mockRoomRepository))

			r := &roomUsecase{
				repo: tt.fields.repo,
			}
			err := r.Book(tt.args.ctx, tt.args.room, tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("roomUsecase.Book() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
