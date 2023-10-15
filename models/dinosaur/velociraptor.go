package models

// VelociraptorFactory creates Velociraptor dinosaurs.
type VelociraptorFactory struct{}

// CreateDinosaur creates a Velociraptor dinosaur with the given name.
func (f VelociraptorFactory) CreateDinosaur(name string) Dinosaur {
	return Dinosaur{
		Name:    name,
		Species: "Velociraptor",
		Diet:    Carnivore,
	}
}
