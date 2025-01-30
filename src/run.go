package src

import (
	infrastructureAlumns "api/src/alumn/infrastructure"
	infrastructureTeachers "api/src/teacher/infraestructure"

	"github.com/gin-gonic/gin"
)

func Init() {
	// cargamos las dependencias de cada modulo
	dbT := infrastructureTeachers.NewMySQL()
	dbA := infrastructureAlumns.NewMySQL()

	router := gin.Default()
	// inicializamos los modulos
	infrastructureTeachers.InitTeachers(dbT, router)
	infrastructureAlumns.InitAlumns(dbA, router)
	// corremos el servidor
	router.Run(":8080")
}
