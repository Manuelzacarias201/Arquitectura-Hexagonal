package application

import (
	"api/src/alumn/domain"
	"api/src/core" // Importamos la encriptación
	"fmt"
)

type SaveAlumn struct {
	db     domain.IAlumn
	bcrypt *core.BcryptRepository
}

func NewSaveAlumn(db domain.IAlumn, bcrypt *core.BcryptRepository) *SaveAlumn {
	return &SaveAlumn{
		db:     db,
		bcrypt: bcrypt,
	}
}

func (sa *SaveAlumn) Execute(name, matricula string) error {
	// Encriptar la matrícula antes de guardarla
	hashedMatricula, err := sa.bcrypt.HashPassword(matricula)
	if err != nil {
		return fmt.Errorf("error al encriptar la matrícula: %v", err)
	}

	// Guardar el alumno con la matrícula encriptada
	return sa.db.Save(name, hashedMatricula)
}
