package controllers

import (
	"api/src/teacher/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteTeacherController struct {
	teacherRemover *application.DeleteTeacher
}

func NewDeleteTeacherController(useCase *application.DeleteTeacher) *DeleteTeacherController {
	return &DeleteTeacherController{teacherRemover: useCase}
}

func (tr *DeleteTeacherController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Teacher ID"})
		return
	}
	err = tr.teacherRemover.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher deleted successfully"})
}
