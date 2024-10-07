package models

type Comment struct {
	ID        int    `json:"id"`
	TodoID    int    `json:"todo_id"`
	Content   string `json:"content"`
	Author    string `json:"author"`
	CreatedAt string `json:"created_at"`
}
