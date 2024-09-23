package buisness

import (
	"database/sql"
	"errors"
	"todo-app/pkg/models"
	"todo-app/pkg/storage"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db *storage.Database
}

func NewUserService(db *storage.Database) *UserService {
	return &UserService{db}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func checkPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (s *UserService) CreateUser(user *models.User) error {
	password_hash, _ := hashPassword(user.Password)

	err := s.db.QueryRow("INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3) RETURNING id, created_at",
		user.Username, user.Email, password_hash).Scan(&user.ID, &user.CreatedAt)

	return err

}

func (s *UserService) GetUser(id int) (*models.User, error) {

	var user models.User

	err := s.db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, err
}

func (s *UserService) GetUsers() ([]models.User, error) {
	rows, err := s.db.Query("SELECT id, username, email, created_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (s *UserService) UpdateUser(user *models.User) error {

	var createdAt string
	err := s.db.QueryRow("UPDATE users SET username = ?, email = ?, password_hash = ? WHERE id = ? RETURNING created_at",
		user.Username, user.Email, user.Password, user.ID).Scan(&createdAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("user not found")
		}
		return err
	}

	return nil
}

func (s *UserService) DeleteUser(id int) error {
	_, err := s.db.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
