package models

import (
	"errors"
	"hotel-back-v1/db"
	"hotel-back-v1/utils"
	"time"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
	Role     string `binding:"required"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u *User) InsertUser() error {
	query := `
		INSERT INTO users
			(email, password, created_at, updated_at)
		VALUES
			(?, ?, ?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, time.Now(), time.Now())
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId

	return nil
}

func (u *User) ValidateCredentials() error {
	query := `
		SELECT 
			id, password
		FROM 
			users
		WHERE 
			email = ?
	`

	var retrievedPassword string

	row := db.DB.QueryRow(query, u.Email)
	if err := row.Err(); err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}
