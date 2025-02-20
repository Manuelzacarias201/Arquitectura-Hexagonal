package controllers

import (
	"api/src/alumn/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddAlumnController struct { //estructura para guardar alumn
	alumnSaver *application.SaveAlumn
}

func NewSaveAlumnController(useCase *application.SaveAlumn) *AddAlumnController { //constructor para la bd
	return &AddAlumnController{alumnSaver: useCase}
}

func (as *AddAlumnController) Run(c *gin.Context) {
	//definimos el cuerpo de la petición
	var body struct {
		Name      string `json:"name"`
		Matricula string `json:"matricula"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := as.alumnSaver.Execute(body.Name, body.Matricula)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Alumn saved successfully"})
}
