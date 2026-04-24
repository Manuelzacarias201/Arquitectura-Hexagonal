package controllers

import (
	"api/src/user/application"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SendPushNotificationController struct {
	useCase *application.SendPushNotification
}

func NewSendPushNotificationController(useCase *application.SendPushNotification) *SendPushNotificationController {
	return &SendPushNotificationController{useCase: useCase}
}

func (sc *SendPushNotificationController) Run(c *gin.Context) {
	var body struct {
		UserID int               `json:"user_id"`
		Title  string            `json:"title"`
		Body   string            `json:"body"`
		Data   map[string]string `json:"data"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		RespondError(c, http.StatusBadRequest, MsgInvalidInput, CodeInvalidInput)
		return
	}

	if err := sc.useCase.Execute(body.UserID, body.Title, body.Body, body.Data); err != nil {
		var appErr *application.AppError
		if errors.As(err, &appErr) {
			RespondError(c, http.StatusBadRequest, appErr.Message, appErr.Code)
			return
		}
		RespondError(c, http.StatusInternalServerError, "No se pudo enviar la notificación push", CodeInternalError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Notificación enviada"})
}
