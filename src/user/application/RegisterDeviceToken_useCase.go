package application

import (
	"api/src/user/domain"
	"strings"
)

type RegisterDeviceToken struct {
	db domain.IUser
}

func NewRegisterDeviceToken(db domain.IUser) *RegisterDeviceToken {
	return &RegisterDeviceToken{db: db}
}

func (r *RegisterDeviceToken) Execute(userID int, token string) error {
	if userID <= 0 {
		return &AppError{Code: CodeInvalidRefresh, Message: "Usuario inválido"}
	}
	token = strings.TrimSpace(token)
	if token == "" {
		return &AppError{Code: CodePushTokenInvalid, Message: "El token de dispositivo FCM es requerido"}
	}
	return r.db.SaveDeviceToken(userID, token)
}
