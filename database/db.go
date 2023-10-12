package database

import (
	"database/sql"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

var db *sql.DB

// InitDB initializes the database connection.
func InitDB() *sql.DB {
	// Set up the PostgreSQL connection string
	connStr := "user=username dbname=jurassicpark sslmode=disable"

	// Open a database connection
	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Verify the database connection
	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}
