// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type User struct {
	user_id      int64
	email        string
	first_name   string
	last_name    string
	dob          string
	address      string
	phone_number string
	roles_id     int32
	password     string //temporarily string will turn into hash later
	CreatedAt    string
}

// Setup dependency injection
type UserModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// TODO
// Creating a Get Method for Users table
func (m *UserModel) Get() (*User, error) {
	var q User

	statement := `
	            SELECT user_id
				FROM users
				ORDER BY RANDOM()
				LIMIT 10
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.user_id, &q.email, &q.first_name,
		&q.last_name, &q.dob, &q.address, &q.phone_number, &q.roles_id, &q.password, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *UserModel) Insert(body string) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO users(body)
				VALUES($1)
				RETURNING user_id				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
