package domain

import "api/src/alumn/domain/entities"

type IAlumn interface { //interfaz para guardar, ver uno, ver todos, eliminar y editar
	Save(name, matricula string) error
	ViewOne(id int) (*entities.Alumn, error)
	ViewAll() ([]entities.Alumn, error)
	Delete(id int) error
	Edit(id int, name, matricula string) error
}
