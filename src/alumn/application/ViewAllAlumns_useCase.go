package application

import (
	"api/src/alumn/domain"
	"api/src/alumn/domain/entities"
)

type ViewAlumns struct {
	db domain.IAlumn
}

func NewViewAlumns(db domain.IAlumn) *ViewAlumns {
	return &ViewAlumns{db: db}
}

func (va *ViewAlumns) Execute() ([]entities.Alumn, error) {
	return va.db.ViewAll()
}
