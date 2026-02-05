package application

import (
	"api/src/user/domain"
	"errors"
)

type GetMe struct {
	db domain.IUser
}

func NewGetMe(db domain.IUser) *GetMe {
	return &GetMe{db: db}
}

// Execute retorna los datos del usuario por ID (sin contraseña)
func (g *GetMe) Execute(userID int) (*domain.UserResponse, error) {
	if userID <= 0 {
		return nil, errors.New("id de usuario inválido")
	}

	user, err := g.db.FindByID(userID)
	if err != nil || user == nil {
		return nil, errors.New("usuario no encontrado")
	}

	return &domain.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
