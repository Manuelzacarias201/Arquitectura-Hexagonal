package application

import "api/src/teacher/domain"

type EditTeacher struct {
	db domain.ITteacher
}

func NewEditTeacher(db domain.ITteacher) *EditTeacher {
	return &EditTeacher{db: db}
}

func (ep *EditTeacher) Execute(id int, name string, asignature string) error {
	return ep.db.Edit(id, name, asignature)
}
