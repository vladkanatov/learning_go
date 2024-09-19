package buisness

import (
	"database/sql"
	"errors"
	"todo-app/pkg/models"
	"todo-app/pkg/storage"

	_ "github.com/mattn/go-sqlite3"
)

type TodoService struct {
	db *storage.Database
}

func NewTodoService(db *storage.Database) *TodoService {
	return &TodoService{db}
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
	rows, err := s.db.Query("SELECT id, description, status, priority FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo
		if err := rows.Scan(&todo.ID, &todo.Description, &todo.Status, &todo.Priority); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}

func (s *TodoService) CreateTodo(todo models.Todo) error {
	_, err := s.db.Exec("INSERT INTO todos (description, status, priority) VALUES (?,?,?)",
		todo.Description, todo.Status, todo.Priority)
	return err
}

func (s *TodoService) GetTodoById(id int) (*models.Todo, error) {
	var todo models.Todo
	err := s.db.QueryRow("SELECT id, description, status, priority FROM todos WHERE id = ?", id).
		Scan(&todo.ID, &todo.Description, &todo.Status, &todo.Priority)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("task not found")
		}
		return nil, err
	}
	return &todo, nil
}

func (s *TodoService) UpdateTodo(todo models.Todo) error {
	result, err := s.db.Exec("UPDATE todos SET description = ?, status = ?, priority = ? WHERE id = ?",
		&todo.Description, &todo.Status, &todo.Priority, &todo.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (s *TodoService) DeleteTodo(id int) error {
	result, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("task not found")
	}

	return nil
}
