package repository

import (
	"database/sql"
	repository "dinosaur-cage/database/repository"
	"testing"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

func TestGetDinosaursInCage(t *testing.T) {
	// Open a test database (in-memory SQLite database)
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	// Initialize the database schema (create tables if needed)
	err = InitSchema(db)
	if err != nil {
		t.Fatal(err)
	}

	// Insert test data into the database
	_, err = db.Exec("INSERT INTO dinosaurs (name, species, diet, cage_id) VALUES (?, ?, ?, ?)", "Dino1", "T-Rex", "carnivore", 1)
	if err != nil {
		t.Fatal(err)
	}

	// Call the function to be tested
	dinosaurs := repository.GetDinosaursInCage(1)
	if err != nil {
		t.Fatal(err)
	}

	// Validate the results
	if len(dinosaurs) != 1 {
		t.Fatalf("Expected 1 dinosaur, but got %d", len(dinosaurs))
	}

	// Ensure the properties of the dinosaur match the expected values
	if dinosaurs[0].Name != "Dino1" || dinosaurs[0].Species != "T-Rex" || dinosaurs[0].Diet != "carnivore" || dinosaurs[0].CageID != 1 {
		t.Fatalf("Dinosaur properties do not match expected values")
	}
}

// You can write similar tests for other functions in your repository.

// InitSchema initializes the database schema (e.g., creating tables).
func InitSchema(db *sql.DB) error {
	// Define your database schema initialization here.
	// For example, create tables if they don't exist.
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS dinosaurs (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			species TEXT,
			diet TEXT,
			cage_id INTEGER
		)
	`)
	return err
}
