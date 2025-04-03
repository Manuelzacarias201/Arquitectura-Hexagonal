package application

import "api/src/alumn/domain"

type DeleteAlumn struct { //estructura para eliminar un alumno
	db domain.IAlumn
}

func NewDeleteAlumn(db domain.IAlumn) *DeleteAlumn { //puntero de delete alumn
	return &DeleteAlumn{db: db}
}

func (da *DeleteAlumn) Execute(id int) error { //funcion para eliminar un alumno con id
	return da.db.Delete(id)
}
