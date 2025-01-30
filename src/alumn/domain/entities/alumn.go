package entities

type Alumn struct {
	ID        int
	Name      string
	Matricula string
}

func NewAlumn(name, matricula string) *Alumn {
	//Id innecesario pq en la base de datos se le asigna uno
	alumn := Alumn{ID: 1, Name: name, Matricula: matricula}
	return &alumn
}
