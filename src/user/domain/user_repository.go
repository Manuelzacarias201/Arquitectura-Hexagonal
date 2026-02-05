package domain

import "api/src/user/domain/entities"

type IUser interface {
	Save(email, hashedPassword, name string) error
	FindByEmail(email string) (*entities.User, error)
	FindByID(id int) (*entities.User, error)
	ViewAll() ([]entities.User, error)
}

// UserResponse representa los datos del usuario sin la contraseña
type UserResponse struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// LoginResponse representa la respuesta del login
type LoginResponse struct {
	User UserResponse `json:"user"`
}
