package entities

type Alumn struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Matricula string `json:"matricula"`
	Email     string `json:"email"`
	PhotoPath string `json:"photo_path,omitempty"`
}

func NewAlumn(name, matricula, email, photoPath string) *Alumn {
	return &Alumn{
		Name:      name,
		Matricula: matricula,
		Email:     email,
		PhotoPath: photoPath,
	}
}
