package buisness

import (
	"database/sql"
	"errors"
	"todo-app/pkg/models"
	"todo-app/pkg/storage"
)

type CommentService struct {
	db *storage.Database
}

func NewCommentService(db *storage.Database) *CommentService {
	return &CommentService{db}
}

func (s *CommentService) GetCommentsByTodoID(todoID int) ([]models.Comment, error) {
	var comments []models.Comment
	rows, err := s.db.Query("SELECT id, todo_id, content, author, created_at FROM comments WHERE todo_id = ?", todoID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		err := rows.Scan(&comment.ID, &comment.TodoID, &comment.Content, &comment.Author, &comment.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func (s *CommentService) GetCommentByID(id int) (*models.Comment, error) {
	var comment models.Comment
	err := s.db.QueryRow("SELECT id, todo_id, content, author, created_at FROM comments WHERE id = ?", id).
		Scan(&comment.ID, &comment.TodoID, &comment.Content, &comment.Author, &comment.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("comment not found")
		}
		return nil, err
	}
	return &comment, nil
}

func (s *CommentService) CreateComment(comment *models.Comment) error {
	_, err := s.db.Query("INSERT INTO comments(todo_id, content, author) VALUES (?,?,?)", comment.TodoID, comment.Content, comment.Author)
	return err
}

func (s *CommentService) DeleteCommentByID()
