package controllers

import (
	"api/src/teacher/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateTeacherController struct {
	teacherSaver *application.AddTeacher
}

func NewSaveTeacherController(useCase *application.AddTeacher) *CreateTeacherController {
	return &CreateTeacherController{teacherSaver: useCase}
}

func (ts *CreateTeacherController) Run(c *gin.Context) {
	var body struct {
		Name       string `json:"name"`
		Asignature string `json:"asignature"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ts.teacherSaver.Execute(body.Name, body.Asignature)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Teacher added successfully"})
}
