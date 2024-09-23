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
	// Запрос для создания таблицы `users`
	userTableQuery := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL UNIQUE,
        email TEXT NOT NULL UNIQUE,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    );
    `

	// Запрос для создания таблицы `categories`
	categoryTableQuery := `
    CREATE TABLE IF NOT EXISTS categories (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );
    `

	// Запрос для создания таблицы `todos` с привязкой к пользователям
	todoTableQuery := `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        description TEXT NOT NULL,
        status BOOLEAN NOT NULL DEFAULT false,
        priority INTEGER NOT NULL DEFAULT 0,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        user_id INTEGER NOT NULL,
        category_id INTEGER,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
        FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL
    );
    `

	// Запрос для создания таблицы `comments` с привязкой к пользователям
	commentTableQuery := `
    CREATE TABLE IF NOT EXISTS comments (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        todo_id INTEGER NOT NULL,
        user_id INTEGER NOT NULL,
        content TEXT NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        FOREIGN KEY (todo_id) REFERENCES todos(id) ON DELETE CASCADE,
        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
    );
    `

	// Выполнение запроса для создания таблицы `users`
	_, err := db.Exec(userTableQuery)
	if err != nil {
		return err
	}

	// Выполнение запроса для создания таблицы `categories`
	_, err = db.Exec(categoryTableQuery)
	if err != nil {
		return err
	}

	// Выполнение запроса для создания таблицы `todos`
	_, err = db.Exec(todoTableQuery)
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
