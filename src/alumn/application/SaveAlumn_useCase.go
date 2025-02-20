package application

import (
	"api/src/alumn/domain"
)

type SaveAlumn struct { //estruc para guardar alumn
	db domain.IAlumn
}

func NewSaveAlumn(db domain.IAlumn) *SaveAlumn { //constructor para guardar alumn
	return &SaveAlumn{db: db}
}
func (sa *SaveAlumn) Execute(name, matricula string) error { //funcion para guardar
	return sa.db.Save(name, matricula)
}
