package routes

import (
	"dinosaur-cage/controllers"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRoutes initializes and sets up routes for the Gin router.
func SetupRoutes(r *gin.Engine) {
	r.GET("/dinosaurs/:id", controllers.GetDinosaur)
	r.POST("/dinosaurs", controllers.CreateDinosaur)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler)) // Serve Swagger UI
	r.GET("/cages/:cage_id", controllers.GetCageByID)
	r.POST("/cages", controllers.CreateCage)
	r.PUT("/cages/:id", controllers.UpdateCage)
	r.POST("/cages/:cage_id/dinosaurs/:dinosaur_id", controllers.AddDinosaurToCage)
	r.GET("/getdinosaurs/:cage_id", controllers.GetDinosaursInCage)
}
