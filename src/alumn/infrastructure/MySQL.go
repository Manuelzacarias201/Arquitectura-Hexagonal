package infrastructure

import (
	"api/src/alumn/domain/entities"
	"api/src/core"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(name, matricula string) error {
	query := "INSERT INTO alumns (name, matricula) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, matricula)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM alumns WHERE id = ?"

	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %w", err)
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) ViewAll() ([]entities.Alumn, error) {
	query := "SELECT * FROM alumns"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var alumns []entities.Alumn
	for rows.Next() {
		var alumn entities.Alumn
		if err := rows.Scan(&alumn.ID, &alumn.Name, &alumn.Matricula); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		alumns = append(alumns, alumn)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return alumns, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Alumn, error) {
	query := "SELECT * FROM alumns WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var alumn entities.Alumn
	if rows.Next() {
		if err := rows.Scan(&alumn.ID, &alumn.Name, &alumn.Matricula); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningun alumno con el ID %d", id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return &alumn, nil
}
