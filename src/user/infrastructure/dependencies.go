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
	fcmRepo := core.NewFCMRepository()
	if err != nil {
		log.Fatalf("Error al inicializar JWT: %v", err)
	}
	authMiddleware := core.AuthMiddleware(jwtRepo)

	// Instanciar casos de uso (Use Cases)
	registerUseCase := application.NewRegister(db, bcryptRepo)
	loginUseCase := application.NewLogin(db, bcryptRepo, jwtRepo)
	refreshUseCase := application.NewRefresh(db, jwtRepo)
	registerDeviceTokenUseCase := application.NewRegisterDeviceToken(db)
	sendPushUseCase := application.NewSendPushNotification(db, fcmRepo)
	getMeUseCase := application.NewGetMe(db)
	viewUsersUseCase := application.NewViewUsers(db)

	// Instanciar controladores (Handlers)
	registerController := controllers.NewRegisterController(registerUseCase)
	loginController := controllers.NewLoginController(loginUseCase)
	refreshController := controllers.NewRefreshController(refreshUseCase)
	registerDeviceTokenController := controllers.NewRegisterDeviceTokenController(registerDeviceTokenUseCase)
	sendPushController := controllers.NewSendPushNotificationController(sendPushUseCase)
	getMeController := controllers.NewGetMeController(getMeUseCase)
	viewUsersController := controllers.NewViewUsersController(viewUsersUseCase)

	// Configurar rutas
	SetupUserRoutes(
		router,
		authMiddleware,
		registerController,
		loginController,
		refreshController,
		registerDeviceTokenController,
		sendPushController,
		getMeController,
		viewUsersController,
	)
}
