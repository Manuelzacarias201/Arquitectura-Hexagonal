package infrastructure

import (
	"api/src/alumn/application"
	"api/src/alumn/infrastructure/controllers"
	"api/src/core"
	userApplication "api/src/user/application"
	userInfrastructure "api/src/user/infrastructure"

	"github.com/gin-gonic/gin"
)

func InitAlumns(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE ALUMNOS")

	// Instanciar casos de uso (Use Cases)
	alumnSaver := application.NewSaveAlumn(db)
	alumnRemover := application.NewDeleteAlumn(db)
	alumnViewer := application.NewViewAlumns(db)
	alumnView := application.NewViewAlumn(db)
	alumnEdit := application.NewEditAlumn(db)
	alumnUpdatePhoto := application.NewUpdateAlumnPhoto(db)
	fcmRepo := core.NewFCMRepository()
	userDB := userInfrastructure.NewMySQL()
	sendBroadcastUseCase := userApplication.NewSendBroadcastNotification(userDB, fcmRepo)

	// Instanciar controladores (Handlers)
	addAlumnController := controllers.NewSaveAlumnController(alumnSaver, sendBroadcastUseCase)
	deleteAlumnController := controllers.NewRemoveAlumnController(alumnRemover)
	viewAlumnsController := controllers.NewViewAllAlumnsController(alumnViewer)
	viewAlumnController := controllers.NewViewOneAlumnController(alumnView)
	editAlumnController := controllers.NewEditAlumnController(alumnEdit)
	uploadAlumnPhotoController := controllers.NewUploadAlumnPhotoController(alumnUpdatePhoto)

	// Configurar rutas
	SetupAlumnRoutes(router, addAlumnController, deleteAlumnController, viewAlumnsController, viewAlumnController, editAlumnController, uploadAlumnPhotoController)
}
