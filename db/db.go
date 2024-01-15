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

func createTables() {}
