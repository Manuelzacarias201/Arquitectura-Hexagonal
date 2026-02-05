package main

import "api/src"

func main() {
	// inicializo la aplicacion
	src.Init()
}

// Endpoints disponibles:
// http://localhost:8080/alumns
// http://localhost:8080/alumns/1
// http://localhost:8080/teachers
// http://localhost:8080/teachers/1
// http://localhost:8080/auth/register  (POST)
// http://localhost:8080/auth/login      (POST)
// http://localhost:8080/auth/users     (GET - listar usuarios registrados)
// http://localhost:8080/auth/me        (GET - mi perfil, requiere header Authorization: Bearer <token>)
