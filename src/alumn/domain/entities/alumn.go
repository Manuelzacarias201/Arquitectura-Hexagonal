package entities

type Alumn struct { //modelo de alumno
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Matricula string `json:"matricula"`
}

func NewAlumn(name, matricula string) *Alumn { //funcion para crear un alumno
	//Id innecesario pq en la base de datos se le asigna uno
	alumn := Alumn{ID: 1, Name: name, Matricula: matricula} //crea un alumno con nombre y matricula
	return &alumn
}
