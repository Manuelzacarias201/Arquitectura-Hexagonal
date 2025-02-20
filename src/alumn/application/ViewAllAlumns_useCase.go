package application

import (
	"api/src/alumn/domain"
	"api/src/alumn/domain/entities"
)

type ViewAlumns struct { //estruct ver todos los alumnos
	db domain.IAlumn
}

func NewViewAlumns(db domain.IAlumn) *ViewAlumns { //constructor para ver a todos
	return &ViewAlumns{db: db}
}

func (va *ViewAlumns) Execute() ([]entities.Alumn, error) { //metod para ver a todos los alumns
	return va.db.ViewAll()
}
