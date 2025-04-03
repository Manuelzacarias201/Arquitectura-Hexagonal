package domain

import "api/src/alumn/domain/entities"

type IAlumn interface {
	Save(name, hashedMatricula string) error // Ahora recibe la matrícula encriptada
	ViewOne(id int) (*entities.Alumn, error)
	ViewAll() ([]entities.Alumn, error)
	Delete(id int) error
	Edit(id int, name, hashedMatricula string) error
}
