package models

// StegosaurusFactory creates Stegosaurus dinosaurs.
type StegosaurusFactory struct{}

// CreateDinosaur creates a Stegosaurus dinosaur with the given name.
func (f StegosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Stegosaurus",
		Diet:    Herbivore,
		// Add other specific attributes for Stegosaurus.
	}
}
