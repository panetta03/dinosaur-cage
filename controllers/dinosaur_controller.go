package controllers

import (
	"dinosaur-cage/database/repository"
	models "dinosaur-cage/models/dinosaur"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Set up console and file log output
	logFile, err := os.OpenFile(wd+"/utils/dinosaur-cage.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	defer logFile.Close()
	log.SetOutput(io.MultiWriter(os.Stdout, logFile))

	var d CreateDinosaurRequest
	if err := c.ShouldBindJSON(&d); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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
	// Insert the dinosaur into the SQLite database
	log.Println("Inserting dinosaur:", newDinosaur.Name, newDinosaur.Species)
	if err := repository.InsertDinosaur(&newDinosaur); err != nil {
		log.Println("Failed to insert dinosaur:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create dinosaur"})
		return
	}

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

	// Retrieve the dinosaur from the SQLite database
	dinosaur, err := repository.GetDinosaur(dinosaurID)
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
