package entities

type Teacher struct {
	Id         int
	Name       string
	Asignature string
}

var increment = 0

func NewTeacher(name, asignature string) *Teacher {
	increment++
	teacher := Teacher{Id: increment, Name: name, Asignature: asignature}
	return &teacher
}
