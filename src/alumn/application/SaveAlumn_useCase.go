package application

import (
	"api/src/alumn/domain"
)

type SaveAlumn struct {
	db domain.IAlumn
}

func NewSaveAlumn(db domain.IAlumn) *SaveAlumn {
	return &SaveAlumn{
		db: db,
	}
}

func (sa *SaveAlumn) Execute(name, matricula string) error {
	// Guardar el alumno con la matrícula en texto plano (visible tal cual)
	return sa.db.Save(name, matricula)
}
