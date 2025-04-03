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
		log.Println("[ERROR] Error al configurar el pool de conexiones:", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(name, asignature string) error {
	query := "INSERT INTO teachers (name, asignature) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, asignature)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta INSERT: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("[MySQL] - Filas insertadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM teachers WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta DELETE: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún maestro con el ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Maestro eliminado con ID: %d", id)
	return nil
}

func (mysql *MySQL) ViewAll() ([]entities.Teacher, error) {
	query := "SELECT id, name, asignature FROM teachers"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

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

	if len(teachers) == 0 {
		log.Println("[INFO] No se encontraron maestros en la base de datos.")
	}

	return teachers, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Teacher, error) {
	query := "SELECT id, name, asignature FROM teachers WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var teacher entities.Teacher
	if rows.Next() {
		if err := rows.Scan(&teacher.Id, &teacher.Name, &teacher.Asignature); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningún maestro con el ID %d", id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return &teacher, nil
}

func (mysql *MySQL) Edit(id int, name, asignature string) error {
	query := "UPDATE teachers SET name = ?, asignature = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, asignature, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta UPDATE: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún maestro con el ID %d para actualizar", id)
	}

	log.Printf("[MySQL] - Maestro actualizado con ID: %d", id)
	return nil
}
