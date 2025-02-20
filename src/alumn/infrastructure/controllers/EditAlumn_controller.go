package controllers

import (
	"api/src/alumn/application"
	"api/src/alumn/domain/entities"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type EditAlumnController struct {
	alumnEditor *application.EditAccessory
}

func NewEditAlumnController(editor *application.EditAccessory) *EditAlumnController {
	return &EditAlumnController{
		alumnEditor: editor,
	}
}

func (cp *EditAlumnController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid alumn ID"})
		return
	}

	var alumno entities.Alumn
	if err := c.ShouldBindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Ejecutar la actualización en la base de datos
	err = cp.alumnEditor.Execute(id, alumno.Name, alumno.Matricula)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumn updated successfully"})
}
