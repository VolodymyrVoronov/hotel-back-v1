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

func SelectAllSubscriptions() ([]Subscription, error) {
	query := `
		SELECT
			id, email, created_at, updated_at
		FROM
			subscriptions
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var subscriptions []Subscription

	for rows.Next() {
		var subscription Subscription

		err = rows.Scan(&subscription.ID, &subscription.Email, &subscription.CreatedAt, &subscription.UpdatedAt)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	return subscriptions, nil
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
