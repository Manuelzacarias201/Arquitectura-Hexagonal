package application

import "api/src/teacher/domain"

type DeleteTeacher struct {
	db domain.ITteacher
}

func NewDeleteTeacher(db domain.ITteacher) *DeleteTeacher {
	return &DeleteTeacher{db: db}
}

func (dt *DeleteTeacher) Execute(id int) error {
	return dt.db.Delete(id)
}
