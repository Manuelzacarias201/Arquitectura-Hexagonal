package application

import "api/src/alumn/domain"

type DeleteAlumn struct {
	db domain.IAlumn
}

func NewDeleteAlumn(db domain.IAlumn) *DeleteAlumn {
	return &DeleteAlumn{db: db}
}

func (da *DeleteAlumn) Execute(id int) error {
	return da.db.Delete(id)
}
