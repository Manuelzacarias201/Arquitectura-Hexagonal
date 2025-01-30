package infrastructure

import (
	"api/src/alumn/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupAlumnRoutes(
	router *gin.Engine,
	addAlumnController *controllers.AddAlumnController,
	deleteAlumnController *controllers.RemoveAlumnController,
	viewAlumnsController *controllers.ViewAllAlumnsController,
	viewAlumnController *controllers.ViewOneAlumnController,
) {
	alumnGroup := router.Group("/alumns")
	{
		alumnGroup.POST("", addAlumnController.Run)
		alumnGroup.GET("", viewAlumnsController.Run)
		alumnGroup.GET("/:id", viewAlumnController.Run)
		alumnGroup.DELETE("/:id", deleteAlumnController.Run)
	}
}
