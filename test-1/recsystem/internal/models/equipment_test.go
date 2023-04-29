package models

// import (
// 	"context"
// 	"database/sql"
// 	"flag"
// 	"log"
// 	"os"
// 	"testing"
// 	"time"

// 	_ "github.com/jackc/pgx/v5/stdlib"
// )

// // Function to open the database connection and setup the database connection pool
// func openDB() (*sql.DB, error) {
// 	db, err := sql.Open("pgx", *flag.String("dsn", os.Getenv("RECSYSTEM_DB_DSN"), "PostgreSQL DSN"))
// 	if err != nil {
// 		return nil, err
// 	}
// 	// use a context to check if the DB is reachable
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //always to this
// 	defer cancel()                                                          // then this to clean up
// 	// let's ping the DB
// 	err = db.PingContext(ctx)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }

// //----------------------------------------------------------------------------------------
// // TESTING RETRIEVING EQUIPMENT LISTINGS FROM THE DB
// // func TestGet(t *testing.T) {

// // 	db, err := openDB()

// // 	if err != nil {
// // 		log.Println(err)
// // 		return
// // 	}
// // 	defer db.Close()

// // 	e := EquipmentModel{
// // 		DB: db,
// // 	}

// // 	Equipment_Struct, err := e.Get()
// // 	if err != nil {
// // 		t.Errorf("unexpected error calling Insert function: %s", err)
// // 	}

// // 	log.Print("Equipmwent Successdul Retrieval")
// // 	log.Printf("Name: %v\nAvailability: %v\nSatus: %v\nEquip_type_id: %v\nimage: %v",
// // 		Equipment_Struct.name, Equipment_Struct.availability, Equipment_Struct.status,
// // 		Equipment_Struct.equipment_type_id, Equipment_Struct.image)

// // }

// // ---------------------------------------------------------------------------------------------
// // TESTING DELETING EQUIPMET FROM THE DATABASE
// // func TestDelete(t *testing.T) {

// // 	db, err := openDB()

// // 	if err != nil {
// // 		log.Println(err)
// // 		return
// // 	}
// // 	defer db.Close()

// // 	e := EquipmentModel{
// // 		DB: db,
// // 	}

// // 	equip_id := 7
// // 	RowsAffected, err := e.Delete(int64(equip_id))
// // 	if RowsAffected != 1 || err != nil {
// // 		t.Errorf("unexpected error calling Insert function: %s", err)
// // 		log.Fatal()
// // 	}

// // 	log.Print("Equipmwent Successdul Retrieval")
// // 	log.Print(RowsAffected)

// // }

// //----------------------------------------------------------------------------------------------
// // TESTING INSERITNG EQUIPMENT IN THE DB
// // func TestInsert(t *testing.T) {

// // 	db, err := openDB()

// // 	if err != nil {
// // 		log.Println(err)
// // 		return
// // 	}
// // 	defer db.Close()

// // 	e := EquipmentModel{
// // 		DB: db,
// // 	}

// // 	img := []byte("8557-0473587653830")
// // 	RowsAffected, err := e.Insert("Basketball", img, 1, true, true)
// // 	if RowsAffected != 1 || err != nil {
// // 		t.Errorf("unexpected error calling Insert function: %s", err)
// // 		log.Fatal(err)
// // 	}
// // 	t.Log("Equipmwent Successdul insertion")

// // }

// // ------------------------------------------------------------------------------------------
// // TESTING FUNCTION FOR MARKING EQUIPMENT AS BORROWED
// func TestBorrow(t *testing.T) {

// 	db, err := openDB()

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer db.Close()

// 	e := EquipmentModel{
// 		DB: db,
// 	}

// 	id, err := e.Borrow(6)
// 	if err != nil {
// 		t.Errorf("unexpected error calling Insert function: %s", err)
// 		log.Fatal(err)
// 	}
// 	t.Log("Equipmwent Marked as Borrowed")
// 	log.Println(id)

// }

// // ------------------------------------------------------------------------------------------
// // TESTING FUNCTION FOR MARKING EQUIPMENT AS BORROWED
// func TestReturn(t *testing.T) {

// 	db, err := openDB()

// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	defer db.Close()

// 	e := EquipmentModel{
// 		DB: db,
// 	}

// 	id, err := e.Borrow(9)
// 	if err != nil {
// 		t.Errorf("unexpected error calling Insert function: %s", err)
// 		log.Fatal(err)
// 	}
// 	t.Log("Equipmwent Marked as Returned")
// 	log.Println(id)

// }
