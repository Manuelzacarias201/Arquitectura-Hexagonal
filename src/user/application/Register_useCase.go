package application

import (
	"api/src/core"
	"api/src/user/domain"
	"errors"
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

func (r *Register) Execute(email, password, name string) error {
	// Validar que el email no esté vacío
	if email == "" {
		return errors.New("el email es requerido")
	}

	// Validar que la contraseña no esté vacía
	if password == "" {
		return errors.New("la contraseña es requerida")
	}

	// Validar que el nombre no esté vacío
	if name == "" {
		return errors.New("el nombre es requerido")
	}

	// Validar longitud mínima de contraseña
	if len(password) < 6 {
		return errors.New("la contraseña debe tener al menos 6 caracteres")
	}

	// Verificar si el usuario ya existe
	existingUser, err := r.db.FindByEmail(email)
	if err == nil && existingUser != nil {
		return errors.New("el email ya está registrado")
	}

	// Encriptar la contraseña
	hashedPassword, err := r.bcrypt.HashPassword(password)
	if err != nil {
		return fmt.Errorf("error al encriptar la contraseña: %v", err)
	}

	// Guardar el usuario
	return r.db.Save(email, hashedPassword, name)
}
