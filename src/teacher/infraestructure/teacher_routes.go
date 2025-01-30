package infrastructure

import (
	"api/src/teacher/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupTeacherRoutes(
	router *gin.Engine,
	saveTeacherController *controllers.CreateTeacherController,
	deleteTeacherController *controllers.DeleteTeacherController,
	viewTeachersController *controllers.ViewTeachersController,
	viewTeacherController *controllers.ViewTeacherController,
) {
	teacherGroup := router.Group("/teachers")
	{
		teacherGroup.POST("", saveTeacherController.Run)
		teacherGroup.GET("", viewTeachersController.Run)
		teacherGroup.GET("/:id", viewTeacherController.Run)
		teacherGroup.DELETE("/:id", deleteTeacherController.Run)
	}
}
