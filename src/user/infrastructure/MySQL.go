package infrastructure

import (
	"api/src/core"
	"api/src/user/domain/entities"
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

func (mysql *MySQL) Save(email, hashedPassword, name string) error {
	query := "INSERT INTO users (email, password, name) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, email, hashedPassword, name)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta INSERT: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	log.Printf("[MySQL] - Usuario insertado: %d filas afectadas", rowsAffected)
	return nil
}

func (mysql *MySQL) FindByEmail(email string) (*entities.User, error) {
	query := "SELECT id, email, password, name FROM users WHERE email = ?"
	rows, err := mysql.conn.FetchRows(query, email)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var user entities.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		return &user, nil
	}

	// No se encontró el usuario
	return nil, fmt.Errorf("usuario no encontrado")
}

func (mysql *MySQL) FindByID(id int) (*entities.User, error) {
	query := "SELECT id, email, password, name FROM users WHERE id = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var user entities.User
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		return &user, nil
	}

	return nil, fmt.Errorf("usuario no encontrado con el ID %d", id)
}

func (mysql *MySQL) ViewAll() ([]entities.User, error) {
	query := "SELECT id, email, password, name FROM users ORDER BY id"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()

	var users []entities.User
	for rows.Next() {
		var user entities.User
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.Name); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %w", err)
	}

	return users, nil
}
