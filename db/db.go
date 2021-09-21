package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-chi/chi"
)

const (
	HOST = "localhost"
	PORT = 5433
)

// ErrNoMatch is returned when we request a row that doesn `t exist
var ErrNoMatch = fmt.Errorf("no matching record")

type Database struct {
	Conn *sql.DB
}

func Initialize(username, password, database string) (Database, error) {
	db := Database{}
	// dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	HOST, PORT, username, password, database)

	conn, err := sql.Open("postgres", "postgres://root:123456@localhost:5433/dbcrud?sslmode=disable")
	// conn, err := sql.Open("postgres", dsn)
	if err != nil {
		return db, err
	}

	db.Conn = conn
	err = db.Conn.Ping()
	if err != nil {
		return db, err
	}
	log.Println("Database connection established")
	return db, nil

}
