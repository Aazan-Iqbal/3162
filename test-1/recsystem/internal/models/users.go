// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// variables for possible errors
var (
	ErrNoRecord           = errors.New("no matching record found")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateEmail     = errors.New("duplicate email")
)

// Let's model the users table
type User struct {
	User_id      int64
	Email        string
	First_name   string
	Last_name    string
	Dob          string
	Address      string
	Phone_number string
	Roles_id     int32
	Password     string //temporarily string will turn into hash later
	CreatedAt    string
}

// Setup dependency injection
type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
	log.Println(password)
	var id int

	var hashedPassword []byte
	//check if there is a row in the table for the email provided
	query := `
			SELECT user_id, password_hash
			FROM users
			WHERE email = $1
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, query, email).Scan(&id, &hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Println("user does not exist/wrong email")
			return 0, ErrInvalidCredentials

		} else {
			return 0, err
		}
	} //handling error
	//the user does exist
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			log.Println("Incorrect Password")
			return 0, ErrInvalidCredentials
		} else {
			return 0, err
		}
	}
	log.Println("password is correct.")
	//password is correct
	return id, nil
}

// Write SQL code to access the database
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
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.User_id, &q.Email, &q.First_name,
		&q.Last_name, &q.Dob, &q.Address, &q.Phone_number, &q.Roles_id, &q.Password, &q.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *UserModel) Insert(user User) error {
	//let's Hash the password
	log.Println(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}
	query := `
			INSERT INTO users(email, first_name, last_name, dob, address,
				phone_number, roles_id, password_hash)
			VALUES($1, $2, $3, $4, $5, $6, $7, $8)
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	_, err = m.DB.ExecContext(ctx, query, user.Email, user.First_name,
		user.Last_name, user.Dob, user.Address, user.Phone_number, user.Roles_id,
		hashedPassword)
	if err != nil {
		switch {
		case err.Error() == `pgx: duplicate key value violates unique constraint "users_email_key"`:
			return ErrDuplicateEmail
		default:
			return err
		}
	}
	return nil
}
