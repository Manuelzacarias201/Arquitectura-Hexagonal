package application

import (
	"api/src/core"
	"api/src/user/domain"
	"fmt"
)

type Refresh struct {
	db  domain.IUser
	jwt *core.JWTRepository
}

func NewRefresh(db domain.IUser, jwt *core.JWTRepository) *Refresh {
	return &Refresh{db: db, jwt: jwt}
}

func (rf *Refresh) Execute(refreshToken string) (string, string, *domain.LoginResponse, error) {
	if refreshToken == "" {
		return "", "", nil, &AppError{Code: CodeInvalidRefresh, Message: "El refresh token es requerido"}
	}

	claims, err := rf.jwt.ValidateToken(refreshToken)
	if err != nil {
		return "", "", nil, &AppError{Code: CodeInvalidRefresh, Message: "Refresh token inválido o expirado"}
	}

	user, err := rf.db.FindByID(claims.UserID)
	if err != nil || user == nil {
		return "", "", nil, &AppError{Code: CodeInvalidRefresh, Message: "Usuario no encontrado"}
	}

	accessToken, err := rf.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", "", nil, fmt.Errorf("error al generar token: %w", err)
	}

	newRefreshToken, err := rf.jwt.GenerateRefreshToken(user.ID, user.Email)
	if err != nil {
		return "", "", nil, fmt.Errorf("error al generar refresh token: %w", err)
	}

	response := &domain.LoginResponse{
		User: domain.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}

	return accessToken, newRefreshToken, response, nil
}
