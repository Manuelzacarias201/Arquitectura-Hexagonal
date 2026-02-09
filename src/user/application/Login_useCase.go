package application

import (
	"api/src/core"
	"api/src/user/domain"
	"fmt"
)

type Login struct {
	db     domain.IUser
	bcrypt *core.BcryptRepository
	jwt    *core.JWTRepository
}

func NewLogin(db domain.IUser, bcrypt *core.BcryptRepository, jwt *core.JWTRepository) *Login {
	return &Login{
		db:     db,
		bcrypt: bcrypt,
		jwt:    jwt,
	}
}

func (l *Login) Execute(email, password string) (string, string, *domain.LoginResponse, error) {
	email = NormalizeEmail(email)

	if email == "" {
		return "", "", nil, &AppError{Code: CodeEmailRequired, Message: "El correo es requerido"}
	}
	if !IsValidEmailFormat(email) {
		return "", "", nil, &AppError{Code: CodeInvalidEmail, Message: "El formato del correo no es válido. Ejemplo: usuario@correo.com"}
	}
	if password == "" {
		return "", "", nil, &AppError{Code: CodePasswordRequired, Message: "La contraseña es requerida"}
	}

	user, err := l.db.FindByEmail(email)
	if err != nil || user == nil {
		return "", "", nil, &AppError{Code: CodeEmailNotFound, Message: "El correo no está registrado"}
	}

	if err := l.bcrypt.ComparePassword(user.Password, password); err != nil {
		return "", "", nil, &AppError{Code: CodeWrongPassword, Message: "Contraseña incorrecta"}
	}

	accessToken, err := l.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", "", nil, fmt.Errorf("error al generar el token: %w", err)
	}

	refreshToken, err := l.jwt.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", nil, fmt.Errorf("error al generar el refresh token: %w", err)
	}

	response := &domain.LoginResponse{
		User: domain.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}

	return accessToken, refreshToken, response, nil
}
