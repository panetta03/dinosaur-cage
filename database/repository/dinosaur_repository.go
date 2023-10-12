package repository

import (
	"database/sql"
	models "dinosaur-cage/models/dinosaur"
)

// DinosaurRepository is responsible for database operations related to dinosaurs.
type DinosaurRepository struct {
	db *sql.DB
}

// NewDinosaurRepository creates a new instance of DinosaurRepository.
func NewDinosaurRepository(db *sql.DB) *DinosaurRepository {
	return &DinosaurRepository{db}
}

// InsertDinosaur inserts a dinosaur into the database.
func (r *DinosaurRepository) InsertDinosaur(dinosaur models.Dinosaur) error {

	// Define the SQL query for inserting a dinosaur
	insertDinosaurQuery := `
        INSERT INTO dinosaurs (name, species, diet, cage_id)
        VALUES ($1, $2, $3, $4)`

	// Execute the SQL query
	_, err := r.db.Exec(
		insertDinosaurQuery,
		dinosaur.Name,
		dinosaur.Species,
		string(dinosaur.Diet),
		dinosaur.CageID,
	)

	return err
}
