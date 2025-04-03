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
		log.Println("[ERROR] Error al configurar el pool de conexiones:", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(name, matricula string) error {
	query := "INSERT INTO alumns (name, matricula) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, matricula)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta INSERT: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("[MySQL] - Filas insertadas: %d", rowsAffected)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM alumns WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta DELETE: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún alumno con el ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Alumno eliminado con ID: %d", id)
	return nil
}

func (mysql *MySQL) ViewAll() ([]entities.Alumn, error) {
	query := "SELECT id, name, matricula FROM alumns"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

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

	if len(alumns) == 0 {
		log.Println("[INFO] No se encontraron alumnos en la base de datos.")
	}

	return alumns, nil
}

func (mysql *MySQL) ViewOne(id int) (*entities.Alumn, error) {
	query := "SELECT id, name, matricula FROM alumns WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var alumn entities.Alumn
	if rows.Next() {
		if err := rows.Scan(&alumn.ID, &alumn.Name, &alumn.Matricula); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
	} else {
		return nil, fmt.Errorf("no se encontró ningún alumno con el ID %d", id)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return &alumn, nil
}

func (mysql *MySQL) Edit(id int, name, matricula string) error {
	query := "UPDATE alumns SET name = ?, matricula = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, name, matricula, id)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta UPDATE: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún alumno con el ID %d para actualizar", id)
	}

	log.Printf("[MySQL] - Alumno actualizado con ID: %d", id)
	return nil
}
