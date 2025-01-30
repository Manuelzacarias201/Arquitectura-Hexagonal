package controllers

import (
	"api/src/alumn/application"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RemoveAlumnController struct {
	AlumnRemover *application.DeleteAlumn
}

func NewRemoveAlumnController(useCase *application.DeleteAlumn) *RemoveAlumnController {
	return &RemoveAlumnController{AlumnRemover: useCase}
}

func (ar *RemoveAlumnController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Alumn ID"})
		return
	}
	err = ar.AlumnRemover.Execute(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Alumn removed successfully"})
}
