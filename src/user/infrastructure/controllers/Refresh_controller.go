package controllers

import (
	"api/src/user/application"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshController struct {
	refresh *application.Refresh
}

func NewRefreshController(refresh *application.Refresh) *RefreshController {
	return &RefreshController{refresh: refresh}
}

func (rc *RefreshController) Run(c *gin.Context) {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		RespondError(c, http.StatusBadRequest, MsgInvalidInput, CodeInvalidInput)
		return
	}

	accessToken, refreshToken, userResponse, err := rc.refresh.Execute(body.RefreshToken)
	if err != nil {
		var appErr *application.AppError
		if errors.As(err, &appErr) {
			RespondError(c, http.StatusUnauthorized, appErr.Message, appErr.Code)
			return
		}
		RespondError(c, http.StatusInternalServerError, "Ha ocurrido un error. Inténtalo más tarde.", CodeInternalError)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Token renovado",
		"token":         accessToken,
		"refresh_token": refreshToken,
		"user":          userResponse.User,
	})
}
