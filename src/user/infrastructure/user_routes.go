package infrastructure

import (
	"api/src/core"
	"api/src/user/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

const authRateLimitPerMinute = 10

func SetupUserRoutes(
	router *gin.Engine,
	authMiddleware gin.HandlerFunc,
	registerController *controllers.RegisterController,
	loginController *controllers.LoginController,
	refreshController *controllers.RefreshController,
	registerDeviceTokenController *controllers.RegisterDeviceTokenController,
	sendPushController *controllers.SendPushNotificationController,
	sendBroadcastController *controllers.SendBroadcastNotificationController,
	getMeController *controllers.GetMeController,
	viewUsersController *controllers.ViewUsersController,
) {
	rateLimit := core.AuthRateLimitMiddleware(authRateLimitPerMinute)

	authGroup := router.Group("/auth")
	{
		authGroup.GET("/register/requirements", registerController.GetPasswordRequirements)
		authGroup.POST("/register", rateLimit, registerController.Run)
		authGroup.POST("/login", rateLimit, loginController.Run)
		authGroup.POST("/refresh", rateLimit, refreshController.Run)
		authGroup.POST("/notifications/token", authMiddleware, registerDeviceTokenController.Run)
		authGroup.POST("/notifications/send", authMiddleware, rateLimit, sendPushController.Run)
		authGroup.POST("/notifications/broadcast", authMiddleware, rateLimit, sendBroadcastController.Run)
		authGroup.GET("/users", viewUsersController.Run)
		authGroup.GET("/me", authMiddleware, getMeController.Run)
	}
}
