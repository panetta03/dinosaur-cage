package repository

import (
	cagemodels "dinosaur-cage/models"
	dinomodels "dinosaur-cage/models/dinosaur"

	"github.com/hashicorp/go-memdb"
)

var (
	Dinosaurs         []dinomodels.Dinosaur
	Cages             []cagemodels.Cage
	CageIDCounter     int
	DinosaurIDCounter int
)

func InsertDinosaur(txn *memdb.Txn, dinosaur *dinomodels.Dinosaur) error {
	// Retrieve the highest ID from the 'dinosaurs' table
	highestID := 0
	raw, err := txn.Last("dinosaurs", "id")
	if err == nil && raw != nil {
		highestID = raw.(*dinomodels.Dinosaur).ID
	}

	// Increment the ID
	dinosaur.ID = highestID + 1

	if err := txn.Insert("dinosaurs", dinosaur); err != nil {
		return err
	}

	return nil
}

func GetDinosaur(txn *memdb.Txn, id int) (*dinomodels.Dinosaur, error) {
	raw, err := txn.First("dinosaurs", "id", id)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	dinosaur := raw.(*dinomodels.Dinosaur)
	return dinosaur, nil
}

func GetCage(txn *memdb.Txn, id int) (*cagemodels.Cage, error) {
	raw, err := txn.First("cages", "id", id)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	cage := raw.(*cagemodels.Cage)
	return cage, nil
}

func InsertCage(txn *memdb.Txn, cage *cagemodels.Cage) error {
	// Retrieve the highest ID from the 'cages' table
	highestID := 0
	raw, err := txn.Last("cages", "id")
	if err == nil && raw != nil {
		highestID = raw.(*cagemodels.Cage).ID
	}

	// Increment the ID
	cage.ID = highestID + 1

	if err := txn.Insert("cages", cage); err != nil {
		return err
	}

	return nil
}
func UpdateCage(updatedCage cagemodels.Cage) cagemodels.Cage {
	for i, cage := range Cages {
		if cage.ID == updatedCage.ID {
			// Create a new instance with the updated values
			updated := cagemodels.Cage{
				ID:          cage.ID,
				Name:        updatedCage.Name,
				MaxCapacity: updatedCage.MaxCapacity,
				PowerStatus: updatedCage.PowerStatus,
				// Include other fields here
			}

			// Update the cage in the slice
			Cages[i] = updated

			// Return the updated instance
			return updated
		}
	}
	return cagemodels.Cage{}
}

func AddDinosaurToCage(dinosaur *dinomodels.Dinosaur, cage *cagemodels.Cage) {

	// Check for nil or create a new CurrentDinosaurs slice
	if cage.CurrentDinosaurs == nil {
		cage.CurrentDinosaurs = &[]dinomodels.Dinosaur{*dinosaur}
	} else {

		// Append the dinosaur to the CurrentDinosaurs of the cage
		*cage.CurrentDinosaurs = append(*cage.CurrentDinosaurs, *dinosaur)
	}

}
