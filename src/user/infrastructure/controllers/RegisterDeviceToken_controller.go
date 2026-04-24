package controllers

import (
	"api/src/core"
	"api/src/user/application"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterDeviceTokenController struct {
	useCase *application.RegisterDeviceToken
}

func NewRegisterDeviceTokenController(useCase *application.RegisterDeviceToken) *RegisterDeviceTokenController {
	return &RegisterDeviceTokenController{useCase: useCase}
}

func (rc *RegisterDeviceTokenController) Run(c *gin.Context) {
	claimsVal, exists := c.Get(string(core.ClaimsContextKey))
	if !exists {
		RespondError(c, http.StatusUnauthorized, "Sesión no encontrada. Inicia sesión de nuevo.", "MISSING_SESSION")
		return
	}

	claims, ok := claimsVal.(*core.Claims)
	if !ok {
		RespondError(c, http.StatusUnauthorized, "Sesión inválida.", "INVALID_SESSION")
		return
	}

	var body struct {
		Token string `json:"token"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		RespondError(c, http.StatusBadRequest, MsgInvalidInput, CodeInvalidInput)
		return
	}

	if err := rc.useCase.Execute(claims.UserID, body.Token); err != nil {
		var appErr *application.AppError
		if errors.As(err, &appErr) {
			RespondError(c, http.StatusBadRequest, appErr.Message, appErr.Code)
			return
		}
		RespondError(c, http.StatusInternalServerError, "No se pudo registrar el token FCM", CodeInternalError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token FCM registrado"})
}
