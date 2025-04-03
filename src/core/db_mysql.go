package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	errorMsg := ""

	// Cargar las variables de entorno
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Obtener las variables
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME") // Cambio de DB_SCHEMA a DB_NAME
	dbPort := os.Getenv("DB_PORT")

	// Construcción del DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		errorMsg = fmt.Sprintf("Error al abrir la base de datos: %v", err)
		return &Conn_MySQL{DB: nil, Err: errorMsg}
	}

	// Configuración del pool de conexiones
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	// Probar la conexión
	if err := db.Ping(); err != nil {
		db.Close()
		errorMsg = fmt.Sprintf("Error al verificar la conexión a la base de datos: %v", err)
		return &Conn_MySQL{DB: nil, Err: errorMsg}
	}

	fmt.Println("✅ Conexión a MySQL establecida correctamente")

	return &Conn_MySQL{DB: db, Err: errorMsg}
}

func (conn *Conn_MySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("Error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

func (conn *Conn_MySQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %w", err)
	}

	return rows, nil
}
