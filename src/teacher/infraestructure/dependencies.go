package infrastructure

import (
	"api/src/teacher/application"
	"api/src/teacher/infraestructure/controllers"

	"github.com/gin-gonic/gin"
)

func InitTeachers(db *MySQL, router *gin.Engine) {
	println("CARGANDO DEPENDENCIAS DE TEACHERS")

	// Instanciar casos de uso (Use Cases)
	teacherSaver := application.NewAddTeacher(db)
	teacherRemover := application.NewDeleteTeacher(db)
	teacherViewer := application.NewViewTeachers(db)
	teacherView := application.NewViewTeacher(db)
	teacherEDit := application.NewEditTeacher(db)

	// Instanciar controladores (Handlers)
	addTeacherController := controllers.NewSaveTeacherController(teacherSaver)
	deleteTeacherController := controllers.NewDeleteTeacherController(teacherRemover)
	viewTeachersController := controllers.NewViewTeachersController(teacherViewer)
	viewTeacherController := controllers.NewViewTeacherController(teacherView)
	editTeacherController := controllers.NewEditTeacherController(teacherEDit)

	// Configurar rutas
	SetupTeacherRoutes(router, addTeacherController, deleteTeacherController, viewTeachersController, viewTeacherController, editTeacherController)
}
