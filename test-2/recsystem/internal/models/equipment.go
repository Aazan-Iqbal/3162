// Filename: internal/models/users.go
package models

import (
	"context"
	"database/sql"
	"fmt"
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

// BASIC CRUD FUNCTIONS ------------------------------------------------------------
// Creating a Read Method for Equipment table
func (m *EquipmentModel) Read() ([]*Equipment, error) {

	statement := `
	            SELECT *
				FROM equipment
				LIMIT 5
	            `

	rows, err := m.DB.Query(statement)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	//defer to close connection before we leave our read method
	defer rows.Close()
	// array to store pointers to the data we get from the form
	equipmentList := []*Equipment{}

	for i := 1; i <= 5; i++ {
		rows.Next()

		equipment := &Equipment{} //variable to hold equipment
		err = rows.Scan(&equipment.equipment_id, &equipment.name, &equipment.image,
			&equipment.equipment_type_id, &equipment.status, &equipment.availability)

		fmt.Println(equipment.name)

		equipmentList = append(equipmentList, equipment)
		if err != nil { // check for errors appending pointers to equipment structs the list
			fmt.Println(err)
			return nil, err
		}
	} // end of for loop

	//check for errors from the query
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return equipmentList, nil
}

// Creating an Insert Method that will add a piece of equipment into the database and return the ID
func (m *EquipmentModel) Insert(name string, image []byte, equipment_type_id int32,
	status bool, availability bool) (int64, error) {

	sql := `
	INSERT INTO equipment (name, image, equipment_type_id, status, availability)
	VALUES($1,$2,$3,$4,$5)
	RETURNING equipment_id				
	`
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	statement, err := m.DB.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(name, image, equipment_type_id, status, availability)
	// err := m.DB.QueryRowContext(ctx, statement, name /*image,*/, equipment_type_id, status, availability).Scan(&id)
	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 1 {
		fmt.Println("Insertion successful")
	} else {
		fmt.Println("Insertion failed")
	}

	return rowsAffected, nil
}

// function to delete a piece of equipment from the database
func (m *EquipmentModel) Delete(equip_id int64) (int64, error) {

	sql := `
	DELETE FROM equipment
	WHERE equipment_id = $1
	`
	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	// defer cancel()
	statement, err := m.DB.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer statement.Close()
	// does the deletion and returns the id of the deleted equipment to be used for
	// confirmation of deletion
	result, err := statement.Exec(equip_id)

	if err != nil {
		return 0, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}

	if rowsAffected == 1 {
		fmt.Println("Insertion successful")
	} else {
		fmt.Println("Insertion failed")
	}

	return rowsAffected, nil
}

func (m *EquipmentModel) Update(equipment *Equipment) (int64, error) {
	var id int64

	//Sql statement that will be ran to insert data, will return the equipment_id which is serial and auto increments
	statement := `
	UPDATE equipment
	SET name =$2, image=$3, equipment_type_id=$4, status=$5, availability=$6
	WHERE equipment_id = $1
	
 `
	//sets the timeout for the DB connection
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	//sets the timeout for the DB connection, passes the statements and associates the arguements with the place holders in the SQL ($1, $2 etc)
	_, err := m.DB.ExecContext(ctx, statement, equipment.equipment_id, equipment.name, equipment.image,
		equipment.equipment_type_id, equipment.status, equipment.availability)

	fmt.Println(equipment.equipment_id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	return id, err

}

// MAIN FUNCTIONALITY
// function to mark a peice of equipment as borrowed in the database
func (m *EquipmentModel) Borrow(equip_id int64) (int64, error) {
	var id int64

	statement := `
	UPDATE equipment
	SET status = false
	WHERE equipment_id = ($1)	
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, statement, equip_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// function to mark a peice of equipment as available for borrowing in the database
func (m *EquipmentModel) Return(equip_id int64) (int64, error) {
	var id int64

	statement := `
	UPDATE equipment
	SET status = true
	WHERE equipment_id = ($1);		
	`
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, statement, equip_id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
