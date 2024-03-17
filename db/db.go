package db

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

func Say() string {
	return "I'm where for db things"
}

var db *sql.DB

func Init(server, user, password, database string, port int) error {
	var err error
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return fmt.Errorf("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("error pinging database: %w", err)
	}

	return nil

}

func Query(server, user, password, database string, port int, query string) (*sql.Rows, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		fmt.Println("error opening database: %w", err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("error pinging database: %w", err)
	}
	if db == nil {
		return nil, fmt.Errorf("database connection not initialized")
	}
	return db.QueryContext(context.Background(), query)
}
