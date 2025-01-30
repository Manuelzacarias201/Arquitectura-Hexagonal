package application

import (
	"api/src/teacher/domain"
)

type AddTeacher struct {
	db domain.ITteacher
}

func NewAddTeacher(db domain.ITteacher) *AddTeacher {
	return &AddTeacher{db: db}
}
func (at *AddTeacher) Execute(name, description string) error {
	return at.db.Save(name, description)
}
