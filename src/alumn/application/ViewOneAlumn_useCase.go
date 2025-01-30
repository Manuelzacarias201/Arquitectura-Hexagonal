package application

import (
	"api/src/alumn/domain"
	"api/src/alumn/domain/entities"
)

type ViewAlumn struct {
	db domain.IAlumn
}

func NewViewAlumn(db domain.IAlumn) *ViewAlumn {
	return &ViewAlumn{db: db}
}

func (va *ViewAlumn) Execute(id int) (*entities.Alumn, error) {
	return va.db.ViewOne(id)
}
