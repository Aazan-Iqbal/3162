package models

import (
	"context"
	"database/sql"
	"flag"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// Function to open the database connection and setup the database connection pool
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	// use a context to check if the DB is reachable
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second) //always to this
	defer cancel()                                                          // then this to clean up
	// let's ping the DB
	err = db.PingContext(ctx)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// TESTING RETRIEVING EQUIPMENT LISTINGS FROM THE DB
func TestGet(t *testing.T) {

	dsn := flag.String("dsn", os.Getenv("RECSYSTEM_DB_DSN"), "PostgreSQL DSN")
	db, err := openDB(*dsn)

	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	e := EquipmentModel{
		DB: db,
	}

	Equipment_Struct, err := e.Get()
	if Equipment_Struct != nil || err != nil {
		t.Errorf("unexpected error calling Insert function: %s", err)
		log.Fatal(err)
	}

	t.Log("Equipmwent Successdul insertions")
	

}

// TESTING INSERITNG EQUIPMENT IN THE DB
func TestInsert(t *testing.T) {

	dsn := flag.String("dsn", os.Getenv("RECSYSTEM_DB_DSN"), "PostgreSQL DSN")
	db, err := openDB(*dsn)

	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	e := EquipmentModel{
		DB: db,
	}

	img := []byte("8557-0473587653830")
	RowsAffected, err := e.Insert("Basketball", img, 1, true, true)
	if RowsAffected != 1 || err != nil {
		t.Errorf("unexpected error calling Insert function: %s", err)
		log.Fatal(err)
	}
	t.Log("Equipmwent Successdul insertion")

}
