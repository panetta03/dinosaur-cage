package utils

import (
	cagemodels "dinosaur-cage/models"
	dinomodels "dinosaur-cage/models/dinosaur"
	"errors"
)

// canAddDinosaurToCage checks if the specified dinosaur can be added to the cage based on the rules.
func CanAddDinosaurToCage(cage *cagemodels.Cage, dinosaur *dinomodels.Dinosaur) (bool, error) {
	// 1. Carnivores can only be in a cage with other dinosaurs of the same species.
	if dinosaur.Diet == "carnivore" {
		if cage.CurrentDinosaurs != nil {
			for _, dino := range *cage.CurrentDinosaurs {
				if dino.Species != dinosaur.Species {
					return false, errors.New("Carnivores can only be in a cage with dinosaurs of the same species")
				}
			}
		}
	}

	// 2. Herbivores cannot be in the same cage as carnivores.
	if dinosaur.Diet == "herbivore" {
		if cage.CurrentDinosaurs != nil {
			for _, dino := range *cage.CurrentDinosaurs {
				if dino.Diet == "carnivore" {
					return false, errors.New("Herbivores cannot be in the same cage as carnivores")
				}
			}
		}
	}

	// 3. Cages have a maximum capacity for how many dinosaurs they can hold.
	if cage.CurrentDinosaurs != nil && len(*cage.CurrentDinosaurs) >= cage.MaxCapacity {
		return false, errors.New("Cage is at maximum capacity for dinosaurs")
	}

	// 4. Dinosaurs cannot be moved into a cage that is powered down.
	if cage.PowerStatus == cagemodels.PowerStatusDown {
		return false, errors.New("Dinosaurs cannot be moved into a powered down cage")
	}

	// All checks passed
	return true, nil
}
