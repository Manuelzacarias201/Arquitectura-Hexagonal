package src

import (
	infrastructureAlumns "api/src/alumn/infrastructure"
	infrastructureTeachers "api/src/teacher/infraestructure"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
)

func Init() {
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:4200"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	router := gin.Default()

	router.Use(cors.New(config))

	// cargamos las dependencias de cada modulo
	dbT := infrastructureTeachers.NewMySQL()
	dbA := infrastructureAlumns.NewMySQL()

	// inicializamos los modulos
	infrastructureTeachers.InitTeachers(dbT, router)
	infrastructureAlumns.InitAlumns(dbA, router)
	// corremos el servidor
	router.Run(":8080")
}
