package models

// Dinosaur represents a dinosaur in the Jurassic Park.
type Dinosaur struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Species string `json:"species"`
	Diet    Diet   `json:"diet"`
	CageID  int    `json:"cage_id"`
	// Add other fields specific to dinosaurs if needed.
}

// Diet is an enumeration of diet types.
type Diet string

const (
	Carnivore Diet = "carnivore"
	Herbivore Diet = "herbivore"
)
