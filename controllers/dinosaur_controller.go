package controllers

import (
	"dinosaur-cage/database"
	"dinosaur-cage/database/repository"
	models "dinosaur-cage/models/dinosaur"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateDinosaurRequest represents the request to create a new dinosaur.
type CreateDinosaurRequest struct {
	Name    string `json:"name" binding:"required" example:"Dino"`
	Species string `json:"species" binding:"required" example:"Tyrannosaurus"`
}

// CreateDinosaur creates a new dinosaur.
// @Summary Create a dinosaur
// @Description Create a new dinosaur
// @ID create-dinosaur
// @Accept json
// @Produce json
// @Param dinosaur body CreateDinosaurRequest true "Dinosaur object"
// @Success 201 {object} models.Dinosaur
// @Failure 400 {object} string "Bad Request"
// @Failure 500 {object} string "Internal Server Error"
// @Router /dinosaurs [post]
func CreateDinosaur(c *gin.Context) {
	var d models.Dinosaur
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	txn := database.GetDB().Txn(true)
	defer txn.Abort()

	// Determine the factory based on the dinosaur's species
	var factory models.DinosaurFactory
	switch d.Species {
	case "Ankylosaurus":
		factory = &models.AnkylosaurusFactory{}
	case "Brachiosaurus":
		factory = &models.BrachiosaurusFactory{}
	case "Megalosaurus":
		factory = &models.MegalosaurusFactory{}
	case "Spinosaurus":
		factory = &models.SpinosaurusFactory{}
	case "Stegosaurus":
		factory = &models.StegosaurusFactory{}
	case "Triceratops":
		factory = &models.TriceratopsFactory{}
	case "Tyrannosaurus":
		factory = &models.TyrannosaurusFactory{}
	case "Velociraptor":
		factory = &models.VelociraptorFactory{}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dinosaur species"})
		return
	}

	// Create the dinosaur with the appropriate factory
	newDinosaur := factory.CreateDinosaur(d.Name)

	// Insert the dinosaur into the go-memdb database
	if err := repository.InsertDinosaur(txn, &newDinosaur); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dinosaur"})
		return
	}

	txn.Commit()
	c.JSON(http.StatusCreated, newDinosaur)
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
