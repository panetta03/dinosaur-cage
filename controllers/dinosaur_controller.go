package controllers

import (
	"dinosaur-cage/database"
	"dinosaur-cage/database/repository"
	models "dinosaur-cage/models/dinosaur"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create a dinosaur
// @Description Create a new dinosaur
// @ID create-dinosaur
// @Accept json
// @Produce json
// @Param dinosaur body models.Dinosaur true "Dinosaur object"
// @Success 201 {object} models.Dinosaur
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dinosaurs [post]

var lastDinosaurID int // Define a package-level variable

func CreateDinosaur(c *gin.Context) {
	var d models.Dinosaur
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Increment the lastDinosaurID
	lastDinosaurID++
	d.ID = lastDinosaurID

	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Insert the dinosaur into the go-memdb database
	if err := repository.InsertDinosaur(txn, &d); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dinosaur"})
		return
	}

	txn.Commit()
	c.JSON(http.StatusCreated, d)
}

// @Summary Get a dinosaur by ID
// @Description Get a dinosaur's details by its ID
// @ID get-dinosaur
// @Produce json
// @Param id path int true "Dinosaur ID"
// @Success 200 {object} models.Dinosaur
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Failure 404 {object} string "Not Found"
// @Router /dinosaurs/{id} [get]
func GetDinosaur(c *gin.Context) {
	id := c.Param("id")
	dinosaurID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dinosaur ID"})
		return
	}

	txn := database.GetDB().Txn(false)
	defer txn.Abort()

	// Retrieve the dinosaur from the go-memdb database
	dinosaur, err := repository.GetDinosaur(txn, dinosaurID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve dinosaur"})
		return
	}

	if dinosaur == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Dinosaur not found"})
		return
	}

	c.JSON(http.StatusOK, dinosaur)
}
