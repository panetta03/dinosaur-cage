package models

// SpinosaurusFactory creates Spinosaurus dinosaurs.
type SpinosaurusFactory struct{}

// CreateDinosaur creates a Spinosaurus dinosaur with the given name.
func (f SpinosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Spinosaurus",
		Diet:    Carnivore,
	}
}
