package controllers

/*
import (
	"net/http"
	"strconv"

	"dinosaur-cage/database"
	"dinosaur-cage/database/repository"
	"dinosaur-cage/models"

	"github.com/gin-gonic/gin"
)

// GetCages returns a list of all cages.
func GetCages(c *gin.Context) {
	c.JSON(http.StatusOK, repository.Cages)
}

// GetCageByID returns a specific cage by ID.
func GetCageByID(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	cage := database.FindCageByID(id)
	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	c.JSON(http.StatusOK, cage)
}

// CreateCage creates a new cage.
func CreateCage(c *gin.Context) {
	var newCage models.Cage
	if err := c.ShouldBindJSON(&newCage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage data"})
		return
	}

	createdCage := repository.CreateCage(newCage)
	c.JSON(http.StatusCreated, createdCage)
}

// UpdateCage updates an existing cage.
func UpdateCage(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	var updatedCage models.Cage
	if err := c.ShouldBindJSON(&updatedCage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage data"})
		return
	}

	cage := database.FindCageByID(id)
	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	updatedCage.ID = cage.ID // Make sure ID remains the same
	updatedCage = database.UpdateCage(updatedCage)
	c.JSON(http.StatusOK, updatedCage)
}

// DeleteCage deletes a cage by ID.
func DeleteCage(c *gin.Context) {
	cageID := c.Param("id")
	id, err := strconv.Atoi(cageID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid cage ID"})
		return
	}

	cage := database.FindCageByID(id)
	if cage == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cage not found"})
		return
	}

	database.DeleteCage(id)
	c.JSON(http.StatusOK, gin.H{"message": "Cage deleted successfully"})
}
*/
