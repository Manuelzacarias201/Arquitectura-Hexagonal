package application

import "api/src/alumn/domain"

type EditAlumn struct {
	db domain.IAlumn
}

func NewEditAlumn(db domain.IAlumn) *EditAlumn {
	return &EditAlumn{db: db}
}

func (ep *EditAlumn) Execute(id int, name, matricula, email, photoPath string) error {
	return ep.db.Edit(id, name, matricula, email, photoPath)
}
