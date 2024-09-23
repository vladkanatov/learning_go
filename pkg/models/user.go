package models

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"usernames"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}
