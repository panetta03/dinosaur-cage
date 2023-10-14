package database

import (
	dinosaur "dinosaur-cage/models/dinosaur"

	"gorm.io/gorm"
)

// DinosaurRepository represents a repository for interacting with the dinosaur model.
type DinosaurRepository struct {
	db *gorm.DB
}

// NewDinosaurRepository creates a new DinosaurRepository.
func NewDinosaurRepository(db *gorm.DB) *DinosaurRepository {
	return &DinosaurRepository{db: db}
}

// InsertDinosaur inserts a dinosaur into the database.
func (r *DinosaurRepository) InsertDinosaur(dinosaur *dinosaur.Dinosaur) error {
	return r.db.Create(dinosaur).Error
}

// GetDinosaur retrieves a dinosaur from the database by ID.
func (r *DinosaurRepository) GetDinosaur(id uint) (*dinosaur.Dinosaur, error) {
	var dino dinosaur.Dinosaur
	err := r.db.Where("id = ?", id).First(&dino).Error
	if err != nil {
		return nil, err
	}
	return &dino, nil
}
