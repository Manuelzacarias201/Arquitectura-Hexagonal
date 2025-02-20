package application

import "api/src/alumn/domain"

type EditAccessory struct { //editar un alumno
	db domain.IAlumn
}

func NewEditAlumn(db domain.IAlumn) *EditAccessory { //fun para editar un alumno
	return &EditAccessory{db: db}
}

func (ep *EditAccessory) Execute(id int, name string, matricula string) error { // funcion para editar un alumn con id
	return ep.db.Edit(id, name, matricula)
}
