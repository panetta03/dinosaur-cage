package main

import (
	"dinosaur-cage/database"
	_ "dinosaur-cage/docs"
	"dinosaur-cage/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the go-memdb database
	database.InitDB()

	// Initialize the Gin router
	r := gin.Default()
	routes.SetupRoutes(r)

	// Run the server
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
