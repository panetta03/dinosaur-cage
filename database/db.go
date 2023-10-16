package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var db *sql.DB

// InitDB initializes the SQLite database.
func InitDB() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Connect to the SQLite database file
	// If the file doesn't exist, it will be created.
	db, err := sql.Open("sqlite3", wd+"/database/dino.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	// Log that the database connection is successfully established.
	log.Println("Database connection established")

	// Create the "dinosaurs" and "cages" tables if they don't exist.
	createTable := `

        CREATE TABLE IF NOT EXISTS dinosaurs (
            id INTEGER PRIMARY KEY,
            name TEXT,
            species TEXT,
            diet TEXT,
            cage_id INTEGER,
            FOREIGN KEY (cage_id) REFERENCES cages(id)
        );
        
        CREATE TABLE IF NOT EXISTS cages (
            id INTEGER PRIMARY KEY,
            name TEXT,
            power_status TEXT,
            max_capacity INTEGER,
            current_capacity INTEGER
        );
    `

	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Log that the tables are created successfully.
	log.Println("Database tables created")
}

// GetDB returns the SQLite database instance.
func GetDB() *sql.DB {

	//Get working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Connect to the SQLite database file
	// If the file doesn't exist, it will be created.
	db, err := sql.Open("sqlite3", wd+"/database/dino.db")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	return db
}
