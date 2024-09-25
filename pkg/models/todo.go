package models

type Todo struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Priority    int    `json:"priority"`
	UserID      int    `json:"user_id"`
}
