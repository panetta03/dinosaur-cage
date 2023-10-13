package repository

import (
	models "dinosaur-cage/models/dinosaur"

	"github.com/hashicorp/go-memdb"
)

func InsertDinosaur(txn *memdb.Txn, dinosaur *models.Dinosaur) error {
	return txn.Insert("dinosaurs", dinosaur)
}

func GetDinosaur(txn *memdb.Txn, id int) (*models.Dinosaur, error) {
	raw, err := txn.First("dinosaurs", "id", id)
	if err != nil {
		return nil, err
	}

	if raw == nil {
		return nil, nil
	}

	dinosaur := raw.(*models.Dinosaur)
	return dinosaur, nil
}
