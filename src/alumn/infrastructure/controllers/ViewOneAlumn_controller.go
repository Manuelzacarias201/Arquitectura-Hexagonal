package controllers

import (
	"net/http"
	"strconv"

	"api/src/alumn/application"

	"github.com/gin-gonic/gin"
)

type ViewOneAlumnController struct {
	alumnVIew *application.ViewAlumn
}

func NewViewOneAlumnController(useCase *application.ViewAlumn) *ViewOneAlumnController {
	return &ViewOneAlumnController{alumnVIew: useCase}
}

func (av *ViewOneAlumnController) Run(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Alumn ID"})
		return
	}

	alumn, err := av.alumnVIew.Execute(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alumn)
}
