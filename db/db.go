package db

import (
	"database/sql"
	"fmt"

	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("sqlite3", "hotel.db")
	if err != nil {
		fmt.Println("Error connecting to database.")
		panic(err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		fmt.Println("Error pinging database.")
		panic(err)
	}

	fmt.Println("Connected to database.")

	createTables()
}

func createTables() {
	createRolesTable := `
		CREATE TABLE IF NOT EXISTS roles (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			role_type TEXT NOT NULL UNIQUE,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
	`
	_, err := DB.Exec(createRolesTable)
	if err != nil {
		fmt.Println("Error creating roles table.")
		panic(err)
	}

	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL,
			role_type TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME	NOT NULL
		);
	`
	_, err = DB.Exec(createUsersTable)
	if err != nil {
		fmt.Println("Error creating users table.")
		panic(err)
	}

	createBookingsTable := `
		CREATE TABLE IF NOT EXISTS bookings (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			room_id TEXT NOT NULL,
			room_price INTEGER NOT NULL,
			total_price INTEGER NOT NULL,
			total_booked_days INTEGER NOT NULL,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			phone TEXT NOT NULL,
			message TEXT NOT NULL,
			start_date DATETIME NOT NULL,
			end_date DATETIME NOT NULL,
			processed BOOLEAN NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
	`
	_, err = DB.Exec(createBookingsTable)
	if err != nil {
		fmt.Println("Error creating bookings table.")
		panic(err)
	}

	createBookingRoomsTable := `
		CREATE TABLE IF NOT EXISTS bookings_rooms (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			room_id TEXT NOT NULL,
			start_date DATETIME NOT NULL,
			end_date DATETIME NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
	`
	_, err = DB.Exec(createBookingRoomsTable)
	if err != nil {
		fmt.Println("Error creating bookings_rooms table.")
		panic(err)
	}

	createContactUsTable := `
		CREATE TABLE IF NOT EXISTS contact_us (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT NOT NULL,
			message TEXT NOT NULL,
			created_at DATETIME NOT NULL,
			updated_at DATETIME NOT NULL
		);
	`
	_, err = DB.Exec(createContactUsTable)
	if err != nil {
		fmt.Println("Error creating contact_us table.")
		panic(err)
	}
}
