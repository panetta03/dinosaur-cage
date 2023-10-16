package repository

import (
	"dinosaur-cage/database"
	cagemodels "dinosaur-cage/models"
	dinomodels "dinosaur-cage/models/dinosaur"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InsertDinosaur(dinosaur *dinomodels.Dinosaur) error {

	stmt, err := database.GetDB().Prepare("INSERT INTO dinosaurs (name, species, diet, cage_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("Error preparing SQL statement:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(dinosaur.Name, dinosaur.Species, dinosaur.Diet, dinosaur.CageID)
	if err != nil {
		log.Println("Error executing SQL statement:", err)
		return err
	}

	return nil
}

func GetDinosaur(id int) (*dinomodels.Dinosaur, error) {
	row := database.GetDB().QueryRow("SELECT id, name, species, diet, cage_id FROM dinosaurs WHERE id = ?", id)
	dinosaur := dinomodels.Dinosaur{}
	err := row.Scan(&dinosaur.ID, &dinosaur.Name, &dinosaur.Species, &dinosaur.Diet, &dinosaur.CageID)
	if err != nil {
		return nil, err
	}
	return &dinosaur, nil
}

func GetCage(id int) (*cagemodels.Cage, error) {
	row := database.GetDB().QueryRow("SELECT id, name, max_capacity, power_status FROM cages WHERE id = ?", id)
	cage := cagemodels.Cage{}
	err := row.Scan(&cage.ID, &cage.Name, &cage.MaxCapacity, &cage.PowerStatus)
	if err != nil {
		return nil, err
	}
	return &cage, nil
}

func InsertCage(cage *cagemodels.Cage) error {
	stmt, err := database.GetDB().Prepare("INSERT INTO cages (name, power_status, max_capacity, current_capacity) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(cage.Name, cage.PowerStatus, cage.MaxCapacity, cage.CurrentCapacity)
	if err != nil {
		return err
	}

	return nil
}

func UpdateCage(updatedCage cagemodels.Cage) error {
	stmt, err := database.GetDB().Prepare("UPDATE cages SET name=?, max_capacity=?, power_status=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedCage.Name, updatedCage.MaxCapacity, updatedCage.PowerStatus, updatedCage.ID)
	if err != nil {
		return err
	}

	return nil
}

func AddDinosaurToCage(updatedDinosaur dinomodels.Dinosaur) error {
	stmt, err := database.GetDB().Prepare("UPDATE dinosaurs SET cage_id=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedDinosaur.CageID, updatedDinosaur.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetDinosaursInCage(cageID int) []dinomodels.Dinosaur {
	query := "SELECT id, name, species, diet, cage_id FROM dinosaurs WHERE cage_id = ?"
	rows, err := database.GetDB().Query(query, cageID)
	if err != nil {
		return nil
	}
	defer rows.Close()

	dinosaurs := []dinomodels.Dinosaur{}

	for rows.Next() {
		var dinosaur dinomodels.Dinosaur
		if err := rows.Scan(&dinosaur.ID, &dinosaur.Name, &dinosaur.Species, &dinosaur.Diet, &dinosaur.CageID); err != nil {
			return nil
		}
		dinosaurs = append(dinosaurs, dinosaur)
	}

	if err := rows.Err(); err != nil {
		return nil
	}

	if len(dinosaurs) == 0 {
		log.Println("Error executing SQL statement:", err)
		return nil
	}

	return dinosaurs
}
