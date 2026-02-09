package controllers

import (
	"api/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ViewUsersController struct {
	viewUsers *application.ViewUsers
}

func NewViewUsersController(viewUsers *application.ViewUsers) *ViewUsersController {
	return &ViewUsersController{viewUsers: viewUsers}
}

func (vu *ViewUsersController) Run(c *gin.Context) {
	users, err := vu.viewUsers.Execute()
	if err != nil {
		RespondError(c, http.StatusInternalServerError, "No se pudieron cargar los usuarios. Inténtalo más tarde.", CodeInternalError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
