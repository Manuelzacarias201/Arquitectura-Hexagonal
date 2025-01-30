package application

import (
	"api/src/teacher/domain"
	"api/src/teacher/domain/entities"
)

type ViewTeacher struct {
	db domain.ITteacher
}

func NewViewTeacher(db domain.ITteacher) *ViewTeacher {
	return &ViewTeacher{db: db}
}

func (vt *ViewTeacher) Execute(id int) (*entities.Teacher, error) {
	return vt.db.ViewOne(id)
}
