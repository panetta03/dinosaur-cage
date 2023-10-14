package repository

import (
	//	cagemodels "dinosaur-cage/models"
	dinomodels "dinosaur-cage/models/dinosaur"
	//cagemodels "dinosaur-cage/models"
	"github.com/hashicorp/go-memdb"
)

type DinosaurRepository struct {
	Dinosaurs []dinomodels.Dinosaur
	//Cages             []cagemodels.Cage
	//CageIDCounter     uint
	DinosaurIDCounter int
}

func (r *DinosaurRepository) CreateDinosaur(dinosaur dinomodels.Dinosaur, txn *memdb.Txn) (dinomodels.Dinosaur, error) {
	r.DinosaurIDCounter++             // Auto-increment the Dinosaur ID
	dinosaur.ID = r.DinosaurIDCounter // Assign the new ID to the dinosaur

	// Insert the new dinosaur into the database using a transaction
	if err := InsertDinosaur(txn, &dinosaur); err != nil {
		return dinomodels.Dinosaur{}, err
	}

	// Update the in-memory slice (if necessary)
	r.Dinosaurs = append(r.Dinosaurs, dinosaur)

	return dinosaur, nil
}

// ...

// InsertDinosaur is similar to InsertCage
func InsertDinosaur(txn *memdb.Txn, dinosaur *dinomodels.Dinosaur) error {
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

/*
func (r *DinosaurRepository) FindCageByID(id int) *cagemodels.Cage {
	for _, cage := range r.Cages {
		if cage.ID == uint(id) {
			return &cage
		}
	}
	return nil
}

func CreateCage(cage cagemodels.Cage, txn *memdb.Txn) (cagemodels.Cage, error) {
    r.CageIDCounter++
    cage.ID = r.CageIDCounter

    // Insert the new cage into the database using a transaction
    if err := InsertCage(txn, &cage); err != nil {
        return cagemodels.Cage{}, err
    }

    // Update the in-memory slice (if necessary)
    r.Cages = append(r.Cages, cage)

    return cage, nil
}

func (r *DinosaurRepository) UpdateCage(updatedCage cagemodels.Cage) cagemodels.Cage {
	for i, cage := range r.Cages {
		if cage.ID == updatedCage.ID {
			r.Cages[i] = updatedCage
			return updatedCage
		}
	}
	return cagemodels.Cage{}
}

func (r *DinosaurRepository) DeleteCage(id int) {
	for i, cage := range r.Cages {
		if cage.ID == uint(id) {
			r.Cages = append(r.Cages[:i], r.Cages[i+1:]...)
			return
		}
	}
}
*/
