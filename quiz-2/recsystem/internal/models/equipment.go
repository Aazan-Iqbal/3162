// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"time"
)

// Let's model the users table
type Equipment struct {
	equipment_id      int64
	name              string
	image             []byte
	equipment_type_id int32
	status            bool
	availability      bool
}

// Setup dependency injection
type EquipmentModel struct {
	DB *sql.DB
}

// Write SQL code to access the database
// Creating a Get Method for Equipment table
func (m *EquipmentModel) Get() (*Equipment, error) {
	var q Equipment

	statement := `
	            SELECT equipment_id, body
				FROM equipment
				LIMIT 20
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement).Scan(&q.equipment_id, &q.name, &q.image,
		&q.equipment_type_id, &q.status, &q.availability)
	if err != nil {
		return nil, err
	}
	return &q, err
}

// Creating an Insert Method that will post users entered into the database
func (m *EquipmentModel) Insert(body string) (int64, error) {
	var id int64

	statement := `
	            INSERT INTO equipments(body)
				VALUES($1)
				RETURNING equipment_id				
	             `
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	err := m.DB.QueryRowContext(ctx, statement, body).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
