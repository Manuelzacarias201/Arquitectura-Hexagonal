package application

import "api/src/alumn/domain"

type UpdateAlumnPhoto struct {
	db domain.IAlumn
}

func NewUpdateAlumnPhoto(db domain.IAlumn) *UpdateAlumnPhoto {
	return &UpdateAlumnPhoto{db: db}
}

func (u *UpdateAlumnPhoto) Execute(id int, photoPath string) error {
	return u.db.UpdatePhoto(id, photoPath)
}
