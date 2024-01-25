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

type UserLogin struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type DeleteUser struct {
	ID int64
}

func SelectAllUsers() ([]User, error) {
	query := `
		SELECT
			id, email, role_type, created_at, updated_at
		FROM
			users
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.ID, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (u *User) InsertUser() error {
	query := `
		INSERT INTO users
			(email, password, role_type, created_at, updated_at)
		VALUES
			(?, ?, ?, ?, ?)
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

	result, err := stmt.Exec(u.Email, hashedPassword, u.Role, time.Now(), time.Now())
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

func (u *DeleteUser) DeleteUserFromDB() error {
	query := `
		DELETE FROM
			users
		WHERE
			id = ?
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.ID)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserLogin) ValidateCredentials() error {
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
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("invalid credentials")
	}

	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func (u *User) CheckRole(role string) bool {
	return u.Role == role
}

func GetUserRole(email string) (string, error) {
	query := `SELECT role_type FROM users WHERE email = ?`

	row := db.DB.QueryRow(query, email)

	var role string

	err := row.Scan(&role)

	if err != nil {
		return "", err
	}

	return role, nil
}
