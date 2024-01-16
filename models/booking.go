package models

import (
	"hotel-back-v1/db"
	"time"
)

type Booking struct {
	ID        int64
	RoomID    int64   `binding:"required"`
	RoomPrice float64 `binding:"required"`
	Name      string  `binding:"required"`
	Email     string  `binding:"required"`
	Phone     string  `binding:"required"`
	Message   string  `binding:"required"`
	StartDate string  `binding:"required"`
	EndDate   string  `binding:"required"`
	Processed bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (b *Booking) InsertBooking() error {
	query := `
		INSERT INTO bookings
			(room_id, room_price, name, email, phone, message, start_date, end_date, processed, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.RoomID, b.RoomPrice, b.Name, b.Email, b.Phone, b.Message, b.StartDate, b.EndDate, b.Processed, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func SelectAllBookings() ([]Booking, error) {
	query := `
		SELECT 
			id, room_id, room_price, name, email, phone, message, start_date, end_date, processed, created_at, updated_at
		FROM 
			bookings
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []Booking

	for rows.Next() {
		var booking Booking
		err = rows.Scan(&booking.ID, &booking.RoomID, &booking.RoomPrice, &booking.Name, &booking.Email, &booking.Phone, &booking.Message, &booking.StartDate, &booking.EndDate, &booking.Processed, &booking.CreatedAt, &booking.UpdatedAt)
		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}
