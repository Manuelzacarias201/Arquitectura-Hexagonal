package controllers

import (
	"api/src/teacher/application"
	"api/src/teacher/domain/entities"
	"net/http"

	"strconv"

	"github.com/gin-gonic/gin"
)

type EditTeacherController struct {
	teacherEditor *application.EditTeacher
}

func NewEditTeacherController(editor *application.EditTeacher) *EditTeacherController {
	return &EditTeacherController{
		teacherEditor: editor,
	}
}

func (cp *EditTeacherController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid teacher ID"})
		return
	}

	var teacher entities.Teacher
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	// Ejecutar la actualización en la base de datos
	err = cp.teacherEditor.Execute(id, teacher.Name, teacher.Asignature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Teacher updated successfully"})
}
