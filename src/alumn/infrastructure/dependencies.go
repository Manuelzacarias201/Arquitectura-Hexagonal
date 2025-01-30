package infrastructure

import (
	"api/src/alumn/application"
	"api/src/alumn/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitAlumns(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE ALUMNOS")

	// Instanciar casos de uso (Use Cases)
	alumnSaver := application.NewSaveAlumn(db)
	alumnRemover := application.NewDeleteAlumn(db)
	alumnViewer := application.NewViewAlumns(db)
	alumnView := application.NewViewAlumn(db)

	// Instanciar controladores (Handlers)
	addAlumnController := controllers.NewSaveAlumnController(alumnSaver)
	deleteAlumnController := controllers.NewRemoveAlumnController(alumnRemover)
	viewAlumnsController := controllers.NewViewAllAlumnsController(alumnViewer)
	viewAlumnController := controllers.NewViewOneAlumnController(alumnView)

	// Configurar rutas
	SetupAlumnRoutes(router, addAlumnController, deleteAlumnController, viewAlumnsController, viewAlumnController)

}
