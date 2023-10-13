package test

import (
	"dinosaur-cage/database"               // Adjust the import path
	"dinosaur-cage/database/repository"    // Adjust the import path
	models "dinosaur-cage/models/dinosaur" // Adjust the import path
	"testing"
)

func TestInsertDinosaur(t *testing.T) {
	// Set up a test database
	database.InitDB()

	// Start a transaction
	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	dinosaur := &models.Dinosaur{
		ID:      1, // Adjust with a unique ID
		Name:    "Test Dino",
		Species: "Test Species",
		Diet:    models.Herbivore,
	}

	if err := repository.InsertDinosaur(txn, dinosaur); err != nil {
		t.Errorf("Error inserting dinosaur: %v", err)
	}

	// Retrieve the inserted dinosaur to verify
	retrievedDino, err := repository.GetDinosaur(txn, 1) // Adjust with the correct ID
	if err != nil {
		t.Errorf("Error retrieving dinosaur: %v", err)
	}

	if retrievedDino == nil {
		t.Error("Expected dinosaur not found in the database")
	}
}

func TestGetDinosaur(t *testing.T) {
	// Set up a test database
	database.InitDB()

	// Start a transaction
	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Insert a dinosaur for testing
	dinosaur := &models.Dinosaur{
		ID:      1, // Adjust with a unique ID
		Name:    "Test Dino",
		Species: "Test Species",
		Diet:    models.Herbivore,
	}
	if err := repository.InsertDinosaur(txn, dinosaur); err != nil {
		t.Errorf("Error inserting dinosaur: %v", err)
	}

	// Retrieve the inserted dinosaur
	retrievedDino, err := repository.GetDinosaur(txn, 1) // Adjust with the correct ID
	if err != nil {
		t.Errorf("Error retrieving dinosaur: %v", err)
	}

	if retrievedDino == nil {
		t.Error("Expected dinosaur not found in the database")
	}
}
