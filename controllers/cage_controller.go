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

// Set enums for Cage Request PowerStatus
type PowerStatus string

const (
	PowerStatusActive PowerStatus = "ACTIVE"
	PowerStatusDown   PowerStatus = "DOWN"
)

// CreateCageRequest represents the request to create a new cage.
type CreateCageRequest struct {
	Name        string      `json:"name" binding:"required" example:"Cage A"`
	MaxCapacity int         `json:"max_capacity" binding:"required" example:"100"`
	PowerStatus PowerStatus `json:"power_status" binding:"required" example:"ACTIVE"`
}

// CreateCage creates a new cage.
// @Summary Create a new cage.
// @Description Create a new cage in the database.
// @ID create-cage
// @Accept json
// @Produce json
// @Param cage body CreateCageRequest true "New Cage object"
// @Success 201 {object} models.Cage
// @Failure 400 {object} string "Bad Request"
// @Router /cages [post]
func CreateCage(c *gin.Context) {
	var newCage cagemodels.Cage
	if err := c.ShouldBindJSON(&newCage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage data"})
		return
	}

	//get the DB
	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Insert the cage into the go-memdb database
	if err := repository.InsertCage(txn, &newCage); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	//Commit Changes to DB
	txn.Commit()
	c.JSON(http.StatusCreated, newCage)
}

type UpdateCageRequest struct {
	MaxCapacity *int                    `json:"max_capacity,omitempty"`
	PowerStatus *cagemodels.PowerStatus `json:"power_status,omitempty"`
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

	//provide
	var updateRequest UpdateCageRequest
	if err := c.ShouldBindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid update data"})
		return
	}

	// Get the DB
	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Find the existing cage
	existingCage, err := repository.GetCage(txn, id)
	if existingCage.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	// Update only the specified fields, if provided
	if updateRequest.MaxCapacity != nil {
		existingCage.MaxCapacity = *updateRequest.MaxCapacity
	}
	if updateRequest.PowerStatus != nil {
		existingCage.PowerStatus = *updateRequest.PowerStatus
	}

	// Update the cage in the repository
	repository.UpdateCage(*existingCage) // Pass the existingCage, not a pointer to it

	c.JSON(http.StatusOK, existingCage)
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
