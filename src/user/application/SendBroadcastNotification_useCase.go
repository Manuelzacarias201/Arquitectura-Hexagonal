package application

import (
	"api/src/core"
	"api/src/user/domain"
	"strings"
)

type SendBroadcastNotification struct {
	db  domain.IUser
	fcm *core.FCMRepository
}

func NewSendBroadcastNotification(db domain.IUser, fcm *core.FCMRepository) *SendBroadcastNotification {
	return &SendBroadcastNotification{db: db, fcm: fcm}
}

func (s *SendBroadcastNotification) Execute(title, body string, data map[string]string) error {
	title = strings.TrimSpace(title)
	body = strings.TrimSpace(body)
	if title == "" || body == "" {
		return &AppError{Code: "INVALID_INPUT", Message: "title y body son requeridos"}
	}

	tokens, err := s.db.GetAllDeviceTokens()
	if err != nil {
		return &AppError{Code: CodePushUnavailable, Message: "No se pudieron obtener tokens de dispositivos"}
	}
	if len(tokens) == 0 {
		return &AppError{Code: CodePushTokenInvalid, Message: "No hay dispositivos registrados para notificar"}
	}

	if data == nil {
		data = map[string]string{}
	}
	if _, ok := data["android_channel_id"]; !ok {
		data["android_channel_id"] = "canal_estudiantes"
	}

	if err := s.fcm.SendToTokens(tokens, title, body, data); err != nil {
		return &AppError{Code: CodePushUnavailable, Message: "No se pudo enviar la notificación push masiva"}
	}
	return nil
}
