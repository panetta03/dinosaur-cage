package controllers

import (
	"dinosaur-cage/database"
	"dinosaur-cage/database/repository"
	cagemodels "dinosaur-cage/models"
	validation "dinosaur-cage/utils"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCages returns a list of cages.
// @Summary Get a list of cages
// @Description Retrieve a list of all cages.
// @ID get-cages
// @Produce json
// @Success 200 {array} cagemodels.Cage
// @Router /cages [get]
func GetCages(c *gin.Context) {
	c.JSON(http.StatusOK, repository.Cages)
}

// GetCageByID returns a specific cage by ID.
// @Summary Get a specific cage by ID
// @Description Retrieve a specific cage based on its ID.
// @ID get-cage-by-id
// @Produce json
// @Param id path int true "Cage ID"
// @Success 200 {object} cagemodels.Cage
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Cage not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /cages/{id} [get]
func GetCageByID(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	txn := database.GetDB().Txn(false)
	defer txn.Abort()

	// Retrieve the cage from the go-memdb database
	cage, err := repository.GetCage(txn, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cage"})
		return
	}

	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	c.JSON(http.StatusOK, cage)
}

// CreateCage creates a new cage.
// @Summary Create a new cage
// @Description Create a new cage in the database.
// @ID create-cage
// @Accept json
// @Produce json
// @Param cage body cagemodels.Cage true "Cage object"
// @Success 201 {object} models.Cage
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /cages [post]
func CreateCage(c *gin.Context) {
	var newCage cagemodels.Cage
	if err := c.ShouldBindJSON(&newCage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage data"})
		return
	}

	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Insert the cage into the go-memdb database
	if err := repository.InsertCage(txn, &newCage); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	txn.Commit()
	c.JSON(http.StatusCreated, newCage)
}

// UpdateCage updates an existing cage.
// @Summary Update an existing cage
// @Description Update an existing cage in the database.
// @ID update-cage
// @Accept json
// @Produce json
// @Param id path int true "Cage ID to update"
// @Param cage body models.Cage true "Updated Cage object"
// @Success 200 {object} models.Cage
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Cage Not Found"
// @Router /cages/{id} [put]
func UpdateCage(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	var updatedCage cagemodels.Cage
	if err := c.ShouldBindJSON(&updatedCage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage data"})
		return
	}

	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	cage, err := repository.GetCage(txn, id)
	if cage == nil {
		log.Printf("Error: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	updatedCage.ID = cage.ID // Make sure ID remains the same
	updatedCage = repository.UpdateCage(updatedCage)
	c.JSON(http.StatusOK, updatedCage)
}

// DeleteCage deletes an existing cage by ID.
// @Summary Delete an existing cage
// @Description Delete an existing cage by ID from the database.
// @ID delete-cage
// @Param id path int true "Cage ID to delete"
// @Success 200 {object} string "Cage deleted successfully"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Cage Not Found"
// @Router /cages/{id} [delete]
func DeleteCage(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	cage, err := repository.GetCage(txn, id)
	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	repository.DeleteCage(id)
	c.JSON(http.StatusOK, gin.H{"message": "Cage deleted successfully"})
}

// AddDinosaurToCage adds a dinosaur to a cage.
// @Summary Add a dinosaur to a cage.
// @Description Add a dinosaur to a cage with specific checks.
// @ID add-dinosaur-to-cage
// @Accept json
// @Produce json
// @Param cage_id path int true "Cage ID"
// @Param dinosaur_id path int true "Dinosaur ID"
// @Success 200 {object} cagemodels.Cage "Updated cage"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Not Found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /cages/{cage_id}/dinosaurs/{dinosaur_id} [post]
func AddDinosaurToCage(c *gin.Context) {
	cageID, err := strconv.Atoi(c.Param("cage_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	dinosaurID, err := strconv.Atoi(c.Param("dinosaur_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dinosaur ID"})
		return
	}

	// Log the incoming cage and dinosaur IDs
	log.Printf("Received request to add Dinosaur ID: %d to Cage ID: %d", dinosaurID, cageID)

	// Get the DB
	txn := database.GetDB().Txn(false)
	defer txn.Abort()

	// Get the cage by cage ID
	cage, err := repository.GetCage(txn, cageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cage"})
		return
	}
	if cage == nil {
		log.Printf("Cage with ID %d not found", cageID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	// Get the dinosaur by dinosaur ID
	dinosaur, err := repository.GetDinosaur(txn, dinosaurID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve dinosaur"})
		return
	}
	if dinosaur == nil {
		log.Printf("Dinosaur with ID %d not found", dinosaurID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Dinosaur not found"})
		return
	}

	// Check the rules for adding a dinosaur to the cage
	canAdd, err := validation.CanAddDinosaurToCage(cage, dinosaur)
	if !canAdd {
		log.Printf("Cannot add Dinosaur ID: %d to Cage ID: %d due to rules: %v", dinosaurID, cageID, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Use the custom error message
		return
	}

	// Add the dinosaur to the cage
	repository.AddDinosaurToCage(dinosaur, cage)

	log.Printf("Dinosaur ID: %d successfully added to Cage ID: %d", dinosaurID, cageID)
	c.JSON(http.StatusOK, cage)
}

/*
// ToggleCagePowerStatus toggles the power status of a cage.
// @Summary Toggle the power status of a cage.
// @Description Toggle the power status of a cage with checks.
// @ID toggle-cage-power-status
// @Accept json
// @Produce json
// @Param id path int true "Cage ID"
// @Success 200 {object} cagemodels.Cage "Updated cage"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Cage not found"
// @Failure 500 {object} string "Internal Server Error"
// @Router /cages/{id}/power [put]
func ToggleCagePowerStatus(c *gin.Context) {
	cageID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	// Get the cage by cage ID
	txn := database.GetDB().Txn(true)
	defer txn.Abort()
	cage, err := repository.GetCage(txn, cageID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve cage"})
		return
	}
	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	// Toggle the PowerStatus of the cage
	if cage.PowerStatus == cagemodels.PowerStatusActive {
		cage.PowerStatus = cagemodels.PowerStatusDown
	} else {
		cage.PowerStatus = cagemodels.PowerStatusActive
	}

	c.JSON(http.StatusOK, cage)
}
*/
