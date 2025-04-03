package core

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// BcryptRepository maneja la encriptación y validación de contraseñas
type BcryptRepository struct{}

// NewBcryptRepository crea una nueva instancia del repositorio bcrypt
func NewBcryptRepository() *BcryptRepository {
	return &BcryptRepository{}
}

// HashPassword genera un hash para la contraseña (matrícula)
func (br *BcryptRepository) HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

// ComparePassword verifica si una contraseña coincide con su hash
func (br *BcryptRepository) ComparePassword(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("contraseña incorrecta")
	}
	return nil
}
