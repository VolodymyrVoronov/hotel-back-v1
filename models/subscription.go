package models

import (
	"database/sql"
	"hotel-back-v1/db"
	"time"
)

type Subscription struct {
	ID        int64
	Email     string `binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Subscription) InsertSubscription() error {
	query := `
		INSERT INTO subscriptions
			(email, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.Email, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func CheckSubscription(email string) (bool, error) {
	query := `
		SELECT
			email
		FROM
			subscriptions
		WHERE
			email = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email)

	var subscriptionEmail string

	err = row.Scan(&subscriptionEmail)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
