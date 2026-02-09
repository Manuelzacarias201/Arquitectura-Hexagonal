package application

import (
	"api/src/core"
	"api/src/user/domain"
	"fmt"
)

type Register struct {
	db     domain.IUser
	bcrypt *core.BcryptRepository
}

func NewRegister(db domain.IUser, bcrypt *core.BcryptRepository) *Register {
	return &Register{
		db:     db,
		bcrypt: bcrypt,
	}
}

const passwordExample = "MiSegura123"

func (r *Register) Execute(email, password, name string) (*domain.UserResponse, error) {
	email = NormalizeEmail(email)

	if email == "" {
		return nil, &AppError{Code: CodeEmailRequired, Message: "El correo es requerido"}
	}
	if !IsValidEmailFormat(email) {
		return nil, &AppError{Code: CodeInvalidEmail, Message: "El formato del correo no es válido. Ejemplo: usuario@correo.com"}
	}
	if password == "" {
		return nil, &AppError{
			Code:    CodePasswordRequired,
			Message: "La contraseña es obligatoria. Debe tener al menos 8 caracteres, una letra y un número. Ejemplo: " + passwordExample,
		}
	}
	if name == "" {
		return nil, &AppError{Code: CodeNameRequired, Message: "El nombre es requerido"}
	}
	if !ValidatePasswordStrength(password) {
		return nil, &AppError{
			Code:    CodePasswordTooWeak,
			Message: "La contraseña debe tener al menos 8 caracteres, una letra y un número. Ejemplo: " + passwordExample,
		}
	}

	existingUser, err := r.db.FindByEmail(email)
	if err == nil && existingUser != nil {
		return nil, &AppError{Code: CodeEmailTaken, Message: "El correo ya está registrado"}
	}

	hashedPassword, err := r.bcrypt.HashPassword(password)
	if err != nil {
		return nil, fmt.Errorf("error al encriptar la contraseña: %w", err)
	}

	id, err := r.db.Save(email, hashedPassword, name)
	if err != nil {
		return nil, err
	}

	user, err := r.db.FindByID(id)
	if err != nil || user == nil {
		return &domain.UserResponse{ID: id, Email: email, Name: name}, nil
	}

	return &domain.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Name:  user.Name,
	}, nil
}
