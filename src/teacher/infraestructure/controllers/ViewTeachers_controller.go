package controllers

import (
	"api/src/teacher/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewTeachersController struct {
	teacherVIewer *application.ViewTeachers
}

func NewViewTeachersController(useCase *application.ViewTeachers) *ViewTeachersController {
	return &ViewTeachersController{teacherVIewer: useCase}
}

func (tv *ViewTeachersController) Run(c *gin.Context) {
	teachers, err := tv.teacherVIewer.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, teachers)
}
