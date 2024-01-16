package models

import (
	"hotel-back-v1/db"
	"time"
)

type Role struct {
	ID        int64
	RoleType  string `binding:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (r *Role) InsertRole() error {
	query := `
		INSERT INTO roles
			(role_type, created_at, updated_at)
		VALUES
			(?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.RoleType, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func SelectAllRoles() ([]Role, error) {
	query := `
		SELECT 
			id, role_type, created_at, updated_at
		FROM 
			roles
	`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []Role

	for rows.Next() {
		var role Role
		err = rows.Scan(&role.ID, &role.RoleType, &role.CreatedAt, &role.UpdatedAt)
		if err != nil {
			return nil, err
		}

		roles = append(roles, role)
	}

	return roles, nil
}
