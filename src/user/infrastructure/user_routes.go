package infrastructure

import (
	"api/src/user/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(
	router *gin.Engine,
	authMiddleware gin.HandlerFunc,
	registerController *controllers.RegisterController,
	loginController *controllers.LoginController,
	getMeController *controllers.GetMeController,
	viewUsersController *controllers.ViewUsersController,
) {
	authGroup := router.Group("/auth")
	{
		authGroup.POST("/register", registerController.Run)
		authGroup.POST("/login", loginController.Run)
		authGroup.GET("/users", viewUsersController.Run)              // GET: listar usuarios registrados
		authGroup.GET("/me", authMiddleware, getMeController.Run)    // GET: mi perfil (requiere token)
	}
}
