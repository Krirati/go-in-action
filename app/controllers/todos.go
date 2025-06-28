package controllers

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Completed   *bool  `json:"completed" binding:"required"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}