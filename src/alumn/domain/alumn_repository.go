package domain

import "api/src/alumn/domain/entities"

type IAlumn interface {
	Save(name, matricula string) error
	ViewOne(id int) (*entities.Alumn, error)
	ViewAll() ([]entities.Alumn, error)
	Delete(id int) error
	Edit(id int, name, matricula string) error
}
