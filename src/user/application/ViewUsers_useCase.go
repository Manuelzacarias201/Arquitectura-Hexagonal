package application

import (
	"api/src/user/domain"
)

type ViewUsers struct {
	db domain.IUser
}

func NewViewUsers(db domain.IUser) *ViewUsers {
	return &ViewUsers{db: db}
}

// Execute retorna la lista de todos los usuarios registrados (sin contraseñas)
func (v *ViewUsers) Execute() ([]domain.UserResponse, error) {
	users, err := v.db.ViewAll()
	if err != nil {
		return nil, err
	}

	result := make([]domain.UserResponse, 0, len(users))
	for _, u := range users {
		result = append(result, domain.UserResponse{
			ID:    u.ID,
			Email: u.Email,
			Name:  u.Name,
		})
	}
	return result, nil
}
