package application

import (
	"api/src/alumn/domain"
	"api/src/alumn/domain/entities"
)

type ViewAlumn struct { //interf para ver un alumno
	db domain.IAlumn
}

func NewViewAlumn(db domain.IAlumn) *ViewAlumn { //constructor para ver un alumno
	return &ViewAlumn{db: db}
}

func (va *ViewAlumn) Execute(id int) (*entities.Alumn, error) { //metodo para ver un alumno
	return va.db.ViewOne(id)
}
