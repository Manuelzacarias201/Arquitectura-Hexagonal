package controllers

import (
	"api/src/alumn/application"
	"api/src/core"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type EditAlumnController struct {
	alumnEditor *application.EditAlumn
}

func NewEditAlumnController(editor *application.EditAlumn) *EditAlumnController {
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

	var alumno struct {
		Name      string `json:"name"`
		Matricula string `json:"matricula"`
		Email     string `json:"email"`
		PhotoPath string `json:"photo_path"`
	}
	if err := c.ShouldBindJSON(&alumno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	photoPath := alumno.PhotoPath
	if strings.HasPrefix(strings.TrimSpace(alumno.PhotoPath), "data:image") {
		saved, saveErr := core.SaveBase64Image(alumno.PhotoPath, "uploads/alumns")
		if saveErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": saveErr.Error()})
			return
		}
		photoPath = saved
	}

	// Ejecutar la actualización en la base de datos
	err = cp.alumnEditor.Execute(id, alumno.Name, alumno.Matricula, alumno.Email, photoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Alumn updated successfully"})
}
