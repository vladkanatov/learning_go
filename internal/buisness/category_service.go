package buisness

import (
	"database/sql"
	"errors"
	"todo-app/pkg/models"
	"todo-app/pkg/storage"
)

type CategoryService struct {
	db *storage.Database
}

func NewCategoryService(db *storage.Database) *CategoryService {
	return &CategoryService{db}
}

func (s *CategoryService) CreateCategory(category *models.Categories) (int, error) {
	var id int

	err := s.db.QueryRow("INSERT INTO categories (name) VALUES ($1) RETURNING id",
		&category.Name).Scan(&id)

	return id, err
}

func (s *CategoryService) GetCategory(id int) (*models.Categories, error) {
	var category models.Categories
	err := s.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", id).Scan(&category.ID, &category.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("category not found")
		}
		return nil, err
	}

	return &category, err
}

func (s *CategoryService) GetCategories() ([]models.Categories, error) {
	var categories []models.Categories
	rows, err := s.db.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var category models.Categories
		if err := rows.Scan(&category.ID, &category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *CategoryService) UpdateCategory(category *models.Categories) error {
	result, err := s.db.Exec("UPDATE categories SET name = ? WHERE id = ?", &category.Name, &category.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}

func (s *CategoryService) DeleteCategory(id int) error {
	result, err := s.db.Exec("DELETE FROM categories WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("category not found")
	}

	return nil
}
