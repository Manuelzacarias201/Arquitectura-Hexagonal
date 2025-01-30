package application

import (
	"api/src/teacher/domain"
	"api/src/teacher/domain/entities"
)

type ViewTeachers struct {
	db domain.ITteacher
}

func NewViewTeachers(db domain.ITteacher) *ViewTeachers {
	return &ViewTeachers{db: db}
}

func (vt *ViewTeachers) Execute() ([]entities.Teacher, error) {
	return vt.db.ViewAll()
}
