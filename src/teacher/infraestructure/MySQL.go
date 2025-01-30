package infrastructure

import (
	"api/src/core"
	"api/src/teacher/domain/entities"
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

func (mysql *MySQL) Save(name, description string) error {
	query := "INSERT INTO teachers (name, description) VALUES (?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, name, description)
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
	query := "DELETE FROM teachers WHERE id = ?"

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

func (mysql *MySQL) ViewAll() ([]entities.Teacher, error) {
	query := "SELECT * FROM teachers"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var teachers []entities.Teacher
	for rows.Next() {
		var teacher entities.Teacher
		if err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Asignature); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		teachers = append(teachers, teacher)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return teachers, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Teacher, error) {
	query := "SELECT * FROM teachers WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	defer rows.Close()

	var teacher entities.Teacher
	if rows.Next() {
		if err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Asignature); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningun Mestro con el ID %d", id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}
	return &teacher, nil
}
