package application // Parte de la logica de el negocio y la app

import "api/src/alumn/domain"

type DeleteAlumn struct {
	db domain.IAlumn //Interfaz para la bd
}

func NewDeleteAlumn(db domain.IAlumn) *DeleteAlumn {
	return &DeleteAlumn{db: db} //puntero a DeleteAlumn
}

func (da *DeleteAlumn) Execute(id int) error { //funcion para eliminar un alumno con id
	return da.db.Delete(id)
}
