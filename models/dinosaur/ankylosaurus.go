package models

// AnkylosaurusFactory creates Ankylosaurus dinosaurs.
type AnkylosaurusFactory struct{}

// CreateDinosaur creates a Ankylosaurus dinosaur with the given name.
func (f AnkylosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Ankylosaurus",
		Diet:    Herbivore,
		// Add other specific attributes for Ankylosaurus.
	}
}
