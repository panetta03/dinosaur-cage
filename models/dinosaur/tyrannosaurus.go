package models

// TyrannosaurusFactory creates Tyrannosaurus dinosaurs.
type TyrannosaurusFactory struct{}

// CreateDinosaur creates a Tyrannosaurus dinosaur with the given name.
func (f TyrannosaurusFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Tyrannosaurus",
		Diet:    Carnivore,
	}
}
