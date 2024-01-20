package models

import (
	"hotel-back-v1/db"
	"time"
)

type ContactUs struct {
	ID        int64
	Name      string `binding:"required"`
	Email     string `binding:"required"`
	Message   string `binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *ContactUs) InsertContactUs() error {
	query := `
		INSERT INTO contact_us
			(name, email, message, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(c.Name, c.Email, c.Message, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}
