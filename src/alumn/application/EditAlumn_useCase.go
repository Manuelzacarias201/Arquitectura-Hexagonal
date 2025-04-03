package application

import "api/src/alumn/domain"

type EditAlumn struct { //estructura para editar un alumno
	db domain.IAlumn
}

func NewEditAlumn(db domain.IAlumn) *EditAlumn { //fun para editar un alumno
	return &EditAlumn{db: db}
}

func (ep *EditAlumn) Execute(id int, name string, matricula string) error { // funcion para editar un alumn con id
	return ep.db.Edit(id, name, matricula)
}
