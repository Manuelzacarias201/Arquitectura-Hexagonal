package application

import (
	"api/src/core"
	"api/src/user/domain"
	"strings"
)

type SendPushNotification struct {
	db  domain.IUser
	fcm *core.FCMRepository
}

func NewSendPushNotification(db domain.IUser, fcm *core.FCMRepository) *SendPushNotification {
	return &SendPushNotification{db: db, fcm: fcm}
}

func (s *SendPushNotification) Execute(userID int, title, body string, data map[string]string) error {
	if userID <= 0 {
		return &AppError{Code: CodeInvalidRefresh, Message: "Usuario destino inválido"}
	}
	title = strings.TrimSpace(title)
	body = strings.TrimSpace(body)
	if title == "" || body == "" {
		return &AppError{Code: "INVALID_INPUT", Message: "title y body son requeridos"}
	}

	token, err := s.db.GetDeviceToken(userID)
	if err != nil {
		return &AppError{Code: CodePushTokenInvalid, Message: "El usuario no tiene token FCM registrado"}
	}

	if err := s.fcm.SendToToken(token, title, body, data); err != nil {
		return &AppError{Code: CodePushUnavailable, Message: "No se pudo enviar la notificación push"}
	}
	return nil
}
