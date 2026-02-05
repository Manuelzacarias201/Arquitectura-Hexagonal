package application

import (
	"api/src/core"
	"api/src/user/domain"
	"errors"
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

func (l *Login) Execute(email, password string) (string, *domain.LoginResponse, error) {
	// Validar que el email no esté vacío
	if email == "" {
		return "", nil, errors.New("el email es requerido")
	}

	// Validar que la contraseña no esté vacía
	if password == "" {
		return "", nil, errors.New("la contraseña es requerida")
	}

	// Buscar el usuario por email
	user, err := l.db.FindByEmail(email)
	if err != nil || user == nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// Verificar la contraseña
	err = l.bcrypt.ComparePassword(user.Password, password)
	if err != nil {
		return "", nil, errors.New("credenciales inválidas")
	}

	// Generar token JWT
	token, err := l.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return "", nil, fmt.Errorf("error al generar el token: %v", err)
	}

	// Retornar token y datos del usuario (sin contraseña)
	response := &domain.LoginResponse{
		User: domain.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}

	return token, response, nil
}
