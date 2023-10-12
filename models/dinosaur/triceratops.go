package models

// TriceratopsFactory creates Triceratops dinosaurs.
type TriceratopsFactory struct{}

// CreateDinosaur creates a Triceratops dinosaur with the given name.
func (f TriceratopsFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Triceratops",
		Diet:    Herbivore,
		// Add other specific attributes for Triceratops.
	}
}
