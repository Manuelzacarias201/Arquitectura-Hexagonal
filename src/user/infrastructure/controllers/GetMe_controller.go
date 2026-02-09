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
		RespondError(c, http.StatusUnauthorized, "Sesión no encontrada. Inicia sesión de nuevo.", "MISSING_SESSION")
		return
	}

	claims, ok := claimsVal.(*core.Claims)
	if !ok {
		RespondError(c, http.StatusInternalServerError, "Ha ocurrido un error. Inténtalo más tarde.", CodeInternalError)
		return
	}

	user, err := gm.getMe.Execute(claims.UserID)
	if err != nil {
		RespondError(c, http.StatusNotFound, "Usuario no encontrado.", "USER_NOT_FOUND")
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}
