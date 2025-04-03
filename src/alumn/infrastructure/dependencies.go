package infrastructure

import (
	"api/src/alumn/application"
	"api/src/alumn/infrastructure/controllers"
	"api/src/core" // Agregamos bcrypt

	"github.com/gin-gonic/gin"
)

func InitAlumns(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE ALUMNOS")

	// Instanciar el repositorio de bcrypt
	bcryptRepo := core.NewBcryptRepository()

	// Instanciar casos de uso (Use Cases)
	alumnSaver := application.NewSaveAlumn(db, bcryptRepo) // Se añade bcryptRepo
	alumnRemover := application.NewDeleteAlumn(db)
	alumnViewer := application.NewViewAlumns(db)
	alumnView := application.NewViewAlumn(db)
	alumnEdit := application.NewEditAlumn(db)

	// Instanciar controladores (Handlers)
	addAlumnController := controllers.NewSaveAlumnController(alumnSaver)
	deleteAlumnController := controllers.NewRemoveAlumnController(alumnRemover)
	viewAlumnsController := controllers.NewViewAllAlumnsController(alumnViewer)
	viewAlumnController := controllers.NewViewOneAlumnController(alumnView)
	editAlumnController := controllers.NewEditAlumnController(alumnEdit)

	// Configurar rutas
	SetupAlumnRoutes(router, addAlumnController, deleteAlumnController, viewAlumnsController, viewAlumnController, editAlumnController)
}
