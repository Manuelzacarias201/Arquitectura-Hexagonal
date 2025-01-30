package domain

import "api/src/teacher/domain/entities"

type ITteacher interface {
	Save(name, description string) error
	ViewOne(id int) (*entities.Teacher, error)
	ViewAll() ([]entities.Teacher, error)
	Delete(id int) error
}
