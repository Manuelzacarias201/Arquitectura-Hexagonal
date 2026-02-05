package controllers

import (
	"api/src/core"
	"api/src/user/application"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetMeController struct {
	getMe *application.GetMe
}

func NewGetMeController(getMe *application.GetMe) *GetMeController {
	return &GetMeController{getMe: getMe}
}

func (gm *GetMeController) Run(c *gin.Context) {
	claimsVal, exists := c.Get(string(core.ClaimsContextKey))
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "no se encontró la sesión"})
		return
	}

	claims, ok := claimsVal.(*core.Claims)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error interno"})
		return
	}

	user, err := gm.getMe.Execute(claims.UserID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
