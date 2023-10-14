package controllers

import (
	"dinosaur-cage/database"
	"dinosaur-cage/database/repository"
	cagemodels "dinosaur-cage/models"
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
