package controllers

import (
	"api/src/alumn/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewAllAlumnsController struct {
	alumnsViewer *application.ViewAlumns
}

func NewViewAllAlumnsController(useCase *application.ViewAlumns) *ViewAllAlumnsController {
	return &ViewAllAlumnsController{alumnsViewer: useCase}
}

func (av *ViewAllAlumnsController) Run(c *gin.Context) {
	alumns, err := av.alumnsViewer.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Alumns": alumns})
}
