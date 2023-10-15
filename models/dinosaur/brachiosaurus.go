package models

// BrachiosaurusFactory creates Brachiosaurus dinosaurs.
type BrachiosaurusFactory struct{}

// CreateDinosaur creates a Brachiosaurus dinosaur with the given name.
func (f BrachiosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Brachiosaurus",
		Diet:    Herbivore,
		// Add other specific attributes for Brachiosaurus.
	}
}
