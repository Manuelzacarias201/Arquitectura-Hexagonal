package controllers

import (
	"net/http"
	"strconv"

	"api/src/teacher/application"

	"github.com/gin-gonic/gin"
)

type ViewTeacherController struct {
	teacherView *application.ViewTeacher
}

func NewViewTeacherController(useCase *application.ViewTeacher) *ViewTeacherController {
	return &ViewTeacherController{teacherView: useCase}
}

func (tv *ViewTeacherController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Teacher ID"})
		return
	}

	teacher, err := tv.teacherView.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Teacher": teacher})
}
