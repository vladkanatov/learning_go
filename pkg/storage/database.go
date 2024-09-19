package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	*sql.DB
}

func SetupDatabase() *Database {
	db, err := sql.Open("sqlite3", "todos.db")
	if err != nil {
		panic(err)
	}

	if err := InitializeDatabase(db); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	return &Database{db}
}

func InitializeDatabase(db *sql.DB) error {
	tableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        status BOOLEAN NOT NULL DEFAULT false,
        priority INTEGER NOT NULL DEFAULT 0
    );
    `

	_, err := db.Exec(tableQuery)
	if err != nil {
		return err
	}

	return nil
}
