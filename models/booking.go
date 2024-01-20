package models

import (
	"hotel-back-v1/db"
	"time"
)

type Booking struct {
	ID              int64
	RoomID          string `binding:"required"`
	RoomPrice       int64  `binding:"required"`
	TotalPrice      int64  `binding:"required"`
	TotalBookedDays int64  `binding:"required"`
	Name            string `binding:"required"`
	Email           string `binding:"required"`
	Phone           string `binding:"required"`
	Message         string `binding:"required"`
	StartDate       string `binding:"required"`
	EndDate         string `binding:"required"`
	Processed       bool

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookedRoom struct {
	ID        int64
	RoomID    string `binding:"required"`
	StartDate string `binding:"required"`
	EndDate   string `binding:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type BookedRoomID struct {
	RoomID string `binding:"required"`
}

type RoomAvailability struct {
	RoomID    string `binding:"required"`
	StartDate string `binding:"required"`
	EndDate   string `binding:"required"`
}

func (b *Booking) InsertBooking() error {
	query := `
		INSERT INTO bookings
			(room_id, room_price, total_price, total_booked_days, name, email, phone, message, start_date, end_date, processed, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(b.RoomID, b.RoomPrice, b.TotalPrice, b.TotalBookedDays, b.Name, b.Email, b.Phone, b.Message, b.StartDate, b.EndDate, b.Processed, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (br *BookedRoom) InsertBookedRoom() error {
	query := `
		INSERT INTO bookings_rooms
			(room_id, start_date, end_date, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(br.RoomID, br.StartDate, br.EndDate, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func SelectAllBookings() ([]Booking, error) {
	query := `
		SELECT 
			id, room_id, room_price, total_price, total_booked_days, name, email, phone, message, start_date, end_date, processed, created_at, updated_at
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

		err = rows.Scan(&booking.ID, &booking.RoomID, &booking.RoomPrice, &booking.TotalPrice, &booking.TotalBookedDays, &booking.Name, &booking.Email, &booking.Phone, &booking.Message, &booking.StartDate, &booking.EndDate, &booking.Processed, &booking.CreatedAt, &booking.UpdatedAt)

		if err != nil {
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func SearchAvailabilityByDatesByRoomID(roomID string, startDate string, endDate string) (bool, error) {
	query := `
		SELECT 
			count(id)
		FROM 
			bookings_rooms
		WHERE 
			room_id = ? AND
			(? < end_date AND ? > start_date) OR 
      (start_date = ? OR end_date = ?)
	`

	row := db.DB.QueryRow(query, roomID, startDate, endDate, startDate, endDate)

	var id int64

	err := row.Scan(&id)
	if err != nil {
		return false, err
	}

	return id == 0, nil
}

func SearchSelectedRoomByRoomID(roomID string) ([]BookedRoom, error) {
	query := `
		SELECT 
			id, room_id, start_date, end_date, created_at, updated_at
		FROM 
			bookings_rooms
		WHERE 
			room_id = ?
	`

	rows, err := db.DB.Query(query, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookedRooms []BookedRoom

	for rows.Next() {
		var bookedRoom BookedRoom

		err = rows.Scan(&bookedRoom.ID, &bookedRoom.RoomID, &bookedRoom.StartDate, &bookedRoom.EndDate, &bookedRoom.CreatedAt, &bookedRoom.UpdatedAt)

		if err != nil {
			return nil, err
		}

		bookedRooms = append(bookedRooms, bookedRoom)
	}

	return bookedRooms, nil
}
