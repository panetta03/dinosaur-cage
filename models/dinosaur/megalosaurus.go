package models

// MegalosaurusFactory creates Megalosaurus dinosaurs.
type MegalosaurusFactory struct{}

// CreateDinosaur creates a Megalosaurus dinosaur with the given name.
func (f MegalosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Megalosaurus",
		Diet:    Carnivore,
	}
}
