package infrastructure

import (
	"api/src/core"
	"api/src/teacher/application"
	"api/src/teacher/infraestructure/controllers"
	userApplication "api/src/user/application"
	userInfrastructure "api/src/user/infrastructure"

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
	fcmRepo := core.NewFCMRepository()
	userDB := userInfrastructure.NewMySQL()
	sendBroadcastUseCase := userApplication.NewSendBroadcastNotification(userDB, fcmRepo)

	// Instanciar controladores (Handlers)
	addTeacherController := controllers.NewSaveTeacherController(teacherSaver, sendBroadcastUseCase)
	deleteTeacherController := controllers.NewDeleteTeacherController(teacherRemover)
	viewTeachersController := controllers.NewViewTeachersController(teacherViewer)
	viewTeacherController := controllers.NewViewTeacherController(teacherView)
	editTeacherController := controllers.NewEditTeacherController(teacherEDit)

	// Configurar rutas
	SetupTeacherRoutes(router, addTeacherController, deleteTeacherController, viewTeachersController, viewTeacherController, editTeacherController)
}
