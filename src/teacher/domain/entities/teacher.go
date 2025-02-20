package entities

type Teacher struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Asignature string `json:"asignature"`
}

var increment = 0

func NewTeacher(name, asignature string) *Teacher {
	increment++
	teacher := Teacher{Id: increment, Name: name, Asignature: asignature}
	return &teacher
}
