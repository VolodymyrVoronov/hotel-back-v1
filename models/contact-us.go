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

func SelectAllContactUs() ([]ContactUs, error) {
	query := `
		SELECT
			id, name, email, message, created_at, updated_at
		FROM
			contact_us
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []ContactUs

	for rows.Next() {
		var message ContactUs

		err := rows.Scan(&message.ID, &message.Name, &message.Email, &message.Message, &message.CreatedAt, &message.UpdatedAt)
		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
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
