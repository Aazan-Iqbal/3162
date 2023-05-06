// Welcome to my page this is my main.go
package main

import (
	"context"
	"crypto/tls"
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Aazan-Iqbal/3161/test-2/recsystem/internal/models"
	"github.com/alexedwards/scs/v2"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// create a new type
type application struct {
	// basic variables for site management
	errorLog        *log.Logger
	infoLog         *log.Logger
	sessionsManager *scs.SessionManager

	// variables for functionality
	users     models.UserModel
	equipment models.EquipmentModel
}

func main() {
	// Create a flag for specifiying the port number
	// when starting the server
	addr := flag.String("port", ":4000", "HTTP network address")
	dsn := flag.String("dsn", os.Getenv("RECSYSTEM"), "PostgreSQL DSN")
	flag.Parse()

	// Create an instance of the connection pool
	db, err := openDB(*dsn)
	if err != nil {
		log.Println(err)
		return
	}

	//create instances of errorLog & infoLog
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// setup a new session manager
	sessionManager := scs.New()
	sessionManager.Lifetime = 1 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.Secure = true                   //false if the cookies aren't secure
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode //Same site

	// create an instance of the application type
	app := &application{
		errorLog:        errorLog,
		infoLog:         infoLog,
		sessionsManager: sessionManager,
		users:           models.UserModel{DB: db},
		equipment:       models.EquipmentModel{DB: db},
	}

	defer db.Close()
	// acquired a database connection pool
	infoLog.Println("database connection pool established")
	// create and start a custom web server
	infoLog.Printf("starting server on %s", *addr)

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256}, // added with security
	}

	srv := &http.Server{
		Addr:         *addr,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig:    tlsConfig,
	}

	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
	log.Fatal(err)
}

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
