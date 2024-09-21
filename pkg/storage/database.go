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
	// Запрос для создания таблицы `todos`
	todoTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        status BOOLEAN NOT NULL DEFAULT false,
        priority INTEGER NOT NULL DEFAULT 0,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	// Запрос для создания таблицы `comment`
	commentTableQuery := `
    CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        todo_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        author TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (todo_id) REFERENCES todos(id) ON DELETE CASCADE
    );
    `

	// Выполнение запроса для создания таблицы `todos`
	_, err := db.Exec(todoTableQuery)
	if err != nil {
		return err
	}

	// Выполнение запроса для создания таблицы `comments`
	_, err = db.Exec(commentTableQuery)
	if err != nil {
		return err
	}

	return nil
}
