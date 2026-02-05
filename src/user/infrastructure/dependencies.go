package infrastructure

import (
	"api/src/core"
	"api/src/user/application"
	"api/src/user/infrastructure/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func InitUsers(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE USUARIOS")

	// Instanciar repositorios compartidos
	bcryptRepo := core.NewBcryptRepository()
	jwtRepo, err := core.NewJWTRepository()
	if err != nil {
		log.Fatalf("Error al inicializar JWT: %v", err)
	}
	authMiddleware := core.AuthMiddleware(jwtRepo)

	// Instanciar casos de uso (Use Cases)
	registerUseCase := application.NewRegister(db, bcryptRepo)
	loginUseCase := application.NewLogin(db, bcryptRepo, jwtRepo)
	getMeUseCase := application.NewGetMe(db)
	viewUsersUseCase := application.NewViewUsers(db)

	// Instanciar controladores (Handlers)
	registerController := controllers.NewRegisterController(registerUseCase)
	loginController := controllers.NewLoginController(loginUseCase)
	getMeController := controllers.NewGetMeController(getMeUseCase)
	viewUsersController := controllers.NewViewUsersController(viewUsersUseCase)

	// Configurar rutas
	SetupUserRoutes(router, authMiddleware, registerController, loginController, getMeController, viewUsersController)
}
